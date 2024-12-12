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
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli/deployments/test/fixture"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/container"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/mocks"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/pointer"
	"github.com/stretchr/testify/assert"
	atlasClustersPinned "go.mongodb.org/atlas-sdk/v20240530005/admin"
)

const (
	expectedLocalDeployment = "localDeployment1"
	expectedAtlasDeployment = "atlasCluster1"
)

func TestRun_ConnectLocal(t *testing.T) {
	ctrl := gomock.NewController(t)
	ctx := context.Background()
	buf := new(bytes.Buffer)

	deploymenTest := fixture.NewMockLocalDeploymentOpts(ctrl, expectedLocalDeployment)
	connectOpts := &ConnectOpts{
		ConnectWith:    "connectionString",
		DeploymentOpts: *deploymenTest.Opts,
		OutputOpts: cli.OutputOpts{
			OutWriter: buf,
		},
	}

	deploymenTest.LocalMockFlow(ctx)

	deploymenTest.MockContainerEngine.
		EXPECT().
		ContainerInspect(ctx, expectedLocalDeployment).
		Return([]*container.InspectData{
			{
				Name: expectedLocalDeployment,
				Config: &container.InspectDataConfig{
					Labels: map[string]string{
						"version": "7.0.1",
					},
				},
				NetworkSettings: &container.NetworkSettings{
					Ports: map[string][]container.InspectDataHostPort{
						"27017/tcp": {
							container.InspectDataHostPort{
								HostIP:   "127.0.0.1",
								HostPort: "27017",
							},
						},
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

	mockAtlasClusterDescriberStarter := mocks.NewMockClusterDescriberStarter(ctrl)
	deploymenTest := fixture.NewMockAtlasDeploymentOpts(ctrl, expectedAtlasDeployment)

	connectOpts := &ConnectOpts{
		ConnectWith:    "connectionString",
		DeploymentOpts: *deploymenTest.Opts,
		ConnectToAtlasOpts: ConnectToAtlasOpts{
			Store: mockAtlasClusterDescriberStarter,
			ProjectOpts: cli.ProjectOpts{
				ProjectID: "projectID",
			},
			ConnectionStringType: "standard",
		},
		OutputOpts: cli.OutputOpts{
			OutWriter: buf,
		},
	}

	expectedAtlasClusters := &atlasClustersPinned.PaginatedAdvancedClusterDescription{
		Results: &[]atlasClustersPinned.AdvancedClusterDescription{
			{
				Name:           pointer.Get(expectedAtlasDeployment),
				Id:             pointer.Get("123"),
				MongoDBVersion: pointer.Get("7.0.0"),
				StateName:      pointer.Get("IDLE"),
				Paused:         pointer.Get(false),
				ConnectionStrings: &atlasClustersPinned.ClusterConnectionStrings{
					StandardSrv: pointer.Get("mongodb://localhost:27017/?directConnection=true"),
				},
			},
		},
	}

	deploymenTest.CommonAtlasMocks(connectOpts.ProjectID)

	mockAtlasClusterDescriberStarter.
		EXPECT().
		AtlasCluster(connectOpts.ProjectID, expectedAtlasDeployment).
		Return(&expectedAtlasClusters.GetResults()[0], nil).
		Times(1)

	if err := Run(ctx, connectOpts); err != nil {
		t.Fatalf("Run() unexpected error: %v", err)
	}

	assert.Equal(t, `mongodb://localhost:27017/?directConnection=true
`, buf.String())
}

func TestPostRun(t *testing.T) {
	ctrl := gomock.NewController(t)
	deploymentsTest := fixture.NewMockLocalDeploymentOpts(ctrl, "localDeployment")
	buf := new(bytes.Buffer)

	opts := &ConnectOpts{
		DeploymentOpts: *deploymentsTest.Opts,
		OutputOpts: cli.OutputOpts{
			OutWriter: buf,
		},
	}

	deploymentsTest.
		MockDeploymentTelemetry.
		EXPECT().
		AppendDeploymentType().
		Times(1)

	PostRun(opts)
}
