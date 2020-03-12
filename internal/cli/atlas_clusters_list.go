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
	atlas "github.com/mongodb/go-client-mongodb-atlas/mongodbatlas"
	"github.com/mongodb/mongocli/internal/flags"
	"github.com/mongodb/mongocli/internal/json"
	"github.com/mongodb/mongocli/internal/store"
	"github.com/mongodb/mongocli/internal/usage"
	"github.com/spf13/cobra"
)

type atlasClustersListOpts struct {
	*globalOpts
	pageNum      int
	itemsPerPage int
	storeAT      store.ClusterLister
	storeOM      store.ListAllClusters
}

func (opts *atlasClustersListOpts) init() error {
	var err error

	if opts.ProjectID() == "" {
		opts.storeOM, err = store.New()
	} else {
		opts.storeAT, err = store.New()
	}

	return err
}

func (opts *atlasClustersListOpts) RunAT() error {
	listOpts := opts.newListOptions()
	result, err := opts.storeAT.ProjectClusters(opts.ProjectID(), listOpts)

	if err != nil {
		return err
	}

	return json.PrettyPrint(result)
}

func (opts *atlasClustersListOpts) RunOM() error {
	result, err := opts.storeOM.ListAllClustersProjects()
	if err != nil {
		return err
	}
	return json.PrettyPrint(result)
}

func (opts *atlasClustersListOpts) newListOptions() *atlas.ListOptions {
	return &atlas.ListOptions{
		PageNum:      opts.pageNum,
		ItemsPerPage: opts.itemsPerPage,
	}
}

// mongocli atlas cluster(s) list --projectId projectId [--page N] [--limit N]
func AtlasClustersListBuilder() *cobra.Command {
	opts := &atlasClustersListOpts{
		globalOpts: newGlobalOpts(),
	}
	cmd := &cobra.Command{
		Use:     "list",
		Short:   "List Atlas clusters for a project.",
		Aliases: []string{"ls"},
		Args:    cobra.NoArgs,
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return opts.init()
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			if opts.storeAT != nil {
				return opts.RunAT()
			}
			return opts.RunOM()
		},
	}

	cmd.Flags().IntVar(&opts.pageNum, flags.Page, 0, usage.Page)
	cmd.Flags().IntVar(&opts.itemsPerPage, flags.Limit, 0, usage.Limit)

	cmd.Flags().StringVar(&opts.projectID, flags.ProjectID, "", usage.ProjectID)

	return cmd
}
