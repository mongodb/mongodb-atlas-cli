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
	"bytes"
	"fmt"
	"strings"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/mongodb/mongocli/internal/config"
	"github.com/mongodb/mongocli/internal/mocks"
	"github.com/spf13/afero"
)

func TestOutputOpts_CheckAvailable(t *testing.T) {
	tests := TestCases()
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v / %v", tt.currentVersion, tt.release.GetTagName()), func(t *testing.T) {
			bin := config.BinName()
			ctrl := gomock.NewController(t)
			mockDescriber := mocks.NewMockReleaseVersionDescriber(ctrl)
			mockStore := mocks.NewMockStore(ctrl)
			defer ctrl.Finish()

			mockStore.
				EXPECT().
				LoadLatestVersion().
				Return("", nil).
				Times(1)

			mockStore.EXPECT().SaveLatestVersion(gomock.Any()).Return(nil)

			mockDescriber.
				EXPECT().
				LatestWithCriteria(gomock.Any(), gomock.Any(), gomock.Any()).
				Return(tt.release, nil).
				Times(1)

			bufOut := new(bytes.Buffer)

			printer := NewPrinter(bufOut, tt.tool, bin)
			finder := NewVersionFinder(mockDescriber, mockStore, tt.tool, tt.currentVersion)
			checker := newCheckerForTest(tt.currentVersion, tt.tool, printer, finder, afero.NewMemMapFs())

			err := checker.CheckAvailable()
			if err != nil {
				t.Errorf("NewVersionAvailable() unexpected error: %v", err)
			}

			want := ""
			if tt.expectNewVersion {
				v := strings.ReplaceAll(tt.release.GetTagName(), tt.tool+"/", "")
				want = fmt.Sprintf(`
A new version of %s is available '%v'!
To upgrade, see: https://dochub.mongodb.org/core/%s-install.

To disable this alert, run "%s config set skip_update_check true".
`, tt.tool, v, tt.tool, bin)
			}

			if got := bufOut.String(); got != want {
				t.Errorf("printNewVersionAvailable() got = %v, want %v", got, want)
			}
		})
	}
}
