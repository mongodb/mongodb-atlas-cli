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
	"strings"

	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/log"
	"github.com/spf13/cobra"
)

func GetAllValidPluginCommands(existingCommands []*cobra.Command) []*cobra.Command {
	existingCommandsMap := make(map[string]bool)

	for _, cmd := range existingCommands {
		existingCommandsMap[cmd.Name()] = true
	}

	var pluginCommands []*cobra.Command

	if pluginsWithCommands, err := getPluginCommandsFromDirectory("./plugins"); err != nil {
		_, _ = log.Warningf("Could not load plugins from directory ./plugins because of error: %s", err.Error())
	} else {
		commands := filterUniqueCommands(pluginsWithCommands, existingCommandsMap)
		pluginCommands = append(pluginCommands, commands...)
	}

	extraPluginDir := os.Getenv("ATLAS_CLI_EXTRA_PLUGIN_DIRECTORY")

	if extraPluginDir != "" {
		if pluginsWithCommands, err := getPluginCommandsFromDirectory(extraPluginDir); err != nil {
			_, _ = log.Warningf("Could not load plugins from folder %s provided in environment variable ATLAS_CLI_EXTRA_PLUGIN_DIRECTORY: %s", extraPluginDir, err.Error())
		} else {
			commands := filterUniqueCommands(pluginsWithCommands, existingCommandsMap)
			pluginCommands = append(pluginCommands, commands...)
		}
	}

	return pluginCommands
}

func filterUniqueCommands(pluginsWithCommands map[*Manifest][]*cobra.Command, existingCommandsMap map[string]bool) []*cobra.Command {
	var filteredPlugins []*cobra.Command

	for pluginManifest, commands := range pluginsWithCommands {
		if hasDuplicateCommand(commands, existingCommandsMap) {
			_, _ = log.Warningf("Could not load plugin %s because it contains a command that already exists in the AtlasCLI or another plugin", pluginManifest.Name)
			continue
		}
		for _, cmd := range commands {
			existingCommandsMap[cmd.Name()] = true
		}
		filteredPlugins = append(filteredPlugins, commands...)
	}

	return filteredPlugins
}

func hasDuplicateCommand(commands []*cobra.Command, existingCommandsMap map[string]bool) bool {
	for _, cmd := range commands {
		if existingCommandsMap[cmd.Name()] {
			return true
		}
	}
	return false
}

func getPluginCommandsFromDirectory(pluginDir string) (map[*Manifest][]*cobra.Command, error) {
	files, err := os.ReadDir(pluginDir)

	if err != nil {
		return nil, err
	}

	var warningLog strings.Builder

	pluginsWithCommands := make(map[*Manifest][]*cobra.Command)
	for _, directory := range files {
		if !directory.IsDir() {
			continue
		}

		pluginDirectoryPath := fmt.Sprintf("%s/%s", pluginDir, directory.Name())

		manifestFileData, err := getManifestFileBytes(pluginDirectoryPath)

		if err != nil {
			continue
		}

		pluginManifest, err := parseManifestFile(manifestFileData)

		if err != nil {
			warningLog.WriteString("\n-- plugin invalid: manifest file could not be parsed\n")
			continue
		}

		if valid, errors := pluginManifest.IsValid(); !valid {
			warningLog.WriteString(fmt.Sprintf("\n-- plugin invalid: plugin in directory %s could not be loaded due to the following error(s) in the manifest.yaml:\n", pluginDirectoryPath))
			for _, err := range errors {
				warningLog.WriteString(fmt.Sprintf("\t- %s\n", err.Error()))
			}
			warningLog.WriteString("\n")
			continue
		}

		binaryPath, err := getPathToExecutableBinary(pluginDirectoryPath, pluginManifest.Binary)

		if err != nil {
			warningLog.WriteString(fmt.Sprintf("\n-- plugin invalid: %s\n", err.Error()))
		}
		pluginsWithCommands[pluginManifest] = createCommandsFromManifest(*pluginManifest, binaryPath)
	}

	_, _ = log.Warning(warningLog.String())

	return pluginsWithCommands, nil
}

func createCommandsFromManifest(pluginManifest Manifest, binaryPath string) []*cobra.Command {
	commands := make([]*cobra.Command, 0, len(pluginManifest.Commands))

	for cmdName, value := range pluginManifest.Commands {
		command := &cobra.Command{
			Use:   cmdName,
			Short: value.Description,
			RunE: func(cmd *cobra.Command, args []string) error {
				execCmd := exec.Command(binaryPath, args...)
				execCmd.Stdout = cmd.OutOrStdout()
				execCmd.Stderr = cmd.OutOrStderr()
				return execCmd.Run()
			},
		}

		commands = append(commands, command)
	}

	return commands
}

func getPathToExecutableBinary(pluginDirectoryPath string, binaryName string) (string, error) {
	binaryPath := fmt.Sprintf("%s/%s", pluginDirectoryPath, binaryName)

	binaryFileInfo, err := os.Stat(binaryPath)

	if err != nil {
		return "", fmt.Errorf("binary %s does not exists", binaryPath)
	}

	// makes sure that the binary file is made executable if it is not already
	binaryFileMode := binaryFileInfo.Mode()
	const executablePermissions = 0o111

	if binaryFileMode & executablePermissions != 0 {
		return binaryPath, nil
	}

	if err := os.Chmod(binaryPath, binaryFileMode | executablePermissions); err != nil {
		return "", err
	}

	return binaryPath, nil
}
