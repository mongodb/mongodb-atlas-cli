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
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/file"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/flag"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/store"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/usage"
	"github.com/spf13/afero"
	"github.com/spf13/cobra"
	atlasv2 "go.mongodb.org/atlas-sdk/v20240530002/admin"
)

type UpdateOpts struct {
	cli.GlobalOpts
	cli.OutputOpts
	cli.InputOpts
	*DescribeOrgConfigsOpts
	federationSettingsID string
	file                 string
	store                store.ConnectedOrgConfigsUpdater
	fs                   afero.Fs
}

const updateTemplate = `Connected Org Config updated.`

func (opts *UpdateOpts) InitStore(ctx context.Context) func() error {
	return func() error {
		if opts.store != nil {
			return nil
		}

		var err error
		opts.store, err = store.New(store.AuthenticatedPreset(config.Default()), store.WithContext(ctx))
		return err
	}
}

func (opts *UpdateOpts) fillReadOnlyValues(c *atlasv2.ConnectedOrgConfig) {
	c.OrgId = opts.ConfigOrgID()
}

func (opts *UpdateOpts) Run() error {
	orgConfig := new(atlasv2.ConnectedOrgConfig)
	if err := file.Load(opts.fs, opts.file, orgConfig); err != nil {
		return err
	}

	opts.fillReadOnlyValues(orgConfig)

	params := &atlasv2.UpdateConnectedOrgConfigApiParams{
		FederationSettingsId: opts.federationSettingsID,
		OrgId:                opts.ConfigOrgID(),
		ConnectedOrgConfig:   orgConfig,
	}

	r, err := opts.store.UpdateConnectedOrgConfig(params)
	if err != nil {
		return err
	}

	return opts.Print(r)
}

// atlas federatedAuthentication federationSettings connectedOrgConfigs update --federationSettingsId federationSettingsId --file file [-o/--output output].
func UpdateBuilder() *cobra.Command {
	opts := &UpdateOpts{
		fs: afero.NewOsFs(),
	}
	cmd := &cobra.Command{
		Use:   "update",
		Short: "Update One Org Config Connected to One Federation Setting.",
		Args:  cobra.NoArgs,
		Example: `  # Update the connected orgs config with the current profile org and federationSettingsId 5d1113b25a115342acc2d1aa using the JSON configuration file config.json
			atlas federatedAuthentication federationSettings connectedOrgConfigs update --federationSettingsId 5d1113b25a115342acc2d1aa --file config.json
		`,
		PreRunE: func(cmd *cobra.Command, _ []string) error {
			return opts.PreRunE(
				opts.ValidateOrgID,
				opts.InitStore(cmd.Context()),
				opts.InitOutput(cmd.OutOrStdout(), updateTemplate),
			)
		},
		RunE: func(_ *cobra.Command, _ []string) error {
			return opts.Run()
		},
	}

	cmd.Flags().StringVar(&opts.federationSettingsID, flag.FederationSettingsID, "", usage.FederationSettingsID)
	cmd.Flags().StringVar(&opts.OrgID, flag.OrgID, "", usage.OrgID)
	cmd.Flags().StringVar(&opts.file, flag.File, "", usage.ConnectedOrgConfigFilename)

	cmd.Flags().StringVarP(&opts.Output, flag.Output, flag.OutputShort, "", usage.FormatOut)
	_ = cmd.RegisterFlagCompletionFunc(flag.Output, opts.AutoCompleteOutputFlag())

	_ = cmd.MarkFlagRequired(flag.FederationSettingsID)
	_ = cmd.MarkFlagRequired(flag.File)

	return cmd
}
