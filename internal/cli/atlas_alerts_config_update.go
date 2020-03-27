// Copyright 2020 MongoDB Inc
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

package cli

import (
	"github.com/mongodb/mongocli/internal/description"
	"github.com/mongodb/mongocli/internal/flags"
	"github.com/mongodb/mongocli/internal/json"
	"github.com/mongodb/mongocli/internal/store"
	"github.com/mongodb/mongocli/internal/usage"
	"github.com/spf13/cobra"
)

type atlasAlertsConfigUpdateOpts struct {
	*globalOpts
	*atlasAlertsConfigOpts
	store   store.AlertConfigurationUpdater
	alertID string
}

func (opts *atlasAlertsConfigUpdateOpts) init() error {
	if opts.ProjectID() == "" {
		return errMissingProjectID
	}

	var err error
	opts.store, err = store.New()
	return err
}

func (opts *atlasAlertsConfigUpdateOpts) Run() error {
	alert := opts.newAlertConfiguration(opts.ProjectID())
	alert.ID = opts.alertID
	result, err := opts.store.UpdateAlertConfiguration(alert)

	if err != nil {
		return err
	}

	return json.PrettyPrint(result)
}

// mongocli atlas alerts config(s) update [--event event] [--enabled enabled][--matcherField fieldName --matcherOperator operator --matcherValue value]
// [--notificationType type --notificationDelayMin min --notificationEmailEnabled --notificationSmsEnabled --notificationUsername username --notificationTeamID id
// [--notificationEmailAddress email --notificationMobileNumber number --notificationChannelName channel --notificationApiToken --notificationRegion region]
// [--projectId projectId]
func AtlasAlertsConfigUpdateBuilder() *cobra.Command {
	opts := &atlasAlertsConfigUpdateOpts{
		globalOpts:            newGlobalOpts(),
		atlasAlertsConfigOpts: newAtlasAlertsConfigOpts(),
	}
	cmd := &cobra.Command{
		Use:     "update",
		Short:   description.UpdateConfig,
		Aliases: []string{"updates"},
		Args:    cobra.ExactArgs(1),
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return opts.init()
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			opts.alertID = args[0]
			return opts.Run()
		},
	}

	cmd.Flags().StringVar(&opts.event, flags.Event, "", usage.Event)
	cmd.Flags().BoolVar(&opts.enabled, flags.Enabled, false, usage.Enabled)
	cmd.Flags().StringVar(&opts.matcherFieldName, flags.MatcherFieldName, "", usage.MatcherFieldName)
	cmd.Flags().StringVar(&opts.matcherOperator, flags.MatcherOperator, "", usage.MatcherOperator)
	cmd.Flags().StringVar(&opts.matcherValue, flags.MatcherValue, "", usage.MatcherValue)
	cmd.Flags().StringVar(&opts.metricThresholdMetricName, flags.MetricName, "", usage.MetricName)
	cmd.Flags().StringVar(&opts.metricThresholdOperator, flags.MetricOperator, "", usage.MetricOperator)
	cmd.Flags().Float64Var(&opts.metricThresholdThreshold, flags.MetricThreshold, 0, usage.MetricThreshold)
	cmd.Flags().StringVar(&opts.metricThresholdUnits, flags.MetricUnits, "", usage.MetricUnits)
	cmd.Flags().StringVar(&opts.metricThresholdMode, flags.MetricMode, "", usage.MetricMode)
	cmd.Flags().StringVar(&opts.notificationToken, flags.NotificationToken, "", usage.NotificationToken)
	cmd.Flags().StringVar(&opts.notificationChannelName, flags.NotificationChannelName, "", usage.NotificationsChannelName)
	cmd.Flags().StringVar(&opts.apiKey, flags.APIKey, "", usage.APIKey)
	cmd.Flags().StringVar(&opts.notificationRegion, flags.NotificationRegion, "", usage.NotificationRegion)
	cmd.Flags().IntVar(&opts.notificationDelayMin, flags.NotificationDelayMin, 0, usage.NotificationDelayMin)
	cmd.Flags().StringVar(&opts.notificationEmailAddress, flags.NotificationEmailAddress, "", usage.NotificationEmailAddress)
	cmd.Flags().BoolVar(&opts.notificationEmailEnabled, flags.NotificationEmailEnabled, false, usage.NotificationEmailEnabled)
	cmd.Flags().StringVar(&opts.notificationFlowName, flags.NotificationFlowName, "", usage.NotificationFlowName)
	cmd.Flags().IntVar(&opts.notificationIntervalMin, flags.NotificationIntervalMin, 0, usage.NotificationIntervalMin)
	cmd.Flags().StringVar(&opts.notificationMobileNumber, flags.NotificationMobileNumber, "", usage.NotificationMobileNumber)
	cmd.Flags().StringVar(&opts.notificationOrgName, flags.NotificationOrgName, "", usage.NotificationOrgName)
	cmd.Flags().StringVar(&opts.notificationServiceKey, flags.NotificationServiceKey, "", usage.NotificationServiceKey)
	cmd.Flags().BoolVar(&opts.notificationSmsEnabled, flags.NotificationSmsEnabled, false, usage.NotificationSmsEnabled)
	cmd.Flags().StringVar(&opts.notificationTeamID, flags.NotificationTeamID, "", usage.NotificationTeamID)
	cmd.Flags().StringVar(&opts.notificationType, flags.NotificationType, "", usage.NotificationType)
	cmd.Flags().StringVar(&opts.notificationUsername, flags.NotificationUsername, "", usage.NotificationUsername)
	cmd.Flags().StringVar(&opts.notificationVictorOpsRoutingKey, flags.NotificationVictorOpsRoutingKey, "", usage.NotificationVictorOpsRoutingKey)

	cmd.Flags().StringVar(&opts.projectID, flags.ProjectID, "", usage.ProjectID)

	return cmd
}
