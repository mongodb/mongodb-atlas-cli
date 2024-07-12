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

const createTemplate = `API Key '{{.Id}}' created.
Public API Key {{.PublicKey}}
Private API Key {{.PrivateKey}}
`

type CreateOpts struct {
	cli.GlobalOpts
	cli.OutputOpts
	desc  string
	roles []string
	store store.OrganizationAPIKeyCreator
}

func (opts *CreateOpts) initStore(ctx context.Context) func() error {
	return func() error {
		var err error
		opts.store, err = store.New(store.AuthenticatedPreset(config.Default()), store.WithContext(ctx))
		return err
	}
}

func (opts *CreateOpts) newAPIKeyInput() *atlasv2.CreateAtlasOrganizationApiKey {
	return &atlasv2.CreateAtlasOrganizationApiKey{
		Desc:  opts.desc,
		Roles: opts.roles,
	}
}

func (opts *CreateOpts) Run() error {
	r, err := opts.store.CreateOrganizationAPIKey(opts.ConfigOrgID(), opts.newAPIKeyInput())

	if err != nil {
		return err
	}

	return opts.Print(r)
}

// atlas organizations|orgs apiKey(s)|apikeys create [--role role][--desc description][--orgId orgId].
func CreateBuilder() *cobra.Command {
	opts := new(CreateOpts)
	cmd := &cobra.Command{
		Use:   "create",
		Short: "Create an API Key for your organization.",
		Long: `MongoDB returns the private API key only once. After you run this command, immediately copy, save, and secure both the public and private API keys.

` + fmt.Sprintf(usage.RequiredRole, "Organization User Admin"),
		Args: require.NoArgs,
		Annotations: map[string]string{
			"output": createTemplate,
		},
		Example: `  # Create an organization API key with organization owner access in the organization with the ID 5a1b39eec902201990f12345:
  atlas organizations apiKeys create --role ORG_OWNER --desc "My API Key" --orgId 5a1b39eec902201990f12345 --output json`,
		PreRunE: func(cmd *cobra.Command, _ []string) error {
			return opts.PreRunE(
				opts.ValidateOrgID,
				opts.initStore(cmd.Context()),
				opts.InitOutput(cmd.OutOrStdout(), createTemplate),
			)
		},
		RunE: func(_ *cobra.Command, _ []string) error {
			return opts.Run()
		},
	}

	cmd.Flags().StringSliceVar(&opts.roles, flag.Role, []string{}, usage.APIKeyRoles)
	cmd.Flags().StringVar(&opts.desc, flag.Description, "", usage.APIKeyDescription)

	cmd.Flags().StringVar(&opts.OrgID, flag.OrgID, "", usage.OrgID)
	cmd.Flags().StringVarP(&opts.Output, flag.Output, flag.OutputShort, "", usage.FormatOut)
	_ = cmd.RegisterFlagCompletionFunc(flag.Output, opts.AutoCompleteOutputFlag())

	_ = cmd.MarkFlagRequired(flag.Description)
	_ = cmd.MarkFlagRequired(flag.Role)

	return cmd
}
