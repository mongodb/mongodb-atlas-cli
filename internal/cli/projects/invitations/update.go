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
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/flag"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/store"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/usage"
	"github.com/spf13/cobra"
	atlasv2 "go.mongodb.org/atlas-sdk/v20240530002/admin"
)

const updateTemplate = "Invitation {{.Id}} updated.\n"

type UpdateOpts struct {
	cli.GlobalOpts
	cli.OutputOpts
	store        store.ProjectInvitationUpdater
	invitationID string
	username     string
	roles        []string
}

func (opts *UpdateOpts) initStore(ctx context.Context) func() error {
	return func() error {
		var err error
		opts.store, err = store.New(store.AuthenticatedPreset(config.Default()), store.WithContext(ctx))
		return err
	}
}

func (opts *UpdateOpts) Run() error {
	r, err := opts.store.UpdateProjectInvitation(opts.ConfigProjectID(), opts.invitationID, opts.newInvitation())
	if err != nil {
		return err
	}

	return opts.Print(r)
}

func (opts *UpdateOpts) newInvitation() *atlasv2.GroupInvitationRequest {
	return &atlasv2.GroupInvitationRequest{
		Username: &opts.username,
		Roles:    &opts.roles,
	}
}

func (opts *UpdateOpts) validate() error {
	if opts.username == "" && opts.invitationID == "" {
		return errors.New("you must provide the email address or the invitationId")
	}

	return nil
}

// atlas project(s) invitation(s) updates [invitationId] --role role  [--orgId orgId] [--email email].
func UpdateBuilder() *cobra.Command {
	opts := &UpdateOpts{}
	cmd := &cobra.Command{
		Use:     "update [invitationId]",
		Aliases: []string{"updates"},
		Short:   "Modifies the details of the specified pending invitation to your project.",
		Long: `You can use either the invitation ID or the user's email address to specify the invitation.

` + fmt.Sprintf(usage.RequiredRole, "Project Owner"),
		Annotations: map[string]string{
			"invitationIdDesc": "Unique 24-digit string that identifies the invitation.",
			"output":           updateTemplate,
		},
		Example: `  # Modify the pending invitation with the ID 5dd56c847a3e5a1f363d424d to grant GROUP_READ_ONLY access the project with the ID 5f71e5255afec75a3d0f96dc:
  atlas projects invitations update 5dd56c847a3e5a1f363d424d --projectId 5f71e5255afec75a3d0f96dc --role GROUP_READ_ONLY --output json
  
  # Modify the invitation for the user with the email address user@example.com to grant GROUP_READ_ONLY access the project with the ID 5f71e5255afec75a3d0f96dc:
  atlas projects invitations update --email user@example.com --projectId 5f71e5255afec75a3d0f96dc --role GROUP_READ_ONLY --output json`,
		PreRunE: func(cmd *cobra.Command, args []string) error {
			if len(args) > 0 {
				opts.invitationID = args[0]
			}

			return opts.PreRunE(
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

	cmd.Flags().StringVar(&opts.username, flag.Email, "", usage.Email)
	cmd.Flags().StringSliceVar(&opts.roles, flag.Role, []string{}, usage.OrgRole+usage.UpdateWarning)
	cmd.Flags().StringVar(&opts.ProjectID, flag.ProjectID, "", usage.ProjectID)
	cmd.Flags().StringVarP(&opts.Output, flag.Output, flag.OutputShort, "", usage.FormatOut)
	_ = cmd.RegisterFlagCompletionFunc(flag.Output, opts.AutoCompleteOutputFlag())

	_ = cmd.MarkFlagRequired(flag.Role)

	return cmd
}
