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

	"github.com/containers/podman/v4/libpod/define"
	"github.com/golang/mock/gomock"
	"github.com/mongodb/mongodb-atlas-cli/internal/cli"
	"github.com/mongodb/mongodb-atlas-cli/internal/cli/atlas/deployments/options"
	"github.com/mongodb/mongodb-atlas-cli/internal/flag"
	"github.com/mongodb/mongodb-atlas-cli/internal/mocks"
	"github.com/mongodb/mongodb-atlas-cli/internal/test"
)

func TestDelete_Run_Atlas(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockAtlasStore := mocks.NewMockClusterDeleter(ctrl)
	mockCredentialsGetter := mocks.NewMockCredentialsGetter(ctrl)
	ctx := context.Background()

	buf := new(bytes.Buffer)
	opts := &DeleteOpts{
		atlasStore: mockAtlasStore,
		DeploymentOpts: options.DeploymentOpts{
			CredStore:      mockCredentialsGetter,
			DeploymentName: "testDeployment",
			DeploymentType: "atlas",
		},
		GlobalOpts: cli.GlobalOpts{
			ProjectID: "64f670f0bf789926667dad1a",
		},
		OutputOpts: cli.OutputOpts{
			OutWriter: buf,
		},
		DeleteOpts: cli.NewDeleteOpts(deleteSuccessMessage, deleteFailMessage),
	}
	opts.Confirm = true

	mockAtlasStore.
		EXPECT().
		DeleteCluster(opts.ProjectID, opts.DeploymentName).
		Return(nil).
		Times(1)

	if err := opts.Run(ctx); err != nil {
		t.Fatalf("Run() unexpected error: %v", err)
	}

	t.Log(buf.String())
}

func TestDelete_Run_Local(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockPodman := mocks.NewMockClient(ctrl)
	ctx := context.Background()
	buf := new(bytes.Buffer)

	expectedLocalDeployment := "testDeployment"
	opts := &DeleteOpts{
		DeploymentOpts: options.DeploymentOpts{
			PodmanClient:   mockPodman,
			DeploymentName: expectedLocalDeployment,
			DeploymentType: "local",
		},
		GlobalOpts: cli.GlobalOpts{
			ProjectID: "64f670f0bf789926667dad1a",
		},
		OutputOpts: cli.OutputOpts{
			OutWriter: buf,
		},
		DeleteOpts: cli.NewDeleteOpts(deleteSuccessMessage, deleteFailMessage),
	}
	opts.Confirm = true

	mockPodman.
		EXPECT().
		Ready(ctx).
		Return(nil).
		Times(1)

	mockPodman.
		EXPECT().
		RemoveContainers(ctx, options.MongodHostnamePrefix+"-"+expectedLocalDeployment).
		Return(nil, nil).
		Times(1)

	mockPodman.
		EXPECT().
		RemoveNetworks(ctx, "mdb-local-"+expectedLocalDeployment).
		Return(nil, nil).
		Times(1)

	mockPodman.
		EXPECT().
		RemoveVolumes(ctx,
			"mongod-local-data-"+expectedLocalDeployment,
			"mongot-local-data-"+expectedLocalDeployment,
			"mongot-local-metrics-"+expectedLocalDeployment,
		).
		Return(nil, nil).
		Times(1)

	mockPodman.
		EXPECT().
		ContainerInspect(ctx, options.MongodHostnamePrefix+"-"+expectedLocalDeployment).
		Return([]*define.InspectContainerData{
			{
				Name: options.MongodHostnamePrefix + "-" + expectedLocalDeployment,
				Config: &define.InspectContainerConfig{
					Labels: map[string]string{
						"version": "7.0.1",
					},
				},
				HostConfig: &define.InspectContainerHostConfig{
					PortBindings: map[string][]define.InspectHostPort{
						"27017/tcp": {
							{
								HostIP:   "127.0.0.1",
								HostPort: "27017",
							},
						},
					},
				},
				Mounts: []define.InspectMount{
					{
						Name: "mongod-local-data-" + expectedLocalDeployment,
					},
				},
			},
		}, nil).
		Times(1)

	if err := opts.Run(ctx); err != nil {
		t.Fatalf("Run() unexpected error: %v", err)
	}

	t.Log(buf.String())
}
func TestDeleteBuilder(t *testing.T) {
	test.CmdValidator(
		t,
		DeleteBuilder(),
		0,
		[]string{flag.TypeFlag, flag.Force, flag.EnableWatch, flag.WatchTimeout, flag.ProjectID},
	)
}
