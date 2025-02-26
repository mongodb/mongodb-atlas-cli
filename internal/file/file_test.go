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
	"encoding/json"
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

func TestDecodePrintWarning(t *testing.T) {
	type Foo struct {
		Bar bool
	}

	t.Run("decode json with unknown fields", func(t *testing.T) {
		foo := &Foo{}
		warning, err := decodeWithWarning([]byte(`{"Bar":true,"Baz":false}`), foo, json.NewDecoder, func(d *json.Decoder) { d.DisallowUnknownFields() })
		require.NoError(t, err)
		require.Equal(t, `json: unknown field "Baz"`, warning)
		require.Equal(t, &Foo{Bar: true}, foo)
	})

	t.Run("decode json with only known fields", func(t *testing.T) {
		foo := &Foo{}
		warning, err := decodeWithWarning([]byte(`{"Bar":true}`), foo, json.NewDecoder, func(d *json.Decoder) { d.DisallowUnknownFields() })
		require.NoError(t, err)
		require.Equal(t, "", warning)
		require.Equal(t, &Foo{Bar: true}, foo)
	})

	t.Run("decode invalid json", func(t *testing.T) {
		foo := &Foo{}
		warning, err := decodeWithWarning([]byte(`{"Bar"`), foo, json.NewDecoder, func(d *json.Decoder) { d.DisallowUnknownFields() })
		require.Error(t, err)
		require.Equal(t, "", warning)
		require.Equal(t, &Foo{}, foo)
	})

	t.Run("decode yaml with unknown fields", func(t *testing.T) {
		foo := &Foo{}
		warning, err := decodeWithWarning([]byte("bar: true\nbaz: false"), foo, yaml.NewDecoder, func(d *yaml.Decoder) { d.KnownFields(true) })
		require.NoError(t, err)
		require.Contains(t, warning, `field baz not found in type file.Foo`)
		require.Equal(t, &Foo{Bar: true}, foo)
	})

	t.Run("decode yaml with only known fields", func(t *testing.T) {
		foo := &Foo{}
		warning, err := decodeWithWarning([]byte(`bar: true`), foo, yaml.NewDecoder, func(d *yaml.Decoder) { d.KnownFields(true) })
		require.NoError(t, err)
		require.Equal(t, "", warning)
		require.Equal(t, &Foo{Bar: true}, foo)
	})

	t.Run("decode invalid yaml", func(t *testing.T) {
		foo := &Foo{}
		warning, err := decodeWithWarning([]byte("bar"), foo, yaml.NewDecoder, func(d *yaml.Decoder) { d.KnownFields(true) })
		require.Error(t, err)
		require.Equal(t, "", warning)
		require.Equal(t, &Foo{}, foo)
	})
}
