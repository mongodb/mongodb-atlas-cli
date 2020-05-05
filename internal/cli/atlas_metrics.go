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
	"github.com/spf13/cobra"
)

type metricsOpts struct {
	listOpts
	granularity     string
	period          string
	start           string
	end             string
	measurementType []string
}

func (opts *metricsOpts) newProcessMetricsListOptions() *atlas.ProcessMeasurementListOptions {
	o := &atlas.ProcessMeasurementListOptions{
		ListOptions: opts.newListOptions(),
	}
	o.Granularity = opts.granularity
	o.Period = opts.period
	o.Start = opts.start
	o.End = opts.end
	o.M = opts.measurementType

	return o
}

func AtlasMetricsBuilder() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "metrics",
		Aliases: []string{"metric", "measurements", "measurement"},
		Short:   description.Metrics,
	}
	cmd.AddCommand(AtlasMetricsProcessBuilder())
	cmd.AddCommand(AtlasMetricsDisksBuilder())
	cmd.AddCommand(AtlasMetricsDatabasesBuilder())

	return cmd
}
