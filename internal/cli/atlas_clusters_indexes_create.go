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

package cli

import (
	"fmt"
	"strings"

	atlas "github.com/mongodb/go-client-mongodb-atlas/mongodbatlas"
	"github.com/mongodb/mongocli/internal/description"
	"github.com/mongodb/mongocli/internal/flags"
	"github.com/mongodb/mongocli/internal/store"
	"github.com/mongodb/mongocli/internal/usage"
	"github.com/spf13/cobra"
)

type atlasClustersIndexesCreateOpts struct {
	*globalOpts
	clusterName string
	name        string
	db          string
	collection  string
	keys        []string
	unique      bool
	sparse      bool
	background  bool
	store       store.IndexCreator
}

func (opts *atlasClustersIndexesCreateOpts) init() error {
	if opts.ProjectID() == "" {
		return errMissingProjectID
	}

	var err error
	opts.store, err = store.New()
	return err
}

func (opts *atlasClustersIndexesCreateOpts) Run() error {
	req, err := opts.newIndex()
	if err != nil {
		return err
	}
	if err := opts.store.CreateIndex(opts.ProjectID(), opts.clusterName, req); err != nil {
		return err
	}
	fmt.Println("Your index is being created")
	return nil
}

func (opts *atlasClustersIndexesCreateOpts) newIndex() (*atlas.IndexConfiguration, error) {
	keys, err := opts.indexKeys()
	if err != nil {
		return nil, err
	}
	i := new(atlas.IndexConfiguration)
	i.DB = opts.db
	i.Collection = opts.collection
	i.Keys = keys
	i.Options = opts.newIndexOptions()
	return i, nil
}

func (opts *atlasClustersIndexesCreateOpts) newIndexOptions() *atlas.IndexOptions {
	return &atlas.IndexOptions{
		Background: opts.background,
		Unique:     opts.unique,
		Sparse:     opts.sparse,
		Name:       opts.name,
	}
}

func (opts *atlasClustersIndexesCreateOpts) indexKeys() ([]map[string]string, error) {
	keys := make([]map[string]string, len(opts.keys))
	for i, key := range opts.keys {
		value := strings.Split(key, ":")
		if len(value) != 2 {
			return nil, fmt.Errorf("unexpected key format: %s", key)
		}
		keys[i] = map[string]string{value[0]: value[1]}
	}

	return keys, nil
}

// AtlasClustersIndexesCreateBuilder builds a cobra.Command that can run as:
// mcli atlas clusters index create [name] --clusterName clusterName  --collectionName collectionName --dbName dbName [--key field:type]
func AtlasClustersIndexesCreateBuilder() *cobra.Command {
	opts := &atlasClustersIndexesCreateOpts{
		globalOpts: newGlobalOpts(),
	}
	cmd := &cobra.Command{
		Use:   "create [name]",
		Short: description.CreateCluster,
		Args:  cobra.MaximumNArgs(1),
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return opts.init()
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) == 1 {
				opts.name = args[0]
			}
			return opts.Run()
		},
	}

	cmd.Flags().StringVar(&opts.clusterName, flags.ClusterName, "", usage.ClusterName)
	cmd.Flags().StringVar(&opts.db, flags.Database, "", usage.Database)
	cmd.Flags().StringVar(&opts.collection, flags.Collection, "", usage.Collection)
	cmd.Flags().StringArrayVar(&opts.keys, flags.Key, nil, usage.Key)
	cmd.Flags().BoolVar(&opts.unique, flags.Unique, false, usage.Unique)
	cmd.Flags().BoolVar(&opts.sparse, flags.Sparse, false, usage.Sparse)
	cmd.Flags().BoolVar(&opts.background, flags.Background, false, usage.Background)

	cmd.Flags().StringVar(&opts.projectID, flags.ProjectID, "", usage.ProjectID)

	_ = cmd.MarkFlagRequired(flags.ClusterName)
	_ = cmd.MarkFlagRequired(flags.Database)
	_ = cmd.MarkFlagRequired(flags.Collection)
	_ = cmd.MarkFlagRequired(flags.Key)

	return cmd
}
