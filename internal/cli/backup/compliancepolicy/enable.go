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
	"net/mail"

	"github.com/AlecAivazis/survey/v2"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/config"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/flag"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/store"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/telemetry"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/usage"
	"github.com/spf13/cobra"
	atlasv2 "go.mongodb.org/atlas-sdk/v20240530002/admin"
)

type EnableOpts struct {
	cli.GlobalOpts
	cli.WatchOpts
	policy                  *atlasv2.DataProtectionSettings20231001
	store                   store.CompliancePolicyEnabler
	authorizedUserFirstName string
	authorizedUserLastName  string
	authorizedEmail         string
	confirm                 bool
}

var enableConfirmationMessage = `Backup compliance policy can not be disabled without MongoDB Support. Please confirm that you want to continue.
Learn more: https://www.mongodb.com/docs/atlas/backup/cloud-backup/backup-compliance-policy/
`

var enableWatchTemplate = `Backup Compliance Policy enabled without any configuration. Run "atlas backups compliancepolicy --help" for configuration options.
`
var enableTemplate = `Backup Compliance Policy is being enabled without any configuration. Run "atlas backups compliancepolicy --help" for configuration options.
`

func (opts *EnableOpts) initStore(ctx context.Context) func() error {
	return func() error {
		var err error
		opts.store, err = store.New(store.AuthenticatedPreset(config.Default()), store.WithContext(ctx))
		return err
	}
}

func (opts *EnableOpts) enableWatcher() (any, bool, error) {
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

func newEnableConfirmationQuestion() survey.Prompt {
	return &survey.Confirm{
		Message: enableConfirmationMessage,
		Default: false,
	}
}

func (opts *EnableOpts) Run() error {
	if _, err := mail.ParseAddress(opts.authorizedEmail); err != nil {
		return fmt.Errorf("unable to enable compliance policy due to invalid email: %w", err)
	}
	if !opts.confirm {
		question := newEnableConfirmationQuestion()
		var confirmation bool
		if err := telemetry.TrackAskOne(question, &confirmation); err != nil {
			return fmt.Errorf("couldn't confirm action: %w", err)
		}
		if !confirmation {
			return errors.New("did not receive confirmation to enable backup compliance policy")
		}
	}
	compliancePolicy, err := opts.store.EnableCompliancePolicy(opts.ConfigProjectID(), opts.authorizedEmail, opts.authorizedUserFirstName, opts.authorizedUserLastName)
	opts.policy = compliancePolicy
	if err != nil {
		return fmt.Errorf("couldn't enable compliance policy: %w", err)
	}
	if opts.EnableWatch {
		if _, err := opts.Watch(opts.enableWatcher); err != nil {
			return err
		}
		opts.Template = enableWatchTemplate
	}

	return opts.Print(opts.policy)
}

func EnableBuilder() *cobra.Command {
	opts := new(EnableOpts)
	cmd := &cobra.Command{
		Use:   "enable",
		Short: "Enable Backup Compliance Policy without any configuration.",
		PreRunE: func(cmd *cobra.Command, _ []string) error {
			return opts.PreRunE(
				opts.ValidateProjectID,
				opts.initStore(cmd.Context()),
				opts.InitOutput(cmd.OutOrStdout(), enableTemplate),
			)
		},
		RunE: func(_ *cobra.Command, _ []string) error {
			return opts.Run()
		},
	}

	cmd.Flags().StringVar(&opts.ProjectID, flag.ProjectID, "", usage.ProjectID)
	cmd.Flags().StringVar(&opts.authorizedUserFirstName, flag.AuthorizedUserFirstName, "", usage.AuthorizedUserFirstName)
	_ = cmd.MarkFlagRequired(flag.AuthorizedUserFirstName)
	cmd.Flags().StringVar(&opts.authorizedUserLastName, flag.AuthorizedUserLastName, "", usage.AuthorizedUserLastName)
	_ = cmd.MarkFlagRequired(flag.AuthorizedUserLastName)
	cmd.Flags().StringVar(&opts.authorizedEmail, flag.AuthorizedEmail, "", usage.AuthorizedEmail)
	_ = cmd.MarkFlagRequired(flag.AuthorizedEmail)
	cmd.Flags().BoolVarP(&opts.EnableWatch, flag.EnableWatch, flag.EnableWatchShort, false, usage.EnableWatchDefault)
	cmd.Flags().StringVarP(&opts.Output, flag.Output, flag.OutputShort, "", usage.FormatOut)
	cmd.Flags().BoolVar(&opts.confirm, flag.Force, false, usage.Force)
	_ = cmd.RegisterFlagCompletionFunc(flag.Output, opts.AutoCompleteOutputFlag())
	return cmd
}
