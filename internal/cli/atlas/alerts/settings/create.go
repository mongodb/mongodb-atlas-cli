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

	"github.com/mongodb/mongodb-atlas-cli/internal/cli"
	"github.com/mongodb/mongodb-atlas-cli/internal/config"
	"github.com/mongodb/mongodb-atlas-cli/internal/flag"
	store "github.com/mongodb/mongodb-atlas-cli/internal/store/atlas"
	"github.com/mongodb/mongodb-atlas-cli/internal/usage"
	"github.com/spf13/cobra"
)

type CreateOpts struct {
	cli.GlobalOpts
	cli.OutputOpts
	ConfigOpts
	store store.AlertConfigurationCreator
}

func (opts *CreateOpts) initStore(ctx context.Context) func() error {
	return func() error {
		var err error
		opts.store, err = store.New(store.AuthenticatedPreset(config.Default()), store.WithContext(ctx))
		return err
	}
}

var createTemplate = "Alert configuration {{.Id}} created.\n"

func (opts *CreateOpts) Run() error {
	alert := opts.NewAlertConfiguration(opts.ConfigProjectID())
	r, err := opts.store.CreateAlertConfiguration(alert)
	if err != nil {
		return err
	}

	return opts.Print(r)
}

// CreateBuilder atlas alerts config(s) create
//
//	[--event event]
//	[--enabled enabled]
//	[--matcherField fieldName --matcherOperator operator --matcherValue value]
//	[--notificationType type --notificationDelayMin min --notificationEmailEnabled --notificationSmsEnabled --notificationUsername username --notificationTeamID id
//	[--notificationEmailAddress email --notificationMobileNumber number --notificationChannelName channel --notificationApiToken --notificationRegion region]
//	[--projectId projectId].
func CreateBuilder() *cobra.Command {
	opts := new(CreateOpts)
	opts.Template = createTemplate
	cmd := &cobra.Command{
		Use:   "create",
		Short: "Create an alert configuration for your project.",
		Long:  fmt.Sprintf(usage.RequiredRole, "Project Owner"),
		Annotations: map[string]string{
			"output": createTemplate,
		},
		Example: fmt.Sprintf(`  # Create an alert configuration that notifies a user when they join a group for the project with the ID 5df90590f10fab5e33de2305:
  %s alerts settings create --event JOINED_GROUP --enabled \
  --notificationType USER --notificationEmailEnabled \
  --notificationUsername john@example.com \
  --output json --projectId 5df90590f10fab5e33de2305`, cli.ExampleAtlasEntryPoint()),
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return opts.PreRunE(
				opts.ValidateProjectID,
				opts.initStore(cmd.Context()),
				opts.InitOutput(cmd.OutOrStdout(), createTemplate),
			)
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return opts.Run()
		},
	}

	cmd.Flags().StringVar(&opts.event, flag.Event, "", usage.Event)
	cmd.Flags().BoolVar(&opts.enabled, flag.Enabled, false, usage.Enabled)
	cmd.Flags().StringVar(&opts.matcherFieldName, flag.MatcherFieldName, "", usage.MatcherFieldName)
	cmd.Flags().StringVar(&opts.matcherOperator, flag.MatcherOperator, "", usage.MatcherOperator)
	cmd.Flags().StringVar(&opts.matcherValue, flag.MatcherValue, "", usage.MatcherValue)
	cmd.Flags().StringVar(&opts.metricThresholdMetricName, flag.MetricName, "", usage.MetricName)
	cmd.Flags().StringVar(&opts.metricThresholdOperator, flag.MetricOperator, "", usage.MetricOperator)
	cmd.Flags().Float64Var(&opts.metricThresholdThreshold, flag.MetricThreshold, 0, usage.MetricThreshold)
	cmd.Flags().StringVar(&opts.metricThresholdUnits, flag.MetricUnits, "", usage.MetricUnits)
	cmd.Flags().StringVar(&opts.metricThresholdMode, flag.MetricMode, "", usage.MetricMode)
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
	cmd.Flags().StringSliceVar(&opts.notificationRoles, flag.NotificationRole, []string{}, usage.NotificationRole)

	cmd.Flags().StringVar(&opts.ProjectID, flag.ProjectID, "", usage.ProjectID)
	cmd.Flags().StringVarP(&opts.Output, flag.Output, flag.OutputShort, "", usage.FormatOut)
	_ = cmd.RegisterFlagCompletionFunc(flag.Output, opts.AutoCompleteOutputFlag())
	_ = cmd.MarkFlagRequired(flag.Event)

	return cmd
}
