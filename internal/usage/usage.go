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
	Shards                          = "Number of shards in the cluster."
	Tier                            = "Tier for each data-bearing server in the cluster."
	DiskSizeGB                      = "Capacity, in gigabytes, of the host’s root volume."
	Backup                          = "If true, enables Continuous Cloud Backup for your cluster."
	MDBVersion                      = "MongoDB version of the cluster to deploy."
	AuthDB                          = "Authentication database name."
	Granularity                     = "Duration in ISO 8601 notation that specifies the interval between measurement data points."
	Page                            = "Page number."
	Forever                         = "Acknowledge an alert “forever”."
	Status                          = "Alert's status."
	Until                           = "Acknowledged until a date."
	Limit                           = "Number of items per page."
	Username                        = "Username for authenticating to MongoDB."
	Password                        = "User’s password." //nolint:gosec // This is just a message not a password
	Period                          = "Duration in ISO 8601 notation that specifies how far back in the past to retrieve measurements."
	Roles                           = "User's roles and the databases or collections on which the roles apply."
	DataLakeRole                    = "Amazon Resource Name (ARN) of the role which Atlas Data Lake uses for accessing the data stores."
	DataLakeRegion                  = "Name of the region to which Data Lake routes client connections for data processing."
	Comment                         = "Optional description or comment for the entry."
	DeleteAfter                     = "ISO-8601-formatted UTC date after which Atlas removes the entry from the whitelist."
	Force                           = "Don't ask for confirmation."
	Email                           = "User’s email address."
	LogOut                          = "Optional output filename, if none given will use the log name."
	DiagnoseOut                     = "Optional output filename, if none given will use diagnose-archive.tar.gz."
	LogStart                        = "Beginning of the period for which to retrieve logs."
	LogEnd                          = "End of the period for which to retrieve logs."
	ArchiveLimit                    = "Max number of entries for the diagnose archive."
	ArchiveMinutes                  = "Beginning of the period for which to retrieve diagnose archive. Ops Manager takes out minutes from the current time. "
	MeasurementStart                = "Beginning of the period for which to retrieve measurements."
	MeasurementEnd                  = "End of the period for which to retrieve measurements."
	MeasurementType                 = "Measurements to return. If it is not specified, all measurements are returned."
	FirstName                       = "User’s first name."
	LastName                        = "User’s last name."
	MaxDate                         = "Returns events whose created date is less than or equal to it."
	MinDate                         = "Returns events whose created date is greater than or equal to it."
	Filename                        = "Filename to use, optional file with a json cluster configuration."
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
	NotificationToken               = "Slack API token or Bot token or Flowdock personal API token." //nolint:gosec // This is just a message not a password
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
	Database                        = "Database name."
	Collection                      = "Collection name."
	RSName                          = "The replica set that the index is built on."
	Key                             = "Index keys. Should be formatted as field:type."
	Unique                          = "Create a unique key index."
	LogTypes                        = "Array of strings specifying the types of logs to collect."
	SizeRequestedPerFileBytes       = "Size for each log file in bytes."
	LogRedacted                     = "If set to true, emails, hostnames, IP addresses, and namespaces in API responses involving this job are replaced with random string values."
	Sparse                          = "Create a sparse index."
	Locale                          = "Locale that the ICU defines."
	CaseLevel                       = "If set to true, the index uses case comparison. This field applies only if the strength level is set to 1 or 2."
	CaseFirst                       = "Determines the sort order of case differences during tertiary level comparisons. "
	Strength                        = "Level of comparison to perform."
	Alternate                       = "Determines whether collation should consider whitespace and punctuation as base characters during comparisons."
	MaxVariable                     = "Determines which characters are are considered ignorable. This field applies only if indexConfigs.collation.alternate is set to shifted."
	NumericOrdering                 = "If set to true, collation compares numeric strings as numbers. If false, collation compares numeric strings as strings."
	Normalization                   = "If true, collation checks if text requires normalization and performs normalization to compare text."
	Backwards                       = "If true, strings with diacritics sort from the back to the front of the string."
	ClusterName                     = "Name of the cluster."
	Verbose                         = "If true, returns all child jobs in the response."
	ClusterID                       = "Unique identifier of the cluster."
	Background                      = "Create the index in the background."
	DateField                       = "Name of an already indexed date field from the documents."
	PartitionFields                 = "Fields to use to partition data. You can specify up to two frequently queried fields to use for partitioning data."
	ArchiveAfter                    = "Number of days that specifies the age limit for the data in the live Atlas cluster."
	TargetProjectID                 = "Unique identifier of the project that contains the destination cluster for the restore job."
	TargetClusterID                 = `Unique identifier of the target cluster.
For use only with automated restore jobs.`
	TargetClusterName = `Name of the target cluster.
For use only with automated restore jobs.`
	CheckpointID = `Unique identifier for the sharded cluster checkpoint that represents the point in time to which your data will be restored.
If you set checkpointId, you cannot set oplogInc, oplogTs, snapshotId, or pointInTimeUTCMillis.`
	OplogTS = `Oplog timestamp given as a timestamp in the number of seconds that have elapsed since the UNIX epoch. 
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
Valid values: cidrBlock|ipAddress|awsSecurityGroup`
	Service = `Type of MongoDB service.
Valid values: cloud|cloud-manager|ops-manager`
	Provider = `Name of your cloud service provider.
Valid values: AWS|AZURE|GCP.`
	ClusterTypes = `Type of the cluster that you want to create.
Valid values: REPLICASET|SHARDED.`
	DataLakeTestBucket = `Name of an S3 data bucket which Data Lake uses to validate the provided role.`
	Region             = `Physical location of your MongoDB cluster.
For a complete list of supported AWS regions, see: https://docs.atlas.mongodb.com/reference/amazon-aws/#amazon-aws
For a complete list of supported Azure regions, see: https://docs.atlas.mongodb.com/reference/microsoft-azure/#microsoft-azure
For a complete list of supported GCP regions, see: https://docs.atlas.mongodb.com/reference/google-gcp/#google-gcp`
)
