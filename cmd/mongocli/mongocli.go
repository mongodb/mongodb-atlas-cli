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

	"github.com/mongodb/mongocli/internal/cli/atlas"
	"github.com/mongodb/mongocli/internal/cli/cloudmanager"
	cliconfig "github.com/mongodb/mongocli/internal/cli/config"
	"github.com/mongodb/mongocli/internal/cli/iam"
	"github.com/mongodb/mongocli/internal/cli/opsmanager"
	"github.com/mongodb/mongocli/internal/config"
	"github.com/mongodb/mongocli/internal/flag"
	"github.com/mongodb/mongocli/internal/usage"
	"github.com/mongodb/mongocli/internal/version"
	"github.com/spf13/cobra"
)

var (
	rootCmd = &cobra.Command{
		Version: version.Version,
		Use:     config.ToolName,
		Short:   "CLI tool to manage your MongoDB Cloud",
		Long:    fmt.Sprintf("Use %s command help for information on a specific command", config.ToolName),
		Example: `
  Display the help menu for the config command
  $ mongocli config --help`,
		SilenceUsage: true,
	}

	completionCmd = &cobra.Command{
		Use:       "completion <name>",
		Args:      cobra.ExactValidArgs(1),
		ValidArgs: []string{"bash", "zsh"},
		Hidden:    true,
		RunE: func(cmd *cobra.Command, args []string) error {
			if args[0] == "bash" {
				return rootCmd.GenBashCompletion(os.Stdout)
			}
			return rootCmd.GenZshCompletion(os.Stdout)
		},
	}
)

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
	// config commands
	rootCmd.AddCommand(cliconfig.Builder())
	// Atlas commands
	rootCmd.AddCommand(atlas.Builder())
	// CM commands
	rootCmd.AddCommand(cloudmanager.Builder())
	// OM commands
	rootCmd.AddCommand(opsmanager.Builder())
	// IAM commands
	rootCmd.AddCommand(iam.Builder())

	rootCmd.AddCommand(completionCmd)

	cobra.EnableCommandSorting = false

	profile := rootCmd.PersistentFlags().StringP(flag.Profile, flag.ProfileShort, config.DefaultProfile, usage.Profile)
	config.SetName(profile)
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if err := config.Load(); err != nil {
		log.Fatalf("Error loading config: %v", err)
	}
}

func main() {
	Execute()
}
