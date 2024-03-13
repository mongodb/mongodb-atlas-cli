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

	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli/require"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/config"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/flag"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/store"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/usage"
	"github.com/spf13/cobra"
	atlasv2 "go.mongodb.org/atlas-sdk/v20231115007/admin"
)

const listTemplate = `ID	NAME{{range valueOrEmptySlice .Results}}
{{.Id}}	{{.Name}}{{end}}
`

type ListOpts struct {
	cli.GlobalOpts
	cli.ListOpts
	cli.OutputOpts
	store              store.OrganizationLister
	name               string
	includeDeletedOrgs bool
}

func (opts *ListOpts) initStore(ctx context.Context) func() error {
	return func() error {
		var err error
		opts.store, err = store.New(store.AuthenticatedPreset(config.Default()), store.WithContext(ctx))
		return err
	}
}

func (opts *ListOpts) Run() error {
	r, err := opts.store.Organizations(opts.newOrganizationListOptions())
	if err != nil {
		return err
	}
	return opts.Print(r)
}

func (opts *ListOpts) newOrganizationListOptions() *atlasv2.ListOrganizationsApiParams {
	params := &atlasv2.ListOrganizationsApiParams{
		Name: &opts.name,
	}
	if listOpt := opts.NewListOptions(); listOpt != nil {
		params.PageNum = &listOpt.PageNum
		params.ItemsPerPage = &listOpt.ItemsPerPage
		params.IncludeCount = &listOpt.IncludeCount
	}
	return params
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
		RunE: func(_ *cobra.Command, _ []string) error {
			return opts.Run()
		},
	}

	cmd.Flags().IntVar(&opts.PageNum, flag.Page, cli.DefaultPage, usage.Page)
	cmd.Flags().IntVar(&opts.ItemsPerPage, flag.Limit, cli.DefaultPageLimit, usage.Limit)
	cmd.Flags().BoolVar(&opts.OmitCount, flag.OmitCount, false, usage.OmitCount)

	cmd.Flags().StringVar(&opts.name, flag.Name, "", usage.OrgNameFilter)
	cmd.Flags().BoolVar(&opts.includeDeletedOrgs, flag.IncludeDeleted, false, usage.OrgIncludeDeleted)

	cmd.Flags().StringVarP(&opts.Output, flag.Output, flag.OutputShort, "", usage.FormatOut)
	_ = cmd.RegisterFlagCompletionFunc(flag.Output, opts.AutoCompleteOutputFlag())

	return cmd
}
