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
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/pledge"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/terminal"
	"github.com/spf13/cobra"
)

type setOpts struct {
	profile    string
	allowedOps []string
	yes        bool
	sessionID  string
}

var errAdminRequiresConfirm = errors.New("use --yes to confirm setting admin pledge")

func (o *setOpts) Run(cmd *cobra.Command) error {
	prof := pledge.Profile(o.profile)

	switch prof {
	case pledge.ProfileReadonly, pledge.ProfileReadWrite, pledge.ProfileAdmin:
	default:
		return fmt.Errorf("unknown profile %q: choose readonly, read-write, or admin", o.profile)
	}

	if prof == pledge.ProfileAdmin && !o.yes {
		if terminal.IsTerminalInput(cmd.InOrStdin()) {
			fmt.Fprintln(cmd.ErrOrStderr(), "WARNING: admin pledge allows all operations including destructive ones.")
			fmt.Fprint(cmd.ErrOrStderr(), "Confirm? [y/N] ")
			var answer string
			if _, err := fmt.Fscan(cmd.InOrStdin(), &answer); err != nil || !strings.EqualFold(answer, "y") {
				return errors.New("aborted")
			}
		} else {
			return errAdminRequiresConfirm
		}
	}

	// Resolve the session key. --session-id (or ATLAS_PLEDGE_SESSION_ID) overrides the resolver.
	var key pledge.SessionKey
	sid := o.sessionID
	if sid == "" {
		sid = os.Getenv("ATLAS_PLEDGE_SESSION_ID")
	}
	if sid != "" {
		if !pledge.IsValidUUID(sid) {
			return fmt.Errorf("invalid --session-id %q: must be a UUID v4 (xxxxxxxx-xxxx-4xxx-yxxx-xxxxxxxxxxxx)", sid)
		}
		var err error
		key, err = pledge.NewSessionKey("claude", sid)
		if err != nil {
			return err
		}
		// Write a breadcrumb so subsequent atlas invocations with CLAUDECODE=1
		// and CLAUDE_PROJECT_DIR can discover the conversation UUID without
		// ATLAS_PLEDGE_SESSION_ID in their environment.
		if projectDir := os.Getenv("CLAUDE_PROJECT_DIR"); projectDir != "" {
			breadcrumbDir := pledge.ClaudeBreadcrumbDir(projectDir)
			if mkErr := os.MkdirAll(breadcrumbDir, 0o700); mkErr == nil {
				_ = os.WriteFile(filepath.Join(breadcrumbDir, sid), []byte(sid), 0o600)
			}
		}
	} else {
		var err error
		key, err = pledge.ResolveSessionKey()
		if err != nil {
			return fmt.Errorf("pledge is not supported on this platform: %w", err)
		}
	}

	pf, err := pledge.NewPledgeFile(prof, o.allowedOps)
	if err != nil {
		return err
	}

	if err := pledge.Narrow(key, pf); err != nil {
		if errors.Is(err, pledge.ErrWouldWiden) {
			return fmt.Errorf("cannot widen an existing pledge (current session already has a more restrictive pledge); open a new terminal to reset")
		}
		return err
	}

	fmt.Fprintf(cmd.OutOrStdout(), "Pledge set: %s\n", prof)
	return nil
}

// SetBuilder returns the cobra command for `atlas pledge <profile>`.
func SetBuilder() *cobra.Command {
	opts := &setOpts{}
	cmd := &cobra.Command{
		Use:   "set <profile>",
		Short: "Restrict the current session to a permission profile.",
		Long: `Sets a pledge for the current shell session. Valid profiles:
  readonly    — allow only read (GET) operations
  read-write  — allow read and write operations (default Atlas user)
  admin       — allow all operations including org-level destructive actions

Pledges are monotonically narrowing: you can restrict further but never widen.
To reset, open a new terminal session.`,
		Args:    cobra.ExactArgs(1),
		Example: "  atlas pledge set readonly\n  atlas pledge set admin --yes --allow deleteCluster",
		RunE: func(cmd *cobra.Command, args []string) error {
			opts.profile = args[0]
			return opts.Run(cmd)
		},
	}

	cmd.Flags().StringSliceVar(&opts.allowedOps, "allow", nil, "Comma-separated list of operationIDs explicitly permitted regardless of tier.")
	cmd.Flags().BoolVar(&opts.yes, "yes", false, "Skip the admin confirmation prompt.")
	cmd.Flags().StringVar(&opts.sessionID, "session-id", "", "Claude Code conversation UUID; overrides automatic session detection.")

	return cmd
}
