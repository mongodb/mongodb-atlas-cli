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

package ondemand

import (
	"context"
	"fmt"

	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/config"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/flag"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/store"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/usage"
	"github.com/spf13/cobra"
	atlasv2 "go.mongodb.org/atlas-sdk/v20231115009/admin"
)

type CreateOpts struct {
	cli.GlobalOpts
	cli.WatchOpts
	store          store.CompliancePolicyOnDemandPolicyCreator
	policy         *atlasv2.DataProtectionSettings20231001
	retentionUnit  string
	retentionValue int
}

const (
	active                = "ACTIVE"
	onDemandFrequencyType = "ondemand"
)

const updateTemplate = `Your backup compliance policy is being updated
`
const updateWatchTemplate = `Your backup compliance policy has been updated
`

func (opts *CreateOpts) initStore(ctx context.Context) func() error {
	return func() (err error) {
		opts.store, err = store.New(store.AuthenticatedPreset(config.Default()), store.WithContext(ctx))
		return
	}
}

func (opts *CreateOpts) watcher() (any, bool, error) {
	res, err := opts.store.DescribeCompliancePolicy(opts.ConfigProjectID())
	if err != nil {
		return nil, false, err
	}
	opts.policy = res
	return nil, res.GetState() == active, nil
}

func (opts *CreateOpts) Run() (err error) {
	policyItem := &atlasv2.BackupComplianceOnDemandPolicyItem{
		FrequencyType:  onDemandFrequencyType,
		RetentionUnit:  opts.retentionUnit,
		RetentionValue: opts.retentionValue,
	}

	if opts.policy, err = opts.store.CreateOnDemandPolicy(opts.ProjectID, policyItem); err != nil {
		return err
	}

	if opts.EnableWatch {
		if _, errW := opts.Watch(opts.watcher); errW != nil {
			return fmt.Errorf("received an error while watching for completion: %w", errW)
		}
		opts.Template = updateWatchTemplate
	}
	return opts.Print(opts.policy)
}

func CreateBuilder() *cobra.Command {
	opts := &CreateOpts{}
	cmd := &cobra.Command{
		Use:   "create",
		Short: "Create the on-demand policy item of the backup compliance policy for your project.",
		Example: `  # Create a backup compliance on-demand policy with a retention of two weeks:
  atlas backups compliancepolicy policies ondemand create --retentionUnit weeks --retentionValue 2`,
		PreRunE: func(cmd *cobra.Command, _ []string) error {
			return opts.PreRunE(
				opts.ValidateProjectID,
				opts.initStore(cmd.Context()),
				opts.InitOutput(cmd.OutOrStdout(), updateTemplate),
			)
		},
		RunE: func(_ *cobra.Command, _ []string) error {
			return opts.Run()
		},
	}

	cmd.Flags().StringVar(&opts.retentionUnit, flag.RetentionUnit, "", usage.RetentionUnit)
	cmd.Flags().IntVar(&opts.retentionValue, flag.RetentionValue, 0, usage.RetentionValue)
	_ = cmd.MarkFlagRequired(flag.RetentionUnit)
	_ = cmd.MarkFlagRequired(flag.RetentionValue)

	cmd.Flags().BoolVarP(&opts.EnableWatch, flag.EnableWatch, flag.EnableWatchShort, false, usage.EnableWatchDefault)
	cmd.Flags().StringVar(&opts.ProjectID, flag.ProjectID, "", usage.ProjectID)
	cmd.Flags().StringVarP(&opts.Output, flag.Output, flag.OutputShort, "", usage.FormatOut)
	_ = cmd.RegisterFlagCompletionFunc(flag.Output, opts.AutoCompleteOutputFlag())

	return cmd
}
