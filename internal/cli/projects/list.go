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

	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli/require"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/config"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/store"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/usage"
	"github.com/spf13/cobra"
	atlasv2 "go.mongodb.org/atlas-sdk/v20250312002/admin"
)

const listTemplate = `ID	NAME{{range valueOrEmptySlice .Results}}
{{.Id}}	{{.Name}}{{end}}
`

//go:generate mockgen -typed -destination=list_mock_test.go -package=projects . OrgProjectLister

type OrgProjectLister interface {
	Projects(*store.ListOptions) (*atlasv2.PaginatedAtlasGroup, error)
	GetOrgProjects(string, *store.ListOptions) (*atlasv2.PaginatedAtlasGroup, error)
}

type ListOpts struct {
	cli.OrgOpts
	cli.OutputOpts
	cli.ListOpts
	store OrgProjectLister
}

func (opts *ListOpts) initStore(ctx context.Context) func() error {
	return func() error {
		var err error
		opts.store, err = store.New(store.AuthenticatedPreset(config.Default()), store.WithContext(ctx))
		return err
	}
}

func (opts *ListOpts) Run() error {
	listOptions := opts.NewAtlasListOptions()
	var r any
	var err error
	if opts.OrgID != "" {
		r, err = opts.store.GetOrgProjects(opts.OrgID, listOptions)
	} else {
		r, err = opts.store.Projects(listOptions)
	}
	if err != nil {
		return err
	}
	return opts.Print(r)
}

// atlas project(s) list [--orgId orgId].
func ListBuilder() *cobra.Command {
	opts := &ListOpts{}
	opts.Template = listTemplate
	cmd := &cobra.Command{
		Use:     "list",
		Aliases: []string{"ls"},
		Annotations: map[string]string{
			"output": listTemplate,
		},
		Short: "Return all projects.",
		Long:  fmt.Sprintf(usage.RequiredRole, "Project Data Access Read/Write"),
		Args:  require.NoArgs,
		Example: `  # Return a JSON-formatted list of all projects:
  atlas projects list --output json`,
		PreRunE: func(cmd *cobra.Command, _ []string) error {
			opts.OutWriter = cmd.OutOrStdout()
			return opts.initStore(cmd.Context())()
		},
		RunE: func(_ *cobra.Command, _ []string) error {
			return opts.Run()
		},
	}

	opts.AddListOptsFlags(cmd)

	opts.AddOrgOptFlags(cmd)
	opts.AddOutputOptFlags(cmd)

	return cmd
}
