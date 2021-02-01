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
	"github.com/mongodb/mongocli/internal/cli"
	"github.com/mongodb/mongocli/internal/cli/require"
	"github.com/mongodb/mongocli/internal/config"
	"github.com/mongodb/mongocli/internal/flag"
	"github.com/mongodb/mongocli/internal/store"
	"github.com/mongodb/mongocli/internal/usage"
	"github.com/spf13/cobra"
)

type DescribeOpts struct {
	cli.GlobalOpts
	cli.OutputOpts
	id                string
	privateEndpointID string
	store             store.InterfaceEndpointDescriber
}

func (opts *DescribeOpts) init() error {
	var err error
	opts.store, err = store.New(config.Default())
	return err
}

var describeTemplate = `ID	IP ADDRESS	STATUS	ERROR
{{.PrivateEndpointResourceID}}	{{.PrivateEndpointIPAddress}}	{{.AzureStatus}}	{{.ErrorMessage}}
`

func (opts *DescribeOpts) Run() error {
	r, err := opts.store.InterfaceEndpoint(opts.ConfigProjectID(), provider, opts.id, opts.privateEndpointID)

	if err != nil {
		return err
	}

	return opts.Print(r)
}

// mongocli atlas privateEndpoint(s) azure interface(s) describe <atlasPrivateEndpointId> [--privateEndpointId privateEndpointID][--projectId projectId]
func DescribeBuilder() *cobra.Command {
	opts := new(DescribeOpts)
	cmd := &cobra.Command{
		Use:     "describe <atlasPrivateEndpointId>",
		Aliases: []string{"get"},
		Args:    require.ExactArgs(1),
		Short:   describe,
		PreRunE: func(cmd *cobra.Command, args []string) error {
			opts.id = args[0]
			return opts.PreRunE(
				opts.ValidateProjectID,
				opts.init,
				opts.InitOutput(cmd.OutOrStdout(), describeTemplate),
			)
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return opts.Run()
		},
	}
	cmd.Flags().StringVar(&opts.privateEndpointID, flag.PrivateEndpointID, "", usage.PrivateEndpointID)

	cmd.Flags().StringVar(&opts.ProjectID, flag.ProjectID, "", usage.ProjectID)
	cmd.Flags().StringVarP(&opts.Output, flag.Output, flag.OutputShort, "", usage.FormatOut)

	_ = cmd.MarkFlagRequired(flag.PrivateEndpointID)

	return cmd
}
