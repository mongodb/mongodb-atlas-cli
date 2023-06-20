// Copyright 2021 MongoDB Inc
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

	"github.com/mongodb/mongodb-atlas-cli/internal/cli"
	"github.com/mongodb/mongodb-atlas-cli/internal/cli/require"
	"github.com/mongodb/mongodb-atlas-cli/internal/config"
	"github.com/mongodb/mongodb-atlas-cli/internal/flag"
	"github.com/mongodb/mongodb-atlas-cli/internal/store"
	"github.com/mongodb/mongodb-atlas-cli/internal/usage"
	"github.com/spf13/cobra"
	atlas "go.mongodb.org/atlas/mongodbatlas"
)

const createTemplate = "User '{{.Username}}' invited.\n"

type InviteOpts struct {
	cli.OutputOpts
	cli.GlobalOpts
	username string
	roles    []string
	teamIds  []string
	store    store.OrganizationInviter
}

func (opts *InviteOpts) initStore(ctx context.Context) func() error {
	return func() error {
		var err error
		opts.store, err = store.New(store.AuthenticatedPreset(config.Default()), store.WithContext(ctx))
		return err
	}
}

func (opts *InviteOpts) Run() error {
	r, err := opts.store.InviteUser(opts.ConfigOrgID(), opts.newInvitation())

	if err != nil {
		return err
	}

	return opts.Print(r)
}

func (opts *InviteOpts) newInvitation() *atlas.Invitation {
	return &atlas.Invitation{
		Username: opts.username,
		Roles:    opts.roles,
		TeamIDs:  opts.teamIds,
	}
}

// mongocli iam organization(s) invitation(s) invite|create <email> --role role [--teamId teamId] [--orgId orgId].
func InviteBuilder() *cobra.Command {
	opts := new(InviteOpts)
	opts.Template = createTemplate
	cmd := &cobra.Command{
		Use:     "invite <email>",
		Short:   "Invite the specified MongoDB user to your organization.",
		Long:    fmt.Sprintf(usage.RequiredRole, "Organization User Admin"),
		Aliases: []string{"create"},
		Args:    require.ExactArgs(1),
		Annotations: map[string]string{
			"emailDesc": "Email address that belongs to the user that you want to invite to the organization.",
		},
		Example: fmt.Sprintf(`  # Invite the MongoDB user with the email user@example.com to the organization with the ID 5f71e5255afec75a3d0f96dc with ORG_OWNER access:
  %s organizations invitations invite user@example.com --orgId 5f71e5255afec75a3d0f96dc --role ORG_OWNER --output json`, cli.ExampleAtlasEntryPoint()),
		PreRunE: func(cmd *cobra.Command, args []string) error {
			opts.OutWriter = cmd.OutOrStdout()
			return opts.PreRunE(opts.ValidateOrgID, opts.initStore(cmd.Context()))
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			opts.username = args[0]
			return opts.Run()
		},
	}

	if config.BinName() == config.MongoCLI {
		cmd.Flags().StringSliceVar(&opts.roles, flag.Role, []string{}, usage.MCLIOrgRole)
	} else {
		cmd.Flags().StringSliceVar(&opts.roles, flag.Role, []string{}, usage.OrgRole)
	}
	cmd.Flags().StringSliceVar(&opts.teamIds, flag.TeamID, []string{}, usage.TeamID)
	cmd.Flags().StringVar(&opts.OrgID, flag.OrgID, "", usage.OrgID)

	cmd.Flags().StringVarP(&opts.Output, flag.Output, flag.OutputShort, "", usage.FormatOut)
	_ = cmd.RegisterFlagCompletionFunc(flag.Output, opts.AutoCompleteOutputFlag())

	_ = cmd.MarkFlagRequired(flag.Role)

	return cmd
}
