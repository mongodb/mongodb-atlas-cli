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

const (
	deploymentName = "localTest2"
	projectID      = "64f670f0bf789926667dad1a"
)

func TestPause_RunLocal(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockStore := mocks.NewMockClusterPauser(ctrl)
	mockCredentialsGetter := mocks.NewMockCredentialsGetter(ctrl)
	mockProfileReader := mocks.NewMockProfileReader(ctrl)
	mockPodman := mocks.NewMockClient(ctrl)
	ctx := context.Background()

	expectedLocalDeployments := []*podman.Container{
		{
			Names:  []string{"localTest2"},
			State:  "running",
			Labels: map[string]string{"version": "6.0.9"},
			ID:     deploymentName,
		},
		{
			Names:  []string{"localTest1"},
			State:  "running",
			Labels: map[string]string{"version": "7.0.0"},
			ID:     deploymentName,
		},
	}

	buf := new(bytes.Buffer)
	pauseOpts := &PauseOpts{
		store:  mockStore,
		config: mockProfileReader,
		DeploymentOpts: options.DeploymentOpts{
			PodmanClient:   mockPodman,
			CredStore:      mockCredentialsGetter,
			DeploymentName: deploymentName,
			DeploymentType: "LOCAL",
		},
		GlobalOpts: cli.GlobalOpts{
			ProjectID: projectID,
		},
		OutputOpts: cli.OutputOpts{
			Template:  pauseTemplate,
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
		Exec(ctx, "-d", pauseOpts.LocalMongodHostname(), "mongod", "--shutdown").
		Return(nil).
		Times(1)

	mockPodman.
		EXPECT().
		StopContainers(ctx, pauseOpts.LocalMongotHostname()).
		Return(nil, nil).
		Times(1)

	if err := pauseOpts.Run(ctx); err != nil {
		t.Fatalf("Run() unexpected error: %v", err)
	}

	assert.Equal(t, fmt.Sprintf("Pausing deployment '%s'.\n", deploymentName), buf.String())
	t.Log(buf.String())
}

func TestPause_RunAtlas(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockStore := mocks.NewMockClusterPauser(ctrl)
	mockCredentialsGetter := mocks.NewMockCredentialsGetter(ctrl)
	mockProfileReader := mocks.NewMockProfileReader(ctrl)
	mockPodman := mocks.NewMockClient(ctrl)
	ctx := context.Background()

	expectedLocalDeployments := []*podman.Container{
		{
			Names:  []string{"localTest2"},
			State:  "running",
			Labels: map[string]string{"version": "6.0.9"},
			ID:     deploymentName,
		},
		{
			Names:  []string{"localTest1"},
			State:  "running",
			Labels: map[string]string{"version": "7.0.0"},
			ID:     deploymentName,
		},
	}

	buf := new(bytes.Buffer)
	listOpts := &PauseOpts{
		store:  mockStore,
		config: mockProfileReader,
		DeploymentOpts: options.DeploymentOpts{
			PodmanClient:   mockPodman,
			CredStore:      mockCredentialsGetter,
			DeploymentName: deploymentName,
			DeploymentType: "ATLAS",
		},
		GlobalOpts: cli.GlobalOpts{
			ProjectID: projectID,
		},
		OutputOpts: cli.OutputOpts{
			Template:  pauseTemplate,
			OutWriter: buf,
		},
	}

	mockPodman.
		EXPECT().
		Ready(ctx).
		Return(nil).
		Times(0)

	mockPodman.
		EXPECT().
		ListContainers(ctx, options.MongotHostnamePrefix).
		Return(expectedLocalDeployments, options.ErrDeploymentNotFound).
		Times(0)

	mockCredentialsGetter.
		EXPECT().
		AuthType().
		Return(config.OAuth).
		Times(1)

	mockStore.
		EXPECT().
		PauseCluster(projectID, deploymentName).
		Return(
			&admin.AdvancedClusterDescription{
				Name: pointer.GetStringPointerIfNotEmpty(deploymentName),
			}, nil).
		Times(1)

	if err := listOpts.Run(ctx); err != nil {
		t.Fatalf("Run() unexpected error: %v", err)
	}

	assert.Equal(t, fmt.Sprintf("Pausing deployment '%s'.\n", deploymentName), buf.String())
	t.Log(buf.String())
}

func TestPauseBuilder(t *testing.T) {
	test.CmdValidator(
		t,
		PauseBuilder(),
		0,
		[]string{flag.ProjectID, flag.TypeFlag},
	)
}
