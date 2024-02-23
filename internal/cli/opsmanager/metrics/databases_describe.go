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

type DatabasesDescribeOpts struct {
	cli.GlobalOpts
	cli.OutputOpts
	cli.MetricsOpts
	hostID string
	name   string
	store  store.HostDatabaseMeasurementsLister
}

func (opts *DatabasesDescribeOpts) initStore(ctx context.Context) func() error {
	return func() error {
		var err error
		opts.store, err = store.New(store.AuthenticatedPreset(config.Default()), store.WithContext(ctx))
		return err
	}
}

var databasesMetricTemplate = `NAME	UNITS	TIMESTAMP		VALUE{{range valueOrEmptySlice .ProcessMeasurements.Measurements}}  {{if .DataPoints}}
{{- $name := .Name }}{{- $unit := .Units }}{{- range valueOrEmptySlice .DataPoints}}	
{{ $name }}	{{ $unit }}	{{.Timestamp}}	{{if .Value }}	{{ .Value }}{{else}}	N/A {{end}}{{end}}{{end}}{{end}}
`

func (opts *DatabasesDescribeOpts) Run() error {
	listOpts := opts.NewProcessMetricsListOptions()
	r, err := opts.store.HostDatabaseMeasurements(opts.ConfigProjectID(), opts.hostID, opts.name, listOpts)
	if err != nil {
		return err
	}

	return opts.Print(r)
}

// mcli om metric(s) database(s) describe <hostId:port> <name> --granularity g --period p --start start --end end [--type type] [--projectId projectId].
func DatabasesDescribeBuilder() *cobra.Command {
	const argsN = 2
	opts := &DatabasesDescribeOpts{}
	cmd := &cobra.Command{
		Use:   "describe <hostId> <name>",
		Short: "Describe database measurements for a given host database.",
		Args:  require.ExactArgs(argsN),
		Example: `# List metrics for the database test of the process e4ac1e57c58cc9c8aaa5a1163a851993
  mongocli ops-manager metrics database describe e4ac1e57c58cc9c8aaa5a1163a851993 test --period P1DT12H --granularity PT5`,
		Annotations: map[string]string{
			"hostIdDesc": "Process identifier. You can use mongocli ops-manager processes list to get the ID.",
			"nameDesc":   "Database name.",
		},
		PreRunE: func(cmd *cobra.Command, _ []string) error {
			return opts.PreRunE(
				opts.ValidateProjectID,
				opts.ValidatePeriodStartEnd,
				opts.initStore(cmd.Context()),
				opts.InitOutput(cmd.OutOrStdout(), databasesMetricTemplate),
			)
		},
		RunE: func(_ *cobra.Command, args []string) error {
			opts.hostID = args[0]
			opts.name = args[1]

			return opts.Run()
		},
	}

	cmd.Flags().IntVar(&opts.PageNum, flag.Page, cli.DefaultPage, usage.Page)
	cmd.Flags().IntVar(&opts.ItemsPerPage, flag.Limit, cli.DefaultPageLimit, usage.Limit)

	cmd.Flags().StringVar(&opts.Granularity, flag.Granularity, "", usage.Granularity)
	cmd.Flags().StringVar(&opts.Period, flag.Period, "", usage.Period)
	cmd.Flags().StringVar(&opts.Start, flag.Start, "", usage.MeasurementStart)
	cmd.Flags().StringVar(&opts.End, flag.End, "", usage.MeasurementEnd)
	cmd.Flags().StringSliceVar(&opts.MeasurementType, flag.TypeFlag, nil, usage.MeasurementType)

	cmd.Flags().StringVar(&opts.ProjectID, flag.ProjectID, "", usage.ProjectID)
	cmd.Flags().StringVarP(&opts.Output, flag.Output, flag.OutputShort, "", usage.FormatOut)
	_ = cmd.RegisterFlagCompletionFunc(flag.Output, opts.AutoCompleteOutputFlag())

	_ = cmd.MarkFlagRequired(flag.Granularity)

	cmd.MarkFlagsRequiredTogether(flag.Start, flag.End)
	cmd.MarkFlagsMutuallyExclusive(flag.Period, flag.Start)
	cmd.MarkFlagsMutuallyExclusive(flag.Period, flag.End)

	return cmd
}
