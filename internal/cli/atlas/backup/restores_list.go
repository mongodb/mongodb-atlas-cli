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

package backup

import (
	"github.com/mongodb/mongocli/internal/cli"
	"github.com/mongodb/mongocli/internal/cli/require"
	"github.com/mongodb/mongocli/internal/config"
	"github.com/mongodb/mongocli/internal/flag"
	"github.com/mongodb/mongocli/internal/store"
	"github.com/mongodb/mongocli/internal/usage"
	"github.com/spf13/cobra"
)

type RestoresListOpts struct {
	cli.GlobalOpts
	cli.OutputOpts
	cli.ListOpts
	clusterName string
	store       store.RestoreJobsLister
}

func (opts *RestoresListOpts) initStore() error {
	var err error
	opts.store, err = store.New(config.Default())
	return err
}

var restoreListTemplate = `ID	SNAPSHOT	CLUSTER	TYPE	EXPIRES AT{{range .Results}}
{{.ID}}	{{.SnapshotID}}	{{.TargetClusterName}}	{{.DeliveryType}}	{{.ExpiresAt}}{{end}}
`

func (opts *RestoresListOpts) Run() error {
	listOpts := opts.NewListOptions()
	r, err := opts.store.RestoreJobs(opts.ConfigProjectID(), opts.clusterName, listOpts)
	if err != nil {
		return err
	}

	return opts.Print(r)
}

// mongocli atlas backup(s) restore(s) job(s) list <clusterName> [--page N] [--limit N]
func RestoresListBuilder() *cobra.Command {
	opts := new(RestoresListOpts)
	cmd := &cobra.Command{
		Use:     "list <clusterName>",
		Aliases: []string{"ls"},
		Short:   listRestores,
		Args:    require.ExactArgs(1),
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return opts.PreRunE(
				opts.ValidateProjectID,
				opts.initStore,
				opts.InitOutput(cmd.OutOrStdout(), restoreListTemplate),
			)
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			opts.clusterName = args[0]

			return opts.Run()
		},
	}

	cmd.Flags().IntVar(&opts.PageNum, flag.Page, 0, usage.Page)
	cmd.Flags().IntVar(&opts.ItemsPerPage, flag.Limit, 0, usage.Limit)

	cmd.Flags().StringVar(&opts.ProjectID, flag.ProjectID, "", usage.ProjectID)
	cmd.Flags().StringVarP(&opts.Output, flag.Output, flag.OutputShort, "", usage.FormatOut)

	return cmd
}
