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

package update

import (
	"context"
	"errors"
	"fmt"

	"github.com/mongodb/mongodb-atlas-cli/internal/cli"
	"github.com/mongodb/mongodb-atlas-cli/internal/config"
	"github.com/mongodb/mongodb-atlas-cli/internal/flag"
	store "github.com/mongodb/mongodb-atlas-cli/internal/store/atlas"
	"github.com/mongodb/mongodb-atlas-cli/internal/usage"
	"github.com/spf13/afero"
	"github.com/spf13/cobra"
	atlasv2 "go.mongodb.org/atlas-sdk/v20230201004/admin"
)

var (
	active = "ACTIVE"
)

type combinedStore interface {
	store.CompliancePolicy
	store.ProjectLister
}

type UpdateOpts struct {
	cli.GlobalOpts
	cli.WatchOpts

	store combinedStore
	fs    afero.Fs
	path  string
}

var updateTemplate = `Your backup compliance policy has been updated with the following policies:

POLICIES
ID	FREQUENCY INTERVAL	FREQUENCY TYPE	RETENTION
{{- range .ScheduledPolicyItems}}
{{.Id}}	{{if eq .FrequencyType "hourly"}}{{.FrequencyInterval}}{{else}}-{{end}}	{{.FrequencyType}}	{{.RetentionValue}} {{.RetentionUnit}}
{{- end}}
{{if .OnDemandPolicyItem}}{{.OnDemandPolicyItem.Id}}	-	{{.OnDemandPolicyItem.FrequencyType}}	{{.OnDemandPolicyItem.RetentionValue}} {{.OnDemandPolicyItem.RetentionUnit}}{{end}}
`

var errorCode500Template = `received an internal error on the server side, but we would encourage you to double check your inputs.
For this command, invalid inputs are known to cause internal errors in some situations`

func (opts *UpdateOpts) initStore(ctx context.Context) func() error {
	return func() error {
		var err error
		opts.store, err = store.New(store.AuthenticatedPreset(config.Default()), store.WithContext(ctx))
		if err != nil {
			return err
		}
		return nil
	}
}

func (opts *UpdateOpts) setupWatcher() (bool, error) {
	res, err := opts.store.DescribeCompliancePolicy(opts.ConfigProjectID())
	// opts.policy = res
	if err != nil {
		return false, err
	}
	if res.GetState() == "" {
		return false, errors.New("could not access State field")
	}
	return (res.GetState() == active), nil
}

func (opts *UpdateOpts) interactiveRun() error {

	projectID, err := opts.askProjectOptions()
	if err != nil {
		return fmt.Errorf("couldn't get the projectID: %w", err)
	}

	compliancePolicy, err := opts.store.DescribeCompliancePolicy(projectID)
	if err != nil {
		return fmt.Errorf("couldn't fetch the backup compliance policy: %w", err)
	}

	item, err := opts.askPolicyOptions(compliancePolicy)
	if err != nil {
		return fmt.Errorf("couldn't get the policy item: %w", err)
	}

	snapshotInterval, err := opts.askForSnapshotInterval(item)
	if err != nil {
		return fmt.Errorf("couldn't get the snapshot interval: %w", err)
	}
	item.SetFrequencyInterval(snapshotInterval)

	retentionUnit, retentionValue, err := opts.askForRetention(item)
	if err != nil {
		return fmt.Errorf("couldn't get the retention data: %w", err)
	}
	item.SetRetentionValue(retentionValue)
	item.SetRetentionUnit(retentionUnit)

	return opts.update(projectID, compliancePolicy, item)
}

func (opts *UpdateOpts) update(projectID string, compliancePolicy *atlasv2.DataProtectionSettings, item *atlasv2.DiskBackupApiPolicyItem) error {
	err := replaceItem(compliancePolicy, item)
	if err != nil {
		return err
	}

	res, httpResponse, err := opts.store.UpdateCompliancePolicyAndGetResponse(projectID, compliancePolicy)
	if err != nil {
		if httpResponse.StatusCode == 500 {
			return fmt.Errorf("%v: %w", errorCode500Template, err)
		}
		return err
	}
	return opts.Print(res)
}

func replaceItem(compliancePolicy *atlasv2.DataProtectionSettings, item *atlasv2.DiskBackupApiPolicyItem) error {
	items := compliancePolicy.GetScheduledPolicyItems()
	for i, existingItem := range items {
		if existingItem.GetId() == item.GetId() {
			items[i] = *item
			return nil
		}
	}
	onDemandItem := compliancePolicy.GetOnDemandPolicyItem()
	if onDemandItem.GetId() == item.GetId() {
		compliancePolicy.SetOnDemandPolicyItem(*item)
		return nil
	}
	return errors.New("could not replace policy item")
}

func UpdateBuilder() *cobra.Command {
	opts := &UpdateOpts{
		fs: afero.NewOsFs(),
	}
	use := "update"
	cmd := &cobra.Command{
		Use:     use,
		Aliases: cli.GenerateAliases(use),
		Short:   "Update the backup compliance policy for your project with a configuration file.",
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return opts.PreRunE(
				opts.ValidateProjectID,
				opts.initStore(cmd.Context()),
				opts.InitOutput(cmd.OutOrStdout(), updateTemplate),
			)
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			// if err := file.Load(opts.fs, opts.path, opts.policy); err != nil {
			// 	return err
			// }
			if opts.path != "" {
				return nil
			}
			return opts.interactiveRun()
		},
	}

	cmd.Flags().StringVar(&opts.ProjectID, flag.ProjectID, "", usage.ProjectID)
	cmd.Flags().StringVarP(&opts.Output, flag.Output, flag.OutputShort, "", usage.FormatOut)
	_ = cmd.RegisterFlagCompletionFunc(flag.Output, opts.AutoCompleteOutputFlag())
	cmd.Flags().StringVarP(&opts.path, flag.File, flag.FileShort, "", usage.BackupCompliancePolicyFile)
	return cmd
}
