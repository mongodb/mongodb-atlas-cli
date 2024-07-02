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

package namespaces

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
	atlasv2 "go.mongodb.org/atlas-sdk/v20240530002/admin"
)

const listTemplate = `NAMESPACE	TYPE{{range valueOrEmptySlice .Namespaces}}
{{.Namespace}}	{{.Type}}{{end}}
`

type ListOpts struct {
	cli.GlobalOpts
	cli.OutputOpts
	cli.PerformanceAdvisorOpts
	store    store.PerformanceAdvisorNamespacesLister
	since    int64
	duration int64
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
	r, err := opts.store.PerformanceAdvisorNamespaces(opts.newNamespaceOptions(opts.ConfigProjectID(), host))
	if err != nil {
		return err
	}

	return opts.Print(r)
}

func (opts *ListOpts) newNamespaceOptions(project, host string) *atlasv2.ListSlowQueryNamespacesApiParams {
	params := &atlasv2.ListSlowQueryNamespacesApiParams{
		GroupId:   project,
		ProcessId: host,
	}
	if opts.since != 0 {
		params.Since = &opts.since
	}
	if opts.duration != 0 {
		params.Duration = &opts.duration
	}
	return params
}

// atlas performanceAdvisor namespace(s) list  --processName processName --since since --duration duration  --projectId projectId.
func ListBuilder() *cobra.Command {
	opts := new(ListOpts)
	cmd := &cobra.Command{
		Use:   "list",
		Short: "Return up to 20 namespaces for collections experiencing slow queries on the specified host.",
		Long: `Namespaces appear in the following format: {database}.{collection}.
		
If you don't set the duration option or the since option, this command returns data from the last 24 hours.

` + fmt.Sprintf(usage.RequiredRole, "Project Read Only"),
		Aliases: []string{"ls"},
		Args:    require.NoArgs,
		Example: `  # Return a JSON-formatted list of namespaces for collections with slow queries for the atlas-111ggi-shard-00-00.111xx.mongodb.net:27017 host in the project with the ID 5e2211c17a3e5a48f5497de3:
  atlas performanceAdvisor namespaces list --processName atlas-111ggi-shard-00-00.111xx.mongodb.net:27017 --projectId 5e2211c17a3e5a48f5497de3 --output json`,
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

	cmd.Flags().StringVar(&opts.ProjectID, flag.ProjectID, "", usage.ProjectID)
	cmd.Flags().StringVarP(&opts.Output, flag.Output, flag.OutputShort, "", usage.FormatOut)
	_ = cmd.RegisterFlagCompletionFunc(flag.Output, opts.AutoCompleteOutputFlag())

	return cmd
}

func Builder() *cobra.Command {
	const use = "namespaces"
	cmd := &cobra.Command{
		Use:     use,
		Aliases: cli.GenerateAliases(use),
		Short:   "Retrieve namespaces for collections experiencing slow queries",
	}
	cmd.AddCommand(
		ListBuilder())

	return cmd
}
