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
)

const describeTemplate = `ID	USERNAME	CREATED AT	EXPIRES AT
{{.Id}}	{{.Username}}	{{.CreatedAt}}	{{.ExpiresAt}}
`

type DescribeOpts struct {
	cli.OutputOpts
	cli.GlobalOpts
	id    string
	store store.OrganizationInvitationDescriber
}

func (opts *DescribeOpts) initStore(ctx context.Context) func() error {
	return func() error {
		var err error
		opts.store, err = store.New(store.AuthenticatedPreset(config.Default()), store.WithContext(ctx))
		return err
	}
}

func (opts *DescribeOpts) Run() error {
	r, err := opts.store.OrganizationInvitation(opts.ConfigOrgID(), opts.id)
	if err != nil {
		return err
	}

	return opts.Print(r)
}

// atlas organizations(s) invitations describe|get <ID> [--orgId orgId].
func DescribeBuilder() *cobra.Command {
	opts := new(DescribeOpts)
	cmd := &cobra.Command{
		Use:     "describe <invitationId>",
		Aliases: []string{"get"},
		Args:    require.ExactArgs(1),
		Short:   "Return the details for the specified pending invitation to your organization.",
		Long:    fmt.Sprintf(usage.RequiredRole, "Organization User Admin"),
		Annotations: map[string]string{
			"invitationIdDesc": "Unique 24-digit string that identifies the invitation.",
		},
		Example: `  # Return the JSON-formatted details of the pending invitation with the ID 5dd56c847a3e5a1f363d424d for the organization with the ID 5f71e5255afec75a3d0f96dc:
  atlas organizations invitations describe 5dd56c847a3e5a1f363d424d --orgId 5f71e5255afec75a3d0f96dc --output json`,
		PreRunE: func(cmd *cobra.Command, _ []string) error {
			return opts.PreRunE(
				opts.ValidateOrgID,
				opts.initStore(cmd.Context()),
				opts.InitOutput(cmd.OutOrStdout(), describeTemplate),
			)
		},
		RunE: func(_ *cobra.Command, args []string) error {
			opts.id = args[0]
			return opts.Run()
		},
	}

	cmd.Flags().StringVar(&opts.OrgID, flag.OrgID, "", usage.OrgID)

	cmd.Flags().StringVarP(&opts.Output, flag.Output, flag.OutputShort, "", usage.FormatOut)
	_ = cmd.RegisterFlagCompletionFunc(flag.Output, opts.AutoCompleteOutputFlag())

	return cmd
}
