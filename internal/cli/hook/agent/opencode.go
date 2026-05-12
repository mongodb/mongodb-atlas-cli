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
	_ "embed"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

//go:embed opencode_plugin.ts
var opencodePluginTemplate string

type opencodeAgent struct {
	pluginPath string
}

// NewOpencode returns an Agent for opencode (~/.config/opencode/plugins/atlas-pledge.ts).
func NewOpencode() Agent {
	home, _ := os.UserHomeDir()
	return &opencodeAgent{
		pluginPath: filepath.Join(home, ".config", "opencode", "plugins", "atlas-pledge.ts"),
	}
}

func (a *opencodeAgent) Name() string { return "opencode" }

func (a *opencodeAgent) Status() State {
	if _, err := os.Stat(a.pluginPath); err == nil {
		return StateInstalled
	}
	return StateUninstalled
}

func (a *opencodeAgent) Install(opts InstallOpts) error {
	profile := opts.Profile
	if profile == "" {
		profile = "readonly"
	}
	content := strings.ReplaceAll(opencodePluginTemplate, "{{.Profile}}", profile)

	// Idempotency: skip if content already matches.
	if existing, err := os.ReadFile(a.pluginPath); err == nil && string(existing) == content {
		fmt.Fprintf(os.Stderr, "opencode plugin: already installed, skipping\n")
		return nil
	}

	if err := os.MkdirAll(filepath.Dir(a.pluginPath), 0o700); err != nil {
		return fmt.Errorf("creating opencode plugins directory: %w", err)
	}

	if err := os.WriteFile(a.pluginPath, []byte(content), 0o600); err != nil {
		return fmt.Errorf("writing opencode plugin: %w", err)
	}

	fmt.Fprintf(os.Stderr, "opencode plugin: installed atlas pledge plugin at %s\n", a.pluginPath)
	return nil
}

func (a *opencodeAgent) Uninstall() error {
	if _, err := os.Stat(a.pluginPath); os.IsNotExist(err) {
		fmt.Fprintf(os.Stderr, "opencode plugin: nothing to remove\n")
		return nil
	}
	if err := os.Remove(a.pluginPath); err != nil {
		return fmt.Errorf("removing opencode plugin: %w", err)
	}
	fmt.Fprintf(os.Stderr, "opencode plugin: removed\n")
	return nil
}
