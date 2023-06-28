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

	"github.com/mongodb/mongodb-atlas-cli/internal/pointer"
	"github.com/spf13/cobra"
	"go.mongodb.org/atlas-sdk/admin"
)

const (
	datadog = "DATADOG"
	slack   = "SLACK"
	victor  = "VICTOR_OPS"
	email   = "EMAIL"
	ops     = "OPS_GENIE"
	pager   = "PAGER_DUTY"
	sms     = "SMS"
	group   = "GROUP"
	user    = "USER"
	org     = "ORG"
	team    = "TEAM"
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
	notificationDelayMin            int
	notificationIntervalMin         int
	notificationSmsEnabled          bool
	enabled                         bool
	notificationEmailEnabled        bool
	metricThresholdThreshold        float64
}

func (opts *ConfigOpts) NewAlertConfiguration(projectID string) *admin.GroupAlertsConfig {
	out := new(admin.GroupAlertsConfig)

	out.GroupId = &projectID
	out.EventTypeName = pointer.Get(strings.ToUpper(opts.event))
	out.Enabled = &opts.enabled

	if opts.matcherFieldName != "" {
		out.Matchers = []map[string]interface{}{opts.newMatcher()}
	}

	if opts.metricThresholdMetricName != "" {
		out.MetricThreshold = opts.newMetricThreshold()
	}

	out.Notifications = []admin.AlertsNotificationRootForGroup{*opts.newNotification()}

	return out
}

func (opts *ConfigOpts) newNotification() *admin.AlertsNotificationRootForGroup {
	out := new(admin.AlertsNotificationRootForGroup)
	out.TypeName = pointer.Get(strings.ToUpper(opts.notificationType))
	out.DelayMin = &opts.notificationDelayMin
	out.IntervalMin = &opts.notificationIntervalMin
	out.TeamId = &opts.notificationTeamID
	out.Username = &opts.notificationUsername
	out.ChannelName = &opts.notificationChannelName

	switch out.GetTypeName() {
	case victor:
		out.VictorOpsApiKey = &opts.apiKey
		out.VictorOpsRoutingKey = &opts.notificationVictorOpsRoutingKey
	case slack:
		out.VictorOpsApiKey = &opts.apiKey
		out.VictorOpsRoutingKey = &opts.notificationVictorOpsRoutingKey
		out.ApiToken = &opts.notificationToken
	case datadog:
		out.DatadogApiKey = &opts.apiKey
		out.DatadogRegion = pointer.Get(strings.ToUpper(opts.notificationRegion))
	case email:
		out.EmailAddress = &opts.notificationEmailAddress
	case sms:
		out.MobileNumber = &opts.notificationMobileNumber
	case group, user, org:
		out.SmsEnabled = &opts.notificationSmsEnabled
		out.EmailEnabled = &opts.notificationEmailEnabled
	case ops:
		out.OpsGenieApiKey = &opts.apiKey
		out.OpsGenieRegion = &opts.notificationRegion
	case pager:
		out.ServiceKey = &opts.notificationServiceKey
	}

	return out
}

func (opts *ConfigOpts) newMetricThreshold() *admin.ServerlessMetricThreshold {
	metricName := strings.ToUpper(opts.metricThresholdMetricName)
	operator := strings.ToUpper(opts.metricThresholdOperator)
	mode := strings.ToUpper(opts.metricThresholdMode)
	units := strings.ToUpper(opts.metricThresholdUnits)
	result := &admin.ServerlessMetricThreshold{
		MetricName: &metricName,
		Operator:   &operator,
		Threshold:  &opts.metricThresholdThreshold,
		Units:      &units,
		Mode:       &mode,
	}

	return result
}

func (opts *ConfigOpts) newMatcher() map[string]interface{} {
	result := make(map[string]interface{})
	result["FieldName"] = strings.ToUpper(opts.matcherFieldName)
	result["Operator"] = strings.ToUpper(opts.matcherOperator)
	result["Value"] = strings.ToUpper(opts.matcherValue)
	return result
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
		DeleteBuilder(),
		FieldsBuilder(),
		UpdateBuilder(),
		EnableBuilder(),
		DisableBuilder(),
	)

	return cmd
}
