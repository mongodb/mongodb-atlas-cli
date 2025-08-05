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

	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
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

func Test_isAlreadyInstalled(t *testing.T) {
	plugins := getTestPlugins(t)

	tests := []struct {
		name             string
		firstClassPlugin *FirstClassPlugin
		expected         bool
	}{
		{
			name:             "Is already installed",
			firstClassPlugin: &FirstClassPlugin{Name: "plugin2"},
			expected:         true,
		},
		{
			name:             "Is not installed",
			firstClassPlugin: &FirstClassPlugin{Name: "plugin4"},
			expected:         false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Call the method and check the result
			result := tt.firstClassPlugin.isAlreadyInstalled(plugins)
			assert.Equal(t, tt.expected, result)
		})
	}
}
