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

	"github.com/mongodb/atlas-cli-core/config"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli/require"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/flag"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/store"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/usage"
	"github.com/spf13/cobra"
	atlasv2 "go.mongodb.org/atlas-sdk/v20250312012/admin"
)

const addTemplate = "User(s) added to the team.\n"

//go:generate go tool go.uber.org/mock/mockgen -typed -destination=add_mock_test.go -package=users . TeamAdder

type TeamAdder interface {
	AddUsersToTeam(string, string, []atlasv2.AddUserToTeam) (*atlasv2.PaginatedApiAppUser, error)
}

type AddOpts struct {
	cli.OrgOpts
	cli.OutputOpts
	store  TeamAdder
	teamID string
	users  []string
}

func (opts *AddOpts) initStore(ctx context.Context) func() error {
	return func() error {
		var err error
		opts.store, err = store.New(store.AuthenticatedPreset(config.Default()), store.WithContext(ctx))
		return err
	}
}

func (opts *AddOpts) Run() error {
	r, err := opts.store.AddUsersToTeam(opts.ConfigOrgID(), opts.teamID, opts.newUsers())
	if err != nil {
		return err
	}

	return opts.Print(r)
}

func (opts *AddOpts) newUsers() []atlasv2.AddUserToTeam {
	usersToAdd := make([]atlasv2.AddUserToTeam, len(opts.users))
	for i, user := range opts.users {
		usersToAdd[i] = atlasv2.AddUserToTeam{
			Id: user,
		}
	}
	return usersToAdd
}

// atlas team(s) user(s) add <userId> [userId]... --teamId teamId --orgId orgId.
func AddBuilder() *cobra.Command {
	opts := &AddOpts{}
	cmd := &cobra.Command{
		Use:   "add <userId>...",
		Args:  require.MinimumNObjectIDArgs(1),
		Short: "Add the specified MongoDB user to a team for your organization.",
		Long: `You can add users that are part of the organization or users that have been sent an invitation to join the organization.

` + fmt.Sprintf(usage.RequiredRole, "Organization User Admin"),
		Annotations: map[string]string{
			"userIdDesc": "Unique 24-digit string that identifies the user. You can add more than one user at a time by specifying multiple user IDs separated by a space.",
			"output":     addTemplate,
		},
		Example: `  # Add the users with the IDs 5dd58c647a3e5a6c5bce46c7 and 5dd56c847a3e5a1f363d424d to the team with the ID 5f6a5c6c713184005d72fe6e for the organization with ID 5e2211c17a3e5a48f5497de3:
  atlas teams users add 5dd58c647a3e5a6c5bce46c7 5dd56c847a3e5a1f363d424d --teamId 5f6a5c6c713184005d72fe6e --orgId 5e1234c17a3e5a48f5497de3 --output json`,
		PreRunE: func(cmd *cobra.Command, args []string) error {
			opts.users = args
			return opts.PreRunE(
				opts.ValidateOrgID,
				opts.initStore(cmd.Context()),
				opts.InitOutput(cmd.OutOrStdout(), addTemplate),
			)
		},
		RunE: func(_ *cobra.Command, _ []string) error {
			return opts.Run()
		},
	}

	cmd.Flags().StringVar(&opts.teamID, flag.TeamID, "", usage.TeamID)

	opts.AddOrgOptFlags(cmd)
	opts.AddOutputOptFlags(cmd)

	_ = cmd.MarkFlagRequired(flag.TeamID)

	return cmd
}
