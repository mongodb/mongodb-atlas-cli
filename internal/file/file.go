// Copyright 2020 MongoDB Inc
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

package file

import (
	"encoding/json"
	"errors"
	"fmt"
	"path/filepath"

	"github.com/mongodb/mongodb-atlas-cli/internal/search"
	"github.com/spf13/afero"
	"gopkg.in/yaml.v2"
)

const (
	yamlName         = "yaml"
	jsonName         = "json"
	ymlName          = "yml"
	configPermission = 0700
	filePermission   = 0600
)

var supportedExts = []string{jsonName, yamlName, ymlName}

var ErrUnsupportedFileType = errors.New("unsupported file type")

// configType gets the config type from a given file path.
func configType(filename string, supported []string) (string, error) {
	ext := filepath.Ext(filename)

	if len(ext) <= 1 {
		return "", fmt.Errorf("filename: %s requires valid extension", filename)
	}

	t := ext[1:]
	if !search.StringInSlice(supported, t) {
		return "", fmt.Errorf("%w: %s", ErrUnsupportedFileType, t)
	}
	return t, nil
}

var ErrFileNotFound = errors.New("file not found")

// Load loads a given filename into the out interface.
// The file should be a valid json or yaml format.
func Load(fs afero.Fs, filename string, out interface{}) error {
	if exists, err := afero.Exists(fs, filename); !exists || err != nil {
		return fmt.Errorf("%w: %s", ErrFileNotFound, filename)
	}

	t, err := configType(filename, supportedExts)
	if err != nil {
		return err
	}

	file, err := afero.ReadFile(fs, filename)
	if err != nil {
		return err
	}

	switch t {
	case yamlName, ymlName:
		if err := yaml.Unmarshal(file, out); err != nil {
			return err
		}
	case jsonName:
		if err := json.Unmarshal(file, out); err != nil {
			return err
		}
	}

	return nil
}

// Save saves a given data interface into a given file path
// The file should be a yaml format.
func Save(fs afero.Fs, filePath string, data interface{}) error {
	var content []byte

	if _, err := configType(filePath, []string{yamlName}); err != nil {
		return err
	}

	content, err := yaml.Marshal(data)
	if err != nil {
		return err
	}

	err = fs.MkdirAll(filepath.Dir(filePath), configPermission)
	if err != nil {
		return err
	}

	return afero.WriteFile(fs, filePath, content, filePermission)
}
