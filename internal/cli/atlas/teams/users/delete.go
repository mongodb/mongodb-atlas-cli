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

package users

import (
	"context"
	"fmt"

	"github.com/andreangiolillo/mongocli-test/internal/cli"
	"github.com/andreangiolillo/mongocli-test/internal/cli/require"
	"github.com/andreangiolillo/mongocli-test/internal/config"
	"github.com/andreangiolillo/mongocli-test/internal/flag"
	store "github.com/andreangiolillo/mongocli-test/internal/store/atlas"
	"github.com/andreangiolillo/mongocli-test/internal/usage"
	"github.com/spf13/cobra"
)

type DeleteOpts struct {
	cli.GlobalOpts
	*cli.DeleteOpts
	store  store.TeamUserRemover
	teamID string
}

func (opts *DeleteOpts) initStore(ctx context.Context) func() error {
	return func() error {
		var err error
		opts.store, err = store.New(store.AuthenticatedPreset(config.Default()), store.WithContext(ctx))
		return err
	}
}

func (opts *DeleteOpts) Run() error {
	return opts.Delete(opts.store.RemoveUserFromTeam, opts.ConfigOrgID(), opts.teamID)
}

// atlas team(s) users(s) delete <ID> [--force] --orgId orgId --teamId teamId.
func DeleteBuilder() *cobra.Command {
	opts := &DeleteOpts{
		DeleteOpts: cli.NewDeleteOpts("User '%s' deleted from the team\n", "User not deleted"),
	}
	cmd := &cobra.Command{
		Use:     "delete <userId>",
		Aliases: []string{"rm"},
		Short:   "Remove the specified user from a team for your organization.",
		Long:    fmt.Sprintf(usage.RequiredRole, "Organization User Admin"),
		Args:    require.ExactArgs(1),
		Example: fmt.Sprintf(`  # Remove the user with the ID 5dd58c647a3e5a6c5bce46c7 from the team with the ID 5f6a5c6c713184005d72fe6e for the organization with ID 5e2211c17a3e5a48f5497de3:
  %s teams users delete 5dd58c647a3e5a6c5bce46c7 --teamId 5f6a5c6c713184005d72fe6e --orgId 5e1234c17a3e5a48f5497de3`, cli.ExampleAtlasEntryPoint()),
		Annotations: map[string]string{
			"userIdDesc": "Unique 24-digit string that identifies the user.",
			"output":     opts.SuccessMessage(),
		},
		PreRunE: func(cmd *cobra.Command, args []string) error {
			opts.Entry = args[0]
			return opts.PreRunE(
				opts.ValidateOrgID,
				opts.initStore(cmd.Context()),
				opts.Prompt,
			)
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return opts.Run()
		},
	}

	cmd.Flags().StringVar(&opts.teamID, flag.TeamID, "", usage.TeamID)

	cmd.Flags().BoolVar(&opts.Confirm, flag.Force, false, usage.Force)
	cmd.Flags().StringVar(&opts.OrgID, flag.OrgID, "", usage.OrgID)

	_ = cmd.MarkFlagRequired(flag.TeamID)

	return cmd
}
