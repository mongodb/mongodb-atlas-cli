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

package usage

const (
	ProjectID                       = "Project ID to use. Overrides configuration file or environment variable settings."
	OrgID                           = "Organization ID to use. Overrides configuration file or environment variable settings."
	Profile                         = "Profile to use from your configuration file."
	Members                         = "Number of members in the replica set."
	InstanceSize                    = "Tier for each data-bearing server in the cluster."
	DiskSizeGB                      = "Capacity, in gigabytes, of the host’s root volume."
	Backup                          = "If true, uses Atlas Continuous Backups to back up cluster data."
	MDBVersion                      = "MongoDB version of the cluster to deploy."
	AuthDB                          = "Authentication database name."
	Page                            = "Page number."
	Until                           = "Acknowledged until a date."
	Limit                           = "Number of items per page."
	Username                        = "Username for authenticating to MongoDB."
	Password                        = "User’s password."
	Roles                           = "User's roles and the databases or collections on which the roles apply."
	Comment                         = "Optional description of the whitelist entry."
	Force                           = "Don't ask for confirmation."
	Email                           = "User’s email address."
	FirstName                       = "User’s first name."
	LastName                        = "User’s last name."
	Filename                        = "Filename to use"
	WhitelistIps                    = "IP addresses to add to the new user’s whitelist."
	Event                           = "Type of event that will trigger an alert."
	Enabled                         = "If set to true, the alert configuration is enabled."
	MatcherFieldName                = "Name of the field in the target object to match on."
	MatcherOperator                 = "The operator to test the field’s value."
	MatcherValue                    = "Value to test with the specified operator."
	MetricName                      = "Name of the metric against which Atlas checks the configured"
	MetricOperator                  = "Operator to apply when checking the current metric value against the threshold value."
	MetricThreshold                 = "Threshold value outside of which an alert will be triggered."
	MetricUnits                     = "The units for the threshold value."
	MetricMode                      = "Atlas computes the current metric value as an average."
	NotificationToken               = "Slack API token or Bot token or Flowdock personal API token."
	NotificationsChannelName        = "Slack channel name. Required for the SLACK notifications type."
	APIKey                          = "Datadog API Key, Opsgenie API Key, VictorOps API key."
	NotificationRegion              = "Region that indicates which API URL to use."
	NotificationDelayMin            = "Number of minutes to wait after an alert condition is detected before sending out the first notification."
	NotificationEmailAddress        = "Email address to which alert notifications are sent."
	NotificationEmailEnabled        = "Flag indicating if email notifications should be sent."
	NotificationFlowName            = "Flowdock flow name in lower-case letters."
	NotificationIntervalMin         = "Number of minutes to wait between successive notifications for unacknowledged alerts that are not resolved."
	NotificationMobileNumber        = "Mobile number to which alert notifications are sent."
	NotificationOrgName             = "Flowdock organization name in lower-case letters."
	NotificationServiceKey          = "PagerDuty service key."
	NotificationSmsEnabled          = "Flag indicating if text message notifications should be sent."
	NotificationTeamID              = "Unique identifier of a team."
	NotificationType                = "Type of alert notification."
	NotificationUsername            = "Name of the Atlas user to which to send notifications."
	NotificationVictorOpsRoutingKey = "VictorOps routing key."
	SnapshotID                      = "Unique identifier of the snapshot to restore."
	ClusterName                     = "Name of the cluster that contains the snapshots that you want to retrieve."
	ClusterID                       = "Unique identifier of the cluster that the job represents."
	TargetProjectID                 = "Unique identifier of the project that contains the destination cluster for the restore job."
	TargetClusterID                 = `Unique identifier of the target cluster.
For use only with automated restore jobs.`
	TargetClusterName = `Name of the target cluster.
For use only with automated restore jobs.`
	CheckpointID = `Unique identifier for the sharded cluster checkpoint that represents the point in time to which your data will be restored.
If you set checkpointId, you cannot set oplogInc, oplogTs, snapshotId, or pointInTimeUTCMillis.`
	OplogTs = `Oplog timestamp given as a timestamp in the number of seconds that have elapsed since the UNIX epoch. 
When paired with oplogInc, they represent the point in time to which your data will be restored.`
	OplogInc = `32-bit incrementing ordinal that represents operations within a given second. 
When paired with oplogTs, they represent the point in time to which your data will be restored.`
	PointInTimeUTCMillis = `Timestamp in the number of milliseconds that have elapsed since the UNIX epoch that represents the point in time to which your data will be restored.
This timestamp must be within last 24 hours of the current time.`
	Expires = `Timestamp in ISO 8601 date and time format after which the URL is no longer available.
For use only with download restore jobs.`
	ExpirationHours = `Number of hours the download URL is valid once the restore job is complete.
For use only with download restore jobs.`
	MaxDownloads = `Number of times the download URL can be used. This must be 1 or greater.
For use only with download restore jobs.`
	Mechanisms = `Authentication mechanism. 
Valid values: SCRAM-SHA-1|SCRAM-SHA-256`
	WhitelistType = `Type of whitelist entry.
Valid values: cidrBlock|ipAddress`
	Service = `Type of MongoDB service.
Valid values: cloud|cloud-manager|ops-manager`
	Provider = `Name of your cloud service provider.
Valid values: AWS|AZURE|GCP.`
	Region = `Physical location of your MongoDB cluster.
For a complete list of supported AWS regions, see: https://docs.atlas.mongodb.com/reference/amazon-aws/#amazon-aws
For a complete list of supported Azure regions, see: https://docs.atlas.mongodb.com/reference/microsoft-azure/#microsoft-azure
For a complete list of supported GCP regions, see: https://docs.atlas.mongodb.com/reference/google-gcp/#google-gcp`
)
