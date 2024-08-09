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
	"encoding/json"
	"fmt"
	"os"

	"github.com/google/go-github/v61/github"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/plugin"
	"github.com/spf13/cobra"
)

const (
	firstClassPluginsFilePath = "first-class-plugins.json"
)

type Command struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

type Github struct {
	Owner string `json:"owner"`
	Name  string `json:"name"`
}

type FirstClassPlugin struct {
	Name     string    `json:"name"`
	Github   Github    `json:"github"`
	Commands []Command `json:"commands"`
}

func firstClassPluginCmdAlreadyExists(firstClassPlugin *FirstClassPlugin, existingCommandsMap map[string]bool) bool {
	for _, firstClassPluginCommand := range firstClassPlugin.Commands {
		if existingCommandsMap[firstClassPluginCommand.Name] {
			return true
		}
	}

	return false
}

func runFirstClassPluginCommand(cmd *cobra.Command, args []string, ghClient *github.Client, firstClassPlugin *FirstClassPlugin, existingCommandsMap map[string]bool) error {
	installOpts := &InstallOpts{}
	installOpts.githubAsset = &GithubAsset{
		ghClient: ghClient,
		owner:    firstClassPlugin.Github.Owner,
		name:     firstClassPlugin.Github.Name,
	}
	installOpts.Print("Installing first class plugin " + firstClassPlugin.Name)

	// check if plugin already exists, if not, install it
	if err := installOpts.checkForDuplicatePlugins(); err != nil {
		return fmt.Errorf("first class plugin %s is already installed, should not install again", firstClassPlugin.Name)
	}
	if err := installOpts.Run(); err != nil {
		return fmt.Errorf("failed to install first class plugin %s: %w", firstClassPlugin.Name, err)
	}

	// remove first class plugin commands from existingCommandsMap so that plugin is considered valid and can be discovered
	for _, firstfirstClassPluginCommand := range firstClassPlugin.Commands {
		delete(existingCommandsMap, firstfirstClassPluginCommand.Name)
	}

	// find and run installed plugin
	installedPlugin, err := plugin.GetPluginWithName(firstClassPlugin.Name, existingCommandsMap)
	if err != nil {
		return err
	}

	return installedPlugin.Run(cmd, args)
}

func getCommandsFromFirstClassPlugin(firstClassPlugin *FirstClassPlugin, existingCommandsMap map[string]bool) []*cobra.Command {
	commands := make([]*cobra.Command, 0, len(firstClassPlugin.Commands))
	ghClient := github.NewClient(nil)

	// for every command listed in the first class plugin, create a cobra command that installs the plugin
	for _, firstClassPluginCommand := range firstClassPlugin.Commands {
		cmd := &cobra.Command{
			Use:   firstClassPluginCommand.Name,
			Short: firstClassPluginCommand.Description,
			RunE: func(cmd *cobra.Command, args []string) error {
				copiedExistingCommandsMap := make(map[string]bool)
				for name, exists := range existingCommandsMap {
					copiedExistingCommandsMap[name] = exists
				}

				return runFirstClassPluginCommand(cmd, args, ghClient, firstClassPlugin, copiedExistingCommandsMap)
			},
		}

		commands = append(commands, cmd)
	}

	return commands
}

func RegisterFirstClassPluginCommands(rootCmd *cobra.Command) error {
	// Read first class plugin json file
	firstClassPluginsFile, err := os.ReadFile(firstClassPluginsFilePath)
	if err != nil {
		return fmt.Errorf("failed to read first class plugin file: %w", err)
	}

	// Read the file contents into a byte slice
	var firstClassPlugins []*FirstClassPlugin
	if err = json.Unmarshal(firstClassPluginsFile, &firstClassPlugins); err != nil {
		return fmt.Errorf("failed to unmarshal first class plugin file: %w", err)
	}

	// create cobra commands to install first class plugins when their commands are run
	existingCommandsMap := makeExistingCommandsMap(rootCmd.Commands())
	for _, firstClassPlugin := range firstClassPlugins {
		if firstClassPluginCmdAlreadyExists(firstClassPlugin, existingCommandsMap) {
			continue
		}

		rootCmd.AddCommand(getCommandsFromFirstClassPlugin(firstClassPlugin, existingCommandsMap)...)
	}

	return nil
}
