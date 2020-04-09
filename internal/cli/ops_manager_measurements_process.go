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
	atlas "github.com/mongodb/go-client-mongodb-atlas/mongodbatlas"
	"github.com/mongodb/mongocli/internal/description"
	"github.com/mongodb/mongocli/internal/flags"
	"github.com/mongodb/mongocli/internal/json"
	"github.com/mongodb/mongocli/internal/store"
	"github.com/mongodb/mongocli/internal/usage"
	"github.com/spf13/cobra"
)

type OpsManagerMeasurementsProcessOpts struct {
	*globalOpts
	pageNum         int
	itemsPerPage    int
	hostID          string
	granularity     string
	period          string
	start           string
	end             string
	measurementType string
	store           store.HostMeasurementLister
}

func (opts *OpsManagerMeasurementsProcessOpts) init() error {
	if opts.ProjectID() == "" {
		return errMissingProjectID
	}

	var err error
	opts.store, err = store.New()
	return err
}

func (opts *OpsManagerMeasurementsProcessOpts) Run() error {
	listOpts := opts.newProcessMeasurementListOptions()
	result, err := opts.store.HostMeasurements(opts.ProjectID(), opts.hostID, listOpts)

	if err != nil {
		return err
	}

	return json.PrettyPrint(result)
}

func (opts *OpsManagerMeasurementsProcessOpts) newProcessMeasurementListOptions() *atlas.ProcessMeasurementListOptions {
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

// mongocli om|cm measurements process(es) hostId [--granularity granularity] [--period period] [--start start] [--end end] [--type type][--projectId projectId]
func OpsManagerMeasurementsProcessBuilder() *cobra.Command {
	opts := &OpsManagerMeasurementsProcessOpts{
		globalOpts: newGlobalOpts(),
	}
	cmd := &cobra.Command{
		Use:     "process",
		Short:   description.ProcessMeasurements,
		Aliases: []string{"processes"},
		Args:    cobra.ExactArgs(1),
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return opts.init()
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			opts.hostID = args[0]
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
