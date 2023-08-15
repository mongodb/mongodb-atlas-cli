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

package compliancepolicy

import (
	"context"
	"errors"

	"github.com/mongodb/mongodb-atlas-cli/internal/cli"
	"github.com/mongodb/mongodb-atlas-cli/internal/config"
	"github.com/mongodb/mongodb-atlas-cli/internal/file"
	"github.com/mongodb/mongodb-atlas-cli/internal/flag"
	store "github.com/mongodb/mongodb-atlas-cli/internal/store/atlas"
	"github.com/mongodb/mongodb-atlas-cli/internal/usage"
	"github.com/spf13/afero"
	"github.com/spf13/cobra"
	atlasv2 "go.mongodb.org/atlas-sdk/v20230201004/admin"
)

type SetupOpts struct {
	cli.GlobalOpts
	cli.WatchOpts
	policy      *atlasv2.DataProtectionSettings
	store       store.CompliancePolicy
	fs          afero.Fs
	path        string
	confirm     bool
	EnableWatch bool
}

var setupWatchTemplate = `Your backup compliance policy has been set up with the following configuration:

Project:	{{.ProjectId}}
Authorized e-mail:	{{.AuthorizedEmail}}
Copy protection enabled:	{{.CopyProtectionEnabled}}
Encryption at rest enabled:	{{.EncryptionAtRestEnabled}}
Point-in-Time restores enabled:	{{.PitEnabled}}
Restore window days:	{{.RestoreWindowDays}}

POLICIES
ID	FREQUENCY INTERVAL	FREQUENCY TYPE	RETENTION
{{- range .ScheduledPolicyItems}}
{{.Id}}	{{if eq .FrequencyType "hourly"}}{{.FrequencyInterval}}{{else}}-{{end}}	{{.FrequencyType}}	{{.RetentionValue}} {{.RetentionUnit}}
{{- end}}
{{if .OnDemandPolicyItem}}{{.OnDemandPolicyItem.Id}}	-	{{.OnDemandPolicyItem.FrequencyType}}	{{.OnDemandPolicyItem.RetentionValue}} {{.OnDemandPolicyItem.RetentionUnit}}{{end}}
`

var setupTemplate = `Your backup compliance policy is being set up with the following configuration:

Project:	{{.ProjectId}}
Authorized e-mail:	{{.AuthorizedEmail}}
Copy protection enabled:	{{.CopyProtectionEnabled}}
Encryption at rest enabled:	{{.EncryptionAtRestEnabled}}
Point-in-Time restores enabled:	{{.PitEnabled}}
Restore window days:	{{.RestoreWindowDays}}

POLICIES
ID	FREQUENCY INTERVAL	FREQUENCY TYPE	RETENTION
{{- range .ScheduledPolicyItems}}
{{.Id}}	{{if eq .FrequencyType "hourly"}}{{.FrequencyInterval}}{{else}}-{{end}}	{{.FrequencyType}}	{{.RetentionValue}} {{.RetentionUnit}}
{{- end}}
{{if .OnDemandPolicyItem}}{{.OnDemandPolicyItem.Id}}	-	{{.OnDemandPolicyItem.FrequencyType}}	{{.OnDemandPolicyItem.RetentionValue}} {{.OnDemandPolicyItem.RetentionUnit}}{{end}}
`

func (opts *SetupOpts) initStore(ctx context.Context) func() error {
	return func() error {
		var err error
		opts.store, err = store.New(store.AuthenticatedPreset(config.Default()), store.WithContext(ctx))
		return err
	}
}

func (opts *SetupOpts) setupWatcher() (bool, error) {
	res, err := opts.store.DescribeCompliancePolicy(opts.ConfigProjectID())
	opts.policy = res
	if err != nil {
		return false, err
	}
	if res.GetState() == "" {
		return false, errors.New("could not access State field")
	}
	return (res.GetState() == active), nil
}

func (opts *SetupOpts) Run() error {
	_, err := opts.store.UpdateCompliancePolicy(opts.ConfigProjectID(), opts.policy)
	if err != nil {
		return err
	}
	if opts.EnableWatch {
		if err := opts.Watch(opts.setupWatcher); err != nil {
			return err
		}
		opts.Template = setupWatchTemplate
	}

	return opts.Print(opts.policy)
}

func SetupBuilder() *cobra.Command {
	opts := &SetupOpts{
		policy: new(atlasv2.DataProtectionSettings),
		fs:     afero.NewOsFs(),
	}
	use := "setup"
	cmd := &cobra.Command{
		Use:     use,
		Aliases: cli.GenerateAliases(use),
		Short:   "Setup the backup compliance policy for your project with a configuration file.",
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return opts.PreRunE(
				opts.ValidateProjectID,
				opts.initStore(cmd.Context()),
				opts.InitOutput(cmd.OutOrStdout(), setupTemplate),
			)
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			if err := file.Load(opts.fs, opts.path, opts.policy); err != nil {
				return err
			}
			return opts.Run()
		},
	}

	cmd.Flags().StringVar(&opts.ProjectID, flag.ProjectID, "", usage.ProjectID)
	cmd.Flags().StringVarP(&opts.Output, flag.Output, flag.OutputShort, "", usage.FormatOut)
	_ = cmd.RegisterFlagCompletionFunc(flag.Output, opts.AutoCompleteOutputFlag())
	cmd.Flags().StringVarP(&opts.path, flag.File, flag.FileShort, "", usage.BackupCompliancePolicyFile)
	cmd.Flags().BoolVar(&opts.confirm, flag.Force, false, usage.Force)
	cmd.Flags().BoolVarP(&opts.EnableWatch, flag.EnableWatch, flag.EnableWatchShort, false, usage.EnableWatchDefault)
	_ = cmd.Flags().MarkHidden(flag.Force)
	_ = cmd.MarkFlagRequired(flag.File)
	return cmd
}
