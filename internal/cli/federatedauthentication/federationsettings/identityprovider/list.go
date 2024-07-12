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

package identityprovider

import (
	"context"
	"fmt"

	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli/require"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/config"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/flag"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/store"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/usage"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/validate"
	"github.com/spf13/cobra"
	atlasv2 "go.mongodb.org/atlas-sdk/v20240530002/admin"
)

type ListOpts struct {
	cli.GlobalOpts
	cli.OutputOpts
	*cli.ListOpts
	store                store.IdentityProviderLister
	federationSettingsID string
	idpType              string
	protocol             string
}

const (
	oidc         = "OIDC"
	workforce    = "WORKFORCE"
	workload     = "WORKLOAD"
	saml         = "SAML"
	listTemplate = `ID	DISPLAY NAME	ISSUER URI	CLIENT ID	IDP TYPE{{range valueOrEmptySlice .Results}}
{{.Id}}	{{.DisplayName}}	{{.IssuerUri}}	{{.ClientId}}	{{.IdpType}}{{end}}
`
)

func (opts *ListOpts) initStore(ctx context.Context) func() error {
	return func() error {
		var err error
		opts.store, err = store.New(store.AuthenticatedPreset(config.Default()), store.WithContext(ctx))
		return err
	}
}

func (opts *ListOpts) Run() error {
	protocol := []string{opts.protocol}
	idpType := []string{opts.idpType}
	params := &atlasv2.ListIdentityProvidersApiParams{
		FederationSettingsId: opts.federationSettingsID,
		ItemsPerPage:         &opts.ItemsPerPage,
		PageNum:              &opts.PageNum,
		Protocol:             &protocol,
		IdpType:              &idpType,
	}
	r, err := opts.store.IdentityProviders(params)
	if err != nil {
		return err
	}

	return opts.Print(r)
}

func (opts *ListOpts) Validate() error {
	if err := validate.FlagInSlice(opts.idpType, flag.IdpType, []string{workforce, workload}); err != nil {
		return err
	}

	return validate.FlagInSlice(opts.protocol, flag.Protocol, []string{oidc, saml})
}

// atlas federatedAuthentication federationSettings identityProvider list --federationSettingsId federationSettingsId -[-idpType idpType] [--page page] [--itemsPerPage itemsPerPage] [--output output].
func ListBuilder() *cobra.Command {
	opts := &ListOpts{
		ListOpts: &cli.ListOpts{},
	}
	cmd := &cobra.Command{
		Use:   "list",
		Short: "List the identity providers from your federation settings.",
		Long:  fmt.Sprintf(usage.RequiredRole, "Org Owner"),
		Args:  require.NoArgs,
		Annotations: map[string]string{
			"identityProviderIdDesc": "ID of the identityProvider.",
			"output":                 listTemplate,
		},
		Example: `  # List the identity providers from your federation settings with federationSettingsId 5d1113b25a115342acc2d1aa and idpType WORKLOAD
	atlas federatedAuthentication federationSettings identityProvider list --federationSettingsId 5d1113b25a115342acc2d1aa --idpType WORKLOAD
`,
		PreRunE: func(cmd *cobra.Command, _ []string) error {
			return opts.PreRunE(
				opts.initStore(cmd.Context()),
				opts.InitOutput(cmd.OutOrStdout(), listTemplate),
				opts.Validate,
			)
		},
		RunE: func(_ *cobra.Command, _ []string) error {
			return opts.Run()
		},
	}

	cmd.Flags().IntVar(&opts.PageNum, flag.Page, cli.DefaultPage, usage.Page)
	cmd.Flags().IntVar(&opts.ItemsPerPage, flag.Limit, cli.DefaultPageLimit, usage.Limit)
	cmd.Flags().StringVar(&opts.federationSettingsID, flag.FederationSettingsID, "", usage.FederationSettingsID)
	cmd.Flags().StringVar(&opts.idpType, flag.IdpType, workforce, usage.IdpType)
	cmd.Flags().StringVarP(&opts.Output, flag.Output, flag.OutputShort, "", usage.FormatOut)
	cmd.Flags().StringVar(&opts.protocol, flag.Protocol, oidc, usage.Protocol)

	_ = cmd.MarkFlagRequired(flag.FederationSettingsID)

	return cmd
}
