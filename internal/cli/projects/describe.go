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

package projects

import (
	"context"
	"fmt"

	"github.com/mongodb/atlas-cli-core/config"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli/require"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/store"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/usage"
	"github.com/spf13/cobra"
	atlasv2 "go.mongodb.org/atlas-sdk/v20250312014/admin"
)

const describeTemplate = `ID	NAME
{{.Id}}	{{.Name}}
`

//go:generate go tool go.uber.org/mock/mockgen -typed -destination=describe_mock_test.go -package=projects . ProjectDescriber

type ProjectDescriber interface {
	Project(string) (*atlasv2.Group, error)
}

type DescribeOpts struct {
	cli.OutputOpts
	id    string
	store ProjectDescriber
}

func (opts *DescribeOpts) initStore(ctx context.Context) func() error {
	return func() error {
		var err error
		opts.store, err = store.New(store.AuthenticatedPreset(config.Default()), store.WithContext(ctx))
		return err
	}
}

func (opts *DescribeOpts) Run() error {
	r, err := opts.store.Project(opts.id)
	if err != nil {
		return err
	}

	return opts.Print(r)
}

// atlas projects(s) describe <ID>.
func DescribeBuilder() *cobra.Command {
	opts := new(DescribeOpts)
	opts.Template = describeTemplate
	cmd := &cobra.Command{
		Use:     "describe <ID>",
		Aliases: []string{"show", "get"},
		Args:    require.ExactArgs(1),
		Short:   "Return the details for the specified project.",
		Long:    fmt.Sprintf(usage.RequiredRole, "Project Read Only"),
		Annotations: map[string]string{
			"IDDesc": "Unique 24-digit string that identifies the project.",
			"output": describeTemplate,
		},
		Example: `  # Return the JSON-formatted details for the project with the ID 5e2211c17a3e5a48f5497de3:
  atlas projects describe 5e2211c17a3e5a48f5497de3 --output json`,
		PreRunE: func(cmd *cobra.Command, _ []string) error {
			opts.OutWriter = cmd.OutOrStdout()
			return opts.initStore(cmd.Context())()
		},
		RunE: func(_ *cobra.Command, args []string) error {
			opts.id = args[0]
			return opts.Run()
		},
	}

	opts.AddOutputOptFlags(cmd)

	return cmd
}
