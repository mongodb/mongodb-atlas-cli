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

package update

import (
	"context"

	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/config"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/flag"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/store"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/usage"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/validate"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	atlasv2 "go.mongodb.org/atlas-sdk/v20240530002/admin"
)

type OidcOpts struct {
	cli.GlobalOpts
	cli.OutputOpts
	cli.InputOpts
	identityProviderID   string
	associatedDomains    []string
	federationSettingsID string
	audience             string
	clientID             string
	authorizationType    string
	description          string
	displayName          string
	idpType              string
	issuerURI            string
	protocol             string
	groupsClaim          string
	userClaim            string
	requestedScopes      []string
	store                store.IdentityProviderUpdater
}

const (
	user           = "USER"
	group          = "GROUP"
	oidc           = "OIDC"
	workflorce     = "WORKFORCE"
	workload       = "WORKLOAD"
	updateTemplate = "Identity provider '{{.Id}}' updated.\n"
)

var (
	validAuthTypeFlagValues   = []string{group, user}
	validIdpTypeValues        = []string{workflorce, workload}
	workloadInvalidValidFlags = []string{flag.AssociatedDomain, flag.ClientID, flag.RequestedScope}
)

func (opts *OidcOpts) InitStore(ctx context.Context) func() error {
	return func() error {
		if opts.store != nil {
			return nil
		}

		var err error
		opts.store, err = store.New(store.AuthenticatedPreset(config.Default()), store.WithContext(ctx))
		return err
	}
}

func (opts *OidcOpts) newIdentityProvider() *atlasv2.UpdateIdentityProviderApiParams {
	params := &atlasv2.UpdateIdentityProviderApiParams{
		FederationSettingsId: opts.federationSettingsID,
		IdentityProviderId:   opts.identityProviderID,
		FederationIdentityProviderUpdate: &atlasv2.FederationIdentityProviderUpdate{
			Protocol: &opts.protocol,
			IdpType:  &opts.idpType,
		},
	}

	if len(opts.associatedDomains) > 0 {
		params.FederationIdentityProviderUpdate.AssociatedDomains = &opts.associatedDomains
	}

	if opts.audience != "" {
		params.FederationIdentityProviderUpdate.Audience = &opts.audience
	}

	if opts.clientID != "" {
		params.FederationIdentityProviderUpdate.ClientId = &opts.clientID
	}

	if opts.authorizationType != "" {
		params.FederationIdentityProviderUpdate.AuthorizationType = &opts.authorizationType
	}

	if opts.description != "" {
		params.FederationIdentityProviderUpdate.Description = &opts.description
	}

	if opts.displayName != "" {
		params.FederationIdentityProviderUpdate.DisplayName = &opts.displayName
	}

	if opts.issuerURI != "" {
		params.FederationIdentityProviderUpdate.IssuerUri = &opts.issuerURI
	}

	if opts.groupsClaim != "" {
		params.FederationIdentityProviderUpdate.GroupsClaim = &opts.groupsClaim
	}

	if len(opts.requestedScopes) > 0 {
		params.FederationIdentityProviderUpdate.RequestedScopes = &opts.requestedScopes
	}

	if opts.userClaim != "" {
		params.FederationIdentityProviderUpdate.UserClaim = &opts.userClaim
	}

	if len(opts.associatedDomains) > 0 {
		params.FederationIdentityProviderUpdate.AssociatedDomains = &opts.associatedDomains
	}

	return params
}

func (opts *OidcOpts) Validate(flagSet *pflag.FlagSet) error {
	var flags []string
	flagSet.Visit(func(f *pflag.Flag) {
		flags = append(flags, f.Name)
	})

	if opts.idpType == workload {
		for _, f := range flags {
			if err := validate.ConditionalFlagNotInSlice(flag.IdpType, opts.idpType, f, workloadInvalidValidFlags); err != nil {
				return err
			}
		}
	}

	if opts.authorizationType != "" {
		if err := validate.FlagInSlice(opts.authorizationType, flag.AuthorizationType, validAuthTypeFlagValues); err != nil {
			return err
		}
	}

	if opts.idpType != "" {
		if err := validate.FlagInSlice(opts.idpType, flag.IdpType, validIdpTypeValues); err != nil {
			return err
		}
	}
	return nil
}

func (opts *OidcOpts) Run() error {
	provider := opts.newIdentityProvider()

	r, err := opts.store.UpdateIdentityProvider(provider)
	if err != nil {
		return err
	}

	return opts.Print(r)
}

// atlas federatedAuthentication federationSettings identityProvider update oidc <identityProviderId> --federationSettingsId federationSettingsId --idpType idpType [--audience audience] [--authorizationType authorizationType] [--clientId clientId] [--description description] [--displayName displayName] [--groupsClaim groupsClaim] [--userClaim userClaim] [--issuerUri issuerUri] [--associatedDomain associatedDomains] [--requestedScope requestedScopes][-o/--output output].
func OIDCBuilder() *cobra.Command {
	opts := &OidcOpts{
		protocol: oidc,
	}
	cmd := &cobra.Command{
		Use:   "oidc [identityProviderId]",
		Short: "Update an OIDC identity provider.",
		Args:  cobra.ExactArgs(1),
		Annotations: map[string]string{
			"identityProviderIdDesc": "The Identity Provider ID.",
			"output":                 updateTemplate,
		},
		Example: `  # Update the audience of the identity provider with ID aa2223b25a115342acc1f108 and from your federation settings with federationSettingsId 5d1113b25a115342acc2d1aa with IdpType WORKFORCE
			atlas federatedAuthentication federationSettings identityProvider update aa2223b25a115342acc1f108 --federationSettingsId 5d1113b25a115342acc2d1aa --idpType WORKFORCE --audience newAudience
		`,
		PreRunE: func(cmd *cobra.Command, _ []string) error {
			opts.protocol = oidc
			flags := cmd.Flags()
			if err := opts.Validate(flags); err != nil {
				return err
			}
			return opts.PreRunE(
				opts.InitStore(cmd.Context()),
				opts.InitOutput(cmd.OutOrStdout(), updateTemplate),
				opts.InitInput(cmd.InOrStdin()),
			)
		},
		RunE: func(_ *cobra.Command, args []string) error {
			opts.identityProviderID = args[0]
			return opts.Run()
		},
	}

	cmd.Flags().StringVar(&opts.federationSettingsID, flag.FederationSettingsID, "", usage.FederationSettingsID)
	cmd.Flags().StringVar(&opts.idpType, flag.IdpType, "", usage.IdpType)
	cmd.Flags().StringVar(&opts.audience, flag.Audience, "", usage.Audience)
	cmd.Flags().StringVar(&opts.authorizationType, flag.AuthorizationType, "", usage.AuthorizationType)
	cmd.Flags().StringVar(&opts.clientID, flag.ClientID, "", usage.ClientID)
	cmd.Flags().StringVar(&opts.description, flag.Description, "", usage.Description)
	cmd.Flags().StringVar(&opts.groupsClaim, flag.GroupsClaim, "", usage.GroupsClaim)
	cmd.Flags().StringVar(&opts.userClaim, flag.UserClaim, "", usage.UserClaim)
	cmd.Flags().StringVar(&opts.issuerURI, flag.IssuerURI, "", usage.IssuerURI)
	cmd.Flags().StringSliceVar(&opts.associatedDomains, flag.AssociatedDomain, []string{}, usage.AssociatedDomains)
	cmd.Flags().StringSliceVar(&opts.requestedScopes, flag.RequestedScope, []string{}, usage.RequestedScopes)

	cmd.Flags().StringVarP(&opts.Output, flag.Output, flag.OutputShort, "", usage.FormatOut)
	_ = cmd.RegisterFlagCompletionFunc(flag.Output, opts.AutoCompleteOutputFlag())
	_ = cmd.MarkFlagRequired(flag.FederationSettingsID)
	_ = cmd.MarkFlagRequired(flag.IdpType)

	return cmd
}
