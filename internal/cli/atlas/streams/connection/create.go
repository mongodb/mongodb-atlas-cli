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

// This code was autogenerated at 2023-07-05T01:21:22Z. Note: Manual updates are allowed, but may be overwritten.

package connection

import (
	"context"
	"fmt"

	"github.com/mongodb/mongodb-atlas-cli/internal/cli"
	"github.com/mongodb/mongodb-atlas-cli/internal/cli/require"
	"github.com/mongodb/mongodb-atlas-cli/internal/config"
	"github.com/mongodb/mongodb-atlas-cli/internal/file"
	"github.com/mongodb/mongodb-atlas-cli/internal/flag"
	"github.com/mongodb/mongodb-atlas-cli/internal/store"
	"github.com/mongodb/mongodb-atlas-cli/internal/usage"
	"github.com/spf13/afero"
	"github.com/spf13/cobra"
	atlasv2 "go.mongodb.org/atlas-sdk/v20230201008/admin"
)

type CreateOpts struct {
	cli.GlobalOpts
	cli.OutputOpts
	store           store.ConnectionCreator
	name            string
	filename        string
	streamsInstance string
	fs              afero.Fs
}

func (opts *CreateOpts) initStore(ctx context.Context) func() error {
	return func() error {
		var err error
		opts.store, err = store.New(store.AuthenticatedPreset(config.Default()), store.WithContext(ctx))
		return err
	}
}

var createTemplate = "Connection {{.Name}} created.\n"

func (opts *CreateOpts) Run() error {
	createRequest, err := opts.newCreateRequest()
	if err != nil {
		return err
	}

	r, err := opts.store.CreateConnection(opts.ConfigProjectID(), opts.streamsInstance, createRequest)
	if err != nil {
		return err
	}

	return opts.Print(r)
}

func (opts *CreateOpts) newCreateRequest() (*atlasv2.StreamsConnection, error) {
	out := atlasv2.NewStreamsConnectionWithDefaults()
	if err := file.Load(opts.fs, opts.filename, out); err != nil {
		return nil, err
	}
	out.Name = &opts.name

	return out, nil
}

// atlas streams connection create <connectionName> [--projectId projectId].
func CreateBuilder() *cobra.Command {
	opts := &CreateOpts{
		fs: afero.NewOsFs(),
	}
	cmd := &cobra.Command{
		Use:   "create <connectionName>",
		Short: "Creates a connection for an Atlas streams instance.",
		Long:  fmt.Sprintf(usage.RequiredRole, "Project Owner"),
		Args:  require.ExactArgs(1),
		Annotations: map[string]string{
			"connectionNameDesc": "Name of the connection",
			"output":             createTemplate,
		},
		Example: `# create a new connection for Atlas Streams:
  atlas streams connection create kafkaprod -i test01 -f kafkaConfig.json
`,
		PreRunE: func(cmd *cobra.Command, args []string) error {
			opts.name = args[0]
			return opts.PreRunE(
				opts.ValidateProjectID,
				opts.initStore(cmd.Context()),
				opts.InitOutput(cmd.OutOrStdout(), createTemplate),
			)
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return opts.Run()
		},
	}

	cmd.Flags().StringVar(&opts.ProjectID, flag.ProjectID, "", usage.ProjectID)
	cmd.Flags().StringVarP(&opts.Output, flag.Output, flag.OutputShort, "", usage.FormatOut)
	cmd.Flags().StringVarP(&opts.streamsInstance, flag.Instance, flag.InstanceShort, "", usage.StreamsInstance)
	_ = cmd.RegisterFlagCompletionFunc(flag.Output, opts.AutoCompleteOutputFlag())

	cmd.Flags().StringVarP(&opts.filename, flag.File, flag.FileShort, "", usage.StreamsConnectionFilename)
	_ = cmd.MarkFlagFilename(flag.File)

	_ = cmd.MarkFlagRequired(flag.Instance)
	_ = cmd.MarkFlagRequired(flag.File)

	return cmd
}
