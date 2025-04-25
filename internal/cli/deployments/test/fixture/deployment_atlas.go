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
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli/deployments/options"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/config"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/mocks"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/pointer"
	atlasClustersPinned "go.mongodb.org/atlas-sdk/v20240530005/admin"
	"go.uber.org/mock/gomock"
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

func (m *MockDeploymentOpts) MockPaginatedAdvancedClusterDescription(state string) *atlasClustersPinned.PaginatedAdvancedClusterDescription {
	return &atlasClustersPinned.PaginatedAdvancedClusterDescription{
		Results: &[]atlasClustersPinned.AdvancedClusterDescription{
			{
				Name:           pointer.Get(m.Opts.DeploymentName),
				Id:             pointer.Get("123"),
				MongoDBVersion: pointer.Get("7.0.0"),
				StateName:      &state,
				Paused:         pointer.Get(false),
			},
		},
	}
}

func (m *MockDeploymentOpts) CommonAtlasMocks(projectID string) {
	m.CommonAtlasMocksWithState(projectID, "IDLE")
}

func (m *MockDeploymentOpts) CommonAtlasMocksWithState(projectID string, state string) {
	m.MockCredentialsGetter.
		EXPECT().
		AuthType().
		Return(config.OAuth).
		Times(1)

	m.MockAtlasClusterListStore.
		EXPECT().
		ProjectClusters(projectID, gomock.Any()).
		Return(m.MockPaginatedAdvancedClusterDescription(state), nil).
		Times(1)

	m.MockDeploymentTelemetry.
		EXPECT().
		AppendDeploymentType().
		Times(1)
}
