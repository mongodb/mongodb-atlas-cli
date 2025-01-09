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
	"fmt"

	"github.com/google/go-github/v61/github"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/plugin"
	"github.com/spf13/cobra"
)

const (
	FirstClassSourceType = "first-class"
	sourceType           = "sourceType"
)

// uncomment example plugin to test first class plugin feature.
var firstClassPlugins = []*FirstClassPlugin{
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
		Name: "atlas-cli-plugin-kubernetes",
		Github: &Github{
			Owner: "mongodb",
			Name:  "atlas-cli-plugin-kubernetes",
		},
		Commands: []*Command{
			{
				Name:        "kubernetes",
				Description: "Root command of the Atlas CLI Kuberenetes plugin",
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
	Name     string
	Github   *Github
	Commands []*Command
}

func IsFirstClassPluginCmd(cmd *cobra.Command) bool {
	if cmdSourceType, ok := cmd.Annotations[sourceType]; ok && cmdSourceType == FirstClassSourceType {
		return true
	}
	return false
}

func (fcp *FirstClassPlugin) isAlreadyInstalled(plugins []*plugin.Plugin) bool {
	for _, p := range plugins {
		if p.Name == fcp.Name {
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
	if err := installOpts.Run(cmd.Context()); err != nil {
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
				sourceType: FirstClassSourceType,
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

func getFirstClassPluginCommands(plugins []*plugin.Plugin) []*cobra.Command {
	var commands []*cobra.Command
	// create cobra commands to install first class plugins when their commands are run
	for _, firstClassPlugin := range firstClassPlugins {
		if firstClassPlugin.isAlreadyInstalled(plugins) {
			continue
		}

		commands = append(commands, firstClassPlugin.getCommands()...)
	}

	return commands
}
