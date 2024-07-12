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
	"regexp"
	"slices"
	"strconv"
	"strings"

	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/config"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/file"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/flag"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/store"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/usage"
	"github.com/spf13/afero"
	"github.com/spf13/cobra"
	atlasv2 "go.mongodb.org/atlas-sdk/v20240530002/admin"
)

var updateTemplate = "Snapshot backup policy for cluster '{{.ClusterName}}' updated.\n"
var ErrFrequencyIntervalNumber = errors.New("invalid frequencyIntervalNumber")

const (
	daily              = "daily"
	hourly             = "hourly"
	weekly             = "weekly"
	monthly            = "monthly"
	yearly             = "yearly"
	ondemand           = "ondemand"
	backupPolicyLength = 6
)

type UpdateOpts struct {
	cli.GlobalOpts
	cli.OutputOpts
	clusterName                         string
	exportBucketID                      string
	exportFrequencyType                 string
	backupPolicy                        []string
	referenceHourOfDay                  int
	referenceMinuteOfHour               int
	restoreWindowDays                   int
	autoExport                          bool
	noAutoExport                        bool
	updateSnapshots                     bool
	noUpdateSnapshots                   bool
	useOrgAndGroupNamesInExportPrefix   bool
	noUseOrgAndGroupNamesInExportPrefix bool
	store                               store.ScheduleDescriberUpdater
	filename                            string
	fs                                  afero.Fs
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

func (opts *UpdateOpts) NewBackupConfig(cmd *cobra.Command, clusterName string) (*atlasv2.DiskBackupSnapshotSchedule, error) {
	out := new(atlasv2.DiskBackupSnapshotSchedule)

	if opts.filename != "" {
		if err := file.Load(opts.fs, opts.filename, out); err != nil {
			return nil, err
		}
		out.ClusterName = &clusterName
		return out, nil
	}

	out.ClusterName = &clusterName

	if opts.exportBucketID != "" {
		checkForExport(out)
		out.Export.ExportBucketId = &opts.exportBucketID
	}

	if cmd.Flags().Changed(flag.ExportFrequencyType) {
		checkForExport(out)
		out.Export.FrequencyType = &opts.exportFrequencyType
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

	if cmd.Flags().Changed(flag.BackupPolicy) {
		currentSchedule, err := opts.store.DescribeSchedule(opts.ConfigProjectID(), opts.clusterName)
		if err != nil {
			return nil, err
		}
		policies := currentSchedule.GetPolicies()

		for _, backupPolicy := range opts.backupPolicy {
			policyItems := strings.Split(backupPolicy, ",")
			frequencyInterval, err := strconv.ParseInt(policyItems[3], 10, 32)
			if err != nil {
				return nil, err
			}
			retentionValue, err := strconv.ParseInt(policyItems[5], 10, 32)
			if err != nil {
				return nil, err
			}

			policyIndex := findPolicyIndex(policyItems[0], policies)

			if policyIndex == -1 {
				return nil, errors.New("incorrect value for parameter policyID. Policy with such ID does not exist")
			}

			policyItem := atlasv2.DiskBackupApiPolicyItem{
				Id:                &policyItems[1],
				FrequencyType:     policyItems[2],
				FrequencyInterval: int(frequencyInterval),
				RetentionUnit:     policyItems[4],
				RetentionValue:    int(retentionValue),
			}
			policyItemIndex := findPolicyItemsIndex(policyItems[1], policies[policyIndex].GetPolicyItems())
			if policyItemIndex == -1 {
				return nil, errors.New("incorrect value for parameter policyItemID. Policy item with such ID does not exist")
			}

			policies[policyIndex].GetPolicyItems()[policyItemIndex] = policyItem
		}

		out.Policies = &policies
	}

	return out, nil
}

func findPolicyIndex(policyID string, policies []atlasv2.AdvancedDiskBackupSnapshotSchedulePolicy) int {
	for index, policy := range policies {
		if policy.GetId() == policyID {
			return index
		}
	}

	return -1
}

func findPolicyItemsIndex(policyItemID string, policyItems []atlasv2.DiskBackupApiPolicyItem) int {
	for index, policyItem := range policyItems {
		if policyItemID == policyItem.GetId() {
			return index
		}
	}

	return -1
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

func (opts *UpdateOpts) validateBackupPolicy(cmd *cobra.Command) func() error {
	return func() error {
		if cmd.Flags().Changed(flag.Policy) {
			for _, policy := range opts.backupPolicy {
				policyItems := strings.Split(policy, ",")
				err := validatePolicyLength(policyItems)
				if err != nil {
					return err
				}
				err = validateID(policyItems[0], policyItems[1])
				if err != nil {
					return err
				}
				err = validateFrequencyType(policyItems[2])
				if err != nil {
					return err
				}
				err = validateFrequencyIntervalNumber(policyItems[3])
				if err != nil {
					return err
				}
				err = validateRetentionType(policyItems[4])
				if err != nil {
					return err
				}
				err = validateRetentionValue(policyItems[5])
				if err != nil {
					return err
				}
			}
		}
		return nil
	}
}

func validatePolicyLength(policyItems []string) error {
	if len(policyItems) != backupPolicyLength {
		return errors.New("error when parsing policy. You must specify it in the following format: '--policy policyID,policyItemID,frequencyType,frequencyIntervalNumber,retentionUnit,retentionValue'")
	}
	return nil
}

func validateID(policyID, policyItemID string) error {
	r := regexp.MustCompile(`^[A-Za-z\d]{24}$`) // Regular expression that matches 24-characters alphanumeric string
	if !r.MatchString(policyID) {
		return errors.New("policyID was provided in an incorrect format. It must be a 24-characters alphanumeric string")
	}
	if !r.MatchString(policyItemID) {
		return errors.New("policyItemID was provided in an incorrect format. It must be a 24-characters alphanumeric string")
	}
	return nil
}

func validateFrequencyType(frequencyType string) error {
	allowedValues := []string{hourly, daily, weekly, monthly, yearly, ondemand}
	if !slices.Contains(allowedValues, frequencyType) {
		return errors.New("frequencyType was provided in an incorrect format. It must be equal to 'hourly', 'daily', 'weekly', 'monthly', 'yearly' or 'ondemand'")
	}
	return nil
}

func validateFrequencyIntervalNumber(frequencyIntervalNumber string) error {
	intervalNumber, err := strconv.Atoi(frequencyIntervalNumber)
	if err != nil {
		return errors.New("frequencyIntervalNumber was provided in an incorrect format. It must be an integer")
	}

	if (intervalNumber >= 1 && intervalNumber <= 28) || intervalNumber == 40 {
		return nil
	}

	return fmt.Errorf("%w: %d", ErrFrequencyIntervalNumber, intervalNumber)
}

func validateRetentionType(retentionType string) error {
	if retentionType != "days" && retentionType != "weeks" && retentionType != "months" {
		return errors.New("retentionType was provided in an incorrect format. It must be equal to 'days', 'weeks' or 'months'")
	}
	return nil
}

func validateRetentionValue(retentionValue string) error {
	_, err := strconv.Atoi(retentionValue)
	if err != nil {
		return errors.New("retentionValue was provided in an incorrect format. It must be an integer")
	}
	return nil
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

func checkForExport(out *atlasv2.DiskBackupSnapshotSchedule) {
	if out.Export == nil {
		out.Export = new(atlasv2.AutoExportPolicy)
	}
}

// atlas backup(s) schedule update --clusterName clusterName [--projectId projectId] [--exportBucketId exportBucketID].
func UpdateBuilder() *cobra.Command {
	opts := &UpdateOpts{
		fs: afero.NewOsFs(),
	}
	cmd := &cobra.Command{
		Use:     "update",
		Aliases: []string{"updates"},
		Short:   "Modify the backup schedule for the specified cluster for your project.",
		Long: `The backup schedule defines when MongoDB takes scheduled snapshots and how long it stores those snapshots.

` + fmt.Sprintf(usage.RequiredRole, "Project Owner"),
		Annotations: map[string]string{
			"output": updateTemplate,
		},
		Example: `  # Update a snapshot backup policy for a cluster named Cluster0 to back up snapshots every 6 hours and, retain for 7 days, and update retention of previously-taken snapshots:
  atlas backup schedule update --clusterName Cluster0 --updateSnapshots --policy 62da8faac84a2721e448d767,62da8faac84a2721e448d768,hourly,6,days,7
  
  # Update a snapshot backup policy for a cluster named Cluster0 to export snapshots monthly to an S3 bucket:
  atlas backup schedule update --clusterName Cluster0 --exportBucketId 62c569f85b7a381c093cc539 --exportFrequencyType monthly`,
		PreRunE: func(cmd *cobra.Command, _ []string) error {
			return opts.PreRunE(
				opts.ValidateProjectID,
				opts.initStore(cmd.Context()),
				opts.InitOutput(cmd.OutOrStdout(), updateTemplate),
				opts.verifyExportFrequencyType(),
				opts.verifyReferenceHourOfDay(cmd),
				opts.verifyReferenceMinuteOfHour(cmd),
				opts.verifyRestoreWindowDays(cmd),
				opts.validateBackupPolicy(cmd),
			)
		},
		RunE: func(cmd *cobra.Command, _ []string) error {
			return opts.Run(cmd)
		},
	}

	cmd.Flags().StringVar(&opts.clusterName, flag.ClusterName, "", usage.ClusterName)
	cmd.Flags().StringVar(&opts.exportBucketID, flag.ExportBucketID, "", usage.BucketID)
	cmd.Flags().StringVar(&opts.exportFrequencyType, flag.ExportFrequencyType, "", usage.ExportFrequencyType)
	cmd.Flags().IntVar(&opts.referenceHourOfDay, flag.ReferenceHourOfDay, 0, usage.ReferenceHourOfDay)
	cmd.Flags().IntVar(&opts.referenceMinuteOfHour, flag.ReferenceMinuteOfHour, 0, usage.ReferenceMinuteOfHour)
	cmd.Flags().IntVar(&opts.restoreWindowDays, flag.RestoreWindowDays, 0, usage.RestoreWindowDays)

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

	cmd.Flags().StringArrayVar(&opts.backupPolicy, flag.BackupPolicy, nil, usage.BackupPolicy)

	cmd.Flags().StringVar(&opts.ProjectID, flag.ProjectID, "", usage.ProjectID)
	cmd.Flags().StringVarP(&opts.Output, flag.Output, flag.OutputShort, "", usage.FormatOut)
	_ = cmd.RegisterFlagCompletionFunc(flag.Output, opts.AutoCompleteOutputFlag())

	cmd.Flags().StringVarP(&opts.filename, flag.File, flag.FileShort, "", usage.BackupFilename)

	_ = cmd.MarkFlagRequired(flag.ClusterName)

	_ = cmd.MarkFlagFilename(flag.File)

	cmd.MarkFlagsMutuallyExclusive(flag.File, flag.ExportBucketID)
	cmd.MarkFlagsMutuallyExclusive(flag.File, flag.ExportFrequencyType)
	cmd.MarkFlagsMutuallyExclusive(flag.File, flag.ReferenceHourOfDay)
	cmd.MarkFlagsMutuallyExclusive(flag.File, flag.ReferenceMinuteOfHour)
	cmd.MarkFlagsMutuallyExclusive(flag.File, flag.RestoreWindowDays)
	cmd.MarkFlagsMutuallyExclusive(flag.File, flag.AutoExport)
	cmd.MarkFlagsMutuallyExclusive(flag.File, flag.NoAutoExport)
	cmd.MarkFlagsMutuallyExclusive(flag.File, flag.UpdateSnapshots)
	cmd.MarkFlagsMutuallyExclusive(flag.File, flag.NoUpdateSnapshots)
	cmd.MarkFlagsMutuallyExclusive(flag.File, flag.UseOrgAndGroupNamesInExportPrefix)
	cmd.MarkFlagsMutuallyExclusive(flag.File, flag.NoUseOrgAndGroupNamesInExportPrefix)
	cmd.MarkFlagsMutuallyExclusive(flag.File, flag.BackupPolicy)

	return cmd
}
