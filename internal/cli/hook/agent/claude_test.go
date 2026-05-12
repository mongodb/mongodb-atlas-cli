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

package agent_test

import (
	"encoding/json"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli/hook/agent"
)

func settingsPath(t *testing.T) string {
	t.Helper()
	return filepath.Join(t.TempDir(), "settings.json")
}

func readJSON(t *testing.T, path string) map[string]any {
	t.Helper()
	raw, err := os.ReadFile(path)
	if err != nil {
		t.Fatalf("reading %s: %v", path, err)
	}
	var m map[string]any
	if err := json.Unmarshal(raw, &m); err != nil {
		t.Fatalf("parsing %s: %v", path, err)
	}
	return m
}

func hasAtlasManagedHook(m map[string]any, event string) bool {
	hooks, ok := m["hooks"].(map[string]any)
	if !ok {
		return false
	}
	entries, ok := hooks[event].([]any)
	if !ok {
		return false
	}
	for _, e := range entries {
		entry, ok := e.(map[string]any)
		if !ok {
			continue
		}
		if managed, _ := entry["_atlas_managed"].(bool); managed {
			return true
		}
	}
	return false
}

func TestClaudeInstallOnEmptyFile(t *testing.T) {
	path := settingsPath(t)
	a := agent.NewClaudeCode(path)

	if err := a.Install(agent.InstallOpts{Profile: "readonly"}); err != nil {
		t.Fatal(err)
	}
	m := readJSON(t, path)
	if !hasAtlasManagedHook(m, "SessionStart") {
		t.Error("expected _atlas_managed SessionStart hook after install on empty file")
	}
}

func TestClaudeInstallOnPopulatedFile(t *testing.T) {
	path := settingsPath(t)
	// Pre-populate with an unrelated hook.
	initial := `{"hooks":{"SessionStart":[{"type":"command","command":"echo hello"}]}}`
	if err := os.WriteFile(path, []byte(initial), 0o600); err != nil {
		t.Fatal(err)
	}

	a := agent.NewClaudeCode(path)
	if err := a.Install(agent.InstallOpts{Profile: "read-write"}); err != nil {
		t.Fatal(err)
	}

	m := readJSON(t, path)
	if !hasAtlasManagedHook(m, "SessionStart") {
		t.Error("expected managed hook after install on populated file")
	}

	// Original hook must still be present.
	hooks := m["hooks"].(map[string]any)
	entries := hooks["SessionStart"].([]any)
	if len(entries) < 2 {
		t.Errorf("expected at least 2 SessionStart hooks, got %d", len(entries))
	}
}

func TestClaudeInstallIdempotent(t *testing.T) {
	path := settingsPath(t)
	a := agent.NewClaudeCode(path)

	opts := agent.InstallOpts{Profile: "readonly"}
	if err := a.Install(opts); err != nil {
		t.Fatal(err)
	}
	if err := a.Install(opts); err != nil {
		t.Fatal(err)
	}

	// Should still have exactly one managed hook.
	m := readJSON(t, path)
	hooks := m["hooks"].(map[string]any)
	entries := hooks["SessionStart"].([]any)
	managed := 0
	for _, e := range entries {
		if entry, ok := e.(map[string]any); ok {
			if v, _ := entry["_atlas_managed"].(bool); v {
				managed++
			}
		}
	}
	if managed != 1 {
		t.Errorf("expected exactly 1 managed hook after idempotent install, got %d", managed)
	}
}

func TestClaudeInstallHookShape(t *testing.T) {
	path := settingsPath(t)
	a := agent.NewClaudeCode(path)

	if err := a.Install(agent.InstallOpts{Profile: "readonly"}); err != nil {
		t.Fatal(err)
	}

	m := readJSON(t, path)
	hooks := m["hooks"].(map[string]any)
	entries := hooks["SessionStart"].([]any)

	var managed map[string]any
	for _, e := range entries {
		entry, ok := e.(map[string]any)
		if !ok {
			continue
		}
		if v, _ := entry["_atlas_managed"].(bool); v {
			managed = entry
			break
		}
	}
	if managed == nil {
		t.Fatal("no managed hook entry found")
	}

	// Must have the nested hooks array, not flat type/command fields.
	innerHooks, ok := managed["hooks"].([]any)
	if !ok || len(innerHooks) == 0 {
		t.Errorf("expected managed entry to have nested hooks array, got: %v", managed)
	}
	if _, hasType := managed["type"]; hasType {
		t.Error("managed entry must not have top-level 'type' field (old flat format)")
	}
	if _, hasMatcher := managed["matcher"]; !hasMatcher {
		t.Error("managed entry must have a 'matcher' field")
	}

	// The inner hook command must use the set-claude-code subcommand.
	inner := innerHooks[0].(map[string]any)
	cmd, _ := inner["command"].(string)
	if !strings.Contains(cmd, "pledge set-claude-code") {
		t.Errorf("installed command must use 'pledge set-claude-code', got: %q", cmd)
	}
	if !strings.Contains(cmd, "--yes") {
		t.Errorf("installed command must include --yes, got: %q", cmd)
	}
}

func TestClaudeUninstallPreservesUnrelated(t *testing.T) {
	path := settingsPath(t)
	initial := `{"hooks":{"SessionStart":[{"type":"command","command":"echo hello"}]}}`
	if err := os.WriteFile(path, []byte(initial), 0o600); err != nil {
		t.Fatal(err)
	}

	a := agent.NewClaudeCode(path)
	if err := a.Install(agent.InstallOpts{Profile: "readonly"}); err != nil {
		t.Fatal(err)
	}
	if err := a.Uninstall(); err != nil {
		t.Fatal(err)
	}

	m := readJSON(t, path)
	if hasAtlasManagedHook(m, "SessionStart") {
		t.Error("managed hook should be gone after uninstall")
	}
	// Unrelated hook must survive.
	hooks := m["hooks"].(map[string]any)
	entries := hooks["SessionStart"].([]any)
	if len(entries) != 1 {
		t.Errorf("expected 1 unrelated hook to survive uninstall, got %d", len(entries))
	}
}
