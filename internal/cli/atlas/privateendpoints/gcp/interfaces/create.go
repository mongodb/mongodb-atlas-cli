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
	"errors"

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
	store store.InterfaceEndpointCreator
	//	privateEndpointID   string
	interfaceEndpointID string
	endpointGroupName   string
	gcpProjectID        string
}

func (opts *CreateOpts) initStore(ctx context.Context) func() error {
	return func() error {
		var err error
		opts.store, err = store.New(store.AuthenticatedPreset(config.Default()), store.WithContext(ctx))
		return err
	}
}

var createTemplate = "Interface endpoint '{{.InterfaceEndpointID}}' created.\n"

/*
func (opts *CreateOpts) Run() error {
	r, err := opts.store.CreateInterfaceEndpoint(opts.ConfigProjectID(), provider, opts.interfaceEndpointID, opts.newInterfaceEndpointConnection())
	if err != nil {
		return err
	}

	return opts.Print(r)
}
*/
// TODO: remove...
func (opts *CreateOpts) Run() error {
	err := errors.New(provider)
	return err
}

/*
func (opts *CreateOpts) newInterfaceEndpointConnection() *atlas.InterfaceEndpointConnection {
	return &atlas.InterfaceEndpointConnection{
		//		ID: opts.privateEndpointID,
		EndpointGroupName: opts.endpointGroupName,
		GCPProjectID:      opts.gcpProjectID,
	}
}
*/

// TODO: Better way to indicate --endpoints list?
// mongocli atlas privateEndpoint(s)|privateendpoint(s) gcp interface(s) create <atlasPrivateEndpointId> --endpointGroupName endpointGroupName --gcpProjectId gcpProjectId --endpoints <list> [--projectId projectId].
func CreateBuilder() *cobra.Command {
	opts := &CreateOpts{}
	cmd := &cobra.Command{
		Use:     "create <atlasPrivateEndpointId>",
		Aliases: []string{"add"},
		Short:   "Create a GCP private endpoint interface.",
		Args:    require.ExactArgs(1),
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return opts.PreRunE(
				opts.ValidateProjectID,
				opts.initStore(cmd.Context()),
				opts.InitOutput(cmd.OutOrStdout(), createTemplate),
			)
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			opts.interfaceEndpointID = args[0]
			return opts.Run()
		},
	}
	cmd.Flags().StringVar(&opts.endpointGroupName, flag.EndpointGroupName, "", usage.EndpointGroupName)
	cmd.Flags().StringVar(&opts.gcpProjectID, flag.GCPProjectID, "", usage.GCPProjectID)

	cmd.Flags().StringVar(&opts.ProjectID, flag.ProjectID, "", usage.ProjectID)
	cmd.Flags().StringVarP(&opts.Output, flag.Output, flag.OutputShort, "", usage.FormatOut)

	_ = cmd.MarkFlagRequired(flag.EndpointGroupName)
	_ = cmd.MarkFlagRequired(flag.GCPProjectID)

	return cmd
}
