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

//go:build unit

package deployments

import (
	"bytes"
	"context"
	"fmt"
	"testing"

	"github.com/containers/podman/v4/libpod/define"
	"github.com/golang/mock/gomock"
	"github.com/mongodb/mongodb-atlas-cli/internal/cli"
	"github.com/mongodb/mongodb-atlas-cli/internal/cli/atlas/deployments/options"
	"github.com/mongodb/mongodb-atlas-cli/internal/config"
	"github.com/mongodb/mongodb-atlas-cli/internal/flag"
	"github.com/mongodb/mongodb-atlas-cli/internal/mocks"
	"github.com/mongodb/mongodb-atlas-cli/internal/podman"
	"github.com/mongodb/mongodb-atlas-cli/internal/pointer"
	"github.com/mongodb/mongodb-atlas-cli/internal/test"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/atlas-sdk/v20230201008/admin"
)

func TestStart_RunLocal_PausedContainers(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockStore := mocks.NewMockClusterStarter(ctrl)
	mockCredentialsGetter := mocks.NewMockCredentialsGetter(ctrl)
	mockProfileReader := mocks.NewMockProfileReader(ctrl)
	mockPodman := mocks.NewMockClient(ctrl)
	ctx := context.Background()

	expectedLocalDeployments := []*podman.Container{
		{
			Names:  []string{"localTest2"},
			State:  "paused",
			Labels: map[string]string{"version": "6.0.9"},
			ID:     deploymentName,
		},
		{
			Names:  []string{"localTest1"},
			State:  "paused",
			Labels: map[string]string{"version": "7.0.0"},
			ID:     deploymentName,
		},
	}

	buf := new(bytes.Buffer)
	startOpts := &StartOpts{
		store: mockStore,
		DeploymentOpts: options.DeploymentOpts{
			Config:         mockProfileReader,
			PodmanClient:   mockPodman,
			CredStore:      mockCredentialsGetter,
			DeploymentName: deploymentName,
			DeploymentType: "LOCAL",
		},
		GlobalOpts: cli.GlobalOpts{
			ProjectID: projectID,
		},
		OutputOpts: cli.OutputOpts{
			Template:  startTemplate,
			OutWriter: buf,
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
		Return(expectedLocalDeployments, nil).
		Times(1)

	mockPodman.
		EXPECT().
		ContainerInspect(ctx, startOpts.LocalMongotHostname()).
		Return([]*define.InspectContainerData{
			{
				NetworkSettings: &define.InspectNetworkSettings{
					Networks: map[string]*define.InspectAdditionalNetwork{
						startOpts.LocalNetworkName(): {
							InspectBasicNetworkConfig: define.InspectBasicNetworkConfig{
								IPAddress: "1.2.3.4",
							},
						},
					},
				},
			},
		}, nil).
		Times(1)

	mockPodman.
		EXPECT().
		Exec(gomock.Any(), startOpts.LocalMongodHostname(), "/bin/sh", "-c", gomock.Any()).
		Return(nil).
		Times(1)

	mockPodman.
		EXPECT().
		UnpauseContainers(ctx, startOpts.LocalMongodHostname(), startOpts.LocalMongotHostname()).
		Return(nil, nil).
		Times(1)

	if err := startOpts.Run(ctx); err != nil {
		t.Fatalf("Run() unexpected error: %v", err)
	}

	assert.Equal(t, fmt.Sprintf("Starting deployment '%s'.\n", deploymentName), buf.String())
	t.Log(buf.String())
}

func TestStart_RunLocal_StoppedContainers(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockStore := mocks.NewMockClusterStarter(ctrl)
	mockCredentialsGetter := mocks.NewMockCredentialsGetter(ctrl)
	mockProfileReader := mocks.NewMockProfileReader(ctrl)
	mockPodman := mocks.NewMockClient(ctrl)
	ctx := context.Background()

	expectedLocalDeployments := []*podman.Container{
		{
			Names:  []string{"localTest2"},
			State:  "exited",
			Labels: map[string]string{"version": "6.0.9"},
			ID:     deploymentName,
		},
		{
			Names:  []string{"localTest1"},
			State:  "exited",
			Labels: map[string]string{"version": "7.0.0"},
			ID:     deploymentName,
		},
	}

	buf := new(bytes.Buffer)
	startOpts := &StartOpts{
		store: mockStore,
		DeploymentOpts: options.DeploymentOpts{
			PodmanClient:   mockPodman,
			CredStore:      mockCredentialsGetter,
			Config:         mockProfileReader,
			DeploymentName: deploymentName,
			DeploymentType: "LOCAL",
		},
		GlobalOpts: cli.GlobalOpts{
			ProjectID: projectID,
		},
		OutputOpts: cli.OutputOpts{
			Template:  startTemplate,
			OutWriter: buf,
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
		Return(expectedLocalDeployments, nil).
		Times(1)

	mockPodman.
		EXPECT().
		ContainerInspect(ctx, startOpts.LocalMongotHostname()).
		Return([]*define.InspectContainerData{
			{
				NetworkSettings: &define.InspectNetworkSettings{
					Networks: map[string]*define.InspectAdditionalNetwork{
						startOpts.LocalNetworkName(): {
							InspectBasicNetworkConfig: define.InspectBasicNetworkConfig{
								IPAddress: "1.2.3.4",
							},
						},
					},
				},
			},
		}, nil).
		Times(1)

	mockPodman.
		EXPECT().
		Exec(gomock.Any(), startOpts.LocalMongodHostname(), "/bin/sh", "-c", gomock.Any()).
		Return(nil).
		Times(1)

	mockPodman.
		EXPECT().
		StartContainers(ctx, startOpts.LocalMongodHostname(), startOpts.LocalMongotHostname()).
		Return(nil, nil).
		Times(1)

	if err := startOpts.Run(ctx); err != nil {
		t.Fatalf("Run() unexpected error: %v", err)
	}

	assert.Equal(t, fmt.Sprintf("Starting deployment '%s'.\n", deploymentName), buf.String())
	t.Log(buf.String())
}

func TestStart_RunAtlas(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockStore := mocks.NewMockClusterStarter(ctrl)
	mockCredentialsGetter := mocks.NewMockCredentialsGetter(ctrl)
	mockProfileReader := mocks.NewMockProfileReader(ctrl)
	mockPodman := mocks.NewMockClient(ctrl)
	mockAtlasListStore := mocks.NewMockClusterLister(ctrl)
	ctx := context.Background()
	deploymentName := "atlasCluster1"

	expectedAtlasClusters := &admin.PaginatedAdvancedClusterDescription{
		Results: []admin.AdvancedClusterDescription{
			{
				Name:           pointer.Get(deploymentName),
				Id:             pointer.Get("123"),
				MongoDBVersion: pointer.Get("7.0.0"),
				StateName:      pointer.Get("IDLE"),
				Paused:         pointer.Get(false),
			},
		},
	}

	buf := new(bytes.Buffer)
	listOpts := &StartOpts{
		store: mockStore,
		DeploymentOpts: options.DeploymentOpts{
			PodmanClient:          mockPodman,
			CredStore:             mockCredentialsGetter,
			Config:                mockProfileReader,
			DeploymentName:        deploymentName,
			DeploymentType:        "ATLAS",
			AtlasClusterListStore: mockAtlasListStore,
		},
		GlobalOpts: cli.GlobalOpts{
			ProjectID: projectID,
		},
		OutputOpts: cli.OutputOpts{
			Template:  startTemplate,
			OutWriter: buf,
		},
	}

	mockPodman.
		EXPECT().
		Ready(ctx).
		Return(nil).
		Times(0)

	mockCredentialsGetter.
		EXPECT().
		AuthType().
		Return(config.OAuth).
		Times(2)

	mockAtlasListStore.
		EXPECT().
		ProjectClusters(listOpts.ProjectID, gomock.Any()).
		Return(expectedAtlasClusters, nil).
		Times(1)

	mockStore.
		EXPECT().
		StartCluster(projectID, deploymentName).
		Return(
			&admin.AdvancedClusterDescription{
				Name: pointer.GetStringPointerIfNotEmpty(deploymentName),
			}, nil).
		Times(1)

	if err := listOpts.Run(ctx); err != nil {
		t.Fatalf("Run() unexpected error: %v", err)
	}

	assert.Equal(t, fmt.Sprintf("Starting deployment '%s'.\n", deploymentName), buf.String())
	t.Log(buf.String())
}

func TestStartBuilder(t *testing.T) {
	test.CmdValidator(
		t,
		StartBuilder(),
		0,
		[]string{flag.ProjectID, flag.TypeFlag},
	)
}
