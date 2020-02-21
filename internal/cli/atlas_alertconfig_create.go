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
	enabled                         bool
	matcherFieldName                string
	matcherOperator                 string
	matcherValue                    string
	metricThresholdMetricName       string
	metricThresholdOperator         string
	metricThresholdThreshold        float64
	metricThresholdUnits            string
	metricThresholdMode             string
	token                           string // notificationsApiToken, notificationsFlowdockApiToken
	notificationChannelName         string
	apiKey                          string // notificationsDatadogApiKey, notificationsOpsGenieApiKey, notificationsVictorOpsApiKey
	notificationDelayMin            int
	notificationEmailAddress        string
	notificationEmailEnabled        bool
	notificationFlowName            string
	notificationIntervalMin         int
	notificationMobileNumber        string
	notificationRegion              string // notificationsOpsGenieRegion, notificationsDatadogRegion
	notificationOrgName             string
	notificationServiceKey          string
	notificationSmsEnabled          bool
	notificationTeamId              string
	notificationTypeName            string
	notificationUsername            string
	notificationVictorOpsRoutingKey string
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

	alert := new(atlas.AlertConfiguration)

	alert.GroupID = opts.ProjectID()
	alert.EventTypeName = strings.ToUpper(opts.event)
	alert.Enabled = &opts.enabled

	if opts.matcherFieldName != "" {
		match := new(atlas.Matcher)
		match.FieldName = strings.ToUpper(opts.matcherFieldName)
		match.Operator = strings.ToUpper(opts.matcherOperator)
		match.Value = strings.ToUpper(opts.matcherValue)
		alert.Matchers = []atlas.Matcher{*match}
	}

	if opts.metricThresholdMetricName != "" {
		metric := new(atlas.MetricThreshold)
		metric.MetricName = strings.ToUpper(opts.metricThresholdMetricName)
		metric.Operator = strings.ToUpper(opts.metricThresholdOperator)
		metric.Threshold = opts.metricThresholdThreshold
		metric.Units = strings.ToUpper(opts.metricThresholdUnits)
		metric.Mode = strings.ToUpper(opts.metricThresholdMode)
		alert.MetricThreshold = metric
	}

	notification := new(atlas.Notification)
	notification.TypeName = strings.ToUpper(opts.notificationTypeName)
	notification.DelayMin = &opts.notificationDelayMin
	notification.IntervalMin = opts.notificationIntervalMin
	notification.TeamID = opts.notificationTeamId
	notification.Username = opts.notificationUsername

	switch notification.TypeName {

	case victor:
		notification.VictorOpsAPIKey = opts.apiKey
		notification.VictorOpsRoutingKey = opts.notificationVictorOpsRoutingKey

	case slack:
		notification.VictorOpsAPIKey = opts.apiKey
		notification.VictorOpsRoutingKey = opts.notificationVictorOpsRoutingKey

	case datadog:
		notification.DatadogAPIKey = opts.apiKey
		notification.DatadogRegion = strings.ToUpper(opts.notificationRegion)

	case email:
		notification.EmailAddress = opts.notificationEmailAddress

	case flowdock:
		notification.FlowdockAPIToken = opts.token
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

	alert.Notifications = []atlas.Notification{*notification}
	return alert
}

// mcli atlas alert-config(s) create -event event --enabled [--matcherField fieldName --matcherOperator operator --matcherValue value]
// [--notificationType type --notificationDelayMin min --notificationEmailEnabled --notificationSmsEnabled --notificationUsername username --notificationTeamId id
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
	cmd.Flags().StringVar(&opts.token, flags.Token, "", usage.Token)
	cmd.Flags().StringVar(&opts.notificationChannelName, flags.NotificationsChannelName, "", usage.NotificationsChannelName)
	cmd.Flags().StringVar(&opts.apiKey, flags.ApiKey, "", usage.ApiKey)
	cmd.Flags().StringVar(&opts.notificationRegion, flags.NotificationsRegion, "", usage.NotificationRegion)
	cmd.Flags().IntVar(&opts.notificationDelayMin, flags.NotificationsDelayMin, 0, usage.NotificationDelayMin)
	cmd.Flags().StringVar(&opts.notificationEmailAddress, flags.NotificationsEmailAddress, "", usage.NotificationEmailAddress)
	cmd.Flags().BoolVar(&opts.notificationEmailEnabled, flags.NotificationsEmailEnabled, false, usage.NotificationEmailEnabled)
	cmd.Flags().StringVar(&opts.notificationFlowName, flags.NotificationsFlowName, "", usage.NotificationFlowName)
	cmd.Flags().IntVar(&opts.notificationIntervalMin, flags.NotificationsIntervalMin, 0, usage.NotificationIntervalMin)
	cmd.Flags().StringVar(&opts.notificationMobileNumber, flags.NotificationsMobileNumber, "", usage.NotificationMobileNumber)
	cmd.Flags().StringVar(&opts.notificationOrgName, flags.NotificationsOrgName, "", usage.NotificationOrgName)
	cmd.Flags().StringVar(&opts.notificationServiceKey, flags.NotificationsServiceKey, "", usage.NotificationServiceKey)
	cmd.Flags().BoolVar(&opts.notificationSmsEnabled, flags.NotificationsSmsEnabled, false, usage.NotificationSmsEnabled)
	cmd.Flags().StringVar(&opts.notificationTeamId, flags.NotificationsTeamId, "", usage.NotificationTeamId)
	cmd.Flags().StringVar(&opts.notificationTypeName, flags.NotificationsTypeName, "", usage.NotificationTypeName)
	cmd.Flags().StringVar(&opts.notificationUsername, flags.NotificationsUsername, "", usage.NotificationUsername)
	cmd.Flags().StringVar(&opts.notificationVictorOpsRoutingKey, flags.NotificationsVictorOpsRoutingKey, "", usage.NotificationVictorOpsRoutingKey)

	cmd.Flags().StringVar(&opts.projectID, flags.ProjectID, "", usage.ProjectID)

	return cmd
}

/*
curl -X POST -u "dozqoqjw:6e89fd16-ebfd-4584-9f3b-d3b692326983" --digest "https://cloud-dev.mongodb.com/api/atlas/v1.0/groups/5e4e593f70dfbf1010295836/alertConfigs" \
   -H "Content-Type: application/json" --data '
   {
	 "groupId": "5e4e593f70dfbf1010295836",
     "eventTypeName" : "NO_PRIMARY",
     "enabled" : true,
     "notifications" : [ {
       "typeName" : "GROUP",
       "intervalMin" : 5,
       "delayMin" : 0,
       "smsEnabled" : false,
       "emailEnabled" : true
     } ]
   }'
*/
