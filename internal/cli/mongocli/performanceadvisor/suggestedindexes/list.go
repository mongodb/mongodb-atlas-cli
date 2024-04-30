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

package suggestedindexes

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

const listTemplate = `ID	NAMESPACE	SUGGESTED INDEX{{range valueOrEmptySlice .SuggestedIndexes}}  
{{ .ID }}	{{ .Namespace}}	{ {{range $i, $element := .Index}}{{range $key, $value := .}}{{if $i}}, {{end}}{{ $key }}: {{ $value }}{{end}}{{end}} }{{end}}
`

type ListOpts struct {
	cli.GlobalOpts
	cli.OutputOpts
	HostID     string
	store      store.PerformanceAdvisorIndexesLister
	since      int64
	duration   int64
	namespaces string
	nIndexes   int64
	nExamples  int64
}

func (opts *ListOpts) initStore(ctx context.Context) func() error {
	return func() error {
		var err error
		opts.store, err = store.New(store.AuthenticatedPreset(config.Default()), store.WithContext(ctx))
		return err
	}
}

func (opts *ListOpts) Run() error {
	r, err := opts.store.PerformanceAdvisorIndexes(opts.ConfigProjectID(), opts.HostID, opts.newSuggestedIndexOptions())
	if err != nil {
		return err
	}

	return opts.Print(r)
}

func (opts *ListOpts) newSuggestedIndexOptions() *opsmngr.SuggestedIndexOptions {
	return &opsmngr.SuggestedIndexOptions{
		Namespaces: opts.namespaces,
		NIndexes:   opts.nIndexes,
		NExamples:  opts.nExamples,
		NamespaceOptions: opsmngr.NamespaceOptions{
			Since:    opts.since,
			Duration: opts.duration,
		},
	}
}

// mongocli atlas performanceAdvisor suggestedIndexes list  --processName processName --nIndexes nIndexes --nExamples nExamples --namespaces namespaces --since since --duration duration  --projectId projectId.
func ListBuilder() *cobra.Command {
	opts := new(ListOpts)
	cmd := &cobra.Command{
		Use:   "list",
		Short: "Return the suggested indexes for collections experiencing slow queries.",
		Long: `The Performance Advisor monitors queries that MongoDB considers slow and suggests new indexes to improve query performance.

` + fmt.Sprintf(usage.RequiredRole, "Project Read Only"),
		Aliases: []string{"ls"},
		Args:    require.NoArgs,
		Example: fmt.Sprintf(`  # Return a JSON-formatted list of suggested indexes for the atlas-111ggi-shard-00-00.111xx.mongodb.net:27017 host in the project with the ID 5e2211c17a3e5a48f5497de3:
  %s performanceAdvisor suggestedIndexes list --processName atlas-111ggi-shard-00-00.111xx.mongodb.net:27017 --projectId 5e2211c17a3e5a48f5497de3 --output json`, cli.ExampleEntryPoint()),
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

	cmd.Flags().StringVar(&opts.HostID, flag.HostID, "", usage.HostID)
	cmd.Flags().Int64Var(&opts.since, flag.Since, 0, usage.Since)
	cmd.Flags().Int64Var(&opts.duration, flag.Duration, 0, usage.Duration)
	cmd.Flags().StringVar(&opts.namespaces, flag.Namespaces, "", usage.SuggestedIndexNamespaces)
	cmd.Flags().Int64Var(&opts.nExamples, flag.NExamples, 0, usage.NExamples)
	cmd.Flags().Int64Var(&opts.nIndexes, flag.NIndexes, 0, usage.NIndexes)

	cmd.Flags().StringVar(&opts.ProjectID, flag.ProjectID, "", usage.ProjectID)
	cmd.Flags().StringVarP(&opts.Output, flag.Output, flag.OutputShort, "", usage.FormatOut)
	_ = cmd.RegisterFlagCompletionFunc(flag.Output, opts.AutoCompleteOutputFlag())

	_ = cmd.MarkFlagRequired(flag.HostID)

	return cmd
}
