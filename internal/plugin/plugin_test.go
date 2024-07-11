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

//go:build unit

package plugin

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func getTestManifest() *Manifest {
	return &Manifest{
		Name:        "testPlugin",
		Description: "Test plugin description",
		BinaryPath:  "/plugins/testPlugin/binary",
		Version:     "1.2.3",
		Commands: map[string]struct {
			Description string `yaml:"description,omitempty"`
		}{
			"testCommand":  {"Test command"},
			"testCommand2": {"Test command 2"},
		},
	}
}

func Test_GetCobraCommands(t *testing.T) {
	manifest := getTestManifest()

	// Create a mock Plugin from the mock Manifest
	plugin := createPluginFromManifest(manifest)

	commands := plugin.GetCobraCommands()

	assert.Len(t, commands, 2)

	index := 0
	for cmd, value := range manifest.Commands {
		assert.Equal(t, commands[index].Use, cmd)
		assert.Equal(t, commands[index].Short, value.Description)
		assert.NotNil(t, commands[index].RunE)
		index++
	}
}

func Test_createPluginFromManifest(t *testing.T) {
	manifest := getTestManifest()

	plugin := createPluginFromManifest(manifest)

	assert.Equal(t, plugin.Name, manifest.Name)
	assert.Equal(t, plugin.Description, manifest.Description)
	assert.Equal(t, plugin.BinaryPath, manifest.BinaryPath)
	assert.Equal(t, plugin.Version, manifest.Version)

	assert.Len(t, plugin.Commands, len(manifest.Commands))
}
