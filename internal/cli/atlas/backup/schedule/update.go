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
	"context"
	"errors"
	"fmt"

	"github.com/mongodb/mongodb-atlas-cli/internal/cli"
	"github.com/mongodb/mongodb-atlas-cli/internal/config"
	"github.com/mongodb/mongodb-atlas-cli/internal/flag"
	"github.com/mongodb/mongodb-atlas-cli/internal/store"
	"github.com/mongodb/mongodb-atlas-cli/internal/usage"
	"github.com/spf13/cobra"
	atlas "go.mongodb.org/atlas/mongodbatlas"
)

var updateTemplate = "Snapshot backup policy for cluster '{{.ClusterName}}' updated.\n"

type UpdateOpts struct {
	cli.GlobalOpts
	cli.OutputOpts
	clusterName                         string
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
	store                               store.ScheduleUpdater
}

func (opts *UpdateOpts) initStore(ctx context.Context) func() error {
	return func() error {
		var err error
		opts.store, err = store.New(store.AuthenticatedPreset(config.Default()), store.WithContext(ctx))
		return err
	}
}

func (opts *UpdateOpts) Run(cmd *cobra.Command) error {
	backupConfig, err := opts.NewBackupConfig(cmd, opts.clusterName)
	if err != nil {
		return err
	}

	r, err := opts.store.UpdateSchedule(opts.ConfigProjectID(), opts.clusterName, backupConfig)
	if err != nil {
		return err
	}

	return opts.Print(r)
}

func (opts *UpdateOpts) NewBackupConfig(cmd *cobra.Command, clusterName string) (*atlas.CloudProviderSnapshotBackupPolicy, error) {
	out := new(atlas.CloudProviderSnapshotBackupPolicy)

	out.ClusterName = clusterName

	if opts.exportBucketID != "" {
		checkForExport(out)
		out.Export.ExportBucketID = opts.exportBucketID
	}

	if cmd.Flags().Changed(flag.ExportFrequencyType) {
		checkForExport(out)
		out.Export.FrequencyType = opts.exportFrequencyType
	}
	if cmd.Flags().Changed(flag.ReferenceHourOfDay) {
		out.ReferenceHourOfDay = &opts.referenceHourOfDay
	}
	if cmd.Flags().Changed(flag.ReferenceMinuteOfHour) {
		out.ReferenceMinuteOfHour = &opts.referenceMinuteOfHour
	}
	if cmd.Flags().Changed(flag.RestoreWindowDays) {
		out.RestoreWindowDays = &opts.restoreWindowDays
	}

	out.AutoExportEnabled = cli.ReturnValueForSetting(opts.autoExport, opts.noAutoExport)
	out.UpdateSnapshots = cli.ReturnValueForSetting(opts.updateSnapshots, opts.noUpdateSnapshots)
	out.UseOrgAndGroupNamesInExportPrefix = cli.ReturnValueForSetting(opts.useOrgAndGroupNamesInExportPrefix, opts.noUseOrgAndGroupNamesInExportPrefix)

	return out, nil
}

func (opts *UpdateOpts) verifyExportFrequencyType() func() error {
	return func() error {
		if opts.exportFrequencyType != "" {
			if opts.exportFrequencyType != "daily" && opts.exportFrequencyType != "weekly" && opts.exportFrequencyType != "monthly" {
				return errors.New("incorrect value for parameter exportFrequencyType. Value must be daily, weekly, or monthly")
			}
		}
		return nil
	}
}

func (opts *UpdateOpts) verifyReferenceHourOfDay(cmd *cobra.Command) func() error {
	return func() error {
		if cmd.Flags().Changed(flag.ReferenceHourOfDay) {
			if opts.referenceHourOfDay < 0 || opts.referenceHourOfDay > 23 {
				return errors.New("incorrect value for parameter referenceHourOfDay. Value must be an integer between 0 and 23 inclusive")
			}
		}
		return nil
	}
}

func (opts *UpdateOpts) verifyReferenceMinuteOfHour(cmd *cobra.Command) func() error {
	return func() error {
		if cmd.Flags().Changed(flag.ReferenceMinuteOfHour) {
			if opts.referenceMinuteOfHour < 0 || opts.referenceMinuteOfHour > 59 {
				return errors.New("incorrect value for parameter referenceMinuteOfHour. Value must be an integer between 0 and 59 inclusive")
			}
		}
		return nil
	}
}

func (opts *UpdateOpts) verifyRestoreWindowDays(cmd *cobra.Command) func() error {
	return func() error {
		if cmd.Flags().Changed(flag.RestoreWindowDays) {
			if opts.restoreWindowDays <= 0 {
				return errors.New("incorrect value for parameter restoreWindowDays. Value must be a positive, non-zero integer")
			}
		}
		return nil
	}
}

func checkForExport(out *atlas.CloudProviderSnapshotBackupPolicy) {
	if out.Export == nil {
		out.Export = new(atlas.Export)
	}
}

func UpdateBuilder() *cobra.Command {
	opts := &UpdateOpts{}
	cmd := &cobra.Command{
		Use:     "update",
		Aliases: []string{"updates"},
		Short:   "Update a snapshot backup policies for a cluster.",
		Example: fmt.Sprintf(`  The following updates a snapshot backup policies for a cluster Cluster0:
  $ %s backup schedule update --clusterName Cluster0 --updateSnapshots --exportBucketId 62c569f85b7a381c093cc539 --exportFrequencyType monthly`, cli.ExampleAtlasEntryPoint()),
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return opts.PreRunE(
				opts.ValidateProjectID,
				opts.initStore(cmd.Context()),
				opts.InitOutput(cmd.OutOrStdout(), updateTemplate),
				opts.verifyExportFrequencyType(),
				opts.verifyReferenceHourOfDay(cmd),
				opts.verifyReferenceMinuteOfHour(cmd),
				opts.verifyRestoreWindowDays(cmd),
			)
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return opts.Run(cmd)
		},
	}

	cmd.Flags().StringVar(&opts.clusterName, flag.ClusterName, "", usage.ClusterName)
	cmd.Flags().StringVar(&opts.exportBucketID, flag.ExportBucketID, "", usage.BucketID)
	cmd.Flags().StringVar(&opts.exportFrequencyType, flag.ExportFrequencyType, "", usage.ExportFrequencyType)
	cmd.Flags().Int64Var(&opts.referenceHourOfDay, flag.ReferenceHourOfDay, 0, usage.ReferenceHourOfDay)
	cmd.Flags().Int64Var(&opts.referenceMinuteOfHour, flag.ReferenceMinuteOfHour, 0, usage.ReferenceMinuteOfHour)
	cmd.Flags().Int64Var(&opts.restoreWindowDays, flag.RestoreWindowDays, 0, usage.RestoreWindowDays)

	cmd.Flags().BoolVar(&opts.autoExport, flag.AutoExport, false, usage.AutoExport)
	cmd.Flags().BoolVar(&opts.noAutoExport, flag.NoAutoExport, false, usage.NoAutoExport)
	cmd.MarkFlagsMutuallyExclusive(flag.AutoExport, flag.NoAutoExport)
	cmd.MarkFlagsRequiredTogether(flag.AutoExport, flag.ExportBucketID, flag.ExportFrequencyType)

	cmd.Flags().BoolVar(&opts.updateSnapshots, flag.UpdateSnapshots, false, usage.UpdateSnapshots)
	cmd.Flags().BoolVar(&opts.noUpdateSnapshots, flag.NoUpdateSnapshots, false, usage.NoUpdateSnapshots)
	cmd.MarkFlagsMutuallyExclusive(flag.UpdateSnapshots, flag.NoUpdateSnapshots)

	cmd.Flags().BoolVar(&opts.useOrgAndGroupNamesInExportPrefix, flag.UseOrgAndGroupNamesInExportPrefix, false, usage.UseOrgAndGroupNamesInExportPrefix)
	cmd.Flags().BoolVar(&opts.noUseOrgAndGroupNamesInExportPrefix, flag.NoUseOrgAndGroupNamesInExportPrefix, false, usage.NoUseOrgAndGroupNamesInExportPrefix)
	cmd.MarkFlagsMutuallyExclusive(flag.UseOrgAndGroupNamesInExportPrefix, flag.NoUseOrgAndGroupNamesInExportPrefix)

	cmd.Flags().StringVar(&opts.ProjectID, flag.ProjectID, "", usage.ProjectID)
	cmd.Flags().StringVarP(&opts.Output, flag.Output, flag.OutputShort, "", usage.FormatOut)

	_ = cmd.MarkFlagRequired(flag.ClusterName)

	return cmd
}
