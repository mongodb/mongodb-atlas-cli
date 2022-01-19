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

	atlas "go.mongodb.org/atlas/mongodbatlas"
)

type MetricsOpts struct {
	ListOpts
	Granularity     string
	Period          string
	Start           string
	End             string
	MeasurementType []string
}

func (opts *MetricsOpts) NewProcessMetricsListOptions() *atlas.ProcessMeasurementListOptions {
	o := &atlas.ProcessMeasurementListOptions{
		ListOptions: opts.NewListOptions(),
	}
	o.Granularity = opts.Granularity
	o.Period = opts.Period
	o.Start = opts.Start
	o.End = opts.End
	o.M = opts.MeasurementType

	return o
}

// ValidatePeriodStartEnd validates period, start and end flags.
func (opts *MetricsOpts) ValidatePeriodStartEnd() error {
	if opts.Period == "" && opts.Start == "" && opts.End == "" {
		return fmt.Errorf("either the --period flag or the --start and --end flags are required")
	}
	if opts.Period != "" && (opts.Start != "" || opts.End != "") {
		return fmt.Errorf("the --period flag is mutually exclusive to the --start and --end flags")
	}
	if (opts.Start != "" && opts.End == "") || (opts.Start == "" && opts.End != "") {
		return fmt.Errorf("the --start and --end flags need to be used together")
	}
	return nil
}
