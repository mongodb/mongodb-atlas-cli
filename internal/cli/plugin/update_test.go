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
)

func Test_extractPluginSpecifierAndVersionFromArg(t *testing.T) {
	var expectedVersion, _ = semver.NewVersion("1.0.0")

	tests := []struct {
		arg                     string
		expectedPluginSpecifier string
		expectVersion           bool
		expectError             bool
	}{
		{
			arg:                     "mongodb/atlas-cli-plugin-example",
			expectedPluginSpecifier: "mongodb/atlas-cli-plugin-example",
			expectVersion:           false,
			expectError:             false,
		},
		{
			arg:                     "atlas-cli-plugin-example@1.0.0",
			expectedPluginSpecifier: "atlas-cli-plugin-example",
			expectVersion:           true,
			expectError:             false,
		},
		{
			arg:                     "atlas-cli-plugin-example@",
			expectedPluginSpecifier: "",
			expectVersion:           false,
			expectError:             true,
		},
		{
			arg:                     "mongodb/atlas-cli-plugin-example/",
			expectedPluginSpecifier: "mongodb/atlas-cli-plugin-example/",
			expectVersion:           false,
			expectError:             false,
		},
		{
			arg:                     "mongodb/atlas-cli-plugin-example/@v1",
			expectedPluginSpecifier: "mongodb/atlas-cli-plugin-example/",
			expectVersion:           true,
			expectError:             false,
		},
		{
			arg:                     "https://github.com/mongodb/atlas-cli-plugin-example",
			expectedPluginSpecifier: "https://github.com/mongodb/atlas-cli-plugin-example",
			expectVersion:           false,
			expectError:             false,
		},
		{
			arg:                     "https://github.com/mongodb/atlas-cli-plugin-example@v1.0",
			expectedPluginSpecifier: "https://github.com/mongodb/atlas-cli-plugin-example",
			expectVersion:           false,
			expectError:             false,
		},
		{
			arg:                     "github.com/mongodb/atlas-cli-plugin-example/",
			expectedPluginSpecifier: "github.com/mongodb/atlas-cli-plugin-example/",
			expectVersion:           false,
			expectError:             false,
		},
		{
			arg:                     "github.com/mongodb/atlas-cli-plugin-example/@v1.0.0",
			expectedPluginSpecifier: "github.com/mongodb/atlas-cli-plugin-example/",
			expectVersion:           true,
			expectError:             false,
		},
		{
			arg:                     "/mongodb/atlas-cli-plugin-example/",
			expectedPluginSpecifier: "/mongodb/atlas-cli-plugin-example/",
			expectVersion:           false,
			expectError:             false,
		},
		{
			arg:                     "mongodb@atlas-cli-plugin-example",
			expectedPluginSpecifier: "",
			expectVersion:           false,
			expectError:             true,
		},
		{
			arg:                     "mongodb@atlas-cli-plugin-example@1.0",
			expectedPluginSpecifier: "",
			expectVersion:           false,
			expectError:             true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.arg, func(t *testing.T) {
			pluginSpecifier, version, err := extractPluginSpecifierAndVersionFromArg(tt.arg)

			if (err != nil) != tt.expectError {
				t.Errorf("expected error: %v, got: %v", tt.expectError, err)
			}

			if pluginSpecifier != tt.expectedPluginSpecifier {
				t.Errorf("expected plugin specifier: %s, got: %s", tt.expectedPluginSpecifier, pluginSpecifier)
			}

			if tt.expectVersion && !expectedVersion.Equal(version) {
				t.Errorf("expected version: %s, got: %s", expectedVersion.String(), version.String())
			}
		})
	}
}
