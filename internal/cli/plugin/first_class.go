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
	_ "embed"
	"encoding/json"
	"fmt"

	"github.com/google/go-github/v61/github"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/plugin"
	"github.com/spf13/cobra"
)

const (
	SourceType = "first-class"
)

//go:embed first-class-plugins.json
var firstClassPluginsFile []byte

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

func IsFirstClassPlugin(cmd *cobra.Command) bool {
	if cmdSourceType, ok := cmd.Annotations["sourceType"]; ok && cmdSourceType == SourceType {
		return true
	}
	return false
}

func (fcp *FirstClassPlugin) isAlreadyInstalled(existingCommands []*cobra.Command) bool {
	for _, cmd := range existingCommands {
		if !plugin.IsPluginCmd(cmd) {
			continue
		}
		if pluginName, ok := cmd.Annotations["sourcePluginName"]; ok && pluginName == fcp.Name {
			return true
		}
	}

	return false
}

func (fcp *FirstClassPlugin) runFirstClassPluginCommand(cmd *cobra.Command, args []string, ghClient *github.Client) error {
	installOpts := &InstallOpts{}
	installOpts.githubAsset = &GithubAsset{
		ghClient: ghClient,
		owner:    fcp.Github.Owner,
		name:     fcp.Github.Name,
	}
	installOpts.Print("Installing first class plugin " + fcp.Name)

	// check if plugin already exists, if not, install it
	if err := installOpts.checkForDuplicatePlugins(); err != nil {
		return fmt.Errorf("first class plugin %s is already installed, should not install again", fcp.Name)
	}
	if err := installOpts.Run(); err != nil {
		return fmt.Errorf("failed to install first class plugin %s: %w", fcp.Name, err)
	}

	// find and run installed plugin
	installedPlugin, err := plugin.GetPluginWithName(fcp.Name, nil, false)
	if err != nil {
		return err
	}

	return installedPlugin.Run(cmd, args)
}

func (fcp *FirstClassPlugin) getCommands() []*cobra.Command {
	commands := make([]*cobra.Command, 0, len(fcp.Commands))
	ghClient := github.NewClient(nil)

	// for every command listed in the first class plugin, create a cobra command that installs the plugin
	for _, firstClassPluginCommand := range fcp.Commands {
		cmd := &cobra.Command{
			Use:   firstClassPluginCommand.Name,
			Short: firstClassPluginCommand.Description,
			Annotations: map[string]string{
				"sourceType": SourceType,
			},
			RunE: func(cmd *cobra.Command, args []string) error {
				return fcp.runFirstClassPluginCommand(cmd, args, ghClient)
			},
			DisableFlagParsing: true,
		}

		commands = append(commands, cmd)
	}

	return commands
}

// there are no first class plugins at the moment which is why we can only test this function with
// an example plugin. To test this copy content from the file "example.first-class-plugins.json" into "first-class-plugins.json".
func RegisterFirstClassPluginCommands(rootCmd *cobra.Command) error {
	// Read the file contents into a byte slice
	var firstClassPlugins []*FirstClassPlugin
	if err := json.Unmarshal(firstClassPluginsFile, &firstClassPlugins); err != nil {
		return fmt.Errorf("failed to unmarshal first class plugin file: %w", err)
	}

	// create cobra commands to install first class plugins when their commands are run
	for _, firstClassPlugin := range firstClassPlugins {
		if firstClassPlugin.isAlreadyInstalled(rootCmd.Commands()) {
			continue
		}

		rootCmd.AddCommand(firstClassPlugin.getCommands()...)
	}

	return nil
}
