// Copyright 2021 MongoDB Inc
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

package serverless

import (
	"context"
	"fmt"

	"github.com/mongodb/atlas-cli-core/config"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli/require"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/store"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/usage"
	"github.com/spf13/cobra"
	atlasv2 "go.mongodb.org/atlas-sdk/v20250312010/admin"
)

var listTemplate = `ID	NAME	MDB VER	STATE{{range valueOrEmptySlice .Results}}
{{.Id}}	{{.Name}}	{{.MongoDBVersion}}	{{.StateName}}{{end}}
`

//go:generate go tool go.uber.org/mock/mockgen -typed -destination=list_mock_test.go -package=serverless . Lister

type Lister interface {
	ServerlessInstances(string, *store.ListOptions) (*atlasv2.PaginatedServerlessInstanceDescription, error)
}

type ListOpts struct {
	cli.ProjectOpts
	cli.OutputOpts
	cli.ListOpts
	store Lister
}

func (opts *ListOpts) initStore(ctx context.Context) func() error {
	return func() error {
		var err error
		opts.store, err = store.New(store.AuthenticatedPreset(config.Default()), store.WithContext(ctx))
		return err
	}
}

func (opts *ListOpts) Run() error {
	listOpts := opts.NewAtlasListOptions()
	r, err := opts.store.ServerlessInstances(opts.ConfigProjectID(), listOpts)
	if err != nil {
		return err
	}

	return opts.Print(r)
}

// atlas serverless|sl list [--projectId projectId] [--page N] [--limit N].
func ListBuilder() *cobra.Command {
	opts := &ListOpts{}
	cmd := &cobra.Command{
		Use:     "list",
		Short:   "Return all serverless instances in the specified project.",
		Long:    fmt.Sprintf(usage.RequiredRole, "Project Read Only"),
		Aliases: []string{"ls"},
		Args:    require.NoArgs,
		Example: `  # Return the JSON-formatted list of all serverlessInstances for the project with the ID 5e2211c17a3e5a48f5497de3:
  atlas serverless list --projectId 5e2211c17a3e5a48f5497de3`,
		Annotations: map[string]string{
			"output": listTemplate,
		},
		PreRunE: func(cmd *cobra.Command, _ []string) error {
			return opts.PreRunE(
				opts.ValidateProjectID,
				opts.initStore(cmd.Context()),
				opts.InitOutput(cmd.OutOrStdout(), listTemplate),
			)
		},
		RunE: func(_ *cobra.Command, _ []string) error {
			return opts.Run()
		},
		Deprecated: "please use the 'atlas cluster list --tier FLEX [--projectId projectId]' command instead. For the migration guide and timeline, visit: https://dochub.mongodb.org/core/flex-migration",
	}

	opts.AddListOptsFlags(cmd)

	opts.AddProjectOptsFlags(cmd)
	opts.AddOutputOptFlags(cmd)

	return cmd
}
