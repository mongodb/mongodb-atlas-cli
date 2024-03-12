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
	ProjectID                       = "projectId"                       // ProjectID flag to use a project ID
	ProcessName                     = "processName"                     // Process Name
	HostID                          = "hostId"                          // HostID flag
	Since                           = "since"                           // Since flag
	Duration                        = "duration"                        // Duration flag
	NLog                            = "nLog"                            // NLog flag
	Namespaces                      = "namespaces"                      // Namespaces flag
	NIndexes                        = "nIndexes"                        // NIndexes flag
	NExamples                       = "nExamples"                       // NExamples flag
	AuthDB                          = "authDB"                          // AuthDB flag
	Forever                         = "forever"                         // Forever flag
	ForeverShort                    = "F"                               // ForeverShort flag
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
	Description                     = "desc"                            // Description flag
	StartDate                       = "startDate"                       // StartDate flag
	EndDate                         = "endDate"                         // EndDate flag
	Format                          = "format"                          // Format flag
	AlertType                       = "alertType"                       // AlertType flag
	Mechanisms                      = "mechanisms"                      // Mechanisms flag
	TypeFlag                        = "type"                            // TypeFlag flag
	Comment                         = "comment"                         // Comment flag
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
	NotificationRegion              = "notificationRegion"              // NotificationRegion flag
	NotificationDelayMin            = "notificationDelayMin"            // NotificationDelayMin flag
	NotificationEmailAddress        = "notificationEmailAddress"        // NotificationEmailAddress flag
	NotificationEmailEnabled        = "notificationEmailEnabled"        // NotificationEmailEnabled flag
	NotificationIntervalMin         = "notificationIntervalMin"         // NotificationIntervalMin flag
	NotificationMobileNumber        = "notificationMobileNumber"        // NotificationMobileNumber flag
	NotificationServiceKey          = "notificationServiceKey"          // NotificationsServiceKey flag
	NotificationSmsEnabled          = "notificationSmsEnabled"          // NotificationsSmsEnabled flag
	NotificationTeamID              = "notificationTeamId"              // NotificationTeamID flag
	NotificationType                = "notificationType"                // NotificationType flag
	NotificationUsername            = "notificationUsername"            // NotificationUsername flag
	NotificationVictorOpsRoutingKey = "notificationVictorOpsRoutingKey" // NotificationVictorOpsRoutingKey flag
	SnapshotID                      = "snapshotId"                      // SnapshotID flag
	ClusterID                       = "clusterId"                       // ClusterID flag
	TargetProjectID                 = "targetProjectId"                 // TargetProjectID flag
	TargetClusterID                 = "targetClusterId"                 // TargetClusterID flag
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
	CollectionName                  = "collectionName"                  // CollectionName flag
	Database                        = "db"                              // Database flag
	RSName                          = "rsName"                          // RSName flag
	Sparse                          = "sparse"                          // Sparse flag
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
	SkipRedaction                   = "skipRedaction"                   // SkipRedaction flag
	Verbose                         = "verbose"                         // Verbose flag
	IP                              = "ip"                              // IP flag
	CIDR                            = "cidr"                            // CIDR flag
	Name                            = "name"                            // Name flag
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
	Timestamp                       = "timestamp"                       // Timestamp flag
	OwnerID                         = "ownerId"                         // OwnerID flag
	LinkToken                       = "linkToken"                       // LinkToken flag
	WithoutDefaultAlertSettings     = "withoutDefaultAlertSettings"     // WithoutDefaultAlertSettings flag
	Version                         = "version"                         // Version flag
	LocalKeyFile                    = "localKeyFile"                    // LocalKeyFile flag
	KMIPServerCAFile                = "kmipServerCAFile"                // KMIPServerCAFile flag
	KMIPClientCertificateFile       = "kmipClientCertificateFile"       // KMIPClientCertificateFile flag
	Debug                           = "debug"                           // Debug flag to set debug log level
	DebugShort                      = "D"                               // DebugShort flag to set debug log level
	KMIPClientCertificatePassword   = "kmipClientCertificatePassword"   //nolint:gosec // KMIPClientCertificatePassword flag
	KMIPUsername                    = "kmipUsername"                    // KMIPUsername flag
	KMIPPassword                    = "kmipPassword"                    // KMIPPassword flag
)
