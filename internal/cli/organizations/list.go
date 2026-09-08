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

package organizations

import (
	"context"
	"fmt"

	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/api"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/api/orgsapi"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli/require"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/flag"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/usage"
	"github.com/spf13/cobra"
)

const listTemplate = `ID	NAME{{range valueOrEmptySlice .Results}}
{{.Id}}	{{.Name}}{{end}}
`

type ListOpts struct {
	cli.ProjectOpts
	cli.ListOpts
	cli.OutputOpts
	executor           api.CommandExecutor
	name               string
	includeDeletedOrgs bool
	includeGlobal      bool
}

func (opts *ListOpts) initStore(_ context.Context) func() error {
	return func() error {
		var err error
		opts.executor, err = api.NewDefaultExecutor(api.NewFormatter())
		return err
	}
}

func (opts *ListOpts) Run(ctx context.Context) error {
	r, err := orgsapi.ListOrgs(ctx, opts.executor, opts.newOrganizationListOptions())
	if err != nil {
		return err
	}
	return opts.Print(r)
}

func (opts *ListOpts) newOrganizationListOptions() orgsapi.ListOrgsOptions {
	options := orgsapi.ListOrgsOptions{
		Name:          opts.name,
		IncludeGlobal: opts.includeGlobal,
	}
	if listOpt := opts.NewAtlasListOptions(); listOpt != nil {
		options.PageNum = listOpt.PageNum
		options.ItemsPerPage = listOpt.ItemsPerPage
		options.IncludeCount = listOpt.IncludeCount
	}
	return options
}

// atlas organizations(s) list --name --includeDeletedOrgs.
func ListBuilder() *cobra.Command {
	opts := new(ListOpts)
	cmd := &cobra.Command{
		Use:     "list",
		Aliases: []string{"ls"},
		Annotations: map[string]string{
			"output": listTemplate,
		},
		Short: "Return all organizations.",
		Long:  fmt.Sprintf(usage.RequiredRole, "Organization Member"),
		Args:  require.NoArgs,
		Example: `  # Return a JSON-formatted list of all organizations:
  atlas organizations list --output json
  
  # Return a JSON-formatted list that includes the organizations named org1 and Org1, but doesn't return org123:
  atlas organizations list --name org1 --output json`,
		PreRunE: func(cmd *cobra.Command, _ []string) error {
			return opts.PreRunE(
				opts.initStore(cmd.Context()),
				opts.InitOutput(cmd.OutOrStdout(), listTemplate),
			)
		},
		RunE: func(cmd *cobra.Command, _ []string) error {
			return opts.Run(cmd.Context())
		},
	}

	opts.AddListOptsFlags(cmd)

	cmd.Flags().StringVar(&opts.name, flag.Name, "", usage.OrgNameFilter)
	cmd.Flags().BoolVar(&opts.includeDeletedOrgs, flag.IncludeDeleted, false, usage.OrgIncludeDeleted)
	// Hidden: not part of the public API and a no-op for customers. See CLOUDP-432111.
	cmd.Flags().BoolVar(&opts.includeGlobal, flag.IncludeGlobal, true, usage.OrgIncludeGlobal)
	_ = cmd.Flags().MarkHidden(flag.IncludeGlobal)

	opts.AddOutputOptFlags(cmd)

	return cmd
}
