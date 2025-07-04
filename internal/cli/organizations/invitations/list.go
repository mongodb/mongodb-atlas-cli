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

	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli/require"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/config"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/flag"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/store"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/usage"
	"github.com/spf13/cobra"
	atlasv2 "go.mongodb.org/atlas-sdk/v20250312005/admin"
)

const listTemplate = `ID	USERNAME	CREATED AT	EXPIRES AT{{range valueOrEmptySlice .}}
{{.Id}}	{{.Username}}	{{.CreatedAt}}	{{.ExpiresAt}}{{end}}
`

//go:generate go tool go.uber.org/mock/mockgen -typed -destination=list_mock_test.go -package=invitations . OrganizationInvitationLister

type OrganizationInvitationLister interface {
	OrganizationInvitations(*atlasv2.ListOrganizationInvitationsApiParams) ([]atlasv2.OrganizationInvitation, error)
}

type ListOpts struct {
	cli.OrgOpts
	cli.OutputOpts
	store    OrganizationInvitationLister
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
		Example: `  # Return a JSON-formatted list of pending invitations to the organization with the ID 5f71e5255afec75a3d0f96dc:
  atlas organizations invitations list --orgId 5f71e5255afec75a3d0f96dc --output json`,
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

	cmd.Flags().StringVar(&opts.username, flag.Email, "", usage.Email)
	opts.AddOrgOptFlags(cmd)

	opts.AddOutputOptFlags(cmd)

	return cmd
}
