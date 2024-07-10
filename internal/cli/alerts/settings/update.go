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

package settings

import (
	"context"
	"fmt"

	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli/require"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/config"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/file"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/flag"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/store"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/usage"
	"github.com/spf13/afero"
	"github.com/spf13/cobra"
	atlasv2 "go.mongodb.org/atlas-sdk/v20240530002/admin"
)

type UpdateOpts struct {
	cli.GlobalOpts
	cli.OutputOpts
	ConfigOpts
	store    store.AlertConfigurationUpdater
	alertID  string
	filename string
	fs       afero.Fs
}

func (opts *UpdateOpts) initStore(ctx context.Context) func() error {
	return func() error {
		var err error
		opts.store, err = store.New(store.AuthenticatedPreset(config.Default()), store.WithContext(ctx))
		return err
	}
}

var updateTemplate = "Alert configuration '{{.Id}}' updated.\n"

func (opts *UpdateOpts) Run() error {
	alert := &atlasv2.GroupAlertsConfig{}

	// File flag has priority over other flags
	projectID := opts.ConfigProjectID()
	if opts.filename != "" {
		if err := file.Load(opts.fs, opts.filename, alert); err != nil {
			return err
		}
		alert.GroupId = &projectID
	} else {
		alert = opts.NewAlertConfiguration(projectID)
	}
	alert.Id = &opts.alertID
	r, err := opts.store.UpdateAlertConfiguration(alert)
	if err != nil {
		return err
	}

	return opts.Print(r)
}

// atlas alerts config(s) update <ID> [--event event] [--enabled enabled][--matcherField fieldName --matcherOperator operator --matcherValue value]
// [--notificationType type --notificationDelayMin min --notificationEmailEnabled --notificationSmsEnabled --notificationUsername username --notificationTeamID id
// [--notificationEmailAddress email --notificationMobileNumber number --notificationChannelName channel --notificationApiToken --notificationRegion region]
// [--projectId projectId].
func UpdateBuilder() *cobra.Command {
	opts := UpdateOpts{fs: afero.NewOsFs()}
	cmd := &cobra.Command{
		Use:   "update <alertConfigId>",
		Short: "Modify the details of the specified alert configuration for your project.",
		Long:  fmt.Sprintf(usage.RequiredRole, "Project Owner"),
		Args:  require.ExactArgs(1),
		Annotations: map[string]string{
			"alertConfigIdDesc": "Unique identifier of the alert configuration you want to update.",
			"output":            updateTemplate,
		},
		Example: `  # Modify the alert configuration with the ID 5d1113b25a115342acc2d1aa so that it notifies a user when they join a group for the project with the ID 5df90590f10fab5e33de2305:
  atlas alerts settings update 5d1113b25a115342acc2d1aa --event JOINED_GROUP --enabled \
		--notificationType USER --notificationEmailEnabled \
		--notificationIntervalMin 60 --notificationUsername john@example.com \
		--output json --projectId 5df90590f10fab5e33de2305
  # Update alert using json file input containing alert configuration
  atlas alerts settings update 5d1113b25a115342acc2d1aa --file alerts.json`,
		PreRunE: func(cmd *cobra.Command, _ []string) error {
			return opts.PreRunE(
				opts.ValidateProjectID,
				func() error {
					if opts.filename == "" {
						return validateConfigOpts(&opts.ConfigOpts)
					}
					return nil
				},
				opts.initStore(cmd.Context()),
				opts.InitOutput(cmd.OutOrStdout(), updateTemplate),
			)
		},
		RunE: func(_ *cobra.Command, args []string) error {
			opts.alertID = args[0]
			return opts.Run()
		},
	}

	cmd.Flags().StringVar(&opts.event, flag.Event, "", usage.Event)
	cmd.Flags().BoolVar(&opts.enabled, flag.Enabled, false, usage.Enabled)

	cmd.Flags().StringVar(&opts.matcherFieldName, flag.MatcherFieldName, "", usage.MatcherFieldName)
	cmd.Flags().StringVar(&opts.matcherOperator, flag.MatcherOperator, "", usage.MatcherOperator)
	cmd.Flags().StringVar(&opts.matcherValue, flag.MatcherValue, "", usage.MatcherValue)
	cmd.MarkFlagsRequiredTogether(flag.MatcherOperator, flag.MatcherValue, flag.MatcherFieldName)

	cmd.Flags().StringVar(&opts.metricThresholdMetricName, flag.MetricName, "", usage.MetricName)
	cmd.Flags().StringVar(&opts.metricThresholdOperator, flag.MetricOperator, "", usage.MetricOperator)
	cmd.Flags().Float64Var(&opts.metricThresholdThreshold, flag.MetricThreshold, 0, usage.MetricThreshold)
	cmd.Flags().StringVar(&opts.metricThresholdUnits, flag.MetricUnits, "", usage.MetricUnits)
	cmd.Flags().StringVar(&opts.metricThresholdMode, flag.MetricMode, "", usage.MetricMode)
	cmd.MarkFlagsRequiredTogether(flag.MetricOperator, flag.MetricName, flag.MetricUnits, flag.MetricThreshold)

	cmd.Flags().StringVar(&opts.notificationToken, flag.NotificationToken, "", usage.NotificationToken)
	cmd.Flags().StringVar(&opts.notificationChannelName, flag.NotificationChannelName, "", usage.NotificationsChannelName)
	cmd.Flags().StringVar(&opts.apiKey, flag.APIKey, "", usage.AlertConfigAPIKey)
	cmd.Flags().StringVar(&opts.notificationRegion, flag.NotificationRegion, "", usage.NotificationRegion)
	cmd.Flags().IntVar(&opts.notificationDelayMin, flag.NotificationDelayMin, 0, usage.NotificationDelayMin)
	cmd.Flags().StringVar(&opts.notificationEmailAddress, flag.NotificationEmailAddress, "", usage.NotificationEmailAddress)
	cmd.Flags().BoolVar(&opts.notificationEmailEnabled, flag.NotificationEmailEnabled, false, usage.NotificationEmailEnabled)
	cmd.Flags().IntVar(&opts.notificationIntervalMin, flag.NotificationIntervalMin, 0, usage.NotificationIntervalMin)
	cmd.Flags().StringVar(&opts.notificationMobileNumber, flag.NotificationMobileNumber, "", usage.NotificationMobileNumber)
	cmd.Flags().StringVar(&opts.notificationServiceKey, flag.NotificationServiceKey, "", usage.NotificationServiceKey)
	cmd.Flags().BoolVar(&opts.notificationSmsEnabled, flag.NotificationSmsEnabled, false, usage.NotificationSmsEnabled)
	cmd.Flags().StringVar(&opts.notificationTeamID, flag.NotificationTeamID, "", usage.NotificationTeamID)
	cmd.Flags().StringVar(&opts.notificationType, flag.NotificationType, "", usage.NotificationTypeAtlas)
	cmd.Flags().StringVar(&opts.notificationUsername, flag.NotificationUsername, "", usage.NotificationUsername)
	cmd.Flags().StringVar(&opts.notificationVictorOpsRoutingKey, flag.NotificationVictorOpsRoutingKey, "", usage.NotificationVictorOpsRoutingKey)
	cmd.Flags().StringVar(&opts.notificationWebhookURL, flag.NotificationWebhookURL, "", usage.NotificationWebhookURL)
	cmd.Flags().StringVar(&opts.notificationWebhookSecret, flag.NotificationWebhookSecret, "", usage.NotificationWebhookSecret)
	cmd.Flags().StringVar(&opts.notifierID, flag.NotifierID, "", usage.NotifierID)
	cmd.Flags().StringSliceVar(&opts.notificationRoles, flag.NotificationRole, []string{}, usage.NotificationRole)
	cmd.Flags().StringVarP(&opts.filename, flag.File, flag.FileShort, "", usage.AlertConfigFilename)

	_ = cmd.MarkFlagFilename(flag.File)

	cmd.Flags().StringVar(&opts.ProjectID, flag.ProjectID, "", usage.ProjectID)
	cmd.Flags().StringVarP(&opts.Output, flag.Output, flag.OutputShort, "", usage.FormatOut)
	_ = cmd.RegisterFlagCompletionFunc(flag.Output, opts.AutoCompleteOutputFlag())

	return cmd
}
