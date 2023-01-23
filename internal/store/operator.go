// Copyright 2022 MongoDB Inc
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

package store

type AtlasOperatorProjectStore interface {
	AtlasOperatorTeamsStore
	ProjectDescriber
	ProjectIPAccessListLister
	ProjectSettingsDescriber
	IntegrationLister
	MaintenanceWindowDescriber
	PrivateEndpointLister
	CloudProviderAccessRoleLister
	PeeringConnectionLister
	EncryptionAtRestDescriber
	AuditingDescriber
	AlertConfigurationLister
	DatabaseRoleLister
}

type AtlasOperatorDBUsersStore interface {
	DatabaseUserLister
}

type AtlasOperatorClusterStore interface {
	AtlasAllClustersLister
	AtlasClusterDescriber
	AtlasClusterConfigurationOptionsDescriber
	ScheduleDescriber
	ServerlessInstanceDescriber
	ServerlessPrivateEndpointsLister
	GlobalClusterDescriber
}

type AtlasAllClustersLister interface {
	ClusterLister
	ServerlessInstanceLister
}

type AtlasOperatorTeamsStore interface {
	TeamDescriber
	ProjectTeamLister
	TeamUserLister
}

type AtlasOperatorGenericStore interface {
	AtlasOperatorProjectStore
	AtlasOperatorClusterStore
	AtlasOperatorDBUsersStore
}
