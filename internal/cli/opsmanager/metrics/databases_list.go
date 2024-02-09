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
	"context"

	"github.com/mongodb/mongodb-atlas-cli/internal/cli"
	"github.com/mongodb/mongodb-atlas-cli/internal/cli/require"
	"github.com/mongodb/mongodb-atlas-cli/internal/config"
	"github.com/mongodb/mongodb-atlas-cli/internal/flag"
	"github.com/mongodb/mongodb-atlas-cli/internal/store"
	"github.com/mongodb/mongodb-atlas-cli/internal/usage"
	"github.com/spf13/cobra"
)

type DatabasesListsOpts struct {
	cli.GlobalOpts
	cli.OutputOpts
	cli.ListOpts
	hostID string
	store  store.HostDatabaseLister
}

func (opts *DatabasesListsOpts) initStore(ctx context.Context) func() error {
	return func() error {
		var err error
		opts.store, err = store.New(store.AuthenticatedPreset(config.Default()), store.WithContext(ctx))
		return err
	}
}

var databasesListTemplate = `{{range valueOrEmptySlice .Results}}
{{.DatabaseName}}{{end}}
`

func (opts *DatabasesListsOpts) Run() error {
	listOpts := opts.NewListOptions()
	r, err := opts.store.HostDatabases(opts.ConfigProjectID(), opts.hostID, listOpts)

	if err != nil {
		return err
	}

	return opts.Print(r)
}

// mongocli om metric(s) process(es) disks lists <HOST_ID>.
func DatabasesListBuilder() *cobra.Command {
	opts := &DatabasesListsOpts{}
	cmd := &cobra.Command{
		Use:     "list <ID>",
		Short:   "List available databases for a given host.",
		Aliases: []string{"ls"},
		Args:    require.ExactArgs(1),
		Annotations: map[string]string{
			"IDDesc": "Process identifier.",
		},
		PreRunE: func(cmd *cobra.Command, _ []string) error {
			return opts.PreRunE(
				opts.ValidateProjectID,
				opts.initStore(cmd.Context()),
				opts.InitOutput(cmd.OutOrStdout(), databasesListTemplate),
			)
		},
		RunE: func(_ *cobra.Command, args []string) error {
			opts.hostID = args[0]

			return opts.Run()
		},
	}

	cmd.Flags().IntVar(&opts.PageNum, flag.Page, cli.DefaultPage, usage.Page)
	cmd.Flags().IntVar(&opts.ItemsPerPage, flag.Limit, cli.DefaultPageLimit, usage.Limit)
	cmd.Flags().StringVar(&opts.ProjectID, flag.ProjectID, "", usage.ProjectID)
	cmd.Flags().StringVarP(&opts.Output, flag.Output, flag.OutputShort, "", usage.FormatOut)
	_ = cmd.RegisterFlagCompletionFunc(flag.Output, opts.AutoCompleteOutputFlag())

	return cmd
}
