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
	atlasv2 "go.mongodb.org/atlas-sdk/v20240530002/admin"
)

const updateTemplate = "Team's roles updated.\n"

type UpdateOpts struct {
	cli.GlobalOpts
	cli.OutputOpts
	store  store.TeamRolesUpdater
	teamID string
	roles  []string
}

func (opts *UpdateOpts) initStore(ctx context.Context) func() error {
	return func() error {
		var err error
		opts.store, err = store.New(store.AuthenticatedPreset(config.Default()), store.WithContext(ctx))
		return err
	}
}

func (opts *UpdateOpts) Run() error {
	r, err := opts.store.UpdateProjectTeamRoles(opts.ConfigProjectID(), opts.teamID, opts.newTeamUpdateRoles())
	if err != nil {
		return err
	}

	return opts.Print(r)
}

func (opts *UpdateOpts) newTeamUpdateRoles() *atlasv2.TeamRole {
	return &atlasv2.TeamRole{
		RoleNames: &opts.roles,
	}
}

// atlas team(s) user(s) updates teamId --projectId projectId --role role.
func UpdateBuilder() *cobra.Command {
	opts := &UpdateOpts{}
	cmd := &cobra.Command{
		Use:     "update <teamId>",
		Aliases: []string{"updates"},
		Args:    require.ExactArgs(1),
		Short:   "Modify the roles for the specified team for your project.",
		Long:    fmt.Sprintf(usage.RequiredRole, "Project User Admin"),
		Annotations: map[string]string{
			"teamIdDesc": "Unique 24-digit string that identifies the team.",
			"output":     updateTemplate,
		},
		Example: `  # Modify the roles for the team with the ID 5dd56c847a3e5a1f363d424d to grant GROUP_READ_ONLY access to the project with the ID 5f71e5255afec75a3d0f96dc:
  atlas projects teams update 5dd56c847a3e5a1f363d424d --projectId 5f71e5255afec75a3d0f96dc --role GROUP_READ_ONLY --output json`,
		PreRunE: func(cmd *cobra.Command, args []string) error {
			opts.teamID = args[0]
			return opts.PreRunE(
				opts.ValidateProjectID,
				opts.initStore(cmd.Context()),
				opts.InitOutput(cmd.OutOrStdout(), updateTemplate),
			)
		},
		RunE: func(_ *cobra.Command, _ []string) error {
			return opts.Run()
		},
	}

	cmd.Flags().StringSliceVar(&opts.roles, flag.Role, []string{}, usage.TeamRole+usage.UpdateWarning)

	cmd.Flags().StringVar(&opts.ProjectID, flag.ProjectID, "", usage.ProjectID)
	cmd.Flags().StringVarP(&opts.Output, flag.Output, flag.OutputShort, "", usage.FormatOut)
	_ = cmd.RegisterFlagCompletionFunc(flag.Output, opts.AutoCompleteOutputFlag())

	_ = cmd.MarkFlagRequired(flag.Role)

	return cmd
}
