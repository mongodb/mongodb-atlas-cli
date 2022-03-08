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
// +build unit

package file_test

import (
	"fmt"
	"testing"

	"github.com/mongodb/mongocli/internal/file"
	"github.com/spf13/afero"
	"gopkg.in/yaml.v2"
)

const (
	noExtFileName  = "test"
	txtFileName    = "test.txt"
	jsonFileName   = "test.json"
	yamlFileName   = "test.yaml"
	unsupportedMsg = "unsupported file type: txt"
)

func TestFile(t *testing.T) {
	t.Run("load file does not exists", func(t *testing.T) {
		appFS := afero.NewMemMapFs()
		filename := "test.xml"
		err := file.Load(appFS, filename, nil)
		if err == nil || err.Error() != fmt.Sprintf("file not found: %s", filename) {
			t.Errorf("Load() unexpected error: %v", err)
		}
	})
	t.Run("load file with no ext", func(t *testing.T) {
		appFS := afero.NewMemMapFs()
		filename := noExtFileName
		_ = afero.WriteFile(appFS, filename, []byte(""), 0600)
		err := file.Load(appFS, filename, nil)
		if err == nil || err.Error() != fmt.Sprintf("filename: %s requires valid extension", filename) {
			t.Errorf("Load() unexpected error: %v", err)
		}
	})
	t.Run("load file with invalid ext", func(t *testing.T) {
		appFS := afero.NewMemMapFs()
		filename := txtFileName
		_ = afero.WriteFile(appFS, filename, []byte(""), 0600)
		err := file.Load(appFS, filename, nil)
		if err == nil || err.Error() != unsupportedMsg {
			t.Errorf("Load() unexpected error: %v", err)
		}
	})
	t.Run("load valid json file", func(t *testing.T) {
		appFS := afero.NewMemMapFs()
		filename := jsonFileName
		_ = afero.WriteFile(appFS, filename, []byte("{}"), 0600)
		out := new(map[string]interface{})
		err := file.Load(appFS, filename, out)
		if err != nil {
			t.Fatalf("Load() unexpected error: %v", err)
		}
	})
	t.Run("load valid yaml file", func(t *testing.T) {
		appFS := afero.NewMemMapFs()
		filename := yamlFileName
		_ = afero.WriteFile(appFS, filename, []byte(""), 0600)
		out := new(map[string]interface{})
		err := file.Load(appFS, filename, out)
		if err != nil {
			t.Fatalf("Load() unexpected error: %v", err)
		}
	})
	t.Run("save file with no ext", func(t *testing.T) {
		appFS := afero.NewMemMapFs()
		filename := "test"
		err := file.Save(appFS, filename, nil)
		if err == nil || err.Error() != "filename: test requires valid extension" {
			t.Errorf("Save() unexpected error: %v", err)
		}
	})
	t.Run("save file with wrong ext", func(t *testing.T) {
		appFS := afero.NewMemMapFs()
		filename := txtFileName
		err := file.Save(appFS, filename, nil)
		if err == nil || err.Error() != unsupportedMsg {
			t.Errorf("Save() unexpected error: %v", err)
		}
	})
	t.Run("save valid yaml file", func(t *testing.T) {
		appFS := afero.NewMemMapFs()
		filename := yamlFileName

		type Test struct {
			name string
			age  int
		}

		tYaml := Test{
			name: "MongoDB",
			age:  100,
		}

		yamlData, _ := yaml.Marshal(&tYaml)

		err := file.Save(appFS, filename, yamlData)
		if err != nil {
			t.Fatalf("Save() unexpected error: %v", err)
		}
	})
}
