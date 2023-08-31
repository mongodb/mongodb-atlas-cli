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

//go:generate mockgen -destination=../mocks/mock_atlas_operator_cluster_store.go -package=mocks github.com/mongodb/mongodb-atlas-cli/internal/store AtlasOperatorClusterStore
//go:generate mockgen -destination=../mocks/mock_atlas_operator_project_store.go -package=mocks github.com/mongodb/mongodb-atlas-cli/internal/store AtlasOperatorProjectStore
//go:generate mockgen -destination=../mocks/mock_atlas_operator_db_users_store.go -package=mocks github.com/mongodb/mongodb-atlas-cli/internal/store AtlasOperatorDBUsersStore
//go:generate mockgen -destination=../mocks/mock_atlas_operator_org_store.go -package=mocks github.com/mongodb/mongodb-atlas-cli/internal/store AtlasOperatorOrgStore
//go:generate mockgen -destination=../mocks/mock_atlas_generic_store.go -package=mocks github.com/mongodb/mongodb-atlas-cli/internal/store AtlasOperatorGenericStore

type AtlasOperatorProjectStore interface {
	AtlasOperatorTeamsStore
	ProjectDescriber
	ProjectCreator
	ProjectLister
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
	ProjectAPIKeyCreator
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

type AtlasOperatorOrgStore interface {
	OrganizationAPIKeyCreator
	ProjectAPIKeyAssigner
}

type AtlasOperatorGenericStore interface {
	AtlasOperatorOrgStore
	AtlasOperatorProjectStore
	AtlasOperatorClusterStore
	AtlasOperatorDBUsersStore
	DataFederationStore
}
