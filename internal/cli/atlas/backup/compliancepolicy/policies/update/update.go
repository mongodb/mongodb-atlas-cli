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
	"fmt"

	"github.com/mongodb/mongodb-atlas-cli/internal/cli"
	"github.com/mongodb/mongodb-atlas-cli/internal/cli/atlas/backup/compliancepolicy/watcher"
	"github.com/mongodb/mongodb-atlas-cli/internal/config"
	"github.com/mongodb/mongodb-atlas-cli/internal/file"
	"github.com/mongodb/mongodb-atlas-cli/internal/flag"
	store "github.com/mongodb/mongodb-atlas-cli/internal/store/atlas"
	"github.com/mongodb/mongodb-atlas-cli/internal/usage"
	"github.com/spf13/afero"
	"github.com/spf13/cobra"
	atlasv2 "go.mongodb.org/atlas-sdk/v20230201004/admin"
)

type updateStore interface {
	store.CompliancePolicyPolicyItemUpdater
	store.ProjectLister
	store.CompliancePolicyDescriber
}

type UpdateOpts struct {
	cli.GlobalOpts
	cli.WatchOpts
	projectID string
	store     updateStore
	fs        afero.Fs
	path      string
}

var updateTemplate = `Your backup compliance policy is being updated with the following policies:
ID	FREQUENCY INTERVAL	FREQUENCY TYPE	RETENTION
{{- range .ScheduledPolicyItems}}
{{.Id}}	{{if eq .FrequencyType "hourly"}}{{.FrequencyInterval}}{{else}}-{{end}}	{{.FrequencyType}}	{{.RetentionValue}} {{.RetentionUnit}}
{{- end}}
{{if .OnDemandPolicyItem}}{{.OnDemandPolicyItem.Id}}	-	{{.OnDemandPolicyItem.FrequencyType}}	{{.OnDemandPolicyItem.RetentionValue}} {{.OnDemandPolicyItem.RetentionUnit}}{{end}}
`
var updateWatchTemplate = `Your backup compliance policy has been updated with the following policies:
ID	FREQUENCY INTERVAL	FREQUENCY TYPE	RETENTION
{{- range .ScheduledPolicyItems}}
{{.Id}}	{{if eq .FrequencyType "hourly"}}{{.FrequencyInterval}}{{else}}-{{end}}	{{.FrequencyType}}	{{.RetentionValue}} {{.RetentionUnit}}
{{- end}}
{{if .OnDemandPolicyItem}}{{.OnDemandPolicyItem.Id}}	-	{{.OnDemandPolicyItem.FrequencyType}}	{{.OnDemandPolicyItem.RetentionValue}} {{.OnDemandPolicyItem.RetentionUnit}}{{end}}
`
var example = `How to run atlas backups compliancepolicy policies update with --file.
As an example, the file should be in the following format:

{
	"frequencyInterval": 1,
	"frequencyType": "daily",
	"id": "stringstringstringstring",
	"retentionUnit": "days",
	"retentionValue": 0
}

To get the ID of a policy item, run "atlas backups compliancepolicy policies describe".

Finally, run the command as such: "atlas backups compliancepolicy policies update --file /path/to/file"
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

func (opts *UpdateOpts) interactiveRun() error {

	projectID, err := opts.askProjectOptions()
	if err != nil {
		return fmt.Errorf("asking for projectID failed: %w", err)
	}
	opts.projectID = projectID

	compliancePolicy, err := opts.store.DescribeCompliancePolicy(projectID)
	if err != nil {
		return fmt.Errorf("couldn't fetch the backup compliance policy: %w", err)
	}

	item, err := opts.askPolicyOptions(compliancePolicy)
	if err != nil {
		return fmt.Errorf("asking for policy item failed: %w", err)
	}

	snapshotInterval, err := opts.askForSnapshotInterval(item)
	if err != nil {
		return fmt.Errorf("asking for the snapshot interval failed: %w", err)
	}
	item.SetFrequencyInterval(snapshotInterval)

	retentionUnit, retentionValue, err := opts.askForRetention(item)
	if err != nil {
		return fmt.Errorf("asking for retention data failed: %w", err)
	}
	item.SetRetentionValue(retentionValue)
	item.SetRetentionUnit(retentionUnit)

	return opts.Run(item)
}

func (opts *UpdateOpts) Run(policyItem *atlasv2.DiskBackupApiPolicyItem) error {
	result, httpResponse, err := opts.store.UpdatePolicyItem(opts.projectID, policyItem)
	if err != nil {
		if httpResponse != nil && httpResponse.StatusCode == 500 {
			return fmt.Errorf("%v: %w", errorCode500Template, err)
		}
		return err
	}

	if opts.EnableWatch {
		watcher := watcher.CompliancePolicyWatcherFactory(opts.projectID, opts.store, result)
		err := opts.Watch(watcher)
		if err != nil {
			return fmt.Errorf("received an error while watching for completion: %w", err)
		}
		opts.Template = updateWatchTemplate
	}
	return opts.Print(result)
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
		Example: example,
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return opts.PreRunE(
				opts.ValidateProjectID,
				opts.initStore(cmd.Context()),
				opts.InitOutput(cmd.OutOrStdout(), updateTemplate),
			)
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			if opts.path != "" {
				opts.projectID = opts.ConfigProjectID()
				policyItem := &atlasv2.DiskBackupApiPolicyItem{}
				if err := file.Load(opts.fs, opts.path, policyItem); err != nil {
					return err
				}
				return opts.Run(policyItem)
			}
			return opts.interactiveRun()
		},
	}

	cmd.Flags().StringVar(&opts.ProjectID, flag.ProjectID, "", usage.ProjectID)
	cmd.Flags().StringVarP(&opts.Output, flag.Output, flag.OutputShort, "", usage.FormatOut)
	_ = cmd.RegisterFlagCompletionFunc(flag.Output, opts.AutoCompleteOutputFlag())
	cmd.Flags().StringVarP(&opts.path, flag.File, flag.FileShort, "", usage.BackupCompliancePolicyItemFile)

	return cmd
}
