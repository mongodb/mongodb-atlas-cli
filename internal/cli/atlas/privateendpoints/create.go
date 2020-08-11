// Copyright 2020 MongoDB Inc
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
	"github.com/mongodb/mongocli/internal/cli"
	"github.com/mongodb/mongocli/internal/config"
	"github.com/mongodb/mongocli/internal/flag"
	"github.com/mongodb/mongocli/internal/output"
	"github.com/mongodb/mongocli/internal/store"
	"github.com/mongodb/mongocli/internal/usage"
	"github.com/spf13/cobra"
	"go.mongodb.org/atlas/mongodbatlas"
)

type CreateOpts struct {
	cli.GlobalOpts
	store    store.PrivateEndpointCreator
	region   string
	provider string
}

func (opts *CreateOpts) initStore() error {
	var err error
	opts.store, err = store.New(config.Default())
	return err
}

var createTemplate = "Private endpoint '{{.ID}}' created.\n"

func (opts *CreateOpts) Run() error {
	createRequest := opts.newPrivateEndpointConnection()

	r, err := opts.store.CreatePrivateEndpoint(opts.ConfigProjectID(), createRequest)
	if err != nil {
		return err
	}

	return output.Print(config.Default(), createTemplate, r)
}

func (opts *CreateOpts) newPrivateEndpointConnection() *mongodbatlas.PrivateEndpointConnection {
	createRequest := &mongodbatlas.PrivateEndpointConnection{
		Region:       opts.region,
		ProviderName: opts.provider,
	}
	return createRequest
}

// mongocli atlas datalake(s) create [--provider AWS] [--region <name>] --projectId projectId
func CreateBuilder() *cobra.Command {
	opts := &CreateOpts{}
	cmd := &cobra.Command{
		Use:   "create",
		Short: createPrivateEndpoint,
		Args:  cobra.NoArgs,
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return opts.PreRunE(opts.initStore)
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return opts.Run()
		},
	}
	cmd.Flags().StringVar(&opts.provider, flag.Provider, "AWS", usage.PrivateEndpointProvider)
	cmd.Flags().StringVar(&opts.region, flag.Region, "", usage.PrivateEndpointRegion)

	cmd.Flags().StringVar(&opts.ProjectID, flag.ProjectID, "", usage.ProjectID)

	_ = cmd.MarkFlagRequired(flag.Region)

	return cmd
}
