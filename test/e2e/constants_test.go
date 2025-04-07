// Copyright 2025 MongoDB Inc
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package e2e

const (
	eventsEntity                  = "events"
	clustersEntity                = "clusters"
	processesEntity               = "processes"
	metricsEntity                 = "metrics"
	searchEntity                  = "search"
	indexEntity                   = "index"
	nodesEntity                   = "nodes"
	datafederationEntity          = "datafederation"
	datalakePipelineEntity        = "datalakepipeline"
	alertsEntity                  = "alerts"
	configEntity                  = "config"
	dbusersEntity                 = "dbusers"
	certsEntity                   = "certs"
	privateEndpointsEntity        = "privateendpoints"
	queryLimitsEntity             = "querylimits"
	onlineArchiveEntity           = "onlineArchives"
	projectEntity                 = "project"
	orgEntity                     = "org"
	invitationsEntity             = "invitations"
	maintenanceEntity             = "maintenanceWindows"
	integrationsEntity            = "integrations"
	securityEntity                = "security"
	ldapEntity                    = "ldap"
	awsEntity                     = "aws"
	azureEntity                   = "azure"
	gcpEntity                     = "gcp"
	customDNSEntity               = "customDns"
	logsEntity                    = "logs"
	cloudProvidersEntity          = "cloudProviders"
	accessRolesEntity             = "accessRoles"
	customDBRoleEntity            = "customDbRoles"
	regionalModeEntity            = "regionalModes"
	serverlessEntity              = "serverless"
	liveMigrationsEntity          = "liveMigrations"
	auditingEntity                = "auditing"
	accessLogsEntity              = "accessLogs"
	accessListEntity              = "accessList"
	performanceAdvisorEntity      = "performanceAdvisor"
	slowQueryLogsEntity           = "slowQueryLogs"
	namespacesEntity              = "namespaces"
	networkingEntity              = "networking"
	networkPeeringEntity          = "peering"
	suggestedIndexesEntity        = "suggestedIndexes"
	slowOperationThresholdEntity  = "slowOperationThreshold"
	tierM10                       = "M10"
	tierM20                       = "M20"
	tierM0                        = "M0"
	tierM2                        = "M2"
	diskSizeGB40                  = "40"
	diskSizeGB30                  = "30"
	projectsEntity                = "projects"
	settingsEntity                = "settings"
	backupsEntity                 = "backups"
	exportsEntity                 = "exports"
	bucketsEntity                 = "buckets"
	jobsEntity                    = "jobs"
	snapshotsEntity               = "snapshots"
	restoresEntity                = "restores"
	compliancePolicyEntity        = "compliancepolicy"
	policiesEntity                = "policies"
	teamsEntity                   = "teams"
	setupEntity                   = "setup"
	deploymentEntity              = "deployments"
	federatedAuthenticationEntity = "federatedAuthentication"
	federationSettingsEntity      = "federationSettings"
	identityProviderEntity        = "identityProvider"
	connectedOrgsConfigsEntity    = "connectedOrgConfigs"
	authEntity                    = "auth"
	streamsEntity                 = "streams"
	apiKeysEntity                 = "apikeys"
	apiKeyAccessListEntity        = "accessLists"
	usersEntity                   = "users"
)

// Auth constants.
const (
	whoami = "whoami"
)

// AlertConfig constants.
const (
	group         = "GROUP"
	eventTypeName = "NO_PRIMARY"
	intervalMin   = 5
	delayMin      = 0
)

// Integration constants.
const (
	datadogEntity   = "DATADOG"
	opsGenieEntity  = "OPS_GENIE"
	pagerDutyEntity = "PAGER_DUTY"
	victorOpsEntity = "VICTOR_OPS"
	webhookEntity   = "WEBHOOK"
)

// Cluster settings.
const (
	// 	e2eClusterTier       = "M10"
	// 	e2eGovClusterTier    = "M20"
	// 	e2eSharedClusterTier = "M2"
	e2eClusterProvider = "AWS" // e2eClusterProvider preferred provider for e2e testing.
)

// Backup compliance policy constants.
const (
	authorizedUserFirstName = "firstname"
	authorizedUserLastName  = "lastname"
	authorizedEmail         = "firstname.lastname@example.com"
)

// Local Development constants.
const (
	collectionName  = "movies"
	databaseName    = "sample_mflix"
	searchIndexName = "indexTest"
	vectorSearchDB  = "sample_mflix"
	vectorSearchCol = "embedded_movies"
)

// CLI Plugins System constants.
const (
	examplePluginRepository = "mongodb/atlas-cli-plugin-example"
	examplePluginName       = "atlas-cli-plugin-example"
)

// Roles constants.
const (
	roleName1   = "GROUP_READ_ONLY"
	roleName2   = "GROUP_DATA_ACCESS_READ_ONLY"
	roleNameOrg = "ORG_READ_ONLY"
)
