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

package root

import (
	"bytes"
	"fmt"
	"testing"

	"github.com/google/go-github/v61/github"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/latestrelease"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/mocks"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/version"
	"github.com/spf13/afero"
	"go.uber.org/mock/gomock"
)

func TestOutputOpts_notifyIfApplicable(t *testing.T) {
	f := false
	atlasV := "atlascli/v2.0.0"
	tests := []struct {
		currentVersion   string
		expectNewVersion bool
		release          *github.RepositoryRelease
	}{
		{
			currentVersion:   "atlascli/v1.0.0",
			expectNewVersion: true,
			release:          &github.RepositoryRelease{TagName: &atlasV, Prerelease: &f, Draft: &f},
		},
		{
			currentVersion:   "v3.0.0",
			expectNewVersion: false,
			release:          &github.RepositoryRelease{TagName: &atlasV, Prerelease: &f, Draft: &f},
		},
		{
			currentVersion:   "v2.0.0",
			expectNewVersion: false,
			release:          &github.RepositoryRelease{TagName: &atlasV, Prerelease: &f, Draft: &f},
		},
		{
			currentVersion:   "v2.0.0-123",
			expectNewVersion: false,
			release:          &github.RepositoryRelease{TagName: &atlasV, Prerelease: &f, Draft: &f},
		},
		{
			currentVersion:   "v3.0.0-123",
			expectNewVersion: false,
			release:          nil,
		},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v/%v", tt.currentVersion, tt.release), func(t *testing.T) {
			prevVersion := version.Version
			version.Version = tt.currentVersion
			t.Cleanup(func() {
				version.Version = prevVersion
			})

			ctrl := gomock.NewController(t)
			mockDescriber := mocks.NewMockReleaseVersionDescriber(ctrl)

			mockDescriber.
				EXPECT().
				LatestWithCriteria(gomock.Any(), gomock.Any()).
				Return(tt.release, nil).
				Times(1)

			bufOut := new(bytes.Buffer)
			fs := afero.NewMemMapFs()
			finder, _ := latestrelease.NewVersionFinder(fs, mockDescriber)

			notifier := &Notifier{
				currentVersion: latestrelease.VersionFromTag(version.Version),
				finder:         finder,
				filesystem:     fs,
				writer:         bufOut,
			}

			if err := notifier.notifyIfApplicable(false); err != nil {
				t.Fatalf("notifyIfApplicable() unexpected error:%v", err)
			}

			v := ""
			if tt.release != nil {
				v = latestrelease.VersionFromTag(tt.release.GetTagName())
			}

			want := ""
			if tt.expectNewVersion {
				want = fmt.Sprintf(`
A new version of atlascli is available %q!
To upgrade, see: https://dochub.mongodb.org/core/install-atlas-cli

To disable this alert, run "atlas config set skip_update_check true"
`, v)
			}

			if got := bufOut.String(); got != want {
				t.Errorf("notifyIfApplicable() got = %v, want %v", got, want)
			}
		})
	}
}
