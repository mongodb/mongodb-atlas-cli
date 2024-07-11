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
)

func TestSetupOpts_PostRun(t *testing.T) {
	ctrl := gomock.NewController(t)
	deploymentTest := fixture.NewMockLocalDeploymentOpts(ctrl, deploymentName)
	buf := new(bytes.Buffer)

	opts := &SetupOpts{
		DeploymentOpts: *deploymentTest.Opts,
		OutputOpts: cli.OutputOpts{
			OutWriter: buf,
		},
	}

	deploymentTest.MockDeploymentTelemetry.
		EXPECT().
		AppendDeploymentType().
		Times(1)

	opts.PostRun()
}

// Happy path, nothing is downloaded yet. No containers exist.
func TestSetupOpts_LocalDev_HappyPathClean(t *testing.T) {
	ctrl := gomock.NewController(t)
	ctx := context.Background()
	deploymentTest := fixture.NewMockLocalDeploymentOpts(ctrl, deploymentName)
	buf := new(bytes.Buffer)

	opts := &SetupOpts{
		DeploymentOpts: *deploymentTest.Opts,
		OutputOpts: cli.OutputOpts{
			OutWriter: buf,
		},
		force: true,
	}

	const dockerImageName = "docker.io/mongodb/mongodb-atlas-local:7.0"

	// Container engine is fine
	deploymentTest.MockContainerEngine.EXPECT().Ready().Return(nil).Times(1)

	// Image does not exist
	deploymentTest.MockContainerEngine.EXPECT().ImageList(ctx, dockerImageName).Return([]container.Image{}, nil).AnyTimes()

	// Image gets pulled
	deploymentTest.MockContainerEngine.EXPECT().ImagePull(ctx, dockerImageName).Return(nil).Times(1)

	// No local dev container exists yet
	deploymentTest.MockContainerEngine.EXPECT().ContainerList(ctx, "mongodb-atlas-local=container").Return([]container.Container{}, nil).Times(1)

	// Container run succeeds
	deploymentTest.MockContainerEngine.EXPECT().ContainerRun(ctx, gomock.Any(), gomock.Any()).Return(deploymentName, nil).Times(1)

	// Image contains a health check
	deploymentTest.MockContainerEngine.EXPECT().ImageHealthCheck(ctx, dockerImageName).Return(&container.ImageHealthCheck{
		Test: []string{"/bin/some-path"},
	}, nil).Times(1)

	// Container is healthy
	deploymentTest.MockContainerEngine.EXPECT().ContainerHealthStatus(ctx, deploymentName).Return(container.DockerHealthcheckStatusHealthy, nil).Times(1)

	// Verify
	if err := opts.Run(ctx); err != nil {
		t.Fatalf("Run() unexpected error: %v", err)
	}
}

// Happy path, image is already downloaded. No containers exist.
func TestSetupOpts_LocalDev_HappyPathImageDownloaded(t *testing.T) {
	ctrl := gomock.NewController(t)
	ctx := context.Background()
	deploymentTest := fixture.NewMockLocalDeploymentOpts(ctrl, deploymentName)
	buf := new(bytes.Buffer)

	opts := &SetupOpts{
		DeploymentOpts: *deploymentTest.Opts,
		OutputOpts: cli.OutputOpts{
			OutWriter: buf,
		},
		force: true,
	}

	const dockerImageName = "docker.io/mongodb/mongodb-atlas-local:7.0"

	// Container engine is fine
	deploymentTest.MockContainerEngine.EXPECT().Ready().Return(nil).Times(1)

	// Image exists
	deploymentTest.MockContainerEngine.EXPECT().ImageList(ctx, dockerImageName).Return([]container.Image{{
		ID: dockerImageName,
	}}, nil).Times(1)

	// No local dev container exists yet
	deploymentTest.MockContainerEngine.EXPECT().ContainerList(ctx, "mongodb-atlas-local=container").Return([]container.Container{}, nil).Times(1)

	// Container run succeeds
	deploymentTest.MockContainerEngine.EXPECT().ContainerRun(ctx, gomock.Any(), gomock.Any()).Return(deploymentName, nil).Times(1)

	// Image contains a health check
	deploymentTest.MockContainerEngine.EXPECT().ImageHealthCheck(ctx, dockerImageName).Return(&container.ImageHealthCheck{
		Test: []string{"/bin/some-path"},
	}, nil).Times(1)

	// Container is healthy
	deploymentTest.MockContainerEngine.EXPECT().ContainerHealthStatus(ctx, deploymentName).Return(container.DockerHealthcheckStatusHealthy, nil).Times(1)

	// Verify
	if err := opts.Run(ctx); err != nil {
		t.Fatalf("Run() unexpected error: %v", err)
	}
}

// Happy path, image is already downloaded. Containers exist.
func TestSetupOpts_LocalDev_HappyPathEverythingAlreadyExists(t *testing.T) {
	ctrl := gomock.NewController(t)
	ctx := context.Background()
	deploymentTest := fixture.NewMockLocalDeploymentOpts(ctrl, deploymentName)
	buf := new(bytes.Buffer)

	opts := &SetupOpts{
		DeploymentOpts: *deploymentTest.Opts,
		OutputOpts: cli.OutputOpts{
			OutWriter: buf,
		},
		force: true,
	}

	const dockerImageName = "docker.io/mongodb/mongodb-atlas-local:7.0"

	// Container engine is fine
	deploymentTest.MockContainerEngine.EXPECT().Ready().Return(nil).Times(1)

	// Image exists
	deploymentTest.MockContainerEngine.EXPECT().ImageList(ctx, dockerImageName).Return([]container.Image{{
		ID: dockerImageName,
	}}, nil).Times(1)

	// No local dev container exists yet
	deploymentTest.MockContainerEngine.EXPECT().ContainerList(ctx, "mongodb-atlas-local=container").Return([]container.Container{
		{
			ID:    "random-container-id",
			Names: []string{deploymentName},
		},
	}, nil).Times(1)

	// Verify
	if err := opts.Run(ctx); err == nil {
		t.Fatal("Run() unexpected success, should fail")
	}
}

func TestSetupOpts_LocalDev_RemoveUnhealthyDeployment(t *testing.T) {
	ctrl := gomock.NewController(t)
	ctx := context.Background()
	deploymentTest := fixture.NewMockLocalDeploymentOpts(ctrl, deploymentName)
	buf := new(bytes.Buffer)

	opts := &SetupOpts{
		DeploymentOpts: *deploymentTest.Opts,
		OutputOpts: cli.OutputOpts{
			OutWriter: buf,
		},
		force: true,
	}

	const dockerImageName = "docker.io/mongodb/mongodb-atlas-local:7.0"

	// Container engine is fine
	deploymentTest.MockContainerEngine.EXPECT().Ready().Return(nil).Times(1)

	// Image exists
	deploymentTest.MockContainerEngine.EXPECT().ImageList(ctx, dockerImageName).Return([]container.Image{{
		ID: dockerImageName,
	}}, nil).Times(1)

	// No local dev container exists yet
	deploymentTest.MockContainerEngine.EXPECT().ContainerList(ctx, "mongodb-atlas-local=container").Return([]container.Container{}, nil).Times(1)

	// Container run succeeds
	deploymentTest.MockContainerEngine.EXPECT().ContainerRun(ctx, gomock.Any(), gomock.Any()).Return(deploymentName, nil).Times(1)

	// Image contains a health check
	deploymentTest.MockContainerEngine.EXPECT().ImageHealthCheck(ctx, dockerImageName).Return(&container.ImageHealthCheck{
		Test: []string{"/bin/some-path"},
	}, nil).Times(1)

	// Container is unhealthy
	deploymentTest.MockContainerEngine.EXPECT().ContainerHealthStatus(ctx, deploymentName).Return(container.DockerHealthcheckStatusUnhealthy, nil).Times(1)

	// Container is removed
	deploymentTest.MockContainerEngine.EXPECT().ContainerRm(ctx, deploymentName).Return(nil).Times(1)

	// Verify
	if err := opts.Run(ctx); err == nil {
		t.Fatal("Run() unexpected success, should fail")
	}
}
