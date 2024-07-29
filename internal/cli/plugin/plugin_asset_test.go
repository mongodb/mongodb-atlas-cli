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
			opts := &AssetOpts{}

			assetID, err := opts.getAssetID(tt.pluginAssets)
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
	var expectedVersion, _ = semver.NewVersion("1.0.0")

	tests := []struct {
		arg           string
		expectVersion bool
		expectError   bool
	}{
		{
			arg:           "mongodb/atlas-cli-plugin-example",
			expectVersion: false,
			expectError:   false,
		},
		{
			arg:           "mongodb/atlas-cli-plugin-example@1.0.0",
			expectVersion: true,
			expectError:   false,
		},
		{
			arg:           "mongodb/atlas-cli-plugin-example@",
			expectVersion: false,
			expectError:   true,
		},
		{
			arg:           "mongodb/atlas-cli-plugin-example/",
			expectVersion: false,
			expectError:   false,
		},
		{
			arg:           "mongodb/atlas-cli-plugin-example/@v1",
			expectVersion: true,
			expectError:   false,
		},
		{
			arg:           "https://github.com/mongodb/atlas-cli-plugin-example",
			expectVersion: false,
			expectError:   false,
		},
		{
			arg:           "https://github.com/mongodb/atlas-cli-plugin-example@v1.0",
			expectVersion: false,
			expectError:   false,
		},
		{
			arg:           "github.com/mongodb/atlas-cli-plugin-example/",
			expectVersion: false,
			expectError:   false,
		},
		{
			arg:           "github.com/mongodb/atlas-cli-plugin-example/@v1.0.0",
			expectVersion: true,
			expectError:   false,
		},
		{
			arg:           "/mongodb/atlas-cli-plugin-example/",
			expectVersion: false,
			expectError:   true,
		},
		{
			arg:           "mongodb@atlas-cli-plugin-example",
			expectVersion: false,
			expectError:   true,
		},
		{
			arg:           "mongodb@atlas-cli-plugin-example@1.0",
			expectVersion: false,
			expectError:   true,
		},
		{
			arg:           "invalidArgString",
			expectVersion: false,
			expectError:   true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.arg, func(t *testing.T) {
			githubRelease, err := parseGithubReleaseValues(tt.arg)
			if (err != nil) != tt.expectError {
				t.Errorf("expected error: %v, got: %v", tt.expectError, err)
			}
			if !tt.expectError {
				if githubRelease.owner != expectedOwner {
					t.Errorf("expected owner: %s, got: %s", expectedOwner, githubRelease.owner)
				}
				if githubRelease.name != expectedName {
					t.Errorf("expected name: %s, got: %s", expectedName, githubRelease.owner)
				}
				if tt.expectVersion && !expectedVersion.Equal(githubRelease.version) {
					t.Errorf("expected version: %s, got: %s", expectedVersion.String(), githubRelease.version.String())
				}
			}
		})
	}
}
