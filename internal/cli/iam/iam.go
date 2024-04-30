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

package iam

import (
	"github.com/andreaangiolillo/mongocli-test/internal/cli"
	"github.com/andreaangiolillo/mongocli-test/internal/cli/iam/globalaccesslists"
	"github.com/andreaangiolillo/mongocli-test/internal/cli/iam/globalapikeys"
	"github.com/andreaangiolillo/mongocli-test/internal/cli/iam/organizations"
	"github.com/andreaangiolillo/mongocli-test/internal/cli/iam/projects"
	"github.com/andreaangiolillo/mongocli-test/internal/cli/iam/teams"
	"github.com/andreaangiolillo/mongocli-test/internal/cli/iam/users"
	"github.com/andreaangiolillo/mongocli-test/internal/config"
	"github.com/andreaangiolillo/mongocli-test/internal/flag"
	"github.com/andreaangiolillo/mongocli-test/internal/log"
	"github.com/andreaangiolillo/mongocli-test/internal/usage"
	"github.com/andreaangiolillo/mongocli-test/internal/validate"
	"github.com/spf13/cobra"
)

func Builder() *cobra.Command {
	opts := &cli.RefresherOpts{}
	var debugLevel bool
	cmd := &cobra.Command{
		Use: "iam",
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			log.SetWriter(cmd.ErrOrStderr())
			if debugLevel {
				log.SetLevel(log.DebugLevel)
			}

			if err := opts.InitFlow(config.Default())(); err != nil {
				return err
			}
			if err := opts.RefreshAccessToken(cmd.Context()); err != nil {
				return err
			}
			if config.Service() == "" {
				config.SetService(config.CloudService)
			}
			return validate.Credentials()
		},
		Short: "Organization and projects operations.",
		Long:  "Identity and Access Management.",
		Annotations: map[string]string{
			"toc": "true",
		},
	}
	cmd.AddCommand(
		projects.Builder(),
		organizations.Builder(),
		globalapikeys.Builder(),
		globalaccesslists.Builder(),
		users.Builder(),
		teams.Builder(),
	)

	cmd.PersistentFlags().BoolVarP(&debugLevel, flag.Debug, flag.DebugShort, false, usage.Debug)
	_ = cmd.PersistentFlags().MarkHidden(flag.Debug)

	return cmd
}
