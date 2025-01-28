// Copyright 2025 MongoDB Inc
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

package processor

import (
	"context"
	"fmt"

	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli/require"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/config"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/flag"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/store"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/usage"
	"github.com/spf13/cobra"
	atlasv2 "go.mongodb.org/atlas-sdk/v20241113004/admin"
)

const jsonOutput = "json"

type DescribeOpts struct {
	cli.ProjectOpts
	cli.OutputOpts
	cli.StreamsOpts
	processorName string
	includeStats  bool
	store         store.ProcessorDescriber
}

func (opts *DescribeOpts) Run() error {
	r, err := opts.store.StreamProcessor(opts.ConfigProjectID(), opts.Instance, opts.processorName)
	if err != nil {
		return err
	}

	opts.Output = jsonOutput
	if opts.includeStats {
		return opts.Print(r)
	}

	// NewStreamsProcessorWithStats will return a Stream Processor without the stats field populated (which gets omitted)
	// but has a State field, unlike NewStreamsProcessor
	sp := atlasv2.NewStreamsProcessorWithStats(r.Id, r.Name, r.Pipeline, r.State)
	return opts.Print(sp)
}

func (opts *DescribeOpts) initStore(ctx context.Context) func() error {
	return func() error {
		var err error
		opts.store, err = store.New(store.AuthenticatedPreset(config.Default()), store.WithContext(ctx))
		return err
	}
}

// atlas streams processor describe <processorName>.
func DescribeBuilder() *cobra.Command {
	opts := &DescribeOpts{}
	cmd := &cobra.Command{
		Use:   "describe <processorName>",
		Short: "Get details about an Atlas Stream Processor in a Stream Processing Instance.",
		Long:  fmt.Sprintf(usage.RequiredOneOfRoles, streamsRoles),
		Example: `# Return a JSON-formatted view of stream processor 'ExampleProcessor' for an instance 'ExampleInstance':
  atlas streams processors describe ExampleProcessor --instance ExampleInstance`,
		Args: require.ExactArgs(1),
		Annotations: map[string]string{
			"processorNameDesc": "Name of the Stream Processor",
		},
		PreRunE: func(cmd *cobra.Command, args []string) error {
			if err := opts.PreRunE(
				opts.ValidateProjectID,
				opts.ValidateInstance,
				opts.initStore(cmd.Context()),
			); err != nil {
				return err
			}
			opts.processorName = args[0]
			return nil
		},
		RunE: func(_ *cobra.Command, _ []string) error {
			return opts.Run()
		},
	}

	opts.AddProjectOptsFlags(cmd)
	opts.AddStreamsOptsFlags(cmd)

	cmd.Flags().BoolVar(&opts.includeStats, flag.IncludeStats, false, usage.IncludeStreamProcessorStats)

	return cmd
}
