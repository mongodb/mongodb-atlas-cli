// Copyright 2026 MongoDB Inc
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

package cmdtree

import "testing"

var (
	oneOperation  = []ApiOperation{{Version: "v20250312", OperationID: "CreateOrg"}}
	twoOperations = []ApiOperation{
		{Version: "v20250312", OperationID: "CreateOrg"},
		{Version: "v20250312", OperationID: "CreateCluster"},
	}
)

func TestCategory(t *testing.T) {
	tests := []struct {
		name string
		cmd  Command
		want Category
	}{
		{"a command that only groups subcommands", Command{Path: "atlas clusters"}, ""},
		{"no operation", Command{Path: "atlas config list", Run: "run"}, CategoryMeta},
		{"one operation", Command{Path: "atlas organizations create", Run: "run", ApiOperationIds: oneOperation}, CategoryWorkflow},
		{"two operations", Command{Path: "atlas clusters create", Run: "run", ApiOperationIds: twoOperations}, CategoryDirect},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := category(&tt.cmd); got != tt.want {
				t.Errorf("category() = %q, want %q", got, tt.want)
			}
		})
	}
}

func TestLegacyCategory(t *testing.T) {
	tests := []struct {
		name string
		cmd  Command
		want LegacyCategory
	}{
		{"the api subtree root", Command{Path: "atlas api"}, LegacyL1},
		{"a command in the api subtree", Command{Path: "atlas api accessTracking getAccessHistoryCluster", Run: "run"}, LegacyL1},
		{"an api command with one operation is still L1", Command{Path: "atlas api accessTracking getAccessHistoryCluster", Run: "run", ApiOperationIds: oneOperation}, LegacyL1},
		{"outside it with one operation", Command{Path: "atlas organizations create", Run: "run", ApiOperationIds: oneOperation}, LegacyL2},
		{"outside it with two operations", Command{Path: "atlas clusters create", Run: "run", ApiOperationIds: twoOperations}, LegacyL3},
		{"outside it with none", Command{Path: "atlas config list", Run: "run"}, LegacyL3},
		{"a command that only groups", Command{Path: "atlas clusters"}, LegacyL3},
		{"a path that only shares a prefix", Command{Path: "atlas apiKeys create", Run: "run", ApiOperationIds: oneOperation}, LegacyL2},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := legacyCategory(&tt.cmd); got != tt.want {
				t.Errorf("legacyCategory() = %q, want %q", got, tt.want)
			}
		})
	}
}
