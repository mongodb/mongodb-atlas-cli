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
	"github.com/mongodb/mongocli/internal/cli/require"
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
	host  string
	port  int
	store store.ProcessDisksLister
}

func (opts *DisksListsOpts) initStore() error {
	var err error
	opts.store, err = store.New(config.Default())
	return err
}

func (opts *DisksListsOpts) Run() error {
	listOpts := opts.NewListOptions()
	r, err := opts.store.ProcessDisks(opts.ConfigProjectID(), opts.host, opts.port, listOpts)
	if err != nil {
		return err
	}

	return opts.Print(r)
}

var listTemplate = `{{range .Results}}
{{.PartitionName}}{{end}}
`

// mongocli atlas metric(s) process(es) disks lists <hostname:port>
func DisksListBuilder() *cobra.Command {
	opts := &DisksListsOpts{}
	cmd := &cobra.Command{
		Use:     "list <hostname:port>",
		Short:   listDisks,
		Aliases: []string{"ls"},
		Args:    require.ExactArgs(1),
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return opts.PreRunE(
				opts.ValidateProjectID,
				opts.initStore,
				opts.InitOutput(cmd.OutOrStdout(), listTemplate),
			)
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			var err error
			opts.host, opts.port, err = getHostnameAndPort(args[0])
			if err != nil {
				return err
			}

			return opts.Run()
		},
	}

	cmd.Flags().IntVar(&opts.PageNum, flag.Page, 0, usage.Page)
	cmd.Flags().IntVar(&opts.ItemsPerPage, flag.Limit, 0, usage.Limit)
	cmd.Flags().StringVar(&opts.ProjectID, flag.ProjectID, "", usage.ProjectID)
	cmd.Flags().StringVarP(&opts.Output, flag.Output, flag.OutputShort, "", usage.FormatOut)

	return cmd
}
