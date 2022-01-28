// Copyright 2022 MongoDB Inc
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
	"strings"

	atlas "go.mongodb.org/atlas/mongodbatlas"

	"github.com/mongodb/mongocli/internal/cli"
	"github.com/mongodb/mongocli/internal/cli/require"
	"github.com/mongodb/mongocli/internal/config"
	"github.com/mongodb/mongocli/internal/flag"
	"github.com/mongodb/mongocli/internal/store"
	"github.com/mongodb/mongocli/internal/usage"
	"github.com/spf13/cobra"
)

type CreateOpts struct {
	cli.GlobalOpts
	cli.OutputOpts
	store                    store.InterfaceEndpointCreator
	privateEndpointServiceID string
	privateEndpointGroupID   string
	gcpProjectID             string
	endpoints                []string
}

func (opts *CreateOpts) initStore(ctx context.Context) func() error {
	return func() error {
		var err error
		opts.store, err = store.New(store.AuthenticatedPreset(config.Default()), store.WithContext(ctx))
		return err
	}
}

func (opts *CreateOpts) parseEndpoints() []*atlas.GCPEndpoint {
	endpoints := make([]*atlas.GCPEndpoint, len(opts.endpoints))
	for i, endpoint := range opts.endpoints {
		s := strings.Split(endpoint, "@")
		endpoints[i] = &atlas.GCPEndpoint{
			EndpointName: s[0],
			IPAddress:    s[1],
		}
	}
	return endpoints
}

var createTemplate = "Interface endpoint '{{.EndpointGroupName}}' created.\n"

func (opts *CreateOpts) Run() error {
	r, err := opts.store.CreateInterfaceEndpoint(opts.ConfigProjectID(), provider, opts.privateEndpointServiceID, opts.newInterfaceEndpointConnection())
	if err != nil {
		return err
	}

	return opts.Print(r)
}

func (opts *CreateOpts) newInterfaceEndpointConnection() *atlas.InterfaceEndpointConnection {
	return &atlas.InterfaceEndpointConnection{
		EndpointGroupName: opts.privateEndpointGroupID,
		GCPProjectID:      opts.gcpProjectID,
		Endpoints:         opts.parseEndpoints(),
	}
}

// TODO: Better way to indicate --endpoints list?
// mongocli atlas privateEndpoint(s) gcp interface(s) create <endpointGroupId> --endpointServiceId endpointServiceId --gcpProjectId gcpProjectId --endpoints <list> [--projectId projectId].
func CreateBuilder() *cobra.Command {
	opts := &CreateOpts{}
	cmd := &cobra.Command{
		Use:     "create <endpointGroupId>",
		Aliases: []string{"add"},
		Short:   "Create a GCP private endpoint interface.",
		Args:    require.ExactArgs(1),
		Annotations: map[string]string{
			"args":                "endpointGroupId",
			"requiredArgs":        "endpointGroupId",
			"endpointGroupIdDesc": "Unique identifier for the endpoint group.",
		},
		Example: `$ mongocli atlas privateEndpoints gcp interfaces create endpoint-1 --endpointServiceId 61eaca605af86411903de1dd --gcpProjectId mcli-private-endpoints --endpoints endpoint-1-0@10.142.0.2,endpoint-1-1@10.142.0.3,endpoint-1-2@10.142.0.4,endpoint-1-3@10.142.0.5,endpoint-1-4@10.142.0.6,endpoint-1-5@10.142.0.7`,
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return opts.PreRunE(
				opts.ValidateProjectID,
				opts.initStore(cmd.Context()),
				opts.InitOutput(cmd.OutOrStdout(), createTemplate),
			)
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			opts.privateEndpointGroupID = args[0]
			return opts.Run()
		},
	}
	cmd.Flags().StringVar(&opts.privateEndpointServiceID, flag.EndpointServiceID, "", usage.EndpointServiceID)
	cmd.Flags().StringVar(&opts.gcpProjectID, flag.GCPProjectID, "", usage.GCPProjectID)
	cmd.Flags().StringSliceVar(&opts.endpoints, flag.Endpoints, []string{}, usage.LinkTokenAccessListCIDREntries)
	cmd.Flags().StringVar(&opts.ProjectID, flag.ProjectID, "", usage.ProjectID)
	cmd.Flags().StringVarP(&opts.Output, flag.Output, flag.OutputShort, "", usage.FormatOut)
	_ = cmd.MarkFlagRequired(flag.EndpointServiceID)
	_ = cmd.MarkFlagRequired(flag.GCPProjectID)
	return cmd
}
