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

	"github.com/Masterminds/semver/v3"
	"github.com/golang/mock/gomock"
	"github.com/mongodb/mongocli/internal/mocks"
	"github.com/mongodb/mongocli/internal/version"
)

func TestOutputOpts_NewVersionAvailable(t *testing.T) {
	tests := TestCases()
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v / %v", tt.currentVersion, tt.release.GetTagName()), func(t *testing.T) {
			prevVersion := version.Version
			version.Version = tt.currentVersion
			defer func() {
				version.Version = prevVersion
			}()

			currVer, _ := semver.NewVersion(tt.currentVersion)
			*currVer, _ = currVer.SetPrerelease("")

			ctrl := gomock.NewController(t)
			mockDescriber := mocks.NewMockReleaseVersionDescriber(ctrl)
			mockStore := mocks.NewMockStore(ctrl)
			defer ctrl.Finish()

			mockStore.EXPECT().SaveLatestVersion(gomock.Any()).Return(nil)

			mockDescriber.
				EXPECT().
				LatestWithCriteria(gomock.Any(), gomock.Any(), gomock.Any()).
				Return(tt.release, nil).
				Times(1)

			finder := NewVersionFinder(mockDescriber, mockStore, tt.tool, tt.currentVersion)
			newV, err := finder.NewVersionAvailable(false)

			expectedV := strings.ReplaceAll(tt.release.GetTagName(), tt.tool+"/", "")

			if err != nil {
				t.Errorf("NewVersionAvailable() unexpected error: %v", err)
			}

			if newV != "" && (!tt.expectNewVersion || newV != expectedV) {
				t.Errorf("want: versionAvailable=%v and newV=%v got: versionAvailable=%v and newV=%v.",
					tt.expectNewVersion, expectedV, newV != "", newV)
			}
		})
	}
}

func TestOutputOpts_StoredLatestVersionAvailable(t *testing.T) {
	tests := []struct {
		tool           string
		currentVersion string
		version        string
		success        bool
	}{
		{
			tool:           version.MongoCLI,
			currentVersion: "v1.0.0",
			success:        true,
		},
		{
			tool:           version.MongoCLI,
			currentVersion: "v1",
			success:        true,
		},
		{
			tool:           version.AtlasCLI,
			currentVersion: "v2.0.0",
			success:        true,
		},
		{
			tool:           version.AtlasCLI,
			currentVersion: "v2",
			success:        false,
		},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v / should succeed: %v", tt.currentVersion, tt.success), func(t *testing.T) {
			ctrl := gomock.NewController(t)
			mockDescriber := mocks.NewMockReleaseVersionDescriber(ctrl)
			mockStore := mocks.NewMockStore(ctrl)
			defer ctrl.Finish()

			mockStore.EXPECT().LoadLatestVersion().Return(tt.currentVersion, nil)

			finder := NewVersionFinder(mockDescriber, mockStore, tt.tool, tt.currentVersion)
			_, _, err := finder.StoredLatestVersionAvailable()

			if err != nil && tt.success {
				t.Errorf("StoredLatestVersionAvailable() unexpected error: %v", err)
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
