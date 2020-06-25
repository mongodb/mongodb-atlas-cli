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
package description

const (
	Atlas                    = "Atlas operations."
	Alerts                   = "Manage alerts for your project."
	AcknowledgeAlerts        = "Acknowledge an alert for your project."
	UnacknowledgeAlerts      = "Unacknowledge an alert for your project."
	Config                   = "Manage alerts configuration for your project."
	ConfigLong               = "The configs command provides access to your alerts configurations. You can create, edit, and delete alert configurations."
	CreateAlertsConfig       = "Create an alert configuration for your project."
	DeleteAlertsConfig       = "Delete an alert configuration from your project."
	AlertsConfigFields       = "Manage alert configuration fields for your project."
	AlertsConfigFieldsType   = "List alert configurations available field types."
	ListAlertsConfigs        = "List alert configurations for a project."
	UpdateAlertsConfig       = "Update an alert configuration for a project."
	DescribeAlert            = "Describe an alert for a project."
	ListAlerts               = "List alerts for a project."
	Backup                   = "Manage backups for your project."
	Checkpoints              = "Manage backup checkpoints for your project."
	ListCheckpoints          = "List continuous backup checkpoints for your project."
	Restores                 = "Manage restore jobs for your project."
	ListRestores             = "Lists restore jobs for a project and cluster."
	StartRestore             = "Start a restore job for a project and cluster."
	Snapshots                = "Manage backup snapshots for your project."
	ListSnapshots            = "List snapshots for a project and cluster."
	Logs                     = "Download host logs for your project."
	LogsDownload             = "Download a host mongodb logs."
	Clusters                 = "Manage clusters for your project."
	ClustersIndexes          = "Manage cluster rolling indexes for your project."
	ClustersLong             = "The clusters command provides access to your cluster configurations. You can create, edit, and delete clusters."
	CreateCluster            = "Create a MongoDB cluster."
	CreateIndex              = "Create a rolling index for your MongoDB cluster."
	ApplyCluster             = "Apply a new cluster configuration for your project."
	DeleteCluster            = "Delete a cluster from your project."
	DescribeCluster          = "Describe a cluster."
	ListClusters             = "List clusters for your project."
	UpdateCluster            = "Update a MongoDB cluster."
	PauseCluster             = "Pause a running MongoDB cluster in Atlas."
	StartCluster             = "Start a paused MongoDB cluster in Atlas."
	WatchCluster             = "Watch for a cluster to be available."
	Processes                = "Manage MongoDB processes for your project."
	ListProcesses            = "List MongoDB processes for a project."
	DBUsers                  = "Manage database users for your project."
	DiagnoseArchive          = "Manage diagnose archives."
	DownloadDiagnoseArchive  = "Download diagnose archives."
	CreateDBUser             = "Create a database user for a project."
	DeleteDBUser             = "Delete a database user for a project."
	ListDBUsers              = "List Atlas database users for a project."
	ListEvents               = "List events for an organization or project"
	UpdateDBUser             = "Update a MongoDB dbuser in Atlas."
	ProcessMeasurements      = "Get measurements for a given host."
	Disks                    = "List available disks or disks measurements for a given host."
	ListDisks                = "List available disks for a given host."
	DescribeDisks            = "Describe disks measurements for a given host partition."
	Databases                = "List available databases or databases measurements for a given host."
	ListDatabases            = "List available databases for a given host."
	Whitelist                = "Manage the IP whitelist for a project."
	CreateWhitelist          = "Create an IP whitelist for a project."
	DeleteWhitelist          = "Delete a database user for a project."
	DescribeWhitelist        = "Describe an Atlas whitelist."
	ListWhitelist            = "List Atlas whitelist for a project."
	CloudManager             = "Cloud Manager operations."
	ShutdownCluster          = "Shutdown a cluster."
	StartUpCluster           = "Start up a cluster."
	ConfigDescription        = "Configure a profile to store access settings for your MongoDB deployment."
	ConfigSetDescription     = "Configure specific properties of a profile."
	ConfigList               = "List available profiles."
	IAM                      = "Organization and projects operations."
	Organization             = "Organization operations."
	OrganizationLong         = "Create, list and manage your MongoDB organizations."
	CreateOrganization       = "Create an organization."
	DeleteOrganization       = "Delete an organization."
	ListOrganizations        = "List organizations."
	DescribeOrganizations    = "Describe an organizations."
	Projects                 = "Project operations."
	ProjectsLong             = "Create, list and manage your MongoDB projects."
	CreateProject            = "Create a project."
	DeleteProject            = "Delete a project."
	ListProjects             = "List projects."
	DescribeProject          = "Describe a project."
	DownloadLogCollectionJob = "Download logs generated by a log collection job."
	OpsManager               = "Ops Manager operations."
	Agents                   = "Manage Ops Manager agents."
	ListGlobalAlerts         = "List global alerts."
	Automation               = "Manage Ops Manager automation config."
	ShowAutomationStatus     = "Show the current status of the automation config."
	WatchAutomationStatus    = "Watch for automation changes."
	Global                   = "Manage Ops Manager global properties."
	Owner                    = "Manage Ops Manager owners."
	CreateOwner              = "Create the first user for Ops Manager."
	Servers                  = "Manage Ops Manager servers."
	ListServer               = "List all available servers running an automation agent for the given project."
	Security                 = "Manage clusters security configuration."
	EnableSecurity           = "Enable authentication mechanisms for the project."
	Events                   = "Manage events for your project."
	Metrics                  = "Get measurements on the state of the MongoDB process."
	LogCollection            = "Manage log collection jobs."
	StartLogCollectionJob    = "Start a job to collect logs."
	ListLogCollectionJobs    = "List log collection jobs."
	DeleteLogCollectionJob   = "Delete a log collection job."
	LogsDownloadLong         = `To download a log you need the name of the host where the log files that you want to download are stored.
To see the hostnames of your Atlas cluster, visit the cluster overview page in the Atlas UI.
The name of the log file must be one of: mongodb.gz, mongos.gz, mongodb-audit-log.gz, mongos-audit-log.gz
`
	DBUsersLong = `The dbusers command retrieves, creates and modifies the MongoDB database users in your cluster.
Each user has a set of roles that provide access to the project’s databases. 
A user’s roles apply to all the clusters in the project.`
	ConfigLongDescription = `Configure default values in a user profile.
All settings are optional, and can be set individually, to find out how run: 
  $ mongocli config set --help 

You can always use environment variables (MCLI_*) when running the tool,
To find out more check, https://docs.mongodb.com/mongocli/stable/configure/environment-variables/.`
	ConfigSetLong = `Configure specific properties of the profile.
Available properties include: %v.`
	CreateClusterLong = `You can create MongoDB clusters using this command.
To quickest way to get started is to just specify a name for your cluster and cloud provider and region to deploy, 
this will create a 3 member replica set with the latest available mongodb server version available.
Some of the cluster configuration options are available via flags but for full control of your deployment you can provide a config file.`
)
