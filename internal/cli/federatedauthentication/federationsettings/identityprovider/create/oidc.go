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

package create

import (
	"context"

	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/config"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/flag"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/store"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/usage"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/validate"
	"github.com/spf13/cobra"
	atlasv2 "go.mongodb.org/atlas-sdk/v20240530002/admin"
)

type OidcOpts struct {
	cli.GlobalOpts
	cli.OutputOpts
	cli.InputOpts
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
	store                store.IdentityProviderCreator
}

const (
	user           = "USER"
	group          = "GROUP"
	oidc           = "OIDC"
	workforce      = "WORKFORCE"
	workload       = "WORKLOAD"
	createTemplate = "Identity provider '{{.Id}}' created.\n"
)

var (
	validAuthTypeFlagValues = []string{group, user}
	validIdpTypeValues      = []string{workforce, workload}
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

func (opts *OidcOpts) newIdentityProvider() *atlasv2.CreateIdentityProviderApiParams {
	return &atlasv2.CreateIdentityProviderApiParams{
		FederationSettingsId: opts.federationSettingsID,
		FederationOidcIdentityProviderUpdate: &atlasv2.FederationOidcIdentityProviderUpdate{
			AssociatedDomains: &opts.associatedDomains,
			Audience:          &opts.audience,
			ClientId:          &opts.clientID,
			AuthorizationType: &opts.authorizationType,
			Description:       &opts.description,
			DisplayName:       &opts.displayName,
			IdpType:           &opts.idpType,
			IssuerUri:         &opts.issuerURI,
			Protocol:          &opts.protocol,
			GroupsClaim:       &opts.groupsClaim,
			RequestedScopes:   &opts.requestedScopes,
			UserClaim:         &opts.userClaim,
		},
	}
}

func (opts *OidcOpts) Validate() error {
	if err := validate.FlagInSlice(opts.authorizationType, flag.AuthorizationType, validAuthTypeFlagValues); err != nil {
		return err
	}

	return validate.FlagInSlice(opts.idpType, flag.IdpType, validIdpTypeValues)
}

func (opts *OidcOpts) Run() error {
	user := opts.newIdentityProvider()

	r, err := opts.store.CreateIdentityProvider(user)
	if err != nil {
		return err
	}

	return opts.Print(r)
}

// atlas federatedAuthentication federationSettings identityProvider oidc create <displayName> --federationSettingsId federationSettingsId --idpType idpType --audience audience --authorizationType authorizationType --clientId clientId --description description --groupsClaim groupsClaim --userClaim userClaim --issuerUri issuerUri [--associatedDomain associatedDomains] [--requestedScope requestedScopes][-o/--output output].
func OIDCBuilder() *cobra.Command {
	opts := &OidcOpts{}
	cmd := &cobra.Command{
		Use:   "oidc [displayName]",
		Short: "Create an OIDC identity provider.",
		Args:  cobra.ExactArgs(1),
		Annotations: map[string]string{
			"displayNameDesc": "The Identity Provider display name.",
			"output":          createTemplate,
		},
		Example: `  # Create an identity provider with name IDPName and from your federation settings with federationSettingsId 5d1113b25a115342acc2d1aa.
		atlas federatedAuthentication federationSettings identityProvider create oidc IDPName --audience "audience" --authorizationType "GROUP" --clientId clientId --desc "IDPName test" --federationSettingsId "5d1113b25a115342acc2d1aa" --groupsClaim "groups" --idpType "WORKLOAD" --issuerUri uri" --userClaim "user"  --associatedDomain "domain"
	`,
		PreRunE: func(cmd *cobra.Command, _ []string) error {
			opts.protocol = oidc
			return opts.PreRunE(
				opts.InitStore(cmd.Context()),
				opts.InitOutput(cmd.OutOrStdout(), createTemplate),
				opts.InitInput(cmd.InOrStdin()),
				opts.Validate,
			)
		},
		RunE: func(_ *cobra.Command, args []string) error {
			opts.displayName = args[0]
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
	_ = cmd.MarkFlagRequired(flag.Username)
	_ = cmd.MarkFlagRequired(flag.FederationSettingsID)
	_ = cmd.MarkFlagRequired(flag.IdpType)
	_ = cmd.MarkFlagRequired(flag.Audience)
	_ = cmd.MarkFlagRequired(flag.AuthorizationType)
	_ = cmd.MarkFlagRequired(flag.Description)
	_ = cmd.MarkFlagRequired(flag.DisplayName)
	_ = cmd.MarkFlagRequired(flag.GroupsClaim)
	_ = cmd.MarkFlagRequired(flag.UserClaim)
	_ = cmd.MarkFlagRequired(flag.IssuerURI)

	return cmd
}
