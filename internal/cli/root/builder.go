// Copyright 2021 MongoDB Inc
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

package root

import (
	"fmt"

	"github.com/mongodb/mongocli/internal/cli"
	"github.com/mongodb/mongocli/internal/cli/atlas"
	"github.com/mongodb/mongocli/internal/cli/cloudmanager"
	cliconfig "github.com/mongodb/mongocli/internal/cli/config"
	"github.com/mongodb/mongocli/internal/cli/iam"
	"github.com/mongodb/mongocli/internal/cli/opsmanager"
	"github.com/mongodb/mongocli/internal/flag"
	"github.com/mongodb/mongocli/internal/search"
	"github.com/mongodb/mongocli/internal/usage"
	"github.com/spf13/cobra"
)

var (
	rootCmd *cobra.Command

	completionCmd = &cobra.Command{
		Use:   "completion <bash|zsh|fish|powershell>",
		Args:  cobra.ExactValidArgs(1),
		Short: "Generate shell completion scripts",
		Long: `Generate shell completion scripts for MongoDB CLI commands.
The output of this command will be computer code and is meant to be saved to a
file or immediately evaluated by an interactive shell.

When installing MongoDB CLI through brew, it's possible that
no additional shell configuration is necessary, see https://docs.brew.sh/Shell-Completion.`,
		ValidArgs: []string{"bash", "zsh", "powershell", "fish"},
		RunE: func(cmd *cobra.Command, args []string) error {
			switch args[0] {
			case "bash":
				return rootCmd.GenBashCompletion(cmd.OutOrStdout())
			case "zsh":
				return rootCmd.GenZshCompletion(cmd.OutOrStdout())
			case "powershell":
				return rootCmd.GenPowerShellCompletion(cmd.OutOrStdout())
			case "fish":
				return rootCmd.GenFishCompletion(cmd.OutOrStdout(), true)
			default:
				return fmt.Errorf("unsupported shell type %q", args[0])
			}
		},
	}
)

// rootBuilder conditionally adds children commands as needed.
// This is important in particular for Atlas as it dynamically sets flags for cluster creation and
// this can be slow to timeout on environments with limited internet access (Ops Manager)
func Builder(profile *string, argsWithoutProg []string) *cobra.Command {
	rootCmd = cli.Builder()
	hasArgs := len(argsWithoutProg) != 0

	if hasArgs && argsWithoutProg[0] == "--version" {
		return nil
	}
	rootCmd.AddCommand(cliconfig.Builder())

	// We realistically only care about not adding Atlas to the chain of commands
	// but given we can reuse the pattern for the rest we can save some initialization time here
	if !hasArgs || search.StringInSlice([]string{atlas.Use, "completion", "__complete"}, argsWithoutProg[0]) {
		rootCmd.AddCommand(atlas.Builder())
	}
	rootCmd.AddCommand(
		cloudmanager.Builder(),
		opsmanager.Builder(),
		iam.Builder(),
		completionCmd,
	)

	rootCmd.PersistentFlags().StringVarP(profile, flag.Profile, flag.ProfileShort, "", usage.Profile)

	return rootCmd
}
