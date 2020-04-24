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
	Atlas                  = "Atlas operations."
	Alerts                 = "Manage alerts for your project."
	AcknowledgeAlerts      = "Acknowledge an alert."
	Config                 = "Manage alerts configuration for your project."
	ConfigLong             = "The configs command provides access to your alerts configurations. You can create, edit, and delete alert configurations."
	CreateAlertsConfig     = "Create an alert configuration for a project."
	DeleteAlertsConfig     = "Delete an alert config."
	AlertsConfigFields     = "Manage alert configuration fields for your project."
	AlertsConfigFieldsType = "List alert configurations available field types."
	ListAlertsConfigs      = "List alert configurations for a project."
	UpdateAlertsConfig     = "Update an alert configuration for a project."
	DescribeAlert          = "Describe an alert for a project."
	ListAlerts             = "List alerts for a project."
	Backup                 = "Manage continuous backups for your project."
	Checkpoints            = "Manage backup checkpoints for your project."
	ListCheckpoints        = "List continuous backup checkpoints."
	Restores               = "Manage restore jobs."
	ListRestores           = "Lists restore jobs for a project and cluster."
	StartRestore           = "Start a restore job."
	Snapshots              = "Manage continuous snapshots for your project."
	ListSnapshots          = "List continuous snapshots for a project."
	Logs                   = "Download host logs."
	Clusters               = "Manage clusters for your project."
	ClustersIndexes        = "Manage cluster indexes for your project."
	ClustersLong           = "The clusters command provides access to your cluster configurations. You can create, edit, and delete clusters."
	CreateCluster          = "Create a MongoDB cluster."
	CreateIndex            = "Create a rolling index for your MongoDB cluster."
	ApplyCluster           = "Apply a new cluster configuration."
	DeleteCluster          = "Delete a cluster."
	DescribeCluster        = "Describe a cluster."
	ListClusters           = "List clusters for a project."
	UpdateCluster          = "Update a MongoDB cluster in Atlas."
	PauseCluster           = "Pause a running MongoDB cluster in Atlas."
	StartCluster           = "Start a paused MongoDB cluster in Atlas."
	Processes              = "Manage processes for your project."
	ListProcesses          = "List processes for a project."
	DBUsers                = "Manage database users for your project."
	CreateDBUser           = "Create a database user for a project."
	DeleteDBUser           = "Delete a database user for a project."
	ListDBUsers            = "List Atlas database users for a project."
	ListEvents             = "List events for an organization or project"
	UpdateDBUser           = "Update a MongoDB dbuser in Atlas."
	ProcessMeasurements    = "Get measurements for a given host."
	Disks                  = "List available disks or disks measurements for a given host."
	ListDisks              = "List available disks for a given host."
	DescribeDisks          = "Describe disks measurements for a given host partition."
	Databases              = "List available databases or databases measurements for a given host."
	ListDatabases          = "List available databases for a given host."
	Whitelist              = "Manage the IP whitelist for a project."
	CreateWhitelist        = "Create an IP whitelist for a project."
	DeleteWhitelist        = "Delete a database user for a project."
	DescribeWhitelist      = "Describe an Atlas whitelist."
	ListWhitelist          = "List Atlas whitelist for a project."
	CloudManager           = "Cloud Manager operations."
	ShutdownOMCluster      = "Shutdown a Cloud Manager cluster."
	StartUpOMCluster       = "Startup a Cloud Manager cluster."
	UpdateOMCluster        = "Update a Cloud Manager cluster."
	ConfigDescription      = "Configure a profile. This let you store access settings to your cloud."
	ConfigSetDescription   = "Configure specific properties of the profile."
	IAM                    = "Authentication operations."
	Organization           = "Organization operations."
	OrganizationLong       = "Create, list and manage your MongoDB organizations."
	CreateOrganization     = "Create an organization."
	DeleteOrganization     = "Delete an organization."
	ListOrganizations      = "List organizations."
	Projects               = "Project operations."
	ProjectsLong           = "Create, list and manage your MongoDB projects."
	CreateProject          = "Create a project."
	DeleteProject          = "Delete a project."
	ListProjects           = "List projects."
	OpsManager             = "Ops Manager operations."
	ListGlobalAlerts       = "List global alerts."
	Automation             = "Manage Ops Manager automation config."
	ShowAutomationStatus   = "Show the current status of the automation config."
	Global                 = "Manage Ops Manager global properties."
	Owner                  = "Manage Ops Manager owners."
	CreateOwner            = "Create the first user for Ops Manager."
	Servers                = "Manage Ops Manager servers."
	ListServer             = "List all available servers running an automation agent for the given project."
	Security               = "Manage clusters security configuration."
	EnableSecurity         = "Enable authentication mechanisms for the project."
	Events                 = "Manage events for your project."
	Measurements           = "Get measurements on the state of the MongoDB process."
	LogCollection          = "Manage log collection jobs."
	StartLogCollectionJob  = "Start a job to collect logs."
	DBUsersLong            = `The dbusers command retrieves, creates and modifies the MongoDB database users in your cluster.
Each user has a set of roles that provide access to the project’s databases. 
A user’s roles apply to all the clusters in the project.`
	ConfigLongDescription = `Setting API keys is optional and env variables can be used instead.
Leaving any option empty won't override existing stored values.`
	ConfigSetLongDescription = `Configure specific properties of the profile.
Available properties include: %v.`
)
