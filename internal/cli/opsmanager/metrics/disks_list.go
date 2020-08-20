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

package metrics

import (
	"github.com/mongodb/mongocli/internal/cli"
	"github.com/mongodb/mongocli/internal/config"
	"github.com/mongodb/mongocli/internal/flag"
	"github.com/mongodb/mongocli/internal/store"
	"github.com/mongodb/mongocli/internal/usage"
	"github.com/spf13/cobra"
)

type DisksListsOpts struct {
	cli.GlobalOpts
	cli.OutputOpts
	cli.ListOpts
	hostID string
	store  store.HostDisksLister
}

func (opts *DisksListsOpts) initStore() error {
	var err error
	opts.store, err = store.New(config.Default())
	return err
}

var listTemplate = `{{range .Results}}
{{.PartitionName}}{{end}}
`

func (opts *DisksListsOpts) Run() error {
	listOpts := opts.NewListOptions()
	r, err := opts.store.HostDisks(opts.ConfigProjectID(), opts.hostID, listOpts)
	if err != nil {
		return err
	}

	return opts.Print(r)
}

// mongocli om metric(s) process(es) disks lists <ID>
func DisksListBuilder() *cobra.Command {
	opts := &DisksListsOpts{}
	cmd := &cobra.Command{
		Use:     "list <ID>",
		Short:   ListDisks,
		Aliases: []string{"ls"},
		Args:    cobra.ExactArgs(1),
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return opts.PreRunE(
				opts.initStore,
				opts.InitOutput(cmd.OutOrStdout(), listTemplate),
			)
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			opts.hostID = args[0]

			return opts.Run()
		},
	}

	cmd.Flags().IntVar(&opts.PageNum, flag.Page, 0, usage.Page)
	cmd.Flags().IntVar(&opts.ItemsPerPage, flag.Limit, 0, usage.Limit)
	cmd.Flags().StringVar(&opts.ProjectID, flag.ProjectID, "", usage.ProjectID)

	return cmd
}
