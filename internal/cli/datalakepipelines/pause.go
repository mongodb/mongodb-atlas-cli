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
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/store"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/usage"
	"github.com/spf13/cobra"
	atlasv2 "go.mongodb.org/atlas-sdk/v20250312005/admin"
)

//go:generate go tool go.uber.org/mock/mockgen -typed -destination=pause_mock_test.go -package=datalakepipelines . PipelinesPauser

type PipelinesPauser interface {
	PipelinePause(string, string) (*atlasv2.DataLakeIngestionPipeline, error)
}

type PauseOpts struct {
	cli.ProjectOpts
	cli.OutputOpts
	id    string
	store PipelinesPauser
}

func (opts *PauseOpts) initStore(ctx context.Context) func() error {
	return func() error {
		var err error
		opts.store, err = store.New(store.AuthenticatedPreset(config.Default()), store.WithContext(ctx))
		return err
	}
}

var pauseTemplate = `ID	NAME	STATE
{{.Id}}	{{.Name}}	{{.State}}
`

func (opts *PauseOpts) Run() error {
	r, err := opts.store.PipelinePause(opts.ConfigProjectID(), opts.id)
	if err != nil {
		return err
	}

	return opts.Print(r)
}

// atlas dataLakePipelines pause <pipelineName> [--projectId projectId].
func PauseBuilder() *cobra.Command {
	opts := new(PauseOpts)
	cmd := &cobra.Command{
		Use:   "pause <pipelineName>",
		Short: "Pause the specified data lake pipeline for your project.",
		Long:  fmt.Sprintf(usage.RequiredRole, "Project Owner"),
		Args:  require.ExactArgs(1),
		Annotations: map[string]string{
			"pipelineNameDesc": "Label that identifies the pipeline",
		},
		Deprecated: "Data Lake Pipelines is deprecated. Please see: https://dochub.mongodb.org/core/data-lake-deprecation.",
		Example: `# pause pipeline 'Pipeline1':
  atlas dataLakePipelines pause Pipeline1
`,
		PreRunE: func(cmd *cobra.Command, _ []string) error {
			return opts.PreRunE(
				opts.ValidateProjectID,
				opts.initStore(cmd.Context()),
				opts.InitOutput(cmd.OutOrStdout(), pauseTemplate),
			)
		},
		RunE: func(_ *cobra.Command, args []string) error {
			opts.id = args[0]

			return opts.Run()
		},
	}

	opts.AddProjectOptsFlags(cmd)
	opts.AddOutputOptFlags(cmd)

	return cmd
}
