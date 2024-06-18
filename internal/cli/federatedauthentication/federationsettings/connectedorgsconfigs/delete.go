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
)

type DeleteOpts struct {
	cli.GlobalOpts
	cli.OutputOpts
	cli.InputOpts
	*cli.DeleteOpts
	federationSettingsID string
	store                store.ConnectedOrgConfigsDeleter
}

func (opts *DeleteOpts) InitStore(ctx context.Context) func() error {
	return func() error {
		if opts.store != nil {
			return nil
		}

		var err error
		opts.store, err = store.New(store.AuthenticatedPreset(config.Default()), store.WithContext(ctx))
		return err
	}
}

func (opts *DeleteOpts) Run() error {
	return opts.Delete(opts.store.DeleteConnectedOrgConfig, opts.federationSettingsID)
}

// atlas federatedAuthentication federationSettings connectedOrgConfigs delete --federationSettingsId federationSettingsId [-o/--output output].
func DeleteBuilder() *cobra.Command {
	opts := &DeleteOpts{
		DeleteOpts: cli.NewDeleteOpts("Connected Org Config with OrgID %s deleted.\n", "Connected Org Config with OrgID %s not deleted."),
	}
	cmd := &cobra.Command{
		Use:   "delete",
		Short: "Delete a connected org config Organization.",
		Args:  cobra.NoArgs,
		Example: `  # Delete a connected org config from the current profile org and federationSettingsId 5d1113b25a115342acc2d1aa 
			atlas federatedAuthentication federationSettings connectedOrgConfigs delete --federationSettingsId 5d1113b25a115342acc2d1aa 
			# Delete a connected org config from the org with ID 7d1113b25a115342acc2d1aa and federationSettingsId 5d1113b25a115342acc2d1aa 
			atlas federatedAuthentication federationSettings connectedOrgConfigs delete --orgId 7d1113b25a115342acc2d1aa --federationSettingsId 5d1113b25a115342acc2d1aa 
		`,
		PreRunE: func(cmd *cobra.Command, _ []string) error {
			err := opts.PreRunE(
				opts.ValidateOrgID,
				opts.InitStore(cmd.Context()),
				opts.InitOutput(cmd.OutOrStdout(), opts.SuccessMessage()))
			if err != nil {
				return err
			}

			opts.Entry = opts.ConfigOrgID()
			return opts.Prompt()
		},
		RunE: func(_ *cobra.Command, _ []string) error {
			return opts.Run()
		},
	}

	cmd.Flags().StringVar(&opts.federationSettingsID, flag.FederationSettingsID, "", usage.FederationSettingsID)
	cmd.Flags().StringVar(&opts.OrgID, flag.OrgID, "", usage.OrgID)
	cmd.Flags().BoolVar(&opts.Confirm, flag.Force, false, usage.Force)

	cmd.Flags().StringVarP(&opts.Output, flag.Output, flag.OutputShort, "", usage.FormatOut)
	_ = cmd.RegisterFlagCompletionFunc(flag.Output, opts.AutoCompleteOutputFlag())

	_ = cmd.MarkFlagRequired(flag.FederationSettingsID)

	return cmd
}
