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

	"github.com/Masterminds/semver/v3"
	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func getTestManifest() *Manifest {
	return &Manifest{
		Name:                "testPlugin",
		Description:         "Test plugin description",
		Binary:              "binary",
		PluginDirectoryPath: "/plugins/testPlugin",
		Version:             "1.2.3",
		Commands: map[string]ManifestCommand{
			"testCommand": {Description: "Test command", Aliases: []string{"T1", "T2"}},
		},
	}
}

func Test_IsPluginCmd(t *testing.T) {
	pluginCmd := cobra.Command{
		Annotations: map[string]string{
			sourceType: PluginSourceType,
		}}
	cmd := cobra.Command{}

	assert.True(t, IsPluginCmd(&pluginCmd))
	assert.False(t, IsPluginCmd(&cmd))
}

func Test_GetCobraCommands(t *testing.T) {
	manifest := getTestManifest()

	// Create a mock Plugin from the mock Manifest
	plugin, err := createPluginFromManifest(manifest)
	require.NoError(t, err)

	commands := plugin.GetCobraCommands()

	assert.Len(t, commands, 1)
	assert.Equal(t, "testCommand", commands[0].Use)
	assert.Equal(t, manifest.Commands["testCommand"].Description, commands[0].Short)
	assert.Equal(t, manifest.Commands["testCommand"].Aliases, commands[0].Aliases)
	assert.NotNil(t, commands[0].RunE)
}

func Test_createPluginFromManifest(t *testing.T) {
	manifest := getTestManifest()

	plugin, err := createPluginFromManifest(manifest)
	require.NoError(t, err)

	assert.Equal(t, plugin.Name, manifest.Name)
	assert.Equal(t, plugin.Description, manifest.Description)
	assert.Equal(t, plugin.BinaryName, manifest.Binary)
	assert.Equal(t, plugin.PluginDirectoryPath, manifest.PluginDirectoryPath)

	manifestSemverVersion, err := semver.NewVersion(manifest.Version)
	require.NoError(t, err)
	assert.True(t, plugin.Version.Equal(manifestSemverVersion))

	assert.Len(t, plugin.Commands, len(manifest.Commands))
	assert.Equal(t, plugin.Commands[0].Description, manifest.Commands["testCommand"].Description)
	assert.Equal(t, plugin.Commands[0].Aliases, manifest.Commands["testCommand"].Aliases)
}
