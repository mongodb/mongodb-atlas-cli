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

package pledge

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/pledge"
	"github.com/spf13/cobra"
)

// claudeHookPayload is the subset of fields Claude Code sends to hook commands
// via stdin for all hook events.
type claudeHookPayload struct {
	SessionID string `json:"session_id"`
}

// readClaudeHookSessionID reads at most 8 KB from r, parses it as the Claude
// Code hook JSON payload, and returns the session_id if it is a valid UUID.
func readClaudeHookSessionID(r io.Reader) (string, bool) {
	buf, err := io.ReadAll(io.LimitReader(r, 8192))
	if err != nil || len(buf) == 0 {
		return "", false
	}
	var p claudeHookPayload
	if err := json.Unmarshal(buf, &p); err != nil {
		return "", false
	}
	if pledge.IsValidUUID(p.SessionID) {
		return p.SessionID, true
	}
	return "", false
}

// SetClaudeCodeBuilder returns the hidden cobra command for
// `atlas pledge set-claude-code <profile> --yes`.
//
// This command is designed to be called from Claude Code's SessionStart hook.
// It reads the hook's JSON stdin payload to extract the conversation UUID
// (session_id) and keys the pledge on that UUID so that subsequent Bash tool
// invocations within the same Claude conversation share the same pledge file.
func SetClaudeCodeBuilder() *cobra.Command {
	var yes bool
	cmd := &cobra.Command{
		Use:    "set-claude-code <profile>",
		Short:  "Set a pledge from a Claude Code SessionStart hook (reads session_id from stdin JSON).",
		Hidden: true,
		Args:   cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			prof := pledge.Profile(args[0])
			switch prof {
			case pledge.ProfileReadonly, pledge.ProfileReadWrite, pledge.ProfileAdmin:
			default:
				return fmt.Errorf("unknown profile %q: choose readonly, read-write, or admin", args[0])
			}
			if prof == pledge.ProfileAdmin && !yes {
				return errors.New("use --yes to confirm setting admin pledge")
			}

			// Extract the conversation UUID from the Claude Code hook JSON stdin.
			uuid, found := readClaudeHookSessionID(cmd.InOrStdin())

			var key pledge.SessionKey
			if found {
				var err error
				key, err = pledge.NewSessionKey("claude", uuid)
				if err != nil {
					return err
				}
				// Write a breadcrumb for any later atlas invocations that resolve
				// via CLAUDE_PROJECT_DIR rather than CLAUDE_CODE_SESSION_ID.
				if projectDir := os.Getenv("CLAUDE_PROJECT_DIR"); projectDir != "" {
					breadcrumbDir := pledge.ClaudeBreadcrumbDir(projectDir)
					if mkErr := os.MkdirAll(breadcrumbDir, 0o700); mkErr == nil {
						_ = os.WriteFile(filepath.Join(breadcrumbDir, uuid), []byte(uuid), 0o600)
					}
				}
			} else {
				// Fall back to the standard resolver (CLAUDE_CODE_SESSION_ID env,
				// breadcrumb, or POSIX SID).
				var err error
				key, err = pledge.ResolveSessionKey()
				if err != nil {
					return fmt.Errorf("pledge is not supported on this platform: %w", err)
				}
			}

			pf, err := pledge.NewPledgeFile(prof, nil)
			if err != nil {
				return err
			}
			if err := pledge.Narrow(key, pf); err != nil {
				if errors.Is(err, pledge.ErrWouldWiden) {
					return fmt.Errorf("cannot widen an existing pledge; open a new terminal to reset")
				}
				return err
			}

			fmt.Fprintf(cmd.OutOrStdout(), "Pledge set: %s\n", prof)
			return nil
		},
	}

	cmd.Flags().BoolVar(&yes, "yes", false, "Skip the admin confirmation prompt.")
	return cmd
}
