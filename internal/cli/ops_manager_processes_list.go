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
	om "github.com/mongodb/go-client-mongodb-ops-manager/opsmngr"
	"github.com/mongodb/mongocli/internal/description"
	"github.com/mongodb/mongocli/internal/flags"
	"github.com/mongodb/mongocli/internal/json"
	"github.com/mongodb/mongocli/internal/store"
	"github.com/mongodb/mongocli/internal/usage"
	"github.com/spf13/cobra"
)

type opsManagerProcessesListOpts struct {
	globalOpts
	listOpts
	clusterID string
	store     store.HostLister
}

func (opts *opsManagerProcessesListOpts) initStore() error {
	var err error
	opts.store, err = store.New()
	return err
}

func (opts *opsManagerProcessesListOpts) Run() error {
	listOpts := opts.newHostListOptions()
	result, err := opts.store.Hosts(opts.ProjectID(), listOpts)

	if err != nil {
		return err
	}

	return json.PrettyPrint(result)
}

func (opts *opsManagerProcessesListOpts) newHostListOptions() *om.HostListOptions {
	return &om.HostListOptions{
		ClusterID:   opts.clusterID,
		ListOptions: *opts.newListOptions(),
	}
}

// mongocli om process(es) list --projectId projectId [--page N] [--limit N]
func OpsManagerProcessListBuilder() *cobra.Command {
	opts := &opsManagerProcessesListOpts{}
	cmd := &cobra.Command{
		Use:     "list",
		Short:   description.ListProcesses,
		Aliases: []string{"ls"},
		Hidden:  true,
		Args:    cobra.NoArgs,
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return opts.PreRunE(opts.initStore)
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return opts.Run()
		},
	}

	cmd.Flags().StringVar(&opts.clusterID, flags.ClusterID, "", usage.ClusterID)
	cmd.Flags().IntVar(&opts.pageNum, flags.Page, 0, usage.Page)
	cmd.Flags().IntVar(&opts.itemsPerPage, flags.Limit, 0, usage.Limit)

	cmd.Flags().StringVar(&opts.projectID, flags.ProjectID, "", usage.ProjectID)

	return cmd
}
