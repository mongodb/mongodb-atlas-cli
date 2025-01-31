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
	Profile                                   = "profile"                                   // Profile flag to use a profile
	ProfileShort                              = "P"                                         // ProfileShort flag to use a profile
	OrgID                                     = "orgId"                                     // OrgID flag to use an Organization ID
	TeamID                                    = "teamId"                                    // TeamID flag
	URL                                       = "url"                                       // URL flag
	Secret                                    = "secret"                                    // Secret flag
	ProjectID                                 = "projectId"                                 // ProjectID flag to use a project ID
	ProcessName                               = "processName"                               // Process Name
	HostID                                    = "hostId"                                    // HostID flag
	Since                                     = "since"                                     // Since flag
	Duration                                  = "duration"                                  // Duration flag
	NLog                                      = "nLog"                                      // NLog flag
	Namespaces                                = "namespaces"                                // Namespaces flag
	NIndexes                                  = "nIndexes"                                  // NIndexes flag
	NExamples                                 = "nExamples"                                 // NExamples flag
	AuthDB                                    = "authDB"                                    // AuthDB flag
	Hostname                                  = "hostname"                                  // Hostname flag
	Port                                      = "port"                                      // Port flag
	BindUsername                              = "bindUsername"                              // BindUsername flag
	BindPassword                              = "bindPassword"                              // BindPassword flag
	CaCertificate                             = "caCertificate"                             // CaCertificate flag
	AuthzQueryTemplate                        = "authzQueryTemplate"                        // AuthzQueryTemplate flag
	MappingMatch                              = "mappingMatch"                              // MappingMatch flag
	MappingLdapQuery                          = "mappingLdapQuery"                          // MappingLdapQuery flag
	MappingSubstitution                       = "mappingSubstitution"                       // MappingSubstitution flag
	AuthenticationEnabled                     = "authenticationEnabled"                     // AuthenticationEnabled flag
	AuthorizationEnabled                      = "authorizationEnabled"                      // AuthorizationEnabled flag
	Provider                                  = "provider"                                  // Provider flag
	Region                                    = "region"                                    // Region flag
	RegionShort                               = "r"                                         // RegionShort flag
	Members                                   = "members"                                   // Members flag
	Shards                                    = "shards"                                    // Shards flag
	MembersShort                              = "m"                                         // MembersShort flag
	ShardsShort                               = "s"                                         // ShardsShort flag
	Tier                                      = "tier"                                      // Tier flag
	Forever                                   = "forever"                                   // Forever flag
	ForeverShort                              = "F"                                         // ForeverShort flag
	DiskSizeGB                                = "diskSizeGB"                                // DiskSizeGB flag
	MDBVersion                                = "mdbVersion"                                // MDBVersion flag
	Backup                                    = "backup"                                    // Backup flag
	BIConnector                               = "biConnector"                               // BIConnector flag
	ID                                        = "id"                                        // ID flag
	EnableTerminationProtection               = "enableTerminationProtection"               // EnableTerminationProtection flag
	DisableTerminationProtection              = "disableTerminationProtection"              // DisableTerminationProtection flag
	Username                                  = "username"                                  // Username flag
	UsernameShort                             = "u"                                         // UsernameShort flag
	Password                                  = "password"                                  // Password flag
	SkipMongosh                               = "skipMongosh"                               // SkipMongosh flag
	SkipSampleData                            = "skipSampleData"                            // SkipSampleData flag
	Country                                   = "country"                                   // Country flag
	X509Type                                  = "x509Type"                                  // X509Type flag
	AWSIAMType                                = "awsIAMType"                                // AWSIAMType flag
	LDAPType                                  = "ldapType"                                  // LDAPType flag
	Period                                    = "period"                                    // Period flag
	PasswordShort                             = "p"                                         // PasswordShort flag
	Email                                     = "email"                                     // Email flag
	Mobile                                    = "mobile"                                    // Mobile flag
	OrgRole                                   = "orgRole"                                   // OrgRole flag
	ProjectRole                               = "projectRole"                               // ProjectRole flag
	Out                                       = "out"                                       // Out flag
	Output                                    = "output"                                    // Output flag
	OutputShort                               = "o"                                         // OutputShort flag
	Status                                    = "status"                                    // Status flag
	Start                                     = "start"                                     // Start flag
	End                                       = "end"                                       // End flag
	AuthResult                                = "authResult"                                // AuthResult flag
	FirstName                                 = "firstName"                                 // FirstName flag
	LastName                                  = "lastName"                                  // LastName flag
	Role                                      = "role"                                      // Role flag
	Scope                                     = "scope"                                     // Scope flag
	IAMAssumedRoleARN                         = "iamAssumedRoleArn"                         // IAMAssumedRoleARN flag
	Description                               = "desc"                                      // Description flag
	TypeFlag                                  = "type"                                      // TypeFlag flag
	Comment                                   = "comment"                                   // Comment flag
	DeleteAfter                               = "deleteAfter"                               // DeleteAfter flag
	ArchiveAfter                              = "archiveAfter"                              // ArchiveAfter flag
	ExpireAfterDays                           = "expireAfterDays"                           // ExpireAfterDays flag
	DayOfWeek                                 = "dayOfWeek"                                 // DayOfWeek flag
	HourOfDay                                 = "hourOfDay"                                 // HourOfDay flag
	StartASAP                                 = "startASAP"                                 // StartASAP flag
	Until                                     = "until"                                     // Until flag
	Page                                      = "page"                                      // Page flag
	Limit                                     = "limit"                                     // Limit flag
	OmitCount                                 = "omitCount"                                 // OmitCount flag
	File                                      = "file"                                      // File flag
	FileShort                                 = "f"                                         // File flag
	Force                                     = "force"                                     // Force flag
	AccessListIP                              = "accessListIp"                              // AccessListIP flag
	Event                                     = "event"                                     // EventTypeName flag
	Enabled                                   = "enabled"                                   // Enabled flag
	MatcherFieldName                          = "matcherFieldName"                          // MatcherFieldName flag
	MatcherOperator                           = "matcherOperator"                           // MatcherOperator flag
	MatcherValue                              = "matcherValue"                              // MatcherValue flag
	MetricName                                = "metricName"                                // MetricName flag
	MetricOperator                            = "metricOperator"                            // MetricOperator flag
	MetricThreshold                           = "metricThreshold"                           // MetricThreshold flag
	MetricUnits                               = "metricUnits"                               // MetricUnits flag
	MetricMode                                = "metricMode"                                // MetricMode flag
	NotificationToken                         = "notificationToken"                         // NotificationToken flag
	NotificationChannelName                   = "notificationChannelName"                   // NotificationChannelName flag
	APIKey                                    = "apiKey"                                    // APIKey flag
	APIKeyRole                                = "apiKeyRole"                                // APIKeyRole flag
	RoutingKey                                = "routingKey"                                // RoutingKey flag
	NotificationRegion                        = "notificationRegion"                        // NotificationRegion flag
	NotificationDelayMin                      = "notificationDelayMin"                      // NotificationDelayMin flag
	NotificationEmailAddress                  = "notificationEmailAddress"                  // NotificationEmailAddress flag
	NotificationEmailEnabled                  = "notificationEmailEnabled"                  // NotificationEmailEnabled flag
	NotificationIntervalMin                   = "notificationIntervalMin"                   // NotificationIntervalMin flag
	NotificationMobileNumber                  = "notificationMobileNumber"                  // NotificationMobileNumber flag
	NotificationServiceKey                    = "notificationServiceKey"                    // NotificationsServiceKey flag
	NotificationSmsEnabled                    = "notificationSmsEnabled"                    // NotificationsSmsEnabled flag
	NotificationTeamID                        = "notificationTeamId"                        // NotificationTeamID flag
	NotificationType                          = "notificationType"                          // NotificationType flag
	NotificationUsername                      = "notificationUsername"                      // NotificationUsername flag
	NotificationVictorOpsRoutingKey           = "notificationVictorOpsRoutingKey"           // NotificationVictorOpsRoutingKey flag
	NotificationWebhookURL                    = "notificationWebhookUrl"                    // NotificationWebhookURL flag
	NotificationWebhookSecret                 = "notificationWebhookSecret"                 // NotificationWebhookSecret flag
	NotifierID                                = "notifierId"                                // NotifierID flag
	NotificationRole                          = "notificationRole"                          // NotificationRoles flag
	SnapshotID                                = "snapshotId"                                // SnapshotID flag
	IndexName                                 = "indexName"                                 // IndexName flag
	ClusterName                               = "clusterName"                               // ClusterName flag
	TargetProjectID                           = "targetProjectId"                           // TargetProjectID flag
	TargetClusterName                         = "targetClusterName"                         // TargetClusterName flag
	OplogTS                                   = "oplogTs"                                   // OplogTS flag
	OplogInc                                  = "oplogInc"                                  // OplogInc flag
	PointInTimeUTCMillis                      = "pointInTimeUTCMillis"                      // PointInTimeUTCMillis flag
	PointInTimeUTCSeconds                     = "pointInTimeUTCSeconds"                     // PointInTimeUTCSeconds flag
	MaxDate                                   = "maxDate"                                   // MaxDate flag
	MinDate                                   = "minDate"                                   // MinDate flag
	Granularity                               = "granularity"                               // Granularity flag
	Key                                       = "key"                                       // Key flag
	Collection                                = "collection"                                // Collection flag
	Append                                    = "append"                                    // Append flag
	Privilege                                 = "privilege"                                 // Privilege flag
	InheritedRole                             = "inheritedRole"                             // InheritedRole flag
	Database                                  = "db"                                        // Database flag
	Sparse                                    = "sparse"                                    // Sparse flag
	TestBucket                                = "testBucket"                                // TestBucket flag
	Partition                                 = "partition"                                 // Partition flag
	DateField                                 = "dateField"                                 // DateField flag
	DateFormat                                = "dateFormat"                                // DateFormat flag
	Analyzer                                  = "analyzer"                                  // Analyzer flag
	SearchAnalyzer                            = "searchAnalyzer"                            // SearchAnalyzer flag
	Dynamic                                   = "dynamic"                                   // Dynamic flag
	Field                                     = "field"                                     // Fields flag
	CASFilePath                               = "casFile"                                   // CASFilePath flag
	MonthsUntilExpiration                     = "monthsUntilExpiration"                     // MonthsUntilExpiration flag
	IP                                        = "ip"                                        // IP flag
	CIDR                                      = "cidr"                                      // CIDR flag
	PrivateEndpointID                         = "privateEndpointId"                         // PrivateEndpointID flag
	EndpointServiceID                         = "endpointServiceId"                         // EndpointServiceId flag
	PrivateEndpointIPAddress                  = "privateEndpointIpAddress"                  // PrivateEndpointIPAddress flag
	Endpoint                                  = "endpoint"                                  // Endpoint flag
	Retention                                 = "retention"                                 // Retention flag
	AtlasCIDRBlock                            = "atlasCidrBlock"                            // AtlasCIDRBlock flag
	DirectoryID                               = "directoryId"                               // DirectoryID flag
	SubscriptionID                            = "subscriptionId"                            // SubscriptionID flag
	ResourceGroup                             = "resourceGroup"                             // ResourceGroup flag
	VNet                                      = "vnet"                                      // VNet flag
	AccountID                                 = "accountId"                                 // AccountID flag
	RouteTableCidrBlock                       = "routeTableCidrBlock"                       // RouteTableCidrBlock flag
	VpcID                                     = "vpcId"                                     // VpcID flag
	GCPProjectID                              = "gcpProjectId"                              // GCPProjectID flag
	Network                                   = "network"                                   // Network flag
	Name                                      = "name"                                      // Name flag
	LicenceKey                                = "licenceKey"                                // LicenceKey flag
	ServiceKey                                = "serviceKey"                                // ServiceKey flag
	WriteToken                                = "writeToken"                                // WriteToken flag
	ReadToken                                 = "readToken"                                 // ReadToken flag
	WriteConcern                              = "writeConcern"                              // WriteConcern flag
	ReadConcern                               = "readConcern"                               // ReadConcern flag
	DisableFailIndexKeyTooLong                = "disableFailIndexKeyTooLong"                // DisableFailIndexKeyTooLong flag
	EnableFailIndexKeyTooLong                 = "enableFailIndexKeyTooLong"                 // EnableFailIndexKeyTooLong flag
	DisableJavascript                         = "disableJavascript"                         // DisableJavascript flag
	EnableJavascript                          = "enableJavascript"                          // EnableJavascript flag
	TLSProtocol                               = "tlsProtocol"                               // TLSProtocol flag
	DisableTableScan                          = "disableTableScan"                          // DisableTableScan flag
	EnableTableScan                           = "enableTableScan"                           // EnableTableScan flag
	OplogMinRetentionHours                    = "oplogMinRetentionHours"                    // OplogMinRetentionHours flag
	OplogSizeMB                               = "oplogSizeMB"                               // OplogSizeMB flag
	SampleRefreshIntervalBIConnector          = "sampleRefreshIntervalBIConnector"          // SampleRefreshIntervalBIConnector flag
	SampleSizeBIConnector                     = "sampleSizeBIConnector"                     // SampleSizeBIConnector flag
	IncludeDeleted                            = "includeDeleted"                            // IncludeDeleted flag
	AWSAccessKey                              = "awsAccessKey"                              // AWSAccessKey flag
	AWSSecretKey                              = "awsSecretKey"                              // AWSSecretKey fag
	ReferenceHourOfDay                        = "referenceHourOfDay"                        // ReferenceHourOfDay flag
	ReferenceMinuteOfHour                     = "referenceMinuteOfHour"                     // ReferenceMinuteOfHour flag
	Policy                                    = "policy"                                    // Policy flag
	Default                                   = "default"                                   // Default flag
	LiveMigrationHost                         = "migrationHost"                             // LiveMigrationHost flag
	LiveMigrationSourceClusterName            = "sourceClusterName"                         // LiveMigrationSourceClusterName flag
	LiveMigrationSourceProjectID              = "sourceProjectId"                           // LiveMigrationSourceProjectID flag
	LiveMigrationSourceSSL                    = "sourceSsl"                                 // LiveMigrationSourceSSL flag
	LiveMigrationSourceCACertificatePath      = "sourceCACertificatePath"                   // LiveMigrationSourceCACertificatePath flag
	LiveMigrationSourceManagedAuthentication  = "sourceManagedAuthentication"               // LiveMigrationSourceManagedAuthentication flag
	LiveMigrationSourceUsername               = "sourceUsername"                            // LiveMigrationSourceUsername flag
	LiveMigrationSourcePassword               = "sourcePassword"                            // LiveMigrationSourcePassword flag
	LiveMigrationDropCollections              = "drop"                                      // LiveMigrationDropCollections flag
	OwnerID                                   = "ownerId"                                   // OwnerID flag
	GovCloudRegionsOnly                       = "govCloudRegionsOnly"                       // GovCloudRegionsOnly flag
	LiveMigrationID                           = "liveMigrationId"                           // LiveMigrationID flag
	LiveMigrationValidationID                 = "validationId"                              // LiveMigrationDropCollections flag
	WithoutDefaultAlertSettings               = "withoutDefaultAlertSettings"               // WithoutDefaultAlertSettings flag
	CurrentIP                                 = "currentIp"                                 // CurrentIP flag
	Gov                                       = "gov"                                       // Gov flag
	Version                                   = "version"                                   // Version flag
	EnableCollectDatabaseSpecificsStatistics  = "enableCollectDatabaseSpecificsStatistics"  // EnableCollectDatabaseSpecificsStatistics flag
	DisableCollectDatabaseSpecificsStatistics = "disableCollectDatabaseSpecificsStatistics" // DisableCollectDatabaseSpecificsStatistics flag
	EnableDataExplorer                        = "enableDataExplorer"                        // EnableDataExplorer flag
	DisableDataExplorer                       = "disableDataExplorer"                       // DisableDataExplorer flag
	EnablePerformanceAdvisor                  = "enablePerformanceAdvisor"                  // EnablePerformanceAdvisor flag
	DisablePerformanceAdvisor                 = "disablePerformanceAdvisor"                 // DisablePerformanceAdvisor flag
	EnableSchemaAdvisor                       = "enableSchemaAdvisor"                       // EnableSchemaAdvisor flag
	DisableSchemaAdvisor                      = "disableSchemaAdvisor"                      // DisableSchemaAdvisor flag
	EnableRealtimePerformancePanel            = "enableRealtimePerformancePanel"            // EnableRealtimePerformancePanel flag
	DisableRealtimePerformancePanel           = "disableRealtimePerformancePanel"           // DisableRealtimePerformancePanel flag
	CloudProvider                             = "cloudProvider"                             // CloudProvider flag
	IAMRoleID                                 = "iamRoleId"                                 // IamRoleID flag
	BucketID                                  = "bucketId"                                  // BucketID flag
	CustomData                                = "customData"                                // CustomData flag
	ExportBucketID                            = "exportBucketId"                            // ExportBucketID flag
	ExportFrequencyType                       = "exportFrequencyType"                       // ExportFrequencyType flag
	RestoreWindowDays                         = "restoreWindowDays"                         // RestoreWindowDays flag
	AutoExport                                = "autoExport"                                // AutoExport flag
	NoAutoExport                              = "noAutoExport"                              // NoAutoExport flag
	UpdateSnapshots                           = "updateSnapshots"                           // UpdateSnapshots flag
	NoUpdateSnapshots                         = "noUpdateSnapshots"                         // NoUpdateSnapshots flag
	UseOrgAndGroupNamesInExportPrefix         = "useOrgAndGroupNamesInExportPrefix"         // UseOrgAndGroupNamesInExportPrefix flag
	NoUseOrgAndGroupNamesInExportPrefix       = "noUseOrgAndGroupNamesInExportPrefix"       // NoUseOrgAndGroupNamesInExportPrefix flag
	BackupPolicy                              = "policy"                                    // BackupPolicy flag
	OperatorIncludeSecrets                    = "includeSecrets"                            // OperatorIncludeSecrets flag
	OperatorTargetNamespace                   = "targetNamespace"                           // OperatorTargetNamespace flag
	OperatorWatchNamespaces                   = "watchNamespaces"                           // OperatorTargetNamespace flag
	OperatorVersion                           = "operatorVersion"                           // OperatorVersion flag
	OperatorProjectName                       = "projectName"                               // OperatorProjectName flag
	OperatorImport                            = "import"                                    // OperatorImport flag
	OperatorResourceDeletionProtection        = "resourceDeletionProtection"                // OperatorResourceDeletionProtection flag
	OperatorSubResourceDeletionProtection     = "subresourceDeletionProtection"             // Operator OperatorSubResourceDeletionProtection flag
	OperatorConfigOnly                        = "configOnly"                                // OperatorConfigOnly config flag
	OperatorAtlasGov                          = "atlasGov"                                  // OperatorAtlasGov flag
	KubernetesClusterConfig                   = "kubeconfig"                                // Kubeconfig flag
	KubernetesClusterContext                  = "kubeContext"                               // KubeContext flag
	ExportID                                  = "exportId"                                  // ExportID flag
	Debug                                     = "debug"                                     // Debug flag to set debug log level
	DebugShort                                = "D"                                         // DebugShort flag to set debug log level
	GCPServiceAccountKey                      = "gcpServiceAccountKey"                      // GCPServiceAccountKey flag
	AzureClientID                             = "azureClientId"                             // AzureClientID flag
	AzureTenantID                             = "azureTenantId"                             // AzureTenantID flag
	AzureSecret                               = "azureSecret"                               // AzureSecret flag
	AWSSessionToken                           = "awsSessionToken"                           // AWSSessionToken flag
	APIKeyDescription                         = "apiKeyDescription"                         //nolint:gosec // APIKeyDescription flag
	RestoreJobID                              = "restoreJobId"                              // ID of the Restore Job
	DeliveryType                              = "deliveryType"                              // Type of the backup restore delivery
	EnableServerlessContinuousBackup          = "enableServerlessContinuousBackup"          // EnableServerlessContinuousBackup flag
	DisableServerlessContinuousBackup         = "disableServerlessContinuousBackup"         // DisableServerlessContinuousBackup flag
	SinkType                                  = "sinkType"                                  // SinkType flag
	SinkMetadataProvider                      = "sinkMetadataProvider"                      // SinkMetadataProvider flag
	SinkMetadataRegion                        = "sinkMetadataRegion"                        // SinkMetadataRegion flag
	SinkPartitionField                        = "sinkPartitionField"                        // SinkPartitionField flag
	SourceType                                = "sourceType"                                // SourceType flag
	SourceClusterName                         = "sourceClusterName"                         // SourceClusterName flag
	SourceCollectionName                      = "sourceCollectionName"                      // SourceCollectionName flag
	SourceDatabaseName                        = "sourceDatabaseName"                        // SourceDatabaseName flag
	SourcePolicyItemID                        = "sourcePolicyItemId"                        // SourcePolicyItemID flag
	Transform                                 = "transform"                                 // Transform flag
	Pipeline                                  = "pipeline"                                  // Pipeline flag
	CompletedAfter                            = "completedAfter"                            // CompletedAfter flag
	Tag                                       = "tag"                                       // Tag flag
	EnableWatch                               = "watch"                                     // EnableWatch flag
	EnableWatchShort                          = "w"                                         // EnableWatchShort flag
	WatchTimeout                              = "watchTimeout"                              // WatchTimeout flag
	CompactResponse                           = "compact"                                   // CompactResponse flag to return compacted list response
	CompactResponseShort                      = "c"                                         // CompactResponseShort flag
	AWSRoleID                                 = "awsRoleId"                                 // AWSRoleID flag
	AWSTestS3Bucket                           = "awsTestS3Bucket"                           // AWSTestS3Bucket flag
	DataFederation                            = "dataFederation"                            // DataFederation flag
	DataFederationName                        = "dataFederationName"                        // DataFederationName flag
	Value                                     = "value"                                     // Value flag
	OverrunPolicy                             = "overrunPolicy"                             // OverrunPolicy flag
	AuthorizedUserFirstName                   = "authorizedUserFirstName"                   // authorizedUserFirstName flag
	AuthorizedUserLastName                    = "authorizedUserLastName"                    // authorizedUserLastName flag
	AuthorizedEmail                           = "authorizedEmail"                           // authorizedEmail flag
	ConnectWith                               = "connectWith"                               // connectWith flag
	DeploymentName                            = "deploymentName"                            // deploymentName flag
	Instance                                  = "instance"                                  // Instance flag
	InstanceShort                             = "i"                                         // InstanceShort flag
	ConnectionStringType                      = "connectionStringType"                      // ConnectionStringType flag
	Decompress                                = "decompress"                                // Decompress flag
	DecompressShort                           = "d"                                         // DecompressShort flag
	BindIPAll                                 = "bindIpAll"                                 // BindIpAll flag
	FrequencyType                             = "frequencyType"                             // FrequencyType flag
	FrequencyInterval                         = "frequencyInterval"                         // FrequencyInterval flag
	All                                       = "all"                                       // All flag
	RetentionUnit                             = "retentionUnit"                             // RetentionUnit flag
	RetentionValue                            = "retentionValue"                            // RetentionValue flag
	FederationSettingsID                      = "federationSettingsId"                      // FederationSettingsId flag
	OIDCType                                  = "oidcType"                                  // OidcType flag
	IdpType                                   = "idpType"                                   // IdpType flag
	Audience                                  = "audience"                                  // Audience flag
	AuthorizationType                         = "authorizationType"                         // AuthorizationType flag
	ClientID                                  = "clientId"                                  // ClientId flag
	DisplayName                               = "displayName"                               // DisplayName flag
	GroupsClaim                               = "groupsClaim"                               // GroupsClaim flag
	UserClaim                                 = "userClaim"                                 // UserClaim flag
	IssuerURI                                 = "issuerUri"                                 // IssuerURI flag
	AssociatedDomain                          = "associatedDomain"                          // AssociatedDomain flag
	RequestedScope                            = "requestedScope"                            // RequestedScope flag
	Protocol                                  = "protocol"                                  // Protocol flag
	IdentityProviderID                        = "identityProviderId"                        // IdentityProviderId flag
	AuditAuthorizationSuccess                 = "auditAuthorizationSuccess"                 // AuditAuthorizationSuccess flag
	AuditFilter                               = "auditFilter"                               // AuditFilter flag
	IndependentResources                      = "independentResources"                      // IndependentResources flag
	InitDB                                    = "initdb"                                    // InitDB flag
)
