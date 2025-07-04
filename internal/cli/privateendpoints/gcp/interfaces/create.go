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
	"fmt"
	"strings"

	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli/require"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/config"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/flag"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/pointer"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/store"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/usage"
	"github.com/spf13/cobra"
	atlasv2 "go.mongodb.org/atlas-sdk/v20250312005/admin"
)

//go:generate go tool go.uber.org/mock/mockgen -typed -destination=create_mock_test.go -package=interfaces . InterfaceEndpointCreator

type InterfaceEndpointCreator interface {
	CreateInterfaceEndpoint(string, string, string, *atlasv2.CreateEndpointRequest) (*atlasv2.PrivateLinkEndpoint, error)
}

type CreateOpts struct {
	cli.ProjectOpts
	cli.OutputOpts
	store                    InterfaceEndpointCreator
	privateEndpointServiceID string
	privateEndpointGroupID   string
	gcpProjectID             string
	Endpoints                []string
}

func (opts *CreateOpts) initStore(ctx context.Context) func() error {
	return func() error {
		var err error
		opts.store, err = store.New(store.AuthenticatedPreset(config.Default()), store.WithContext(ctx))
		return err
	}
}

func (opts *CreateOpts) validateEndpoints() error {
	for _, endpoint := range opts.Endpoints {
		index := strings.Index(endpoint, "@")
		if index < 1 || index >= len(endpoint)-1 {
			return fmt.Errorf("invalid endpoint: %s\nRequired format is: <endpointName>@<ipAddress>, eg: foo@127.0.0.1", endpoint)
		}
	}
	return nil
}

func (opts *CreateOpts) parseEndpoints() []atlasv2.CreateGCPForwardingRuleRequest {
	endpoints := make([]atlasv2.CreateGCPForwardingRuleRequest, len(opts.Endpoints))
	for i, endpoint := range opts.Endpoints {
		s := strings.Split(endpoint, "@")
		endpoints[i] = atlasv2.CreateGCPForwardingRuleRequest{
			EndpointName: &s[0],
			IpAddress:    &s[1],
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

func (opts *CreateOpts) newInterfaceEndpointConnection() *atlasv2.CreateEndpointRequest {
	return &atlasv2.CreateEndpointRequest{
		EndpointGroupName: &opts.privateEndpointGroupID,
		GcpProjectId:      &opts.gcpProjectID,
		Endpoints:         pointer.Get(opts.parseEndpoints()),
	}
}

// atlas privateEndpoint(s) gcp interface(s) create <endpointGroupId> --endpointServiceId endpointServiceId --gcpProjectId gcpProjectId --endpoint endpointName1@ipAddress1,...,endpointNameN@ipAddressN [--projectId projectId].
func CreateBuilder() *cobra.Command {
	opts := &CreateOpts{}
	cmd := &cobra.Command{
		Use:     "create <endpointGroupId>",
		Aliases: []string{"add"},
		Short:   "Create a GCP private endpoint interface.",
		Long:    fmt.Sprintf(usage.RequiredRole, "Project Owner"),
		Args:    require.ExactArgs(1),
		Annotations: map[string]string{
			"endpointGroupIdDesc": "Unique identifier for the endpoint group.",
			"output":              createTemplate,
		},
		Example: `  atlas privateEndpoints gcp interfaces create endpoint-1 \
  --endpointServiceId 61eaca605af86411903de1dd \
  --gcpProjectId mcli-private-endpoints \
  --endpoint endpoint-0@10.142.0.2,endpoint-1@10.142.0.3,endpoint-2@10.142.0.4,endpoint-3@10.142.0.5,endpoint-4@10.142.0.6,endpoint-5@10.142.0.7`,
		PreRunE: func(cmd *cobra.Command, _ []string) error {
			return opts.PreRunE(
				opts.validateEndpoints,
				opts.ValidateProjectID,
				opts.initStore(cmd.Context()),
				opts.InitOutput(cmd.OutOrStdout(), createTemplate),
			)
		},
		RunE: func(_ *cobra.Command, args []string) error {
			opts.privateEndpointGroupID = args[0]
			return opts.Run()
		},
	}
	cmd.Flags().StringVar(&opts.privateEndpointServiceID, flag.EndpointServiceID, "", usage.EndpointServiceID)
	cmd.Flags().StringVar(&opts.gcpProjectID, flag.GCPProjectID, "", usage.GCPProjectID)
	cmd.Flags().StringSliceVar(&opts.Endpoints, flag.Endpoint, []string{}, usage.Endpoint)
	opts.AddProjectOptsFlags(cmd)
	opts.AddOutputOptFlags(cmd)

	_ = cmd.MarkFlagRequired(flag.EndpointServiceID)
	_ = cmd.MarkFlagRequired(flag.GCPProjectID)
	return cmd
}
