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
)

type DeleteOpts struct {
	cli.GlobalOpts
	*cli.DeleteOpts
	privateEndpointServiceID string
	store                    store.InterfaceEndpointDeleter
}

func (opts *DeleteOpts) initStore(ctx context.Context) func() error {
	return func() error {
		var err error
		opts.store, err = store.New(store.AuthenticatedPreset(config.Default()), store.WithContext(ctx))
		return err
	}
}

func (opts *DeleteOpts) Run() error {
	return opts.Delete(opts.store.DeleteInterfaceEndpoint, opts.ConfigProjectID(), provider, opts.privateEndpointServiceID)
}

// atlas privateEndpoint(s) azure interface(s) delete <endpointId> [--endpointServiceID endpointServiceID][--projectId projectId].
func DeleteBuilder() *cobra.Command {
	opts := &DeleteOpts{
		DeleteOpts: cli.NewDeleteOpts("Interface endpoint '%s' deleted\n", "Interface endpoint not deleted"),
	}
	cmd := &cobra.Command{
		Use:     "delete <privateEndpointResourceId>",
		Aliases: []string{"rm"},
		Short:   "Remove the specified Azure private endpoint interface and related service from your project.",
		Long:    fmt.Sprintf(usage.RequiredRole, "Project Owner"),
		Args:    require.ExactArgs(1),
		Annotations: map[string]string{
			"privateEndpointResourceIdDesc": "Unique string that identifies the Azure private endpoint interface in Azure.",
			"output":                        opts.SuccessMessage(),
		},
		Example: `  # Remove the Azure private endpoint interface with the ID /subscriptions/4e133d35-e734-4385-a565-c0945567ae346/resourceGroups/rg_95847a959b876e255dbb9b33_dfragd7w/providers/Microsoft.Network/privateEndpoints/cli-test in Azure from the project with the ID 5e2211c17a3e5a48f5497de3:
  atlas privateEndpoints azure interfaces delete /subscriptions/4e133d35-e734-4385-a565-c0945567ae346/resourceGroups/rg_95847a959b876e255dbb9b33_dfragd7w/providers/Microsoft.Network/privateEndpoints/cli-test --projectId 5e2211c17a3e5a48f5497de3`,
		PreRunE: func(cmd *cobra.Command, _ []string) error {
			return opts.PreRunE(opts.ValidateProjectID, opts.initStore(cmd.Context()))
		},
		RunE: func(_ *cobra.Command, args []string) error {
			opts.Entry = args[0]
			if err := opts.Prompt(); err != nil {
				return err
			}
			return opts.Run()
		},
	}

	cmd.Flags().BoolVar(&opts.Confirm, flag.Force, false, usage.Force)
	cmd.Flags().StringVar(&opts.privateEndpointServiceID, flag.EndpointServiceID, "", usage.EndpointServiceID)

	cmd.Flags().StringVar(&opts.ProjectID, flag.ProjectID, "", usage.ProjectID)

	_ = cmd.MarkFlagRequired(flag.PrivateEndpointID)

	return cmd
}
