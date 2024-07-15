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
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/flag"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/mocks"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/pointer"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/test"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/atlas-sdk/v20240530002/admin"
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
				HostConfig: &container.InspectDataHostConfig{
					PortBindings: map[string][]container.InspectDataHostPort{
						"27017/tcp": {
							{
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

	mockAtlasClusterDescriber := mocks.NewMockClusterDescriber(ctrl)
	deploymenTest := fixture.NewMockAtlasDeploymentOpts(ctrl, expectedAtlasDeployment)

	connectOpts := &ConnectOpts{
		ConnectWith:    "connectionString",
		DeploymentOpts: *deploymenTest.Opts,
		ConnectToAtlasOpts: ConnectToAtlasOpts{
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

	deploymenTest.CommonAtlasMocks(connectOpts.ProjectID)

	mockAtlasClusterDescriber.
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

func TestConnectBuilder(t *testing.T) {
	test.CmdValidator(
		t,
		ConnectBuilder(),
		0,
		// List flags that this command uses
		[]string{flag.ConnectWith, flag.ProjectID, flag.TypeFlag, flag.Username, flag.Password, flag.ConnectionStringType},
	)
}
