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
	"fmt"
	"strconv"
	"strings"

	atlas "github.com/mongodb/go-client-mongodb-atlas/mongodbatlas"
	"github.com/mongodb/mongocli/internal/description"
	"github.com/mongodb/mongocli/internal/flags"
	"github.com/mongodb/mongocli/internal/json"
	"github.com/mongodb/mongocli/internal/store"
	"github.com/mongodb/mongocli/internal/usage"
	"github.com/spf13/cobra"
)

type atlasMeasurementsProcessOpts struct {
	*globalOpts
	pageNum         int
	itemsPerPage    int
	host            string
	port            int
	granularity     string
	period          string
	start           string
	end             string
	measurementType string
	store           store.ProcessMeasurementLister
}

func (opts *atlasMeasurementsProcessOpts) init() error {
	if opts.ProjectID() == "" {
		return errMissingProjectID
	}

	var err error
	opts.store, err = store.New()
	return err
}

func (opts *atlasMeasurementsProcessOpts) Run() error {
	listOpts := opts.newProcessMeasurementListOptions()
	result, err := opts.store.ListProcessMeasurements(opts.ProjectID(), opts.host, opts.port, listOpts)

	if err != nil {
		return err
	}

	return json.PrettyPrint(result)
}

func (opts *atlasMeasurementsProcessOpts) getHostNameAndPort(hostInfo string) (string, int, error) {
	host := strings.SplitN(hostInfo, ":", -1)
	if len(host) != 2 {
		return "", 0, fmt.Errorf("expected hostname:port, got %s", host)
	}

	port, err := strconv.Atoi(host[1])
	if err != nil {
		return "", 0, err
	}

	return host[0], port, nil
}

func (opts *atlasMeasurementsProcessOpts) newProcessMeasurementListOptions() *atlas.ProcessMeasurementListOptions {
	return &atlas.ProcessMeasurementListOptions{
		ListOptions: &atlas.ListOptions{
			PageNum:      opts.pageNum,
			ItemsPerPage: opts.itemsPerPage,
		},
		Granularity: opts.granularity,
		Period:      opts.period,
		Start:       opts.start,
		End:         opts.end,
		M:           opts.measurementType,
	}
}

// mongocli atlas measurements process(es) host:port [--granularity granularity] [--period period] [--start start] [--end end] [--type type][--projectId projectId]
func AtlasMeasurementsProcessBuilder() *cobra.Command {
	opts := &atlasMeasurementsProcessOpts{
		globalOpts: newGlobalOpts(),
	}
	cmd := &cobra.Command{
		Use:   "process",
		Short: description.ProcessMeasurements,
		Args:  cobra.ExactArgs(1),
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return opts.init()
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			var err error
			opts.host, opts.port, err = opts.getHostNameAndPort(args[0])
			if err != nil {
				return err
			}

			return opts.Run()
		},
	}

	cmd.Flags().StringVar(&opts.granularity, flags.Granularity, "", usage.Granularity)
	cmd.Flags().StringVar(&opts.period, flags.Period, "", usage.Period)
	cmd.Flags().StringVar(&opts.start, flags.Start, "", usage.Start)
	cmd.Flags().StringVar(&opts.end, flags.End, "", usage.End)
	cmd.Flags().StringVar(&opts.measurementType, flags.MeasurementType, "", usage.MeasurementType)

	cmd.Flags().StringVar(&opts.projectID, flags.ProjectID, "", usage.ProjectID)

	return cmd
}
