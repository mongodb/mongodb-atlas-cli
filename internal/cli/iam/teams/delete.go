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

package teams

import (
	"github.com/mongodb/mongocli/internal/cli"
	"github.com/mongodb/mongocli/internal/cli/require"
	"github.com/mongodb/mongocli/internal/config"
	"github.com/mongodb/mongocli/internal/flag"
	"github.com/mongodb/mongocli/internal/store"
	"github.com/mongodb/mongocli/internal/usage"
	"github.com/spf13/cobra"
)

type DeleteOpts struct {
	cli.GlobalOpts
	*cli.DeleteOpts
	store store.TeamDeleter
}

func (opts *DeleteOpts) init() error {
	var err error
	opts.store, err = store.New(config.Default())
	return err
}

func (opts *DeleteOpts) Run() error {
	return opts.Delete(opts.store.DeleteTeam, opts.ConfigOrgID())
}

// mongocli iam team(s) delete <ID> [--force] [--orgId orgId]
func DeleteBuilder() *cobra.Command {
	opts := &DeleteOpts{
		DeleteOpts: cli.NewDeleteOpts("Team '%s' deleted\n", "Team not deleted"),
	}
	cmd := &cobra.Command{
		Use:     "delete <TeamID>",
		Aliases: []string{"rm"},
		Short:   deleteTeam,
		Args:    require.ExactArgs(1),
		PreRunE: func(cmd *cobra.Command, args []string) error {
			opts.Entry = args[0]
			return opts.PreRunE(
				opts.ValidateOrgID,
				opts.init,
				opts.Prompt,
			)
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return opts.Run()
		},
	}

	cmd.Flags().BoolVar(&opts.Confirm, flag.Force, false, usage.Force)

	cmd.Flags().StringVar(&opts.OrgID, flag.OrgID, "", usage.OrgID)

	return cmd
}
