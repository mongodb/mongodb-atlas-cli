// Copyright 2023 MongoDB Inc
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

package slowquerylogs

import (
	"context"
	"fmt"

	"github.com/mongodb/mongodb-atlas-cli/mongocli/v2/internal/cli"
	"github.com/mongodb/mongodb-atlas-cli/mongocli/v2/internal/cli/require"
	"github.com/mongodb/mongodb-atlas-cli/mongocli/v2/internal/config"
	"github.com/mongodb/mongodb-atlas-cli/mongocli/v2/internal/flag"
	"github.com/mongodb/mongodb-atlas-cli/mongocli/v2/internal/store"
	"github.com/mongodb/mongodb-atlas-cli/mongocli/v2/internal/usage"
	"github.com/spf13/cobra"
	"go.mongodb.org/ops-manager/opsmngr"
)

const listTemplate = `NAMESPACE	LINE{{range valueOrEmptySlice .SlowQuery}}
{{.Namespace}}	{{.Line}}{{end}}
`

type ListOpts struct {
	cli.GlobalOpts
	cli.OutputOpts
	HostID     string
	store      store.PerformanceAdvisorSlowQueriesLister
	since      int64
	duration   int64
	namespaces string
	nLog       int64
}

func (opts *ListOpts) initStore(ctx context.Context) func() error {
	return func() error {
		var err error
		opts.store, err = store.New(store.AuthenticatedPreset(config.Default()), store.WithContext(ctx))
		return err
	}
}

func (opts *ListOpts) Run() error {
	r, err := opts.store.PerformanceAdvisorSlowQueries(opts.ConfigProjectID(), opts.HostID, opts.newSlowQueryOptions())
	if err != nil {
		return err
	}

	return opts.Print(r)
}

func (opts *ListOpts) newSlowQueryOptions() *opsmngr.SlowQueryOptions {
	return &opsmngr.SlowQueryOptions{
		Namespaces: opts.namespaces,
		NLogs:      opts.nLog,
		NamespaceOptions: opsmngr.NamespaceOptions{
			Since:    opts.since,
			Duration: opts.duration,
		},
	}
}

// mongocli atlas performanceAdvisor slowQueryLogs list  --processName processName --since since --duration duration  --projectId projectId.
func ListBuilder() *cobra.Command {
	opts := new(ListOpts)
	opts.Template = listTemplate
	cmd := &cobra.Command{
		Use:   "list",
		Short: "Return log lines for slow queries that the Performance Advisor and Query Profiler identified.",
		Long: `The Performance Advisor monitors queries that MongoDB considers slow and suggests new indexes to improve query performance. The threshold for slow queries varies based on the average time of operations on your cluster to provide recommendations pertinent to your workload.
		
If you don't set the duration option or the since option, this command returns data from the last 24 hours.

` + fmt.Sprintf(usage.RequiredRole, "Project Data Access Read/Write"),
		Aliases: []string{"ls"},
		Args:    require.NoArgs,
		Example: fmt.Sprintf(`  # Return a JSON-formatted list of log lines for collections with slow queries for the atlas-111ggi-shard-00-00.111xx.mongodb.net:27017 host in the project with the ID 5e2211c17a3e5a48f5497de3:
  %s performanceAdvisor slowQueryLogs list --processName atlas-111ggi-shard-00-00.111xx.mongodb.net:27017 --projectId 5e2211c17a3e5a48f5497de3 --output json`, cli.ExampleEntryPoint()),
		PreRunE: func(cmd *cobra.Command, _ []string) error {
			return opts.PreRunE(
				opts.ValidateProjectID,
				opts.initStore(cmd.Context()),
				opts.InitOutput(cmd.OutOrStdout(), listTemplate),
			)
		},
		RunE: func(_ *cobra.Command, _ []string) error {
			return opts.Run()
		},
	}

	const defaultLogLines = 20000

	cmd.Flags().StringVar(&opts.HostID, flag.HostID, "", usage.HostID)
	cmd.Flags().Int64Var(&opts.since, flag.Since, 0, usage.Since)
	cmd.Flags().Int64Var(&opts.duration, flag.Duration, 0, usage.Duration)
	cmd.Flags().Int64Var(&opts.nLog, flag.NLog, defaultLogLines, usage.NLog)
	cmd.Flags().StringVar(&opts.namespaces, flag.Namespaces, "", usage.SlowQueryNamespaces)

	cmd.Flags().StringVar(&opts.ProjectID, flag.ProjectID, "", usage.ProjectID)
	cmd.Flags().StringVarP(&opts.Output, flag.Output, flag.OutputShort, "", usage.FormatOut)
	_ = cmd.RegisterFlagCompletionFunc(flag.Output, opts.AutoCompleteOutputFlag())
	_ = cmd.MarkFlagRequired(flag.HostID)

	return cmd
}
