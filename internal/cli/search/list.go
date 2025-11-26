// Copyright 2020 MongoDB Inc
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

package search

import (
	"context"
	"fmt"

	"github.com/mongodb/atlas-cli-core/config"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli/require"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/flag"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/store"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/usage"
	"github.com/spf13/cobra"
	atlasv2 "go.mongodb.org/atlas-sdk/v20250312010/admin"
)

//go:generate go tool go.uber.org/mock/mockgen -typed -destination=list_mock_test.go -package=search . Lister

type Lister interface {
	SearchIndexesDeprecated(string, string, string, string) ([]atlasv2.ClusterSearchIndex, error)
}

type ListOpts struct {
	cli.ProjectOpts
	cli.OutputOpts
	cli.ListOpts
	clusterName string
	dbName      string
	collName    string
	store       Lister
}

func (opts *ListOpts) initStore(ctx context.Context) func() error {
	return func() error {
		var err error
		opts.store, err = store.New(store.AuthenticatedPreset(config.Default()), store.WithContext(ctx))
		return err
	}
}

var listTemplate = `ID	NAME	DATABASE	COLLECTION	TYPE{{range valueOrEmptySlice .}}
{{.IndexID}}	{{.Name}}	{{.Database}}	{{.CollectionName}}	{{if .Type }}{{.Type}}{{else}}` + DefaultType + `{{end}}{{end}}
`

func (opts *ListOpts) Run() error {
	r, err := opts.store.SearchIndexesDeprecated(opts.ConfigProjectID(), opts.clusterName, opts.dbName, opts.collName)
	if err != nil {
		return err
	}

	return opts.Print(r)
}

// atlas clusters search(s) list [--projectId projectId] [--clusterName name][--db database][--collection collName].
func ListBuilder() *cobra.Command {
	opts := &ListOpts{}
	cmd := &cobra.Command{
		Use:   "list",
		Short: "List all Atlas Search indexes for a cluster.",
		Long:  fmt.Sprintf(usage.RequiredRole, "Project Data Access Read/Write"),
		Annotations: map[string]string{
			"output": listTemplate,
		},
		Aliases: []string{"ls"},
		Args:    require.NoArgs,
		Example: `  # Return the JSON-formatted list of Atlas search indexes on the sample_mflix.movies database in the cluster named myCluster:
  atlas clusters search indexes list --clusterName myCluster --db sample_mflix --collection movies --output json`,
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
	}

	cmd.Flags().StringVar(&opts.clusterName, flag.ClusterName, "", usage.ClusterName)
	cmd.Flags().StringVar(&opts.dbName, flag.Database, "", usage.Database)
	cmd.Flags().StringVar(&opts.collName, flag.Collection, "", usage.Collection)

	opts.AddListOptsFlagsWithoutOmitCount(cmd)
	_ = cmd.Flags().MarkDeprecated(flag.Page, deprecatedFlagMessage)
	_ = cmd.Flags().MarkDeprecated(flag.Limit, deprecatedFlagMessage)

	opts.AddProjectOptsFlags(cmd)
	opts.AddOutputOptFlags(cmd)

	_ = cmd.MarkFlagRequired(flag.ClusterName)
	_ = cmd.MarkFlagRequired(flag.Database)
	_ = cmd.MarkFlagRequired(flag.Collection)

	return cmd
}
