// Copyright 2021 MongoDB Inc
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

package federatedauthentication

import (
	"context"
	"fmt"

	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/config"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/flag"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/store"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/usage"
	"github.com/spf13/cobra"
	atlasv2 "go.mongodb.org/atlas-sdk/v20231115009/admin"
)

type CreateOpts struct {
	cli.GlobalOpts
	cli.OutputOpts
	cli.InputOpts
	associatedDomains []string
	Audience          string
	ClientId          string
	AuthorizationType string
	Description       string
	DisplayName       string
	IdpType           string
	IssuerUri         string
	Protocol          string
	scopes            []string
	GroupsClaim       string
	store             store.IdentityProviderCreator
}

const (
	user             = "USER"
	role             = "ROLE"
	group            = "GROUP"
	X509TypeManaged  = "MANAGED"
	X509TypeCustomer = "CUSTOMER"
	none             = "NONE"
	createTemplate   = "Database user '{{.Username}}' successfully created.\n"
)

func (opts *CreateOpts) InitStore(ctx context.Context) func() error {
	return func() error {
		if opts.store != nil {
			return nil
		}

		var err error
		opts.store, err = store.New(store.AuthenticatedPreset(config.Default()), store.WithContext(ctx))
		return err
	}
}
func (opts *CreateOpts) newIdentityProvider() *atlasv2.CreateIdentityProviderApiParams {
	opts.associatedDomains = []string{"iam-test-domain-dev.com"}
	opts.Audience = "audience3"
	opts.Description = "CLI TEST PROvider"
	opts.DisplayName = "CLI test provider"
	// opts.IdpType = "WORKFORCE"
	opts.GroupsClaim = "groups"
	opts.Protocol = "OIDC"
	opts.IssuerUri = "https://accounts.google.com"

	return &atlasv2.CreateIdentityProviderApiParams{
		FederationSettingsId: "6531abc7cbe54754fd77f068",
		FederationOidcIdentityProviderUpdate: &atlasv2.FederationOidcIdentityProviderUpdate{
			AssociatedDomains: &opts.associatedDomains,
			Audience:          &opts.Audience,
			ClientId:          &opts.ClientId,
			AuthorizationType: &opts.AuthorizationType,
			Description:       &opts.Description,
			DisplayName:       &opts.DisplayName,
			IdpType:           &opts.IdpType,
			IssuerUri:         &opts.IssuerUri,
			Protocol:          &opts.Protocol,
			GroupsClaim:       &opts.GroupsClaim,
		},
	}
}

func (opts *CreateOpts) Run() error {
	user := opts.newIdentityProvider()

	r, err := opts.store.CreateIdentityProvider(user)
	if err != nil {
		return err
	}

	return opts.Print(r)
}

// CreateBuilder
// mongocli atlas dbuser(s) create
//
//	--username username --password password
//	--role roleName@dbName
//	--scope resourceName@resourceType
//	[--projectId projectId]
//	[--x509Type NONE|MANAGED|CUSTOMER]
//	[--awsIAMType NONE|ROLE|USER]
//	[--ldapType NONE|USER|GROUP]
func CreateBuilder() *cobra.Command {
	opts := &CreateOpts{}
	cmd := &cobra.Command{
		Use:   "create [builtInRole]...",
		Short: "Create a database user for your project.",
		Long: `If you set --ldapType, --x509Type, and --awsIAMType to NONE, Atlas authenticates this user through SCRAM-SHA. To learn more, see https://www.mongodb.com/docs/manual/core/security-scram/.

` + fmt.Sprintf(usage.RequiredRole, "Project Owner"),
		Example: `  # Create an Atlas database admin user named myAdmin for the project with ID 5e2211c17a3e5a48f5497de3:
  atlas dbusers create atlasAdmin --username myAdmin  --projectId 5e2211c17a3e5a48f5497de3

  # Create a database user named myUser with read/write access to any database for the project with ID 5e2211c17a3e5a48f5497de3:
  atlas dbusers create readWriteAnyDatabase --username myUser --projectId 5e2211c17a3e5a48f5497de3

  # Create a database user named myUser with multiple roles for the project with ID 5e2211c17a3e5a48f5497de3:
  atlas dbusers create --username myUser --role clusterMonitor,backup --projectId 5e2211c17a3e5a48f5497de3

  # Create a database user named myUser with multiple scopes for the project with ID 5e2211c17a3e5a48f5497de3:
  atlas dbusers create --username myUser --role clusterMonitor --scope <REPLICA-SET ID>,<storeName> --projectId 5e2211c17a3e5a48f5497de3`,
		Args: cobra.OnlyValidArgs,
		Annotations: map[string]string{
			"builtInRoleDesc": "Atlas built-in role that you want to assign to the user.",
			"output":          createTemplate,
		},
		ValidArgs: []string{"atlasAdmin", "readWriteAnyDatabase", "readAnyDatabase", "clusterMonitor", "backup", "dbAdminAnyDatabase", "enableSharding"},
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return opts.PreRunE(
				opts.ValidateProjectID,
				opts.InitStore(cmd.Context()),
				opts.InitOutput(cmd.OutOrStdout(), createTemplate),
				opts.InitInput(cmd.InOrStdin()),
			)
		},
		RunE: func(_ *cobra.Command, _ []string) error {

			return opts.Run()
		},
	}

	cmd.Flags().StringVar(&opts.ProjectID, flag.ProjectID, "", usage.ProjectID)
	cmd.Flags().StringVarP(&opts.Output, flag.Output, flag.OutputShort, "", usage.FormatOut)
	_ = cmd.RegisterFlagCompletionFunc(flag.Output, opts.AutoCompleteOutputFlag())

	_ = cmd.MarkFlagRequired(flag.Username)

	return cmd
}
