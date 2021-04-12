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
	"runtime"

	"github.com/mongodb/mongocli/internal/cli/atlas"
	"github.com/mongodb/mongocli/internal/cli/cloudmanager"
	cliconfig "github.com/mongodb/mongocli/internal/cli/config"
	"github.com/mongodb/mongocli/internal/cli/iam"
	"github.com/mongodb/mongocli/internal/cli/opsmanager"
	"github.com/mongodb/mongocli/internal/cli/require"
	"github.com/mongodb/mongocli/internal/config"
	"github.com/mongodb/mongocli/internal/flag"
	"github.com/mongodb/mongocli/internal/search"
	"github.com/mongodb/mongocli/internal/usage"
	"github.com/mongodb/mongocli/internal/version"
	"github.com/spf13/cobra"
)

var (
	rootCmd *cobra.Command

	completionCmd = &cobra.Command{
		Use:   "completion <bash|zsh|fish|powershell>",
		Args:  require.ExactValidArgs(1),
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
	rootCmd := &cobra.Command{
		Version: version.Version,
		Use:     config.ToolName,
		Short:   "CLI tool to manage your MongoDB Cloud",
		Long:    fmt.Sprintf("Use %s command help for information on a specific command", config.ToolName),
		Example: `
  Display the help menu for the config command
  $ mongocli config --help`,
		SilenceUsage: true,
	}
	rootCmd.SetVersionTemplate(formattedVersion())
	hasArgs := len(argsWithoutProg) != 0

	if hasArgs && (argsWithoutProg[0] == "--version" || argsWithoutProg[0] == "-v") {
		return rootCmd
	}
	rootCmd.AddCommand(cliconfig.Builder())

	shouldIncludeAtlas := []string{
		atlas.Use,
		"help",
		"--help",
		"-h",
		"completion",
		"__complete",
	}
	if !hasArgs || search.StringInSlice(shouldIncludeAtlas, argsWithoutProg[0]) {
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

const verTemplate = `%s version: %s
git version: %s
Go version: %s
   os: %s
   arch: %s
   compiler: %s
`

func formattedVersion() string {
	return fmt.Sprintf(verTemplate,
		config.ToolName,
		version.Version,
		version.GitCommit,
		runtime.Version(),
		runtime.GOOS,
		runtime.GOARCH,
		runtime.Compiler)
}
