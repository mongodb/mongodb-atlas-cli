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

type ListOpts struct {
	cli.GlobalOpts
	cli.OutputOpts
	cli.InputOpts
	cli.ListOpts
	store store.ConnectedOrgConfigsLister

	federationSettingsID string
}

const listTemplate = `ORG ID	IDENTITY PROVIDER ID	DATA ACCESS IDENTITY PRODIVER IDs{{range valueOrEmptySlice .Results}}
{{.OrgId}}	{{if .IdentityProviderId }}	{{ .IdentityProviderId }}{{else}}N/A{{end}}	{{if and .DataAccessIdentityProviderIds (gt (len .DataAccessIdentityProviderIds) 0)}}{{range $index, $element := .DataAccessIdentityProviderIds}}{{if $index}}, {{end}}{{$element}}{{end}}{{else}}N/A{{end}}{{end}}`

func (opts *ListOpts) Run() error {
	params := &atlasv2.ListConnectedOrgConfigsApiParams{
		FederationSettingsId: opts.federationSettingsID,
		ItemsPerPage:         &opts.ItemsPerPage,
		PageNum:              &opts.PageNum,
	}

	r, err := opts.store.ListConnectedOrgConfigs(params)
	if err != nil {
		return err
	}

	return opts.Print(r)
}

func (opts *ListOpts) InitStore(ctx context.Context) func() error {
	return func() error {
		if opts.store != nil {
			return nil
		}

		var err error
		opts.store, err = store.New(store.AuthenticatedPreset(config.Default()), store.WithContext(ctx))
		return err
	}
}

// atlas federatedAuthentication federationSettings connectedOrgsConfig list --federationSettingsId federationSettingsId [-o/--output output].
func ListBuilder() *cobra.Command {
	opts := new(ListOpts)
	cmd := &cobra.Command{
		Use:   "list",
		Short: "Describe a Connected Org Config.",
		Args:  cobra.NoArgs,
		Example: `  # List all connected org config with federationSettingsId 5d1113b25a115342acc2d1aa 
  atlas federatedAuthentication federationSettings connectedOrgsConfig list --federationSettingsId 5d1113b25a115342acc2d1aa 
`,
		PreRunE: func(cmd *cobra.Command, _ []string) error {
			return opts.PreRunE(
				opts.ValidateOrgID,
				opts.InitStore(cmd.Context()),
				opts.InitOutput(cmd.OutOrStdout(), listTemplate),
			)
		},
		RunE: func(_ *cobra.Command, _ []string) error {
			return opts.Run()
		},
	}

	cmd.Flags().StringVar(&opts.federationSettingsID, flag.FederationSettingsID, "", usage.FederationSettingsID)
	cmd.Flags().StringVar(&opts.OrgID, flag.OrgID, "", usage.OrgID)
	cmd.Flags().IntVar(&opts.PageNum, flag.Page, cli.DefaultPage, usage.Page)
	cmd.Flags().IntVar(&opts.ItemsPerPage, flag.Limit, cli.DefaultPageLimit, usage.Limit)

	cmd.Flags().StringVarP(&opts.Output, flag.Output, flag.OutputShort, "", usage.FormatOut)
	_ = cmd.RegisterFlagCompletionFunc(flag.Output, opts.AutoCompleteOutputFlag())

	_ = cmd.MarkFlagRequired(flag.FederationSettingsID)

	return cmd
}
