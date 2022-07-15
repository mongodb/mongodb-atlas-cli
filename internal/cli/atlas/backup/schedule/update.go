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
	"strconv"
	"strings"

	"github.com/mongodb/mongodb-atlas-cli/internal/cli"
	"github.com/mongodb/mongodb-atlas-cli/internal/config"
	"github.com/mongodb/mongodb-atlas-cli/internal/flag"
	"github.com/mongodb/mongodb-atlas-cli/internal/store"
	"github.com/mongodb/mongodb-atlas-cli/internal/usage"
	"github.com/spf13/cobra"
	atlas "go.mongodb.org/atlas/mongodbatlas"
)

var updateTemplate = "Snapshot backup policy for cluster '{{.ClusterName}}' updated.\n"

const (
	Daily              = "daily"
	Hourly             = "hourly"
	Weekly             = "weekly"
	Monthly            = "monthly"
	BackupPolicyLength = 6
)

type UpdateOpts struct {
	cli.GlobalOpts
	cli.OutputOpts
	clusterName                         string
	exportBucketID                      string
	exportFrequencyType                 string
	backupPolicy                        []string
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
	opts.verifyExportBucketID(out)

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

	out.AutoExportEnabled = returnValueForSetting(opts.autoExport, opts.noAutoExport)
	out.UpdateSnapshots = returnValueForSetting(opts.updateSnapshots, opts.noUpdateSnapshots)
	out.UseOrgAndGroupNamesInExportPrefix = returnValueForSetting(opts.useOrgAndGroupNamesInExportPrefix, opts.noUseOrgAndGroupNamesInExportPrefix)

	var policies []atlas.Policy

	for _, backupPolicy := range opts.backupPolicy {
		policyItems := strings.Split(backupPolicy, ",")
		frequencyInterval, err := strconv.Atoi(policyItems[3])
		if err != nil {
			return nil, err
		}
		retentionValue, err := strconv.Atoi(policyItems[5])
		if err != nil {
			return nil, err
		}

		policyIndex := checkIfPolicyExists(policyItems[1], policies)

		if policyIndex != -1 {
			policyItem := atlas.PolicyItem{
				ID:                policyItems[1],
				FrequencyType:     policyItems[2],
				FrequencyInterval: frequencyInterval,
				RetentionUnit:     policyItems[4],
				RetentionValue:    retentionValue,
			}
			policies[policyIndex].PolicyItems = append(policies[policyIndex].PolicyItems, policyItem)
		} else {
			policy := atlas.Policy{
				ID: policyItems[0],
				PolicyItems: []atlas.PolicyItem{
					{
						ID:                policyItems[1],
						FrequencyType:     policyItems[2],
						FrequencyInterval: frequencyInterval,
						RetentionUnit:     policyItems[4],
						RetentionValue:    retentionValue,
					},
				},
			}
			policies = append(policies, policy)
		}
	}

	out.Policies = policies

	return out, nil
}

func checkIfPolicyExists(currentPolicyID string, policies []atlas.Policy) int {
	for index, policy := range policies {
		if policy.ID == currentPolicyID {
			return index
		}
	}

	return -1
}

func (opts *UpdateOpts) verifyExportBucketID(out *atlas.CloudProviderSnapshotBackupPolicy) {
	if opts.exportBucketID != "" {
		checkForExport(out)
		out.Export.ExportBucketID = opts.exportBucketID
	}
}

func (opts *UpdateOpts) verifyExportFrequencyType() error {
	if opts.exportFrequencyType != "" {
		if opts.exportFrequencyType != Daily && opts.exportFrequencyType != Weekly && opts.exportFrequencyType != Monthly {
			return errors.New("incorrect value for parameter exportFrequencyType. Value must be daily, weekly, or monthly")
		}
	}
	return nil
}

func (opts *UpdateOpts) verifyReferenceHourOfDay(cmd *cobra.Command) error {
	if cmd.Flags().Changed(flag.ReferenceHourOfDay) {
		if opts.referenceHourOfDay < 0 || opts.referenceHourOfDay > 23 {
			return errors.New("incorrect value for parameter referenceHourOfDay. Value must be an integer between 0 and 23 inclusive")
		}
	}
	return nil
}

func (opts *UpdateOpts) verifyReferenceMinuteOfHour(cmd *cobra.Command) error {
	if cmd.Flags().Changed(flag.ReferenceMinuteOfHour) {
		if opts.referenceMinuteOfHour < 0 || opts.referenceMinuteOfHour > 59 {
			return errors.New("incorrect value for parameter referenceMinuteOfHour. Value must be an integer between 0 and 59 inclusive")
		}
	}
	return nil
}

func (opts *UpdateOpts) verifyRestoreWindowDays(cmd *cobra.Command) error {
	if cmd.Flags().Changed(flag.RestoreWindowDays) {
		if opts.restoreWindowDays <= 0 {
			return errors.New("incorrect value for parameter restoreWindowDays. Value must be a positive, non-zero integer")
		}
	}
	return nil
}

func (opts *UpdateOpts) validateBackupPolicy(cmd *cobra.Command) error {
	if cmd.Flags().Changed(flag.Policy) {
		for _, policy := range opts.backupPolicy {
			policyItems := strings.Split(policy, ",")
			fmt.Println(policyItems)
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
			err = validateFrequencyIntervalNumber(policyItems[2], policyItems[3])
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

func validatePolicyLength(policyItems []string) error {
	if len(policyItems) != BackupPolicyLength {
		return errors.New("error when parsing policy. You must specify it in a format: '--policy policyID,policyItemID,frequencyType,frequencyIntervalNumber,retentionUnit,retentionValue'")
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
	if frequencyType != Hourly && frequencyType != Daily && frequencyType != Weekly && frequencyType != Monthly {
		return errors.New("frequencyType was provided in an incorrect format. It must be equal to 'hourly', 'daily', 'weekly' or 'monthly'")
	}
	return nil
}

func validateFrequencyIntervalNumber(frequencyType, frequencyIntervalNumber string) error {
	intervalNumber, err := strconv.Atoi(frequencyIntervalNumber)
	if err != nil {
		return errors.New("frequencyIntervalNumber was provided in an incorrect format. It must be an integer")
	}

	hourlyAllowedValues := []int{1, 2, 4, 6, 8, 12}
	dailyAllowedValues := []int{1}
	weeklyAllowedValues := []int{1, 7}
	monthlyAllowedValues := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 40}

	switch frequencyType {
	case Hourly:
		if !intInSlice(intervalNumber, hourlyAllowedValues) {
			return errors.New("frequencyIntervalNumber was provided in an incorrect format for 'hourly' frequencyType")
		}
	case Daily:
		if !intInSlice(intervalNumber, dailyAllowedValues) {
			return errors.New("frequencyIntervalNumber was provided in an incorrect format for 'daily' frequencyType")
		}
	case Weekly:
		if !intInSlice(intervalNumber, weeklyAllowedValues) {
			return errors.New("frequencyIntervalNumber was provided in an incorrect format for 'weekly' frequencyType")
		}
	case Monthly:
		if !intInSlice(intervalNumber, monthlyAllowedValues) {
			return errors.New("frequencyIntervalNumber was provided in an incorrect format for 'monthly' frequencyType")
		}
	}

	return nil
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

func intInSlice(a int, list []int) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

func checkForExport(out *atlas.CloudProviderSnapshotBackupPolicy) {
	if out.Export == nil {
		out.Export = new(atlas.Export)
	}
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

func UpdateBuilder() *cobra.Command {
	opts := &UpdateOpts{}
	cmd := &cobra.Command{
		Use:     "update",
		Aliases: []string{"updates"},
		Short:   "Update a snapshot backup policies for a cluster.",
		Example: fmt.Sprintf(`  The following updates a snapshot backup policies for a cluster Cluster0:
  $ %s backup schedule update --clusterName Cluster0 --updateSnapshots --exportBucketId 62c569f85b7a381c093cc539 --exportFrequencyType monthly`, cli.ExampleAtlasEntryPoint()),
		PreRunE: func(cmd *cobra.Command, args []string) error {
			preRun := opts.PreRunE(
				opts.ValidateProjectID,
				opts.initStore(cmd.Context()),
				opts.InitOutput(cmd.OutOrStdout(), updateTemplate),
			)
			err := opts.verifyExportFrequencyType()
			if err != nil {
				return err
			}
			err = opts.verifyReferenceHourOfDay(cmd)
			if err != nil {
				return err
			}
			err = opts.verifyReferenceMinuteOfHour(cmd)
			if err != nil {
				return err
			}
			err = opts.verifyRestoreWindowDays(cmd)
			if err != nil {
				return err
			}
			err = opts.validateBackupPolicy(cmd)
			if err != nil {
				return err
			}
			return preRun
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

	cmd.Flags().StringArrayVar(&opts.backupPolicy, flag.BackupPolicy, nil, usage.BackupPolicy)

	cmd.Flags().StringVar(&opts.ProjectID, flag.ProjectID, "", usage.ProjectID)
	cmd.Flags().StringVarP(&opts.Output, flag.Output, flag.OutputShort, "", usage.FormatOut)

	_ = cmd.MarkFlagRequired(flag.ClusterName)

	return cmd
}
