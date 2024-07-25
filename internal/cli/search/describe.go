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

	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli/require"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/config"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/flag"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/store"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/usage"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/validate"
	"github.com/spf13/cobra"
)

type DescribeOpts struct {
	cli.GlobalOpts
	cli.OutputOpts
	clusterName string
	indexID     string
	store       store.SearchIndexDescriber
}

func (opts *DescribeOpts) initStore(ctx context.Context) func() error {
	return func() error {
		var err error
		opts.store, err = store.New(store.AuthenticatedPreset(config.Default()), store.WithContext(ctx))
		return err
	}
}

var describeTemplate = `ID	NAME	DATABASE	COLLECTION	TYPE
{{.IndexID}}	{{.Name}}	{{.Database}}	{{.CollectionName}}	{{if .Type }}{{.Type}}{{else}}` + DefaultType + `{{end}}
`

func (opts *DescribeOpts) Run() error {
	r, err := opts.store.SearchIndex(opts.ConfigProjectID(), opts.clusterName, opts.indexID)
	if err != nil {
		return err
	}

	return opts.Print(r)
}

// atlas cluster(s) search indexes describe <ID> [--clusterName name][--projectId projectId].
func DescribeBuilder() *cobra.Command {
	opts := &DescribeOpts{}
	cmd := &cobra.Command{
		Use:   "describe <indexId>",
		Short: "Return the details for the search index for a cluster.",
		Long:  fmt.Sprintf(usage.RequiredRole, "Project Data Access Read/Write"),
		Args:  require.ExactArgs(1),
		Annotations: map[string]string{
			"indexIdDesc": "ID of the index.",
			"output":      describeTemplate,
		},
		Example: `  # Return the JSON-formatted details for the search index with the ID 5f1f40842f2ac35f49190c20 for the cluster named myCluster:
  atlas clusters search indexes describe 5f1f40842f2ac35f49190c20 --clusterName myCluster --output json`,
		PreRunE: func(cmd *cobra.Command, _ []string) error {
			return opts.PreRunE(
				opts.ValidateProjectID,
				opts.initStore(cmd.Context()),
				opts.InitOutput(cmd.OutOrStdout(), describeTemplate),
			)
		},
		RunE: func(_ *cobra.Command, args []string) error {
			if err := validate.ObjectID(args[0]); err != nil {
				return err
			}
			opts.indexID = args[0]

			return opts.Run()
		},
	}

	cmd.Flags().StringVar(&opts.clusterName, flag.ClusterName, "", usage.ClusterName)

	cmd.Flags().StringVar(&opts.ProjectID, flag.ProjectID, "", usage.ProjectID)
	cmd.Flags().StringVarP(&opts.Output, flag.Output, flag.OutputShort, "", usage.FormatOut)
	_ = cmd.RegisterFlagCompletionFunc(flag.Output, opts.AutoCompleteOutputFlag())

	_ = cmd.MarkFlagRequired(flag.ClusterName)

	return cmd
}
