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
	"testing"

	"github.com/Masterminds/semver/v3"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/plugin"
	"github.com/spf13/cobra"
	"github.com/stretchr/testify/require"
)

func getTestPlugins(t *testing.T) *plugin.ValidatedPlugins {
	t.Helper()
	version1, err := semver.NewVersion("1.4.5")
	require.NoError(t, err)
	version2, err := semver.NewVersion("1.2.3")
	require.NoError(t, err)

	return &plugin.ValidatedPlugins{
		ValidPlugins: []*plugin.Plugin{
			{
				Name:        "plugin1",
				Description: "plugin1 description",
				Version:     version1,
				Commands: []*plugin.Command{
					{Name: "command1"},
					{Name: "command 2"},
				},
				Github: &plugin.Github{
					Owner: "owner1",
					Name:  "repo1",
				},
			},
			{
				Name:        "plugin2",
				Description: "plugin2 description",
				Version:     version2,
				Commands: []*plugin.Command{
					{Name: "command3"},
					{Name: "command4"},
				},
				Github: &plugin.Github{
					Owner: "owner2",
					Name:  "repo2",
				},
			},
			{
				Name:        "plugin3",
				Description: "plugin3 description",
				Version:     version2,
				Commands: []*plugin.Command{
					{Name: "command5"},
					{Name: "command6"},
				},
				Github: &plugin.Github{
					Owner: "owner3",
					Name:  "repo3",
				},
			},
			{
				Name:        "plugin5",
				Description: "plugin5 description",
				Version:     version2,
				Commands: []*plugin.Command{
					{Name: "command7"},
					{Name: "command8"},
				},
				Github: &plugin.Github{
					Owner: "owner5",
					Name:  "repo5",
				},
			},
		},
		PluginsWithDuplicateManifestName: []*plugin.Plugin{
			{
				Name:        "plugin5",
				Description: "plugin5 duplicate",
				Version:     version2,
				Commands: []*plugin.Command{
					{Name: "command9"},
				},
				Github: &plugin.Github{
					Owner: "owner5-duplicate",
					Name:  "repo5-duplicate",
				},
			},
		},
		PluginsWithDuplicateCommands: []*plugin.Plugin{
			{
				Name:        "plugin6",
				Description: "plugin6 description",
				Version:     version2,
				Commands: []*plugin.Command{
					{Name: "command7"},
				},
				Github: &plugin.Github{
					Owner: "owner6",
					Name:  "repo6",
				},
			},
		},
	}
}

func getTestCommands(t *testing.T) []*cobra.Command {
	t.Helper()
	return []*cobra.Command{
		{
			Use:     "testcommand1",
			Aliases: []string{"testcommand1alias"},
			Annotations: map[string]string{
				sourceType:       "plugin",
				sourcePluginName: "testplugin1",
			},
		},
		{
			Use: "testcommand2",
			Annotations: map[string]string{
				sourceType:       "plugin",
				sourcePluginName: "testplugin2",
			},
		},
		{
			Use: "testcommand3",
			Annotations: map[string]string{
				sourceType:       "plugin",
				sourcePluginName: "testplugin3",
			},
		},
	}
}

func Test_findPluginWithArg(t *testing.T) {
	tests := []struct {
		name    string
		arg     string
		want    string
		wantErr bool
	}{
		{
			name:    "Plugin found with plugin name",
			arg:     "plugin3",
			want:    "plugin3",
			wantErr: false,
		},
		{
			name:    "Plugin found with github values",
			arg:     "owner1/repo1@v1",
			want:    "plugin1",
			wantErr: false,
		},
		{
			name:    "Plugin found with github URL",
			arg:     "https://github.com/owner2/repo2/",
			want:    "plugin2",
			wantErr: false,
		},
		{
			name:    "Plugin not found",
			arg:     "this is not a valid args",
			want:    "",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			opts := &Opts{
				plugins: getTestPlugins(t),
			}
			got, err := opts.findPluginWithArg(tt.arg)
			if (err != nil) != tt.wantErr {
				t.Errorf("findPluginByGithubValues() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != nil && got.Name != tt.want {
				t.Errorf("findPluginByGithubValues() = %v, want %v", got.Name, tt.want)
			}
			if got == nil && tt.want != "" {
				t.Errorf("findPluginByGithubValues() = nil, want %v", tt.want)
			}
		})
	}
}

func Test_createExistingCommandsSet(t *testing.T) {
	existingCommandsSet := createExistingCommandsSet(getTestCommands(t))

	require.True(t, existingCommandsSet.Contains("testcommand1"))
	require.True(t, existingCommandsSet.Contains("testcommand1alias"))
	require.True(t, existingCommandsSet.Contains("testcommand2"))
	require.False(t, existingCommandsSet.Contains("testcommand2alias"))
	require.True(t, existingCommandsSet.Contains("testcommand3"))
	require.False(t, existingCommandsSet.Contains("testcommand4"))
}

func Test_findPluginWithGithubValues(t *testing.T) {
	tests := []struct {
		name      string
		repoOwner string
		repoName  string
		want      string
		wantErr   bool
	}{
		{
			name:      "Plugin found",
			repoOwner: "owner3",
			repoName:  "repo3",
			want:      "plugin3",
			wantErr:   false,
		},
		{
			name:      "Plugin not found",
			repoOwner: "owner4",
			repoName:  "repo4",
			want:      "",
			wantErr:   true,
		},
		{
			name:      "Plugin found with duplicate name",
			repoOwner: "owner5-duplicate",
			repoName:  "repo5-duplicate",
			want:      "plugin5",
			wantErr:   false,
		},
		{
			name:      "Plugin found with duplicate command",
			repoOwner: "owner6",
			repoName:  "repo6",
			want:      "plugin6",
			wantErr:   false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			opts := &Opts{
				plugins: getTestPlugins(t),
			}
			got, err := opts.findPluginWithGithubValues(tt.repoOwner, tt.repoName)
			if (err != nil) != tt.wantErr {
				t.Errorf("findPluginByGithubValues() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != nil && got.Name != tt.want {
				t.Errorf("findPluginByGithubValues() = %v, want %v", got.Name, tt.want)
			}
			if got == nil && tt.want != "" {
				t.Errorf("findPluginByGithubValues() = nil, want %v", tt.want)
			}
		})
	}
}

func Test_findPluginWithName(t *testing.T) {
	tests := []struct {
		testName string
		name     string
		want     string
		wantErr  bool
	}{
		{
			testName: "Plugin found",
			name:     "plugin3",
			want:     "plugin3",
			wantErr:  false,
		},
		{
			testName: "Plugin not found",
			name:     "plugin4",
			want:     "",
			wantErr:  true,
		},
		{
			testName: "Plugin found with duplicate name",
			name:     "plugin5",
			want:     "",
			wantErr:  true,
		},
		{
			testName: "Plugin found with duplicate command",
			name:     "plugin6",
			want:     "plugin6",
			wantErr:  false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.testName, func(t *testing.T) {
			opts := &Opts{
				plugins: getTestPlugins(t),
			}
			got, err := opts.findPluginWithName(tt.name)
			if (err != nil) != tt.wantErr {
				t.Errorf("findPluginByName() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != nil && got.Name != tt.want {
				t.Errorf("findPluginByName() = %v, want %v", got.Name, tt.want)
			}
			if got == nil && tt.want != "" {
				t.Errorf("findPluginByName() = nil, want %v", tt.want)
			}
		})
	}
}

func Test_validateManifest(t *testing.T) {
	validManifest := &plugin.Manifest{
		Name:        "name",
		Description: "description",
		Binary:      "binary",
		Version:     "1.0.0",
		Commands: map[string]plugin.ManifestCommand{
			"command1": {Description: "command description"},
		},
	}
	invalidManifest := &plugin.Manifest{}

	require.NoError(t, validateManifest(validManifest))
	require.Error(t, validateManifest(invalidManifest))
}
