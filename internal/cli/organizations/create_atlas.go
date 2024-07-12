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

package organizations

import (
	"context"
	"fmt"
	"strings"

	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli/require"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/config"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/flag"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/prerun"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/store"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/usage"
	"github.com/spf13/cobra"
	atlasv2 "go.mongodb.org/atlas-sdk/v20240530002/admin"
)

var createAtlasTemplate = "Organization '{{.Organization.Id}}' created.\n"

type CreateAtlasOpts struct {
	cli.OutputOpts
	name                 string
	ownerID              string
	apiKeyDescription    string
	apiKeyRole           []string
	federationSettingsID string
	store                store.OrganizationCreator
}

func (opts *CreateAtlasOpts) initStore(ctx context.Context) func() error {
	return func() error {
		var err error
		opts.store, err = store.New(store.AuthenticatedPreset(config.Default()), store.WithContext(ctx))
		return err
	}
}

func (opts *CreateAtlasOpts) Run() error {
	o := &atlasv2.CreateOrganizationRequest{
		Name: opts.name,
	}
	if opts.ownerID != "" {
		o.OrgOwnerId = &opts.ownerID
	}

	if opts.federationSettingsID != "" {
		o.FederationSettingsId = &opts.federationSettingsID
	}

	if len(opts.apiKeyRole) > 0 {
		o.ApiKey = &atlasv2.CreateAtlasOrganizationApiKey{}
		o.ApiKey.Roles = opts.apiKeyRole
		o.ApiKey.Desc = opts.apiKeyDescription
	}

	r, err := opts.store.CreateAtlasOrganization(o)
	if err != nil {
		return err
	}

	return opts.Print(r)
}

func (opts *CreateAtlasOpts) validateAPIKeyRequirements() error {
	required := make([]string, 0)
	if len(opts.apiKeyRole) == 0 {
		required = append(required, flag.APIKeyRole)
	}
	if opts.apiKeyDescription == "" {
		required = append(required, flag.APIKeyDescription)
	}
	if opts.ownerID == "" {
		required = append(required, flag.OwnerID)
	}
	if len(required) > 0 {
		return fmt.Errorf(
			"%s are required when using API keys to authenticate",
			strings.Join(required, ", "),
		)
	}
	return nil
}

func (opts *CreateAtlasOpts) validateOAuthRequirements() error {
	disallowed := make([]string, 0)
	if len(opts.apiKeyRole) > 0 {
		disallowed = append(disallowed, flag.APIKeyRole)
	}
	if opts.ownerID != "" {
		disallowed = append(disallowed, flag.OwnerID)
	}
	if len(disallowed) > 0 {
		return fmt.Errorf(
			"%s are not allowed when using account to authenticate",
			strings.Join(disallowed, ", "),
		)
	}
	return nil
}

func (opts *CreateAtlasOpts) validateAuthType() error {
	switch config.AuthType() {
	case config.APIKeys:
		return opts.validateAPIKeyRequirements()
	case config.OAuth:
		return opts.validateOAuthRequirements()
	case config.NotLoggedIn:
		return nil // should not happen
	default:
		return nil
	}
}

// CreateAtlasBuilder atlas organization(s) create <name>.
func CreateAtlasBuilder() *cobra.Command {
	opts := new(CreateAtlasOpts)

	cmd := &cobra.Command{
		Use:   "create <name>",
		Short: "Create an organization.",
		Long:  "When authenticating using API keys, the organization to which the API keys belong must have cross-organization billing enabled. The resulting org will be linked to the paying org.",
		Args:  require.ExactArgs(1),
		Annotations: map[string]string{
			"nameDesc": "Label that identifies the organization.",
		},
		Example: `  # Create an Atlas organization with the name myOrg:
  atlas organizations create myOrg --output json`,
		PreRunE: func(cmd *cobra.Command, _ []string) error {
			if len(opts.apiKeyRole) > 0 {
				createAtlasTemplate += `API Key '{{.APIKey.ID}}' created.
Public API Key '{{.APIKey.PublicKey}}'
Private API Key '{{.APIKey.PrivateKey}}'
`
			}
			return prerun.ExecuteE(
				opts.InitOutput(cmd.OutOrStdout(), createAtlasTemplate),
				opts.validateAuthType,
				opts.initStore(cmd.Context()),
			)
		},
		RunE: func(_ *cobra.Command, args []string) error {
			opts.name = args[0]

			return opts.Run()
		},
	}
	cmd.Flags().StringVar(&opts.ownerID, flag.OwnerID, "", usage.OrgOwnerID)
	cmd.Flags().StringVar(&opts.apiKeyDescription, flag.APIKeyDescription, "", usage.AtlasAPIKeyDescription)
	cmd.Flags().StringVar(&opts.federationSettingsID, flag.FederationSettingsID, "", usage.FederationSettingsID)
	cmd.Flags().StringSliceVar(&opts.apiKeyRole, flag.APIKeyRole, []string{}, usage.AtlasAPIKeyRoles)
	cmd.Flags().StringVarP(&opts.Output, flag.Output, flag.OutputShort, "", usage.FormatOut)
	_ = cmd.RegisterFlagCompletionFunc(flag.Output, opts.AutoCompleteOutputFlag())

	return cmd
}
