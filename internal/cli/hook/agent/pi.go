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

//go:embed pi_extension.ts
var piExtensionTemplate string

type piAgent struct {
	extensionPath string
}

// NewPi returns an Agent for pi (~/.pi/agent/extensions/atlas-pledge.ts).
func NewPi() Agent {
	home, _ := os.UserHomeDir()
	return &piAgent{
		extensionPath: filepath.Join(home, ".pi", "agent", "extensions", "atlas-pledge.ts"),
	}
}

func (a *piAgent) Name() string { return "pi" }

func (a *piAgent) Status() State {
	if _, err := os.Stat(a.extensionPath); err == nil {
		return StateInstalled
	}
	return StateUninstalled
}

func (a *piAgent) Install(opts InstallOpts) error {
	profile := opts.Profile
	if profile == "" {
		profile = "readonly"
	}
	content := strings.ReplaceAll(piExtensionTemplate, "{{.Profile}}", profile)

	// Idempotency: skip if content already matches.
	if existing, err := os.ReadFile(a.extensionPath); err == nil && string(existing) == content {
		fmt.Fprintf(os.Stderr, "pi extension: already installed, skipping\n")
		return nil
	}

	if err := os.MkdirAll(filepath.Dir(a.extensionPath), 0o700); err != nil {
		return fmt.Errorf("creating pi extensions directory: %w", err)
	}

	if err := os.WriteFile(a.extensionPath, []byte(content), 0o600); err != nil {
		return fmt.Errorf("writing pi extension: %w", err)
	}

	fmt.Fprintf(os.Stderr, "pi extension: installed atlas pledge extension\n")
	fmt.Fprintf(os.Stderr, "\nAtlas pledge extension installed for pi. Extension: %s\nReload pi with `/reload` or restart pi.\n", a.extensionPath)
	return nil
}

func (a *piAgent) Uninstall() error {
	if _, err := os.Stat(a.extensionPath); os.IsNotExist(err) {
		fmt.Fprintf(os.Stderr, "pi extension: nothing to remove\n")
		return nil
	}
	if err := os.Remove(a.extensionPath); err != nil {
		return fmt.Errorf("removing pi extension: %w", err)
	}
	fmt.Fprintf(os.Stderr, "pi extension: removed\n")
	return nil
}
