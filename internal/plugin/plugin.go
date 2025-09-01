// Copyright 2024 MongoDB Inc
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package plugin

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"path"

	"github.com/Masterminds/semver/v3"
	"github.com/mongodb/atlas-cli-core/config"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/log"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/set"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/telemetry"
	"github.com/spf13/cobra"
)

var (
	errCreateDefaultPluginDir = errors.New("failed to create default plugin directory")
	minimumPluginVersions     = map[Github]string{
		{Owner: "mongodb", Name: "atlas-cli-plugin-kubernetes"}: "v1.1.7",
		{Owner: "mongodb", Name: "atlas-cli-plugin-gsa"}:        "v0.0.2", // TODO: ensure this version is correct version after work in CLOUDP-333246
	}
)

const (
	PluginSourceType = "plugin"
	sourceType       = "sourceType"
	sourcePluginName = "sourcePluginName"
)

func IsPluginCmd(cmd *cobra.Command) bool {
	if cmdSourceType, ok := cmd.Annotations[sourceType]; ok && cmdSourceType == PluginSourceType {
		return true
	}
	return false
}

func GetPluginWithName(name string, existingCommandsSet set.Set[string], onlySearchValidPlugins bool) (*Plugin, error) {
	var plugins []*Plugin
	if onlySearchValidPlugins {
		plugins = GetAllPluginsValidated(existingCommandsSet).GetValidPlugins()
	} else {
		plugins = getAllPlugins()
	}

	for _, plugin := range plugins {
		if plugin.Name == name {
			return plugin, nil
		}
	}

	return nil, fmt.Errorf("could not find plugin %s", name)
}

func getAllPlugins() []*Plugin {
	// Load manifests from plugin directories
	manifests := loadManifestsFromPluginDirectories()

	// Convert manifests to plugins
	plugins := convertManifestsToPlugins(manifests)

	return plugins
}

type ValidatedPlugins struct {
	ValidPlugins                     []*Plugin
	PluginsWithDuplicateManifestName []*Plugin
	PluginsWithDuplicateCommands     []*Plugin
}

func (v *ValidatedPlugins) GetValidPlugins() []*Plugin {
	return v.ValidPlugins
}

func (v *ValidatedPlugins) GetValidAndInvalidPlugins() []*Plugin {
	return append(append(v.ValidPlugins, v.PluginsWithDuplicateManifestName...), v.PluginsWithDuplicateCommands...)
}

func GetAllPluginsValidated(existingCommandsSet set.Set[string]) *ValidatedPlugins {
	// Load manifests from plugin directories
	manifests := loadManifestsFromPluginDirectories()
	duplicateManifests := make([]*Manifest, 0)
	duplicateCommands := make([]*Manifest, 0)

	// Remove manifests with duplicate names
	manifests, duplicateManifestNames := removeManifestsWithDuplicateNames(manifests)
	for _, duplicate := range duplicateManifestNames {
		duplicateManifests = append(duplicateManifests, duplicate.Manifest)
		logPluginWarning(`could not load plugin "%s" because there are multiple plugins with that name`, duplicate.DuplicateName)
	}

	// Remove manifests that contain already existing commands
	manifests, duplicateManifest := getUniqueManifests(manifests, existingCommandsSet)
	for _, manifest := range duplicateManifest {
		duplicateCommands = append(duplicateCommands, manifest)
		logPluginWarning(`could not load plugin "%s" because it contains a command that already exists in the AtlasCLI or another plugin`, manifest.Name)
	}

	// Convert manifests to validated plugins
	return &ValidatedPlugins{
		ValidPlugins:                     convertManifestsToPlugins(manifests),
		PluginsWithDuplicateManifestName: convertManifestsToPlugins(duplicateManifests),
		PluginsWithDuplicateCommands:     convertManifestsToPlugins(duplicateCommands),
	}
}

func GetDefaultPluginDirectory() (string, error) {
	configHome, err := config.CLIConfigHome()

	if err != nil {
		return "", fmt.Errorf("failed to retrieve CLI config home: %w", err)
	}

	pluginDirectoryPath := path.Join(configHome, "plugins")

	err = os.MkdirAll(pluginDirectoryPath, os.ModePerm)
	if err != nil {
		return "", errCreateDefaultPluginDir
	}

	return pluginDirectoryPath, nil
}

type Command struct {
	Name        string
	Description string
	Aliases     []string
}

type Github struct {
	Owner string
	Name  string
}

func (g *Github) Equals(owner string, name string) bool {
	if g.Owner == owner && g.Name == name {
		return true
	}

	return false
}

type Plugin struct {
	Name                string
	Description         string
	PluginDirectoryPath string
	BinaryName          string
	Version             *semver.Version
	Commands            []*Command
	Github              *Github
}

func (p *Plugin) Run(cmd *cobra.Command, args []string) error {
	// TODO:  uncomment after plugin release
	// if err := ValidateVersion(*p.Github, p.Version); err != nil {
	// 	return err
	// }

	p.setTelemetry()

	binaryPath := path.Join(p.PluginDirectoryPath, p.BinaryName)
	// suppressing lint error flagging potential tainted input or cmd arguments
	// we are this can happen, it is by design
	// #nosec G204
	execCmd := exec.Command(binaryPath, append([]string{cmd.Use}, args...)...)
	execCmd.Stdin = cmd.InOrStdin()
	execCmd.Stdout = cmd.OutOrStdout()
	execCmd.Stderr = cmd.OutOrStderr()
	execCmd.Env = os.Environ()
	if err := execCmd.Run(); err != nil {
		var exitErr *exec.ExitError
		if errors.As(err, &exitErr) {
			cmd.SilenceErrors = true
			_, _ = log.Debugf("Silenced error: %v", exitErr)
		}
		return err
	}
	return nil
}

func (p *Plugin) Uninstall() error {
	return os.RemoveAll(p.PluginDirectoryPath)
}

func (p *Plugin) HasGithub() bool {
	return p.Github != nil && p.Github.Name != "" && p.Github.Owner != ""
}

func (p *Plugin) GetCobraCommands() []*cobra.Command {
	commands := make([]*cobra.Command, 0, len(p.Commands))

	for _, pluginCmd := range p.Commands {
		command := &cobra.Command{
			Use:   pluginCmd.Name,
			Short: pluginCmd.Description,
			Annotations: map[string]string{
				sourceType:       PluginSourceType,
				sourcePluginName: p.Name,
			},
			RunE:    p.Run,
			Aliases: pluginCmd.Aliases,
		}

		// Disable the default cobra help function.
		// Instead redirect help to the plugin.
		// Example: atlas example-plugin --help -> [example-binary] example-plugin --help
		command.SetHelpFunc(func(cmd *cobra.Command, args []string) {
			// args contains all arguments + the name of the command
			// we don't need the name of the subcommand
			if err := p.Run(cmd, args[1:]); err != nil {
				_, _ = log.Warningf("failed to generate help for plugin command '%v': %v", args[0], err)
			}
		})

		command.DisableFlagParsing = true

		commands = append(commands, command)
	}

	return commands
}

func convertManifestsToPlugins(manifests []*Manifest) []*Plugin {
	plugins := make([]*Plugin, 0, len(manifests))
	for _, manifest := range manifests {
		plugin, err := createPluginFromManifest(manifest)
		if err != nil {
			logPluginWarning(err.Error())
			continue
		}
		plugins = append(plugins, plugin)
	}

	return plugins
}

func createPluginFromManifest(manifest *Manifest) (*Plugin, error) {
	version, err := semver.NewVersion(manifest.Version)
	if err != nil {
		return nil, fmt.Errorf("invalid version in manifest file %s", manifest.Name)
	}

	plugin := Plugin{
		Name:                manifest.Name,
		Description:         manifest.Description,
		PluginDirectoryPath: manifest.PluginDirectoryPath,
		BinaryName:          manifest.Binary,
		Version:             version,
		Commands:            make([]*Command, 0, len(manifest.Commands)),
	}

	if manifest.Github != nil {
		plugin.Github = &Github{
			Owner: manifest.Github.Owner,
			Name:  manifest.Github.Name,
		}
	}

	for cmdName, value := range manifest.Commands {
		plugin.Commands = append(plugin.Commands, &Command{Name: cmdName, Description: value.Description, Aliases: value.Aliases})
	}

	return &plugin, nil
}

func logPluginWarning(message string, args ...any) {
	_, _ = log.Warningf(fmt.Sprintf("-- plugin warning: %s\n", message), args...)
}

func (p *Plugin) setTelemetry() {
	info := telemetry.PluginExecutionInfo{
		Version: p.Version,
	}

	if p.Github != nil {
		info.GithubOwner = &p.Github.Owner
		info.GithubRepository = &p.Github.Name
	}

	telemetry.AppendOption(telemetry.WithPluginExecutionInfo(info))
}

// ValidateVersion validates the version of a plugin against the minimum required version.
// If a plugin is not listed in the minimumPluginVersions map, it is considered valid.
func ValidateVersion(gh Github, version *semver.Version) error {
	minVersionStr, exists := minimumPluginVersions[gh]
	if !exists {
		return nil // No version requirement for this plugin
	}

	// TODO check if version is optional, if so, check if version is available and return nil
	minVersion, err := semver.NewVersion(minVersionStr)
	if err != nil {
		return err
	}

	if version.LessThan(minVersion) {
		return fmt.Errorf("plugin %s/%s version v%s is below minimum required version %s for this version of AtlasCLI.\nPlease update the plugin using 'atlas plugin update %s'",
			gh.Owner, gh.Name, version.String(), minVersionStr, gh.Name)
	}

	return nil
}
