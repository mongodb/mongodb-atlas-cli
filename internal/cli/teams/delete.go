// Copyright 2023 MongoDB Inc
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
	"context"
	"fmt"

	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli/require"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/config"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/flag"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/store"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/usage"
	"github.com/spf13/cobra"
)

type DeleteOpts struct {
	cli.GlobalOpts
	*cli.DeleteOpts
	store store.TeamDeleter
}

func (opts *DeleteOpts) initStore(ctx context.Context) func() error {
	return func() error {
		var err error
		opts.store, err = store.New(store.AuthenticatedPreset(config.Default()), store.WithContext(ctx))
		return err
	}
}

func (opts *DeleteOpts) Run() error {
	return opts.Delete(opts.store.DeleteTeam, opts.ConfigOrgID())
}

// atlas team(s) delete <ID> [--force] [--orgId orgId].
func DeleteBuilder() *cobra.Command {
	opts := &DeleteOpts{
		DeleteOpts: cli.NewDeleteOpts("Team '%s' deleted\n", "Team not deleted"),
	}
	cmd := &cobra.Command{
		Use:     "delete <teamId>",
		Aliases: []string{"rm"},
		Short:   "Remove the specified team from your organization.",
		Long:    fmt.Sprintf(usage.RequiredRole, "Organization User Admin"),
		Args:    require.ExactArgs(1),
		Annotations: map[string]string{
			"teamIdDesc": "Unique 24-digit string that identifies the team.",
			"output":     opts.SuccessMessage(),
		},
		Example: `  # Remove the team with the ID 5e44445ef10fab20b49c0f31 from the organization with ID 5e2211c17a3e5a48f5497de3:
  atlas teams delete 5e44445ef10fab20b49c0f31 --orgId 5e1234c17a3e5a48f5497de3`,
		PreRunE: func(cmd *cobra.Command, args []string) error {
			opts.Entry = args[0]
			return opts.PreRunE(
				opts.ValidateOrgID,
				opts.initStore(cmd.Context()),
				opts.Prompt,
			)
		},
		RunE: func(_ *cobra.Command, _ []string) error {
			return opts.Run()
		},
	}

	cmd.Flags().BoolVar(&opts.Confirm, flag.Force, false, usage.Force)

	cmd.Flags().StringVar(&opts.OrgID, flag.OrgID, "", usage.OrgID)

	return cmd
}
