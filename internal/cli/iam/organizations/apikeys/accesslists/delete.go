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

package accesslists

import (
	"context"
	"fmt"

	"github.com/mongodb/mongodb-atlas-cli/internal/cli"
	"github.com/mongodb/mongodb-atlas-cli/internal/cli/require"
	"github.com/mongodb/mongodb-atlas-cli/internal/config"
	"github.com/mongodb/mongodb-atlas-cli/internal/flag"
	"github.com/mongodb/mongodb-atlas-cli/internal/store"
	"github.com/mongodb/mongodb-atlas-cli/internal/usage"
	"github.com/spf13/cobra"
)

type DeleteOpts struct {
	*cli.DeleteOpts
	cli.GlobalOpts
	apiKey string
	store  store.OrganizationAPIKeyAccessListDeleter
}

func (opts *DeleteOpts) initStore(ctx context.Context) func() error {
	return func() error {
		var err error
		opts.store, err = store.New(store.AuthenticatedPreset(config.Default()), store.WithContext(ctx))
		return err
	}
}

func (opts *DeleteOpts) Run() error {
	return opts.Delete(opts.store.DeleteOrganizationAPIKeyAccessList, opts.ConfigOrgID(), opts.apiKey)
}

// DeleteBuilder mongocli iam organizations|orgs apiKey(s)|apikey(s) accesslist delete <IP> [--orgId orgId] [--apiKey apiKey] --force.
func DeleteBuilder() *cobra.Command {
	opts := &DeleteOpts{
		DeleteOpts: cli.NewDeleteOpts("Access list entry '%s' deleted\n", "Access list entry not deleted"),
	}

	cmd := &cobra.Command{
		Use:     "delete <entry>",
		Aliases: []string{"rm"},
		Short:   "Remove the specified IP access list entry from your API Key.",
		Long:    fmt.Sprintf(usage.RequiredRole, "Read Write"),
		Args:    require.ExactArgs(1),
		Annotations: map[string]string{
			"entryDesc": "IP or CIDR address that you want to remove from the access list. If the entry includes a subnet mask, such as 192.0.2.0/24, use the URL-encoded value %2F for the forward slash /.",
			"output":    opts.SuccessMessage(),
		},
		Example: fmt.Sprintf(`  # Remove the IP address 192.0.2.0 from the access list for the API key with the ID 5f24084d8dbffa3ad3f21234 in the organization with the ID 5a1b39eec902201990f12345:
  %s organizations apiKeys accessLists delete 192.0.2.0 --apiKey 5f24084d8dbffa3ad3f21234 --orgId 5a1b39eec902201990f12345`, cli.ExampleAtlasEntryPoint()),
		PreRunE: func(cmd *cobra.Command, args []string) error {
			if err := opts.PreRunE(opts.ValidateOrgID, opts.initStore(cmd.Context())); err != nil {
				return err
			}
			opts.Entry = args[0]
			return opts.Prompt()
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return opts.Run()
		},
	}
	cmd.Flags().BoolVar(&opts.Confirm, flag.Force, false, usage.Force)
	cmd.Flags().StringVar(&opts.apiKey, flag.APIKey, "", usage.APIKey)

	cmd.Flags().StringVar(&opts.OrgID, flag.OrgID, "", usage.OrgID)

	return cmd
}
