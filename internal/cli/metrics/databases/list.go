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

package databases

import (
	"context"
	"fmt"

	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli/require"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/config"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/flag"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/store"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/usage"
	"github.com/spf13/cobra"
)

type ListsOpts struct {
	cli.GlobalOpts
	cli.OutputOpts
	cli.MetricsOpts
	cli.ListOpts
	host  string
	port  int
	store store.ProcessDatabaseLister
}

func (opts *ListsOpts) initStore(ctx context.Context) func() error {
	return func() error {
		var err error
		opts.store, err = store.New(store.AuthenticatedPreset(config.Default()), store.WithContext(ctx))
		return err
	}
}

var databasesListTemplate = `DATABASE NAME{{range valueOrEmptySlice .Results}}
{{.DatabaseName}}{{end}}
`

func (opts *ListsOpts) Run() error {
	listOpts := opts.NewListOptions()
	r, err := opts.store.ProcessDatabases(opts.ConfigProjectID(), opts.host, opts.port, listOpts)
	if err != nil {
		return err
	}

	return opts.Print(r)
}

// atlas metric(s) database(s) lists <hostname:port>.
func ListBuilder() *cobra.Command {
	opts := &ListsOpts{}
	cmd := &cobra.Command{
		Use: "list <hostname:port>",
		Long: `To return the hostname and port needed for this command, run
atlas processes list

` + fmt.Sprintf(usage.RequiredRole, "Project Read Only"),
		Short:   "Return all databases running on the specified host for your project.",
		Aliases: []string{"ls"},
		Annotations: map[string]string{
			"hostname:portDesc": "Hostname and port number of the instance running the MongoDB process.",
			"output":            databasesListTemplate,
		},
		Example: `  # Return a JSON-formatted list of available databases for the host atlas-lnmtkm-shard-00-00.ajlj3.mongodb.net:27017
  atlas metrics databases list atlas-lnmtkm-shard-00-00.ajlj3.mongodb.net:27017 --output json`,
		Args: require.ExactArgs(1),
		PreRunE: func(cmd *cobra.Command, _ []string) error {
			return opts.PreRunE(
				opts.ValidateProjectID,
				opts.initStore(cmd.Context()),
				opts.InitOutput(cmd.OutOrStdout(), databasesListTemplate),
			)
		},
		RunE: func(_ *cobra.Command, args []string) error {
			var err error
			if opts.host, opts.port, err = cli.GetHostnameAndPort(args[0]); err != nil {
				return err
			}

			return opts.Run()
		},
	}

	cmd.Flags().IntVar(&opts.PageNum, flag.Page, cli.DefaultPage, usage.Page)
	cmd.Flags().IntVar(&opts.ItemsPerPage, flag.Limit, cli.DefaultPageLimit, usage.Limit)
	cmd.Flags().BoolVar(&opts.OmitCount, flag.OmitCount, false, usage.OmitCount)

	cmd.Flags().StringVar(&opts.ProjectID, flag.ProjectID, "", usage.ProjectID)
	cmd.Flags().StringVarP(&opts.Output, flag.Output, flag.OutputShort, "", usage.FormatOut)
	_ = cmd.RegisterFlagCompletionFunc(flag.Output, opts.AutoCompleteOutputFlag())

	return cmd
}
