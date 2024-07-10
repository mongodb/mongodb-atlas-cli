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
	"github.com/spf13/cobra"
	atlasv2 "go.mongodb.org/atlas-sdk/v20240530002/admin"
)

type DescribeOpts struct {
	cli.GlobalOpts
	cli.OutputOpts
	store                store.IdentityProviderDescriber
	FederationSettingsID string
	IdentityProviderID   string
}

const describeTemplate = `ID	DISPLAY NAME	ISSUER URI	CLIENT ID	IDP TYPE
{{.Id}}	{{.DisplayName}}	{{.IssuerUri}}	{{.ClientId}}	{{.IdpType}}
`

func (opts *DescribeOpts) initStore(ctx context.Context) func() error {
	return func() error {
		var err error
		opts.store, err = store.New(store.AuthenticatedPreset(config.Default()), store.WithContext(ctx))
		return err
	}
}

func (opts *DescribeOpts) Run() error {
	params := &atlasv2.GetIdentityProviderApiParams{
		FederationSettingsId: opts.FederationSettingsID,
		IdentityProviderId:   opts.IdentityProviderID,
	}

	r, err := opts.store.IdentityProvider(params)
	if err != nil {
		return err
	}

	return opts.Print(r)
}

// atlas federatedAuthentication federationSettings identityProvider describe <identityProviderId> --federationSettingsId federationSettingsId [--output output].
func DescribeBuilder() *cobra.Command {
	opts := &DescribeOpts{}
	cmd := &cobra.Command{
		Use:   "describe <identityProviderId>",
		Short: "Describe the specified identity provider from your federation settings.",
		Long:  fmt.Sprintf(usage.RequiredRole, "Org Owner"),
		Args:  require.ExactArgs(1),
		Annotations: map[string]string{
			"identityProviderIdDesc": "ID of the identityProvider.",
			"output":                 describeTemplate,
		},
		Example: `  # Describe the identity provider with ID aa2223b25a115342acc1f108 with federationSettingsId 5d1113b25a115342acc2d1aa.
	atlas federatedAuthentication federationSettings identityProvider describe aa2223b25a115342acc1f108 --federationSettingsId 5d1113b25a115342acc2d1aa
`,
		PreRunE: func(cmd *cobra.Command, args []string) error {
			opts.IdentityProviderID = args[0]
			return opts.PreRunE(
				opts.initStore(cmd.Context()),
				opts.InitOutput(cmd.OutOrStdout(), describeTemplate),
			)
		},
		RunE: func(_ *cobra.Command, _ []string) error {
			return opts.Run()
		},
	}

	cmd.Flags().StringVar(&opts.FederationSettingsID, flag.FederationSettingsID, "", usage.FederationSettingsID)
	cmd.Flags().StringVarP(&opts.Output, flag.Output, flag.OutputShort, "", usage.FormatOut)

	_ = cmd.MarkFlagRequired(flag.FederationSettingsID)

	return cmd
}
