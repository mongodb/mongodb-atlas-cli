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
	"path"
	"path/filepath"
	"strings"

	"github.com/Masterminds/semver/v3"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/set"
	"gopkg.in/yaml.v3"
)

const (
	ExtraPluginDirectoryEnvKey = "ATLAS_CLI_EXTRA_PLUGIN_DIRECTORY"
)

type ManifestGithubValues struct {
	Owner string `yaml:"owner,omitempty"`
	Name  string `yaml:"name,omitempty"`
}

type ManifestCommand struct {
	Aliases     []string `yaml:"aliases,omitempty"`
	Description string   `yaml:"description,omitempty"`
}

type Manifest struct {
	Name                string                     `yaml:"name,omitempty"`
	Description         string                     `yaml:"description,omitempty"`
	Binary              string                     `yaml:"binary,omitempty"`
	Version             string                     `yaml:"version,omitempty"`
	Github              *ManifestGithubValues      `yaml:"github"`
	Commands            map[string]ManifestCommand `yaml:"commands,omitempty"`
	PluginDirectoryPath string
}

func (m *Manifest) IsValid() (bool, []error) {
	var errorsList []error
	errorMessage := `value "%s" is not defined`

	if m.Name == "" {
		errorsList = append(errorsList, fmt.Errorf(errorMessage, "name"))
	}
	if m.Description == "" {
		errorsList = append(errorsList, fmt.Errorf(errorMessage, "description"))
	}
	if m.Binary == "" {
		errorsList = append(errorsList, fmt.Errorf(errorMessage, "binary"))
	}
	if m.Version == "" {
		errorsList = append(errorsList, fmt.Errorf(errorMessage, "version"))
	} else if _, err := semver.NewVersion(m.Version); err != nil {
		errorsList = append(errorsList, errors.New(`value in field "version" is not a valid semantic version`))
	}

	if m.Github != nil {
		if m.Github.Owner == "" {
			errorsList = append(errorsList, fmt.Errorf(errorMessage, "github owner"))
		}
		if m.Github.Name == "" {
			errorsList = append(errorsList, fmt.Errorf(errorMessage, "github name"))
		}
	}

	switch {
	case m.Commands == nil:
		errorsList = append(errorsList, fmt.Errorf(errorMessage, "commands"))
	case len(m.Commands) == 0:
		errorsList = append(errorsList, errors.New("the plugin needs to contain at least one command"))
	default:
		for command, value := range m.Commands {
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

func loadManifestsFromPluginDirectories() []*Manifest {
	var manifests []*Manifest

	// Load manifests from plugin directories
	if defaultPluginDirectory, err := GetDefaultPluginDirectory(); err == nil {
		if loadedManifests, err := getManifestsFromPluginsDirectory(defaultPluginDirectory); err != nil {
			logPluginWarning(`could not load manifests from directory "%s" because of error: %s`, defaultPluginDirectory, err.Error())
		} else {
			manifests = append(manifests, loadedManifests...)
		}
	}

	if extraPluginDir := os.Getenv(ExtraPluginDirectoryEnvKey); extraPluginDir != "" {
		if loadedManifests, err := getManifestsFromPluginsDirectory(extraPluginDir); err != nil {
			logPluginWarning(`could not load plugins from folder "%s" provided in environment variable ATLAS_CLI_EXTRA_PLUGIN_DIRECTORY: %s`, extraPluginDir, err.Error())
		} else {
			manifests = append(manifests, loadedManifests...)
		}
	}

	return manifests
}

func getManifestsFromPluginsDirectory(pluginsDirectory string) ([]*Manifest, error) {
	files, err := os.ReadDir(pluginsDirectory)

	if err != nil {
		return nil, err
	}

	manifests := make([]*Manifest, 0, len(files))

	for _, directory := range files {
		if !directory.IsDir() {
			continue
		}

		pluginDirectoryPath := path.Join(pluginsDirectory, directory.Name())
		manifest, err := GetManifestFromPluginDirectory(pluginDirectoryPath)
		if err != nil {
			logPluginWarning(err.Error())
			continue
		}

		manifests = append(manifests, manifest)
	}

	return manifests, nil
}

func GetManifestFromPluginDirectory(pluginDirectoryPath string) (*Manifest, error) {
	manifestFileData, err := getManifestFileBytes(pluginDirectoryPath)
	if err != nil {
		return nil, err
	}

	manifest, err := parseManifestFile(manifestFileData)
	if err != nil {
		logPluginWarning(`manifest file of plugin in directory "%s"could not be parsed`, pluginDirectoryPath)
		return nil, err
	}

	if valid, errorList := manifest.IsValid(); !valid {
		var manifestErrorLog strings.Builder
		manifestErrorLog.WriteString(fmt.Sprintf("plugin in directory \"%s\" could not be loaded due to the following error(s) in the manifest.yaml:\n", pluginDirectoryPath))
		for _, err := range errorList {
			manifestErrorLog.WriteString(fmt.Sprintf("\t- %s\n", err.Error()))
		}
		return nil, errors.New(manifestErrorLog.String())
	}

	manifest.PluginDirectoryPath = pluginDirectoryPath
	_, err = getPathToExecutableBinary(manifest)
	if err != nil {
		logPluginWarning(err.Error())
		return nil, err
	}

	return manifest, nil
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

func removeManifestsWithDuplicateNames(manifests []*Manifest) ([]*Manifest, []string) {
	manifestCountMap := make(map[string]int)
	uniqueManifestMap := make(map[string]*Manifest)
	var duplicateManifestNames []string

	// iterate through all manifests and add them to a map, if a manifest with the same name
	// has alread been added remove it from said map and add name to slice of duplicate manifest names
	for _, manifest := range manifests {
		manifestCountMap[manifest.Name]++

		if manifestCountMap[manifest.Name] == 1 {
			uniqueManifestMap[manifest.Name] = manifest
		} else if _, ok := uniqueManifestMap[manifest.Name]; ok {
			delete(uniqueManifestMap, manifest.Name)
			duplicateManifestNames = append(duplicateManifestNames, manifest.Name)
		}
	}

	uniqueManifests := make([]*Manifest, 0, len(uniqueManifestMap))
	for _, manifest := range uniqueManifestMap {
		uniqueManifests = append(uniqueManifests, manifest)
	}

	return uniqueManifests, duplicateManifestNames
}

func getUniqueManifests(manifests []*Manifest, existingCommandsSet set.Set[string]) ([]*Manifest, []*Manifest) {
	uniqueManifests := make([]*Manifest, 0, len(manifests))
	var duplicateManifests []*Manifest

	existingCommandsSet.Add("plugin")

	for _, manifest := range manifests {
		if manifest.HasDuplicateCommand(existingCommandsSet) {
			duplicateManifests = append(duplicateManifests, manifest)
			continue
		}
		for cmdName := range manifest.Commands {
			existingCommandsSet.Add(cmdName)
		}
		uniqueManifests = append(uniqueManifests, manifest)
	}

	return uniqueManifests, duplicateManifests
}

func (m *Manifest) HasDuplicateCommand(existingCommandsSet set.Set[string]) bool {
	for cmdName, cmd := range m.Commands {
		if existingCommandsSet.Contains(cmdName) {
			return true
		}
		for _, alias := range cmd.Aliases {
			if existingCommandsSet.Contains(alias) {
				return true
			}
		}
	}
	return false
}

func getPathToExecutableBinary(manifest *Manifest) (string, error) {
	binaryPath := path.Join(manifest.PluginDirectoryPath, manifest.Binary)

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
