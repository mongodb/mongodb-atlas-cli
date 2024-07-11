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
	"os"
	"os/exec"

	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/log"
	"github.com/spf13/cobra"
)

func GetAllValidPlugins(existingCommands []*cobra.Command) []*Plugin {
	var manifests []*Manifest

	// Load manifests from plugin directories
	if loadedManifests, err := getManifestsFromPluginDirectory("./plugins"); err != nil {
		logPluginWarning(`could not load manifests from directory "./plugins" because of error: %s`, err.Error())
	} else {
		manifests = append(manifests, loadedManifests...)
	}

	extraPluginDir := os.Getenv("ATLAS_CLI_EXTRA_PLUGIN_DIRECTORY")

	if extraPluginDir != "" {
		if loadedManifests, err := getManifestsFromPluginDirectory(extraPluginDir); err != nil {
			logPluginWarning(`could not load plugins from folder "%s" provided in environment variable ATLAS_CLI_EXTRA_PLUGIN_DIRECTORY: %s`, extraPluginDir, err.Error())
		} else {
			manifests = append(manifests, loadedManifests...)
		}
	}

	// Remove manifests that contain already existing commands
	manifests, duplicateManifest := getUniqueManifests(manifests, existingCommands)

	for _, manifest := range duplicateManifest {
		logPluginWarning(`could not load plugin "%s" because it contains a command that already exists in the AtlasCLI or another plugin`, manifest.Name)
	}

	// Convert manifests to plugins
	plugins := make([]*Plugin, 0, len(manifests))
	for _, manifest := range manifests {
		plugins = append(plugins, createPluginFromManifest(manifest))
	}

	return plugins
}

type Plugin struct {
	Name        string
	Description string
	BinaryPath  string
	Version     string
	Commands    []*Command
}

type Command struct {
	Name        string
	Description string
}

func (p *Plugin) Run(cmd *cobra.Command, args []string) error {
	// suppressing lint error flagging potential tainted input or cmd arguments
	// we are this can happen, it is by design
	// #nosec G204
	execCmd := exec.Command(p.BinaryPath, args...)
	execCmd.Stdout = cmd.OutOrStdout()
	execCmd.Stderr = cmd.OutOrStderr()
	return execCmd.Run()
}

func (p *Plugin) GetCobraCommands() []*cobra.Command {
	commands := make([]*cobra.Command, 0, len(p.Commands))

	for _, command := range p.Commands {
		command := &cobra.Command{
			Use:   command.Name,
			Short: command.Description,
			RunE:  p.Run,
		}

		commands = append(commands, command)
	}

	return commands
}

func createPluginFromManifest(manifest *Manifest) *Plugin {
	plugin := Plugin{
		Name:        manifest.Name,
		Description: manifest.Description,
		BinaryPath:  manifest.BinaryPath,
		Version:     manifest.Version,
		Commands:    make([]*Command, 0, len(manifest.Commands)),
	}

	for cmdName, value := range manifest.Commands {
		plugin.Commands = append(plugin.Commands, &Command{Name: cmdName, Description: value.Description})
	}

	return &plugin
}

func logPluginWarning(message string, args ...any) {
	_, _ = log.Warningf(fmt.Sprintf("-- plugin warning: %s\n", message), args...)
}
