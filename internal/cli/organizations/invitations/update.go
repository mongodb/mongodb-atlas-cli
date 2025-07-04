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
	"errors"
	"fmt"

	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/config"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/file"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/flag"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/store"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/usage"
	"github.com/spf13/afero"
	"github.com/spf13/cobra"
	atlasv2 "go.mongodb.org/atlas-sdk/v20250312005/admin"
)

const updateTemplate = "Invitation {{.Id}} updated.\n"

//go:generate go tool go.uber.org/mock/mockgen -typed -destination=update_mock_test.go -package=invitations . OrganizationInvitationUpdater

type OrganizationInvitationUpdater interface {
	UpdateOrganizationInvitation(string, string, *atlasv2.OrganizationInvitationRequest) (*atlasv2.OrganizationInvitation, error)
}

type UpdateOpts struct {
	cli.OrgOpts
	cli.ProjectOpts
	cli.OutputOpts
	store        OrganizationInvitationUpdater
	invitationID string
	username     string
	roles        []string
	filename     string
	fs           afero.Fs
}

func (opts *UpdateOpts) initStore(ctx context.Context) func() error {
	return func() error {
		var err error
		opts.store, err = store.New(store.AuthenticatedPreset(config.Default()), store.WithContext(ctx))
		return err
	}
}

func (opts *UpdateOpts) Run() error {
	request, err := opts.newInvitation()
	if err != nil {
		return err
	}

	r, err := opts.store.UpdateOrganizationInvitation(opts.ConfigOrgID(), opts.invitationID, request)
	if err != nil {
		return err
	}

	return opts.Print(r)
}

func (opts *UpdateOpts) newInvitation() (*atlasv2.OrganizationInvitationRequest, error) {
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
	}, nil
}

func (opts *UpdateOpts) validate() error {
	if opts.username == "" && opts.invitationID == "" {
		return errors.New("you must provide the email address or the invitationId")
	}

	return nil
}

// atlas organization(s) invitation(s) updates [invitationId] --role role  [--orgId orgId] [--email email].
func UpdateBuilder() *cobra.Command {
	opts := &UpdateOpts{}
	opts.fs = afero.NewOsFs()
	cmd := &cobra.Command{
		Use:     "update [invitationId]",
		Aliases: []string{"updates"},
		Short:   "Modifies the details of the specified pending invitation to your organization.",
		Long: `You can use either the invitation ID or the user's email address to specify the invitation.

` + fmt.Sprintf(usage.RequiredRole, "Organization Owner"),
		Annotations: map[string]string{
			"invitationIdDesc": "Unique 24-digit string that identifies the invitation.",
			"output":           updateTemplate,
		},
		Example: `  # Modify the pending invitation with the ID 5dd56c847a3e5a1f363d424d to grant ORG_OWNER access the organization with the ID 5f71e5255afec75a3d0f96dc:
  atlas organizations invitations update 5dd56c847a3e5a1f363d424d --orgId 5f71e5255afec75a3d0f96dc --role ORG_OWNER --output json
		
  # Modify the invitation for the user with the email address user@example.com to grant ORG_OWNER access the organization with the ID 5f71e5255afec75a3d0f96dc:
  atlas organizations invitations update --email user@example.com --orgId 5f71e5255afec75a3d0f96dc --role ORG_OWNER --output json`,
		PreRunE: func(cmd *cobra.Command, args []string) error {
			if len(args) > 0 {
				opts.invitationID = args[0]
			}

			return opts.OrgOpts.PreRunE(
				opts.ValidateOrgID,
				opts.ValidateProjectID,
				opts.validate,
				opts.initStore(cmd.Context()),
				opts.InitOutput(cmd.OutOrStdout(), updateTemplate),
			)
		},
		RunE: func(_ *cobra.Command, _ []string) error {
			return opts.Run()
		},
	}

	cmd.Flags().StringVarP(&opts.filename, flag.File, flag.FileShort, "", usage.InvitationFile)
	cmd.Flags().StringVar(&opts.username, flag.Email, "", usage.Email)
	cmd.Flags().StringSliceVar(&opts.roles, flag.Role, []string{}, usage.OrgRole+usage.UpdateWarning)
	opts.AddOrgOptFlags(cmd)
	opts.AddOutputOptFlags(cmd)

	_ = cmd.MarkFlagFilename(flag.File)
	cmd.MarkFlagsMutuallyExclusive(flag.File, flag.Role)

	return cmd
}
