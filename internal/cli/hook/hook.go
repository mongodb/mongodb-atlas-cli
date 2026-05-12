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

package hook

import (
	"github.com/spf13/cobra"
)

func Builder() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "hook",
		Short: "Manage AI agent session hooks for pledge auto-install.",
		Long: `Installs or removes pledge hooks in AI coding agent configurations.

When installed, the hook runs "atlas pledge set <profile>" at the start of each
new agent session, automatically restricting the CLI to the configured profile.

Claude Code note: each Bash tool invocation in Claude runs in a separate shell
session. The installed hook reads Claude's session_id from the SessionStart
payload and keys the pledge on the conversation UUID rather than the POSIX
session ID. This means the pledge applies consistently across all Bash calls
within a single Claude conversation.

When you resume a conversation with "claude --resume <uuid>", atlas pledge
recognises the same UUID and the existing pledge namespace is reused.

Supported agents: claude-code, codex, pi, opencode, shell`,
	}

	cmd.AddCommand(
		InstallBuilder(),
		UninstallBuilder(),
	)

	return cmd
}
