// Copyright 2023 MongoDB Inc
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//	http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
package fixture

import (
	"github.com/golang/mock/gomock"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli/deployments/options"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/config"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/mocks"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/pointer"
	"go.mongodb.org/atlas-sdk/v20240530002/admin"
)

func NewMockAtlasDeploymentOpts(ctrl *gomock.Controller, deploymentName string) MockDeploymentOpts {
	mockCredentialsGetter := mocks.NewMockCredentialsGetter(ctrl)
	mockAtlasClusterListStore := mocks.NewMockClusterLister(ctrl)
	mockDeploymentTelemetry := mocks.NewMockDeploymentTelemetry(ctrl)

	return MockDeploymentOpts{
		ctrl:                      ctrl,
		MockCredentialsGetter:     mockCredentialsGetter,
		MockAtlasClusterListStore: mockAtlasClusterListStore,
		MockDeploymentTelemetry:   mockDeploymentTelemetry,
		Opts: &options.DeploymentOpts{
			CredStore:             mockCredentialsGetter,
			AtlasClusterListStore: mockAtlasClusterListStore,
			DeploymentName:        deploymentName,
			DeploymentType:        "atlas",
			DeploymentTelemetry:   mockDeploymentTelemetry,
		},
	}
}

func (m *MockDeploymentOpts) MockPaginatedAdvancedClusterDescription() *admin.PaginatedAdvancedClusterDescription {
	return &admin.PaginatedAdvancedClusterDescription{
		Results: &[]admin.AdvancedClusterDescription{
			{
				Name:           pointer.Get(m.Opts.DeploymentName),
				Id:             pointer.Get("123"),
				MongoDBVersion: pointer.Get("7.0.0"),
				StateName:      pointer.Get("IDLE"),
				Paused:         pointer.Get(false),
			},
		},
	}
}

func (m *MockDeploymentOpts) CommonAtlasMocks(projectID string) {
	m.MockCredentialsGetter.
		EXPECT().
		AuthType().
		Return(config.OAuth).
		Times(1)

	m.MockAtlasClusterListStore.
		EXPECT().
		ProjectClusters(projectID, gomock.Any()).
		Return(m.MockPaginatedAdvancedClusterDescription(), nil).
		Times(1)

	m.MockDeploymentTelemetry.
		EXPECT().
		AppendDeploymentType().
		Times(1)
}
