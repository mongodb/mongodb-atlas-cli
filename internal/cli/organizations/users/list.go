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

	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli/require"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/config"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/flag"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/store"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/usage"
	"github.com/spf13/cobra"
)

const listTemplate = `ID	FIRST NAME	LAST NAME	USERNAME{{range valueOrEmptySlice .Results}}
{{.Id}}	{{.FirstName}}	{{.LastName}}	{{.Username}}{{end}}
`

type ListOpts struct {
	cli.GlobalOpts
	cli.ListOpts
	cli.OutputOpts
	store store.UserLister
}

func (opts *ListOpts) initStore(ctx context.Context) func() error {
	return func() error {
		var err error
		opts.store, err = store.New(store.AuthenticatedPreset(config.Default()), store.WithContext(ctx))
		return err
	}
}

func (opts *ListOpts) Run() error {
	listOptions := opts.NewListOptions()
	r, err := opts.store.OrganizationUsers(opts.ConfigOrgID(), listOptions)

	if err != nil {
		return err
	}

	return opts.Print(r)
}

// ListBuilder atlas organizations(s) users list --orgId orgId.
func ListBuilder() *cobra.Command {
	opts := new(ListOpts)

	cmd := &cobra.Command{
		Use:     "list",
		Aliases: []string{"ls"},
		Annotations: map[string]string{
			"output": listTemplate,
		},
		Short: "Return all users for an organization.",
		Long:  fmt.Sprintf(usage.RequiredRole, "Organization Member"),
		Args:  require.NoArgs,
		Example: `  # Return a JSON-formatted list of all users for the organization with the ID 5e2211c17a3e5a48f5497de3:
  atlas organizations users list --orgId 5e2211c17a3e5a48f5497de3 --output json`,
		PreRunE: func(cmd *cobra.Command, _ []string) error {
			opts.OutWriter = cmd.OutOrStdout()
			return opts.PreRunE(
				opts.ValidateOrgID,
				opts.InitOutput(cmd.OutOrStdout(), listTemplate),
				opts.initStore(cmd.Context()),
			)
		},
		RunE: func(_ *cobra.Command, _ []string) error {
			return opts.Run()
		},
	}

	cmd.Flags().IntVar(&opts.PageNum, flag.Page, cli.DefaultPage, usage.Page)
	cmd.Flags().IntVar(&opts.ItemsPerPage, flag.Limit, cli.DefaultPageLimit, usage.Limit)
	cmd.Flags().BoolVar(&opts.OmitCount, flag.OmitCount, false, usage.OmitCount)

	cmd.Flags().StringVarP(&opts.Output, flag.Output, flag.OutputShort, "", usage.FormatOut)
	_ = cmd.RegisterFlagCompletionFunc(flag.Output, opts.AutoCompleteOutputFlag())

	cmd.Flags().StringVar(&opts.OrgID, flag.OrgID, "", usage.OrgID)

	return cmd
}

func Builder() *cobra.Command {
	const use = "users"
	cmd := &cobra.Command{
		Use:     use,
		Short:   "Manage your Atlas users.",
		Aliases: cli.GenerateAliases(use),
	}

	cmd.AddCommand(
		ListBuilder(),
	)

	return cmd
}
