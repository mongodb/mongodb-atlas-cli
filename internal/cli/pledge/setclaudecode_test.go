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

//go:build !windows

package pledge_test

import (
	"os"
	"strings"
	"testing"

	clipledge "github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli/pledge"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/pledge"
)

func TestSetClaudeCode_ReadsSessionIDFromStdin(t *testing.T) {
	dir := withTempStateDir(t)
	clearSessionEnv(t)

	hookJSON := `{"session_id":"` + validUUID + `","hook_event_name":"SessionStart","source":"startup"}`

	cmd := clipledge.SetClaudeCodeBuilder()
	cmd.SetIn(strings.NewReader(hookJSON))
	cmd.SetArgs([]string{"readonly", "--yes"})
	if err := cmd.Execute(); err != nil {
		t.Fatalf("Execute failed: %v", err)
	}

	k, _ := pledge.NewSessionKey("claude", validUUID)
	pf, err := pledge.Load(k)
	if err != nil {
		t.Fatalf("Load failed: %v", err)
	}
	if pf.Profile != pledge.ProfileReadonly {
		t.Errorf("expected readonly, got %q", pf.Profile)
	}
	expectedFile := dir + "/claude-" + validUUID + ".json"
	if _, err := os.Stat(expectedFile); err != nil {
		t.Errorf("expected pledge file at %s: %v", expectedFile, err)
	}
}

func TestSetClaudeCode_FallsBackToResolverOnEmptyStdin(t *testing.T) {
	withTempStateDir(t)
	clearSessionEnv(t)
	t.Setenv("ATLAS_PLEDGE_SESSION_KEY", "claude:"+validUUID)

	cmd := clipledge.SetClaudeCodeBuilder()
	cmd.SetIn(strings.NewReader(""))
	cmd.SetArgs([]string{"readonly", "--yes"})
	if err := cmd.Execute(); err != nil {
		t.Fatalf("Execute failed: %v", err)
	}

	k, _ := pledge.NewSessionKey("claude", validUUID)
	pf, err := pledge.Load(k)
	if err != nil {
		t.Fatalf("Load failed: %v", err)
	}
	if pf.Profile != pledge.ProfileReadonly {
		t.Errorf("expected readonly, got %q", pf.Profile)
	}
}

func TestSetClaudeCode_InvalidProfile(t *testing.T) {
	withTempStateDir(t)
	clearSessionEnv(t)

	cmd := clipledge.SetClaudeCodeBuilder()
	cmd.SetIn(strings.NewReader("{}"))
	cmd.SetArgs([]string{"badprofile", "--yes"})
	if err := cmd.Execute(); err == nil {
		t.Error("expected error for unknown profile")
	}
}

func TestSetClaudeCode_AdminRequiresYes(t *testing.T) {
	withTempStateDir(t)
	clearSessionEnv(t)
	t.Setenv("ATLAS_PLEDGE_SESSION_KEY", "claude:"+validUUID)

	hookJSON := `{"session_id":"` + validUUID + `"}`
	cmd := clipledge.SetClaudeCodeBuilder()
	cmd.SetIn(strings.NewReader(hookJSON))
	cmd.SetArgs([]string{"admin"}) // no --yes
	if err := cmd.Execute(); err == nil {
		t.Error("expected error for admin without --yes")
	}
}

func TestSetClaudeCode_IgnoresMalformedStdinJSON(t *testing.T) {
	withTempStateDir(t)
	clearSessionEnv(t)
	t.Setenv("ATLAS_PLEDGE_SESSION_KEY", "claude:"+validUUID)

	cmd := clipledge.SetClaudeCodeBuilder()
	cmd.SetIn(strings.NewReader("not json at all"))
	cmd.SetArgs([]string{"readonly", "--yes"})
	if err := cmd.Execute(); err != nil {
		t.Fatalf("Execute failed on malformed stdin: %v", err)
	}

	k, _ := pledge.NewSessionKey("claude", validUUID)
	if _, err := pledge.Load(k); err != nil {
		t.Errorf("expected pledge saved via resolver fallback: %v", err)
	}
}
