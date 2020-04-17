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
	"github.com/mongodb/mongocli/internal/description"
	"github.com/mongodb/mongocli/internal/flags"
	"github.com/mongodb/mongocli/internal/json"
	"github.com/mongodb/mongocli/internal/store"
	"github.com/mongodb/mongocli/internal/usage"
	"github.com/spf13/cobra"
)

type atlasMeasurementsDatabasesListsOpts struct {
	globalOpts
	listOpts
	host  string
	port  int
	store store.ProcessDatabaseLister
}

func (opts *atlasMeasurementsDatabasesListsOpts) initStore() error {
	var err error
	opts.store, err = store.New()
	return err
}

func (opts *atlasMeasurementsDatabasesListsOpts) Run() error {
	listOpts := opts.newListOptions()
	result, err := opts.store.ProcessDatabases(opts.ProjectID(), opts.host, opts.port, listOpts)

	if err != nil {
		return err
	}

	return json.PrettyPrint(result)
}

// mongocli atlas measurements process(es) disks lists [host:port]
func AtlasMeasurementsDatabasesListBuilder() *cobra.Command {
	opts := &atlasMeasurementsDatabasesListsOpts{}
	cmd := &cobra.Command{
		Use:     "list [host:port]",
		Short:   description.ListDatabases,
		Aliases: []string{"ls"},
		Args:    cobra.ExactArgs(1),
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return opts.PreRunE(opts.initStore)
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			var err error
			if opts.host, opts.port, err = getHostNameAndPort(args[0]); err != nil {
				return err
			}

			return opts.Run()
		},
	}

	cmd.Flags().IntVar(&opts.pageNum, flags.Page, 0, usage.Page)
	cmd.Flags().IntVar(&opts.itemsPerPage, flags.Limit, 0, usage.Limit)
	cmd.Flags().StringVar(&opts.projectID, flags.ProjectID, "", usage.ProjectID)

	return cmd
}
