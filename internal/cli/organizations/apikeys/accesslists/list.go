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

	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli/require"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/config"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/store"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/usage"
	"github.com/spf13/cobra"
	"go.mongodb.org/atlas-sdk/v20250312005/admin"
)

const listTemplate = `IP ADDRESS	CIDR BLOCK	CREATED AT{{range valueOrEmptySlice .Results}}
{{.IpAddress}}	{{.CidrBlock}}	{{.Created}}{{end}}
`

//go:generate go tool go.uber.org/mock/mockgen -typed -destination=list_mock_test.go -package=accesslists . OrganizationAPIKeyAccessListLister

type OrganizationAPIKeyAccessListLister interface {
	OrganizationAPIKeyAccessLists(*admin.ListApiKeyAccessListsEntriesApiParams) (*admin.PaginatedApiUserAccessListResponse, error)
}

type ListOpts struct {
	cli.OrgOpts
	cli.OutputOpts
	cli.ListOpts
	id    string
	store OrganizationAPIKeyAccessListLister
}

func (opts *ListOpts) initStore(ctx context.Context) func() error {
	return func() error {
		var err error
		opts.store, err = store.New(store.AuthenticatedPreset(config.Default()), store.WithContext(ctx))
		return err
	}
}

func (opts *ListOpts) Run() error {
	listOpts := opts.NewAtlasListOptions()
	params := &admin.ListApiKeyAccessListsEntriesApiParams{
		OrgId:        opts.ConfigOrgID(),
		ApiUserId:    opts.id,
		PageNum:      &listOpts.PageNum,
		ItemsPerPage: &listOpts.ItemsPerPage,
		IncludeCount: &listOpts.IncludeCount,
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
		Long: `To view possible values for the apiKeyID argument, run atlas organizations apiKeys list.

` + fmt.Sprintf(usage.RequiredRole, "Organization Member"),
		Annotations: map[string]string{
			"apiKeyIDDesc": "Unique 24-digit string that identifies your API key.",
		},
		Example: `  # Return a JSON-formatted list of access list entries for the API key with the ID 5f24084d8dbffa3ad3f21234 in the organization with the ID 5a1b39eec902201990f12345:
  atlas organizations apiKeys accessLists list --apiKey 5f24084d8dbffa3ad3f21234 --orgId 5a1b39eec902201990f12345 --output json`,
		PreRunE: func(cmd *cobra.Command, _ []string) error {
			return opts.PreRunE(
				opts.ValidateOrgID,
				opts.initStore(cmd.Context()),
				opts.InitOutput(cmd.OutOrStdout(), listTemplate),
			)
		},
		RunE: func(_ *cobra.Command, args []string) error {
			opts.id = args[0]
			return opts.Run()
		},
	}

	opts.AddListOptsFlags(cmd)

	opts.AddOrgOptFlags(cmd)
	opts.AddOutputOptFlags(cmd)

	return cmd
}
