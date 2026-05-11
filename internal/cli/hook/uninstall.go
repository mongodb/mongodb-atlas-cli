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

type uninstallOpts struct {
	shellPath string
}

func (o *uninstallOpts) agentFor(name string) (agent.Agent, error) {
	switch name {
	case "claude-code":
		return agent.NewClaudeCode(""), nil
	case "codex":
		return agent.NewCodex(), nil
	case "shell":
		return agent.NewShell(o.shellPath), nil
	default:
		return nil, fmt.Errorf("unknown agent %q; supported agents: claude-code, codex, shell", name)
	}
}

// UninstallBuilder returns the cobra command for `atlas hook uninstall <agent>`.
func UninstallBuilder() *cobra.Command {
	opts := &uninstallOpts{}
	cmd := &cobra.Command{
		Use:     "uninstall <agent>",
		Short:   "Remove a pledge hook from an AI agent's configuration.",
		Args:    cobra.ExactArgs(1),
		Example: "  atlas hook uninstall claude-code\n  atlas hook uninstall shell --write ~/.bashrc",
		RunE: func(cmd *cobra.Command, args []string) error {
			a, err := opts.agentFor(args[0])
			if err != nil {
				return err
			}
			if err := a.Uninstall(); err != nil {
				return err
			}
			fmt.Fprintf(cmd.OutOrStdout(), "Hook removed for %s\n", a.Name())
			return nil
		},
	}

	cmd.Flags().StringVar(&opts.shellPath, "write", "", "For shell: path to shell config file to remove the snippet from.")

	return cmd
}
