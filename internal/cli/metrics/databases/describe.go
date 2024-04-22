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
	"strconv"

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
	cli.MetricsOpts
	host  string
	port  int
	name  string
	store store.ProcessDatabaseMeasurementsLister
}

func (opts *DescribeOpts) initStore(ctx context.Context) func() error {
	return func() error {
		var err error
		opts.store, err = store.New(store.AuthenticatedPreset(config.Default()), store.WithContext(ctx))
		return err
	}
}

var databasesMetricTemplate = `NAME	UNITS	TIMESTAMP		VALUE{{range valueOrEmptySlice .Measurements}}  {{if .DataPoints}}
{{- $name := .Name }}{{- $unit := .Units }}{{- range valueOrEmptySlice .DataPoints}}	
{{ $name }}	{{ $unit }}	{{.Timestamp}}	{{if .Value }}	{{ .Value }}{{else}}	N/A {{end}}{{end}}{{end}}{{end}}
`

func (opts *DescribeOpts) Run() error {
	processID := opts.host + ":" + strconv.Itoa(opts.port)
	params := opts.NewDatabaseMeasurementsAPIParams(opts.ConfigProjectID(), processID, opts.name)

	r, err := opts.store.ProcessDatabaseMeasurements(params)
	if err != nil {
		return err
	}

	return opts.Print(r)
}

// mcli atlas metric(s) database(s) describe <host:port> <databaseName> --granularity g --period p --start start --end end [--type type] [--projectId projectId].
func DescribeBuilder() *cobra.Command {
	const argsN = 2
	opts := &DescribeOpts{}
	cmd := &cobra.Command{
		Use: "describe <hostname:port> <databaseName>",
		Long: `To return the hostname and port needed for this command, run
atlas process list`,
		Short: "Return the measurements of a database on the specified host.",
		Args:  require.ExactArgs(argsN),
		Annotations: map[string]string{
			"hostname:portDesc": "Hostname and port number of the instance running the MongoDB process.",
			"databaseNameDesc":  "Label that identifies the database from which you want to retrieve metrics.",
		},
		Example: `  # Return the JSON-formatted database metrics from the last 36 hours with 5-minute granularity for the database named testDB in the host atlas-lnmtkm-shard-00-00.ajlj3.mongodb.net:27017 
  atlas metrics databases describe atlas-lnmtkm-shard-00-00.ajlj3.mongodb.net:27017 testDB --granularity PT1M --period P1DT12H --output json`,
		PreRunE: func(cmd *cobra.Command, _ []string) error {
			return opts.PreRunE(
				opts.ValidateProjectID,
				opts.ValidatePeriodStartEnd,
				opts.initStore(cmd.Context()),
				opts.InitOutput(cmd.OutOrStdout(), databasesMetricTemplate),
			)
		},
		RunE: func(_ *cobra.Command, args []string) error {
			var err error
			opts.host, opts.port, err = cli.GetHostnameAndPort(args[0])
			if err != nil {
				return err
			}
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
