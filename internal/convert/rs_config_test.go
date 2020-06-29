// Copyright 2020 MongoDB Inc
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

package convert

import "testing"

func TestProtocolVersion(t *testing.T) {
	testCases := map[string]struct {
		config          *RSConfig
		protocolVersion string
	}{
		"empty fcv": {
			config:          &RSConfig{},
			protocolVersion: "",
		},
		"post 4.0": {
			config:          &RSConfig{FCVersion: "4.0"},
			protocolVersion: "1",
		},
		"pre 4.0": {
			config:          &RSConfig{FCVersion: "3.6"},
			protocolVersion: "0",
		},
	}
	for name, tc := range testCases {
		m := tc.config
		expected := tc.protocolVersion
		t.Run(name, func(t *testing.T) {
			ver, err := m.protocolVer()
			if err != nil {
				t.Fatalf("protocolVer() unexpected error: %v\n", err)
			}
			if ver != expected {
				t.Errorf("protocolVer() expected: %s but got: %s", expected, ver)
			}
		})
	}
}
