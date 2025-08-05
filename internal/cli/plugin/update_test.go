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
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_extractPluginSpecifierAndVersionFromArg(t *testing.T) {
	var (
		v1, _       = semver.NewVersion("1.0.0")
		v1pre, _    = semver.NewVersion("1.0.0-prerelease")
		v1pseudo, _ = semver.NewVersion("1.0.0-beta+very-meta")
	)

	tests := []struct {
		arg                     string
		expectedPluginSpecifier string
		expectedVersion         *semver.Version
		expectError             require.ErrorAssertionFunc
	}{
		{
			arg:                     "mongodb/atlas-cli-plugin-example",
			expectedPluginSpecifier: "mongodb/atlas-cli-plugin-example",
			expectedVersion:         nil,
			expectError:             require.NoError,
		},
		{
			arg:                     "atlas-cli-plugin-example@1.0.0",
			expectedPluginSpecifier: "atlas-cli-plugin-example",
			expectedVersion:         v1,
			expectError:             require.NoError,
		},
		{
			arg:                     "atlas-cli-plugin-example@1.0.0-prerelease",
			expectedPluginSpecifier: "atlas-cli-plugin-example",
			expectedVersion:         v1pre,
			expectError:             require.NoError,
		},
		{
			arg:                     "atlas-cli-plugin-example@1.0.0-beta+very-meta",
			expectedPluginSpecifier: "atlas-cli-plugin-example",
			expectedVersion:         v1pseudo,
			expectError:             require.NoError,
		},
		{
			arg:                     "atlas-cli-plugin-example@",
			expectedPluginSpecifier: "",
			expectedVersion:         nil,
			expectError:             require.Error,
		},
		{
			arg:                     "mongodb/atlas-cli-plugin-example/",
			expectedPluginSpecifier: "mongodb/atlas-cli-plugin-example/",
			expectedVersion:         nil,
			expectError:             require.NoError,
		},
		{
			arg:                     "mongodb/atlas-cli-plugin-example/@v1",
			expectedPluginSpecifier: "mongodb/atlas-cli-plugin-example/",
			expectedVersion:         v1,
			expectError:             require.NoError,
		},
		{
			arg:                     "https://github.com/mongodb/atlas-cli-plugin-example",
			expectedPluginSpecifier: "https://github.com/mongodb/atlas-cli-plugin-example",
			expectedVersion:         nil,
			expectError:             require.NoError,
		},
		{
			arg:                     "https://github.com/mongodb/atlas-cli-plugin-example@v1.0",
			expectedPluginSpecifier: "https://github.com/mongodb/atlas-cli-plugin-example",
			expectedVersion:         v1,
			expectError:             require.NoError,
		},
		{
			arg:                     "github.com/mongodb/atlas-cli-plugin-example/",
			expectedPluginSpecifier: "github.com/mongodb/atlas-cli-plugin-example/",
			expectedVersion:         nil,
			expectError:             require.NoError,
		},
		{
			arg:                     "github.com/mongodb/atlas-cli-plugin-example/@v1.0.0",
			expectedPluginSpecifier: "github.com/mongodb/atlas-cli-plugin-example/",
			expectedVersion:         v1,
			expectError:             require.NoError,
		},
		{
			arg:                     "/mongodb/atlas-cli-plugin-example/",
			expectedPluginSpecifier: "/mongodb/atlas-cli-plugin-example/",
			expectedVersion:         nil,
			expectError:             require.NoError,
		},
		{
			arg:                     "mongodb@atlas-cli-plugin-example",
			expectedPluginSpecifier: "",
			expectedVersion:         nil,
			expectError:             require.Error,
		},
		{
			arg:                     "mongodb@atlas-cli-plugin-example@1.0",
			expectedPluginSpecifier: "",
			expectedVersion:         nil,
			expectError:             require.Error,
		},
	}

	for _, tt := range tests {
		t.Run(tt.arg, func(t *testing.T) {
			t.Parallel()
			pluginSpecifier, version, err := extractPluginSpecifierAndVersionFromArg(tt.arg)
			tt.expectError(t, err)
			assert.Equal(t, tt.expectedPluginSpecifier, pluginSpecifier)

			if tt.expectedVersion != nil && !tt.expectedVersion.Equal(version) {
				t.Errorf("expected version: %s, got: %s", tt.expectedVersion.String(), version.String())
			}
			if tt.expectedVersion == nil && version != nil {
				t.Errorf("expected version to be nil, got: %s", version.String())
			}
		})
	}
}
