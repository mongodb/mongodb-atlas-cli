// Copyright 2021 MongoDB Inc
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

package flag

const (
	Service                         = "service"                         // Service flag to set service
	Profile                         = "profile"                         // Profile flag to use a profile
	ProfileShort                    = "P"                               // ProfileShort flag to use a profile
	OrgID                           = "orgId"                           // OrgID flag to use an Organization ID
	TeamID                          = "teamId"                          // TeamID flag
	URL                             = "url"                             // URL flag
	Secret                          = "secret"                          // Secret flag
	ProjectID                       = "projectId"                       // ProjectID flag to use a project ID
	ProcessName                     = "processName"                     // Process Name
	HostID                          = "hostId"                          // HostID flag
	Since                           = "since"                           // Since flag
	Duration                        = "duration"                        // Duration Flag
	NLog                            = "nLog"                            // NLog flag
	Namespaces                      = "namespaces"                      // Namespaces flag 	// Namespaces flag
	NIndexes                        = "nIndexes"                        // NIndexes flag
	NExamples                       = "nExamples"                       // NExamples flag
	AuthDB                          = "authDB"                          // AuthDB flag
	Hostname                        = "hostname"                        // Hostname flag
	Port                            = "port"                            // Port flag
	BindUsername                    = "bindUsername"                    // BindUsername flag
	BindPassword                    = "bindPassword"                    // BindPassword flag
	CaCertificate                   = "caCertificate"                   // CaCertificate flag
	AuthzQueryTemplate              = "authzQueryTemplate"              // AuthzQueryTemplate flag
	AuthenticationEnabled           = "authenticationEnabled"           // AuthenticationEnabled flag
	AuthorizationEnabled            = "authorizationEnabled"            // AuthorizationEnabled flag
	Provider                        = "provider"                        // Provider flag to set the cloud provider
	Region                          = "region"                          // Region flag
	RegionShort                     = "r"                               // RegionShort flag
	Members                         = "members"                         // Members flag
	Shards                          = "shards"                          // Shards flag
	MembersShort                    = "m"                               // MembersShort flag
	ShardsShort                     = "s"                               // ShardsShort flag
	Tier                            = "tier"                            // Tier flag
	Forever                         = "forever"                         // Forever flag
	ForeverShort                    = "F"                               // ForeverShort flag
	DiskSizeGB                      = "diskSizeGB"                      // DiskSizeGB flag
	MDBVersion                      = "mdbVersion"                      // MDBVersion flag
	Backup                          = "backup"                          // Backup flag
	ID                              = "id"                              // ID flag
	Username                        = "username"                        // Username flag
	StorageEngine                   = "storageEngine"                   // StorageEngine flag
	AuthMechanism                   = "authMechanism"                   // AuthMechanism flag
	Provisioned                     = "provisioned"                     // Provisioned flag
	Encryption                      = "encryption"                      // Encryption flag
	SSL                             = "ssl"                             // SSL flag
	SyncSource                      = "syncSource"                      // SyncSource flag
	ExcludedNamespace               = "excludedNamespace"               // ExcludedNamespace flag
	IncludedNamespace               = "includedNamespace"               // IncludedNamespace flag
	UsernameShort                   = "u"                               // UsernameShort flag
	Password                        = "password"                        // Password flag
	Country                         = "country"                         // Country flag
	X509Type                        = "x509Type"                        // X509Type flag
	AWSIAMType                      = "awsIAMType"                      // AWSIAMType flag
	LDAPType                        = "ldapType"                        // LDAPType flag
	Period                          = "period"                          // Period flag
	PasswordShort                   = "p"                               // PasswordShort flag
	Email                           = "email"                           // Email flag
	Mobile                          = "mobile"                          // Mobile flag
	OrgRole                         = "orgRole"                         // OrgRole flag
	ProjectRole                     = "projectRole"                     // ProjectRole flag
	Out                             = "out"                             // Out flag
	Output                          = "output"                          // Output flag
	OutputShort                     = "o"                               // OutputShort flag
	Minutes                         = "minutes"                         // Minutes flag
	Status                          = "status"                          // Status flag
	Start                           = "start"                           // Start flag
	End                             = "end"                             // End flag
	FirstName                       = "firstName"                       // FirstName flag
	LastName                        = "lastName"                        // LastName flag
	Role                            = "role"                            // Role flag
	RoleID                          = "roleId"                          // RoleID flag
	IAMAssumedRoleARN               = "iamAssumedRoleArn"               // IAMAssumedRoleARN flag
	Description                     = "desc"                            // Description flag
	StartDate                       = "startDate"                       // StartDate flag
	EndDate                         = "endDate"                         // EndDate flag
	Format                          = "format"                          // Format flag
	AlertType                       = "alertType"                       // AlertType flag
	Mechanisms                      = "mechanisms"                      // Mechanisms flag
	Type                            = "type"                            // Type flag
	Comment                         = "comment"                         // Comment flag
	DeleteAfter                     = "deleteAfter"                     // DeleteAfter flag
	ArchiveAfter                    = "archiveAfter"                    // ArchiveAfter flag
	DayOfWeek                       = "dayOfWeek"                       // DayOfWeek flag
	HourOfDay                       = "hourOfDay"                       // HourOfDay flag
	StartASAP                       = "startASAP"                       // StartASAP flag
	Until                           = "until"                           // Until flag
	Page                            = "page"                            // Page flag
	Limit                           = "limit"                           // Limit flag
	File                            = "file"                            // File flag
	FileShort                       = "f"                               // File flag
	Force                           = "force"                           // Force flag
	WhitelistIP                     = "whitelistIp"                     // WhitelistIP flag
	AccessListIP                    = "accessListIp"                    // AccessListIP flag
	Event                           = "event"                           // EventTypeName flag
	Enabled                         = "enabled"                         // Enabled flag
	MatcherFieldName                = "matcherFieldName"                // MatcherFieldName flag
	MatcherOperator                 = "matcherOperator"                 // MatcherOperator flag
	MatcherValue                    = "matcherValue"                    // MatcherValue flag
	MetricName                      = "metricName"                      // MetricName flag
	MetricOperator                  = "metricOperator"                  // MetricOperator flag
	MetricThreshold                 = "metricThreshold"                 // MetricThreshold flag
	MetricUnits                     = "metricUnits"                     // MetricUnits flag
	MetricMode                      = "metricMode"                      // MetricMode flag
	NotificationToken               = "notificationToken"               // NotificationToken flag
	NotificationChannelName         = "notificationChannelName"         // NotificationChannelName flag
	APIKey                          = "apiKey"                          // APIKey flag
	APIToken                        = "apiToken"                        // APIToken flag
	RoutingKey                      = "routingKey"                      // RoutingKey flag
	OrgName                         = "orgName"                         // OrgName flag
	FlowName                        = "flowName"                        // FlowName flag
	NotificationRegion              = "notificationRegion"              // NotificationRegion flag
	NotificationDelayMin            = "notificationDelayMin"            // NotificationDelayMin flag
	NotificationEmailAddress        = "notificationEmailAddress"        // NotificationEmailAddress flag
	NotificationEmailEnabled        = "notificationEmailEnabled"        // NotificationEmailEnabled flag
	NotificationFlowName            = "notificationFlowName"            // NotificationFlowName flag
	NotificationIntervalMin         = "notificationIntervalMin"         // NotificationIntervalMin flag
	NotificationMobileNumber        = "notificationMobileNumber"        // NotificationMobileNumber flag
	NotificationOrgName             = "notificationOrgName"             // NotificationsOrgName flag
	NotificationServiceKey          = "notificationServiceKey"          // NotificationsServiceKey flag
	NotificationSmsEnabled          = "notificationSmsEnabled"          // NotificationsSmsEnabled flag
	NotificationTeamID              = "notificationTeamId"              // NotificationTeamID flag
	NotificationType                = "notificationType"                // NotificationType flag
	NotificationUsername            = "notificationUsername"            // NotificationUsername flag
	NotificationVictorOpsRoutingKey = "notificationVictorOpsRoutingKey" // NotificationVictorOpsRoutingKey flag
	SnapshotID                      = "snapshotId"                      // SnapshotID flag
	IndexName                       = "indexName"                       // IndexName flag
	ClusterName                     = "clusterName"                     // ClusterName flag
	ClusterID                       = "clusterId"                       // ClusterID flag
	TargetProjectID                 = "targetProjectId"                 // TargetProjectID flag
	TargetClusterID                 = "targetClusterId"                 // TargetClusterID flag
	TargetClusterName               = "targetClusterName"               // TargetClusterName flag
	CheckpointID                    = "checkpointId"                    // CheckpointID flag
	OplogTS                         = "oplogTs"                         // OplogTS flag
	OplogInc                        = "oplogInc"                        // OplogInc flag
	PointInTimeUTCMillis            = "pointInTimeUTCMillis"            // PointInTimeUTCMillis flag
	Expires                         = "expires"                         // Expires flag
	MaxDownloads                    = "maxDownloads"                    // MaxDownloads flag
	ExpirationHours                 = "expirationHours"                 // ExpirationHours flag
	MaxDate                         = "maxDate"                         // MaxDate flag
	MinDate                         = "minDate"                         // MinDate flag
	Granularity                     = "granularity"                     // Granularity flag
	Key                             = "key"                             // Key flag
	Collection                      = "collection"                      // Collection flag
	CollectionName                  = "collectionName"                  // CollectionName flag
	Database                        = "db"                              // Database flag
	Unique                          = "unique"                          // Unique flag
	RSName                          = "rsName"                          // RSName flag
	Sparse                          = "sparse"                          // Sparse flag
	Background                      = "background"                      // Background flag
	Locale                          = "locale"                          // Locale flag
	CaseLevel                       = "caseLevel"                       // CaseLevel flag
	CaseFirst                       = "caseFirst"                       // CaseFirst flag
	Alternate                       = "alternate"                       // Alternate flag
	MaxVariable                     = "MaxVariable"                     // MaxVariable flag
	NumericOrdering                 = "numericOrdering"                 // NumericOrdering flag
	Normalization                   = "normalization"                   // Normalization flag
	Backwards                       = "backwards"                       // Backwards flag
	Strength                        = "strength"                        // Strength flag
	SizeRequestedPerFileBytes       = "sizeRequestedPerFileBytes"       // SizeRequestedPerFileBytes flag
	Redacted                        = "redacted"                        // Redacted flag
	Verbose                         = "verbose"                         // Verbose flag
	TestBucket                      = "testBucket"                      // TestBucket flag
	Partition                       = "partition"                       // Partition flag
	DateField                       = "dateField"                       // DateField flag
	DateFormat                      = "dateFormat"                      // DateFormat flag
	Analyzer                        = "analyzer"                        // Analyzer flag
	SearchAnalyzer                  = "searchAnalyzer"                  // SearchAnalyzer flag
	Dynamic                         = "dynamic"                         // Dynamic flag
	Field                           = "field"                           // Fields flag
	CASFilePath                     = "casFile"                         // CASFilePath flag
	MonthsUntilExpiration           = "monthsUntilExpiration"           // MonthsUntilExpiration flag
	IP                              = "ip"                              // IP flag
	CIDR                            = "cidr"                            // CIDR flag
	PrivateEndpointID               = "privateEndpointId"               // PrivateEndpointID flag
	Retention                       = "retention"                       // Retention flag
	AtlasCIDRBlock                  = "atlasCidrBlock"                  // AtlasCIDRBlock flag
	DirectoryID                     = "directoryId"                     // DirectoryID flag
	SubscriptionID                  = "subscriptionId"                  // SubscriptionID flag
	ResourceGroup                   = "resourceGroup"                   // ResourceGroup flag
	VNet                            = "vnet"                            // VNet flag
	AccountID                       = "accountId"                       // AccountID flag
	RouteTableCidrBlock             = "routeTableCidrBlock"             // RouteTableCidrBlock flag
	VpcID                           = "vpcId"                           // VpcID flag
	GCPProjectID                    = "gcpProjectId"                    // GCPProjectID flag
	Network                         = "network"                         // Network flag
	Name                            = "name"                            // Name flag
	LicenceKey                      = "licenceKey"                      // LicenceKey flag
	ServiceKey                      = "serviceKey"                      // ServiceKey flag
	WriteToken                      = "writeToken"                      // WriteToken flag
	ReadToken                       = "readToken"                       // ReadToken flag
	Assignment                      = "assignment"                      // Assignment flag
	EncryptedCredentials            = "encryptedCredentials"            // EncryptedCredentials flag
	Label                           = "label"                           // Label flag
	LoadFactor                      = "loadFactor"                      // LoadFactor flag
	MMAPV1CompressionSetting        = "mmapv1CompressionSetting"        // MMAPV1CompressionSetting flag
	WTCompressionSetting            = "wtCompressionSetting"            // WTCompressionSetting flag
	StorePath                       = "storePath"                       // StorePath flag
	MaxCapacityGB                   = "maxCapacityGB"                   // MaxCapacityGB flag
	URI                             = "uri"                             // URI flag
	WriteConcern                    = "writeConcern"                    // WriteConcern flag
	IncludeDeleted                  = "includeDeleted"                  // IncludeDeleted flag
	AWSAccessKey                    = "awsAccessKey"                    // AWSAccessKey flag
	AWSSecretKey                    = "awsSecretKey"                    // AWSSecretKey fag
	S3AuthMethod                    = "s3AuthMethod"                    // S3AuthMethod flag
	S3BucketEndpoint                = "s3BucketEndpoint"                // S3BucketEndpoint flag
	S3BucketName                    = "s3BucketName"                    // S3BucketName flag
	S3MaxConnections                = "s3MaxConnections"                // S3MaxConnections flag
	DisableProxyS3                  = "disableProxyS3"                  // DisableProxyS3 flag
	AcceptedTos                     = "acceptedTos"                     // AcceptedTos flag
	SSEEnabled                      = "sseEnabled"                      // SSEEnabled flag
	PathStyleAccessEnabled          = "pathStyleAccessEnabled"          // PathStyleAccessEnabled flag
	ReferenceTimeZoneOffset         = "referenceTimeZoneOffset"         // ReferenceTimeZoneOffset flag
	DailySnapshotRetentionDays      = "dailySnapshotRetentionDays"      // DailySnapshotRetentionDays flag
	ClusterCheckpointIntervalMin    = "clusterCheckpointIntervalMin"    // ClusterCheckpointIntervalMin flag
	SnapshotIntervalHours           = "snapshotIntervalHours"           // SnapshotIntervalHours flag
	SnapshotRetentionDays           = "snapshotRetentionDays"           // SnapshotRetentionDays flag
	WeeklySnapshotRetentionWeeks    = "weeklySnapshotRetentionWeeks"    // WeeklySnapshotRetentionWeeks flag
	PointInTimeWindowHours          = "pointInTimeWindowHours"          // PointInTimeWindowHours flag
	ReferenceHourOfDay              = "referenceHourOfDay"              // ReferenceHourOfDay flag
	ReferenceMinuteOfHour           = "referenceMinuteOfHour"           // ReferenceMinuteOfHour flag
	MonthlySnapshotRetentionMonths  = "monthlySnapshotRetentionMonths"  // MonthlySnapshotRetentionMonths flag
	Policy                          = "policy"                          // Policy flag
	SystemID                        = "systemId"                        // SystemID flag

)
