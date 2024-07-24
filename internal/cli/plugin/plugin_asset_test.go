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
	"fmt"
	"runtime"
	"testing"

	"github.com/Masterminds/semver/v3"
	"github.com/google/go-github/v61/github"
	"github.com/stretchr/testify/assert"
)

func Test_repository(t *testing.T) {
	opts := &AssetOpts{githubRelease: &GithubRelease{owner: "repoOwner", name: "repoName"}}

	assert.Equal(t, opts.githubRelease.owner+"/"+opts.githubRelease.name, opts.repository())
}

func Test_getAssetID(t *testing.T) {
	validAssetName := fmt.Sprintf("plugin_%s_%s", runtime.GOOS, runtime.GOARCH)

	tests := []struct {
		name            string
		pluginAssets    []*github.ReleaseAsset
		expectedAssetID int64
		expectError     bool
	}{
		{
			name: "Valid asset",
			pluginAssets: []*github.ReleaseAsset{
				{
					ID:          github.Int64(1),
					Name:        github.String(validAssetName + ".tar.gz"),
					ContentType: github.String("application/gzip"),
				},
			},
			expectedAssetID: 1,
			expectError:     false,
		},
		{
			name: "No matching asset",
			pluginAssets: []*github.ReleaseAsset{
				{
					ID:          github.Int64(2),
					Name:        github.String("plugin_invalid_assetname.tar.gz"),
					ContentType: github.String("application/gzip"),
				},
			},
			expectedAssetID: 0,
			expectError:     true,
		},
		{
			name: "Non-gzip asset",
			pluginAssets: []*github.ReleaseAsset{
				{
					ID:          github.Int64(3),
					Name:        github.String(validAssetName + ".zip"),
					ContentType: github.String("application/zip"),
				},
			},
			expectedAssetID: 0,
			expectError:     true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			opts := &AssetOpts{
				pluginAssets: tt.pluginAssets,
			}

			assetID, err := opts.getAssetID()
			if (err != nil) != tt.expectError {
				t.Errorf("expected error: %v, got: %v", tt.expectError, err)
			}
			if assetID != tt.expectedAssetID {
				t.Errorf("expected asset ID: %d, got: %d", tt.expectedAssetID, assetID)
			}
		})
	}
}

func Test_parseGithubRepoValues(t *testing.T) {
	const (
		expectedOwner = "mongodb"
		expectedName  = "atlas-cli-plugin-example"
	)

	tests := []struct {
		arg         string
		expectError bool
	}{
		{
			arg:         "mongodb/atlas-cli-plugin-example",
			expectError: false,
		},
		{
			arg:         "mongodb/atlas-cli-plugin-example",
			expectError: false,
		},
		{
			arg:         "mongodb/atlas-cli-plugin-example",
			expectError: false,
		},
		{
			arg:         "mongodb/atlas-cli-plugin-example/",
			expectError: false,
		},
		{
			arg:         "https://github.com/mongodb/atlas-cli-plugin-example",
			expectError: false,
		},
		{
			arg:         "github.com/mongodb/atlas-cli-plugin-example/",
			expectError: false,
		},
		{
			arg:         "/mongodb/atlas-cli-plugin-example/",
			expectError: false,
		},
		{
			arg:         "mongodb@atlas-cli-plugin-example",
			expectError: true,
		},
		{
			arg:         "mongodb@atlas-cli-plugin-example@1.2.3",
			expectError: true,
		},
		{
			arg:         "invalidArgString",
			expectError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.arg, func(t *testing.T) {
			owner, name, err := parseGithubRepoValues(tt.arg)
			if (err != nil) != tt.expectError {
				t.Errorf("expected error: %v, got: %v", tt.expectError, err)
			}
			if !tt.expectError {
				if owner != expectedOwner {
					t.Errorf("expected owner: %s, got: %s", expectedOwner, owner)
				}
				if name != expectedName {
					t.Errorf("expected name: %s, got: %s", expectedName, name)
				}
			}
		})
	}
}

func Test_parseGithubReleaseVersion(t *testing.T) {
	tests := []struct {
		arg             string
		expectedVersion string
		expectError     bool
	}{
		{
			arg:             "1.2.3",
			expectedVersion: "1.2.3",
			expectError:     false,
		},
		{
			arg:             "v1.2.3",
			expectedVersion: "1.2.3",
			expectError:     false,
		},
		{
			arg:             "1.2",
			expectedVersion: "1.2.0",
			expectError:     false,
		},
		{
			arg:             "invalid-version",
			expectedVersion: "",
			expectError:     true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.arg, func(t *testing.T) {
			version, err := parseGithubReleaseVersion(tt.arg)
			if (err != nil) != tt.expectError {
				t.Errorf("expected error: %v, got: %v", tt.expectError, err)
				return
			}
			if version == nil {
				if tt.expectedVersion != "" {
					t.Errorf("expected version to be nil, got: %s", version)
				}
				return
			}
			expectedVersion, _ := semver.NewVersion(tt.expectedVersion)
			if !version.Equal(expectedVersion) {
				t.Errorf("expected version: %s, got: %s", tt.expectedVersion, version)
			}
		})
	}
}
