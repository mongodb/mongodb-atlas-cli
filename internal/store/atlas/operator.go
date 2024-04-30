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

package atlas

import (
	"errors"

	"github.com/andreaangiolillo/mongocli-test/internal/pointer"
)

//go:generate mockgen -destination=../../mocks/atlas/mock_atlas_operator_cluster_store.go -package=atlas github.com/andreaangiolillo/mongocli-test/internal/store/atlas OperatorClusterStore
//go:generate mockgen -destination=../../mocks/atlas/mock_atlas_operator_project_store.go -package=atlas github.com/andreaangiolillo/mongocli-test/internal/store/atlas OperatorProjectStore
//go:generate mockgen -destination=../../mocks/atlas/mock_atlas_operator_db_users_store.go -package=atlas github.com/andreaangiolillo/mongocli-test/internal/store/atlas OperatorDBUsersStore
//go:generate mockgen -destination=../../mocks/atlas/mock_atlas_operator_org_store.go -package=atlas github.com/andreaangiolillo/mongocli-test/internal/store/atlas OperatorOrgStore
//go:generate mockgen -destination=../../mocks/atlas/mock_atlas_generic_store.go -package=atlas github.com/andreaangiolillo/mongocli-test/internal/store/atlas OperatorGenericStore

type ListOptions struct {
	PageNum      int
	ItemsPerPage int
}

type ContainersListOptions struct {
	ListOptions
	ProviderName string
}

var errUnsupportedService = errors.New("unsupported service")

func StringOrEmpty(s *string) string {
	return pointer.GetOrDefault(s, "")
}

type OperatorProjectStore interface {
	OperatorTeamsStore
	ProjectDescriber
	ProjectCreator
	ProjectLister
	OrgProjectLister
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

type OperatorDBUsersStore interface {
	DatabaseUserLister
}

type OperatorClusterStore interface {
	AllClustersLister
	ClusterDescriber
	ClusterConfigurationOptionsDescriber
	ScheduleDescriber
	ServerlessInstanceDescriber
	ServerlessPrivateEndpointsLister
	GlobalClusterDescriber
}

type AllClustersLister interface {
	ClusterLister
	ServerlessInstanceLister
}

type OperatorTeamsStore interface {
	TeamDescriber
	ProjectTeamLister
	TeamUserLister
}

type OperatorOrgStore interface {
	OrganizationAPIKeyCreator
	ProjectAPIKeyAssigner
}

type OperatorGenericStore interface {
	OperatorOrgStore
	OperatorProjectStore
	OperatorClusterStore
	OperatorDBUsersStore
	DataFederationStore
}
