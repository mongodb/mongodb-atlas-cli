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
	atlas "go.mongodb.org/atlas/mongodbatlas"
)

const addTemplate = "Team added to the project.\n"

type AddOpts struct {
	cli.GlobalOpts
	cli.OutputOpts
	store  store.ProjectTeamAdder
	teamID string
	roles  []string
}

func (opts *AddOpts) init() error {
	var err error
	opts.store, err = store.New(config.Default())
	return err
}

func (opts *AddOpts) Run() error {
	r, err := opts.store.AddTeamsToProject(opts.ConfigProjectID(), opts.newProjectTeam())
	if err != nil {
		return err
	}

	return opts.Print(r)
}

func (opts *AddOpts) newProjectTeam() []*atlas.ProjectTeam {
	return []*atlas.ProjectTeam{
		{
			TeamID:    opts.teamID,
			RoleNames: opts.roles,
		},
	}
}

// mongocli iam project(s) team(s) add teamId --projectId projectId --role role
func AddBuilder() *cobra.Command {
	opts := &AddOpts{}
	cmd := &cobra.Command{
		Use:   "add <teamId>",
		Args:  require.ExactArgs(1),
		Short: addTeam,
		PreRunE: func(cmd *cobra.Command, args []string) error {
			opts.teamID = args[0]
			return opts.PreRunE(
				opts.ValidateProjectID,
				opts.init,
				opts.InitOutput(cmd.OutOrStdout(), addTemplate),
			)
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return opts.Run()
		},
	}

	cmd.Flags().StringSliceVar(&opts.roles, flag.Role, []string{}, usage.TeamRole)

	cmd.Flags().StringVar(&opts.ProjectID, flag.ProjectID, "", usage.ProjectID)
	cmd.Flags().StringVarP(&opts.Output, flag.Output, flag.OutputShort, "", usage.FormatOut)

	_ = cmd.MarkFlagRequired(flag.Role)

	return cmd
}
