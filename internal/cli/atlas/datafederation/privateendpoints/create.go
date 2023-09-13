// Copyright 2023 MongoDB Inc
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

// This code was autogenerated at 2023-06-22T17:46:28+01:00. Note: Manual updates are allowed, but may be overwritten.

package privateendpoints

import (
	"context"
	"fmt"

	"github.com/mongodb/mongodb-atlas-cli/internal/cli"
	"github.com/mongodb/mongodb-atlas-cli/internal/cli/require"
	"github.com/mongodb/mongodb-atlas-cli/internal/config"
	"github.com/mongodb/mongodb-atlas-cli/internal/flag"
	store "github.com/mongodb/mongodb-atlas-cli/internal/store/atlas"
	"github.com/mongodb/mongodb-atlas-cli/internal/usage"
	"github.com/spf13/cobra"
	"go.mongodb.org/atlas-sdk/v20230201007/admin"
)

type CreateOpts struct {
	cli.GlobalOpts
	cli.OutputOpts
	store      store.DataFederationPrivateEndpointCreator
	endpointID string
	comment    string
}

func (opts *CreateOpts) initStore(ctx context.Context) func() error {
	return func() error {
		var err error
		opts.store, err = store.New(store.AuthenticatedPreset(config.Default()), store.WithContext(ctx))
		return err
	}
}

var createTemplate = `Data federation private endpoint {{(index .Results 0).EndpointId}} created.`

func (opts *CreateOpts) Run() error {
	createRequest := opts.newCreateRequest()

	r, err := opts.store.CreateDataFederationPrivateEndpoint(opts.ConfigProjectID(), createRequest)
	if err != nil {
		return err
	}

	return opts.Print(r)
}

func (opts *CreateOpts) newCreateRequest() *admin.PrivateNetworkEndpointIdEntry {
	p := admin.NewPrivateNetworkEndpointIdEntry(opts.endpointID)
	p.Comment = &opts.comment
	return p
}

// atlas dataFederation privateEndpoints create <endpointId> [--projectId projectId].
func CreateBuilder() *cobra.Command {
	opts := &CreateOpts{}
	cmd := &cobra.Command{
		Use:   "create <endpointId>",
		Short: "Creates a new Data Federation private endpoint.",
		Long:  fmt.Sprintf(usage.RequiredRole, "Project Owner"),
		Args:  require.ExactArgs(1),
		Annotations: map[string]string{
			"endpointIdDesc": "Endpoint identifier of the data federation private endpoint.",
			"output":         createTemplate,
		},
		Example: `# create data federation private endpoint:
  atlas dataFederation privateEndpoints create 507f1f77bcf86cd799439011 --comment "comment"
`,
		PreRunE: func(cmd *cobra.Command, args []string) error {
			opts.endpointID = args[0]
			return opts.PreRunE(
				opts.ValidateProjectID,
				opts.initStore(cmd.Context()),
				opts.InitOutput(cmd.OutOrStdout(), createTemplate),
			)
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return opts.Run()
		},
	}

	cmd.Flags().StringVar(&opts.comment, flag.Comment, "", usage.Comment)

	cmd.Flags().StringVar(&opts.ProjectID, flag.ProjectID, "", usage.ProjectID)
	cmd.Flags().StringVarP(&opts.Output, flag.Output, flag.OutputShort, "", usage.FormatOut)
	_ = cmd.RegisterFlagCompletionFunc(flag.Output, opts.AutoCompleteOutputFlag())

	return cmd
}
