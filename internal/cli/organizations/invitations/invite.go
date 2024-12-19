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
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/file"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/flag"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/store"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/usage"
	"github.com/spf13/afero"
	"github.com/spf13/cobra"
	atlasv2 "go.mongodb.org/atlas-sdk/v20241113004/admin"
)

const createTemplate = "User '{{.Username}}' invited.\n"

type InviteOpts struct {
	cli.OutputOpts
	cli.OrgOpts
	username string
	roles    []string
	teamIDs  []string
	store    store.OrganizationInviter
	filename string
	fs       afero.Fs
}

func (opts *InviteOpts) initStore(ctx context.Context) func() error {
	return func() error {
		var err error
		opts.store, err = store.New(store.AuthenticatedPreset(config.Default()), store.WithContext(ctx))
		return err
	}
}

func (opts *InviteOpts) Run() error {
	request, err := opts.newInvitation()
	if err != nil {
		return err
	}

	r, err := opts.store.InviteUser(opts.ConfigOrgID(), request)
	if err != nil {
		return err
	}

	return opts.Print(r)
}

func (opts *InviteOpts) newInvitation() (*atlasv2.OrganizationInvitationRequest, error) {
	if opts.filename != "" {
		pipeline := &atlasv2.OrganizationInvitationRequest{}
		if err := file.Load(opts.fs, opts.filename, pipeline); err != nil {
			return nil, err
		}
		return pipeline, nil
	}

	return &atlasv2.OrganizationInvitationRequest{
		Username: &opts.username,
		Roles:    &opts.roles,
		TeamIds:  &opts.teamIDs,
	}, nil
}

// atlas organization(s) invitation(s) invite|create <email> --role role [--teamId teamId] [--orgId orgId].
func InviteBuilder() *cobra.Command {
	opts := new(InviteOpts)
	opts.Template = createTemplate
	cmd := &cobra.Command{
		Use:     "invite <email>",
		Short:   "Invite the specified MongoDB user to your organization.",
		Long:    fmt.Sprintf(usage.RequiredRole, "Organization User Admin"),
		Aliases: []string{"create"},
		Args:    require.MaximumNArgs(1),
		Annotations: map[string]string{
			"emailDesc": "Email address that belongs to the user that you want to invite to the organization.",
			"output":    createTemplate,
		},
		Example: `  # Invite the MongoDB user with the email user@example.com to the organization with the ID 5f71e5255afec75a3d0f96dc with ORG_OWNER access:
  atlas organizations invitations invite user@example.com --orgId 5f71e5255afec75a3d0f96dc --role ORG_OWNER --output json`,
		PreRunE: func(cmd *cobra.Command, _ []string) error {
			opts.OutWriter = cmd.OutOrStdout()
			return opts.PreRunE(opts.ValidateOrgID, opts.initStore(cmd.Context()))
		},
		RunE: func(_ *cobra.Command, args []string) error {
			if len(args) == 1 {
				opts.username = args[0]
			}
			return opts.Run()
		},
	}

	cmd.Flags().StringVarP(&opts.filename, flag.File, flag.FileShort, "", usage.InvitationFile)
	cmd.Flags().StringSliceVar(&opts.roles, flag.Role, []string{}, usage.OrgRole)
	cmd.Flags().StringSliceVar(&opts.teamIDs, flag.TeamID, []string{}, usage.TeamID)
	opts.AddOrgOptFlags(cmd)

	opts.AddOutputOptFlags(cmd)

	_ = cmd.MarkFlagFilename(flag.File)
	cmd.MarkFlagsMutuallyExclusive(flag.File, flag.Role)
	cmd.MarkFlagsMutuallyExclusive(flag.File, flag.TeamID)

	return cmd
}
