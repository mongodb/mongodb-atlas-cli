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

	"github.com/mongodb/mongodb-atlas-cli/internal/cli"
	"github.com/mongodb/mongodb-atlas-cli/internal/cli/require"
	"github.com/mongodb/mongodb-atlas-cli/internal/config"
	"github.com/mongodb/mongodb-atlas-cli/internal/flag"
	"github.com/mongodb/mongodb-atlas-cli/internal/store"
	"github.com/mongodb/mongodb-atlas-cli/internal/usage"
	"github.com/spf13/cobra"
	atlasv2 "go.mongodb.org/atlas-sdk/admin"
)

type CreateOpts struct {
	cli.GlobalOpts
	cli.OutputOpts
	store                    store.InterfaceEndpointCreator
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
	r := atlasv2.CreateGCPEndpointGroupRequestAsCreateEndpointRequest(
		&atlasv2.CreateGCPEndpointGroupRequest{
			EndpointGroupName: opts.privateEndpointGroupID,
			GcpProjectId:      opts.gcpProjectID,
			Endpoints:         opts.parseEndpoints(),
		},
	)
	return &r
}

// mongocli atlas privateEndpoint(s) gcp interface(s) create <endpointGroupId> --endpointServiceId endpointServiceId --gcpProjectId gcpProjectId --endpoint endpointName1@ipAddress1,...,endpointNameN@ipAddressN [--projectId projectId].
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
		Example: fmt.Sprintf(
			`  %s privateEndpoints gcp interfaces create endpoint-1 \
  --endpointServiceId 61eaca605af86411903de1dd \
  --gcpProjectId mcli-private-endpoints \
  --endpoint endpoint-0@10.142.0.2,endpoint-1@10.142.0.3,endpoint-2@10.142.0.4,endpoint-3@10.142.0.5,endpoint-4@10.142.0.6,endpoint-5@10.142.0.7`,
			cli.ExampleAtlasEntryPoint()),
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return opts.PreRunE(
				opts.validateEndpoints,
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
	cmd.Flags().StringSliceVar(&opts.Endpoints, flag.Endpoint, []string{}, usage.Endpoint)
	cmd.Flags().StringVar(&opts.ProjectID, flag.ProjectID, "", usage.ProjectID)
	cmd.Flags().StringVarP(&opts.Output, flag.Output, flag.OutputShort, "", usage.FormatOut)
	_ = cmd.RegisterFlagCompletionFunc(flag.Output, opts.AutoCompleteOutputFlag())
	_ = cmd.MarkFlagRequired(flag.EndpointServiceID)
	_ = cmd.MarkFlagRequired(flag.GCPProjectID)
	return cmd
}
