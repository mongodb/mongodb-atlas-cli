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
	"strings"
	"testing"

	"github.com/spf13/cobra"
)

func TestWalk(t *testing.T) {
	root := &cobra.Command{Use: "atlas"}
	group := &cobra.Command{Use: "organizations", Short: "Manage organizations"}
	group.AddCommand(
		&cobra.Command{Use: "list", RunE: func(*cobra.Command, []string) error { return nil }},
		&cobra.Command{Use: "purge", Hidden: true, Run: func(*cobra.Command, []string) {}},
	)
	root.AddCommand(group, &cobra.Command{
		Use:   "version",
		Short: "Print the version",
		RunE:  func(*cobra.Command, []string) error { return nil },
	})

	got := Walk(root)

	if got.Path != "atlas" {
		t.Fatalf("root path = %q, want %q", got.Path, "atlas")
	}
	if got.Run != "" {
		t.Errorf("root run = %q, want no run function", got.Run)
	}
	if len(got.Children) != 2 {
		t.Fatalf("root has %d children, want 2", len(got.Children))
	}

	// Subcommands come out sorted by name.
	orgs, version := got.Children[0], got.Children[1]
	if orgs.Path != "atlas organizations" {
		t.Errorf("first child path = %q, want %q", orgs.Path, "atlas organizations")
	}
	if orgs.Run != "" {
		t.Errorf("%s run = %q, want no run function", orgs.Path, orgs.Run)
	}
	if orgs.Short != "Manage organizations" {
		t.Errorf("%s short = %q", orgs.Path, orgs.Short)
	}
	if len(orgs.Children) != 2 {
		t.Fatalf("%s has %d children, want 2", orgs.Path, len(orgs.Children))
	}

	list, purge := orgs.Children[0], orgs.Children[1]
	if list.Path != "atlas organizations list" {
		t.Errorf("leaf path = %q, want %q", list.Path, "atlas organizations list")
	}
	if list.Run == "" {
		t.Errorf("%s has no run function", list.Path)
	}
	if !strings.Contains(list.Run, "TestWalk") {
		t.Errorf("%s run = %q, want a symbol from this test", list.Path, list.Run)
	}
	if len(list.Children) != 0 {
		t.Errorf("%s has %d children, want none", list.Path, len(list.Children))
	}

	if purge.Path != "atlas organizations purge" {
		t.Errorf("leaf path = %q, want %q", purge.Path, "atlas organizations purge")
	}
	if !purge.Hidden {
		t.Errorf("%s should be hidden", purge.Path)
	}
	if purge.Run == "" {
		t.Errorf("%s has no run function", purge.Path)
	}

	if version.Path != "atlas version" || version.Short != "Print the version" {
		t.Errorf("second child = %q %q", version.Path, version.Short)
	}
	if version.Run == "" {
		t.Errorf("%s has no run function", version.Path)
	}
}
