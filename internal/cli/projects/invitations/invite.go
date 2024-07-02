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
	atlasv2 "go.mongodb.org/atlas-sdk/v20240530002/admin"
)

const createTemplate = "User '{{.Username}}' invited.\n"

type InviteOpts struct {
	cli.OutputOpts
	cli.GlobalOpts
	username string
	roles    []string
	teamIDs  []string
	store    store.ProjectInviter
}

func (opts *InviteOpts) initStore(ctx context.Context) func() error {
	return func() error {
		var err error
		opts.store, err = store.New(store.AuthenticatedPreset(config.Default()), store.WithContext(ctx))
		return err
	}
}

func (opts *InviteOpts) Run() error {
	r, err := opts.store.InviteUserToProject(opts.ConfigProjectID(), opts.newInvitation())

	if err != nil {
		return err
	}

	return opts.Print(r)
}

func (opts *InviteOpts) newInvitation() *atlasv2.GroupInvitationRequest {
	return &atlasv2.GroupInvitationRequest{
		Username: &opts.username,
		Roles:    &opts.roles,
	}
}

// InviteBuilder atlas project(s) invitation(s) invite|create <email> --role role [--teamId teamId] [--orgId orgId].
func InviteBuilder() *cobra.Command {
	opts := new(InviteOpts)
	opts.Template = createTemplate
	cmd := &cobra.Command{
		Use:     "invite <email>",
		Short:   "Invite the specified MongoDB user to your project.",
		Long:    fmt.Sprintf(usage.RequiredRole, "Project User Admin"),
		Aliases: []string{"create"},
		Args:    require.ExactArgs(1),
		Annotations: map[string]string{
			"emailDesc": "Email address that belongs to the user that you want to invite to the project.",
			"output":    createTemplate,
		},
		Example: `  # Invite the MongoDB user with the email user@example.com to the project with the ID 5f71e5255afec75a3d0f96dc with GROUP_READ_ONLY access:
  atlas projects invitations invite user@example.com --projectId 5f71e5255afec75a3d0f96dc --role GROUP_READ_ONLY --output json`,
		PreRunE: func(cmd *cobra.Command, _ []string) error {
			opts.OutWriter = cmd.OutOrStdout()
			return opts.PreRunE(
				opts.initStore(cmd.Context()),
			)
		},
		RunE: func(_ *cobra.Command, args []string) error {
			opts.username = args[0]
			return opts.Run()
		},
	}

	cmd.Flags().StringSliceVar(&opts.roles, flag.Role, []string{}, usage.ProjectRole)
	cmd.Flags().StringSliceVar(&opts.teamIDs, flag.TeamID, []string{}, usage.TeamID)
	cmd.Flags().StringVar(&opts.ProjectID, flag.ProjectID, "", usage.ProjectID)

	cmd.Flags().StringVarP(&opts.Output, flag.Output, flag.OutputShort, "", usage.FormatOut)
	_ = cmd.RegisterFlagCompletionFunc(flag.Output, opts.AutoCompleteOutputFlag())

	_ = cmd.MarkFlagRequired(flag.Role)

	return cmd
}
