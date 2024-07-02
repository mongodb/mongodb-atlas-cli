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
	"strings"

	"github.com/spf13/cobra"
	atlasv2 "go.mongodb.org/atlas-sdk/v20240530002/admin"
)

const (
	datadog        = "DATADOG"
	slack          = "SLACK"
	victor         = "VICTOR_OPS"
	email          = "EMAIL"
	opsGenie       = "OPS_GENIE"
	pagerDuty      = "PAGER_DUTY"
	sms            = "SMS"
	group          = "GROUP"
	user           = "USER"
	org            = "ORG"
	team           = "TEAM"
	webhook        = "WEBHOOK"
	microsoftTeams = "MICROSOFT_TEAMS"
)

// ConfigOpts contains all the information and functions to manage an alert configuration.
type ConfigOpts struct {
	event                           string
	matcherFieldName                string
	matcherOperator                 string
	matcherValue                    string
	metricThresholdMetricName       string
	metricThresholdOperator         string
	metricThresholdUnits            string
	metricThresholdMode             string
	notifierID                      string
	notificationToken               string // notificationsApiToken, notificationsFlowdockApiToken
	notificationChannelName         string
	apiKey                          string // notificationsDatadogApiKey, notificationsOpsGenieApiKey, notificationsVictorOpsApiKey
	notificationEmailAddress        string
	notificationMobileNumber        string
	notificationRegion              string // notificationsOpsGenieRegion, notificationsDatadogRegion
	notificationServiceKey          string
	notificationTeamID              string
	notificationType                string
	notificationUsername            string
	notificationVictorOpsRoutingKey string
	notificationWebhookURL          string
	notificationWebhookSecret       string
	notificationRoles               []string
	notificationDelayMin            int
	notificationIntervalMin         int
	notificationSmsEnabled          bool
	enabled                         bool
	notificationEmailEnabled        bool
	metricThresholdThreshold        float64
}

func (opts *ConfigOpts) NewAlertConfiguration(projectID string) *atlasv2.GroupAlertsConfig {
	out := new(atlasv2.GroupAlertsConfig)

	out.GroupId = &projectID
	eventType := strings.ToUpper(opts.event)
	out.EventTypeName = &eventType
	out.Enabled = &opts.enabled

	if opts.matcherFieldName != "" {
		out.Matchers = &[]atlasv2.StreamsMatcher{opts.newMatcher()}
	}

	if opts.metricThresholdMetricName != "" {
		out.MetricThreshold = opts.newMetricThreshold()
	}

	notification := opts.newNotification()
	out.Notifications = &[]atlasv2.AlertsNotificationRootForGroup{*notification}

	return out
}

func (opts *ConfigOpts) newNotification() *atlasv2.AlertsNotificationRootForGroup {
	out := new(atlasv2.AlertsNotificationRootForGroup)
	notificationType := strings.ToUpper(opts.notificationType)
	out.TypeName = &notificationType
	out.DelayMin = &opts.notificationDelayMin
	out.IntervalMin = &opts.notificationIntervalMin

	if opts.notifierID != "" {
		out.NotifierId = &opts.notifierID
	}

	// write set of functions that contain code from switch cases and return error if required fields are not provided
	switch out.GetTypeName() {
	case datadog:
		out.DatadogApiKey = &opts.apiKey
		region := strings.ToUpper(opts.notificationRegion)
		out.DatadogRegion = &region
	case email:
		out.EmailAddress = &opts.notificationEmailAddress
	case group, org:
		out.EmailEnabled = &opts.notificationEmailEnabled
		out.SmsEnabled = &opts.notificationSmsEnabled
		if len(opts.notificationRoles) > 0 {
			out.Roles = &opts.notificationRoles
		}
	case microsoftTeams:
		out.MicrosoftTeamsWebhookUrl = &opts.notificationWebhookURL
	case opsGenie:
		out.OpsGenieApiKey = &opts.apiKey
		region := strings.ToUpper(opts.notificationRegion)
		out.OpsGenieRegion = &region
	case pagerDuty:
		region := strings.ToUpper(opts.notificationRegion)
		out.Region = &region
		out.ServiceKey = &opts.notificationServiceKey
	case slack:
		out.ApiToken = &opts.notificationToken
		out.ChannelName = &opts.notificationChannelName
	case sms:
		out.MobileNumber = &opts.notificationMobileNumber
	case team:
		out.EmailEnabled = &opts.notificationEmailEnabled
		out.SmsEnabled = &opts.notificationSmsEnabled
		out.TeamId = &opts.notificationTeamID
	case user:
		out.EmailEnabled = &opts.notificationEmailEnabled
		out.SmsEnabled = &opts.notificationSmsEnabled
		out.Username = &opts.notificationUsername
	case victor:
		out.VictorOpsApiKey = &opts.apiKey
		out.VictorOpsRoutingKey = &opts.notificationVictorOpsRoutingKey
	case webhook:
		out.WebhookUrl = &opts.notificationWebhookURL
		out.WebhookSecret = &opts.notificationWebhookSecret
	}

	return out
}

func (opts *ConfigOpts) newMetricThreshold() *atlasv2.ServerlessMetricThreshold {
	operator := strings.ToUpper(opts.metricThresholdOperator)
	mode := strings.ToUpper(opts.metricThresholdMode)
	units := strings.ToUpper(opts.metricThresholdUnits)
	result := &atlasv2.ServerlessMetricThreshold{
		MetricName: strings.ToUpper(opts.metricThresholdMetricName),
		Operator:   &operator,
		Threshold:  &opts.metricThresholdThreshold,
		Units:      &units,
		Mode:       &mode,
	}

	return result
}

func (opts *ConfigOpts) newMatcher() atlasv2.StreamsMatcher {
	return atlasv2.StreamsMatcher{
		FieldName: strings.ToUpper(opts.matcherFieldName),
		Operator:  strings.ToUpper(opts.matcherOperator),
		Value:     strings.ToUpper(opts.matcherValue),
	}
}

func Builder() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "settings",
		Aliases: []string{"config"},
		Short:   "Manages alerts configuration for your project.",
		Long:    `Use this command to list, create, edit, delete, enable and disable alert configurations.`,
	}

	cmd.AddCommand(
		CreateBuilder(),
		ListBuilder(),
		describeBuilder(),
		DeleteBuilder(),
		FieldsBuilder(),
		UpdateBuilder(),
		EnableBuilder(),
		DisableBuilder(),
	)

	return cmd
}
