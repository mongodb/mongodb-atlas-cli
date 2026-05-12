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

type codexAgent struct {
	hooksPath string
}

// NewCodex returns an Agent for Codex (~/.codex/hooks.json).
func NewCodex() Agent {
	home, _ := os.UserHomeDir()
	return &codexAgent{hooksPath: filepath.Join(home, ".codex", "hooks.json")}
}

func (a *codexAgent) Name() string { return "codex" }

func (a *codexAgent) Status() State {
	raw, err := os.ReadFile(a.hooksPath)
	if err != nil {
		return StateUninstalled
	}
	var hooks map[string][]map[string]any
	if err := json.Unmarshal(raw, &hooks); err != nil {
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

func (a *codexAgent) Install(opts InstallOpts) error {
	if err := os.MkdirAll(filepath.Dir(a.hooksPath), 0o700); err != nil {
		return fmt.Errorf("creating .codex directory: %w", err)
	}

	var hooks map[string][]map[string]any
	if raw, err := os.ReadFile(a.hooksPath); err == nil {
		if jsonErr := json.Unmarshal(raw, &hooks); jsonErr != nil {
			return fmt.Errorf("parsing %s: %w", a.hooksPath, jsonErr)
		}
	}
	if hooks == nil {
		hooks = make(map[string][]map[string]any)
	}

	// Check for existing managed entry.
	for _, e := range hooks["PreToolUse"] {
		if managed, _ := e[managedTag].(bool); managed {
			fmt.Fprintf(os.Stderr, "hooks.PreToolUse: already installed, skipping\n")
			return nil
		}
	}

	profile := opts.Profile
	if profile == "" {
		profile = "readonly"
	}

	entry := map[string]any{
		managedTag: true,
		"matcher":  "",
		"command":  fmt.Sprintf("atlas pledge set %s --yes", profile),
	}
	hooks["PreToolUse"] = append(hooks["PreToolUse"], entry)

	return writeCodexHooks(a.hooksPath, hooks)
}

func (a *codexAgent) Uninstall() error {
	raw, err := os.ReadFile(a.hooksPath)
	if errors.Is(err, os.ErrNotExist) {
		return nil
	}
	if err != nil {
		return err
	}

	var hooks map[string][]map[string]any
	if err := json.Unmarshal(raw, &hooks); err != nil {
		return fmt.Errorf("parsing %s: %w", a.hooksPath, err)
	}

	anyRemoved := false
	for event, entries := range hooks {
		filtered := removeManaged(entries)
		if len(filtered) != len(entries) {
			anyRemoved = true
		}
		hooks[event] = filtered
	}

	if !anyRemoved {
		return nil
	}

	return writeCodexHooks(a.hooksPath, hooks)
}

func writeCodexHooks(path string, hooks map[string][]map[string]any) error {
	data, err := json.MarshalIndent(hooks, "", "  ")
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
		return fmt.Errorf("writing hooks: %w", err)
	}
	return os.Rename(tmp, path)
}
