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

package invitations

import (
	"context"
	"fmt"

	"github.com/andreaangiolillo/mongocli-test/internal/cli"
	"github.com/andreaangiolillo/mongocli-test/internal/cli/require"
	"github.com/andreaangiolillo/mongocli-test/internal/config"
	"github.com/andreaangiolillo/mongocli-test/internal/flag"
	store "github.com/andreaangiolillo/mongocli-test/internal/store/atlas"
	"github.com/andreaangiolillo/mongocli-test/internal/usage"
	"github.com/spf13/cobra"
	atlasv2 "go.mongodb.org/atlas-sdk/v20231115002/admin"
)

const listTemplate = `ID	USERNAME	CREATED AT	EXPIRES AT{{range .}}
{{.Id}}	{{.Username}}	{{.CreatedAt}}	{{.ExpiresAt}}{{end}}
`

type ListOpts struct {
	cli.GlobalOpts
	cli.OutputOpts
	store    store.OrganizationInvitationLister
	username string
}

func (opts *ListOpts) initStore(ctx context.Context) func() error {
	return func() error {
		var err error
		opts.store, err = store.New(store.AuthenticatedPreset(config.Default()), store.WithContext(ctx))
		return err
	}
}

func (opts *ListOpts) Run() error {
	r, err := opts.store.OrganizationInvitations(opts.newInvitationOptions())
	if err != nil {
		return err
	}
	return opts.Print(r)
}

func (opts *ListOpts) newInvitationOptions() *atlasv2.ListOrganizationInvitationsApiParams {
	return &atlasv2.ListOrganizationInvitationsApiParams{
		OrgId:    opts.ConfigOrgID(),
		Username: &opts.username,
	}
}

// atlas organizations(s) invitations list [--email email]  [--orgId orgId].
func ListBuilder() *cobra.Command {
	opts := new(ListOpts)
	cmd := &cobra.Command{
		Use:     "list",
		Aliases: []string{"ls"},
		Short:   "Return all pending invitations to your organization.",
		Long:    fmt.Sprintf(usage.RequiredRole, "Organization User Admin"),
		Args:    require.NoArgs,
		Example: fmt.Sprintf(`  # Return a JSON-formatted list of pending invitations to the organization with the ID 5f71e5255afec75a3d0f96dc:
  %s organizations invitations list --orgId 5f71e5255afec75a3d0f96dc --output json`, cli.ExampleAtlasEntryPoint()),
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return opts.PreRunE(
				opts.ValidateOrgID,
				opts.initStore(cmd.Context()),
				opts.InitOutput(cmd.OutOrStdout(), listTemplate),
			)
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return opts.Run()
		},
	}

	cmd.Flags().StringVar(&opts.username, flag.Email, "", usage.Email)
	cmd.Flags().StringVar(&opts.OrgID, flag.OrgID, "", usage.OrgID)

	cmd.Flags().StringVarP(&opts.Output, flag.Output, flag.OutputShort, "", usage.FormatOut)
	_ = cmd.RegisterFlagCompletionFunc(flag.Output, opts.AutoCompleteOutputFlag())

	return cmd
}
