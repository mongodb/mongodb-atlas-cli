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
	"errors"
	"fmt"
	"strings"

	"github.com/mongodb/mongocli/internal/cli"
	"github.com/mongodb/mongocli/internal/config"
	"github.com/mongodb/mongocli/internal/description"
	"github.com/mongodb/mongocli/internal/flag"
	"github.com/mongodb/mongocli/internal/json"
	"github.com/mongodb/mongocli/internal/store"
	"github.com/mongodb/mongocli/internal/usage"
	"github.com/spf13/cobra"
	atlas "go.mongodb.org/atlas/mongodbatlas"
)

type CreateOpts struct {
	cli.GlobalOpts
	clusterName    string
	name           string
	dbName         string
	collection     string
	analyzer       string
	searchAnalyzer string
	dynamic        bool
	fields         []string
	store          store.SearchIndexCreator
}

func (opts *CreateOpts) initStore() error {
	var err error
	opts.store, err = store.New(config.Default())
	return err
}

func (opts *CreateOpts) Run() error {
	index, err := opts.newSearchIndex()
	if err != nil {
		return err
	}
	result, err := opts.store.CreateSearchIndexes(opts.ConfigProjectID(), opts.clusterName, index)
	if err != nil {
		return err
	}

	return json.PrettyPrint(result)
}

func (opts *CreateOpts) newSearchIndex() (*atlas.SearchIndex, error) {
	f, err := opts.indexFields()
	if err != nil {
		return nil, err
	}
	i := &atlas.SearchIndex{
		Analyzer:       opts.analyzer,
		CollectionName: opts.collection,
		Database:       opts.dbName,
		Mappings: &atlas.IndexMapping{
			Dynamic: opts.dynamic,
			Fields:  &f,
		},
		Name:           opts.name,
		SearchAnalyzer: opts.searchAnalyzer,
	}
	return i, nil
}
func (opts *CreateOpts) indexFields() (map[string]atlas.IndexField, error) {
	if len(opts.fields) == 0 {
		return nil, nil
	}
	fields := make(map[string]atlas.IndexField, len(opts.fields))
	for _, p := range opts.fields {
		f := strings.Split(p, ":")
		if len(f) != 3 {
			return nil, fmt.Errorf("partition should be fieldName:analyzer:fieldType, got: %s", p)
		}
		fields[f[0]] = atlas.IndexField{
			Analyzer: f[1],
			Type:     f[2],
		}
	}
	return fields, nil
}

// CreateBuilder
// Create an online archive for a cluster.
//
// Usage:
//   mongocli atlas clusters search create <name> [flags]
//
// Flags:
//      --analyzer string         Analyzer to use when creating the index (default "lucene.standard")
//      --clusterName string      Name of the cluster.
//      --collection string       Collection name.
//      --db string               Database name.
//      --dynamic                 Indicates whether the index uses dynamic or static mappings.
//      --field strings           Static field specifications.
//  -h, --help                    help for create
//      --projectId string        Project ID to use. Overrides configuration file or environment variable settings.
//      --searchAnalyzer string   Analyzer to use when searching the index. (default "lucene.standard")
//
// Global Flags:
//  -P, --profile string   Profile to use from your configuration file.
func CreateBuilder() *cobra.Command {
	opts := &CreateOpts{}
	cmd := &cobra.Command{
		Use:   "create <name>",
		Short: description.CreateOnlineArchive,
		Args:  cobra.ExactArgs(1),
		PreRunE: func(cmd *cobra.Command, args []string) error {
			if !opts.dynamic && len(opts.fields) == 0 {
				return errors.New("you need to specify fields for the index or use a dynamic index")
			}
			if opts.dynamic && len(opts.fields) > 0 {
				return errors.New("you can't specify fields and dynamic at the same time")
			}
			return opts.PreRunE(opts.initStore)
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			opts.name = args[0]
			return opts.Run()
		},
	}

	cmd.Flags().StringVar(&opts.clusterName, flag.ClusterName, "", usage.ClusterName)
	cmd.Flags().StringVar(&opts.dbName, flag.Database, "", usage.Database)
	cmd.Flags().StringVar(&opts.collection, flag.Collection, "", usage.Collection)
	cmd.Flags().StringVar(&opts.analyzer, flag.Analyzer, "lucene.standard", usage.Analyzer)
	cmd.Flags().StringVar(&opts.searchAnalyzer, flag.SearchAnalyzer, "lucene.standard", usage.SearchAnalyzer)
	cmd.Flags().BoolVar(&opts.dynamic, flag.Dynamic, false, usage.Dynamic)
	cmd.Flags().StringSliceVar(&opts.fields, flag.Field, nil, usage.SearchFields)

	cmd.Flags().StringVar(&opts.ProjectID, flag.ProjectID, "", usage.ProjectID)

	_ = cmd.MarkFlagRequired(flag.ClusterName)
	_ = cmd.MarkFlagRequired(flag.Database)
	_ = cmd.MarkFlagRequired(flag.Collection)

	return cmd
}
