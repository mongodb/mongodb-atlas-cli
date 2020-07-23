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
	"github.com/mongodb/mongocli/internal/cli"
	"github.com/mongodb/mongocli/internal/config"
	"github.com/mongodb/mongocli/internal/description"
	"github.com/mongodb/mongocli/internal/flag"
	"github.com/mongodb/mongocli/internal/output"
	"github.com/mongodb/mongocli/internal/store"
	"github.com/mongodb/mongocli/internal/usage"
	"github.com/spf13/cobra"
)

type ListOpts struct {
	cli.GlobalOpts
	cli.ListOpts
	clusterName string
	dbName      string
	collName    string
	store       store.SearchIndexLister
}

func (opts *ListOpts) initStore() error {
	var err error
	opts.store, err = store.New(config.Default())
	return err
}

func (opts *ListOpts) Run() error {
	r, err := opts.store.SearchIndexes(opts.ConfigProjectID(), opts.clusterName, opts.dbName, opts.collName, opts.NewListOptions())
	if err != nil {
		return err
	}

	return output.Print(config.Default(), "", r)
}

// mongocli atlas clusters search(s) list [--projectId projectId] [--clusterName name][--db database][--collection collName]
func ListBuilder() *cobra.Command {
	opts := &ListOpts{}
	cmd := &cobra.Command{
		Use:     "list",
		Short:   description.ListSearchIndexes,
		Aliases: []string{"ls"},
		Args:    cobra.NoArgs,
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return opts.initStore()
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return opts.Run()
		},
	}

	cmd.Flags().StringVar(&opts.clusterName, flag.ClusterName, "", usage.ClusterName)
	cmd.Flags().StringVar(&opts.dbName, flag.Database, "", usage.Database)
	cmd.Flags().StringVar(&opts.collName, flag.Collection, "", usage.Collection)

	cmd.Flags().IntVar(&opts.PageNum, flag.Page, 0, usage.Page)
	cmd.Flags().IntVar(&opts.ItemsPerPage, flag.Limit, 0, usage.Limit)
	cmd.Flags().StringVar(&opts.ProjectID, flag.ProjectID, "", usage.ProjectID)

	_ = cmd.MarkFlagRequired(flag.ClusterName)
	_ = cmd.MarkFlagRequired(flag.Database)
	_ = cmd.MarkFlagRequired(flag.Collection)

	return cmd
}
