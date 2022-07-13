// Copyright 2022 MongoDB Inc
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

package schedule

import (
	"errors"
	"math"

	"github.com/mongodb/mongodb-atlas-cli/internal/cli"
	"github.com/spf13/cobra"
	atlas "go.mongodb.org/atlas/mongodbatlas"
)

type ExportFrequency string

type ConfigOpts struct {
	exportBucketID                      string
	exportFrequencyType                 string
	referenceHourOfDay                  int64
	referenceMinuteOfHour               int64
	restoreWindowDays                   int64
	autoExport                          bool
	noAutoExport                        bool
	updateSnapshots                     bool
	noUpdateSnapshots                   bool
	useOrgAndGroupNamesInExportPrefix   bool
	noUseOrgAndGroupNamesInExportPrefix bool
}

// TODO: Implement adding the missing fields of the CloudProviderSnapshotBackupPolicy struct.
func (opts *UpdateOpts) NewBackupPolicy(clusterName string) (*atlas.CloudProviderSnapshotBackupPolicy, error) {
	out := new(atlas.CloudProviderSnapshotBackupPolicy)
	out.Export = new(atlas.Export)

	out.ClusterName = clusterName

	if opts.configOpts.exportBucketID != "" {
		out.Export.ExportBucketID = opts.configOpts.exportBucketID
	}

	if opts.configOpts.exportFrequencyType != "" {
		if opts.configOpts.exportFrequencyType != "daily" && opts.configOpts.exportFrequencyType != "weekly" && opts.configOpts.exportFrequencyType != "monthly" {
			return nil, errors.New("incorrect value for parameter exportFrequencyType. Value must be daily, weekly, or monthly")
		}
		out.Export.FrequencyType = opts.configOpts.exportFrequencyType
	}

	if opts.configOpts.referenceHourOfDay != math.MinInt64 {
		if opts.configOpts.referenceHourOfDay < 0 || opts.configOpts.referenceHourOfDay > 23 {
			return nil, errors.New("incorrect value for parameter referenceHourOfDay. Value must be an integer between 0 and 23 inclusive")
		}
		out.ReferenceHourOfDay = &opts.configOpts.referenceHourOfDay
	}

	if opts.configOpts.referenceMinuteOfHour != math.MinInt64 {
		if opts.configOpts.referenceMinuteOfHour < 0 || opts.configOpts.referenceMinuteOfHour > 59 {
			return nil, errors.New("incorrect value for parameter referenceMinuteOfHour. Value must be an integer between 0 and 59 inclusive")
		}
		out.ReferenceMinuteOfHour = &opts.configOpts.referenceMinuteOfHour
	}

	if opts.configOpts.restoreWindowDays != math.MinInt64 {
		if opts.configOpts.restoreWindowDays <= 0 {
			return nil, errors.New("incorrect value for parameter restoreWindowDays. Value must be a positive, non-zero integer")
		}
		out.RestoreWindowDays = &opts.configOpts.restoreWindowDays
	}

	out.AutoExportEnabled = returnValueForSetting(opts.configOpts.autoExport, opts.configOpts.noAutoExport)
	out.UpdateSnapshots = returnValueForSetting(opts.configOpts.updateSnapshots, opts.configOpts.noUpdateSnapshots)
	out.UseOrgAndGroupNamesInExportPrefix = returnValueForSetting(opts.configOpts.useOrgAndGroupNamesInExportPrefix, opts.configOpts.noUseOrgAndGroupNamesInExportPrefix)

	return out, nil
}

func returnValueForSetting(enableFlag, disableFlag bool) *bool {
	var valueToSet bool
	if enableFlag && disableFlag {
		return nil
	}
	if enableFlag {
		valueToSet = true
		return &valueToSet
	}
	if disableFlag {
		valueToSet = false
		return &valueToSet
	}
	return nil
}

func Builder() *cobra.Command {
	const use = "schedule"
	cmd := &cobra.Command{
		Use:     use,
		Short:   "Return a cloud backup schedule for the cluster you specify.",
		Aliases: cli.GenerateAliases(use),
	}

	cmd.AddCommand(
		DescribeBuilder(),
		UpdateBuilder(),
	)

	return cmd
}
