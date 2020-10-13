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
	ProcessName                     = "The unique identifier for the host of a MongoDB process in the following format: {hostname}:{port}."
	Since                           = "Point in time, specified as milliseconds since the Unix Epoch, from which you want to receive results."
	HostID                          = "The unique identifier for the host of a MongoDB process. "
	Duration                        = "Length of time from the since parameter, in milliseconds, for which you want to receive results."
	Tier                            = "Tier for each data-bearing server in the cluster."
	NLog                            = "Maximum number of log lines to return."
	SlowQueryNamespaces             = "Namespaces from which to retrieve suggested slow query logs."
	DiskSizeGB                      = "Capacity, in gigabytes, of the host’s root volume."
	Backup                          = "If true, enables Continuous Cloud Backup for your cluster."
	SuggestedIndexNamespaces        = "Namespaces from which to retrieve suggested indexes. "
	NExamples                       = "Maximum number of examples queries to provide that will be improved by a suggested index."
	NIndexes                        = "Maximum number of indexes to suggest."
	MDBVersion                      = "MongoDB version of the cluster to deploy."
	AuthDB                          = "Authentication database name."
	Granularity                     = "Duration in ISO 8601 notation that specifies the interval between measurement data points."
	Page                            = "Page number."
	Forever                         = "Acknowledge an alert “forever”."
	Status                          = "Alert's status."
	Until                           = "Acknowledged until a date."
	Limit                           = "Number of items per page."
	Username                        = "Username of the user."
	BackupStatus                    = "Current (or desired) status of the backup configuration."
	StorageEngine                   = "Storage engine used for the backup."
	AuthMechanism                   = "Authentication mechanism needed to connect to the sync source database."
	Provisioned                     = "Flag that indicates if Ops Manager has provisioned the resources needed to store a backup."
	Encryption                      = "Flag that indicates if encryption is enabled for the backup configuration."
	SSL                             = "Flag that indicates if TLS is enabled for the sync source database."
	OplogSSL                        = "Flag indicating whether this oplog store only accepts connections encrypted using TLS."
	SyncSource                      = "mongod instance from which you retrieve backup data."
	ExcludedNamespace               = "List of database names and collection names to omit from the backup."
	IncludedNamespace               = "List of database names and collection names to include in the backup."
	TeamUsername                    = "List of usernames to add to the new team."
	DBUsername                      = "Username for authenticating to MongoDB."
	TeamName                        = "Name of the team."
	UserID                          = "The ID of the user."
	LDAPHostname                    = "The hostname or IP address of the LDAP server."
	LDAPPort                        = "The port to which the LDAP server listens for client connections."
	BindUsername                    = "The user DN that Atlas uses to connect to the LDAP server."
	BindPassword                    = "The password used to authenticate the bindUsername."
	CaCertificate                   = "CA certificate used to verify the identify of the LDAP server."
	AuthzQueryTemplate              = "An LDAP query template that Atlas executes to obtain the LDAP groups to which the authenticated user belongs."
	AuthenticationEnabled           = "Specifies whether user authentication with LDAP is enabled."
	AuthorizationEnabled            = "Specifies whether user authorization with LDAP is enabled."
	TeamID                          = "The ID of the team."
	Password                        = "User’s password." //nolint:gosec // This is just a message not a password
	Country                         = "The ISO 3166-1 alpha 2 country code of the user’s country of residence."
	Mobile                          = "The user’s mobile or cell phone number."
	Period                          = "Duration in ISO 8601 notation that specifies how far back in the past to retrieve measurements."
	Roles                           = "User's roles and the databases or collections on which the roles apply."
	DataLakeRole                    = "Amazon Resource Name (ARN) of the role which Atlas Data Lake uses for accessing the data stores."
	DataLakeRegion                  = "Name of the region to which Data Lake routes client connections for data processing."
	DataLakeTestBucket              = `Name of an S3 data bucket which Data Lake uses to validate the provided role.`
	PrivateEndpointRegion           = "Cloud provider region in which you want to create the private endpoint connection."
	PrivateEndpointProvider         = "Name of the cloud provider you want to create the private endpoint connection for."
	Comment                         = "Optional description or comment for the entry."
	AccessListsDeleteAfter          = "ISO-8601-formatted UTC date after which Atlas removes the entry from the entry."
	BDUsersDeleteAfter              = "Timestamp in ISO 8601 date and time format in UTC after which Atlas deletes the user."
	Force                           = "Don't ask for confirmation."
	ForceFile                       = "Overwrite the destination file."
	Email                           = "User’s email address."
	LogOut                          = "Optional output filename, if none given will use the log name."
	DiagnoseOut                     = "Optional output filename, if none given will use diagnose-archive.tar.gz."
	LogStart                        = "Beginning of the period for which to retrieve logs as UNIX Epoch time."
	LogEnd                          = "End of the period for which to retrieve logs as UNIX Epoch time."
	ArchiveLimit                    = "Max number of entries for the diagnose archive."
	ArchiveMinutes                  = "Beginning of the period for which to retrieve diagnose archive. Ops Manager takes out minutes from the current time. "
	MeasurementStart                = "Beginning of the period for which to retrieve measurements."
	MeasurementEnd                  = "End of the period for which to retrieve measurements."
	MeasurementType                 = "Measurements to return. If it is not specified, all measurements are returned."
	FirstName                       = "User’s first name."
	LastName                        = "User’s last name."
	OrgRole                         = "User's roles  for the associated organization."
	ProjectRole                     = "User's roles  for the associated project."
	TeamRole                        = "Project role you want to assign to the team."
	MaxDate                         = "Returns events whose created date is less than or equal to it."
	MinDate                         = "Returns events whose created date is greater than or equal to it."
	Filename                        = "Filename to use, optional file with a json cluster configuration."
	AccessListIps                   = "IP addresses to add to the new user’s access list."
	StartDate                       = "Timestamp in ISO 8601 date and time format in UTC when the maintenance window starts."
	EndDate                         = "Timestamp in ISO 8601 date and time format in UTC when the maintenance window ends."
	AlertType                       = "Alert types to silence during maintenance window. For example: HOST, REPLICA_SET, CLUSTER, AGENT, BACKUP."
	MaintenanceDescription          = "Description of the maintenance window."
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
	AlertConfigAPIKey               = "Datadog API Key, Opsgenie API Key, VictorOps API key."
	APIKey                          = "API Key."
	RoutingKey                      = "An optional field for your Routing Key."
	IntegrationAPIToken             = "Your API Token." //nolint:gosec // This is just a message not a credential
	OrgName                         = "Your Flowdock organization name."
	OrgNameFilter                   = "Performs a case-insensitive search for organizations which exactly match the specified name"
	OrgIncludeDeleted               = "Include deleted organization."
	FlowName                        = "Your Flowdock Flow name."
	BlockstoreAssignment            = "Flag indicating whether this blockstore can be assigned backup jobs."
	OplogAssignment                 = "Flag indicating whether this oplog can be assigned backup jobs."
	FileSystemAssignment            = "Flag indicating whether this file system store can be assigned backup jobs."
	EncryptedCredentials            = "Flag indicating whether the username and password were encrypted using the credentialstool."
	MMAPV1CompressionSetting        = "The compression setting for the MMAPv1 storage engine snaphots."
	WTCompressionSetting            = "The compression setting for the WiredTiger storage engine snaphots."
	StorePath                       = "The location where file system-based backups are stored on the file system store host."
	Label                           = "Array of tags to manage which backup jobs Ops Manager can assign to which blockstores."
	LoadFactor                      = "A positive, non-zero integer that expresses how much backup work this snapshot store should perform compared to another snapshot store."
	MaxCapacityGB                   = "The maximum amount of data in GB this blockstore can store."
	BlockstoreURI                   = "A comma-separated list of hosts in the <hostname:port> format that can be used to access this blockstore."
	BlockstoreSSL                   = "Flag indicating whether this blockstore only accepts connections encrypted using TLS."
	BlockstoreName                  = "The unique name that labels this blockstore."
	OplogName                       = "The unique name that labels this oplog store."
	FileSystemName                  = "The unique name that labels this file system store configuration."
	WriteConcern                    = "The write concern used for this blockstore."
	AWSAccessKey                    = "AWS Access Key ID that can access the S3 bucket specified in s3BucketName."
	AWSSecretKey                    = "AWS Secret Access Key that can access the S3 bucket specified in s3BucketName." //nolint:gosec // This is just a message not a credential
	DisableProxyS3                  = "Flag indicating whether the HTTP proxy should be used when connecting to S3."
	S3AuthMethod                    = "Method used to authorize access to the S3 bucket specified in s3BucketName. Accepted values for this option are: KEYS, IAM_ROLE."
	S3BucketEndpoint                = "URL that Ops Manager uses to access this AWS S3 or S3-compatible bucket."
	S3BucketName                    = "Name of the S3 bucket that hosts the S3 blockstore."
	S3MaxConnections                = "Positive integer indicating the maximum number of connections to this S3 blockstore."
	AcceptedTos                     = "Flag indicating whether or not you accepted the terms of service for using S3-compatible stores with Ops Manager."
	SSEEnabled                      = "Flag indicating whether this S3 blockstore enables server-side encryption."
	PathStyleAccessEnabled          = "Flag indicating the style of this endpoint."
	APIKeyDescription               = "Description of the API key."
	APIKeyRoles                     = "List of roles for the API key."
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
	SnapshotDescription             = "Description of the on-demand snapshot."
	Database                        = "Database name."
	DatabaseUser                    = "Username of a database user."
	MonthsUntilExpiration           = "Number of months that the certificate is valid for."
	Collection                      = "Collection name."
	Analyzer                        = "Analyzer to use when creating the index"
	SearchAnalyzer                  = "Analyzer to use when searching the index."
	Dynamic                         = "Indicates whether the index uses dynamic or static mappings."
	SearchFields                    = "Static field specifications."
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
	CASFilePath                     = "Path to a PEM file containing one or more CAs for database user authentication."
	Verbose                         = "If true, returns all child jobs in the response."
	ClusterID                       = "Unique identifier of the cluster."
	ReferenceTimeZoneOffset         = "The ISO-8601 timezone offset where the Ops Manager host resides."
	DailySnapshotRetentionDays      = "Number of days to retain daily snapshots. Ops Manager may return values between 1 and 365, inclusive."
	ClusterCheckpointIntervalMin    = "Number of minutes between successive cluster checkpoints. Ops Manager may return values of 15, 30, or 60."
	SnapshotIntervalHours           = "Number of hours between snapshots. Ops Manager may return values of 6, 8, 12, or 24."
	SnapshotRetentionDays           = "Number of days to keep recent snapshots. Ops Manager may return values between 2 and 5, inclusive."
	WeeklySnapshotRetentionWeeks    = "Number of weeks to retain weekly snapshots. Ops Manager may return values between 1 and 52, inclusive."
	PointInTimeWindowHours          = "Number of hours in the past for which a point-in-time snapshot can be created."
	ReferenceHourOfDay              = "Hour of the day to schedule snapshots using a 24 hour clock. Ops Manager may return values between 0 and 23, inclusive."
	ReferenceMinuteOfHour           = "Minute of the hour to schedule snapshots. Ops Manager may return values between 0 and 59, inclusive."
	MonthlySnapshotRetentionMonths  = "Number of months to retain monthly snapshots. Ops Manager may return values between 1 and 36, inclusive."
	Background                      = "Create the index in the background."
	DateField                       = "Name of an already indexed date field from the documents."
	PartitionFields                 = "Fields to use to partition data. You can specify up to two frequently queried fields to use for partitioning data."
	ArchiveAfter                    = "Number of days that specifies the age limit for the data in the live Atlas cluster."
	TargetProjectID                 = "Unique identifier of the project that contains the destination cluster for the restore job."
	AccessListIPEntry               = "IP address to be allowed for a given API key."
	AccessListCIDREntry             = "Whitelist entry in CIDR notation to be added for a given API key."
	PrivateEndpointID               = "Unique identifier of the AWS PrivateLink connection."
	AccountID                       = "Account ID of the owner of the peer VPC."
	NewRelicAccountID               = "Unique identifier of your New Relic account."
	LicenceKey                      = "Your License Key."
	ServiceKey                      = "Your Service Key."
	URL                             = "Your webhook URL."
	Secret                          = "An optional field for your webhook secret." //nolint:gosec // This is just a message not a credential
	WriteToken                      = "Your Insights Insert Key."
	DayOfWeek                       = "Day of the week that you want the maintenance window to start, as a 1-based integer."
	HourOfDay                       = "Hour of the day that you want the maintenance window to start. This parameter uses the 24-hour clock, where midnight is 0 and noon is 12."
	StartASAP                       = "Start maintenance immediately upon receiving this request."
	ReadToken                       = "Your Insights Query Key."
	RouteTableCidrBlock             = "Peer VPC CIDR block or subnet."
	VpcID                           = "Unique identifier of the peer VPC."
	AtlasCIDRBlock                  = "CIDR block that Atlas uses for your clusters."
	VNet                            = "Name of your Azure VNet."
	ResourceGroup                   = "Name of your Azure resource group."
	DirectoryID                     = "Unique identifier for an Azure AD directory."
	SubscriptionID                  = "Unique identifier of the Azure subscription in which the VNet resides."
	GCPProjectID                    = "Unique identifier of the GCP project in which the network peer resides."
	Network                         = "Unique identifier of the Network Peering connection in the Atlas project."
	APIRegion                       = "Indicates which API URL to use, either US or EU. The integration service will use US by default."
	FormatOut                       = `Output format. 
Valid values: json`
	TargetClusterID = `Unique identifier of the target cluster.
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
	AccessListType = `Type of access list entry.
Valid values: cidrBlock|ipAddress|awsSecurityGroup`
	Service = `Type of MongoDB service.
Valid values: cloud|cloud-manager|ops-manager`
	Provider = `Name of your cloud service provider.
Valid values: AWS|AZURE|GCP.`
	ClusterTypes = `Type of the cluster that you want to create.
Valid values: REPLICASET|SHARDED.`
	Region = `Physical location of your MongoDB cluster.
For a complete list of supported AWS regions, see: https://docs.atlas.mongodb.com/reference/amazon-aws/#amazon-aws
For a complete list of supported Azure regions, see: https://docs.atlas.mongodb.com/reference/microsoft-azure/#microsoft-azure
For a complete list of supported GCP regions, see: https://docs.atlas.mongodb.com/reference/google-gcp/#google-gcp`
	AWSIAMType = `AWS IAM method by which the provided username is authenticated. 
Valid values: NONE|USER|ROLE.`
	X509Type = `X.509 method by which the provided username is authenticated. 
Valid values: NONE|MANAGED|CUSTOMER.`
	LDAPType = `LDAP method by which the provided username is authenticated. 
Valid values: NONE|USER|GROUP.`
)
