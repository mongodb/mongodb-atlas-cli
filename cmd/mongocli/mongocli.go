// Copyright 2020 MongoDB Inc
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

package main

import (
	"fmt"
	"log"
	"os"

	"github.com/mongodb/mongocli/internal/cli"
	"github.com/mongodb/mongocli/internal/cli/atlas"
	"github.com/mongodb/mongocli/internal/cli/cloudmanager"
	cliconfig "github.com/mongodb/mongocli/internal/cli/config"
	"github.com/mongodb/mongocli/internal/cli/iam"
	"github.com/mongodb/mongocli/internal/cli/opsmanager"
	"github.com/mongodb/mongocli/internal/config"
	"github.com/mongodb/mongocli/internal/flag"
	"github.com/mongodb/mongocli/internal/usage"
	"github.com/spf13/cobra"
)

var (
	profile string

	rootCmd = cli.Builder()

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

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	rootBuilder()
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

// rootBuilder conditionally adds children commands as needed.
// This is important in particular for Atlas as it dynamically sets flags for cluster creation and
// this can be slow to timeout on environments with limited internet access (Ops Manager)
func rootBuilder() {
	argsWithoutProg := os.Args[1:]
	if len(argsWithoutProg) != 0 && argsWithoutProg[0] == "--version" {
		return
	}
	rootCmd.AddCommand(cliconfig.Builder())

	// We realistically only care about not adding Atlas to the chain of commands
	// but given we can reuse the pattern for the rest we can save some initialization time here
	if len(argsWithoutProg) == 0 || argsWithoutProg[0] == atlas.Use {
		rootCmd.AddCommand(atlas.Builder())
	}
	if len(argsWithoutProg) == 0 || argsWithoutProg[0] == cloudmanager.Use {
		cloudmanager.Builder()
	}
	if len(argsWithoutProg) == 0 || argsWithoutProg[0] == opsmanager.Use {
		opsmanager.Builder()
	}
	if len(argsWithoutProg) == 0 || argsWithoutProg[0] == iam.Use {
		iam.Builder()
	}
	rootCmd.AddCommand(
		completionCmd,
	)
	rootCmd.PersistentFlags().StringVarP(&profile, flag.Profile, flag.ProfileShort, "", usage.Profile)
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if err := config.Load(); err != nil {
		log.Fatalf("Error loading config: %v", err)
	}

	availableProfiles := config.List()
	if profile != "" {
		config.SetName(profile)
	} else if len(availableProfiles) == 1 {
		config.SetName(availableProfiles[0])
	}
}

func main() {
	cobra.EnableCommandSorting = false
	cobra.OnInitialize(initConfig)

	Execute()
}
