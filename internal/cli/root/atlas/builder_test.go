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

package atlas

import (
	"bytes"
	"fmt"
	"strings"
	"testing"

	"github.com/Masterminds/semver/v3"
	"github.com/golang/mock/gomock"
	"github.com/google/go-github/v38/github"
	"github.com/mongodb/mongocli/internal/mocks"
	"github.com/mongodb/mongocli/internal/test"
	"github.com/mongodb/mongocli/internal/version"
)

func TestBuilder(t *testing.T) {
	var profile string
	test.CmdValidator(
		t,
		Builder(&profile),
		34,
		[]string{},
	)
}

func TestOutputOpts_printNewVersionAvailable(t *testing.T) {
	f := false
	v2 := "atlascli/v2.0.0"

	tests := []struct {
		currentVersion string
		version        string
		wantPrint      bool
		release        *github.RepositoryRelease
	}{
		{
			currentVersion: "v1.0.0",
			wantPrint:      true,
			release:        &github.RepositoryRelease{TagName: &v2, Prerelease: &f, Draft: &f},
		},
		{
			currentVersion: "v3.0.0",
			wantPrint:      false,
			release:        &github.RepositoryRelease{TagName: &v2, Prerelease: &f, Draft: &f},
		},
		{
			currentVersion: "v3.0.0-123",
			wantPrint:      false,
			release:        &github.RepositoryRelease{TagName: &v2, Prerelease: &f, Draft: &f},
		},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v / %v", tt.currentVersion, tt.release), func(t *testing.T) {
			prevVersion := version.Version
			version.Version = tt.currentVersion
			defer func() {
				version.Version = prevVersion
			}()

			currVer, _ := semver.NewVersion(tt.currentVersion)
			*currVer, _ = currVer.SetPrerelease("")

			ctrl := gomock.NewController(t)
			mockStore := mocks.NewMockReleaseVersionDescriber(ctrl)
			defer ctrl.Finish()

			mockStore.
				EXPECT().
				LatestWithCriteria(gomock.Any(), gomock.Any(), gomock.Any()).
				Return(tt.release, nil).
				Times(1)

			bufOut := new(bytes.Buffer)
			opts := &BuilderOpts{
				store: version.NewLatestVersionFinder(mockStore),
			}

			err := opts.store.PrintNewVersionAvailable(
				bufOut,
				tt.currentVersion,
				"atlascli",
				"atlas",
			)

			if err != nil {
				t.Errorf("printNewVersionAvailable() unexpected error: %v", err)
			}

			want := ""
			if tt.wantPrint {
				want = fmt.Sprintf(`
A new version of %s is available '%v'!
To upgrade, see: https://dochub.mongodb.org/core/atlascli-install.

To disable this alert, run "atlas config set skip_update_check true".
`, "atlascli", strings.ReplaceAll(tt.release.GetTagName(), "atlascli/", ""))
			}

			if got := bufOut.String(); got != want {
				t.Errorf("printNewVersionAvailable() got = %v, want %v", got, want)
			}
		})
	}
}

func TestOutputOpts_testIsValidTag(t *testing.T) {
	tests := []struct {
		tag     string
		isValid bool
	}{
		{"atlascli/v1.0.0", true},
		{"mongocli/v1.0.0", false},
		{"v1.0.0", false},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v_%v", tt.tag, tt.isValid), func(t *testing.T) {
			ctrl := gomock.NewController(t)
			mockStore := mocks.NewMockReleaseVersionDescriber(ctrl)
			defer ctrl.Finish()

			opts := &BuilderOpts{
				store: version.NewLatestVersionFinder(mockStore),
			}

			if result := opts.store.IsValidTagForTool(tt.tag, "atlascli"); result != tt.isValid {
				t.Errorf("got = %v, want %v", result, tt.isValid)
			}
		})
	}
}
