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
	"fmt"

	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli/hook/agent"
	"github.com/spf13/cobra"
)

type installOpts struct {
	profile      string
	projectLevel bool
	shellPath    string
}

func (o *installOpts) agentFor(name string) (agent.Agent, error) {
	switch name {
	case "claude-code":
		var path string
		if o.projectLevel {
			path = ".claude/settings.json"
		}
		return agent.NewClaudeCode(path), nil
	case "codex":
		return agent.NewCodex(), nil
	case "pi":
		return agent.NewPi(), nil
	case "opencode":
		return agent.NewOpencode(), nil
	case "shell":
		return agent.NewShell(o.shellPath), nil
	default:
		return nil, fmt.Errorf("unknown agent %q; supported agents: claude-code, codex, pi, opencode, shell", name)
	}
}

// InstallBuilder returns the cobra command for `atlas hook install <agent>`.
func InstallBuilder() *cobra.Command {
	opts := &installOpts{}
	cmd := &cobra.Command{
		Use:   "install <agent>",
		Short: "Install a pledge hook in an AI agent's configuration.",
		Long: `Installs a hook that runs "atlas pledge set <profile> --yes" at the start of every new agent session.

Supported agents:
  claude-code  Edits ~/.claude/settings.json (or ./.claude/settings.json with --project).
  codex        Edits ~/.codex/hooks.json (PreToolUse, since Codex has no SessionStart event).
  pi           Writes ~/.pi/agent/extensions/atlas-pledge.ts. Reload pi with /reload after install.
  opencode     Writes ~/.config/opencode/plugins/atlas-pledge.ts. Loaded automatically by opencode.
  shell        Prints a sourceable snippet; use --write to inject into a shell config file.`,
		Args:      cobra.ExactArgs(1),
		ValidArgs: []string{"claude-code", "codex", "pi", "opencode", "shell"},
		Example:   "  atlas hook install claude-code\n  atlas hook install claude-code --profile read-write\n  atlas hook install codex\n  atlas hook install pi\n  atlas hook install opencode\n  atlas hook install shell --write ~/.bashrc",
		RunE: func(cmd *cobra.Command, args []string) error {
			a, err := opts.agentFor(args[0])
			if err != nil {
				return err
			}
			if err := a.Install(agent.InstallOpts{
				Profile:      opts.profile,
				ProjectLevel: opts.projectLevel,
			}); err != nil {
				return err
			}
			fmt.Fprintf(cmd.OutOrStdout(), "Hook installed for %s (profile: %s)\n", a.Name(), opts.profile)
			return nil
		},
	}

	cmd.Flags().StringVar(&opts.profile, "profile", "readonly", "Pledge profile to set at session start (readonly, read-write, admin).")
	cmd.Flags().BoolVar(&opts.projectLevel, "project", false, "For claude-code: install in ./.claude/settings.json instead of ~/.claude/settings.json.")
	cmd.Flags().StringVar(&opts.shellPath, "write", "", "For shell: path to shell config file to inject the snippet into.")

	return cmd
}
