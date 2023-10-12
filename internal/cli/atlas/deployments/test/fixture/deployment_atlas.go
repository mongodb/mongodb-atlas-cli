package fixture

import (
	"github.com/golang/mock/gomock"
	"github.com/mongodb/mongodb-atlas-cli/internal/cli/atlas/deployments/options"
	"github.com/mongodb/mongodb-atlas-cli/internal/config"
	"github.com/mongodb/mongodb-atlas-cli/internal/mocks"
	"github.com/mongodb/mongodb-atlas-cli/internal/pointer"
	"go.mongodb.org/atlas-sdk/v20230201008/admin"
)

func NewMockAtlasDeploymentOpts(ctrl *gomock.Controller, deploymentName string) MockDeploymentOpts {
	mockCredentialsGetter := mocks.NewMockCredentialsGetter(ctrl)
	mockAtlasClusterListStore := mocks.NewMockClusterLister(ctrl)

	return MockDeploymentOpts{
		ctrl:                      ctrl,
		MockCredentialsGetter:     mockCredentialsGetter,
		MockAtlasClusterListStore: mockAtlasClusterListStore,
		Opts: &options.DeploymentOpts{
			CredStore:             mockCredentialsGetter,
			AtlasClusterListStore: mockAtlasClusterListStore,
			DeploymentName:        deploymentName,
			DeploymentType:        "atlas",
		},
	}
}

func (m *MockDeploymentOpts) MockPaginatedAdvancedClusterDescription() *admin.PaginatedAdvancedClusterDescription {
	return &admin.PaginatedAdvancedClusterDescription{
		Results: []admin.AdvancedClusterDescription{
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
}
