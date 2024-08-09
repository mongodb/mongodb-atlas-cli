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
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/config"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/log"
	"github.com/spf13/cobra"
)

var (
	errCreateDefaultPluginDir = errors.New("failed to create default plugin directory")
)

const (
	ExtraPluginDirectoryEnvKey = "ATLAS_CLI_EXTRA_PLUGIN_DIRECTORY"
	SourceType                 = "plugin"
)

func IsPluginCmd(cmd *cobra.Command) bool {
	if cmdSourceType, ok := cmd.Annotations["sourceType"]; ok && cmdSourceType == SourceType {
		return true
	}
	return false
}

func GetPluginWithName(name string, existingCommandsMap map[string]bool) (*Plugin, error) {
	plugins := GetAllValidPlugins(existingCommandsMap)

	for _, plugin := range plugins {
		if plugin.Name == name {
			return plugin, nil
		}
	}

	return nil, fmt.Errorf("could not find plugin %s", name)
}

func GetAllValidPlugins(existingCommandsMap map[string]bool) []*Plugin {
	var manifests []*Manifest

	// Load manifests from plugin directories
	if defaultPluginDirectory, err := GetDefaultPluginDirectory(); err == nil {
		if loadedManifests, err := getManifestsFromPluginsDirectory(defaultPluginDirectory); err != nil {
			logPluginWarning(`could not load manifests from directory "%s" because of error: %s`, defaultPluginDirectory, err.Error())
		} else {
			manifests = append(manifests, loadedManifests...)
		}
	}

	if extraPluginDir := os.Getenv("ATLAS_CLI_EXTRA_PLUGIN_DIRECTORY"); extraPluginDir != "" {
		if loadedManifests, err := getManifestsFromPluginsDirectory(extraPluginDir); err != nil {
			logPluginWarning(`could not load plugins from folder "%s" provided in environment variable ATLAS_CLI_EXTRA_PLUGIN_DIRECTORY: %s`, extraPluginDir, err.Error())
		} else {
			manifests = append(manifests, loadedManifests...)
		}
	}

	// Remove manifests with duplicate names
	manifests, duplicateManifestNames := removeManifestsWithDuplicateNames(manifests)
	for _, name := range duplicateManifestNames {
		logPluginWarning(`could not load plugin "%s" because there are multiple plugins with that name`, name)
	}

	// Remove manifests that contain already existing commands
	manifests, duplicateManifest := getUniqueManifests(manifests, existingCommandsMap)
	for _, manifest := range duplicateManifest {
		logPluginWarning(`could not load plugin "%s" because it contains a command that already exists in the AtlasCLI or another plugin`, manifest.Name)
	}

	// Convert manifests to plugins
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
	binaryPath := path.Join(p.PluginDirectoryPath, p.BinaryName)
	// suppressing lint error flagging potential tainted input or cmd arguments
	// we are this can happen, it is by design
	// #nosec G204
	execCmd := exec.Command(binaryPath, args...)
	execCmd.Stdin = cmd.InOrStdin()
	execCmd.Stdout = cmd.OutOrStdout()
	execCmd.Stderr = cmd.OutOrStderr()
	execCmd.Env = os.Environ()
	return execCmd.Run()
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
				"sourceType":       SourceType,
				"sourcePluginName": p.Name,
			},
			RunE: p.Run,
		}

		commands = append(commands, command)
	}

	return commands
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
		plugin.Commands = append(plugin.Commands, &Command{Name: cmdName, Description: value.Description})
	}

	return &plugin, nil
}

func logPluginWarning(message string, args ...any) {
	_, _ = log.Warningf(fmt.Sprintf("-- plugin warning: %s\n", message), args...)
}
