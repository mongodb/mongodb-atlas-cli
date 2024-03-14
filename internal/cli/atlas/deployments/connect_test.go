// Copyright 2023 MongoDB Inc
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

package deployments

import (
	"bytes"
	"context"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli/atlas/deployments/options"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/config"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/flag"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/mocks"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/podman"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/pointer"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/store"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/test"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/atlas-sdk/v20231115008/admin"
)

const (
	expectedLocalDeployment = "localDeployment1"
	expectedAtlasDeployment = "atlasCluster1"
)

func TestRun_ConnectLocal(t *testing.T) {
	ctrl := gomock.NewController(t)
	ctx := context.Background()
	mockPodman := mocks.NewMockClient(ctrl)
	buf := new(bytes.Buffer)

	connectOpts := &options.ConnectOpts{
		ConnectWith: "connectionString",
		DeploymentOpts: options.DeploymentOpts{
			PodmanClient:   mockPodman,
			DeploymentName: expectedLocalDeployment,
			DeploymentType: "local",
		},
		OutputOpts: cli.OutputOpts{
			OutWriter: buf,
		},
	}

	expectedContainers := []*podman.Container{
		{
			Names:  []string{expectedLocalDeployment},
			State:  "running",
			Labels: map[string]string{"version": "6.0.9"},
			ID:     expectedLocalDeployment,
		},
	}

	mockPodman.
		EXPECT().
		Ready(ctx).
		Return(nil).
		Times(1)

	mockPodman.
		EXPECT().
		ListContainers(ctx, options.MongodHostnamePrefix).
		Return(expectedContainers, nil).
		Times(1)

	mockPodman.
		EXPECT().
		ContainerInspect(ctx, options.MongodHostnamePrefix+"-"+expectedLocalDeployment).
		Return([]*podman.InspectContainerData{
			{
				Name: options.MongodHostnamePrefix + "-" + expectedLocalDeployment,
				Config: &podman.InspectContainerConfig{
					Labels: map[string]string{
						"version": "7.0.1",
					},
				},
				HostConfig: &podman.InspectContainerHostConfig{
					PortBindings: map[string][]podman.InspectHostPort{
						"27017/tcp": {
							{
								HostIP:   "127.0.0.1",
								HostPort: "27017",
							},
						},
					},
				},
				Mounts: []podman.InspectMount{
					{
						Name: connectOpts.DeploymentOpts.LocalMongodDataVolume(),
					},
				},
			},
		}, nil).
		Times(1)

	if err := Run(ctx, connectOpts); err != nil {
		t.Fatalf("Run() unexpected error: %v", err)
	}

	assert.Equal(t, `mongodb://localhost:27017/?directConnection=true
`, buf.String())
}

func TestRun_ConnectAtlas(t *testing.T) {
	ctrl := gomock.NewController(t)
	ctx := context.Background()
	buf := new(bytes.Buffer)

	mockAtlasClusterListStore := mocks.NewMockClusterLister(ctrl)
	mockCredentialsGetter := mocks.NewMockCredentialsGetter(ctrl)
	mockAtlasClusterDescriber := mocks.NewMockClusterDescriber(ctrl)

	connectOpts := &options.ConnectOpts{
		ConnectWith: "connectionString",
		DeploymentOpts: options.DeploymentOpts{
			AtlasClusterListStore: mockAtlasClusterListStore,
			DeploymentName:        expectedAtlasDeployment,
			DeploymentType:        "atlas",
			CredStore:             mockCredentialsGetter,
		},
		ConnectToAtlasOpts: options.ConnectToAtlasOpts{
			Store: mockAtlasClusterDescriber,
			GlobalOpts: cli.GlobalOpts{
				ProjectID: "projectID",
			},
			ConnectionStringType: "standard",
		},
		OutputOpts: cli.OutputOpts{
			OutWriter: buf,
		},
	}

	expectedAtlasClusters := &admin.PaginatedAdvancedClusterDescription{
		Results: &[]admin.AdvancedClusterDescription{
			{
				Name:           pointer.Get(expectedAtlasDeployment),
				Id:             pointer.Get("123"),
				MongoDBVersion: pointer.Get("7.0.0"),
				StateName:      pointer.Get("IDLE"),
				Paused:         pointer.Get(false),
				ConnectionStrings: &admin.ClusterConnectionStrings{
					StandardSrv: pointer.Get("mongodb://localhost:27017/?directConnection=true"),
				},
			},
		},
	}

	mockAtlasClusterListStore.
		EXPECT().
		ProjectClusters(connectOpts.ProjectID,
			&store.ListOptions{
				PageNum:      cli.DefaultPage,
				ItemsPerPage: options.MaxItemsPerPage,
			},
		).
		Return(expectedAtlasClusters, nil).
		Times(1)

	mockAtlasClusterDescriber.
		EXPECT().
		AtlasCluster(connectOpts.ProjectID, expectedAtlasDeployment).
		Return(&expectedAtlasClusters.GetResults()[0], nil).
		Times(1)

	mockCredentialsGetter.
		EXPECT().
		AuthType().
		Return(config.OAuth).
		Times(1)

	if err := Run(ctx, connectOpts); err != nil {
		t.Fatalf("Run() unexpected error: %v", err)
	}

	assert.Equal(t, `mongodb://localhost:27017/?directConnection=true
`, buf.String())
}

func TestConnectBuilder(t *testing.T) {
	test.CmdValidator(
		t,
		ConnectBuilder(),
		0,
		// List flags that this command uses
		[]string{flag.ConnectWith, flag.ProjectID, flag.TypeFlag, flag.Username, flag.Password, flag.ConnectionStringType},
	)
}
