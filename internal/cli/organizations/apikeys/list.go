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

package apikeys

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
	atlasv2 "go.mongodb.org/atlas-sdk/v20250312005/admin"
)

const listTemplate = `ID	DESCRIPTION	PUBLIC KEY	PRIVATE KEY{{range valueOrEmptySlice .Results}}
{{.Id}}	{{.Desc}}	{{.PublicKey}}	{{.PrivateKey}}{{end}}
`

//go:generate go tool go.uber.org/mock/mockgen -typed -destination=list_mock_test.go -package=apikeys . OrganizationAPIKeyLister

type OrganizationAPIKeyLister interface {
	OrganizationAPIKeys(string, *store.ListOptions) (*atlasv2.PaginatedApiApiUser, error)
}

type ListOpts struct {
	cli.OrgOpts
	cli.OutputOpts
	cli.ListOpts
	CompactResponse bool
	store           OrganizationAPIKeyLister
}

func (opts *ListOpts) initStore(ctx context.Context) func() error {
	return func() error {
		var err error
		opts.store, err = store.New(store.AuthenticatedPreset(config.Default()), store.WithContext(ctx))
		return err
	}
}

func (opts *ListOpts) Run() error {
	r, err := opts.store.OrganizationAPIKeys(opts.ConfigOrgID(), opts.NewAtlasListOptions())

	if err != nil {
		return err
	}

	if opts.CompactResponse {
		return opts.PrintForCompactResultsResponse(r)
	}

	return opts.Print(r)
}

// atlas organizations|orgs apiKey(s)|apikey(s) list|ls [--orgId orgId].
func ListBuilder() *cobra.Command {
	opts := new(ListOpts)
	cmd := &cobra.Command{
		Use:     "list",
		Aliases: []string{"ls"},
		Short:   "Return all API keys for your organization.",
		Long:    fmt.Sprintf(usage.RequiredRole, "Organization Member"),
		Args:    require.NoArgs,
		Example: `  # Return a JSON-formatted list of organization API keys for the organization with the ID 5a1b39eec902201990f12345:
  atlas organizations apiKeys list --orgId 5a1b39eec902201990f12345 --output json`,
		PreRunE: func(cmd *cobra.Command, _ []string) error {
			return opts.PreRunE(
				opts.ValidateOrgID,
				opts.initStore(cmd.Context()),
				opts.InitOutput(cmd.OutOrStdout(), listTemplate),
			)
		},
		RunE: func(_ *cobra.Command, _ []string) error {
			return opts.Run()
		},
	}

	opts.AddListOptsFlags(cmd)

	opts.AddOrgOptFlags(cmd)
	opts.AddOutputOptFlags(cmd)
	cmd.Flags().BoolVarP(&opts.CompactResponse, flag.CompactResponse, flag.CompactResponseShort, false, usage.CompactResponse)

	return cmd
}
