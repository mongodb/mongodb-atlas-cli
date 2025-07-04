// Copyright 2021 MongoDB Inc
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//	http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package serverless

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

var describeTemplate = `ID	NAME	MDB VER	STATE
{{.Id}}	{{.Name}}	{{.MongoDBVersion}}	{{.StateName}}
`

//go:generate go tool go.uber.org/mock/mockgen -typed -destination=describe_mock_test.go -package=serverless . Describer

type Describer interface {
	GetServerlessInstance(string, string) (*atlasv2.ServerlessInstanceDescription, error)
}

type DescribeOpts struct {
	cli.ProjectOpts
	cli.OutputOpts
	store        Describer
	instanceName string
}

func (opts *DescribeOpts) initStore(ctx context.Context) func() error {
	return func() error {
		var err error
		opts.store, err = store.New(store.AuthenticatedPreset(config.Default()), store.WithContext(ctx))
		return err
	}
}

func (opts *DescribeOpts) Run() error {
	r, err := opts.store.GetServerlessInstance(opts.ConfigProjectID(), opts.instanceName)
	if err != nil {
		return err
	}

	return opts.Print(r)
}

// atlas serverless|sl describe <instanceName> --projectId projectId.
func DescribeBuilder() *cobra.Command {
	opts := new(DescribeOpts)
	cmd := &cobra.Command{
		Use:   "describe <instanceName>",
		Short: "Return one serverless instance in the specified project.",
		Long:  fmt.Sprintf(usage.RequiredRole, "Project Read Only"),
		Args:  require.ExactArgs(1),
		Example: `  # Return the JSON-formatted details for the serverlessInstance named myInstance in the project with the ID 5e2211c17a3e5a48f5497de3:
  atlas serverless describe myInstance --projectId 5e2211c17a3e5a48f5497de3`,
		Annotations: map[string]string{
			"instanceNameDesc": "Human-readable label that identifies your serverless instance.",
			"output":           describeTemplate,
		},
		Aliases: []string{"get"},
		PreRunE: func(cmd *cobra.Command, args []string) error {
			opts.instanceName = args[0]
			return opts.PreRunE(
				opts.ValidateProjectID,
				opts.initStore(cmd.Context()),
				opts.InitOutput(cmd.OutOrStdout(), describeTemplate),
			)
		},
		RunE: func(_ *cobra.Command, _ []string) error {
			return opts.Run()
		},
		Deprecated: "please use the 'atlas cluster describe <ClusterName> [--projectId projectId]' command instead. For the migration guide and timeline, visit: https://dochub.mongodb.org/core/flex-migration",
	}

	opts.AddProjectOptsFlags(cmd)
	opts.AddOutputOptFlags(cmd)

	return cmd
}
