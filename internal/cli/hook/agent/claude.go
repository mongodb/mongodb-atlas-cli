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
	"errors"
	"fmt"
	"os"
	"path/filepath"
)

// managedTag is written into every hook entry created by atlas hook install.
// atlas hook uninstall removes only entries with this tag.
const managedTag = "_atlas_managed"

// claudeSettings is a minimal representation of ~/.claude/settings.json.
// Unknown fields are preserved via Extra.
type claudeSettings struct {
	Hooks map[string][]claudeHook `json:"hooks,omitempty"`
	Extra map[string]any          `json:"-"`
}

type claudeHook struct {
	Matcher      string            `json:"matcher"`
	Hooks        []claudeHookEntry `json:"hooks"`
	AtlasManaged bool              `json:"_atlas_managed,omitempty"`
}

type claudeHookEntry struct {
	Type    string `json:"type"`
	Command string `json:"command"`
}

// claudeCodeAgent implements Agent for Claude Code (~/.claude/settings.json).
type claudeCodeAgent struct {
	settingsPath string
}

// NewClaudeCode returns an Agent for Claude Code.
// settingsPath defaults to ~/.claude/settings.json when empty.
func NewClaudeCode(settingsPath string) Agent {
	if settingsPath == "" {
		home, _ := os.UserHomeDir()
		settingsPath = filepath.Join(home, ".claude", "settings.json")
	}
	return &claudeCodeAgent{settingsPath: settingsPath}
}

func (a *claudeCodeAgent) Name() string { return "claude-code" }

func (a *claudeCodeAgent) Status() State {
	raw, err := os.ReadFile(a.settingsPath)
	if err != nil {
		return StateUninstalled
	}
	var top map[string]json.RawMessage
	if err := json.Unmarshal(raw, &top); err != nil {
		return StateUnknown
	}
	hooksRaw, ok := top["hooks"]
	if !ok {
		return StateUninstalled
	}
	var hooks map[string][]map[string]any
	if err := json.Unmarshal(hooksRaw, &hooks); err != nil {
		return StateUnknown
	}
	for _, entries := range hooks {
		for _, e := range entries {
			if managed, _ := e[managedTag].(bool); managed {
				return StateInstalled
			}
		}
	}
	return StateUninstalled
}

func (a *claudeCodeAgent) Install(opts InstallOpts) error {
	if err := os.MkdirAll(filepath.Dir(a.settingsPath), 0o700); err != nil {
		return fmt.Errorf("creating .claude directory: %w", err)
	}

	// Read existing settings (or start empty).
	var top map[string]json.RawMessage
	if raw, err := os.ReadFile(a.settingsPath); err == nil {
		if jsonErr := json.Unmarshal(raw, &top); jsonErr != nil {
			return fmt.Errorf("parsing %s: %w", a.settingsPath, jsonErr)
		}
	}
	if top == nil {
		top = make(map[string]json.RawMessage)
	}

	// Parse existing hooks.
	var hooks map[string][]map[string]any
	if hooksRaw, ok := top["hooks"]; ok {
		if err := json.Unmarshal(hooksRaw, &hooks); err != nil {
			return fmt.Errorf("parsing hooks in %s: %w", a.settingsPath, err)
		}
	}
	if hooks == nil {
		hooks = make(map[string][]map[string]any)
	}

	// Remove any existing managed entries (idempotent re-install).
	hooks["SessionStart"] = removeManaged(hooks["SessionStart"])

	profile := opts.Profile
	if profile == "" {
		profile = "readonly"
	}

	entry := map[string]any{
		managedTag: true,
		"matcher":  "",
		"hooks": []map[string]any{
			{
				"type":    "command",
				"command": fmt.Sprintf("atlas pledge set %s --yes", profile),
			},
		},
	}
	hooks["SessionStart"] = append(hooks["SessionStart"], entry)

	hooksBytes, err := json.Marshal(hooks)
	if err != nil {
		return err
	}
	top["hooks"] = hooksBytes

	return writeSettings(a.settingsPath, top)
}

func (a *claudeCodeAgent) Uninstall() error {
	raw, err := os.ReadFile(a.settingsPath)
	if errors.Is(err, os.ErrNotExist) {
		return nil
	}
	if err != nil {
		return err
	}

	var top map[string]json.RawMessage
	if err := json.Unmarshal(raw, &top); err != nil {
		return fmt.Errorf("parsing %s: %w", a.settingsPath, err)
	}

	hooksRaw, ok := top["hooks"]
	if !ok {
		return nil
	}
	var hooks map[string][]map[string]any
	if err := json.Unmarshal(hooksRaw, &hooks); err != nil {
		return fmt.Errorf("parsing hooks: %w", err)
	}

	for event, entries := range hooks {
		hooks[event] = removeManaged(entries)
	}

	hooksBytes, err := json.Marshal(hooks)
	if err != nil {
		return err
	}
	top["hooks"] = hooksBytes

	return writeSettings(a.settingsPath, top)
}

// removeManaged returns entries with _atlas_managed entries removed.
func removeManaged(entries []map[string]any) []map[string]any {
	var out []map[string]any
	for _, e := range entries {
		if managed, _ := e[managedTag].(bool); managed {
			continue
		}
		out = append(out, e)
	}
	return out
}

// writeSettings writes settings atomically with a .bak backup.
func writeSettings(path string, top map[string]json.RawMessage) error {
	data, err := json.MarshalIndent(top, "", "  ")
	if err != nil {
		return err
	}

	// Backup existing file.
	if _, statErr := os.Stat(path); statErr == nil {
		if existing, readErr := os.ReadFile(path); readErr == nil {
			_ = os.WriteFile(path+".bak", existing, 0o600)
		}
	}

	tmp := path + ".tmp"
	if err := os.WriteFile(tmp, data, 0o600); err != nil {
		return fmt.Errorf("writing settings: %w", err)
	}
	return os.Rename(tmp, path)
}
