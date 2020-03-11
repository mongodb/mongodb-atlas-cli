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
	"github.com/mongodb/mongocli/internal/flags"
	"github.com/mongodb/mongocli/internal/json"
	"github.com/mongodb/mongocli/internal/store"
	"github.com/mongodb/mongocli/internal/usage"
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

	out := new(atlas.AlertConfiguration)

	out.GroupID = opts.ProjectID()
	out.EventTypeName = strings.ToUpper(opts.event)
	out.Enabled = &opts.enabled

	if opts.matcherFieldName != "" {
		out.Matchers = []atlas.Matcher{*newMatcher(opts)}
	}

	if opts.metricThresholdMetricName != "" {
		out.MetricThreshold = newMetricThreshold(opts)
	}

	out.Notifications = []atlas.Notification{*newNotification(opts)}

	return out
}

func newNotification(opts *atlasAlertConfigCreateOpts) *atlas.Notification {

	out := new(atlas.Notification)
	out.TypeName = strings.ToUpper(opts.notificationType)
	out.DelayMin = &opts.notificationDelayMin
	out.IntervalMin = opts.notificationIntervalMin
	out.TeamID = opts.notificationTeamID
	out.Username = opts.notificationUsername

	switch out.TypeName {

	case victor:
		out.VictorOpsAPIKey = opts.apiKey
		out.VictorOpsRoutingKey = opts.notificationVictorOpsRoutingKey

	case slack:
		out.VictorOpsAPIKey = opts.apiKey
		out.VictorOpsRoutingKey = opts.notificationVictorOpsRoutingKey
		out.APIToken = opts.notificationToken

	case datadog:
		out.DatadogAPIKey = opts.apiKey
		out.DatadogRegion = strings.ToUpper(opts.notificationRegion)

	case email:
		out.EmailAddress = opts.notificationEmailAddress

	case flowdock:
		out.FlowdockAPIToken = opts.notificationToken
		out.FlowName = opts.notificationFlowName
		out.OrgName = opts.notificationOrgName

	case sms:
		out.MobileNumber = opts.notificationMobileNumber

	case group, user, org:
		out.SMSEnabled = &opts.notificationSmsEnabled
		out.EmailEnabled = &opts.notificationEmailEnabled

	case ops:
		out.OpsGenieAPIKey = opts.apiKey
		out.OpsGenieRegion = opts.notificationRegion

	case pager:
		out.ServiceKey = opts.notificationServiceKey

	}

	return out
}

func newMetricThreshold(opts *atlasAlertConfigCreateOpts) *atlas.MetricThreshold {
	return &atlas.MetricThreshold{
		MetricName: strings.ToUpper(opts.metricThresholdMetricName),
		Operator:   strings.ToUpper(opts.metricThresholdOperator),
		Threshold:  opts.metricThresholdThreshold,
		Units:      strings.ToUpper(opts.metricThresholdUnits),
		Mode:       strings.ToUpper(opts.metricThresholdMode),
	}
}

func newMatcher(opts *atlasAlertConfigCreateOpts) *atlas.Matcher {
	return &atlas.Matcher{
		FieldName: strings.ToUpper(opts.matcherFieldName),
		Operator:  strings.ToUpper(opts.matcherOperator),
		Value:     strings.ToUpper(opts.matcherValue),
	}
}

// mongocli atlas alerts config(s) create -event event --enabled [--matcherField fieldName --matcherOperator operator --matcherValue value]
// [--notificationType type --notificationDelayMin min --notificationEmailEnabled --notificationSmsEnabled --notificationUsername username --notificationTeamID id
// --notificationEmailAddress email --notificationMobileNumber number --notificationChannelName channel --notificationApiToken --notificationRegion region] [--projectId projectId]
func AtlasAlertConfigCreateBuilder() *cobra.Command {
	opts := &atlasAlertConfigCreateOpts{
		globalOpts: newGlobalOpts(),
	}
	cmd := &cobra.Command{
		Use:   "create",
		Short: "Create an alert configuration for a project.",
		Args:  cobra.NoArgs,
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
