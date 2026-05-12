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
	"time"

	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/pledge"
)

const (
	uuid1 = "47a415e1-c599-401f-9bac-65e71a7232b2"
	uuid2 = "550e8400-e29b-41d4-a716-446655440000"
)

// clearClaudeEnv removes Claude-related env vars for a clean resolver test.
func clearClaudeEnv(t *testing.T) {
	t.Helper()
	t.Setenv("CLAUDECODE", "")
	t.Setenv("ATLAS_PLEDGE_SESSION_ID", "")
	t.Setenv("ATLAS_PLEDGE_SESSION_KEY", "")
	t.Setenv("CLAUDE_CODE_SESSION_ID", "")
	t.Setenv("CLAUDE_PROJECT_DIR", "")
	t.Setenv("XDG_RUNTIME_DIR", "")
}

func TestResolveSessionKey_POSIXDefault(t *testing.T) {
	clearClaudeEnv(t)
	k, err := pledge.ResolveSessionKey()
	if err != nil {
		t.Fatal(err)
	}
	if k.Kind != "sid" {
		t.Errorf("expected kind=sid, got %q", k.Kind)
	}
	if k.Value == "" {
		t.Error("expected non-empty SID value")
	}
}

func TestResolveSessionKey_ExplicitOverride(t *testing.T) {
	clearClaudeEnv(t)
	t.Setenv("ATLAS_PLEDGE_SESSION_KEY", "claude:"+uuid1)
	k, err := pledge.ResolveSessionKey()
	if err != nil {
		t.Fatal(err)
	}
	if k.Kind != "claude" || k.Value != uuid1 {
		t.Errorf("expected claude:%s, got %s", uuid1, k.String())
	}
}

func TestResolveSessionKey_ExplicitOverrideBadFormat(t *testing.T) {
	clearClaudeEnv(t)
	t.Setenv("ATLAS_PLEDGE_SESSION_KEY", "notacolon")
	_, err := pledge.ResolveSessionKey()
	if err == nil {
		t.Error("expected error for malformed ATLAS_PLEDGE_SESSION_KEY")
	}
}

func TestResolveSessionKey_ClaudeEnvWithSessionIDEnvVar(t *testing.T) {
	clearClaudeEnv(t)
	t.Setenv("CLAUDECODE", "1")
	t.Setenv("ATLAS_PLEDGE_SESSION_ID", uuid1)
	k, err := pledge.ResolveSessionKey()
	if err != nil {
		t.Fatal(err)
	}
	if k.Kind != "claude" || k.Value != uuid1 {
		t.Errorf("expected claude:%s, got %s", uuid1, k.String())
	}
}

func TestResolveSessionKey_ClaudeCodeSessionIDEnvVar(t *testing.T) {
	clearClaudeEnv(t)
	t.Setenv("CLAUDECODE", "1")
	t.Setenv("CLAUDE_CODE_SESSION_ID", uuid1)
	k, err := pledge.ResolveSessionKey()
	if err != nil {
		t.Fatal(err)
	}
	if k.Kind != "claude" || k.Value != uuid1 {
		t.Errorf("expected claude:%s, got %s", uuid1, k.String())
	}
}

// TestResolveSessionKey_ClaudeCodeSessionIDWithoutCLAUDECODE covers the
// SessionStart hook case: Claude Code spawns the hook subprocess with
// CLAUDE_CODE_SESSION_ID set but does not propagate CLAUDECODE. The resolver
// must still pick up the UUID so the pledge file is keyed on the conversation
// rather than the POSIX SID.
func TestResolveSessionKey_ClaudeCodeSessionIDWithoutCLAUDECODE(t *testing.T) {
	clearClaudeEnv(t)
	t.Setenv("CLAUDE_CODE_SESSION_ID", uuid1)
	k, err := pledge.ResolveSessionKey()
	if err != nil {
		t.Fatal(err)
	}
	if k.Kind != "claude" || k.Value != uuid1 {
		t.Errorf("expected claude:%s, got %s", uuid1, k.String())
	}
}

// TestResolveSessionKey_AtlasPledgeSessionIDWithoutCLAUDECODE verifies the
// explicit UUID override works without CLAUDECODE — e.g. when a wrapper script
// exports ATLAS_PLEDGE_SESSION_ID before invoking atlas.
func TestResolveSessionKey_AtlasPledgeSessionIDWithoutCLAUDECODE(t *testing.T) {
	clearClaudeEnv(t)
	t.Setenv("ATLAS_PLEDGE_SESSION_ID", uuid1)
	k, err := pledge.ResolveSessionKey()
	if err != nil {
		t.Fatal(err)
	}
	if k.Kind != "claude" || k.Value != uuid1 {
		t.Errorf("expected claude:%s, got %s", uuid1, k.String())
	}
}

func TestResolveSessionKey_AtlasPledgeSessionIDTakesPrecedenceOverClaudeCode(t *testing.T) {
	clearClaudeEnv(t)
	t.Setenv("CLAUDECODE", "1")
	t.Setenv("ATLAS_PLEDGE_SESSION_ID", uuid1)
	t.Setenv("CLAUDE_CODE_SESSION_ID", uuid2)
	k, err := pledge.ResolveSessionKey()
	if err != nil {
		t.Fatal(err)
	}
	if k.Kind != "claude" || k.Value != uuid1 {
		t.Errorf("expected ATLAS_PLEDGE_SESSION_ID to win: claude:%s, got %s", uuid1, k.String())
	}
}

func TestResolveSessionKey_ClaudeEnvWithBreadcrumb(t *testing.T) {
	clearClaudeEnv(t)
	tmpDir := t.TempDir()
	t.Setenv("CLAUDECODE", "1")
	t.Setenv("CLAUDE_PROJECT_DIR", "/home/user/myproject")
	t.Setenv("XDG_RUNTIME_DIR", tmpDir)

	// Write a breadcrumb manually in the expected location.
	breadcrumbDir := pledge.ClaudeBreadcrumbDir("/home/user/myproject")
	if err := os.MkdirAll(breadcrumbDir, 0o700); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(breadcrumbDir, uuid1), []byte(uuid1), 0o600); err != nil {
		t.Fatal(err)
	}

	k, err := pledge.ResolveSessionKey()
	if err != nil {
		t.Fatal(err)
	}
	if k.Kind != "claude" || k.Value != uuid1 {
		t.Errorf("expected claude:%s, got %s", uuid1, k.String())
	}
}

func TestResolveSessionKey_ClaudeMultipleBreadcrumbsNewestWins(t *testing.T) {
	clearClaudeEnv(t)
	tmpDir := t.TempDir()
	t.Setenv("CLAUDECODE", "1")
	t.Setenv("CLAUDE_PROJECT_DIR", "/home/user/myproject")
	t.Setenv("XDG_RUNTIME_DIR", tmpDir)

	breadcrumbDir := pledge.ClaudeBreadcrumbDir("/home/user/myproject")
	if err := os.MkdirAll(breadcrumbDir, 0o700); err != nil {
		t.Fatal(err)
	}

	// Write two breadcrumbs; uuid2 is the newest.
	if err := os.WriteFile(filepath.Join(breadcrumbDir, uuid1), []byte(uuid1), 0o600); err != nil {
		t.Fatal(err)
	}
	time.Sleep(10 * time.Millisecond)
	if err := os.WriteFile(filepath.Join(breadcrumbDir, uuid2), []byte(uuid2), 0o600); err != nil {
		t.Fatal(err)
	}

	k, err := pledge.ResolveSessionKey()
	if err != nil {
		t.Fatal(err)
	}
	if k.Kind != "claude" || k.Value != uuid2 {
		t.Errorf("expected newest breadcrumb uuid2 (%s), got %s", uuid2, k.String())
	}
}

func TestResolveSessionKey_ClaudeNoUUIDFallsBackToSID(t *testing.T) {
	withTempStateDir(t)
	clearClaudeEnv(t)
	t.Setenv("CLAUDECODE", "1")
	// No ATLAS_PLEDGE_SESSION_ID and no breadcrumb dir → fallback to SID.
	k, err := pledge.ResolveSessionKey()
	if err != nil {
		t.Fatal(err)
	}
	if k.Kind != "sid" {
		t.Errorf("expected fallback to sid, got %q", k.Kind)
	}
}

func TestResolveSessionKey_MalformedEnvVarUUIDFallback(t *testing.T) {
	clearClaudeEnv(t)
	t.Setenv("CLAUDECODE", "1")
	t.Setenv("ATLAS_PLEDGE_SESSION_ID", "not-a-uuid")
	// Malformed UUID → should fall back to SID or breadcrumb, not error.
	k, err := pledge.ResolveSessionKey()
	if err != nil {
		t.Fatal(err)
	}
	// Should not have picked up the malformed UUID.
	if k.Value == "not-a-uuid" {
		t.Error("malformed UUID should not be used as session key value")
	}
}
