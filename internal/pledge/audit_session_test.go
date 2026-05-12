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
	"bufio"
	"encoding/json"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/pledge"
)

func readAuditLog(t *testing.T, dir string) []map[string]any {
	t.Helper()
	data, err := os.ReadFile(filepath.Join(dir, "audit.log"))
	if err != nil {
		return nil
	}
	var entries []map[string]any
	sc := bufio.NewScanner(strings.NewReader(string(data)))
	for sc.Scan() {
		line := sc.Text()
		if line == "" {
			continue
		}
		var m map[string]any
		if err := json.Unmarshal([]byte(line), &m); err == nil {
			entries = append(entries, m)
		}
	}
	return entries
}

func TestAuditLogsSessionKeyOnSave(t *testing.T) {
	withTempStateDir(t)
	dir := os.Getenv("ATLAS_PLEDGE_STATE_DIR")

	k, _ := pledge.NewSessionKey("claude", testUUID)
	pf, _ := pledge.NewPledgeFile(pledge.ProfileReadonly, nil)
	if err := pledge.Save(k, pf); err != nil {
		t.Fatal(err)
	}

	entries := readAuditLog(t, dir)
	if len(entries) == 0 {
		t.Fatal("expected audit log entries after Save")
	}
	found := false
	for _, e := range entries {
		if sk, _ := e["sessionKey"].(string); strings.HasPrefix(sk, "claude-") {
			found = true
		}
	}
	if !found {
		t.Errorf("expected audit entry with claude-keyed sessionKey, got: %v", entries)
	}
}

func TestAuditLogsClaudeKeyMissing(t *testing.T) {
	withTempStateDir(t)
	dir := os.Getenv("ATLAS_PLEDGE_STATE_DIR")

	// Set CLAUDECODE=1 but no UUID available — resolver should log claude-key-missing.
	t.Setenv("CLAUDECODE", "1")
	t.Setenv("ATLAS_PLEDGE_SESSION_ID", "")
	t.Setenv("ATLAS_PLEDGE_SESSION_KEY", "")
	t.Setenv("CLAUDE_CODE_SESSION_ID", "")
	t.Setenv("CLAUDE_PROJECT_DIR", "")
	t.Setenv("XDG_RUNTIME_DIR", "")

	_, err := pledge.ResolveSessionKey()
	if err != nil {
		t.Fatal(err)
	}

	entries := readAuditLog(t, dir)
	found := false
	for _, e := range entries {
		if opID, _ := e["operationID"].(string); opID == "<claude-key-missing>" {
			found = true
		}
	}
	if !found {
		t.Errorf("expected claude-key-missing audit entry when CLAUDECODE=1 but no UUID; got: %v", entries)
	}
}
