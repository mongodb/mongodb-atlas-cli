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
	"bufio"
	"bytes"
	"fmt"
	"os"
	"strings"
)

const (
	shellBlockBegin = "# >>> atlas pledge >>>"
	shellBlockEnd   = "# <<< atlas pledge <<<"
)

// shellAgent implements Agent for generic POSIX shells via snippet injection.
type shellAgent struct {
	// writePath is the file to write the snippet to (e.g., ~/.bashrc).
	// If empty, Install prints the snippet to stdout and returns nil.
	writePath string
}

// NewShell returns a shell agent. writePath is the file to modify (may be empty).
func NewShell(writePath string) Agent {
	return &shellAgent{writePath: writePath}
}

func (a *shellAgent) Name() string { return "shell" }

func snippet(profile string) string {
	if profile == "" {
		profile = "readonly"
	}
	return fmt.Sprintf("%s\nif command -v atlas > /dev/null 2>&1; then\n  atlas pledge set %s --yes 2>/dev/null || true\nfi\n%s",
		shellBlockBegin, profile, shellBlockEnd)
}

func (a *shellAgent) Status() State {
	if a.writePath == "" {
		return StateUnknown
	}
	raw, err := os.ReadFile(a.writePath)
	if err != nil {
		return StateUninstalled
	}
	if strings.Contains(string(raw), shellBlockBegin) {
		return StateInstalled
	}
	return StateUninstalled
}

func (a *shellAgent) Install(opts InstallOpts) error {
	snip := snippet(opts.Profile)

	if a.writePath == "" {
		fmt.Println(snip)
		return nil
	}

	// Read existing file.
	var existing []byte
	if raw, err := os.ReadFile(a.writePath); err == nil {
		existing = raw
	}

	// Replace existing block or append.
	updated := replaceOrAppendBlock(string(existing), snip)

	tmp := a.writePath + ".tmp"
	if err := os.WriteFile(tmp, []byte(updated), 0o600); err != nil {
		return fmt.Errorf("writing shell config: %w", err)
	}
	return os.Rename(tmp, a.writePath)
}

func (a *shellAgent) Uninstall() error {
	if a.writePath == "" {
		return nil
	}
	raw, err := os.ReadFile(a.writePath)
	if err != nil {
		return nil
	}
	updated := removeBlock(string(raw))
	tmp := a.writePath + ".tmp"
	if err := os.WriteFile(tmp, []byte(updated), 0o600); err != nil {
		return fmt.Errorf("writing shell config: %w", err)
	}
	return os.Rename(tmp, a.writePath)
}

// replaceOrAppendBlock replaces the managed block if present, or appends it.
func replaceOrAppendBlock(content, newBlock string) string {
	start := strings.Index(content, shellBlockBegin)
	end := strings.Index(content, shellBlockEnd)
	if start >= 0 && end >= 0 && end > start {
		return content[:start] + newBlock + content[end+len(shellBlockEnd):]
	}
	if content != "" && !strings.HasSuffix(content, "\n") {
		content += "\n"
	}
	return content + newBlock + "\n"
}

// removeBlock removes the managed block from content.
func removeBlock(content string) string {
	var buf bytes.Buffer
	inBlock := false
	scanner := bufio.NewScanner(strings.NewReader(content))
	for scanner.Scan() {
		line := scanner.Text()
		if strings.TrimSpace(line) == shellBlockBegin {
			inBlock = true
			continue
		}
		if inBlock && strings.TrimSpace(line) == shellBlockEnd {
			inBlock = false
			continue
		}
		if !inBlock {
			buf.WriteString(line)
			buf.WriteByte('\n')
		}
	}
	return buf.String()
}
