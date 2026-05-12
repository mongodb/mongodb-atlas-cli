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
	"path/filepath"
	"testing"

	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/pledge"
)

// TestClaudeFlowClaudeCodeSessionIDToResolver simulates the simplified hook flow:
//  1. The SessionStart hook calls `atlas pledge set readonly --yes`.
//     CLAUDE_CODE_SESSION_ID is in the environment, so ResolveSessionKey picks it
//     up directly — no --session-id flag or stdin parsing required.
//  2. A subsequent Bash tool invocation also has CLAUDE_CODE_SESSION_ID set and
//     resolves to the same key, loading the pledge correctly.
func TestClaudeFlowClaudeCodeSessionIDToResolver(t *testing.T) {
	withTempStateDir(t)
	t.Setenv("ATLAS_PLEDGE_SESSION_KEY", "")
	t.Setenv("ATLAS_PLEDGE_SESSION_ID", "")
	t.Setenv("CLAUDE_PROJECT_DIR", "")
	t.Setenv("XDG_RUNTIME_DIR", "")

	const sessionUUID = "47a415e1-c599-401f-9bac-65e71a7232b2"

	// --- Step 1: SessionStart hook fires with CLAUDE_CODE_SESSION_ID in env ---
	t.Setenv("CLAUDECODE", "1")
	t.Setenv("CLAUDE_CODE_SESSION_ID", sessionUUID)

	k, err := pledge.ResolveSessionKey()
	if err != nil {
		t.Fatal(err)
	}
	if k.Kind != "claude" || k.Value != sessionUUID {
		t.Fatalf("expected claude:%s, got %s", sessionUUID, k.String())
	}
	pf, _ := pledge.NewPledgeFile(pledge.ProfileReadonly, nil)
	if err := pledge.Save(k, pf); err != nil {
		t.Fatal(err)
	}

	// --- Step 2: Subsequent Bash tool call — same env, same resolution ---
	resolvedKey, err := pledge.ResolveSessionKey()
	if err != nil {
		t.Fatal(err)
	}
	if resolvedKey.Kind != "claude" || resolvedKey.Value != sessionUUID {
		t.Errorf("expected resolver to find claude:%s via CLAUDE_CODE_SESSION_ID, got %s", sessionUUID, resolvedKey.String())
	}

	loadedPf, err := pledge.Load(resolvedKey)
	if err != nil {
		t.Fatalf("Load failed: %v", err)
	}
	if loadedPf.Profile != pledge.ProfileReadonly {
		t.Errorf("expected readonly profile, got %q", loadedPf.Profile)
	}
}

// TestClaudeFlowBreadcrumbToResolver simulates the Claude pledge flow:
//  1. The SessionStart hook calls `atlas pledge set --session-id <uuid>`.
//     In set.go, this writes the pledge file AND the breadcrumb.
//  2. A subsequent Bash invocation (new POSIX SID) runs with CLAUDECODE=1.
//     The resolver should find the UUID via the breadcrumb and load the pledge.
func TestClaudeFlowBreadcrumbToResolver(t *testing.T) {
	withTempStateDir(t)
	dir := os.Getenv("ATLAS_PLEDGE_STATE_DIR")
	xdgDir := t.TempDir()
	t.Setenv("XDG_RUNTIME_DIR", xdgDir)
	t.Setenv("ATLAS_PLEDGE_SESSION_KEY", "")
	t.Setenv("ATLAS_PLEDGE_SESSION_ID", "")
	t.Setenv("CLAUDE_CODE_SESSION_ID", "") // must be clear so breadcrumb path is exercised

	const projectDir = "/home/user/myproject"
	const sessionUUID = "47a415e1-c599-401f-9bac-65e71a7232b2"

	// --- Step 1: SessionStart hook fires ---
	// Simulate `atlas pledge set readonly --yes --session-id <uuid>`:
	// write the pledge file under claude-<uuid>.json.
	k, _ := pledge.NewSessionKey("claude", sessionUUID)
	pf, _ := pledge.NewPledgeFile(pledge.ProfileReadonly, nil)
	if err := pledge.Save(k, pf); err != nil {
		t.Fatal(err)
	}
	// Verify pledge file created.
	pledgeFile := filepath.Join(dir, "claude-"+sessionUUID+".json")
	if _, err := os.Stat(pledgeFile); err != nil {
		t.Fatalf("pledge file not created: %v", err)
	}

	// Simulate the breadcrumb written by `atlas pledge set --session-id`.
	// In real flow, set.go writes this; here we write it directly.
	breadcrumbDir := pledge.ClaudeBreadcrumbDir(projectDir)
	if err := os.MkdirAll(breadcrumbDir, 0o700); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(breadcrumbDir, sessionUUID), []byte(sessionUUID), 0o600); err != nil {
		t.Fatal(err)
	}

	// --- Step 2: Subsequent Bash invocation (new POSIX SID) ---
	// The new process has CLAUDECODE=1 and CLAUDE_PROJECT_DIR set,
	// but NO ATLAS_PLEDGE_SESSION_ID (it wasn't set in the env).
	t.Setenv("CLAUDECODE", "1")
	t.Setenv("CLAUDE_PROJECT_DIR", projectDir)

	resolvedKey, err := pledge.ResolveSessionKey()
	if err != nil {
		t.Fatal(err)
	}
	if resolvedKey.Kind != "claude" || resolvedKey.Value != sessionUUID {
		t.Errorf("expected resolver to find claude:%s via breadcrumb, got %s", sessionUUID, resolvedKey.String())
	}

	// Load the pledge using the resolved key — must find the pledge set in step 1.
	loadedPf, err := pledge.Load(resolvedKey)
	if err != nil {
		t.Fatalf("Load failed with resolver key: %v", err)
	}
	if loadedPf.Profile != pledge.ProfileReadonly {
		t.Errorf("expected readonly profile, got %q", loadedPf.Profile)
	}
}
