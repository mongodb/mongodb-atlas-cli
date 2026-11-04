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
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_IsFirstClassPluginCmd(t *testing.T) {
	tests := []struct {
		name        string
		annotations map[string]string
		expected    bool
	}{
		{
			name: "Correct sourceType annotation",
			annotations: map[string]string{
				sourceType: FirstClassSourceType,
			},
			expected: true,
		},
		{
			name: "Incorrect sourceType annotation",
			annotations: map[string]string{
				sourceType: "anotherType",
			},
			expected: false,
		},
		{
			name:        "Missing sourceType annotation",
			annotations: map[string]string{},
			expected:    false,
		},
		{
			name:        "No annotations",
			annotations: nil,
			expected:    false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cmd := &cobra.Command{
				Annotations: tt.annotations,
			}

			result := IsFirstClassPluginCmd(cmd)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func Test_getInstalledPlugin(t *testing.T) {
	plugins := getTestPlugins(t)

	tests := []struct {
		name             string
		firstClassPlugin *FirstClassPlugin
		expectFound      bool
	}{
		{
			name:             "Is already installed",
			firstClassPlugin: &FirstClassPlugin{Name: "plugin2"},
			expectFound:      true,
		},
		{
			name:             "Is not installed",
			firstClassPlugin: &FirstClassPlugin{Name: "plugin4"},
			expectFound:      false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Call the method and check the result
			result := tt.firstClassPlugin.getInstalledPlugin(plugins)
			if tt.expectFound {
				assert.NotNil(t, result)
				assert.Equal(t, tt.firstClassPlugin.Name, result.Name)
			} else {
				assert.Nil(t, result)
			}
		})
	}
}

func Test_needsUpdate(t *testing.T) {
	version123, err := semver.NewVersion("1.2.3")
	require.NoError(t, err)

	tests := []struct {
		name             string
		firstClassPlugin *FirstClassPlugin
		installedPlugin  *plugin.Plugin
		expected         bool
	}{
		{
			name:             "No MinVersion specified",
			firstClassPlugin: &FirstClassPlugin{Name: "test", MinVersion: ""},
			installedPlugin:  &plugin.Plugin{Name: "test", Version: version123},
			expected:         false,
		},
		{
			name:             "Installed version equals MinVersion",
			firstClassPlugin: &FirstClassPlugin{Name: "test", MinVersion: "1.2.3"},
			installedPlugin:  &plugin.Plugin{Name: "test", Version: version123},
			expected:         false,
		},
		{
			name:             "Installed version greater than MinVersion",
			firstClassPlugin: &FirstClassPlugin{Name: "test", MinVersion: "1.0.0"},
			installedPlugin:  &plugin.Plugin{Name: "test", Version: version123},
			expected:         false,
		},
		{
			name:             "Installed version less than MinVersion",
			firstClassPlugin: &FirstClassPlugin{Name: "test", MinVersion: "2.0.0"},
			installedPlugin:  &plugin.Plugin{Name: "test", Version: version123},
			expected:         true,
		},
		{
			name:             "Invalid MinVersion format",
			firstClassPlugin: &FirstClassPlugin{Name: "test", MinVersion: "invalid"},
			installedPlugin:  &plugin.Plugin{Name: "test", Version: version123},
			expected:         false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.firstClassPlugin.needsUpdate(tt.installedPlugin)
			assert.Equal(t, tt.expected, result)
		})
	}
}
