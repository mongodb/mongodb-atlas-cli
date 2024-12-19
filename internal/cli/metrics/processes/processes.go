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

package processes

import (
	"context"
	"fmt"
	"strconv"

	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli/require"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/config"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/convert"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/flag"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/pointer"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/store"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/usage"
	"github.com/spf13/cobra"
	atlasv2 "go.mongodb.org/atlas-sdk/v20241113004/admin"
)

type Opts struct {
	cli.ProjectOpts
	cli.OutputOpts
	cli.MetricsOpts
	cli.ListOpts
	host  string
	port  int
	store store.ProcessMeasurementLister
}

func (opts *Opts) initStore(ctx context.Context) func() error {
	return func() error {
		var err error
		opts.store, err = store.New(store.AuthenticatedPreset(config.Default()), store.WithContext(ctx))
		return err
	}
}

func (opts *Opts) NewProcessMeasurementsAPIParams(groupID string, processID string) *atlasv2.GetHostMeasurementsApiParams {
	p := &atlasv2.GetHostMeasurementsApiParams{
		GroupId:   groupID,
		ProcessId: processID,
	}
	if opts.Granularity != "" {
		p.Granularity = &opts.Granularity
	}
	if len(opts.MeasurementType) > 0 {
		p.M = &opts.MeasurementType
	}
	if opts.Period != "" {
		p.Period = &opts.Period
	}
	if start, err := convert.ParseTimestamp(opts.Start); err == nil {
		p.Start = pointer.Get(start)
	}
	if end, err := convert.ParseTimestamp(opts.End); err == nil {
		p.End = pointer.Get(end)
	}
	return p
}

func (opts *Opts) Run() error {
	processID := opts.host + ":" + strconv.Itoa(opts.port)
	params := opts.NewProcessMeasurementsAPIParams(opts.ConfigProjectID(), processID)

	r, err := opts.store.ProcessMeasurements(params)
	if err != nil {
		return err
	}

	return opts.Print(r)
}

var metricTemplate = `NAME	UNITS	TIMESTAMP		VALUE{{range valueOrEmptySlice .Measurements}} {{if .DataPoints}}
{{- $name := .Name }}{{- $unit := .Units }}{{- range valueOrEmptySlice .DataPoints}}	
{{ $name }}	{{ $unit }}	{{.Timestamp}}	{{if .Value }}	{{ .Value }}{{else}}	N/A {{end}}{{end}}{{end}}{{end}}
`

// atlas metric(s) process(es) <hostname:port> [--granularity granularity] [--period period] [--start start] [--end end] [--type type][--projectId projectId].
func Builder() *cobra.Command {
	opts := &Opts{}
	cmd := &cobra.Command{
		Use:   "processes <hostname:port>",
		Short: "Return the process measurements for the specified host.",
		Long: `To return the hostname and port needed for this command, run
atlas processes list

` + fmt.Sprintf(usage.RequiredRole, "Project Read Only"),
		Aliases: []string{"process"},
		Args:    require.ExactArgs(1),
		Annotations: map[string]string{
			"hostname:portDesc": "Hostname and port number of the instance running the MongoDB process.",
		},
		Example: `  # Return the JSON-formatted process metrics for the host atlas-lnmtkm-shard-00-00.ajlj3.mongodb.net:27017
  atlas metrics processes atlas-lnmtkm-shard-00-00.ajlj3.mongodb.net:27017 --output json`,
		PreRunE: func(cmd *cobra.Command, _ []string) error {
			return opts.PreRunE(
				opts.ValidateProjectID,
				opts.ValidatePeriodStartEnd,
				opts.initStore(cmd.Context()),
				opts.InitOutput(cmd.OutOrStdout(), metricTemplate),
			)
		},
		RunE: func(_ *cobra.Command, args []string) error {
			var err error
			opts.host, opts.port, err = cli.GetHostnameAndPort(args[0])
			if err != nil {
				return err
			}

			return opts.Run()
		},
	}
	opts.AddListOptsFlagsWithoutOmitCount(cmd)

	opts.AddMetricsOptsFlags(cmd)
	cmd.Flag(flag.TypeFlag).Usage = usage.MetricsMeasurementType

	opts.AddProjectOptsFlags(cmd)
	opts.AddOutputOptFlags(cmd)

	_ = cmd.MarkFlagRequired(flag.Granularity)

	cmd.MarkFlagsRequiredTogether(flag.Start, flag.End)
	cmd.MarkFlagsMutuallyExclusive(flag.Period, flag.Start)
	cmd.MarkFlagsMutuallyExclusive(flag.Period, flag.End)

	return cmd
}
