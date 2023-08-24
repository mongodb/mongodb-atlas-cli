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

// This code was autogenerated at 2023-07-04T05:39:19Z. Note: Manual updates are allowed, but may be overwritten.

package connection

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
	id              string
	streamsInstance string
	store           store.StreamsConnectionDescriber
}

func (opts *DescribeOpts) initStore(ctx context.Context) func() error {
	return func() error {
		var err error
		opts.store, err = store.New(store.AuthenticatedPreset(config.Default()), store.WithContext(ctx))
		return err
	}
}

var describeTemplate = `NAME	TYPE	INSTANCE	SERVERS
{{.Name}}	{{.Type}}	{{.Instance}}	{{.Servers}}
`

func (opts *DescribeOpts) Run() error {
	r, err := opts.store.StreamConnection(opts.ConfigProjectID(), opts.streamsInstance, opts.id)
	if err != nil {
		return err
	}

	if opts.IsJSONOutput() || opts.IsJSONPathOutput() {
		return opts.Print(r.StreamsConnection)
	}

	return opts.Print(r)
}

// atlas streams connection describe <streamConnectionName> --instance <instanceName> [--projectId projectId].
func DescribeBuilder() *cobra.Command {
	opts := new(DescribeOpts)
	cmd := &cobra.Command{
		Use:   "describe <streamConnectionName>",
		Short: "Return the details for the specified Atlas Streams Connection.",
		Long:  fmt.Sprintf(usage.RequiredRole, "Project Owner"),
		Args:  require.ExactArgs(1),
		Annotations: map[string]string{
			"streamConnectionNameDesc": "Name of the Atlas Streams Connection",
		},
		Example: `# retrieves stream connection 'ExampleConnection' in instance 'ExampleInstance':
atlas streams connection describe ExampleConnection --instance ExampleInstance
`,
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return opts.PreRunE(
				opts.ValidateProjectID,
				opts.initStore(cmd.Context()),
				opts.InitOutput(cmd.OutOrStdout(), describeTemplate),
			)
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			opts.id = args[0]

			return opts.Run()
		},
	}

	cmd.Flags().StringVar(&opts.ProjectID, flag.ProjectID, "", usage.ProjectID)
	cmd.Flags().StringVarP(&opts.streamsInstance, flag.Instance, flag.InstanceShort, "", usage.StreamsInstance)
	cmd.Flags().StringVarP(&opts.Output, flag.Output, flag.OutputShort, "", usage.FormatOut)
	_ = cmd.RegisterFlagCompletionFunc(flag.Output, opts.AutoCompleteOutputFlag())
	_ = cmd.MarkFlagRequired(flag.Instance)

	return cmd
}
