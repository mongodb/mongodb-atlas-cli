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

func GetAllValidPluginCommands(existingCommands []*cobra.Command) ([]*cobra.Command) {
	
}

func getPluginCommandsFromDirectory(pluginDir string) (map[*PluginManifest][]*cobra.Command, error) {
	files, err := os.ReadDir(pluginDir)

	if err != nil {
		return nil, err
	}

	pluginsWithCommands := make(map[*PluginManifest][]*cobra.Command)
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
			log.Warningf("\n-- plugin invalid: manifest file could not be parsed\n")
			continue
		}

		if valid, errors := pluginManifest.IsValid(); !valid {
			log.Warningf("\n-- plugin invalid: plugin in directory %s could not be loaded due to the following error(s) in the manifest.yaml:\n", pluginDirectoryPath)
			for _, err := range errors {
				log.Warningf("\t- %s\n", err.Error())
			}
			log.Warning("\n")
			continue
		}

		binaryPath, err := getPathToExecutableBinary(pluginDirectoryPath, pluginManifest.Binary)

		if err != nil {
			log.Warningf("\n-- plugin invalid: %s\n", err.Error())
		}
		pluginsWithCommands[pluginManifest] = createCommandsFromManifest(*pluginManifest, binaryPath)
	}

	return pluginsWithCommands, nil
}

func createCommandsFromManifest(pluginManifest PluginManifest, binaryPath string) []*cobra.Command {
	var commands []*cobra.Command

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
	if binaryFileMode&0111 == 0 {
		os.Chmod(binaryPath, binaryFileMode|0111)
	}

	return binaryPath, nil
}
