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

package disks

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
	atlasv2 "go.mongodb.org/atlas-sdk/v20250312005/admin"
)

//go:generate go tool go.uber.org/mock/mockgen -typed -destination=describe_mock_test.go -package=disks . ProcessDiskMeasurementsLister

type ProcessDiskMeasurementsLister interface {
	ProcessDiskMeasurements(*atlasv2.GetDiskMeasurementsApiParams) (*atlasv2.ApiMeasurementsGeneralViewAtlas, error)
}

type DescribeOpts struct {
	cli.ProjectOpts
	cli.OutputOpts
	cli.ListOpts
	cli.MetricsOpts
	host  string
	port  int
	name  string
	store ProcessDiskMeasurementsLister
}

func (opts *DescribeOpts) initStore(ctx context.Context) func() error {
	return func() error {
		var err error
		opts.store, err = store.New(store.AuthenticatedPreset(config.Default()), store.WithContext(ctx))
		return err
	}
}

func (opts *DescribeOpts) NewDiskMeasurementsAPIParams(groupID string, processID string, partitionName string) *atlasv2.GetDiskMeasurementsApiParams {
	p := &atlasv2.GetDiskMeasurementsApiParams{
		GroupId:       groupID,
		ProcessId:     processID,
		PartitionName: partitionName,
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

func (opts *DescribeOpts) Run() error {
	processID := opts.host + ":" + strconv.Itoa(opts.port)
	params := opts.NewDiskMeasurementsAPIParams(opts.ConfigProjectID(), processID, opts.name)

	r, err := opts.store.ProcessDiskMeasurements(params)
	if err != nil {
		return err
	}

	return opts.Print(r)
}

var diskMetricTemplate = `NAME	UNITS	TIMESTAMP		VALUE{{range valueOrEmptySlice .Measurements}}  {{if .DataPoints}}
{{- $name := .Name }}{{- $unit := .Units }}{{- range valueOrEmptySlice .DataPoints}}	
{{ $name }}	{{ $unit }}	{{.Timestamp}}	{{if .Value }}	{{ .Value }}{{else}}	N/A {{end}}{{end}}{{end}}{{end}}
`

// mcli atlas metric(s) disk(s) describe <host:port> <diskName> --granularity g --period p --start start --end end [--type type] [--projectId projectId].
func DescribeBuilder() *cobra.Command {
	const argsN = 2
	opts := &DescribeOpts{}
	cmd := &cobra.Command{
		Use: "describe <hostname:port> <diskName>",
		Long: `To return the hostname and port needed for this command, run
atlas processes list

` + fmt.Sprintf(usage.RequiredRole, "Project Read Only"),
		Short: "Return the measurements of a disk or partition on the specified host.",
		Args:  require.ExactArgs(argsN),
		Annotations: map[string]string{
			"hostname:portDesc": "Hostname and port number of the instance running the MongoDB process.",
			"diskNameDesc":      "Label that identifies the disk or partition from which you want to retrieve metrics.",
		},
		Example: `  # Return the JSON-formatted disk metrics from the last 36 hours with 5-minute granularity for the database named testDB in the host atlas-lnmtkm-shard-00-00.ajlj3.mongodb.net:27017
  atlas metrics disks describe atlas-lnmtkm-shard-00-00.ajlj3.mongodb.net:27017 testDB --granularity PT1M --period P1DT12H --output json`,
		PreRunE: func(cmd *cobra.Command, _ []string) error {
			return opts.PreRunE(
				opts.ValidateProjectID,
				opts.ValidatePeriodStartEnd,
				opts.initStore(cmd.Context()),
				opts.InitOutput(cmd.OutOrStdout(), diskMetricTemplate),
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

	opts.AddListOptsFlagsWithoutOmitCount(cmd)

	opts.AddMetricsOptsFlags(cmd)

	opts.AddProjectOptsFlags(cmd)
	opts.AddOutputOptFlags(cmd)

	_ = cmd.MarkFlagRequired(flag.Granularity)

	cmd.MarkFlagsRequiredTogether(flag.Start, flag.End)
	cmd.MarkFlagsMutuallyExclusive(flag.Period, flag.Start)
	cmd.MarkFlagsMutuallyExclusive(flag.Period, flag.End)

	return cmd
}
