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

package aws

import (
	"context"
	"fmt"

	"github.com/mongodb/atlas-cli-core/config"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli/require"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/pointer"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/store"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/usage"
	"github.com/spf13/cobra"
	atlasv2 "go.mongodb.org/atlas-sdk/v20250312014/admin"
)

var listTemplate = `ID	ENDPOINT PROVIDER	TYPE	COMMENT{{range valueOrEmptySlice .Results}}
{{.EndpointId}}	{{.Provider}}	{{.Type}}	{{.Comment}}{{end}}
`

//go:generate go tool go.uber.org/mock/mockgen -typed -destination=list_mock_test.go -package=aws . DataLakePrivateEndpointLister

type DataLakePrivateEndpointLister interface {
	DataLakePrivateEndpoints(*atlasv2.ListPrivateEndpointIdsApiParams) (*atlasv2.PaginatedPrivateNetworkEndpointIdEntry, error)
}

type ListOpts struct {
	cli.ProjectOpts
	cli.OutputOpts
	cli.ListOpts
	store DataLakePrivateEndpointLister
}

func (opts *ListOpts) initStore(ctx context.Context) func() error {
	return func() error {
		var err error
		opts.store, err = store.New(store.AuthenticatedPreset(config.Default()), store.WithContext(ctx))
		return err
	}
}

func (opts *ListOpts) Run() error {
	r, err := opts.store.DataLakePrivateEndpoints(opts.newDatalakePrivateEndpointsListOpts())

	if err != nil {
		return err
	}

	return opts.Print(r)
}

func (opts *ListOpts) newDatalakePrivateEndpointsListOpts() *atlasv2.ListPrivateEndpointIdsApiParams {
	return &atlasv2.ListPrivateEndpointIdsApiParams{
		GroupId:      opts.ConfigProjectID(),
		PageNum:      pointer.Get(opts.PageNum),
		ItemsPerPage: pointer.Get(opts.ItemsPerPage),
		IncludeCount: pointer.Get(!opts.OmitCount),
	}
}

// atlas privateEndpoint(s)|privateendpoint(s) dataLakes aws list|ls [--projectId projectId].
func ListBuilder() *cobra.Command {
	opts := new(ListOpts)
	cmd := &cobra.Command{
		Use:        "list",
		Aliases:    []string{"ls"},
		Short:      "List Data Lake private endpoints for your project.",
		Long:       fmt.Sprintf(usage.RequiredRole, "Project Read Only"),
		Args:       require.NoArgs,
		Deprecated: "Please use 'atlas datafederation privateendpoints list'",
		Annotations: map[string]string{
			"output": listTemplate,
		},
		PreRunE: func(cmd *cobra.Command, _ []string) error {
			return opts.PreRunE(
				opts.ValidateProjectID,
				opts.initStore(cmd.Context()),
				opts.InitOutput(cmd.OutOrStdout(), listTemplate),
			)
		},
		RunE: func(_ *cobra.Command, _ []string) error {
			return opts.Run()
		},
	}

	opts.AddListOptsFlags(cmd)

	opts.AddProjectOptsFlags(cmd)
	opts.AddOutputOptFlags(cmd)

	return cmd
}
