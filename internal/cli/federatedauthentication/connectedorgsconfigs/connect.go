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
	atlasv2 "go.mongodb.org/atlas-sdk/v20231115010/admin"
)

type ConnectOpts struct {
	cli.GlobalOpts
	cli.OutputOpts
	cli.InputOpts
	identityProviderID   string
	federationSettingsID string
	store                store.ConnectedOrgConfigsUpdater
}

const (
	updateTemplate = "Connected Org Config '{{.Id}}' connected.\n"
)

func (opts *ConnectOpts) InitStore(ctx context.Context) func() error {
	return func() error {
		if opts.store != nil {
			return nil
		}

		var err error
		opts.store, err = store.New(store.AuthenticatedPreset(config.Default()), store.WithContext(ctx))
		return err
	}
}

func (opts *ConnectOpts) Run() error {
	params := &atlasv2.UpdateConnectedOrgConfigApiParams{
		FederationSettingsId: opts.federationSettingsID,
		OrgId:                opts.ConfigOrgID(),
		ConnectedOrgConfig: &atlasv2.ConnectedOrgConfig{
			IdentityProviderId: opts.identityProviderID,
			OrgId:              opts.ConfigOrgID(),
		},
	}

	r, err := opts.store.UpdateConnectedOrgConfig(params)
	if err != nil {
		return err
	}

	return opts.Print(r)
}

// atlas federatedAuthentication connectedOrgs connect --identityProviderId identityProviderId --federationSettingsId federationSettingsId [-o/--output output].
func ConnectBuilder() *cobra.Command {
	opts := &ConnectOpts{}
	cmd := &cobra.Command{
		Use:   "connect",
		Short: "Connect an Identity Provider to an Organization.",
		Args:  cobra.NoArgs,
		Annotations: map[string]string{
			"output": updateTemplate,
		},
		Example: `  # Connect the current profile Org with Identity Provider with ID 5d1113b25a115342acc2d1aa and  federationSettingsId 5d1113b25a115342acc2d1aa 
			atlas federatedAuthentication connectedOrgs connect --identityProviderId 5d1113b25a115342acc2d1aa --federationSettingsId 5d1113b25a115342acc2d1aa 
			# Connect the Org with ID 7d1113b25a115342acc2d1aa with Identity Provider with ID 5d1113b25a115342acc2d1aa and  federationSettingsId 5d1113b25a115342acc2d1aa 
			atlas federatedAuthentication connectedOrgs connect --orgId 7d1113b25a115342acc2d1aa --identityProviderId 5d1113b25a115342acc2d1aa --federationSettingsId 5d1113b25a115342acc2d1aa 
		`,
		PreRunE: func(cmd *cobra.Command, _ []string) error {
			return opts.PreRunE(
				opts.ValidateOrgID,
				opts.InitStore(cmd.Context()),
				opts.InitOutput(cmd.OutOrStdout(), updateTemplate),
			)
		},
		RunE: func(_ *cobra.Command, args []string) error {
			return opts.Run()
		},
	}

	cmd.Flags().StringVar(&opts.federationSettingsID, flag.FederationSettingsID, "", usage.FederationSettingsID)
	cmd.Flags().StringVar(&opts.identityProviderID, flag.IdentityProviderID, "", usage.IdentityProviderID)
	cmd.Flags().StringVar(&opts.OrgID, flag.OrgID, "", usage.OrgID)

	cmd.Flags().StringVarP(&opts.Output, flag.Output, flag.OutputShort, "", usage.FormatOut)
	_ = cmd.RegisterFlagCompletionFunc(flag.Output, opts.AutoCompleteOutputFlag())

	_ = cmd.MarkFlagRequired(flag.FederationSettingsID)
	_ = cmd.MarkFlagRequired(flag.IdentityProviderID)

	return cmd
}
