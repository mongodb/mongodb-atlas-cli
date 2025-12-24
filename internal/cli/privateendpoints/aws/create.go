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

package aws

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
	atlasv2 "go.mongodb.org/atlas-sdk/v20250312011/admin"
)

//go:generate go tool go.uber.org/mock/mockgen -typed -destination=create_mock_test.go -package=aws . PrivateEndpointCreator

type PrivateEndpointCreator interface {
	CreatePrivateEndpoint(string, *atlasv2.CloudProviderEndpointServiceRequest) (*atlasv2.EndpointService, error)
}

type CreateOpts struct {
	cli.ProjectOpts
	cli.PreRunOpts
	cli.OutputOpts
	store  PrivateEndpointCreator
	region string
}

func (opts *CreateOpts) initStore(ctx context.Context) func() error {
	return func() error {
		var err error
		opts.store, err = store.New(store.AuthenticatedPreset(config.Default()), store.WithContext(ctx))
		return err
	}
}

var createTemplate = "Private endpoint '{{.Id}}' created.\n"

func (opts *CreateOpts) Run() error {
	createRequest := opts.newPrivateEndpointConnection()

	r, err := opts.store.CreatePrivateEndpoint(opts.ConfigProjectID(), createRequest)
	if err != nil {
		return err
	}

	return opts.Print(r)
}

func (opts *CreateOpts) newPrivateEndpointConnection() *atlasv2.CloudProviderEndpointServiceRequest {
	createRequest := &atlasv2.CloudProviderEndpointServiceRequest{
		Region:       opts.region,
		ProviderName: provider,
	}
	return createRequest
}

// atlas privateEndpoint(s) aws create [--region <name>] --projectId projectId.
func CreateBuilder() *cobra.Command {
	opts := &CreateOpts{}
	cmd := &cobra.Command{
		Use:   "create",
		Short: "Create a new AWS private endpoint for your project.",
		Long: `To learn more about how to set up private endpoints with the Atlas CLI, see the tutorial on the Atlas CLI tab here: https://www.mongodb.com/docs/atlas/security-cluster-private-endpoint/.

` + fmt.Sprintf(usage.RequiredRole, "Project Owner"),
		Args: require.NoArgs,
		Annotations: map[string]string{
			"output": createTemplate,
		},
		Example: `  # Create a private endpoint connection for AWS in the us-east-1 region for the project with the ID 5e2211c17a3e5a48f5497de3:
  atlas privateEndpoints aws create --region us-east-1 --projectId 5e2211c17a3e5a48f5497de3 --output json`,
		PreRunE: func(cmd *cobra.Command, _ []string) error {
			return opts.PreRunE(
				opts.ValidateProjectID,
				opts.initStore(cmd.Context()),
				opts.InitOutput(cmd.OutOrStdout(), createTemplate),
			)
		},
		RunE: func(_ *cobra.Command, _ []string) error {
			return opts.Run()
		},
	}

	cmd.Flags().StringVar(&opts.region, flag.Region, "", usage.PrivateEndpointRegion)

	opts.AddProjectOptsFlags(cmd)
	opts.AddOutputOptFlags(cmd)

	_ = cmd.MarkFlagRequired(flag.Region)

	return cmd
}
