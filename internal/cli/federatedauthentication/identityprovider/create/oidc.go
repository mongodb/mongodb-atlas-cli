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
	atlasv2 "go.mongodb.org/atlas-sdk/v20231115010/admin"
)

type OidcOpts struct {
	cli.GlobalOpts
	cli.OutputOpts
	cli.InputOpts
	AssociatedDomains    []string
	FederationSettingsID string
	Audience             string
	ClientID             string
	AuthorizationType    string
	Description          string
	DisplayName          string
	IdpType              string
	IssuerURI            string
	Protocol             string
	GroupsClaim          string
	UserClaim            string
	RequestedScopes      []string
	store                store.IdentityProviderCreator
}

const (
	user           = "USER"
	group          = "GROUP"
	oidc           = "OIDC"
	workflorce     = "WORKFORCE"
	workload       = "WORKLOAD"
	createTemplate = "Identity provider '{{.Id}}' created.\n"
)

var (
	validAuthTypeFlagValues = []string{group, user}
	validIdpTypeValues      = []string{workflorce, workload}
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
		FederationSettingsId: opts.FederationSettingsID,
		FederationOidcIdentityProviderUpdate: &atlasv2.FederationOidcIdentityProviderUpdate{
			AssociatedDomains: &opts.AssociatedDomains,
			Audience:          &opts.Audience,
			ClientId:          &opts.ClientID,
			AuthorizationType: &opts.AuthorizationType,
			Description:       &opts.Description,
			DisplayName:       &opts.DisplayName,
			IdpType:           &opts.IdpType,
			IssuerUri:         &opts.IssuerURI,
			Protocol:          &opts.Protocol,
			GroupsClaim:       &opts.GroupsClaim,
			RequestedScopes:   &opts.RequestedScopes,
			UserClaim:         &opts.UserClaim,
		},
	}
}

func (opts *OidcOpts) Validate() error {
	if err := validate.FlagInSlice(opts.AuthorizationType, flag.AuthorizationType, validAuthTypeFlagValues); err != nil {
		return err
	}

	return validate.FlagInSlice(opts.IdpType, flag.IdpType, validIdpTypeValues)
}

func (opts *OidcOpts) Run() error {
	user := opts.newIdentityProvider()

	r, err := opts.store.CreateIdentityProvider(user)
	if err != nil {
		return err
	}

	return opts.Print(r)
}

// atlas federatedAuthentication identityProvider oidc create <displayName> --federationSettingsId federationSettingsId --idpType idpType --audience audience --authorizationType authorizationType --clientId clientId --description description --groupsClaim groupsClaim --userClaim userClaim --issuerUri issuerUri [--associatedDomains associatedDomains] [--requestedScopes	 requestedScopes][-o/--output output]
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
		PreRunE: func(cmd *cobra.Command, args []string) error {
			opts.Protocol = oidc
			return opts.PreRunE(
				opts.InitStore(cmd.Context()),
				opts.InitOutput(cmd.OutOrStdout(), createTemplate),
				opts.InitInput(cmd.InOrStdin()),
				opts.Validate,
			)
		},
		RunE: func(_ *cobra.Command, args []string) error {
			opts.DisplayName = args[0]
			return opts.Run()
		},
	}

	cmd.Flags().StringVar(&opts.FederationSettingsID, flag.FederationSettingsID, "", usage.FederationSettingsID)
	cmd.Flags().StringVar(&opts.IdpType, flag.IdpType, group, usage.IdpType)
	cmd.Flags().StringVar(&opts.Audience, flag.Audience, "", usage.Audience)
	cmd.Flags().StringVar(&opts.AuthorizationType, flag.AuthorizationType, "", usage.AuthorizationType)
	cmd.Flags().StringVar(&opts.ClientID, flag.ClientID, "", usage.ClientID)
	cmd.Flags().StringVar(&opts.Description, flag.Description, "", usage.Description)
	cmd.Flags().StringVar(&opts.GroupsClaim, flag.GroupsClaim, "", usage.GroupsClaim)
	cmd.Flags().StringVar(&opts.UserClaim, flag.UserClaim, "", usage.UserClaim)
	cmd.Flags().StringVar(&opts.IssuerURI, flag.IssuerURI, "", usage.IssuerURI)
	cmd.Flags().StringSliceVar(&opts.AssociatedDomains, flag.AssociatedDomains, []string{}, usage.AssociatedDomains)
	cmd.Flags().StringSliceVar(&opts.RequestedScopes, flag.RequestedScopes, []string{}, usage.RequestedScopes)

	cmd.Flags().StringVarP(&opts.Output, flag.Output, flag.OutputShort, "", usage.FormatOut)
	_ = cmd.RegisterFlagCompletionFunc(flag.Output, opts.AutoCompleteOutputFlag())
	_ = cmd.MarkFlagRequired(flag.Username)
	_ = cmd.MarkFlagRequired(flag.FederationSettingsID)
	_ = cmd.MarkFlagRequired(flag.IdpType)
	_ = cmd.MarkFlagRequired(flag.Audience)
	_ = cmd.MarkFlagRequired(flag.AuthorizationType)
	_ = cmd.MarkFlagRequired(flag.ClientID)
	_ = cmd.MarkFlagRequired(flag.Description)
	_ = cmd.MarkFlagRequired(flag.DisplayName)
	_ = cmd.MarkFlagRequired(flag.GroupsClaim)
	_ = cmd.MarkFlagRequired(flag.UserClaim)
	_ = cmd.MarkFlagRequired(flag.IssuerURI)

	return cmd
}
