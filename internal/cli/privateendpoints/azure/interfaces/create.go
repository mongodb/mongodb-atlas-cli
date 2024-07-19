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

package interfaces

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

type CreateOpts struct {
	cli.GlobalOpts
	cli.OutputOpts
	store                    store.InterfaceEndpointCreator
	privateEndpointID        string
	interfaceEndpointID      string
	privateEndpointIPAddress string
}

func (opts *CreateOpts) initStore(ctx context.Context) func() error {
	return func() error {
		var err error
		opts.store, err = store.New(store.AuthenticatedPreset(config.Default()), store.WithContext(ctx))
		return err
	}
}

var createTemplate = "Interface endpoint '{{.PrivateEndpointResourceId}}' created.\n"

func (opts *CreateOpts) Run() error {
	r, err := opts.store.CreateInterfaceEndpoint(opts.ConfigProjectID(), provider, opts.interfaceEndpointID, opts.newInterfaceEndpointConnection())
	if err != nil {
		return err
	}

	return opts.Print(r)
}

func (opts *CreateOpts) newInterfaceEndpointConnection() *atlasv2.CreateEndpointRequest {
	r := atlasv2.CreateEndpointRequest{
		Id:                       &opts.privateEndpointID,
		PrivateEndpointIPAddress: &opts.privateEndpointIPAddress,
	}
	return &r
}

// atlas privateEndpoint(s)|privateendpoint(s) azure interface(s) create <atlasPrivateEndpointId> [--privateEndpointId privateEndpointID][--privateEndpointIPAddress privateEndpointIPAddress][--projectId projectId].
func CreateBuilder() *cobra.Command {
	opts := &CreateOpts{}
	cmd := &cobra.Command{
		Use:     "create <endpointServiceId>",
		Aliases: []string{"add"},
		Short:   "Create a new interface for the specified Azure private endpoint.",
		Long: `To learn more about how to set up private endpoints with the Atlas CLI, see the tutorial on the Atlas CLI tab here: https://www.mongodb.com/docs/atlas/security-cluster-private-endpoint/.

` + fmt.Sprintf(usage.RequiredRole, "Project Owner"),
		Args: require.ExactArgs(1),
		Annotations: map[string]string{
			"endpointServiceIdDesc": "Unique 24-character alphanumeric string that identifies the private endpoint in Atlas.",
			"output":                createTemplate,
		},
		Example: `  # Create a new interface for an Azure private endpoint with the ID 5f4fc14da2b47835a58c63a2 in Atlas and the ID /subscriptions/4e133d35-e734-4385-a565-c0945567ae346/resourceGroups/rg_95847a959b876e255dbb9b33_dfragd7w/providers/Microsoft.Network/privateEndpoints/test-endpoint in Azure for the project with the ID 5e2211c17a3e5a48f5497de3:
  atlas privateEndpoints azure interfaces create 5f4fc14da2b47835a58c63a2 --privateEndpointId /subscriptions/4e133d35-e734-4385-a565-c0945567ae346/resourceGroups/rg_95847a959b876e255dbb9b33_dfragd7w/providers/Microsoft.Network/privateEndpoints/test-endpoint --projectId 5e2211c17a3e5a48f5497de3 --privateEndpointIpAddress 192.0.2.5
  --output json`,
		PreRunE: func(cmd *cobra.Command, _ []string) error {
			return opts.PreRunE(
				opts.ValidateProjectID,
				opts.initStore(cmd.Context()),
				opts.InitOutput(cmd.OutOrStdout(), createTemplate),
			)
		},
		RunE: func(_ *cobra.Command, args []string) error {
			opts.interfaceEndpointID = args[0]
			return opts.Run()
		},
	}
	cmd.Flags().StringVar(&opts.privateEndpointID, flag.PrivateEndpointID, "", usage.PrivateEndpointIDAzure)
	cmd.Flags().StringVar(&opts.privateEndpointIPAddress, flag.PrivateEndpointIPAddress, "", usage.PrivateEndpointIPAddressAzure)

	cmd.Flags().StringVar(&opts.ProjectID, flag.ProjectID, "", usage.ProjectID)
	cmd.Flags().StringVarP(&opts.Output, flag.Output, flag.OutputShort, "", usage.FormatOut)
	_ = cmd.RegisterFlagCompletionFunc(flag.Output, opts.AutoCompleteOutputFlag())

	_ = cmd.MarkFlagRequired(flag.PrivateEndpointID)
	_ = cmd.MarkFlagRequired(flag.PrivateEndpointIPAddress)

	return cmd
}
