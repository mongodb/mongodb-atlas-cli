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
)

type RevokeOpts struct {
	cli.GlobalOpts
	cli.OutputOpts
	*cli.DeleteOpts
	store                store.IdentityProviderJwkRevoker
	FederationSettingsID string
}

func (opts *RevokeOpts) initStore(ctx context.Context) func() error {
	return func() error {
		var err error
		opts.store, err = store.New(store.AuthenticatedPreset(config.Default()), store.WithContext(ctx))
		return err
	}
}

func (opts *RevokeOpts) Run() error {
	return opts.Delete(opts.store.RevokeJwksFromIdentityProvider, opts.FederationSettingsID)
}

// atlas federatedAuthentication federationSettings identityProvider revokeJwk <identityProviderId> --federationSettingsId federationSettingsId [--output output].
func RevokeBuilder() *cobra.Command {
	opts := &RevokeOpts{
		DeleteOpts: cli.NewDeleteOpts("Identity Provider %s JWK token revoked.\n", "Identity Provider %s JWK token not revoked.\n"),
	}
	cmd := &cobra.Command{
		Use:   "revokeJwk <identityProviderId>",
		Short: "Revoke the JWK token from the specified identity provider from your federation settings.",
		Long:  fmt.Sprintf(usage.RequiredRole, "Org Owner"),
		Args:  require.ExactArgs(1),
		Annotations: map[string]string{
			"identityProviderIdDesc": "ID of the identityProvider.",
		},
		Example: `  # Revoke the Jwk from the identity provider with ID aa2223b25a115342acc1f108 and federation settings with federationSettingsId 5d1113b25a115342acc2d1aa.
	atlas federatedAuthentication federationSettings identityProvider revokeJwk aa2223b25a115342acc1f108 --federationSettingsId 5d1113b25a115342acc2d1aa
`,
		PreRunE: func(cmd *cobra.Command, args []string) error {
			opts.Entry = args[0]
			return opts.PreRunE(
				opts.initStore(cmd.Context()),
			)
		},
		RunE: func(_ *cobra.Command, _ []string) error {
			opts.Confirm = true
			return opts.Run()
		},
	}

	cmd.Flags().StringVar(&opts.FederationSettingsID, flag.FederationSettingsID, "", usage.FederationSettingsID)
	cmd.Flags().StringVarP(&opts.Output, flag.Output, flag.OutputShort, "", usage.FormatOut)

	_ = cmd.MarkFlagRequired(flag.FederationSettingsID)

	return cmd
}
