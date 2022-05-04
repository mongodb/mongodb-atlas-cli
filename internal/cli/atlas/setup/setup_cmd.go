// Copyright 2022 MongoDB Inc
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

package setup

import (
	"github.com/mongodb/mongocli/internal/cli"
	"github.com/mongodb/mongocli/internal/cli/atlas/quickstart"
	"github.com/mongodb/mongocli/internal/cli/auth"
	"github.com/mongodb/mongocli/internal/flag"
	"github.com/mongodb/mongocli/internal/usage"
	"github.com/spf13/cobra"
)

type Opts struct {
	cli.GlobalOpts
	cli.WatchOpts
	// quickstart
	quickstart quickstart.Opts
	// register
	register auth.RegisterOpts
	// login
	login auth.LoginOpts
}

func (opts *Opts) Run() error {
	//TODO: Registration flow
	//TODO: Quickstart flow
	return nil
}

// Builder
// atlas setup
//	[--clusterName clusterName]
//	[--provider provider]
//	[--region regionName]
//	[--projectId projectId]
//	[--username username]
//	[--password password]
//	[--skipMongosh skipMongosh]
//	[--default]
func Builder() *cobra.Command {
	opts := &Opts{}
	cmd := &cobra.Command{
		Use: "setup",
		Example: `Override default cluster settings like name, provider or database username by using the command options
  $ atlas setup --clusterName Test --provider GCP --username dbuserTest
`,
		Short:  "Register, authenticate, create and access an Atlas Cluster.",
		Long:   "This command takes you through registration, login, default profile creation, creating your first free tier cluster and connecting to it using MongoDB Shell.",
		Hidden: true,
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return opts.PreRunE(
				opts.InitOutput(cmd.OutOrStdout(), ""),
			)
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return opts.Run()
		},
	}

	// Register and login related
	cmd.Flags().BoolVar(&opts.login.IsGov, "gov", false, "Register to Atlas for Government.")
	cmd.Flags().BoolVar(&opts.login.NoBrowser, "noBrowser", false, "Don't try to open a browser session.")
	cmd.Flags().BoolVar(&opts.login.SkipConfig, "skipConfig", false, "Skip profile configuration.")
	// Quickstart related
	cmd.Flags().StringVar(&opts.quickstart.ClusterName, flag.ClusterName, "", usage.ClusterName)
	cmd.Flags().StringVar(&opts.quickstart.Tier, flag.Tier, quickstart.DefaultAtlasTier, usage.Tier)
	cmd.Flags().StringVar(&opts.quickstart.Provider, flag.Provider, "", usage.Provider)
	cmd.Flags().StringVarP(&opts.quickstart.Region, flag.Region, flag.RegionShort, "", usage.Region)
	cmd.Flags().StringSliceVar(&opts.quickstart.IPAddresses, flag.AccessListIP, []string{}, usage.NetworkAccessListIPEntry)
	cmd.Flags().StringVar(&opts.quickstart.DBUsername, flag.Username, "", usage.DBUsername)
	cmd.Flags().StringVar(&opts.quickstart.DBUserPassword, flag.Password, "", usage.Password)
	cmd.Flags().BoolVar(&opts.quickstart.SkipSampleData, flag.SkipSampleData, false, usage.SkipSampleData)
	cmd.Flags().BoolVar(&opts.quickstart.SkipMongosh, flag.SkipMongosh, false, usage.SkipMongosh)
	cmd.Flags().BoolVar(&opts.quickstart.Confirm, flag.Force, false, usage.Force)

	return cmd
}
