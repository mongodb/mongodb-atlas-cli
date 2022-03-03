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

	"github.com/mongodb/mongocli/internal/version"
)

func TestOutputOpts_printNewVersionAvailable(t *testing.T) {
	tests := TestCases()
	for _, tt := range tests {
		if !tt.expectNewVersion {
			continue
		}
		t.Run(fmt.Sprintf("%v / %v", tt.currentVersion, tt.release), func(t *testing.T) {
			prevVersion := version.Version
			version.Version = tt.currentVersion
			defer func() {
				version.Version = prevVersion
			}()

			var bin string
			if tt.tool == "atlascli" {
				bin = "atlas"
			} else {
				bin = tt.tool
			}

			bufOut := new(bytes.Buffer)
			v := strings.ReplaceAll(tt.release.GetTagName(), tt.tool+"/", "")
			err := NewPrinter(bufOut, tt.tool, bin).PrintNewVersionAvailable(
				v,
				"",
			)

			if err != nil {
				t.Errorf("printNewVersionAvailable() unexpected error: %v", err)
			}

			want := ""
			if tt.expectNewVersion {
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
