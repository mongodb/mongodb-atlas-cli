// Copyright 2023 MongoDB Inc
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

package scheduled

import (
	"context"
	"fmt"

	"github.com/andreangiolillo/mongocli-test/internal/cli"
	"github.com/andreangiolillo/mongocli-test/internal/config"
	"github.com/andreangiolillo/mongocli-test/internal/flag"
	store "github.com/andreangiolillo/mongocli-test/internal/store/atlas"
	"github.com/andreangiolillo/mongocli-test/internal/usage"
	"github.com/spf13/cobra"
	atlasv2 "go.mongodb.org/atlas-sdk/v20231115002/admin"
)

type UpdateOpts struct {
	cli.GlobalOpts
	cli.WatchOpts
	store             store.CompliancePolicyScheduledPolicyUpdater
	policy            *atlasv2.DataProtectionSettings20231001
	scheduledPolicyID string
	frequencyType     string
	frequencyInterval int
	retentionUnit     string
	retentionValue    int
}

func (opts *UpdateOpts) initStore(ctx context.Context) func() error {
	return func() (err error) {
		opts.store, err = store.New(store.AuthenticatedPreset(config.Default()), store.WithContext(ctx))
		return
	}
}

func (opts *UpdateOpts) watcher() (bool, error) {
	res, err := opts.store.DescribeCompliancePolicy(opts.ConfigProjectID())
	if err != nil {
		return false, err
	}
	opts.policy = res
	return res.GetState() == active, nil
}

func (opts *UpdateOpts) Run() (err error) {
	policyItem := &atlasv2.BackupComplianceScheduledPolicyItem{
		FrequencyType:     opts.frequencyType,
		FrequencyInterval: opts.frequencyInterval,
		RetentionUnit:     opts.retentionUnit,
		RetentionValue:    opts.retentionValue,
		Id:                &opts.scheduledPolicyID,
	}

	if opts.policy, err = opts.store.UpdateScheduledPolicy(opts.ProjectID, policyItem); err != nil {
		return err
	}

	if opts.EnableWatch {
		if errW := opts.Watch(opts.watcher); errW != nil {
			return fmt.Errorf("received an error while watching for completion: %w", errW)
		}
		opts.Template = updateWatchTemplate
	}
	return opts.Print(opts.policy)
}

func UpdateBuilder() *cobra.Command {
	opts := &UpdateOpts{}
	cmd := &cobra.Command{
		Use:   "update",
		Short: "Update a scheduled policy item for the backup compliance policy for your project.",
		Example: `  # Update a backup compliance schedule policy with a weekly frequency, where the snapshot occurs on Monday and has a retention of two months:
  atlas backups compliancepolicy policies scheduled update --scheduledPolicyId 6567f8ad00e6a55f9e448087 --frequencyType weekly --frequencyInterval 1 --retentionUnit months --retentionValue 2`,
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return opts.PreRunE(
				opts.ValidateProjectID,
				opts.initStore(cmd.Context()),
				opts.InitOutput(cmd.OutOrStdout(), updateTemplate),
			)
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return opts.Run()
		},
	}

	cmd.Flags().StringVar(&opts.scheduledPolicyID, flag.ScheduledPolicyID, "", usage.ScheduledPolicyID)
	cmd.Flags().StringVar(&opts.frequencyType, flag.FrequencyType, "", usage.FrequencyType)
	cmd.Flags().IntVar(&opts.frequencyInterval, flag.FrequencyInterval, 0, usage.FrequencyInterval)
	cmd.Flags().StringVar(&opts.retentionUnit, flag.RetentionUnit, "", usage.RetentionUnit)
	cmd.Flags().IntVar(&opts.retentionValue, flag.RetentionValue, 0, usage.RetentionValue)
	_ = cmd.MarkFlagRequired(flag.ScheduledPolicyID)
	_ = cmd.MarkFlagRequired(flag.FrequencyType)
	_ = cmd.MarkFlagRequired(flag.FrequencyInterval)
	_ = cmd.MarkFlagRequired(flag.RetentionUnit)
	_ = cmd.MarkFlagRequired(flag.RetentionValue)

	cmd.Flags().BoolVarP(&opts.EnableWatch, flag.EnableWatch, flag.EnableWatchShort, false, usage.EnableWatchDefault)
	cmd.Flags().StringVar(&opts.ProjectID, flag.ProjectID, "", usage.ProjectID)
	cmd.Flags().StringVarP(&opts.Output, flag.Output, flag.OutputShort, "", usage.FormatOut)
	_ = cmd.RegisterFlagCompletionFunc(flag.Output, opts.AutoCompleteOutputFlag())

	return cmd
}
