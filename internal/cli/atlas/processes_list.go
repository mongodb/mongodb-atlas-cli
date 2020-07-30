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

package atlas

import (
	"github.com/mongodb/mongocli/internal/cli"
	"github.com/mongodb/mongocli/internal/config"
	"github.com/mongodb/mongocli/internal/description"
	"github.com/mongodb/mongocli/internal/flag"
	"github.com/mongodb/mongocli/internal/output"
	"github.com/mongodb/mongocli/internal/store"
	"github.com/mongodb/mongocli/internal/usage"
	"github.com/spf13/cobra"
	atlas "go.mongodb.org/atlas/mongodbatlas"
)

const listTemplate = `ID	REPLICA SET NAME	SHARD NAME	VERSION{{range .}}
{{.ID}}	{{.ReplicaSetName}}	{{.ShardName}}	{{.Version}}{{end}}
`

type ProcessesListOpts struct {
	cli.GlobalOpts
	cli.ListOpts
	clusterID string
	store     store.ProcessLister
}

func (opts *ProcessesListOpts) initStore() error {
	var err error
	opts.store, err = store.New(config.Default())
	return err
}

func (opts *ProcessesListOpts) Run() error {
	listOpts := opts.newProcessesListOptions()
	r, err := opts.store.Processes(opts.ConfigProjectID(), listOpts)
	if err != nil {
		return err
	}

	return output.Print(config.Default(), listTemplate, r)
}

func (opts *ProcessesListOpts) newProcessesListOptions() *atlas.ProcessesListOptions {
	return &atlas.ProcessesListOptions{
		ClusterID:   opts.clusterID,
		ListOptions: *opts.NewListOptions(),
	}
}

// mongocli atlas process(es) list --projectId projectId [--page N] [--limit N]
func ProcessListBuilder() *cobra.Command {
	opts := &ProcessesListOpts{}
	cmd := &cobra.Command{
		Use:     "list",
		Short:   description.ListProcesses,
		Aliases: []string{"ls"},
		Args:    cobra.NoArgs,
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return opts.PreRunE(opts.initStore)
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return opts.Run()
		},
	}

	cmd.Flags().StringVar(&opts.clusterID, flag.ClusterID, "", usage.ClusterID)
	cmd.Flags().IntVar(&opts.PageNum, flag.Page, 0, usage.Page)
	cmd.Flags().IntVar(&opts.ItemsPerPage, flag.Limit, 0, usage.Limit)

	cmd.Flags().StringVar(&opts.ProjectID, flag.ProjectID, "", usage.ProjectID)

	return cmd
}
