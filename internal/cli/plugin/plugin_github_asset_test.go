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
	"errors"
	"testing"

	"github.com/Masterminds/semver/v3"
	"github.com/google/go-github/v61/github"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/plugin"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/pointer"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_repository(t *testing.T) {
	opts := &GithubAsset{owner: "repoOwner", name: "repoName"}

	assert.Equal(t, opts.owner+"/"+opts.name, opts.repository())
}

func Test_getIDsForOSArch(t *testing.T) {
	tests := []struct {
		name                string
		pluginAssets        []*github.ReleaseAsset
		expectedAssetID     int64
		expectedSignatureID int64
		expectedPubKeyID    int64
		expectError         bool
		os                  string
		arch                string
	}{
		{
			name: "Valid asset linux amd64",
			pluginAssets: []*github.ReleaseAsset{
				{
					ID:          github.Int64(1),
					Name:        github.String("plugin_linux_amd64.tar.gz"),
					ContentType: github.String("application/gzip"),
				},
				{
					ID:          github.Int64(2),
					Name:        github.String("plugin_linux_amd64.tar.gz.sig"),
					ContentType: github.String("application/pgp"),
				},
				{
					ID:          github.Int64(3),
					Name:        github.String("signature.asc"),
					ContentType: github.String("application/text/plain"),
				},
			},
			expectedAssetID:     1,
			expectedSignatureID: 2,
			expectedPubKeyID:    3,
			expectError:         false,
			os:                  "linux",
			arch:                "amd64",
		},
		{
			name: "Valid asset windows amd64",
			pluginAssets: []*github.ReleaseAsset{
				{
					ID:          github.Int64(2),
					Name:        github.String("plugin_windows_amd64.zip"),
					ContentType: github.String("application/zip"),
				},
				{
					ID:          github.Int64(3),
					Name:        github.String("plugin_windows_amd64.zip.sig"),
					ContentType: github.String("application/pgp"),
				},
				{
					ID:          github.Int64(4),
					Name:        github.String("signature.asc"),
					ContentType: github.String("application/text/plain"),
				},
			},
			expectedAssetID:     2,
			expectedSignatureID: 3,
			expectedPubKeyID:    4,
			expectError:         false,
			os:                  "windows",
			arch:                "amd64",
		},
		{
			name: "Valid asset darwin arm64",
			pluginAssets: []*github.ReleaseAsset{
				{
					ID:          github.Int64(3),
					Name:        github.String("plugin_darwin_arm64.tar.gz"),
					ContentType: github.String("application/gzip"),
				},
				{
					ID:          github.Int64(4),
					Name:        github.String("plugin_darwin_arm64.tar.gz.sig"),
					ContentType: github.String("application/pgp"),
				},
				{
					ID:          github.Int64(5),
					Name:        github.String("signature.asc"),
					ContentType: github.String("application/text/plain"),
				},
			},
			expectedAssetID:     3,
			expectedSignatureID: 4,
			expectedPubKeyID:    5,
			expectError:         false,
			os:                  "darwin",
			arch:                "arm64",
		},
		{
			name: "Valid asset with x86_64 linux",
			pluginAssets: []*github.ReleaseAsset{
				{
					ID:          github.Int64(4),
					Name:        github.String("plugin_linux_x86_64.tar.gz"),
					ContentType: github.String("application/gzip"),
				},
				{
					ID:          github.Int64(5),
					Name:        github.String("plugin_linux_x86_64.tar.gz.sig"),
					ContentType: github.String("application/pgp"),
				},
				{
					ID:          github.Int64(6),
					Name:        github.String("signature.asc"),
					ContentType: github.String("application/text/plain"),
				},
			},
			expectedAssetID:     4,
			expectedSignatureID: 5,
			expectedPubKeyID:    6,
			expectError:         false,
			os:                  "linux",
			arch:                "amd64",
		},
		{
			name: "Valid asset with x86_64 darwin",
			pluginAssets: []*github.ReleaseAsset{
				{
					ID:          github.Int64(5),
					Name:        github.String("plugin_darwin_x86_64.tar.gz"),
					ContentType: github.String("application/gzip"),
				},
				{
					ID:          github.Int64(6),
					Name:        github.String("plugin_darwin_x86_64.tar.gz.sig"),
					ContentType: github.String("application/pgp"),
				},
				{
					ID:          github.Int64(7),
					Name:        github.String("signature.asc"),
					ContentType: github.String("application/text/plain"),
				},
			},
			expectedAssetID:     5,
			expectedSignatureID: 6,
			expectedPubKeyID:    7,
			expectError:         false,
			os:                  "darwin",
			arch:                "amd64",
		},
		{
			name: "Valid asset with aarch64 linux",
			pluginAssets: []*github.ReleaseAsset{
				{
					ID:          github.Int64(6),
					Name:        github.String("plugin_linux_aarch64.tar.gz"),
					ContentType: github.String("application/gzip"),
				},
				{
					ID:          github.Int64(7),
					Name:        github.String("plugin_linux_aarch64.tar.gz.sig"),
					ContentType: github.String("application/pgp"),
				},
				{
					ID:          github.Int64(8),
					Name:        github.String("signature.asc"),
					ContentType: github.String("application/text/plain"),
				},
			},
			expectedAssetID:     6,
			expectedSignatureID: 7,
			expectedPubKeyID:    8,
			expectError:         false,
			os:                  "linux",
			arch:                "arm64",
		},
		{
			name: "Valid asset with aarch64 darwin",
			pluginAssets: []*github.ReleaseAsset{
				{
					ID:          github.Int64(7),
					Name:        github.String("plugin_darwin_aarch64.tar.gz"),
					ContentType: github.String("application/gzip"),
				},
				{
					ID:          github.Int64(8),
					Name:        github.String("plugin_darwin_aarch64.tar.gz.sig"),
					ContentType: github.String("application/pgp"),
				},
				{
					ID:          github.Int64(9),
					Name:        github.String("signature.asc"),
					ContentType: github.String("application/text/plain"),
				},
			},
			expectedAssetID:     7,
			expectedSignatureID: 8,
			expectedPubKeyID:    9,
			expectError:         false,
			os:                  "darwin",
			arch:                "arm64",
		},
		{
			name: "No matching asset",
			pluginAssets: []*github.ReleaseAsset{
				{
					ID:          github.Int64(8),
					Name:        github.String("plugin_invalid_assetname.tar.gz"),
					ContentType: github.String("application/gzip"),
				},
			},
			expectedAssetID:     0,
			expectedSignatureID: 0,
			expectedPubKeyID:    0,
			expectError:         true,
			os:                  "linux",
			arch:                "amd64",
		},
		{
			name: "Non-gzip asset",
			pluginAssets: []*github.ReleaseAsset{
				{
					ID:          github.Int64(9),
					Name:        github.String("plugin_linux_amd64.json"),
					ContentType: github.String("application/json"),
				},
			},
			expectedAssetID:     0,
			expectedSignatureID: 0,
			expectedPubKeyID:    0,
			expectError:         true,
			os:                  "linux",
			arch:                "amd64",
		},
		{
			name: "Zip asset",
			pluginAssets: []*github.ReleaseAsset{
				{
					ID:          github.Int64(10),
					Name:        github.String("plugin_linux_amd64.zip"),
					ContentType: github.String("application/zip"),
				},
				{
					ID:          github.Int64(11),
					Name:        github.String("plugin_linux_amd64.zip.sig"),
					ContentType: github.String("application/pgp"),
				},
				{
					ID:          github.Int64(12),
					Name:        github.String("signature.asc"),
					ContentType: github.String("application/text/plain"),
				},
			},
			expectedAssetID:     10,
			expectedSignatureID: 11,
			expectedPubKeyID:    12,
			expectError:         false,
			os:                  "linux",
			arch:                "amd64",
		},
		{
			name: "No signature file present",
			pluginAssets: []*github.ReleaseAsset{
				{
					ID:          github.Int64(11),
					Name:        github.String("plugin_linux_amd64.zip"),
					ContentType: github.String("application/zip"),
				},
			},
			expectedAssetID:     11,
			expectedSignatureID: 0,
			expectedPubKeyID:    0,
			expectError:         false,
			os:                  "linux",
			arch:                "amd64",
		},
		{
			name: "No valid public key present",
			pluginAssets: []*github.ReleaseAsset{
				{
					ID:          github.Int64(12),
					Name:        github.String("plugin_linux_amd64.zip"),
					ContentType: github.String("application/zip"),
				},
				{
					ID:          github.Int64(13),
					Name:        github.String("plugin_linux_amd64.zip.sig"),
					ContentType: github.String("application/pgp"),
				},
			},
			expectedAssetID:     0,
			expectedSignatureID: 0,
			expectedPubKeyID:    0,
			expectError:         true,
			os:                  "linux",
			arch:                "amd64",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			opts := &GithubAsset{}

			assetID, signatureID, pubKeyID, err := opts.getIDsForOSArch(tt.pluginAssets, tt.os, tt.arch)
			if (err != nil) != tt.expectError {
				t.Errorf("expected error: %v, got: %v", tt.expectError, err)
			}
			if assetID != tt.expectedAssetID {
				t.Errorf("expected asset ID: %d, got: %d", tt.expectedAssetID, assetID)
			}
			if signatureID != tt.expectedSignatureID {
				t.Errorf("expected asset ID: %d, got: %d", tt.expectedSignatureID, signatureID)
			}
			if pubKeyID != tt.expectedPubKeyID {
				t.Errorf("expected asset ID: %d, got: %d", tt.expectedPubKeyID, pubKeyID)
			}
		})
	}
}

func Test_parseGithubRepoValues(t *testing.T) {
	const (
		expectedOwner = "mongodb"
		expectedName  = "atlas-cli-plugin-example"
	)
	var (
		v1, _       = semver.NewVersion("1.0.0")
		v1pre, _    = semver.NewVersion("1.0.0-prerelease")
		v1pseudo, _ = semver.NewVersion("1.0.0-beta+very-meta")
	)

	tests := []struct {
		arg             string
		expectedVersion *semver.Version
		expectError     bool
	}{
		{
			arg:             "mongodb/atlas-cli-plugin-example",
			expectedVersion: nil,
			expectError:     false,
		},
		{
			arg:             "mongodb/atlas-cli-plugin-example@1.0.0",
			expectedVersion: v1,
			expectError:     false,
		},
		{
			arg:             "mongodb/atlas-cli-plugin-example@v1.0.0",
			expectedVersion: v1,
			expectError:     false,
		},
		{
			arg:             "mongodb/atlas-cli-plugin-example@1.0.0-prerelease",
			expectedVersion: v1pre,
			expectError:     false,
		},
		{
			arg:             "mongodb/atlas-cli-plugin-example@1.0.0-beta+very-meta",
			expectedVersion: v1pseudo,
			expectError:     false,
		},
		{
			arg:             "mongodb/atlas-cli-plugin-example@",
			expectedVersion: nil,
			expectError:     true,
		},
		{
			arg:             "mongodb/atlas-cli-plugin-example/",
			expectedVersion: nil,
			expectError:     false,
		},
		{
			arg:             "mongodb/atlas-cli-plugin-example/@v1",
			expectedVersion: v1,
			expectError:     false,
		},
		{
			arg:             "https://github.com/mongodb/atlas-cli-plugin-example",
			expectedVersion: nil,
			expectError:     false,
		},
		{
			arg:             "https://github.com/mongodb/atlas-cli-plugin-example@v1.0",
			expectedVersion: v1,
			expectError:     false,
		},
		{
			arg:             "github.com/mongodb/atlas-cli-plugin-example/",
			expectedVersion: nil,
			expectError:     false,
		},
		{
			arg:             "github.com/mongodb/atlas-cli-plugin-example/@v1.0.0",
			expectedVersion: v1,
			expectError:     false,
		},
		{
			arg:             "/mongodb/atlas-cli-plugin-example/",
			expectedVersion: nil,
			expectError:     true,
		},
		{
			arg:             "mongodb@atlas-cli-plugin-example",
			expectedVersion: nil,
			expectError:     true,
		},
		{
			arg:             "mongodb@atlas-cli-plugin-example@1.0",
			expectedVersion: nil,
			expectError:     true,
		},
		{
			arg:             "invalidArgString",
			expectedVersion: nil,
			expectError:     true,
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
				if tt.expectedVersion != nil && !tt.expectedVersion.Equal(githubRelease.version) {
					t.Errorf("expected version: %s, got: %s", tt.expectedVersion.String(), githubRelease.version.String())
				}

				if tt.expectedVersion == nil && githubRelease.version != nil {
					t.Errorf("expected version to be nil, got: %s", githubRelease.version.String())
				}
			}
		})
	}
}

func Test_createGithubAssetFromPlugin(t *testing.T) {
	var expectedVersion, _ = semver.NewVersion("1.0.0")

	tests := []struct {
		name        string
		plugin      *plugin.Plugin
		expectedErr error
		expected    *GithubAsset
	}{
		{
			name: "Plugin with GitHub values",
			plugin: &plugin.Plugin{
				Github: &plugin.Github{
					Owner: "test-owner",
					Name:  "test-repo",
				},
			},
			expectedErr: nil,
			expected: &GithubAsset{
				owner:   "test-owner",
				name:    "test-repo",
				version: expectedVersion,
			},
		},
		{
			name:        "Plugin without GitHub values",
			plugin:      &plugin.Plugin{},
			expectedErr: errCreatePluginAssetFromPlugin,
			expected:    nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := createGithubAssetFromPlugin(tt.plugin, expectedVersion)

			if !errors.Is(err, tt.expectedErr) {
				t.Errorf("expected error: %v, got: %v", tt.expectedErr, err)
			}

			if err == nil && tt.expected != nil {
				if got.owner != tt.expected.owner || got.name != tt.expected.name || !got.version.Equal(expectedVersion) {
					t.Errorf("expected: %v, got: %v", tt.expected, got)
				}
			}
		})
	}
}

func Test_getPluginDirectoryName(t *testing.T) {
	githubAsset := &GithubAsset{owner: "owner", name: "name"}
	require.Equal(t, "owner@name", githubAsset.getPluginDirectoryName())
}

func Test_getLatestStableRelease(t *testing.T) {
	tests := []struct {
		name     string
		releases []*github.RepositoryRelease
		expected *github.RepositoryRelease
	}{
		{
			name: "Single valid value",
			releases: []*github.RepositoryRelease{
				{
					TagName: pointer.Get("v1.0.0"),
				},
			},
			expected: &github.RepositoryRelease{
				TagName: pointer.Get("v1.0.0"),
			},
		},
		{
			name: "Single invalid value",
			releases: []*github.RepositoryRelease{
				{
					TagName: pointer.Get("test"),
				},
			},
			expected: nil,
		},
		{
			name: "Single valid pre-release value",
			releases: []*github.RepositoryRelease{
				{
					TagName: pointer.Get("v1.0.0-pre"),
				},
			},
			expected: nil,
		},
		{
			name: "Multiple",
			releases: []*github.RepositoryRelease{
				{
					TagName: pointer.Get("v2.0.0-pre"),
				},
				{
					TagName: pointer.Get("v2.0.0-beta"),
				},
				{
					TagName: pointer.Get("v1.2.1"),
				},
				{
					TagName: pointer.Get("v1.2.0"),
				},
				{
					TagName: pointer.Get("v1.1.0"),
				},
				{
					TagName: pointer.Get("v1.0.1"),
				},
				{
					TagName: pointer.Get("v1.0.0"),
				},
			},
			expected: &github.RepositoryRelease{
				TagName: pointer.Get("v1.2.1"),
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := getLatestStableRelease(tt.releases)

			assert.Equal(t, tt.expected, actual)
		})
	}
}
