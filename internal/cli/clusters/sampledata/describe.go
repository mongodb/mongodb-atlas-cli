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

package sampledata

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

type DescribeOpts struct {
	cli.GlobalOpts
	cli.OutputOpts
	id    string
	store store.SampleDataStatusDescriber
}

func (opts *DescribeOpts) initStore(ctx context.Context) func() error {
	return func() error {
		var err error
		opts.store, err = store.New(store.AuthenticatedPreset(config.Default()), store.WithContext(ctx))
		return err
	}
}

var describeTemplate = `ID	CLUSTER NAME	STATE
{{.Id}}	{{.ClusterName}}	{{.State}}
`

func (opts *DescribeOpts) Run() error {
	r, err := opts.store.SampleDataStatus(opts.ConfigProjectID(), opts.id)
	if err != nil {
		return err
	}
	return opts.Print(r)
}

// atlas cluster(s) sampleData describe <clusterName> --projectId projectId.
func DescribeBuilder() *cobra.Command {
	opts := &DescribeOpts{}
	cmd := &cobra.Command{
		Use:     "describe <id>",
		Aliases: []string{"get"},
		Short:   "Return the details for the specified sample data load job.",
		Long:    fmt.Sprintf(usage.RequiredRole, "Project Read Only"),
		Args:    require.ExactArgs(1),
		Annotations: map[string]string{
			"idDesc": "Unique identifier of the sample data job.",
			"output": describeTemplate,
		},
		Example: `  # Return the JSON-formatted details for the sample data load job:
  atlas clusters sampleData describe 5e98249d937cfc52efdc2a9f --output json`,
		PreRunE: func(cmd *cobra.Command, _ []string) error {
			return opts.PreRunE(
				opts.ValidateProjectID,
				opts.initStore(cmd.Context()),
				opts.InitOutput(cmd.OutOrStdout(), describeTemplate),
			)
		},
		RunE: func(_ *cobra.Command, args []string) error {
			opts.id = args[0]
			return opts.Run()
		},
	}

	cmd.Flags().StringVar(&opts.ProjectID, flag.ProjectID, "", usage.ProjectID)
	cmd.Flags().StringVarP(&opts.Output, flag.Output, flag.OutputShort, "", usage.FormatOut)
	_ = cmd.RegisterFlagCompletionFunc(flag.Output, opts.AutoCompleteOutputFlag())

	return cmd
}
