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
	"path/filepath"

	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/log"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"
)

type PluginManifest struct {
	Name        string `yaml:"name"`
	Description string `yaml:"description"`
	Binary      string `yaml:"binary"`
	Version     string `yaml:"version"`
	Commands    map[string]struct {
		Description string `yaml:"description"`
	} `yaml:"commands"`
}

func GetPluginCommands(pluginDir string) ([]*cobra.Command, error) {
	files, err := os.ReadDir(pluginDir)

	if err != nil {
		return nil, err
	}

	var commands []*cobra.Command

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
			log.Warningf("plugin invalid: manifest file could not be parsed")
			continue
		}

		log.Debugf("manifest name %s\n", pluginManifest.Name)
		log.Debugf("manifest description %s\n", pluginManifest.Description)
		log.Debugf("manifest binary %s\n", pluginManifest.Binary)
		log.Debugf("manifest version %s\n", pluginManifest.Version)

		for command, value := range pluginManifest.Commands {
			log.Debugf("\tCommand %s, Description: %s\n", command, value.Description)
		}

		// pluginCmd := &cobra.Command{
		// 	Use: pluginManifest.Name,
		// }

	}

	return commands, nil
}

func getManifestFileBytes(pluginDirectoryPath string) ([]byte, error) {
	validManifestFilenames := []string{"manifest.yml", "manifest.yaml"}

	for _, filename := range validManifestFilenames {
		manifestFilePath := filepath.Join(pluginDirectoryPath, filename)

		info, err := os.Stat(manifestFilePath)
		if os.IsNotExist(err) || info.IsDir() {
			continue
		}

		manifestFileData, err := os.ReadFile(manifestFilePath)

		if err != nil {
			continue
		}

		return manifestFileData, nil
	}

	return nil, fmt.Errorf("plugin invalid: manifest file does not exist in plugin folder %s", pluginDirectoryPath)
}

func parseManifestFile(manifestFileData []byte) (*PluginManifest, error) {
	var pluginManifest PluginManifest

	if err := yaml.Unmarshal(manifestFileData, &pluginManifest); err != nil {
		return nil, err
	}

	return &pluginManifest, nil
}