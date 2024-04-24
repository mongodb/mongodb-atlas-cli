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

type DisconnectOpts struct {
	cli.GlobalOpts
	cli.OutputOpts
	cli.InputOpts
	federationSettingsID string
	store                store.ConnectedOrgConfigsUpdater
}

const (
	disconnectTemplate = "Connected Org Config '{{.Id}}' disconnected.\n"
)

func (opts *DisconnectOpts) InitStore(ctx context.Context) func() error {
	return func() error {
		if opts.store != nil {
			return nil
		}

		var err error
		opts.store, err = store.New(store.AuthenticatedPreset(config.Default()), store.WithContext(ctx))
		return err
	}
}

func (opts *DisconnectOpts) Run() error {
	params := &atlasv2.UpdateConnectedOrgConfigApiParams{
		FederationSettingsId: opts.federationSettingsID,
		OrgId:                opts.ConfigOrgID(),
		ConnectedOrgConfig: &atlasv2.ConnectedOrgConfig{
			// IdentityProviderId: "",
			OrgId: opts.ConfigOrgID(),
		},
	}

	r, err := opts.store.UpdateConnectedOrgConfig(params)
	if err != nil {
		return err
	}

	return opts.Print(r)
}

// atlas federatedAuthentication connectedOrgs disconnect --orgId orgId --federationSettingsId federationSettingsId [-o/--output output].
func DisconnectBuilder() *cobra.Command {
	opts := &DisconnectOpts{}
	cmd := &cobra.Command{
		Use:   "disconnect",
		Short: "Disconnect an Identity Provider to an Organization.",
		Args:  cobra.NoArgs,
		Annotations: map[string]string{
			"output": disconnectTemplate,
		},
		Example: `  # Disconnect the current profile Org with federationSettingsId 5d1113b25a115342acc2d1aa from the connected Identity Provider
			atlas federatedAuthentication connectedOrgs disconnect --federationSettingsId 5d1113b25a115342acc2d1aa 
			# Connect the Org with ID 7d1113b25a115342acc2d1aa and federationSettingsId 5d1113b25a115342acc2d1aa  from the connected Identity Provider
			atlas federatedAuthentication connectedOrgs connect --orgId 7d1113b25a115342acc2d1aa --federationSettingsId 5d1113b25a115342acc2d1aa 
		`,
		PreRunE: func(cmd *cobra.Command, _ []string) error {
			return opts.PreRunE(
				opts.ValidateOrgID,
				opts.InitStore(cmd.Context()),
				opts.InitOutput(cmd.OutOrStdout(), disconnectTemplate),
			)
		},
		RunE: func(_ *cobra.Command, args []string) error {
			return opts.Run()
		},
	}

	cmd.Flags().StringVar(&opts.federationSettingsID, flag.FederationSettingsID, "", usage.FederationSettingsID)
	cmd.Flags().StringVar(&opts.OrgID, flag.OrgID, "", usage.OrgID)

	cmd.Flags().StringVarP(&opts.Output, flag.Output, flag.OutputShort, "", usage.FormatOut)
	_ = cmd.RegisterFlagCompletionFunc(flag.Output, opts.AutoCompleteOutputFlag())

	_ = cmd.MarkFlagRequired(flag.FederationSettingsID)

	return cmd
}
