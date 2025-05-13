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
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/log"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/store"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/telemetry"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/usage"
	"github.com/spf13/afero"
	"github.com/spf13/cobra"
	atlasv2 "go.mongodb.org/atlas-sdk/v20250312002/admin"
)

//go:generate go tool go.uber.org/mock/mockgen -typed -destination=update_mock_test.go -package=search . Updater

type Updater interface {
	UpdateSearchIndexesDeprecated(string, string, string, *atlasv2.ClusterSearchIndex) (*atlasv2.ClusterSearchIndex, error)
	UpdateSearchIndexes(string, string, string, *atlasv2.SearchIndexUpdateRequest) (*atlasv2.SearchIndexResponse, error)
}

type UpdateOpts struct {
	cli.ProjectOpts
	cli.OutputOpts
	IndexOpts
	id          string
	clusterName string
	store       Updater
}

func (opts *UpdateOpts) initStore(ctx context.Context) func() error {
	return func() error {
		var err error
		opts.store, err = store.New(store.AuthenticatedPreset(config.Default()), store.WithContext(ctx))
		return err
	}
}

var updateTemplate = "Index {{.Name}} updated.\n"

func (opts *UpdateOpts) Run() error {
	i, err := opts.UpdateSearchIndex()
	if err != nil {
		return err
	}

	switch index := i.(type) {
	case *atlasv2.SearchIndexUpdateRequest:
		r, err := opts.store.UpdateSearchIndexes(opts.ConfigProjectID(), opts.clusterName, opts.id, index)
		if err != nil {
			return err
		}

		return opts.Print(r)
	case *atlasv2.ClusterSearchIndex:
		_, _ = log.Warningln("you're using an old search index definition")
		telemetry.AppendOption(telemetry.WithSearchIndexType(index.GetType()))
		r, err := opts.store.UpdateSearchIndexesDeprecated(opts.ConfigProjectID(), opts.clusterName, opts.id, index)
		if err != nil {
			return err
		}

		return opts.Print(r)
	default:
		return errInvalidIndex
	}
}

// UpdateBuilder
// Update a search index for a cluster.
//
// Usage:
//
//	atlas clusters search indexes update <ID> [flags]
//
// Flags:
//
//	    --analyzer string         Analyzer to use when creating the index (default "lucene.standard")
//	    --clusterName string      Name of the cluster.
//	    --collection string       Collection name.
//	    --db string               Database name.
//	    --dynamic                 Indicates whether the index uses dynamic or static mappings.
//	    --field strings           Static field specifications.
//	-h, --help                    help for update
//	    --indexName string        Name of the cluster.
//	    --projectId string        Project ID to use. Overrides configuration file or environment variable settings.
//	    --searchAnalyzer string   Analyzer to use when searching the index. (default "lucene.standard")
//	-f, --file string             JSON file to use in order to update the index
//
// Global Flags:
//
//	-P, --profile string   Profile to use from your configuration file.
func UpdateBuilder() *cobra.Command {
	opts := &UpdateOpts{}
	opts.Fs = afero.NewOsFs()

	cmd := &cobra.Command{
		Use:   "update <indexId>",
		Short: "Modify a search index for a cluster.",
		Long:  fmt.Sprintf(usage.RequiredRole, "Project Data Access Admin"),
		Annotations: map[string]string{
			"indexIdDesc": "ID of the index.",
			"output":      updateTemplate,
		},
		Args: require.ExactArgs(1),
		Example: `  # Modify the search index with the ID 5f2099cd683fc55fbb30bef6 for the cluster named myCluster:
  atlas clusters search indexes update 5f2099cd683fc55fbb30bef6 --clusterName myCluster --output json`,
		PreRunE: func(cmd *cobra.Command, _ []string) error {
			if opts.Filename == "" {
				_ = cmd.MarkFlagRequired(flag.IndexName)
				_ = cmd.MarkFlagRequired(flag.Database)
				_ = cmd.MarkFlagRequired(flag.Collection)
			}

			return opts.PreRunE(
				opts.validateOpts,
				opts.ValidateProjectID,
				opts.initStore(cmd.Context()),
				opts.InitOutput(cmd.OutOrStdout(), updateTemplate),
			)
		},
		RunE: func(_ *cobra.Command, args []string) error {
			opts.id = args[0]
			return opts.Run()
		},
	}

	cmd.Flags().StringVar(&opts.clusterName, flag.ClusterName, "", usage.ClusterName)
	cmd.Flags().StringVar(&opts.Name, flag.IndexName, "", usage.IndexName)
	cmd.Flags().StringVar(&opts.DBName, flag.Database, "", usage.Database)
	cmd.Flags().StringVar(&opts.Collection, flag.Collection, "", usage.Collection)
	cmd.Flags().StringVar(&opts.Analyzer, flag.Analyzer, DefaultAnalyzer, usage.Analyzer)
	cmd.Flags().StringVar(&opts.SearchAnalyzer, flag.SearchAnalyzer, DefaultAnalyzer, usage.SearchAnalyzer)
	cmd.Flags().BoolVar(&opts.Dynamic, flag.Dynamic, false, usage.Dynamic)
	cmd.Flags().StringSliceVar(&opts.fields, flag.Field, nil, usage.SearchFields+usage.UpdateWarning)
	cmd.Flags().StringVarP(&opts.Filename, flag.File, flag.FileShort, "", usage.SearchFilename)

	_ = cmd.MarkFlagFilename(flag.File)

	opts.AddProjectOptsFlags(cmd)
	opts.AddOutputOptFlags(cmd)

	_ = cmd.MarkFlagRequired(flag.ClusterName)

	_ = cmd.Flags().MarkDeprecated(flag.IndexName, deprecatedFlagMessage)
	_ = cmd.Flags().MarkDeprecated(flag.Database, deprecatedFlagMessage)
	_ = cmd.Flags().MarkDeprecated(flag.Collection, deprecatedFlagMessage)
	_ = cmd.Flags().MarkDeprecated(flag.Analyzer, deprecatedFlagMessage)
	_ = cmd.Flags().MarkDeprecated(flag.SearchAnalyzer, deprecatedFlagMessage)
	_ = cmd.Flags().MarkDeprecated(flag.Dynamic, deprecatedFlagMessage)
	_ = cmd.Flags().MarkDeprecated(flag.Field, deprecatedFlagMessage)

	return cmd
}
