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
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/AlecAivazis/survey/v2"
	atlas "github.com/mongodb/go-client-mongodb-atlas/mongodbatlas"
	"github.com/mongodb/mongocli/internal/config"
	"github.com/mongodb/mongocli/internal/prompts"
	"github.com/mongodb/mongocli/internal/validate"
)

const (
	fallbackSuccessMessage = "'%s' deleted\n"
	fallbackFailMessage    = "entry not deleted"
)

type globalOpts struct {
	orgID     string
	projectID string
}

func deploymentStatus(baseURL, projectID string) string {
	return fmt.Sprintf("Changes are being applied, please check %sv2/%s#deployment/topology for status\n", baseURL, projectID)
}

// ProjectID returns the project id.
// If the id is empty, it caches it after querying config.
func (opts *globalOpts) ProjectID() string {
	if opts.projectID != "" {
		return opts.projectID
	}
	opts.projectID = config.ProjectID()
	return opts.projectID
}

type cmdOpt func() error

// PreRunE is a function to call before running the command,
// this will validate the project ID and call any additional function pass as a callback
func (opts *globalOpts) PreRunE(cbs ...cmdOpt) error {
	if opts.ProjectID() == "" {
		return errMissingProjectID
	}
	if err := validate.ObjectID(opts.ProjectID()); err != nil {
		return err
	}
	for _, f := range cbs {
		if err := f(); err != nil {
			return err
		}
	}

	return nil
}

// OrgID returns the organization id.
// If the id is empty, it caches it after querying config.
func (opts *globalOpts) OrgID() string {
	if opts.orgID != "" {
		return opts.orgID
	}
	opts.orgID = config.OrgID()
	return opts.orgID
}

// deleteOpts options required when deleting a resource.
// A command can compose this struct and then safely rely on the methods Confirm, or Delete
// to manage the interactions with the user
type deleteOpts struct {
	entry          string
	confirm        bool
	successMessage string
	failMessage    string
}

// Delete deletes a resource not associated to a project, it expects a callback
// that should perform the deletion from the store.
func (opts *deleteOpts) Delete(d interface{}, a ...string) error {
	if !opts.confirm {
		fmt.Println(opts.FailMessage())
		return nil
	}

	var err error
	switch f := d.(type) {
	case func(string) error:
		err = f(opts.entry)
	case func(string, string) error:
		err = f(a[0], opts.entry)
	case func(string, string, string) error:
		err = f(a[0], a[1], opts.entry)
	default:
		return errors.New("invalid")
	}

	if err != nil {
		return err
	}

	fmt.Printf(opts.SuccessMessage(), opts.entry)

	return nil
}

// Confirm confirms that the resource should be deleted
func (opts *deleteOpts) Confirm() error {
	if opts.confirm {
		return nil
	}

	prompt := prompts.NewDeleteConfirm(opts.entry)
	return survey.AskOne(prompt, &opts.confirm)
}

// SuccessMessage gets the set success message or the default value
func (opts *deleteOpts) SuccessMessage() string {
	if opts.successMessage != "" {
		return opts.successMessage
	}
	return fallbackSuccessMessage
}

// FailMessage gets the set fail message or the default value
func (opts *deleteOpts) FailMessage() string {
	if opts.failMessage != "" {
		return opts.failMessage
	}
	return fallbackFailMessage
}

// atlasAlertsConfigOpts contains all the information and functions to manage an alert configuration
type atlasAlertsConfigOpts struct {
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

func (opts *atlasAlertsConfigOpts) newAlertConfiguration(projectID string) *atlas.AlertConfiguration {
	out := new(atlas.AlertConfiguration)

	out.GroupID = projectID
	out.EventTypeName = strings.ToUpper(opts.event)
	out.Enabled = &opts.enabled

	if opts.matcherFieldName != "" {
		out.Matchers = []atlas.Matcher{*opts.newMatcher()}
	}

	if opts.metricThresholdMetricName != "" {
		out.MetricThreshold = opts.newMetricThreshold()
	}

	out.Notifications = []atlas.Notification{*opts.newNotification()}

	return out
}

func (opts *atlasAlertsConfigOpts) newNotification() *atlas.Notification {
	out := new(atlas.Notification)
	out.TypeName = strings.ToUpper(opts.notificationType)
	out.DelayMin = &opts.notificationDelayMin
	out.IntervalMin = opts.notificationIntervalMin
	out.TeamID = opts.notificationTeamID
	out.Username = opts.notificationUsername
	out.ChannelName = opts.notificationChannelName

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

func (opts *atlasAlertsConfigOpts) newMetricThreshold() *atlas.MetricThreshold {
	return &atlas.MetricThreshold{
		MetricName: strings.ToUpper(opts.metricThresholdMetricName),
		Operator:   strings.ToUpper(opts.metricThresholdOperator),
		Threshold:  opts.metricThresholdThreshold,
		Units:      strings.ToUpper(opts.metricThresholdUnits),
		Mode:       strings.ToUpper(opts.metricThresholdMode),
	}
}

func (opts *atlasAlertsConfigOpts) newMatcher() *atlas.Matcher {
	return &atlas.Matcher{
		FieldName: strings.ToUpper(opts.matcherFieldName),
		Operator:  strings.ToUpper(opts.matcherOperator),
		Value:     strings.ToUpper(opts.matcherValue),
	}
}

type listOpts struct {
	pageNum      int
	itemsPerPage int
}

func (opts *listOpts) newListOptions() *atlas.ListOptions {
	return &atlas.ListOptions{
		PageNum:      opts.pageNum,
		ItemsPerPage: opts.itemsPerPage,
	}
}

// getHostnameAndPort return the hostname and the port starting from the string hostname:port
func getHostnameAndPort(hostInfo string) (hostname string, port int, err error) {
	host := strings.Split(hostInfo, ":")
	if len(host) != 2 {
		return "", 0, fmt.Errorf("expected hostname:port, got %s", host)
	}

	port, err = strconv.Atoi(host[1])
	if err != nil {
		return "", 0, fmt.Errorf("invalid port number, got %s", host[1])
	}

	return host[0], port, nil
}
