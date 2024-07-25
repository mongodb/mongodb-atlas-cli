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

//go:build unit

package file

import (
	"testing"

	"github.com/spf13/afero"
	"github.com/stretchr/testify/require"
	"gopkg.in/yaml.v3"
)

const (
	noExtFileName = "test"
	txtFileName   = "test.txt"
	jsonFileName  = "test.json"
	yamlFileName  = "test.yaml"
	xmlFileName   = "test.xml"
)

func TestLoad(t *testing.T) {
	t.Run("load file does not exists", func(t *testing.T) {
		appFS := afero.NewMemMapFs()
		err := Load(appFS, xmlFileName, nil)
		require.ErrorIs(t, err, ErrFileNotFound)
	})
	t.Run("load file with no ext", func(t *testing.T) {
		appFS := afero.NewMemMapFs()
		_ = afero.WriteFile(appFS, noExtFileName, []byte(""), 0600)
		err := Load(appFS, noExtFileName, nil)
		require.ErrorIs(t, err, ErrMissingFileType)
	})
	t.Run("load file with invalid ext", func(t *testing.T) {
		appFS := afero.NewMemMapFs()
		require.NoError(t, afero.WriteFile(appFS, txtFileName, []byte(""), 0600))
		err := Load(appFS, txtFileName, nil)
		require.ErrorIs(t, err, ErrUnsupportedFileType)
	})
	t.Run("load valid json file", func(t *testing.T) {
		appFS := afero.NewMemMapFs()
		_ = afero.WriteFile(appFS, jsonFileName, []byte("{}"), 0600)
		out := new(map[string]any)
		require.NoError(t, Load(appFS, jsonFileName, out))
	})
	t.Run("load valid yaml file", func(t *testing.T) {
		appFS := afero.NewMemMapFs()
		_ = afero.WriteFile(appFS, yamlFileName, []byte(""), 0600)
		out := new(map[string]any)
		err := Load(appFS, yamlFileName, out)
		require.NotErrorIs(t, err, ErrMissingFileType)
	})
}

func TestSave(t *testing.T) {
	t.Run("save file with no ext", func(t *testing.T) {
		appFS := afero.NewMemMapFs()
		filename := "test"
		err := Save(appFS, filename, nil)
		require.ErrorIs(t, err, ErrMissingFileType)
	})
	t.Run("save file with wrong ext", func(t *testing.T) {
		appFS := afero.NewMemMapFs()
		err := Save(appFS, txtFileName, nil)
		require.ErrorIs(t, err, ErrUnsupportedFileType)
	})
	t.Run("save valid yaml file", func(t *testing.T) {
		appFS := afero.NewMemMapFs()
		type Test struct {
			name string
			age  int
		}
		tYaml := Test{
			name: "MongoDB",
			age:  100,
		}

		yamlData, _ := yaml.Marshal(&tYaml)
		require.NoError(t, Save(appFS, yamlFileName, yamlData))
	})
}
