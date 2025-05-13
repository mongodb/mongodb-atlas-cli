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

package datalakepipelines

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
	atlasv2 "go.mongodb.org/atlas-sdk/v20250312002/admin"
)

//go:generate go tool go.uber.org/mock/mockgen -typed -destination=trigger_mock_test.go -package=datalakepipelines . PipelinesTriggerer

type PipelinesTriggerer interface {
	PipelineTrigger(string, string, string) (*atlasv2.IngestionPipelineRun, error)
}

type TriggerOpts struct {
	cli.ProjectOpts
	cli.OutputOpts
	id         string
	snapshotID string
	store      PipelinesTriggerer
}

func (opts *TriggerOpts) initStore(ctx context.Context) func() error {
	return func() error {
		var err error
		opts.store, err = store.New(store.AuthenticatedPreset(config.Default()), store.WithContext(ctx))
		return err
	}
}

var triggerTemplate = `ID	DATASET NAME	STATE
{{.Id}}	{{.DatasetName}}	{{.State}}
`

func (opts *TriggerOpts) Run() error {
	r, err := opts.store.PipelineTrigger(opts.ConfigProjectID(), opts.id, opts.snapshotID)
	if err != nil {
		return err
	}

	return opts.Print(r)
}

// atlas dataLakePipelines trigger <pipelineName> [--projectId projectId].
func TriggerBuilder() *cobra.Command {
	opts := new(TriggerOpts)
	cmd := &cobra.Command{
		Use:   "trigger <pipelineName>",
		Short: "Trigger the specified data lake pipeline for your project.",
		Long:  fmt.Sprintf(usage.RequiredRole, "Project Owner"),
		Args:  require.ExactArgs(1),
		Annotations: map[string]string{
			"pipelineNameDesc": "Label that identifies the pipeline",
		},
		Deprecated: "Data Lake Pipelines is deprecated. Please see: https://dochub.mongodb.org/core/data-lake-deprecation.",
		Example: `# trigger pipeline 'Pipeline1':
  atlas dataLakePipelines trigger Pipeline1
`,
		PreRunE: func(cmd *cobra.Command, _ []string) error {
			return opts.PreRunE(
				opts.ValidateProjectID,
				opts.initStore(cmd.Context()),
				opts.InitOutput(cmd.OutOrStdout(), triggerTemplate),
			)
		},
		RunE: func(_ *cobra.Command, args []string) error {
			opts.id = args[0]

			return opts.Run()
		},
	}

	cmd.Flags().StringVar(&opts.snapshotID, flag.SnapshotID, "", usage.SnapshotID)
	_ = cmd.MarkFlagRequired(flag.SnapshotID)

	opts.AddProjectOptsFlags(cmd)
	opts.AddOutputOptFlags(cmd)

	return cmd
}
