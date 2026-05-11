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

package root_test

import (
	"strings"
	"testing"

	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli/root"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/tools/shared/api"
	"github.com/spf13/cobra"
)

// metaCommandPaths are commands that are exempt from the permission annotation requirement.
// These are infrastructure commands that never dispatch Atlas API calls and don't mutate
// Atlas state, so there is nothing to restrict.
var metaCommandPaths = map[string]struct{}{
	"atlas completion":          {},
	"atlas completion bash":     {},
	"atlas completion fish":     {},
	"atlas completion powershell": {},
	"atlas completion zsh":      {},
	"atlas help":                {},
	"atlas pledge":              {},
	"atlas pledge allow":        {},
	"atlas pledge set":          {},
	"atlas pledge show":         {},
	"atlas hook":                {},
	"atlas hook install":        {},
	"atlas hook uninstall":      {},
}

var validTiers = map[string]struct{}{
	string(api.PermissionRead):       {},
	string(api.PermissionWrite):      {},
	string(api.PermissionAdmin):      {},
	string(api.PermissionLocalWrite): {},
}

// TestAllLeafCommandsHavePermissionAnnotation walks the full command tree and
// fails for any leaf command that lacks a valid atlas.permission annotation.
// This is the build-time guarantee that no hand-written command silently bypasses pledge.
func TestAllLeafCommandsHavePermissionAnnotation(t *testing.T) {
	rootCmd := root.Builder()

	var missing []string
	var invalid []string

	var walk func(cmd *cobra.Command)
	walk = func(cmd *cobra.Command) {
		for _, sub := range cmd.Commands() {
			walk(sub)
		}
		if cmd.HasSubCommands() {
			return
		}

		path := cmd.CommandPath()
		if _, exempt := metaCommandPaths[path]; exempt {
			return
		}

		// Generated API commands under "atlas api" carry their tier inline;
		// skip the annotation check for them (enforced in executor.go).
		if strings.HasPrefix(path, "atlas api ") {
			return
		}

		tier, ok := cmd.Annotations[cli.AnnotationKeyPermission]
		if !ok {
			missing = append(missing, path)
			return
		}
		if _, valid := validTiers[tier]; !valid {
			invalid = append(invalid, path+" (has: "+tier+")")
		}
	}
	walk(rootCmd)

	for _, p := range missing {
		t.Errorf("leaf command %q has no %q annotation; add cli.SetPermission(cmd, ...) in its Builder()", p, cli.AnnotationKeyPermission)
	}
	for _, p := range invalid {
		t.Errorf("leaf command %q has an invalid permission tier: %s", p, p)
	}
}
