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
	"errors"
	"fmt"

	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli/require"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/config"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/file"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/flag"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/store"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/usage"
	"github.com/spf13/afero"
	"github.com/spf13/cobra"
	atlasv2 "go.mongodb.org/atlas-sdk/v20241113004/admin"
)

var createTemplate = "Processor {{.Name}} created.\n"

type CreateOpts struct {
	cli.ProjectOpts
	cli.OutputOpts
	cli.StreamsOpts
	store    store.ProcessorCreator
	filename string
	fs       afero.Fs
}

func (opts *CreateOpts) Run() error {
	createParams, err := opts.newCreateRequest()
	if err != nil {
		return err
	}

	result, err := opts.store.CreateStreamProcessor(createParams.GroupId, createParams.TenantName, createParams.StreamsProcessor)
	if err != nil {
		return err
	}

	return opts.Print(result)
}

func (opts *CreateOpts) initStore(ctx context.Context) func() error {
	return func() error {
		var err error
		opts.store, err = store.New(store.AuthenticatedPreset(config.Default()), store.WithContext(ctx))
		return err
	}
}

func (opts *CreateOpts) newCreateRequest() (*atlasv2.CreateStreamProcessorApiParams, error) {
	processor := atlasv2.NewStreamsProcessorWithDefaults()
	if err := file.Load(opts.fs, opts.filename, processor); err != nil {
		return nil, err
	}

	if processor.Name == nil || len(*processor.Name) == 0 {
		return nil, errors.New("streams processor name missing")
	}

	createParams := new(atlasv2.CreateStreamProcessorApiParams)
	createParams.GroupId = opts.ConfigProjectID()
	createParams.TenantName = opts.Instance
	createParams.StreamsProcessor = processor

	return createParams, nil
}

// atlas streams processor create [--projectId projectId].
func CreateBuilder() *cobra.Command {
	opts := &CreateOpts{
		fs: afero.NewOsFs(),
	}
	cmd := &cobra.Command{
		Use:   "create",
		Short: "Creates a stream processor for an Atlas Stream Processing instance.",
		Long:  fmt.Sprintf(usage.RequiredOneOfRoles, streamsRoles),
		Args:  require.NoArgs,
		Annotations: map[string]string{
			"output": createTemplate,
		},
		Example: `# create a new stream processor for Atlas Stream Processing Instance:
  atlas streams processor create -i test01 -f processorConfig.json
`,
		PreRunE: func(cmd *cobra.Command, _ []string) error {
			return opts.PreRunE(
				opts.ValidateProjectID,
				opts.ValidateInstance,
				opts.initStore(cmd.Context()),
				opts.InitOutput(cmd.OutOrStdout(), createTemplate),
			)
		},
		RunE: func(_ *cobra.Command, _ []string) error {
			return opts.Run()
		},
	}

	opts.AddProjectOptsFlags(cmd)
	opts.AddOutputOptFlags(cmd)
	opts.AddStreamsOptsFlags(cmd)

	cmd.Flags().StringVarP(&opts.filename, flag.File, flag.FileShort, "", usage.StreamsProcessorFilename)
	_ = cmd.MarkFlagFilename(flag.File)

	_ = cmd.MarkFlagRequired(flag.File)

	return cmd
}
