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

	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli/processes"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli/require"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/config"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/flag"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/store"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/usage"
	"github.com/spf13/cobra"
	atlasv2 "go.mongodb.org/atlas-sdk/v20240530002/admin"
)

const listTemplate = `ID	NAMESPACE	SUGGESTED INDEX{{range valueOrEmptySlice .SuggestedIndexes}}  
{{ .Id }}	{{ .Namespace}}	{ {{range $i, $element := valueOrEmptySlice .Index}}{{range $key, $value := .}}{{if $i}}, {{end}}{{ $key }}: {{ $value }}{{end}}{{end}} }{{end}}
`

type ListOpts struct {
	cli.GlobalOpts
	cli.OutputOpts
	cli.PerformanceAdvisorOpts
	store      store.PerformanceAdvisorIndexesLister
	since      int64
	duration   int64
	namespaces []string
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
	host, err := opts.Host()
	if err != nil {
		return err
	}
	r, err := opts.store.PerformanceAdvisorIndexes(opts.newSuggestedIndexOptions(opts.ConfigProjectID(), host))
	if err != nil {
		return err
	}

	return opts.Print(r)
}

func (opts *ListOpts) newSuggestedIndexOptions(project, host string) *atlasv2.ListSuggestedIndexesApiParams {
	params := &atlasv2.ListSuggestedIndexesApiParams{
		GroupId:   project,
		ProcessId: host,
	}
	if opts.since != 0 {
		params.Since = &opts.since
	}
	if opts.duration != 0 {
		params.Duration = &opts.duration
	}
	if opts.nExamples != 0 {
		params.NExamples = &opts.nExamples
	}
	if opts.nIndexes != 0 {
		params.NIndexes = &opts.nIndexes
	}
	if len(opts.namespaces) > 0 {
		params.Namespaces = &opts.namespaces
	}
	return params
}

// atlas performanceAdvisor suggestedIndexes list  --processName processName --nIndexes nIndexes --nExamples nExamples --namespaces namespaces --since since --duration duration  --projectId projectId.
func ListBuilder() *cobra.Command {
	opts := new(ListOpts)
	cmd := &cobra.Command{
		Use:   "list",
		Short: "Return the suggested indexes for collections experiencing slow queries.",
		Long: `The Performance Advisor monitors queries that MongoDB considers slow and suggests new indexes to improve query performance.

` + fmt.Sprintf(usage.RequiredRole, "Project Read Only"),
		Aliases: []string{"ls"},
		Args:    require.NoArgs,
		Example: `  # Return a JSON-formatted list of suggested indexes for the atlas-111ggi-shard-00-00.111xx.mongodb.net:27017 host in the project with the ID 5e2211c17a3e5a48f5497de3:
  atlas performanceAdvisor suggestedIndexes list --processName atlas-111ggi-shard-00-00.111xx.mongodb.net:27017 --projectId 5e2211c17a3e5a48f5497de3 --output json`,
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
	_ = cmd.Flags().MarkDeprecated(flag.HostID, "Flag is invalid for MongoDB Atlas")
	cmd.Flags().StringVar(&opts.ProcessName, flag.ProcessName, "", usage.ProcessNameAtlasCLI)
	_ = cmd.MarkFlagRequired(flag.ProcessName)
	cmd.Flags().Int64Var(&opts.since, flag.Since, 0, usage.Since)
	cmd.Flags().Int64Var(&opts.duration, flag.Duration, 0, usage.Duration)
	cmd.Flags().StringSliceVar(&opts.namespaces, flag.Namespaces, []string{}, usage.SuggestedIndexNamespaces)
	cmd.Flags().Int64Var(&opts.nExamples, flag.NExamples, 0, usage.NExamples)
	cmd.Flags().Int64Var(&opts.nIndexes, flag.NIndexes, 0, usage.NIndexes)

	cmd.Flags().StringVar(&opts.ProjectID, flag.ProjectID, "", usage.ProjectID)
	cmd.Flags().StringVarP(&opts.Output, flag.Output, flag.OutputShort, "", usage.FormatOut)
	_ = cmd.RegisterFlagCompletionFunc(flag.Output, opts.AutoCompleteOutputFlag())

	autocomplete := &processes.AutoCompleteOpts{}
	_ = cmd.RegisterFlagCompletionFunc(flag.ProcessName, autocomplete.AutocompleteProcesses())

	return cmd
}

func Builder() *cobra.Command {
	const use = "suggestedIndexes"
	cmd := &cobra.Command{
		Use:     use,
		Aliases: cli.GenerateAliases(use),
		Short:   "Get suggested indexes for collections experiencing slow queries",
	}
	cmd.AddCommand(
		ListBuilder())

	return cmd
}
