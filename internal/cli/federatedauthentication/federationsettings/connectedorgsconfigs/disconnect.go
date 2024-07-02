// Copyright 2024 MongoDB Inc
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

package connectedorgsconfigs

import (
	"context"

	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/config"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/flag"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/store"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/usage"
	"github.com/spf13/cobra"
	atlasv2 "go.mongodb.org/atlas-sdk/v20240530002/admin"
)

type DisconnectOpts struct {
	cli.GlobalOpts
	cli.OutputOpts
	cli.InputOpts
	*DescribeOrgConfigsOpts
	identityProviderID   string
	protocol             string
	federationSettingsID string
	store                store.ConnectedOrgConfigsUpdater
}

const disconnectTemplate = `Org Configs disconnected.`

func (opts *DisconnectOpts) InitStore(ctx context.Context) func() error {
	return func() error {
		if opts.store != nil && opts.describeStore != nil {
			return nil
		}

		var err error
		opts.store, err = store.New(store.AuthenticatedPreset(config.Default()), store.WithContext(ctx))
		if err != nil {
			return err
		}

		return opts.InitDescribeStore(ctx)()
	}
}

func (opts *DisconnectOpts) removeIdpFromOrgConfig(orgConfig *atlasv2.ConnectedOrgConfig) {
	if opts.protocol == oidc {
		if orgConfig.DataAccessIdentityProviderIds == nil {
			return
		}
		// search and remove id from orgConfig.DataAccessIdentityProviderIds
		newIdps := []string{}
		for _, idp := range *orgConfig.DataAccessIdentityProviderIds {
			if idp != opts.identityProviderID {
				newIdps = append(newIdps, idp)
			}
		}
		orgConfig.DataAccessIdentityProviderIds = &newIdps
	} else if opts.protocol == saml {
		emptyID := ""
		orgConfig.IdentityProviderId = &emptyID
	}
}

func (opts *DisconnectOpts) Run() error {
	var orgConfig *atlasv2.ConnectedOrgConfig
	var err error
	if orgConfig, err = opts.GetConnectedOrgConfig(opts.federationSettingsID, opts.ConfigOrgID()); err != nil {
		return err
	}

	if len(orgConfig.GetRoleMappings()) == 0 {
		orgConfig.RoleMappings = nil
	}

	params := &atlasv2.UpdateConnectedOrgConfigApiParams{
		FederationSettingsId: opts.federationSettingsID,
		OrgId:                opts.ConfigOrgID(),
		ConnectedOrgConfig:   orgConfig,
	}

	opts.removeIdpFromOrgConfig(orgConfig)

	r, err := opts.store.UpdateConnectedOrgConfig(params)
	if err != nil {
		return err
	}

	return opts.Print(r)
}

// atlas federatedAuthentication federationSettings connectedOrgConfigs disconnect --identityProviderId identityProviderId --protocol protocol --federationSettingsId federationSettingsId [-o/--output output].
func DisconnectBuilder() *cobra.Command {
	opts := &DisconnectOpts{
		DescribeOrgConfigsOpts: &DescribeOrgConfigsOpts{},
	}
	cmd := &cobra.Command{
		Use:   "disconnect",
		Short: "Connect an Identity Provider to an Organization.",
		Args:  cobra.NoArgs,
		Annotations: map[string]string{
			"output": disconnectTemplate,
		},
		Example: `  # Disconnect the current profile org from identity provider with ID 5d1113b25a115342acc2d1aa and federationSettingsId 7d1113b25a115342acc2d1aa 
			atlas federatedAuthentication federationSettings connectedOrgConfigs disconnect --identityProviderId 5d1113b25a115342acc2d1aa --federationSettingsId 7d1113b25a115342acc2d1aa 
			# Disconnect the org with ID 7d1113b25a115342acc2d1aa to identity provider with ID 5d1113b25a115342acc2d1aa and federationSettingsId 7d1113b25a115342acc2d1aa 
			atlas federatedAuthentication federationSettings connectedOrgConfigs disconnect --orgId 7d1113b25a115342acc2d1aa --identityProviderId 5d1113b25a115342acc2d1aa --federationSettingsId 7d1113b25a115342acc2d1aa 
		`,
		PreRunE: func(cmd *cobra.Command, _ []string) error {
			return opts.PreRunE(
				opts.ValidateOrgID,
				opts.InitStore(cmd.Context()),
				opts.InitOutput(cmd.OutOrStdout(), disconnectTemplate),
			)
		},
		RunE: func(_ *cobra.Command, _ []string) error {
			return opts.Run()
		},
	}

	cmd.Flags().StringVar(&opts.federationSettingsID, flag.FederationSettingsID, "", usage.FederationSettingsID)
	cmd.Flags().StringVar(&opts.identityProviderID, flag.IdentityProviderID, "", usage.IdentityProviderID)
	cmd.Flags().StringVar(&opts.protocol, flag.Protocol, oidc, usage.Protocol)
	cmd.Flags().StringVar(&opts.OrgID, flag.OrgID, "", usage.OrgID)

	cmd.Flags().StringVarP(&opts.Output, flag.Output, flag.OutputShort, "", usage.FormatOut)
	_ = cmd.RegisterFlagCompletionFunc(flag.Output, opts.AutoCompleteOutputFlag())

	_ = cmd.MarkFlagRequired(flag.FederationSettingsID)
	_ = cmd.MarkFlagRequired(flag.IdentityProviderID)

	return cmd
}
