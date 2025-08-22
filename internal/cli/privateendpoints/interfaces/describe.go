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

package interfaces

import (
	"context"

	"github.com/mongodb/atlas-cli-core/config"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli/require"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/flag"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/store"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/usage"
	"github.com/spf13/cobra"
	atlas "go.mongodb.org/atlas/mongodbatlas"
)

//go:generate go tool go.uber.org/mock/mockgen -typed -destination=describe_mock_test.go -package=interfaces . InterfaceEndpointDescriberDeprecated

type InterfaceEndpointDescriberDeprecated interface {
	InterfaceEndpointDeprecated(string, string, string) (*atlas.InterfaceEndpointConnectionDeprecated, error)
}

type DescribeOpts struct {
	cli.ProjectOpts
	cli.OutputOpts
	id                string
	privateEndpointID string
	store             InterfaceEndpointDescriberDeprecated
}

func (opts *DescribeOpts) initStore(ctx context.Context) func() error {
	return func() error {
		var err error
		opts.store, err = store.New(store.AuthenticatedPreset(config.Default()), store.WithContext(ctx))
		return err
	}
}

var describeTemplate = `ID	STATUS	ERROR
{{.ID}}	{{.ConnectionStatus}}	{{.ErrorMessage}}
`

func (opts *DescribeOpts) Run() error {
	r, err := opts.store.InterfaceEndpointDeprecated(opts.ConfigProjectID(), opts.privateEndpointID, opts.id)

	if err != nil {
		return err
	}

	return opts.Print(r)
}

// atlas privateEndpoint(s) interface(s) describe <atlasPrivateEndpointId> [--privateEndpointId privateEndpointID][--projectId projectId].
func DescribeBuilder() *cobra.Command {
	opts := new(DescribeOpts)
	cmd := &cobra.Command{
		Use:         "describe <atlasPrivateEndpointId>",
		Aliases:     []string{"get"},
		Args:        require.ExactArgs(1),
		Short:       "Return a specific private endpoint interface for your project.",
		Annotations: map[string]string{"output": describeTemplate},
		PreRunE: func(cmd *cobra.Command, args []string) error {
			opts.id = args[0]
			return opts.PreRunE(
				opts.ValidateProjectID,
				opts.initStore(cmd.Context()),
				opts.InitOutput(cmd.OutOrStdout(), describeTemplate),
			)
		},
		RunE: func(_ *cobra.Command, _ []string) error {
			return opts.Run()
		},
		Deprecated: "Please use atlas privateEndpoints aws interfaces describe <atlasPrivateEndpointId> [--privateEndpointId privateEndpointID] [--projectId projectId]",
	}
	cmd.Flags().StringVar(&opts.privateEndpointID, flag.PrivateEndpointID, "", usage.PrivateEndpointID)

	opts.AddProjectOptsFlags(cmd)
	opts.AddOutputOptFlags(cmd)

	_ = cmd.MarkFlagRequired(flag.PrivateEndpointID)

	return cmd
}
