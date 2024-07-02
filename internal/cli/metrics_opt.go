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

	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/convert"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/flag"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/pointer"
	atlasv2 "go.mongodb.org/atlas-sdk/v20240530002/admin"
)

type MetricsOpts struct {
	ListOpts
	Granularity     string
	Period          string
	Start           string
	End             string
	MeasurementType []string
}

func (opts *MetricsOpts) NewProcessMeasurementsAPIParams(groupID string, processID string) *atlasv2.GetHostMeasurementsApiParams {
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

func (opts *MetricsOpts) NewDiskMeasurementsAPIParams(groupID string, processID string, partitionName string) *atlasv2.GetDiskMeasurementsApiParams {
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

func (opts *MetricsOpts) NewDatabaseMeasurementsAPIParams(groupID string, processID string, dbName string) *atlasv2.GetDatabaseMeasurementsApiParams {
	p := &atlasv2.GetDatabaseMeasurementsApiParams{
		GroupId:      groupID,
		ProcessId:    processID,
		DatabaseName: dbName,
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

// ValidatePeriodStartEnd validates period, start and end flags.
func (opts *MetricsOpts) ValidatePeriodStartEnd() error {
	if opts.Period == "" && opts.Start == "" && opts.End == "" {
		return fmt.Errorf("either the --%s flag or the --%s and --%s flags are required", flag.Period, flag.Start, flag.End)
	}
	return nil
}

// GetHostnameAndPort return the hostname and the port starting from the string hostname:port.
func GetHostnameAndPort(hostInfo string) (hostname string, port int, err error) {
	const hostnameParts = 2
	host := strings.Split(hostInfo, ":")
	if len(host) != hostnameParts {
		return "", 0, fmt.Errorf("expected hostname:port, got %s", host)
	}

	port, err = strconv.Atoi(host[1])
	if err != nil {
		return "", 0, fmt.Errorf("invalid port number, got %s", host[1])
	}

	return host[0], port, nil
}
