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

package agent

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"
)

func newTestCodex(t *testing.T) *codexAgent {
	t.Helper()
	dir := t.TempDir()
	return &codexAgent{hooksPath: filepath.Join(dir, ".codex", "hooks.json")}
}

func TestCodexInstall_MissingFile(t *testing.T) {
	a := newTestCodex(t)

	if err := a.Install(InstallOpts{Profile: "readonly"}); err != nil {
		t.Fatalf("Install: %v", err)
	}

	raw, err := os.ReadFile(a.hooksPath)
	if err != nil {
		t.Fatalf("reading hooks.json: %v", err)
	}
	var hooks map[string][]map[string]any
	if err := json.Unmarshal(raw, &hooks); err != nil {
		t.Fatalf("parsing hooks.json: %v", err)
	}
	entries := hooks["PreToolUse"]
	if len(entries) != 1 {
		t.Fatalf("expected 1 PreToolUse entry, got %d", len(entries))
	}
	if managed, _ := entries[0][managedTag].(bool); !managed {
		t.Error("entry missing _atlas_managed tag")
	}
	if cmd, _ := entries[0]["command"].(string); cmd != "atlas pledge set readonly --yes" {
		t.Errorf("unexpected command: %q", cmd)
	}
}

func TestCodexInstall_PreservesExistingEntries(t *testing.T) {
	a := newTestCodex(t)

	// Pre-populate with a user entry.
	existing := map[string][]map[string]any{
		"PreToolUse": {{"matcher": "^Bash$", "command": "other-tool"}},
	}
	if err := os.MkdirAll(filepath.Dir(a.hooksPath), 0o700); err != nil {
		t.Fatal(err)
	}
	data, _ := json.Marshal(existing)
	if err := os.WriteFile(a.hooksPath, data, 0o600); err != nil {
		t.Fatal(err)
	}

	if err := a.Install(InstallOpts{Profile: "readonly"}); err != nil {
		t.Fatalf("Install: %v", err)
	}

	raw, _ := os.ReadFile(a.hooksPath)
	var hooks map[string][]map[string]any
	_ = json.Unmarshal(raw, &hooks)

	entries := hooks["PreToolUse"]
	if len(entries) != 2 {
		t.Fatalf("expected 2 entries, got %d", len(entries))
	}
	// First entry should be the original user entry (order preserved by append).
	if cmd, _ := entries[0]["command"].(string); cmd != "other-tool" {
		t.Errorf("user entry not preserved, got: %q", cmd)
	}
}

func TestCodexInstall_Idempotent(t *testing.T) {
	a := newTestCodex(t)

	if err := a.Install(InstallOpts{Profile: "readonly"}); err != nil {
		t.Fatalf("first Install: %v", err)
	}
	if err := a.Install(InstallOpts{Profile: "readonly"}); err != nil {
		t.Fatalf("second Install: %v", err)
	}

	raw, _ := os.ReadFile(a.hooksPath)
	var hooks map[string][]map[string]any
	_ = json.Unmarshal(raw, &hooks)

	if len(hooks["PreToolUse"]) != 1 {
		t.Errorf("expected 1 entry after idempotent install, got %d", len(hooks["PreToolUse"]))
	}
}

func TestCodexUninstall_RemovesManagedEntry(t *testing.T) {
	a := newTestCodex(t)

	existing := map[string][]map[string]any{
		"PreToolUse": {
			{"matcher": "^Bash$", "command": "other-tool"},
			{managedTag: true, "command": "atlas pledge set readonly --yes"},
		},
	}
	if err := os.MkdirAll(filepath.Dir(a.hooksPath), 0o700); err != nil {
		t.Fatal(err)
	}
	data, _ := json.Marshal(existing)
	_ = os.WriteFile(a.hooksPath, data, 0o600)

	if err := a.Uninstall(); err != nil {
		t.Fatalf("Uninstall: %v", err)
	}

	raw, _ := os.ReadFile(a.hooksPath)
	var hooks map[string][]map[string]any
	_ = json.Unmarshal(raw, &hooks)

	entries := hooks["PreToolUse"]
	if len(entries) != 1 {
		t.Fatalf("expected 1 remaining entry, got %d", len(entries))
	}
	if cmd, _ := entries[0]["command"].(string); cmd != "other-tool" {
		t.Errorf("unexpected remaining entry: %q", cmd)
	}
}

func TestCodexUninstall_NoopWhenAbsent(t *testing.T) {
	a := newTestCodex(t)
	if err := a.Uninstall(); err != nil {
		t.Fatalf("Uninstall on missing file: %v", err)
	}
}

func TestCodexStatus(t *testing.T) {
	a := newTestCodex(t)

	if got := a.Status(); got != StateUninstalled {
		t.Errorf("expected StateUninstalled before install, got %q", got)
	}

	_ = a.Install(InstallOpts{Profile: "readonly"})

	if got := a.Status(); got != StateInstalled {
		t.Errorf("expected StateInstalled after install, got %q", got)
	}
}
