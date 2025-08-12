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
	atlasv2 "go.mongodb.org/atlas-sdk/v20250312006/admin"
)

const describeTemplate = `ID	NAME
{{.Id}}	{{.Name}}
`

//go:generate go tool go.uber.org/mock/mockgen -typed -destination=describe_mock_test.go -package=teams . TeamDescriber

type TeamDescriber interface {
	TeamByID(string, string) (*atlasv2.TeamResponse, error)
	TeamByName(string, string) (*atlasv2.TeamResponse, error)
}

type DescribeOpts struct {
	cli.OrgOpts
	cli.OutputOpts
	store TeamDescriber
	name  string
	id    string
}

func (opts *DescribeOpts) initStore(ctx context.Context) func() error {
	return func() error {
		var err error
		opts.store, err = store.New(store.AuthenticatedPreset(config.Default()), store.WithContext(ctx))
		return err
	}
}

func (opts *DescribeOpts) Run() error {
	var r any
	var err error

	if opts.name != "" {
		r, err = opts.store.TeamByName(opts.ConfigOrgID(), opts.name)
	}

	if opts.id != "" {
		r, err = opts.store.TeamByID(opts.ConfigOrgID(), opts.id)
	}

	if err != nil {
		return err
	}

	return opts.Print(r)
}

// atlas team(s) describe --id id --name name --orgId orgId.
func DescribeBuilder() *cobra.Command {
	opts := &DescribeOpts{}
	cmd := &cobra.Command{
		Use:         "describe",
		Aliases:     []string{"get"},
		Annotations: map[string]string{"output": describeTemplate},
		Example: `  # Return the JSON-formatted details for the the team with the ID 5e44445ef10fab20b49c0f31 in the organization with ID 5e2211c17a3e5a48f5497de3:
  atlas teams describe --id 5e44445ef10fab20b49c0f31 --projectId 5e1234c17a3e5a48f5497de3 --output json
  
  # Return the JSON-formatted details for the the team with the name myTeam in the organization with ID 5e2211c17a3e5a48f5497de3:
  atlas teams describe --name myTeam --projectId 5e1234c17a3e5a48f5497de3 --output json`,
		Short: "Return the details for the specified team for your organization.",
		Long: `You can return the details for a team using the team's ID or the team's name. You must specify either the id option or the name option.

` + fmt.Sprintf(usage.RequiredRole, "Organization Member"),
		Args: require.NoArgs,
		PreRunE: func(cmd *cobra.Command, _ []string) error {
			return opts.PreRunE(
				opts.ValidateOrgID,
				opts.initStore(cmd.Context()),
				opts.InitOutput(cmd.OutOrStdout(), describeTemplate),
			)
		},
		RunE: func(_ *cobra.Command, _ []string) error {
			return opts.Run()
		},
	}

	cmd.Flags().StringVar(&opts.name, flag.Name, "", usage.TeamName)
	cmd.Flags().StringVar(&opts.id, flag.ID, "", usage.TeamID)

	cmd.MarkFlagsOneRequired(flag.Name, flag.ID)
	cmd.MarkFlagsMutuallyExclusive(flag.Name, flag.ID)

	opts.AddOrgOptFlags(cmd)
	opts.AddOutputOptFlags(cmd)

	return cmd
}
