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
}

func (opts *ConfigOpts) NewAlertConfiguration(projectID string) *admin.AlertConfigViewForNdsGroup {
	out := new(admin.AlertConfigViewForNdsGroup)

	out.GroupId = &projectID
	out.EventTypeName, _ = admin.NewServerlessEventTypeViewAlertableFromValue(strings.ToUpper(opts.event))
	out.Enabled = &opts.enabled

	if opts.matcherFieldName != "" {
		out.Matchers = []map[string]interface{}{opts.newMatcher()}
	}

	if opts.metricThresholdMetricName != "" {
		out.MetricThreshold = opts.newMetricThreshold()
	}

	out.Notifications = []admin.NotificationViewForNdsGroup{opts.newNotification()}

	return out
}

func (opts *ConfigOpts) newNotification() admin.NotificationViewForNdsGroup {
	out := admin.NotificationViewForNdsGroup{}
	typeName := strings.ToUpper(opts.notificationType)

	switch typeName {
	case victor:
		out.VictorOpsNotification.VictorOpsApiKey = &opts.apiKey
		out.VictorOpsNotification.VictorOpsRoutingKey = &opts.notificationVictorOpsRoutingKey
		out.VictorOpsNotification.TypeName = typeName
		out.VictorOpsNotification.DelayMin = &opts.notificationDelayMin
		out.VictorOpsNotification.IntervalMin = &opts.notificationIntervalMin

	case slack:
		out.SlackNotification.ApiToken = &opts.apiKey
		out.SlackNotification.TypeName = typeName
		out.SlackNotification.DelayMin = &opts.notificationDelayMin
		out.SlackNotification.IntervalMin = &opts.notificationIntervalMin
		out.SlackNotification.ChannelName = &opts.notificationChannelName

	case datadog:
		out.DatadogNotification.DatadogApiKey = &opts.apiKey
		out.DatadogNotification.DatadogRegion = pointer.Get(strings.ToUpper(opts.notificationRegion))
		out.DatadogNotification.TypeName = typeName
		out.DatadogNotification.DelayMin = &opts.notificationDelayMin
		out.DatadogNotification.IntervalMin = &opts.notificationIntervalMin

	case email:
		out.EmailNotification.EmailAddress = &opts.notificationEmailAddress
		out.EmailNotification.TypeName = typeName
		out.EmailNotification.DelayMin = &opts.notificationDelayMin
		out.EmailNotification.IntervalMin = &opts.notificationIntervalMin

	case sms:
		out.SMSNotification.MobileNumber = &opts.notificationMobileNumber
		out.SMSNotification.TypeName = typeName
		out.SMSNotification.DelayMin = &opts.notificationDelayMin
		out.SMSNotification.IntervalMin = &opts.notificationIntervalMin

	case group:
		out.GroupNotification.SmsEnabled = &opts.notificationSmsEnabled
		out.GroupNotification.EmailEnabled = &opts.notificationEmailEnabled
		out.GroupNotification.TypeName = typeName
		out.GroupNotification.DelayMin = &opts.notificationDelayMin
		out.GroupNotification.IntervalMin = &opts.notificationIntervalMin

	case user:
		out.UserNotification.SmsEnabled = &opts.notificationSmsEnabled
		out.UserNotification.EmailEnabled = &opts.notificationEmailEnabled
		out.UserNotification.TypeName = typeName
		out.UserNotification.DelayMin = &opts.notificationDelayMin
		out.UserNotification.IntervalMin = &opts.notificationIntervalMin
		out.UserNotification.Username = &opts.notificationUsername

	case org:
		out.OrgNotification.SmsEnabled = &opts.notificationSmsEnabled
		out.OrgNotification.EmailEnabled = &opts.notificationEmailEnabled
		out.OrgNotification.TypeName = typeName
		out.OrgNotification.DelayMin = &opts.notificationDelayMin
		out.OrgNotification.IntervalMin = &opts.notificationIntervalMin

	case ops:
		out.OpsGenieNotification.OpsGenieApiKey = &opts.apiKey
		out.OpsGenieNotification.OpsGenieRegion = &opts.notificationRegion
		out.OpsGenieNotification.TypeName = typeName
		out.OpsGenieNotification.DelayMin = &opts.notificationDelayMin
		out.OpsGenieNotification.IntervalMin = &opts.notificationIntervalMin

	case pager:
		out.PagerDutyNotification.ServiceKey = &opts.notificationServiceKey
		out.PagerDutyNotification.Region = &opts.notificationRegion
		out.PagerDutyNotification.TypeName = typeName
		out.PagerDutyNotification.DelayMin = &opts.notificationDelayMin
		out.PagerDutyNotification.IntervalMin = &opts.notificationIntervalMin

	case team:
		out.TeamNotification.EmailEnabled = &opts.notificationEmailEnabled
		out.TeamNotification.SmsEnabled = &opts.notificationSmsEnabled
		out.TeamNotification.TypeName = typeName
		out.TeamNotification.DelayMin = &opts.notificationDelayMin
		out.TeamNotification.IntervalMin = &opts.notificationIntervalMin
		out.TeamNotification.TeamId = &opts.notificationTeamID

	}

	return out
}

func (opts *ConfigOpts) newMetricThreshold() *admin.ServerlessMetricThreshold {
	metricName := strings.ToUpper(opts.metricThresholdMetricName)
	operator, _ := admin.NewOperatorFromValue(strings.ToUpper(opts.metricThresholdOperator))
	mode := strings.ToUpper(opts.metricThresholdMode)
	result := &admin.ServerlessMetricThreshold{}
	switch metricName {
	case "DATA":
		result.DataMetricThreshold = &admin.DataMetricThreshold{
			MetricName: &metricName,
			Operator:   operator,
			Threshold:  &opts.metricThresholdThreshold,
			Units:      pointer.Get(admin.DataMetricUnits(strings.ToUpper(opts.metricThresholdUnits))),
			Mode:       &mode,
		}
	case "RPU":
		result.RPUMetricThreshold = &admin.RPUMetricThreshold{
			MetricName: &metricName,
			Operator:   operator,
			Threshold:  &opts.metricThresholdThreshold,
			Units:      pointer.Get(admin.ServerlessMetricUnits(strings.ToUpper(opts.metricThresholdUnits))),
			Mode:       &mode,
		}
	case "RAW":
		result.RawMetricThreshold = &admin.RawMetricThreshold{
			MetricName: &metricName,
			Operator:   operator,
			Threshold:  &opts.metricThresholdThreshold,
			Units:      pointer.Get(admin.RawMetricUnits(strings.ToUpper(opts.metricThresholdUnits))),
			Mode:       &mode,
		}
	case "TIME":
		result.TimeMetricThreshold = &admin.TimeMetricThreshold{
			MetricName: &metricName,
			Operator:   operator,
			Threshold:  &opts.metricThresholdThreshold,
			Units:      pointer.Get(admin.TimeMetricUnits(strings.ToUpper(opts.metricThresholdUnits))),
			Mode:       &mode,
		}

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
