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

package privateendpoints

import (
	"context"

	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli/require"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/config"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/flag"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/store"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/usage"
	"github.com/spf13/cobra"
	atlas "go.mongodb.org/atlas/mongodbatlas"
)

//go:generate mockgen -typed -destination=create_mock_test.go -package=privateendpoints . PrivateEndpointCreatorDeprecated

type PrivateEndpointCreatorDeprecated interface {
	CreatePrivateEndpointDeprecated(string, *atlas.PrivateEndpointConnectionDeprecated) (*atlas.PrivateEndpointConnectionDeprecated, error)
}

type CreateOpts struct {
	cli.ProjectOpts
	cli.OutputOpts
	store    PrivateEndpointCreatorDeprecated
	region   string
	provider string
}

func (opts *CreateOpts) initStore(ctx context.Context) func() error {
	return func() error {
		var err error
		opts.store, err = store.New(store.AuthenticatedPreset(config.Default()), store.WithContext(ctx))
		return err
	}
}

var createTemplate = "Private endpoint '{{.ID}}' created.\n"

func (opts *CreateOpts) Run() error {
	createRequest := opts.newPrivateEndpointConnection()

	r, err := opts.store.CreatePrivateEndpointDeprecated(opts.ConfigProjectID(), createRequest)
	if err != nil {
		return err
	}

	return opts.Print(r)
}

func (opts *CreateOpts) newPrivateEndpointConnection() *atlas.PrivateEndpointConnectionDeprecated {
	createRequest := &atlas.PrivateEndpointConnectionDeprecated{
		Region:       opts.region,
		ProviderName: opts.provider,
	}
	return createRequest
}

// atlas privateEndpoint(s) create [--provider AWS] [--region <name>] --projectId projectId.
func CreateBuilder() *cobra.Command {
	opts := &CreateOpts{}
	cmd := &cobra.Command{
		Use:   "create",
		Short: "Create a new private endpoint for your project.",
		Annotations: map[string]string{
			"output": createTemplate,
		},
		Args: require.NoArgs,
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
		Deprecated: "Please use atlas privateEndpoints aws create [--region region] [--projectId projectId]",
	}
	cmd.Flags().StringVar(&opts.provider, flag.Provider, "AWS", usage.PrivateEndpointProvider)
	cmd.Flags().StringVar(&opts.region, flag.Region, "", usage.PrivateEndpointRegion)

	opts.AddProjectOptsFlags(cmd)
	opts.AddOutputOptFlags(cmd)

	_ = cmd.MarkFlagRequired(flag.Region)

	return cmd
}
