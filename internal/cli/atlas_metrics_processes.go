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

package cli

import (
	"github.com/mongodb/mongocli/internal/config"
	"github.com/mongodb/mongocli/internal/description"
	"github.com/mongodb/mongocli/internal/flags"
	"github.com/mongodb/mongocli/internal/json"
	"github.com/mongodb/mongocli/internal/store"
	"github.com/mongodb/mongocli/internal/usage"
	"github.com/spf13/cobra"
)

type atlasMetricsProcessOpts struct {
	globalOpts
	metricsOpts
	host  string
	port  int
	store store.ProcessMeasurementLister
}

func (opts *atlasMetricsProcessOpts) initStore() error {
	var err error
	opts.store, err = store.New(config.Default())
	return err
}

func (opts *atlasMetricsProcessOpts) Run() error {
	listOpts := opts.newProcessMetricsListOptions()
	result, err := opts.store.ProcessMeasurements(opts.ProjectID(), opts.host, opts.port, listOpts)

	if err != nil {
		return err
	}

	return json.PrettyPrint(result)
}

// mongocli atlas metric(s) process(es) host:port [--granularity granularity] [--period period] [--start start] [--end end] [--type type][--projectId projectId]
func AtlasMetricsProcessBuilder() *cobra.Command {
	opts := &atlasMetricsProcessOpts{}
	cmd := &cobra.Command{
		Use:     "processes [host:port]",
		Short:   description.ProcessMeasurements,
		Aliases: []string{"process"},
		Args:    cobra.ExactArgs(1),
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return opts.PreRunE(opts.initStore)
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			var err error
			opts.host, opts.port, err = getHostnameAndPort(args[0])
			if err != nil {
				return err
			}

			return opts.Run()
		},
	}
	cmd.Flags().IntVar(&opts.pageNum, flags.Page, 0, usage.Page)
	cmd.Flags().IntVar(&opts.itemsPerPage, flags.Limit, 0, usage.Limit)

	cmd.Flags().StringVar(&opts.granularity, flags.Granularity, "", usage.Granularity)
	cmd.Flags().StringVar(&opts.period, flags.Period, "", usage.Period)
	cmd.Flags().StringVar(&opts.start, flags.Start, "", usage.MeasurementStart)
	cmd.Flags().StringVar(&opts.end, flags.End, "", usage.MeasurementEnd)
	cmd.Flags().StringSliceVar(&opts.measurementType, flags.Type, nil, usage.MeasurementType)

	cmd.Flags().StringVar(&opts.projectID, flags.ProjectID, "", usage.ProjectID)

	_ = cmd.MarkFlagRequired(flags.Granularity)

	return cmd
}
