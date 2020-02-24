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
	"strings"

	atlas "github.com/mongodb/go-client-mongodb-atlas/mongodbatlas"
	"github.com/mongodb/mcli/internal/flags"
	"github.com/mongodb/mcli/internal/json"
	"github.com/mongodb/mcli/internal/store"
	"github.com/mongodb/mcli/internal/usage"
	"github.com/spf13/cobra"
)

const (
	datadog  = "DATADOG"
	slack    = "SLACK"
	victor   = "VICTOR_OPS"
	flowdock = "FLOWDOCK"
	email    = "EMAIL"
	ops      = "OPS_GENIE"
	org      = "ORG"
	pager    = "PAGER_DUTY"
	sms      = "SMS"
	group    = "GROUP"
	user     = "USER"
)

type atlasAlertConfigCreateOpts struct {
	*globalOpts
	event                           string
	matcherFieldName                string
	matcherOperator                 string
	matcherValue                    string
	metricThresholdMetricName       string
	metricThresholdOperator         string
	metricThresholdUnits            string
	metricThresholdMode             string
	notificationToken               string // notificationsApiToken, notificationsFlowdockApiToken
	notificationChannelName         string
	apiKey                          string // notificationsDatadogApiKey, notificationsOpsGenieApiKey, notificationsVictorOpsApiKey
	notificationEmailAddress        string
	notificationFlowName            string
	notificationMobileNumber        string
	notificationRegion              string // notificationsOpsGenieRegion, notificationsDatadogRegion
	notificationOrgName             string
	notificationServiceKey          string
	notificationTeamID              string
	notificationType                string
	notificationUsername            string
	notificationVictorOpsRoutingKey string
	notificationDelayMin            int
	notificationIntervalMin         int
	notificationSmsEnabled          bool
	enabled                         bool
	notificationEmailEnabled        bool
	metricThresholdThreshold        float64
	store                           store.AlertConfigurationCreator
}

func (opts *atlasAlertConfigCreateOpts) init() error {
	if opts.ProjectID() == "" {
		return errMissingProjectID
	}

	var err error
	opts.store, err = store.New()
	return err
}

func (opts *atlasAlertConfigCreateOpts) Run() error {
	alert := opts.buildAlertConfiguration()
	result, err := opts.store.CreateAlertConfiguration(alert)

	if err != nil {
		return err
	}

	return json.PrettyPrint(result)
}

func (opts *atlasAlertConfigCreateOpts) buildAlertConfiguration() *atlas.AlertConfiguration {

	alertConfig := new(atlas.AlertConfiguration)

	alertConfig.GroupID = opts.ProjectID()
	alertConfig.EventTypeName = strings.ToUpper(opts.event)
	alertConfig.Enabled = &opts.enabled

	buildMatcher(opts, alertConfig)
	buildMetricThreshold(opts, alertConfig)
	buildNotification(opts, alertConfig)

	return alertConfig
}

func buildNotification(opts *atlasAlertConfigCreateOpts, alertConfig *atlas.AlertConfiguration) {

	notification := atlas.Notification{}
	notification.TypeName = strings.ToUpper(opts.notificationType)
	notification.DelayMin = &opts.notificationDelayMin
	notification.IntervalMin = opts.notificationIntervalMin
	notification.TeamID = opts.notificationTeamID
	notification.Username = opts.notificationUsername

	switch notification.TypeName {

	case victor:
		notification.VictorOpsAPIKey = opts.apiKey
		notification.VictorOpsRoutingKey = opts.notificationVictorOpsRoutingKey

	case slack:
		notification.VictorOpsAPIKey = opts.apiKey
		notification.VictorOpsRoutingKey = opts.notificationVictorOpsRoutingKey
		notification.APIToken = opts.notificationToken

	case datadog:
		notification.DatadogAPIKey = opts.apiKey
		notification.DatadogRegion = strings.ToUpper(opts.notificationRegion)

	case email:
		notification.EmailAddress = opts.notificationEmailAddress

	case flowdock:
		notification.FlowdockAPIToken = opts.notificationToken
		notification.FlowName = opts.notificationFlowName
		notification.OrgName = opts.notificationOrgName

	case sms:
		notification.MobileNumber = opts.notificationMobileNumber

	case group, user, org:
		notification.SMSEnabled = &opts.notificationSmsEnabled
		notification.EmailEnabled = &opts.notificationEmailEnabled

	case ops:
		notification.OpsGenieAPIKey = opts.apiKey
		notification.OpsGenieRegion = opts.notificationRegion

	case pager:
		notification.ServiceKey = opts.notificationServiceKey

	}

	alertConfig.Notifications = []atlas.Notification{notification}
}

func buildMetricThreshold(opts *atlasAlertConfigCreateOpts, alertConfig *atlas.AlertConfiguration) {
	if opts.metricThresholdMetricName != "" {
		metric := new(atlas.MetricThreshold)
		metric.MetricName = strings.ToUpper(opts.metricThresholdMetricName)
		metric.Operator = strings.ToUpper(opts.metricThresholdOperator)
		metric.Threshold = opts.metricThresholdThreshold
		metric.Units = strings.ToUpper(opts.metricThresholdUnits)
		metric.Mode = strings.ToUpper(opts.metricThresholdMode)
		alertConfig.MetricThreshold = metric
	}
}

func buildMatcher(opts *atlasAlertConfigCreateOpts, alertConfig *atlas.AlertConfiguration) {
	if opts.matcherFieldName != "" {
		match := new(atlas.Matcher)
		match.FieldName = strings.ToUpper(opts.matcherFieldName)
		match.Operator = strings.ToUpper(opts.matcherOperator)
		match.Value = strings.ToUpper(opts.matcherValue)
		alertConfig.Matchers = []atlas.Matcher{*match}
	}
}

// mcli atlas alert-config(s) create -event event --enabled [--matcherField fieldName --matcherOperator operator --matcherValue value]
// [--notificationType type --notificationDelayMin min --notificationEmailEnabled --notificationSmsEnabled --notificationUsername username --notificationTeamID id
// --notificationEmailAddress email --notificationMobileNumber number --notificationChannelName channel --notificationApiToken --notificationRegion region] [--projectId projectId]
func AtlasAlertConfigCreateBuilder() *cobra.Command {
	opts := &atlasAlertConfigCreateOpts{
		globalOpts: newGlobalOpts(),
	}
	cmd := &cobra.Command{
		Use:     "create",
		Short:   "Create an Atlas alert configuration for a project.",
		Aliases: []string{"cr", "Create"},
		Args:    cobra.NoArgs,
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return opts.init()
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return opts.Run()
		},
	}

	cmd.Flags().StringVar(&opts.event, flags.Event, "", usage.Event)
	cmd.Flags().BoolVar(&opts.enabled, flags.Enabled, false, usage.Enabled)
	cmd.Flags().StringVar(&opts.matcherFieldName, flags.MatcherFieldName, "", usage.MatcherFieldName)
	cmd.Flags().StringVar(&opts.matcherOperator, flags.MatcherOperator, "", usage.MatcherOperator)
	cmd.Flags().StringVar(&opts.matcherValue, flags.MatcherValue, "", usage.MatcherValue)
	cmd.Flags().StringVar(&opts.metricThresholdMetricName, flags.MetricThresholdMetricName, "", usage.MetricThresholdMetricName)
	cmd.Flags().StringVar(&opts.metricThresholdOperator, flags.MetricThresholdOperator, "", usage.MetricThresholdOperator)
	cmd.Flags().Float64Var(&opts.metricThresholdThreshold, flags.MetricThresholdThreshold, 0, usage.MetricThresholdThreshold)
	cmd.Flags().StringVar(&opts.metricThresholdUnits, flags.MetricThresholdUnits, "", usage.MetricThresholdUnits)
	cmd.Flags().StringVar(&opts.metricThresholdMode, flags.MetricThresholdMode, "", usage.MetricThresholdMode)
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
