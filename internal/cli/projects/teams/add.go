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

const addTemplate = "Team added to the project.\n"

type AddOpts struct {
	cli.GlobalOpts
	cli.OutputOpts
	store  store.ProjectTeamAdder
	teamID string
	roles  []string
}

func (opts *AddOpts) initStore(ctx context.Context) func() error {
	return func() error {
		var err error
		opts.store, err = store.New(store.AuthenticatedPreset(config.Default()), store.WithContext(ctx))
		return err
	}
}

func (opts *AddOpts) Run() error {
	r, err := opts.store.AddTeamsToProject(opts.ConfigProjectID(), opts.newProjectTeam())
	if err != nil {
		return err
	}

	return opts.Print(r)
}

func (opts *AddOpts) newProjectTeam() []atlasv2.TeamRole {
	return []atlasv2.TeamRole{
		{
			TeamId:    &opts.teamID,
			RoleNames: &opts.roles,
		},
	}
}

// atlas project(s) team(s) add teamId --projectId projectId --role role.
func AddBuilder() *cobra.Command {
	opts := &AddOpts{}
	cmd := &cobra.Command{
		Use:   "add <teamId>",
		Args:  require.ExactArgs(1),
		Short: "Add the specified team to your project.",
		Long: `All members of the team share the same project access.

` + fmt.Sprintf(usage.RequiredRole, "Project Owner"),
		Annotations: map[string]string{
			"teamIdDesc": "Unique 24-digit string that identifies the team.",
			"output":     addTemplate,
		},
		Example: `  # Add the team with the ID 5dd58c647a3e5a6c5bce46c7 to the project with the ID 5e2211c17a3e5a48f5497de3 with GROUP_READ_ONLY project access:
  atlas projects teams add 5dd58c647a3e5a6c5bce46c7 --projectId 5e2211c17a3e5a48f5497de3 --role GROUP_READ_ONLY`,
		PreRunE: func(cmd *cobra.Command, args []string) error {
			opts.teamID = args[0]
			return opts.PreRunE(
				opts.ValidateProjectID,
				opts.initStore(cmd.Context()),
				opts.InitOutput(cmd.OutOrStdout(), addTemplate),
			)
		},
		RunE: func(_ *cobra.Command, _ []string) error {
			return opts.Run()
		},
	}

	cmd.Flags().StringSliceVar(&opts.roles, flag.Role, []string{}, usage.TeamRole)

	cmd.Flags().StringVar(&opts.ProjectID, flag.ProjectID, "", usage.ProjectID)
	cmd.Flags().StringVarP(&opts.Output, flag.Output, flag.OutputShort, "", usage.FormatOut)
	_ = cmd.RegisterFlagCompletionFunc(flag.Output, opts.AutoCompleteOutputFlag())

	_ = cmd.MarkFlagRequired(flag.Role)

	return cmd
}
