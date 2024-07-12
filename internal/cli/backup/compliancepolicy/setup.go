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
	"fmt"

	"github.com/AlecAivazis/survey/v2"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/config"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/file"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/flag"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/store"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/telemetry"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/usage"
	"github.com/spf13/afero"
	"github.com/spf13/cobra"
	atlasv2 "go.mongodb.org/atlas-sdk/v20240530002/admin"
)

type SetupOpts struct {
	cli.GlobalOpts
	cli.WatchOpts
	policy  *atlasv2.DataProtectionSettings20231001
	store   store.CompliancePolicyUpdater
	fs      afero.Fs
	path    string
	confirm bool
}

const bcpTemplate = `Project:	{{.ProjectId}}
Authorized user first name:	{{.AuthorizedUserFirstName}}
Authorized user last name:	{{.AuthorizedUserLastName}}
Authorized e-mail:	{{.AuthorizedEmail}}
Copy protection enabled:	{{.CopyProtectionEnabled}}
Encryption at rest enabled:	{{.EncryptionAtRestEnabled}}
Point-in-Time restores enabled:	{{.PitEnabled}}
Restore window days:	{{.RestoreWindowDays}}

POLICIES
ID	FREQUENCY INTERVAL	FREQUENCY TYPE	RETENTION
{{- range valueOrEmptySlice .ScheduledPolicyItems}}
{{.Id}}	{{if eq .FrequencyType "hourly"}}{{.FrequencyInterval}}{{else}}-{{end}}	{{.FrequencyType}}	{{.RetentionValue}} {{.RetentionUnit}}
{{- end}}
{{if .OnDemandPolicyItem}}{{.OnDemandPolicyItem.Id}}	-	{{.OnDemandPolicyItem.FrequencyType}}	{{.OnDemandPolicyItem.RetentionValue}} {{.OnDemandPolicyItem.RetentionUnit}}{{end}}
`

const setupWatchTemplate = `Your backup compliance policy has been set up with the following configuration:
` + bcpTemplate

const setupTemplate = `Your backup compliance policy is being set up with the following configuration:
` + bcpTemplate

const confirmationMessage = `Backup compliance policy can not be disabled without MongoDB Support. Please confirm that you want to continue.
Learn more: https://www.mongodb.com/docs/atlas/backup/cloud-backup/backup-compliance-policy/
`

func (opts *SetupOpts) initStore(ctx context.Context) func() error {
	return func() error {
		var err error
		opts.store, err = store.New(store.AuthenticatedPreset(config.Default()), store.WithContext(ctx))
		return err
	}
}

func (opts *SetupOpts) setupWatcher() (any, bool, error) {
	res, err := opts.store.DescribeCompliancePolicy(opts.ConfigProjectID())
	if err != nil {
		return nil, false, err
	}
	opts.policy = res
	if res.GetState() == "" {
		return nil, false, errors.New("could not access State field")
	}
	return nil, res.GetState() == active, nil
}

func newSetupConfirmationQuestion() survey.Prompt {
	return &survey.Confirm{
		Message: confirmationMessage,
		Default: false,
	}
}

func (opts *SetupOpts) Run() error {
	if !opts.confirm {
		question := newSetupConfirmationQuestion()
		var confirmation bool
		if err := telemetry.TrackAskOne(question, &confirmation); err != nil {
			return fmt.Errorf("couldn't confirm action: %w", err)
		}
		if !confirmation {
			return errors.New("did not receive confirmation to enable backup compliance policy")
		}
	}
	_, err := opts.store.UpdateCompliancePolicy(opts.ConfigProjectID(), opts.policy)
	if err != nil {
		return err
	}
	if opts.EnableWatch {
		if _, err := opts.Watch(opts.setupWatcher); err != nil {
			return err
		}
		opts.Template = setupWatchTemplate
	}

	return opts.Print(opts.policy)
}

func SetupBuilder() *cobra.Command {
	opts := &SetupOpts{
		policy: new(atlasv2.DataProtectionSettings20231001),
		fs:     afero.NewOsFs(),
	}
	use := "setup"
	cmd := &cobra.Command{
		Use:     use,
		Aliases: cli.GenerateAliases(use),
		Short:   "Setup the backup compliance policy for your project with a configuration file.",
		PreRunE: func(cmd *cobra.Command, _ []string) error {
			return opts.PreRunE(
				opts.ValidateProjectID,
				opts.initStore(cmd.Context()),
				opts.InitOutput(cmd.OutOrStdout(), setupTemplate),
			)
		},
		RunE: func(_ *cobra.Command, _ []string) error {
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
	_ = cmd.MarkFlagRequired(flag.File)
	return cmd
}
