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

//go:build unit

package indexes

import (
	"context"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli/atlas/deployments/options"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli/atlas/deployments/test/fixture"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/flag"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/mocks"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/podman"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/test"
)

func TestDelete_RunLocal(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockMongodbClient := mocks.NewMockMongoDBClient(ctrl)
	mockStore := mocks.NewMockSearchIndexDeleter(ctrl)
	ctx := context.Background()

	const (
		expectedLocalDeployment = "localDeployment1"
		indexID                 = "1"
	)

	deploymentTest := fixture.NewMockLocalDeploymentOpts(ctrl, expectedLocalDeployment)
	mockPodman := deploymentTest.MockPodman

	opts := &DeleteOpts{
		DeploymentOpts: *deploymentTest.Opts,
		DeleteOpts: &cli.DeleteOpts{
			Entry:   indexID,
			Confirm: true,
		},
		mongodbClient: mockMongodbClient,
		store:         mockStore,
	}

	deploymentTest.LocalMockFlow(ctx)

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
						Name: opts.DeploymentOpts.LocalMongodDataVolume(),
					},
				},
			},
		}, nil).
		Times(1)
	mockMongodbClient.
		EXPECT().
		Disconnect().
		Times(1)

	mockMongodbClient.
		EXPECT().
		Connect("mongodb://localhost:27017/?directConnection=true", int64(10)).
		Return(nil).
		Times(1)

	mockMongodbClient.
		EXPECT().
		DeleteSearchIndex(indexID).
		Return(nil).
		Times(1)

	if err := opts.Run(ctx); err != nil {
		t.Fatalf("Run() unexpected error: %v", err)
	}
}

func TestDelete_RunAtlas(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockMongodbClient := mocks.NewMockMongoDBClient(ctrl)
	mockStore := mocks.NewMockSearchIndexDeleter(ctrl)
	ctx := context.Background()

	const (
		expectedLocalDeployment = "localDeployment1"
		indexID                 = "1"
		projectID               = "1"
	)

	deploymentTest := fixture.NewMockAtlasDeploymentOpts(ctrl, expectedLocalDeployment)

	opts := &DeleteOpts{
		DeploymentOpts: *deploymentTest.Opts,
		DeleteOpts: &cli.DeleteOpts{
			Entry:   indexID,
			Confirm: true,
		},
		GlobalOpts: cli.GlobalOpts{
			ProjectID: projectID,
		},
		mongodbClient: mockMongodbClient,
		store:         mockStore,
	}

	deploymentTest.CommonAtlasMocks(projectID)

	mockStore.
		EXPECT().
		DeleteSearchIndex(opts.ProjectID, opts.DeploymentName, opts.Entry).
		Return(nil).
		Times(1)

	if err := opts.Run(ctx); err != nil {
		t.Fatalf("Run() unexpected error: %v", err)
	}
}

func TestDeleteBuilder(t *testing.T) {
	test.CmdValidator(
		t,
		DeleteBuilder(),
		0,
		[]string{flag.DeploymentName, flag.ProjectID, flag.Force, flag.TypeFlag},
	)
}
