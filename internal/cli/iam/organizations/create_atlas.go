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

	"github.com/mongodb/mongodb-atlas-cli/internal/cli"
	"github.com/mongodb/mongodb-atlas-cli/internal/cli/require"
	"github.com/mongodb/mongodb-atlas-cli/internal/config"
	"github.com/mongodb/mongodb-atlas-cli/internal/flag"
	"github.com/mongodb/mongodb-atlas-cli/internal/prerun"
	"github.com/mongodb/mongodb-atlas-cli/internal/store"
	"github.com/mongodb/mongodb-atlas-cli/internal/usage"
	"github.com/spf13/cobra"
	"go.mongodb.org/atlas/mongodbatlas"
)

var createAtlasTemplate = "Organization '{{.Organization.ID}}' created.\n"

type CreateAtlasOpts struct {
	cli.OutputOpts
	name              string
	ownerID           string
	apiKeyDescription string
	apiKeyRole        []string
	store             store.AtlasOrganizationCreator
}

func (opts *CreateAtlasOpts) initStore(ctx context.Context) func() error {
	return func() error {
		var err error
		opts.store, err = store.New(store.AuthenticatedPreset(config.Default()), store.WithContext(ctx))
		return err
	}
}

func (opts *CreateAtlasOpts) Run() error {
	o := &mongodbatlas.CreateOrganizationRequest{
		Name: opts.name,
	}
	if opts.ownerID != "" {
		o.OrgOwnerID = &opts.ownerID
	}
	if len(opts.apiKeyRole) > 0 {
		o.APIKey = &mongodbatlas.APIKeyInput{}
		o.APIKey.Roles = opts.apiKeyRole
		o.APIKey.Desc = opts.apiKeyDescription
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
		Short: "Create an Ops Manager or Cloud Manager organization. This command is unavailable for Atlas.",
		Args:  require.ExactArgs(1),
		Annotations: map[string]string{
			"nameDesc": "Label that identifies the organization.",
		},
		Example: `  # Create an Atlas organization with the name myOrg:
  atlas organizations create myOrg --output json`,
		PreRunE: func(cmd *cobra.Command, args []string) error {
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
		RunE: func(cmd *cobra.Command, args []string) error {
			opts.name = args[0]

			return opts.Run()
		},
	}
	cmd.Flags().StringVar(&opts.ownerID, flag.OwnerID, "", usage.OrgOwnerID)
	cmd.Flags().StringVar(&opts.apiKeyDescription, flag.APIKeyDescription, "", usage.APIKeyDescription)
	cmd.Flags().StringSliceVar(&opts.apiKeyRole, flag.APIKeyRole, []string{}, usage.APIKeyRoles)
	cmd.Flags().StringVarP(&opts.Output, flag.Output, flag.OutputShort, "", usage.FormatOut)
	_ = cmd.RegisterFlagCompletionFunc(flag.Output, opts.AutoCompleteOutputFlag())

	return cmd
}
