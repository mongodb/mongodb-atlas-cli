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

package cmd

import (
	"fmt"
	"log"

	"github.com/mongodb/mongocli/internal/cli"
	"github.com/mongodb/mongocli/internal/config"
	"github.com/mongodb/mongocli/internal/flags"
	"github.com/mongodb/mongocli/internal/usage"
	"github.com/mongodb/mongocli/internal/version"
	"github.com/spf13/cobra"
)

var (
	rootCmd = &cobra.Command{
		Version: version.Version,
		Use:     config.Name,
		Short:   "CLI tool to manage your MongoDB Cloud",
		Long:    fmt.Sprintf("Use %s command help for information on a specific command", config.Name),
	}
)

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	_ = rootCmd.Execute()
}

func init() {
	cobra.OnInitialize(initConfig)
	// config commands
	rootCmd.AddCommand(cli.ConfigBuilder())
	// Atlas commands
	rootCmd.AddCommand(cli.AtlasBuilder())
	// CM commands
	rootCmd.AddCommand(cli.CloudManagerBuilder())
	// OM commands
	rootCmd.AddCommand(cli.OpsManagerBuilder())
	// IAM commands
	rootCmd.AddCommand(cli.IAMBuilder())

	cobra.EnableCommandSorting = false

	profile := rootCmd.PersistentFlags().StringP(flags.Profile, flags.ProfileShort, config.DefaultProfile, usage.Profile)
	config.SetName(profile)
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if err := config.Load(); err != nil {
		log.Fatalf("Error loading config: %v", err)
	}
}
