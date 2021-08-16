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
	"errors"

	"github.com/mongodb/mongocli/internal/cli"
	"github.com/mongodb/mongocli/internal/config"
	"github.com/mongodb/mongocli/internal/flag"
	"github.com/mongodb/mongocli/internal/store"
	"github.com/mongodb/mongocli/internal/usage"
	"github.com/spf13/cobra"
	atlas "go.mongodb.org/atlas/mongodbatlas"
)

const updateTemplate = "Invitation {{.ID}} updated.\n"

type UpdateOpts struct {
	cli.GlobalOpts
	cli.OutputOpts
	store        store.OrganizationInvitationUpdater
	invitationID string
	username     string
	roles        []string
}

func (opts *UpdateOpts) init(ctx context.Context) func() error {
	return func() error {
		var err error
		opts.store, err = store.New(store.AuthenticatedPreset(config.Default()), store.WithContext(ctx))
		return err
	}
}

func (opts *UpdateOpts) Run() error {
	r, err := opts.store.UpdateOrganizationInvitation(opts.ConfigOrgID(), opts.invitationID, opts.newInvitation())
	if err != nil {
		return err
	}

	return opts.Print(r)
}

func (opts *UpdateOpts) newInvitation() *atlas.Invitation {
	return &atlas.Invitation{
		Username: opts.username,
		Roles:    opts.roles,
	}
}

func (opts *UpdateOpts) validate() error {
	if opts.username == "" && opts.invitationID == "" {
		return errors.New("you must provide the email address or the invitationId")
	}

	return nil
}

// mongocli iam organization(s) invitation(s) updates [invitationId] --role role  [--orgId orgId] [--email email].
func UpdateBuilder() *cobra.Command {
	opts := &UpdateOpts{}
	cmd := &cobra.Command{
		Use:     "update [invitationId]",
		Aliases: []string{"updates"},
		Short:   "Updates one pending invitation by invitationId or email to the project that you specify.",
		Annotations: map[string]string{
			"args":             "invitationId",
			"invitationIdDesc": "Unique 24-hexadecimal digit string that identifies the invitation.",
		},
		PreRunE: func(cmd *cobra.Command, args []string) error {
			if len(args) > 0 {
				opts.invitationID = args[0]
			}

			return opts.PreRunE(
				opts.ValidateProjectID,
				opts.validate,
				opts.init(cmd.Context()),
				opts.InitOutput(cmd.OutOrStdout(), updateTemplate),
			)
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return opts.Run()
		},
	}

	cmd.Flags().StringVar(&opts.username, flag.Email, "", usage.Email)
	cmd.Flags().StringSliceVar(&opts.roles, flag.Role, []string{}, usage.OrgRole)

	cmd.Flags().StringVar(&opts.OrgID, flag.OrgID, "", usage.OrgID)
	cmd.Flags().StringVarP(&opts.Output, flag.Output, flag.OutputShort, "", usage.FormatOut)

	_ = cmd.MarkFlagRequired(flag.Role)

	return cmd
}
