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
	"context"
	"fmt"
	"github.com/mongodb/mongocli/internal/validate"

	"github.com/mongodb/mongocli/internal/cli"
	"github.com/mongodb/mongocli/internal/cli/atlas/quickstart"
	"github.com/mongodb/mongocli/internal/cli/auth"
	"github.com/mongodb/mongocli/internal/config"
	"github.com/mongodb/mongocli/internal/flag"
	"github.com/mongodb/mongocli/internal/usage"
	"github.com/spf13/cobra"
)

const (
	WithProfileMsg = `Run "atlas auth setup --profile <profile_name>" to create a new Atlas account on a new Atlas CLI profile.`
)

type Opts struct {
	cli.GlobalOpts
	cli.WatchOpts
	// quickstart
	quickstart quickstart.Flow
	// register
	register auth.RegisterFlow
	// login
	login *auth.LoginOpts
	// control
	skipRegister bool
	skipLogin    bool
}

func (opts *Opts) Run(ctx context.Context) error {
	if !opts.skipRegister {
		if err := opts.register.Run(ctx); err != nil {
			return err
		}
	}

	if !opts.skipLogin {
		if err := opts.login.Run(ctx); err != nil {
			return err
		}
	}

	if err := opts.quickstart.PreRun(ctx, opts.OutWriter); err != nil {
		return err
	}

	return opts.quickstart.Run()
}

func (opts *Opts) PreRun(ctx context.Context) error {
	opts.skipRegister = false
	opts.skipLogin = true

	if config.PublicAPIKey() != "" && config.PrivateAPIKey() != "" {
		opts.skipRegister = true
		msg := fmt.Sprintf(auth.AlreadyAuthenticatedMsg, config.PublicAPIKey())
		_, _ = fmt.Fprintf(opts.OutWriter, `
%s

%s

`, msg, WithProfileMsg)
	}

	if account, err := auth.AccountWithAccessToken(); err == nil {
		opts.skipRegister = true
		msg := fmt.Sprintf(auth.AlreadyAuthenticatedEmailMsg, account)
		// token exists but it is not refreshed
		if err := cli.RefreshToken(ctx); err != nil || validate.Token() != nil {
			return fmt.Errorf(`%s

%s
%s
`, msg, auth.LoginMsg, WithProfileMsg)
		} else {
			opts.skipLogin = false
			_, _ = fmt.Fprintf(opts.OutWriter, `%s

%s

`, msg, WithProfileMsg)
		}
	}

	return nil
}

// Builder
// atlas setup
//	[--clusterName clusterName]
//	[--provider provider]
//	[--region regionName]
//	[--username username]
//	[--password password]
//	[--skipMongosh skipMongosh]
func Builder() *cobra.Command {
	loginOpts := &auth.LoginOpts{}
	qsOpts := &quickstart.Opts{}
	opts := &Opts{
		register:   auth.NewRegisterFlow(loginOpts),
		login:      loginOpts,
		quickstart: qsOpts,
	}

	cmd := &cobra.Command{
		Use: "setup",
		Example: `Override default cluster settings like name, provider or database username by using the command options
  $ atlas setup --clusterName Test --provider GCP --username dbuserTest
`,
		Short:  "Register, authenticate, create and access an Atlas Cluster.",
		Long:   "This command takes you through registration, login, default profile creation, creating your first free tier cluster and connecting to it using MongoDB Shell.",
		Hidden: true,
		PreRunE: func(cmd *cobra.Command, args []string) error {
			opts.OutWriter = cmd.OutOrStdout()
			// setup pre run
			if err := opts.PreRun(cmd.Context()); err != nil {
				return err
			}

			// registration pre run if applicable
			if !opts.skipRegister {
				if err := opts.register.PreRun(opts.OutWriter); err != nil {
					return err
				}
			}

			if !opts.skipLogin {
				if err := opts.login.PreRun(); err != nil {
					return err
				}
			}
			return opts.PreRunE(
				opts.InitOutput(opts.OutWriter, ""),
			)
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return opts.Run(cmd.Context())
		},
	}

	// Register and login related
	cmd.Flags().BoolVar(&loginOpts.IsGov, "gov", false, "Register to Atlas for Government.")
	cmd.Flags().BoolVar(&loginOpts.NoBrowser, "noBrowser", false, "Don't try to open a browser session.")
	cmd.Flags().BoolVar(&loginOpts.SkipConfig, "skipConfig", false, "Skip profile configuration.")
	// Quickstart related
	cmd.Flags().StringVar(&qsOpts.ClusterName, flag.ClusterName, "", usage.ClusterName)
	cmd.Flags().StringVar(&qsOpts.Tier, flag.Tier, quickstart.DefaultAtlasTier, usage.Tier)
	cmd.Flags().StringVar(&qsOpts.Provider, flag.Provider, "", usage.Provider)
	cmd.Flags().StringVarP(&qsOpts.Region, flag.Region, flag.RegionShort, "", usage.Region)
	cmd.Flags().StringSliceVar(&qsOpts.IPAddresses, flag.AccessListIP, []string{}, usage.NetworkAccessListIPEntry)
	cmd.Flags().StringVar(&qsOpts.DBUsername, flag.Username, "", usage.DBUsername)
	cmd.Flags().StringVar(&qsOpts.DBUserPassword, flag.Password, "", usage.Password)
	cmd.Flags().BoolVar(&qsOpts.SkipSampleData, flag.SkipSampleData, false, usage.SkipSampleData)
	cmd.Flags().BoolVar(&qsOpts.SkipMongosh, flag.SkipMongosh, false, usage.SkipMongosh)
	cmd.Flags().BoolVar(&qsOpts.Confirm, flag.Force, false, usage.Force)

	return cmd
}
