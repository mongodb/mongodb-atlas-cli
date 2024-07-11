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
	"path/filepath"
	"regexp"
	"strings"

	"github.com/spf13/cobra"
	"gopkg.in/yaml.v2"
)

type Manifest struct {
	Name        string `yaml:"name,omitempty"`
	Description string `yaml:"description,omitempty"`
	Binary      string `yaml:"binary,omitempty"`
	Version     string `yaml:"version,omitempty"`
	Commands    map[string]struct {
		Description string `yaml:"description,omitempty"`
	} `yaml:"commands,omitempty"`
	BinaryPath string
}

func (p *Manifest) IsValid() (bool, []error) {
	var errorsList []error
	errorMessage := `value "%s" is not defined`

	if p.Name == "" {
		errorsList = append(errorsList, fmt.Errorf(errorMessage, "name"))
	}
	if p.Description == "" {
		errorsList = append(errorsList, fmt.Errorf(errorMessage, "description"))
	}
	if p.Binary == "" {
		errorsList = append(errorsList, fmt.Errorf(errorMessage, "binary"))
	}
	if p.Version == "" {
		errorsList = append(errorsList, fmt.Errorf(errorMessage, "version"))
	} else if valid, _ := regexp.MatchString(`^\d+\.\d+\.\d+$`, p.Version); !valid {
		errorsList = append(errorsList, errors.New(`value in field "version" is not a valid semantic version`))
	}
	if p.Commands == nil {
		errorsList = append(errorsList, fmt.Errorf(errorMessage, "commands"))
	} else {
		for command, value := range p.Commands {
			if value.Description == "" {
				errorsList = append(errorsList, fmt.Errorf(`value "description" in command "%s" is not defined`, command))
			}
		}
	}

	if len(errorsList) > 0 {
		return false, errorsList
	}
	return true, nil
}

func getManifestsFromPluginDirectory(pluginDirectory string) ([]*Manifest, error) {
	files, err := os.ReadDir(pluginDirectory)

	if err != nil {
		return nil, err
	}

	manifests := make([]*Manifest, 0, len(files))

	for _, directory := range files {
		if !directory.IsDir() {
			continue
		}

		pluginDirectoryPath := fmt.Sprintf("%s/%s", pluginDirectory, directory.Name())

		manifestFileData, err := getManifestFileBytes(pluginDirectoryPath)

		if err != nil {
			continue
		}
		pluginManifest, err := parseManifestFile(manifestFileData)

		if err != nil {
			logPluginWarning(`manifest file of plugin in directory "%s"could not be parsed`, pluginDirectoryPath)
			continue
		}

		if valid, errors := pluginManifest.IsValid(); !valid {
			var manifestErrorLog strings.Builder
			manifestErrorLog.WriteString(fmt.Sprintf("plugin in directory \"%s\" could not be loaded due to the following error(s) in the manifest.yaml:\n", pluginDirectoryPath))
			for _, err := range errors {
				manifestErrorLog.WriteString(fmt.Sprintf("\t- %s\n", err.Error()))
			}
			logPluginWarning(manifestErrorLog.String())
			continue
		}

		binaryPath, err := getPathToExecutableBinary(pluginDirectoryPath, pluginManifest.Binary)

		if err != nil {
			logPluginWarning(err.Error())
			continue
		}

		pluginManifest.BinaryPath = binaryPath

		manifests = append(manifests, pluginManifest)
	}

	return manifests, nil
}

func parseManifestFile(manifestFileData []byte) (*Manifest, error) {
	var pluginManifest Manifest

	if err := yaml.Unmarshal(manifestFileData, &pluginManifest); err != nil {
		return nil, err
	}

	return &pluginManifest, nil
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

func getUniqueManifests(manifests []*Manifest, existingCommands []*cobra.Command) ([]*Manifest, []*Manifest) {
	existingCommandsMap := make(map[string]bool)
	uniqueManifests := make([]*Manifest, 0, len(manifests))
	var duplicateManifests []*Manifest

	for _, cmd := range existingCommands {
		existingCommandsMap[cmd.Name()] = true
	}
	// add reserved keyword for command plugin
	existingCommandsMap["plugin"] = true

	for _, manifest := range manifests {
		if hasDuplicateCommand(manifest, existingCommandsMap) {
			duplicateManifests = append(duplicateManifests, manifest)
			continue
		}
		for cmdName := range manifest.Commands {
			existingCommandsMap[cmdName] = true
		}
		uniqueManifests = append(uniqueManifests, manifest)
	}

	return uniqueManifests, duplicateManifests
}

func hasDuplicateCommand(manifest *Manifest, existingCommandsMap map[string]bool) bool {
	for cmdName := range manifest.Commands {
		if existingCommandsMap[cmdName] {
			return true
		}
	}
	return false
}

func getPathToExecutableBinary(pluginDirectoryPath string, binaryName string) (string, error) {
	binaryPath := fmt.Sprintf("%s/%s", pluginDirectoryPath, binaryName)

	binaryFileInfo, err := os.Stat(binaryPath)

	if err != nil {
		return "", fmt.Errorf(`binary "%s" does not exists`, binaryPath)
	}

	// makes sure that the binary file is made executable if it is not already
	binaryFileMode := binaryFileInfo.Mode()
	const executablePermissions = 0o111

	if binaryFileMode&executablePermissions != 0 {
		return binaryPath, nil
	}

	if err := os.Chmod(binaryPath, binaryFileMode|executablePermissions); err != nil {
		return "", err
	}

	return binaryPath, nil
}
