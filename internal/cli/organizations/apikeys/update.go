// Copyright 2023 MongoDB Inc
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//	http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package apikeys

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

type UpdateOpts struct {
	cli.GlobalOpts
	cli.OutputOpts
	id    string
	desc  string
	roles []string
	store store.OrganizationAPIKeyUpdater
}

func (opts *UpdateOpts) initStore(ctx context.Context) func() error {
	return func() error {
		var err error
		opts.store, err = store.New(store.AuthenticatedPreset(config.Default()), store.WithContext(ctx))
		return err
	}
}

func (opts *UpdateOpts) newAPIKeyInput() *atlasv2.UpdateAtlasOrganizationApiKey {
	return &atlasv2.UpdateAtlasOrganizationApiKey{
		Desc:  &opts.desc,
		Roles: &opts.roles,
	}
}

const updateTemplate = "API Key '{{.Id}}' successfully updated.\n"

func (opts *UpdateOpts) Run() error {
	r, err := opts.store.UpdateOrganizationAPIKey(opts.ConfigOrgID(), opts.id, opts.newAPIKeyInput())
	if err != nil {
		return err
	}

	return opts.Print(r)
}

// atlas organizations|orgs apiKey(s)|apikey(s) update <ID> [--role role][--desc description][--orgId orgId].
func UpdateBuilder() *cobra.Command {
	opts := new(UpdateOpts)
	cmd := &cobra.Command{
		Use:     "assign <apiKeyId>",
		Aliases: []string{"updates"},
		Args:    require.ExactArgs(1),
		Short:   "Modify the roles or description for the specified organization API key.",
		Long: `When you modify the roles for an organization API key with this command, the values you specify overwrite the existing roles assigned to the API key.
		
To view possible values for the apiKeyId argument, run atlas organizations apiKeys list.

` + fmt.Sprintf(usage.RequiredRole, "Organization User Admin"),
		Annotations: map[string]string{
			"apiKeyIdDesc": "Unique 24-digit string that identifies your API key.",
			"output":       updateTemplate,
		},
		Example: `  # Modify the role and description for the API key with the ID 5f24084d8dbffa3ad3f21234 for the organization with the ID 5a1b39eec902201990f12345:
  atlas organizations apiKeys assign 5f24084d8dbffa3ad3f21234 --role ORG_MEMBER --desc "User1 Member Key" --orgId 5a1b39eec902201990f12345 --output json`,
		PreRunE: func(cmd *cobra.Command, _ []string) error {
			return opts.PreRunE(
				opts.ValidateOrgID,
				opts.initStore(cmd.Context()),
				opts.InitOutput(cmd.OutOrStdout(), updateTemplate),
			)
		},
		RunE: func(_ *cobra.Command, args []string) error {
			opts.id = args[0]
			return opts.Run()
		},
	}

	cmd.Flags().StringSliceVar(&opts.roles, flag.Role, []string{}, usage.APIKeyRoles+usage.UpdateWarning)
	cmd.Flags().StringVar(&opts.desc, flag.Description, "", usage.APIKeyDescription)

	cmd.Flags().StringVar(&opts.OrgID, flag.OrgID, "", usage.OrgID)
	cmd.Flags().StringVarP(&opts.Output, flag.Output, flag.OutputShort, "", usage.FormatOut)
	_ = cmd.RegisterFlagCompletionFunc(flag.Output, opts.AutoCompleteOutputFlag())

	return cmd
}
