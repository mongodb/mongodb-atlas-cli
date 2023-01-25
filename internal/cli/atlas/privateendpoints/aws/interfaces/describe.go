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

	"github.com/mongodb/mongodb-atlas-cli/internal/cli"
	"github.com/mongodb/mongodb-atlas-cli/internal/cli/require"
	"github.com/mongodb/mongodb-atlas-cli/internal/config"
	"github.com/mongodb/mongodb-atlas-cli/internal/flag"
	"github.com/mongodb/mongodb-atlas-cli/internal/store"
	"github.com/mongodb/mongodb-atlas-cli/internal/usage"
	"github.com/spf13/cobra"
)

type DescribeOpts struct {
	cli.GlobalOpts
	cli.OutputOpts
	privateEndpointServiceID string
	privateEndpointID        string
	store                    store.InterfaceEndpointDescriber
}

func (opts *DescribeOpts) initStore(ctx context.Context) func() error {
	return func() error {
		var err error
		opts.store, err = store.New(store.AuthenticatedPreset(config.Default()), store.WithContext(ctx))
		return err
	}
}

var describeTemplate = `ID	STATUS	ERROR
{{.InterfaceEndpointID}}	{{.AWSConnectionStatus}}	{{.ErrorMessage}}
`

func (opts *DescribeOpts) Run() error {
	r, err := opts.store.InterfaceEndpoint(opts.ConfigProjectID(), provider, opts.privateEndpointServiceID, opts.privateEndpointID)

	if err != nil {
		return err
	}

	return opts.Print(r)
}

// mongocli atlas privateEndpoint(s) aws interface(s) describe <privateEndpointId> --endpointServiceId endpointServiceId [--projectId projectId].
func DescribeBuilder() *cobra.Command {
	opts := new(DescribeOpts)
	cmd := &cobra.Command{
		Use:     "describe <interfaceEndpointId>",
		Aliases: []string{"get"},
		Args:    require.ExactArgs(1),
		Short:   "Return the details for the specified AWS private endpoint interface for your project.",
		Annotations: map[string]string{
			"interfaceEndpointIdDesc": "Unique string that identifies the AWS private endpoint interface in AWS.",
		},
		Example: fmt.Sprintf(`  # Return the JSON-formatted details of the AWS private endpoint interface with the ID 	
		vpce-00713b5e644e830a3 in AWS for an AWS private endpoint with the ID 5f4fc14da2b47835a58c63a2 in Atlas:
  %s privateEndpoints aws interfaces describe 	
  vpce-00713b5e644e830a3 --endpointServiceId 5f4fc14da2b47835a58c63a2`, cli.ExampleAtlasEntryPoint()),
		PreRunE: func(cmd *cobra.Command, args []string) error {
			opts.privateEndpointID = args[0]
			return opts.PreRunE(
				opts.ValidateProjectID,
				opts.initStore(cmd.Context()),
				opts.InitOutput(cmd.OutOrStdout(), describeTemplate),
			)
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return opts.Run()
		},
	}
	cmd.Flags().StringVar(&opts.privateEndpointServiceID, flag.EndpointServiceID, "", usage.EndpointServiceID)

	cmd.Flags().StringVar(&opts.ProjectID, flag.ProjectID, "", usage.ProjectID)
	cmd.Flags().StringVarP(&opts.Output, flag.Output, flag.OutputShort, "", usage.FormatOut)

	_ = cmd.MarkFlagRequired(flag.EndpointServiceID)

	return cmd
}
