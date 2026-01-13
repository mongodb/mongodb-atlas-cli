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

package instance

import (
	"context"
	"errors"
	"fmt"

	"github.com/mongodb/atlas-cli-core/config"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli/require"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/store"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/usage"
	"github.com/spf13/cobra"
	atlasv2 "go.mongodb.org/atlas-sdk/v20250312012/admin"
)

//go:generate go tool go.uber.org/mock/mockgen -typed -destination=describe_mock_test.go -package=instance . StreamsDescriber

type StreamsDescriber interface {
	AtlasStream(string, string) (*atlasv2.StreamsTenant, error)
}

type DescribeOpts struct {
	cli.ProjectOpts
	cli.OutputOpts
	name  string
	store StreamsDescriber
}

const (
	describeTemplate = `ID	NAME	CLOUD	REGION
{{.Id}}	{{.Name}}	{{.DataProcessRegion.CloudProvider}}	{{.DataProcessRegion.Region}}
`
)

func (opts *DescribeOpts) Run() error {
	r, err := opts.store.AtlasStream(opts.ConfigProjectID(), opts.name)
	if err != nil {
		return err
	}

	return opts.Print(r)
}

func (opts *DescribeOpts) initStore(ctx context.Context) func() error {
	return func() error {
		var err error
		opts.store, err = store.New(store.AuthenticatedPreset(config.Default()), store.WithContext(ctx))
		return err
	}
}

// DescribeBuilder
// atlas streams instance describe [name].
func DescribeBuilder() *cobra.Command {
	opts := &DescribeOpts{}
	cmd := &cobra.Command{
		Use:   "describe <name>",
		Short: "Describe an Atlas Stream Processing instance for your project.",
		Long:  fmt.Sprintf(usage.RequiredRole, "Project Read Only"),
		Example: `  # Return an Atlas Stream Processing instance with a specific name:
  atlas streams instance describe myProcessor
  # Return a JSON-formatted Atlas Stream Processing instance with a specific name:
  atlas streams instance describe myProcessor --output json`,
		Args: require.ExactArgs(1),
		Annotations: map[string]string{
			"nameDesc": "Name of the Atlas Streams processor instance.",
			"output":   describeTemplate,
		},
		PreRunE: func(cmd *cobra.Command, args []string) error {
			if len(args) == 0 {
				return errors.New("the Atlas Streams Processor instance name is missing")
			}

			if len(args) != 0 {
				opts.name = args[0]
			}

			return opts.PreRunE(
				opts.ValidateProjectID,
				opts.initStore(cmd.Context()),
				opts.InitOutput(cmd.OutOrStdout(), describeTemplate),
			)
		},
		RunE: func(_ *cobra.Command, _ []string) error {
			return opts.Run()
		},
	}

	opts.AddProjectOptsFlags(cmd)
	opts.AddOutputOptFlags(cmd)

	return cmd
}
