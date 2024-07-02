// Copyright 2024 MongoDB Inc
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

var renameTemplate = "Team '{{.Name}}' updated.\n"

type renameOpts struct {
	cli.GlobalOpts
	cli.OutputOpts
	name   string
	teamID string
	store  store.TeamRenamer
}

func (opts *renameOpts) initStore(ctx context.Context) func() error {
	return func() error {
		var err error
		opts.store, err = store.New(store.AuthenticatedPreset(config.Default()), store.WithContext(ctx))
		return err
	}
}

func (opts *renameOpts) Run() error {
	r, err := opts.store.RenameTeam(opts.ConfigOrgID(), opts.teamID, opts.newTeam())
	if err != nil {
		return err
	}
	return opts.Print(r)
}

func (opts *renameOpts) newTeam() *atlasv2.TeamUpdate {
	return &atlasv2.TeamUpdate{
		Name: opts.name,
	}
}

// RenameBuilder atlas team(s) update <newName> [--orgId orgId].
func RenameBuilder() *cobra.Command {
	opts := &renameOpts{}
	cmd := &cobra.Command{
		Use:     "rename <newName>",
		Aliases: []string{"update"},
		Short:   "Rename a team in your organization.",
		Long:    fmt.Sprintf(usage.RequiredRole, "Organization Owner"),
		Args:    require.ExactArgs(1),
		Annotations: map[string]string{
			"newNameDesc": "New label that identifies the team.",
			"output":      renameTemplate,
		},
		Example: `  # Rename a team in the organization with ID 5e2211c17a3e5a48f5497de3:
  atlas teams rename newName --teamId 5e1234c17a3e5a48f5497de3 --orgId 5e1234c17a3e5a48f5497de3 --output json`,
		PreRunE: func(cmd *cobra.Command, _ []string) error {
			return opts.PreRunE(
				opts.ValidateOrgID,
				opts.initStore(cmd.Context()),
				opts.InitOutput(cmd.OutOrStdout(), renameTemplate),
			)
		},
		RunE: func(_ *cobra.Command, args []string) error {
			opts.name = args[0]
			return opts.Run()
		},
	}

	cmd.Flags().StringVar(&opts.teamID, flag.TeamID, "", usage.TeamID)

	cmd.Flags().StringVar(&opts.OrgID, flag.OrgID, "", usage.OrgID)
	cmd.Flags().StringVarP(&opts.Output, flag.Output, flag.OutputShort, "", usage.FormatOut)
	_ = cmd.RegisterFlagCompletionFunc(flag.Output, opts.AutoCompleteOutputFlag())

	_ = cmd.MarkFlagRequired(flag.TeamID)

	return cmd
}
