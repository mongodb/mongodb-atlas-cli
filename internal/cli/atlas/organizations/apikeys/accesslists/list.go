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

package accesslists

import (
	"context"
	"fmt"

	"github.com/mongodb/mongodb-atlas-cli/internal/cli"
	"github.com/mongodb/mongodb-atlas-cli/internal/cli/require"
	"github.com/mongodb/mongodb-atlas-cli/internal/config"
	"github.com/mongodb/mongodb-atlas-cli/internal/flag"
	"github.com/mongodb/mongodb-atlas-cli/internal/pointer"
	store "github.com/mongodb/mongodb-atlas-cli/internal/store/atlas"
	"github.com/mongodb/mongodb-atlas-cli/internal/usage"
	"github.com/spf13/cobra"
	"go.mongodb.org/atlas-sdk/v20231115004/admin"
)

const listTemplate = `IP ADDRESS	CIDR BLOCK	CREATED AT{{range .Results}}
{{.IpAddress}}	{{.CidrBlock}}	{{.Created}}{{end}}
`

type ListOpts struct {
	cli.GlobalOpts
	cli.OutputOpts
	cli.ListOpts
	id    string
	store store.OrganizationAPIKeyAccessListLister
}

func (opts *ListOpts) initStore(ctx context.Context) func() error {
	return func() error {
		var err error
		opts.store, err = store.New(store.AuthenticatedPreset(config.Default()), store.WithContext(ctx))
		return err
	}
}

func (opts *ListOpts) Run() error {
	listOpts := opts.NewListOptions()
	params := &admin.ListApiKeyAccessListsEntriesApiParams{
		OrgId:        opts.ConfigOrgID(),
		ApiUserId:    opts.id,
		PageNum:      pointer.Get(listOpts.PageNum),
		ItemsPerPage: pointer.Get(listOpts.ItemsPerPage),
	}
	result, err := opts.store.OrganizationAPIKeyAccessLists(params)
	if err != nil {
		return err
	}
	return opts.Print(result)
}

// atlas organizations|orgs apiKey(s)|apikey(s) accessList list|ls <apiKeyID> [--orgId orgId].
func ListBuilder() *cobra.Command {
	opts := new(ListOpts)
	cmd := &cobra.Command{
		Use:     "list <apiKeyID>",
		Aliases: []string{"ls"},
		Args:    require.ExactArgs(1),
		Short:   "Return all IP access list entries for your API Key.",
		Long: fmt.Sprintf(`To view possible values for the apiKeyID argument, run %s organizations apiKeys list.

`+fmt.Sprintf(usage.RequiredRole, "Organization Member"), cli.ExampleAtlasEntryPoint()),
		Annotations: map[string]string{
			"apiKeyIDDesc": "Unique 24-digit string that identifies your API key.",
		},
		Example: fmt.Sprintf(`  # Return a JSON-formatted list of access list entries for the API key with the ID 5f24084d8dbffa3ad3f21234 in the organization with the ID 5a1b39eec902201990f12345:
  %s organizations apiKeys accessLists list --apiKey 5f24084d8dbffa3ad3f21234 --orgId 5a1b39eec902201990f12345 --output json`, cli.ExampleAtlasEntryPoint()),
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return opts.PreRunE(
				opts.ValidateOrgID,
				opts.initStore(cmd.Context()),
				opts.InitOutput(cmd.OutOrStdout(), listTemplate),
			)
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			opts.id = args[0]
			return opts.Run()
		},
	}

	cmd.Flags().IntVar(&opts.PageNum, flag.Page, cli.DefaultPage, usage.Page)
	cmd.Flags().IntVar(&opts.ItemsPerPage, flag.Limit, cli.DefaultPageLimit, usage.Limit)

	cmd.Flags().StringVar(&opts.OrgID, flag.OrgID, "", usage.OrgID)
	cmd.Flags().StringVarP(&opts.Output, flag.Output, flag.OutputShort, "", usage.FormatOut)
	_ = cmd.RegisterFlagCompletionFunc(flag.Output, opts.AutoCompleteOutputFlag())

	return cmd
}
