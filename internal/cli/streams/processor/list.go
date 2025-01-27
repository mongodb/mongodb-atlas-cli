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

type ListOpts struct {
	cli.ProjectOpts
	cli.OutputOpts
	cli.ListOpts
	cli.StreamsOpts
	includeStats bool
	store        store.ProcessorLister
}

func (opts *ListOpts) Run() error {
	r, err := opts.store.ListProcessors(opts.ProjectID, opts.Instance)
	if err != nil {
		return err
	}

	sps := make([]atlasv2.StreamsProcessorWithStats, 0, len(*r.Results))
	opts.Output = jsonOutput
	if opts.includeStats {
		return opts.Print(*r.Results)
	}

	for _, res := range *r.Results {
		sps = append(sps, *atlasv2.NewStreamsProcessorWithStats(res.Id, res.Name, res.Pipeline, res.State))
	}
	return opts.Print(sps)
}

func (opts *ListOpts) initStore(ctx context.Context) func() error {
	return func() error {
		var err error
		opts.store, err = store.New(store.AuthenticatedPreset(config.Default()), store.WithContext(ctx))
		return err
	}
}

// atlas streams processor list.
func ListBuilder() *cobra.Command {
	opts := &ListOpts{}
	cmd := &cobra.Command{
		Use:   "list",
		Short: "List all the Atlas Stream Processing Processors for your project.",
		Long:  fmt.Sprintf(usage.RequiredOneOfRoles, streamsRoles),
		Example: `  # Return a JSON-formatted list of all Atlas Stream Processors for an instance 'ExampleInstance' for the project with ID 5e2211c17a3e5a48f5497de3:
  atlas streams processors list --projectId 5e2211c17a3e5a48f5497de3 --instance ExampleInstance`,
		Args: require.NoArgs,
		PreRunE: func(cmd *cobra.Command, _ []string) error {
			return opts.PreRunE(
				opts.ValidateProjectID,
				opts.ValidateInstance,
				opts.initStore(cmd.Context()),
			)
		},
		RunE: func(_ *cobra.Command, _ []string) error {
			return opts.Run()
		},
	}

	opts.AddProjectOptsFlags(cmd)
	opts.AddStreamsOptsFlags(cmd)
	opts.AddListOptsFlags(cmd)

	cmd.Flags().BoolVar(&opts.includeStats, flag.IncludeStats, false, usage.IncludeStreamProcessorStats)

	return cmd
}
