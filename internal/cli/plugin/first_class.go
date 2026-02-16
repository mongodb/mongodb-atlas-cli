// Copyright 2025 MongoDB Inc
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
	"fmt"

	"github.com/Masterminds/semver/v3"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/plugin"
	"github.com/spf13/cobra"
)

const (
	FirstClassSourceType = "first-class"
	sourceType           = "sourceType"
)

// uncomment example plugin to test first class plugin feature.
var FirstClassPlugins = []*FirstClassPlugin{
	// {
	// 	Name: "atlas-cli-first-class-plugin-example",
	// 	Github: &Github{
	// 		Owner: "stefan4h",
	// 		Name: "atlas-cli-first-class-plugin-example",
	// 	},
	// 	Commands: []*Command{
	// 		{
	// 			Name: "first-class",
	// 			Description: "Root command of the Atlas CLI first class plugin example",
	// 		},
	// 	},
	// },
	{
		Name:       "atlas-cli-plugin-kubernetes",
		MinVersion: "1.2.4",
		Github: &Github{
			Owner: "mongodb",
			Name:  "atlas-cli-plugin-kubernetes",
		},
		Commands: []*Command{
			{
				Name:        "kubernetes",
				Description: "Manage Kubernetes resources.",
			},
		},
	},
	{
		Name:       "atlas-local-plugin",
		MinVersion: "0.0.10",
		Github: &Github{
			Owner: "mongodb",
			Name:  "atlas-local-cli",
		},
		Commands: []*Command{
			{
				Name:        "local",
				Description: "Manage MongoDB Atlas local instances",
			},
		},
	},
}

type Command struct {
	Name        string
	Description string
}

type Github struct {
	Owner string
	Name  string
}

type FirstClassPlugin struct {
	Name       string
	MinVersion string
	Github     *Github
	Commands   []*Command
}

func IsFirstClassPluginCmd(cmd *cobra.Command) bool {
	if cmdSourceType, ok := cmd.Annotations[sourceType]; ok && cmdSourceType == FirstClassSourceType {
		return true
	}
	return false
}

func (fcp *FirstClassPlugin) getInstalledPlugin(plugins *plugin.ValidatedPlugins) *plugin.Plugin {
	for _, p := range plugins.GetValidPlugins() {
		if p.Name == fcp.Name {
			return p
		}
	}

	return nil
}

func (fcp *FirstClassPlugin) needsUpdate(installedPlugin *plugin.Plugin) bool {
	if fcp.MinVersion == "" {
		return false
	}

	minVersion, err := semver.NewVersion(fcp.MinVersion)
	if err != nil {
		return false
	}

	return installedPlugin.Version.LessThan(minVersion)
}

func (fcp *FirstClassPlugin) installPlugin(cmd *cobra.Command, plugins *plugin.ValidatedPlugins) error {
	installOpts := &InstallOpts{
		Opts: Opts{
			plugins: plugins,
		},
		ghClient: NewAuthenticatedGithubClient(),
	}
	installOpts.githubAsset = &GithubAsset{
		owner: fcp.Github.Owner,
		name:  fcp.Github.Name,
	}
	installOpts.Print("Installing first class plugin " + fcp.Name)

	// check if plugin already exists, if not, install it
	if err := installOpts.checkForDuplicatePlugins(); err != nil {
		return fmt.Errorf("first class plugin %s is already installed, should not install again", fcp.Name)
	}
	if err := installOpts.Run(cmd.Context()); err != nil {
		return fmt.Errorf("failed to install first class plugin %s: %w", fcp.Name, err)
	}

	return nil
}

func (fcp *FirstClassPlugin) updatePlugin(cmd *cobra.Command, plugins *plugin.ValidatedPlugins, existingPlugin *plugin.Plugin) error {
	updateOpts := &UpdateOpts{
		Opts: Opts{
			plugins: plugins,
		},
		ghClient: NewAuthenticatedGithubClient(),
	}

	githubAsset := &GithubAsset{
		owner: fcp.Github.Owner,
		name:  fcp.Github.Name,
	}

	updateOpts.Print(fmt.Sprintf("Updating first class plugin %s to minimum required version %s", fcp.Name, fcp.MinVersion))

	if err := updateOpts.updatePlugin(cmd.Context(), githubAsset, existingPlugin); err != nil {
		return fmt.Errorf("failed to update first class plugin %s: %w", fcp.Name, err)
	}

	return nil
}

func (fcp *FirstClassPlugin) runFirstClassPluginCommand(cmd *cobra.Command, args []string, plugins *plugin.ValidatedPlugins) error {
	// Check if plugin is already installed
	installedPlugin := fcp.getInstalledPlugin(plugins)

	if installedPlugin == nil {
		// Plugin not installed, install it
		if err := fcp.installPlugin(cmd, plugins); err != nil {
			return err
		}
	} else if fcp.needsUpdate(installedPlugin) {
		// Plugin installed but below minimum version, update it
		if err := fcp.updatePlugin(cmd, plugins, installedPlugin); err != nil {
			return err
		}
	}

	// find and run installed plugin
	installedPlugin, err := plugin.GetPluginWithName(fcp.Name, nil, false)
	if err != nil {
		return err
	}
	// validate plugin version is compatible
	if err := plugin.ValidateVersion(*installedPlugin.Github, installedPlugin.Version); err != nil {
		return err
	}

	return installedPlugin.Run(cmd, args)
}

func (fcp *FirstClassPlugin) getCommands(plugins *plugin.ValidatedPlugins) []*cobra.Command {
	commands := make([]*cobra.Command, 0, len(fcp.Commands))

	// for every command listed in the first class plugin, create a cobra command that installs the plugin
	for _, firstClassPluginCommand := range fcp.Commands {
		cmd := &cobra.Command{
			Use:   firstClassPluginCommand.Name,
			Short: firstClassPluginCommand.Description,
			Annotations: map[string]string{
				sourceType: FirstClassSourceType,
			},
			RunE: func(cmd *cobra.Command, args []string) error {
				return fcp.runFirstClassPluginCommand(cmd, args, plugins)
			},
			DisableFlagParsing: true,
		}

		commands = append(commands, cmd)
	}

	return commands
}

func getFirstClassPluginCommands(plugins *plugin.ValidatedPlugins) []*cobra.Command {
	var commands []*cobra.Command
	// create cobra commands to install/update first class plugins when their commands are run
	for _, firstClassPlugin := range FirstClassPlugins {
		installedPlugin := firstClassPlugin.getInstalledPlugin(plugins)

		// Skip if plugin is already installed and doesn't need updating
		if installedPlugin != nil && !firstClassPlugin.needsUpdate(installedPlugin) {
			continue
		}

		commands = append(commands, firstClassPlugin.getCommands(plugins)...)
	}

	return commands
}

// getFirstClassPluginsNeedingUpdate returns a set of plugin names that are first-class plugins
// which are installed but need updating to meet the minimum version requirement.
func getFirstClassPluginsNeedingUpdate(plugins *plugin.ValidatedPlugins) map[string]bool {
	result := make(map[string]bool)

	for _, firstClassPlugin := range FirstClassPlugins {
		installedPlugin := firstClassPlugin.getInstalledPlugin(plugins)
		if installedPlugin != nil && firstClassPlugin.needsUpdate(installedPlugin) {
			result[installedPlugin.Name] = true
		}
	}

	return result
}
