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

import (
	"sync"
	"testing"
)

const clustersList = "atlas clusters list"

// Building the tree runs the CLI and a call graph over it, which is too slow to
// repeat for every test.
var (
	buildOnce sync.Once
	buildTree Command
	buildErr  error
)

func atlasCommands(t *testing.T) Command {
	t.Helper()
	buildOnce.Do(func() { buildTree, buildErr = Build() })
	if buildErr != nil {
		t.Fatalf("Build() error: %v", buildErr)
	}
	return buildTree
}

// indexByPath flattens the tree so a test can look a command up by its path.
func indexByPath(root Command) map[string]Command {
	paths := make(map[string]Command)
	each(&root, func(cmd *Command) { paths[cmd.Path] = *cmd })
	return paths
}

func calls(values []ApiOperation, want ApiOperation) bool {
	for _, value := range values {
		if value == want {
			return true
		}
	}
	return false
}

func TestBuild(t *testing.T) {
	if testing.Short() {
		t.Skip("builds the CLI and a call graph over it")
	}
	paths := indexByPath(atlasCommands(t))

	// The api subtree is generated from the OpenAPI spec at runtime, so finding
	// it proves the tree came from the running program and not from source.
	if _, ok := paths["atlas api"]; !ok {
		t.Error("atlas api missing from the tree")
	}

	// atlas clusters list branches to flex clusters at runtime, so it reaches
	// the request for each of them.
	list, ok := paths[clustersList]
	if !ok {
		t.Fatalf("%s missing from the tree", clustersList)
	}
	if list.Run == "" {
		t.Errorf("%s has no run function", clustersList)
	}
	for _, want := range []ApiOperation{
		{Version: "v20240530", OperationID: "ListClusters"},
		{Version: "v20250312", OperationID: "ListClusters"},
		{Version: "v20250312", OperationID: "ListFlexClusters"},
	} {
		if !calls(list.ApiOperationIds, want) {
			t.Errorf("%s ApiOperationIds = %v, want %+v", clustersList, list.ApiOperationIds, want)
		}
	}
}

// TestBuildSharedHelper pins the walk to the callback each command passes.
// cli.DeleteOpts.Delete takes the function to call as an argument and calls it
// through a type switch, so an analysis that ignores the argument reports every
// delete in the program for every command that uses the helper.
func TestBuildSharedHelper(t *testing.T) {
	if testing.Short() {
		t.Skip("builds the CLI and a call graph over it")
	}
	paths := indexByPath(atlasCommands(t))

	del, ok := paths["atlas accessLists delete"]
	if !ok {
		t.Fatal("atlas accessLists delete missing from the tree")
	}
	want := ApiOperation{Version: "v20250312", OperationID: "DeleteAccessListEntry"}
	if !calls(del.ApiOperationIds, want) {
		t.Errorf("%s ApiOperationIds = %v, want %+v", del.Path, del.ApiOperationIds, want)
	}
	for _, unwanted := range []string{"DeleteAlertConfig", "DeleteCluster", "DeleteOrg"} {
		for _, call := range del.ApiOperationIds {
			if call.OperationID == unwanted {
				t.Errorf("%s ApiOperationIds = %v, want no %q", del.Path, del.ApiOperationIds, unwanted)
			}
		}
	}
}

func TestBuildCategories(t *testing.T) {
	if testing.Short() {
		t.Skip("builds the CLI and a call graph over it")
	}
	paths := indexByPath(atlasCommands(t))

	tests := []struct {
		path   string
		want   Category
		legacy LegacyCategory
	}{
		// One operation.
		{"atlas organizations create", CategoryWorkflow, LegacyL2},
		// Two: GetAccessHistoryCluster and GetAccessHistoryProcess, chosen by
		// --clusterName at runtime.
		{"atlas accessLogs list", CategoryDirect, LegacyL3},
		// No Atlas SDK operation at all.
		{"atlas config list", CategoryMeta, LegacyL3},
		// A command that only groups subcommands runs nothing.
		{"atlas clusters", "", LegacyL3},
		// The generated subtree, which reaches no Atlas SDK operation because
		// it goes through the hand-written executor.
		{"atlas api", "", LegacyL1},
		{"atlas api accessTracking getAccessHistoryCluster", CategoryMeta, LegacyL1},
	}

	for _, tt := range tests {
		command, ok := paths[tt.path]
		if !ok {
			t.Errorf("%s missing from the tree", tt.path)
			continue
		}
		if command.Category != tt.want {
			t.Errorf("%s category = %q, want %q (%d operations)",
				tt.path, command.Category, tt.want, len(command.ApiOperationIds))
		}
		if command.LegacyCategory != tt.legacy {
			t.Errorf("%s legacy_category = %q, want %q", tt.path, command.LegacyCategory, tt.legacy)
		}
	}
}

func TestBuildFlags(t *testing.T) {
	if testing.Short() {
		t.Skip("builds the CLI and a call graph over it")
	}
	paths := indexByPath(atlasCommands(t))

	list, ok := paths["atlas accessLogs list"]
	if !ok {
		t.Fatal("atlas accessLogs list missing from the tree")
	}
	for _, want := range []Flag{
		{Name: "clusterName", Type: "string"},
		{Name: "hostname", Type: "string"},
		{Name: "projectId", Type: "string"},
	} {
		if !hasFlag(list.Flags, want) {
			t.Errorf("%s flags = %+v, want %+v", list.Path, list.Flags, want)
		}
	}

	// The root's persistent flags belong to the root, not to its children.
	for _, flag := range list.Flags {
		if flag.Name == "profile" {
			t.Errorf("%s flags = %+v, want no inherited --profile", list.Path, list.Flags)
		}
	}
}

func hasFlag(values []Flag, want Flag) bool {
	for _, value := range values {
		if value == want {
			return true
		}
	}
	return false
}
