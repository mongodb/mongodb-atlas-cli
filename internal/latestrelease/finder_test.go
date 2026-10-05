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

package latestrelease

import (
	"fmt"
	"strings"
	"testing"

	"github.com/google/go-github/v61/github"
	"github.com/mongodb/atlas-cli-core/config"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/mocks"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/version"
	"github.com/spf13/afero"
	"go.uber.org/mock/gomock"
)

const (
	versionV1      = "v1.0.0"
	versionAtlasV1 = "atlascli/v1.0.0"
	versionV2      = "v2.0.0"
	versionAtlasV2 = "atlascli/v2.0.0"
	versionV3      = "v3.0.0"
	versionV3Build = "v3.0.0-123"
)

type testCase struct {
	currentVersion   string
	expectNewVersion bool
	release          *github.RepositoryRelease
}

func testCases() []testCase {
	f := false
	atlasV := versionAtlasV2
	bareV := versionV2

	tests := []testCase{
		{
			currentVersion:   versionV1,
			expectNewVersion: true,
			release:          &github.RepositoryRelease{TagName: &atlasV, Prerelease: &f, Draft: &f},
		},
		{
			currentVersion:   versionAtlasV1,
			expectNewVersion: true,
			release:          &github.RepositoryRelease{TagName: &atlasV, Prerelease: &f, Draft: &f},
		},
		{
			currentVersion:   versionV1,
			expectNewVersion: true,
			release:          &github.RepositoryRelease{TagName: &bareV, Prerelease: &f, Draft: &f},
		},
		{
			currentVersion:   versionAtlasV1,
			expectNewVersion: true,
			release:          &github.RepositoryRelease{TagName: &bareV, Prerelease: &f, Draft: &f},
		},
		{
			currentVersion:   versionV2,
			expectNewVersion: false,
			release:          &github.RepositoryRelease{TagName: &atlasV, Prerelease: &f, Draft: &f},
		},
		{
			currentVersion:   versionV2,
			expectNewVersion: false,
			release:          &github.RepositoryRelease{TagName: &bareV, Prerelease: &f, Draft: &f},
		},
		{
			currentVersion:   versionV3,
			expectNewVersion: false,
			release:          &github.RepositoryRelease{TagName: &atlasV, Prerelease: &f, Draft: &f},
		},
		{
			currentVersion:   versionV3Build,
			expectNewVersion: false,
			release:          &github.RepositoryRelease{TagName: &atlasV, Prerelease: &f, Draft: &f},
		},
	}
	return tests
}

func TestOutputOpts_Find_NoCache(t *testing.T) {
	tests := testCases()
	for _, tt := range tests {
		prevVersion := version.Version
		version.Version = tt.currentVersion
		t.Cleanup(func() {
			version.Version = prevVersion
		})
		t.Run(fmt.Sprintf("%v / %v", tt.currentVersion, tt.release.GetTagName()), func(t *testing.T) {
			ctrl := gomock.NewController(t)
			mockDescriber := mocks.NewMockReleaseVersionDescriber(ctrl)

			mockDescriber.
				EXPECT().
				LatestWithCriteria(gomock.Any(), gomock.Any()).
				Return(tt.release, nil).
				Times(1)

			f, err := NewVersionFinder(afero.NewMemMapFs(), mockDescriber)
			if err != nil {
				t.Errorf("NewVersionFinder() unexpected error: %v", err)
			}

			newV, err := f.Find()
			if err != nil {
				t.Errorf("Find() unexpected error: %v", err)
			}

			expectedV := strings.ReplaceAll(tt.release.GetTagName(), config.AtlasCLI+"/", "")

			if newV != nil && (!tt.expectNewVersion || newV.Version != expectedV) {
				t.Errorf("want: versionAvailable=%v and newV=%v got: versionAvailable=%v and newV=%v.",
					tt.expectNewVersion, expectedV, newV != nil, newV)
			}
		})
	}
}

func TestOutputOpts_testIsValidTag(t *testing.T) {
	tests := []struct {
		tag     string
		isValid bool
	}{
		// Legacy format.
		{versionAtlasV1, true},
		{"atlascli/v1.0.0-rc0", true},
		// New bare format.
		{versionV1, true},
		{"v2.3.4", true},
		// Other tools' tags. mongocli shares this repo's release feed and its
		// latest tag (v2.x) outranks atlascli's (v1.x), so it must never match.
		{"mongocli/v1.0.0", false},
		{"mongocli/v2.0.3", false},
		// Tags that do not carry a version.
		{"nightly", false},
		{"atlascli/nightly", false},
		{"2024", false},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v_%v", tt.tag, tt.isValid), func(t *testing.T) {
			if result := isValidTagForTool(tt.tag); result != tt.isValid {
				t.Errorf("got = %v, want %v", result, tt.isValid)
			}
		})
	}
}
