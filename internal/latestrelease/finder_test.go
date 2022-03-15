// Copyright 2022 MongoDB Inc
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
// +build unit

package latestrelease

import (
	"fmt"
	"strings"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/google/go-github/v42/github"
	"github.com/mongodb/mongocli/internal/config"
	"github.com/mongodb/mongocli/internal/mocks"
	"github.com/mongodb/mongocli/internal/version"
	"github.com/spf13/afero"
)

type testCase struct {
	tool             string
	currentVersion   string
	expectNewVersion bool
	release          *github.RepositoryRelease
}

func testCases() []testCase {
	f := false
	atlasV := "atlascli/v2.0.0"
	mcliV := "mongocli/v2.0.0"
	mcliOldV := "v2.0.0"

	tests := []testCase{
		{
			tool:             "atlascli",
			currentVersion:   "v1.0.0",
			expectNewVersion: true,
			release:          &github.RepositoryRelease{TagName: &atlasV, Prerelease: &f, Draft: &f},
		},
		{
			tool:             "atlascli",
			currentVersion:   "atlascli/v1.0.0",
			expectNewVersion: true,
			release:          &github.RepositoryRelease{TagName: &atlasV, Prerelease: &f, Draft: &f},
		},
		{
			tool:             "atlascli",
			currentVersion:   "v2.0.0",
			expectNewVersion: false,
			release:          &github.RepositoryRelease{TagName: &atlasV, Prerelease: &f, Draft: &f},
		},
		{
			tool:             "atlascli",
			currentVersion:   "v3.0.0",
			expectNewVersion: false,
			release:          &github.RepositoryRelease{TagName: &atlasV, Prerelease: &f, Draft: &f},
		},
		{
			tool:             "atlascli",
			currentVersion:   "v3.0.0-123",
			expectNewVersion: false,
			release:          &github.RepositoryRelease{TagName: &atlasV, Prerelease: &f, Draft: &f},
		},
		{
			tool:             "mongocli",
			currentVersion:   "v1.0.0",
			expectNewVersion: true,
			release:          &github.RepositoryRelease{TagName: &mcliOldV, Prerelease: &f, Draft: &f},
		},
		{
			tool:             "mongocli",
			currentVersion:   "v2.0.0",
			expectNewVersion: false,
			release:          &github.RepositoryRelease{TagName: &mcliOldV, Prerelease: &f, Draft: &f},
		},
		{
			tool:             "mongocli",
			currentVersion:   "mongocli/v1.0.0",
			expectNewVersion: true,
			release:          &github.RepositoryRelease{TagName: &mcliOldV, Prerelease: &f, Draft: &f},
		},
		{
			tool:             "mongocli",
			currentVersion:   "v1.0.0",
			expectNewVersion: true,
			release:          &github.RepositoryRelease{TagName: &mcliV, Prerelease: &f, Draft: &f},
		},
		{
			tool:             "mongocli",
			currentVersion:   "v3.0.0",
			expectNewVersion: false,
			release:          &github.RepositoryRelease{TagName: &mcliOldV, Prerelease: &f, Draft: &f},
		},
		{
			tool:             "mongocli",
			currentVersion:   "v3.0.0-123",
			expectNewVersion: false,
			release:          &github.RepositoryRelease{TagName: &mcliV, Prerelease: &f, Draft: &f},
		},
	}
	return tests
}

func TestOutputOpts_Find_NoCache(t *testing.T) {
	tests := testCases()
	for _, tt := range tests {
		config.ToolName = tt.tool
		prevVersion := version.Version
		version.Version = tt.currentVersion
		defer func() {
			version.Version = prevVersion
		}()
		t.Run(fmt.Sprintf("%v / %v", tt.currentVersion, tt.release.GetTagName()), func(t *testing.T) {
			ctrl := gomock.NewController(t)
			mockDescriber := mocks.NewMockReleaseVersionDescriber(ctrl)
			defer ctrl.Finish()

			mockDescriber.
				EXPECT().
				LatestWithCriteria(gomock.Any(), gomock.Any(), gomock.Any()).
				Return(tt.release, nil).
				Times(1)

			finder, err := NewVersionFinder(afero.NewMemMapFs(), mockDescriber)
			if err != nil {
				t.Errorf("NewVersionFinder() unexpected error: %v", err)
			}

			newV, err := finder.Find()
			if err != nil {
				t.Errorf("Find() unexpected error: %v", err)
			}

			expectedV := strings.ReplaceAll(tt.release.GetTagName(), tt.tool+"/", "")

			if newV != nil && (!tt.expectNewVersion || newV.Version != expectedV) {
				t.Errorf("want: versionAvailable=%v and newV=%v got: versionAvailable=%v and newV=%v.",
					tt.expectNewVersion, expectedV, newV != nil, newV)
			}
		})
	}
}

func TestOutputOpts_testIsValidTag(t *testing.T) {
	tests := []struct {
		tool    string
		tag     string
		isValid bool
	}{
		{"atlascli", "atlascli/v1.0.0", true},
		{"atlascli", "mongocli/v1.0.0", false},
		{"atlascli", "v1.0.0", false},
		{"mongocli", "atlascli/v1.0.0", false},
		{"mongocli", "mongocli/v1.0.0", true},
		{"mongocli", "v1.0.0", true},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v_%v", tt.tag, tt.isValid), func(t *testing.T) {
			if result := isValidTagForTool(tt.tag, tt.tool); result != tt.isValid {
				t.Errorf("got = %v, want %v", result, tt.isValid)
			}
		})
	}
}
