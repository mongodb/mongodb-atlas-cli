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

	"github.com/mongodb/atlas-cli-core/config"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli/require"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/flag"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/store"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/usage"
	"github.com/spf13/cobra"
	atlasv2 "go.mongodb.org/atlas-sdk/v20250312012/admin"
)

//go:generate go tool go.uber.org/mock/mockgen -typed -destination=describe_mock_test.go -package=interfaces . InterfaceEndpointDescriber

type InterfaceEndpointDescriber interface {
	InterfaceEndpoint(projectID, cloudProvider, privateEndpointID, endpointServiceID string) (*atlasv2.PrivateLinkEndpoint, error)
}

type DescribeOpts struct {
	cli.ProjectOpts
	cli.OutputOpts
	privateEndpointServiceID string
	privateEndpointGroupID   string
	store                    InterfaceEndpointDescriber
}

func (opts *DescribeOpts) initStore(ctx context.Context) func() error {
	return func() error {
		var err error
		opts.store, err = store.New(store.AuthenticatedPreset(config.Default()), store.WithContext(ctx))
		return err
	}
}

var describeTemplate = `ENDPOINT	STATUS	DELETE REQUESTED
{{.EndpointGroupName}}	{{.Status}}	{{.DeleteRequested}}
`

func (opts *DescribeOpts) Run() error {
	r, err := opts.store.InterfaceEndpoint(opts.ConfigProjectID(), provider, opts.privateEndpointGroupID, opts.privateEndpointServiceID)

	if err != nil {
		return err
	}

	return opts.Print(r)
}

// atlas privateEndpoint(s) gcp interface(s) describe <endpointGroupId> --endpointServiceId endpointServiceId [--projectId projectId].
func DescribeBuilder() *cobra.Command {
	opts := new(DescribeOpts)
	cmd := &cobra.Command{
		Use:     "describe <id>",
		Aliases: []string{"get"},
		Args:    require.ExactArgs(1),
		Short:   "Return a specific GCP private endpoint interface for your project.",
		Long:    fmt.Sprintf(usage.RequiredRole, "Project Read Only"),
		Annotations: map[string]string{
			"idDesc": "Unique identifier of the private endpoint you want to retrieve.",
			"output": describeTemplate,
		},
		Example: `  atlas privateEndpoints gcp interfaces describe endpoint-1 \
  --endpointServiceId 61eaca605af86411903de1dd`,
		PreRunE: func(cmd *cobra.Command, args []string) error {
			opts.privateEndpointGroupID = args[0]
			return opts.PreRunE(
				opts.ValidateProjectID,
				opts.initStore(cmd.Context()),
				opts.InitOutput(cmd.OutOrStdout(), describeTemplate),
			)
		},
		RunE: func(_ *cobra.Command, _ []string) error {
			return opts.Run()
		},
	}
	cmd.Flags().StringVar(&opts.privateEndpointServiceID, flag.EndpointServiceID, "", usage.EndpointServiceID)

	opts.AddProjectOptsFlags(cmd)
	opts.AddOutputOptFlags(cmd)

	_ = cmd.MarkFlagRequired(flag.EndpointServiceID)

	return cmd
}
