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

const (
	validUUID   = "47a415e1-c599-401f-9bac-65e71a7232b2"
	invalidUUID = "not-a-valid-uuid"
)

func withTempStateDir(t *testing.T) string {
	t.Helper()
	dir := t.TempDir()
	t.Setenv("ATLAS_PLEDGE_STATE_DIR", dir)
	return dir
}

func clearSessionEnv(t *testing.T) {
	t.Helper()
	t.Setenv("ATLAS_PLEDGE_SESSION_KEY", "")
	t.Setenv("ATLAS_PLEDGE_SESSION_ID", "")
	t.Setenv("CLAUDECODE", "")
}

func TestSetWithExplicitSessionID_WritesClaudeKeyedFile(t *testing.T) {
	dir := withTempStateDir(t)
	clearSessionEnv(t)

	cmd := clipledge.SetBuilder()
	cmd.SetArgs([]string{"readonly", "--yes", "--session-id", validUUID})
	if err := cmd.Execute(); err != nil {
		t.Fatalf("Execute failed: %v", err)
	}

	k, _ := pledge.NewSessionKey("claude", validUUID)
	pf, err := pledge.Load(k)
	if err != nil {
		t.Fatalf("Load failed: %v", err)
	}
	if pf.Profile != pledge.ProfileReadonly {
		t.Errorf("expected profile readonly, got %q", pf.Profile)
	}
	// Verify the file is at the expected path.
	expectedFile := dir + "/claude-" + validUUID + ".json"
	if _, err := os.Stat(expectedFile); err != nil {
		t.Errorf("expected pledge file at %s: %v", expectedFile, err)
	}
}

func TestSetWithInvalidSessionID_RejectsAndWritesNothing(t *testing.T) {
	dir := withTempStateDir(t)
	clearSessionEnv(t)

	cmd := clipledge.SetBuilder()
	cmd.SetArgs([]string{"readonly", "--yes", "--session-id", invalidUUID})
	err := cmd.Execute()
	if err == nil {
		t.Fatal("expected error for invalid UUID, got nil")
	}
	if !strings.Contains(err.Error(), "invalid") {
		t.Errorf("expected 'invalid' in error message, got %q", err.Error())
	}
	// No pledge file should be written.
	entries, _ := os.ReadDir(dir)
	for _, e := range entries {
		if strings.HasSuffix(e.Name(), ".json") {
			t.Errorf("unexpected pledge file written: %s", e.Name())
		}
	}
}

func TestSetWithSessionIDEnvVar_WritesClaudeKeyedFile(t *testing.T) {
	dir := withTempStateDir(t)
	clearSessionEnv(t)
	t.Setenv("ATLAS_PLEDGE_SESSION_ID", validUUID)

	cmd := clipledge.SetBuilder()
	cmd.SetArgs([]string{"readonly", "--yes"})
	if err := cmd.Execute(); err != nil {
		t.Fatalf("Execute failed: %v", err)
	}

	expectedFile := dir + "/claude-" + validUUID + ".json"
	if _, err := os.Stat(expectedFile); err != nil {
		t.Errorf("expected claude-keyed pledge file via env var: %v", err)
	}
}

func TestSetFlagWinsOverEnvVar(t *testing.T) {
	dir := withTempStateDir(t)
	clearSessionEnv(t)
	otherUUID := "550e8400-e29b-41d4-a716-446655440000"
	t.Setenv("ATLAS_PLEDGE_SESSION_ID", otherUUID)

	cmd := clipledge.SetBuilder()
	cmd.SetArgs([]string{"readonly", "--yes", "--session-id", validUUID})
	if err := cmd.Execute(); err != nil {
		t.Fatalf("Execute failed: %v", err)
	}

	// Flag UUID should win.
	expectedFile := dir + "/claude-" + validUUID + ".json"
	if _, err := os.Stat(expectedFile); err != nil {
		t.Errorf("expected flag UUID to win, file %s missing: %v", expectedFile, err)
	}
	wrongFile := dir + "/claude-" + otherUUID + ".json"
	if _, err := os.Stat(wrongFile); err == nil {
		t.Errorf("env-var UUID file should not exist: %s", wrongFile)
	}
}
