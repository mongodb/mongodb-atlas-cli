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
)

type DeleteOpts struct {
	cli.ProjectOpts
	*cli.DeleteOpts
	cli.StreamsOpts
	store store.ProcessorDeleter
}

func (opts *DeleteOpts) initStore(ctx context.Context) func() error {
	return func() error {
		var err error
		opts.store, err = store.New(store.AuthenticatedPreset(config.Default()), store.WithContext(ctx))
		return err
	}
}

func (opts *DeleteOpts) Run() error {
	return opts.Delete(opts.store.DeleteStreamProcessor, opts.ProjectID, opts.Instance)
}

// atlas streams processor delete <processorName>.
func DeleteBuilder() *cobra.Command {
	opts := &DeleteOpts{
		DeleteOpts: cli.NewDeleteOpts("Atlas Stream Processor '%s' deleted\n", "Atlas Stream Processor not deleted"),
	}
	cmd := &cobra.Command{
		Use:   "delete <processorName>",
		Short: "Delete an Atlas Stream Processor in a Stream Processing Instance.",
		Long:  fmt.Sprintf(usage.RequiredRole, "Project Owner"),
		Example: `# deletes stream processor 'ExampleSP' from instance 'ExampleInstance':
  atlas streams processors delete ExampleSP --instance ExampleInstance`,
		Args: require.ExactArgs(1),
		Annotations: map[string]string{
			"processorNameDesc": "Name of the Stream Processor",
			"output":            opts.SuccessMessage(),
		},
		PreRunE: func(cmd *cobra.Command, args []string) error {
			if err := opts.PreRunE(
				opts.ValidateProjectID,
				opts.ValidateInstance,
				opts.initStore(cmd.Context()),
			); err != nil {
				return err
			}
			opts.Entry = args[0]
			return opts.Prompt()
		},
		RunE: func(_ *cobra.Command, _ []string) error {
			return opts.Run()
		},
	}

	opts.AddProjectOptsFlags(cmd)
	opts.AddStreamsOptsFlags(cmd)

	cmd.Flags().BoolVar(&opts.Confirm, flag.Force, false, usage.Force)

	return cmd
}
