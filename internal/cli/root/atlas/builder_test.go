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
	"testing"
	"time"

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
	tests := []struct {
		currentVersion string
		latestVersion  *version.ReleaseInformation
		wantPrint      bool
	}{
		{
			currentVersion: "v1.0.0",
			latestVersion:  &version.ReleaseInformation{Version: "v2.0.0", PublishedAt: time.Now()},
			wantPrint:      true,
		},
		{
			currentVersion: "v1.0.0",
			latestVersion:  &version.ReleaseInformation{Version: "v1.0.0", PublishedAt: time.Now()},
			wantPrint:      false,
		},
		{
			currentVersion: "v1.0.0-123",
			latestVersion:  &version.ReleaseInformation{Version: "v1.0.0", PublishedAt: time.Now()},
			wantPrint:      false,
		},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v / %v", tt.currentVersion, tt.latestVersion), func(t *testing.T) {
			prevVersion := version.Version
			f := false
			version.Version = tt.currentVersion
			defer func() {
				version.Version = prevVersion
			}()

			tagName := "atlascli/" + tt.latestVersion.Version

			releases := []*github.RepositoryRelease{
				{
					TagName:    &tagName,
					Draft:      &f,
					Prerelease: &f,
				},
			}

			currVer, _ := semver.NewVersion(tt.currentVersion)
			*currVer, _ = currVer.SetPrerelease("")

			ctrl := gomock.NewController(t)
			mockStore := mocks.NewMockReleaseVersionDescriber(ctrl)
			defer ctrl.Finish()

			mockStore.
				EXPECT().
				AllVersions().
				Return(releases, nil).
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
`, "atlascli", tt.latestVersion.Version)
			}

			if got := bufOut.String(); got != want {
				t.Errorf("printNewVersionAvailable() got = %v, want %v", got, want)
			}
		})
	}
}
