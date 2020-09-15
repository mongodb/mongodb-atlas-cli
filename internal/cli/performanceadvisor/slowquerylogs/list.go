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
package slowquerylogs

import (
	"github.com/mongodb/mongocli/internal/cli"
	"github.com/mongodb/mongocli/internal/config"
	"github.com/mongodb/mongocli/internal/flag"
	"github.com/mongodb/mongocli/internal/store"
	"github.com/mongodb/mongocli/internal/usage"
	"github.com/spf13/cobra"
	atlas "go.mongodb.org/atlas/mongodbatlas"
)

const listTemplate = `NAMESPACE	LINE{{range .SlowQuery}}
{{.Namespace}}	{{.Line}}{{end}}
`

type ListOpts struct {
	cli.GlobalOpts
	cli.OutputOpts
	cli.PerformanceAdvisorOpts
	store      store.PerformanceAdvisorSlowQueriesLister
	since      int64
	duration   int64
	namespaces string
	nLog       int64
}

func (opts *ListOpts) initStore() error {
	var err error
	opts.store, err = store.New(config.Default())
	return err
}

func (opts *ListOpts) Run() error {
	host, err := opts.host()
	if err != nil {
		return err
	}
	r, err := opts.store.PerformanceAdvisorSlowQueries(opts.ConfigProjectID(), host, opts.newSlowQueryOptions())
	if err != nil {
		return err
	}

	return opts.Print(r)
}

func (opts *ListOpts) newSlowQueryOptions() *atlas.SlowQueryOptions {
	return &atlas.SlowQueryOptions{
		Namespaces: opts.namespaces,
		NLogs:      opts.nLog,
		NamespaceOptions: atlas.NamespaceOptions{
			Since:    opts.since,
			Duration: opts.duration,
		},
	}
}

func (opts *ListOpts) host() (string, error) {
	if opts.ProcessName == "" {
		return opts.HostID, nil
	}
	err := opts.ValidateProcessName()
	if err != nil {
		return "", err
	}
	return opts.ProcessName, nil
}

// mongocli atlas performanceAdvisor slowQueryLogs list  --processName processName --since since --duration duration  --projectId projectId
func ListBuilder() *cobra.Command {
	opts := new(ListOpts)
	opts.Template = listTemplate
	cmd := &cobra.Command{
		Use:     "list",
		Short:   list,
		Aliases: []string{"ls"},
		Args:    cobra.NoArgs,
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return opts.PreRunE(
				opts.initStore,
				opts.InitOutput(cmd.OutOrStdout(), listTemplate),
				opts.MarkProcessNameHostIDRequired(cmd),
			)
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return opts.Run()
		},
	}

	cmd.Flags().StringVar(&opts.HostID, flag.HostID, "", usage.HostID)
	cmd.Flags().StringVar(&opts.ProcessName, flag.ProcessName, "", usage.ProcessName)
	cmd.Flags().Int64Var(&opts.since, flag.Since, 0, usage.Since)
	cmd.Flags().Int64Var(&opts.duration, flag.Duration, 0, usage.Duration)
	cmd.Flags().Int64Var(&opts.nLog, flag.NLog, 20000, usage.NLog)
	cmd.Flags().StringVar(&opts.namespaces, flag.Namespaces, "", usage.Namespaces)

	cmd.Flags().StringVar(&opts.ProjectID, flag.ProjectID, "", usage.ProjectID)
	cmd.Flags().StringVarP(&opts.Output, flag.Output, flag.OutputShort, "", usage.FormatOut)

	return cmd
}
