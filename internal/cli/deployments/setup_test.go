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
	"errors"
	"testing"

	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli/deployments/test/fixture"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/container"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/log"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
)

const dockerImageName = "docker.io/mongodb/mongodb-atlas-local:8"

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
	deploymentTest.MockDeploymentTelemetry.
		EXPECT().
		AppendDeploymentUUID().
		Times(1)

	opts.PostRun()
}

// Happy path. No containers exist.
func TestSetupOpts_LocalDev_HappyPathClean(t *testing.T) {
	ctrl := gomock.NewController(t)
	ctx := t.Context()
	deploymentTest := fixture.NewMockLocalDeploymentOpts(ctrl, deploymentName)
	buf := new(bytes.Buffer)

	opts := &SetupOpts{
		DeploymentOpts: *deploymentTest.Opts,
		OutputOpts: cli.OutputOpts{
			OutWriter: buf,
		},
		force: true,
	}

	// Container engine is fine
	deploymentTest.MockContainerEngine.EXPECT().Ready().Return(nil).Times(1)

	// Verify version should always succeed
	deploymentTest.MockContainerEngine.EXPECT().VerifyVersion(ctx).Return(nil).Times(1)

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

	// Because no port is specified docker inspect will happen
	deploymentTest.MockContainerEngine.EXPECT().ContainerInspect(ctx, opts.LocalMongodHostname()).Return([]*container.InspectData{{
		Config: &container.InspectDataConfig{
			Labels: map[string]string{
				"version": "7.0.12",
			},
		},
		NetworkSettings: &container.NetworkSettings{
			Ports: map[string][]container.InspectDataHostPort{
				"27017/tcp": {
					container.InspectDataHostPort{
						HostIP:   "127.0.0.1",
						HostPort: "12345",
					},
				},
			},
		},
	}}, nil).Times(1)

	// Verify
	if err := opts.Run(ctx); err != nil {
		t.Fatalf("Run() unexpected error: %v", err)
	}
}

// Happy path. Image exists, image update fails. No containers exist.
func TestSetupOpts_LocalDev_HappyPathOfflinePull(t *testing.T) {
	ctrl := gomock.NewController(t)
	ctx := t.Context()
	deploymentTest := fixture.NewMockLocalDeploymentOpts(ctrl, deploymentName)
	buf := new(bytes.Buffer)

	opts := &SetupOpts{
		DeploymentOpts: *deploymentTest.Opts,
		OutputOpts: cli.OutputOpts{
			OutWriter: buf,
		},
		force: true,
	}

	// Container engine is fine
	deploymentTest.MockContainerEngine.EXPECT().Ready().Return(nil).Times(1)

	// Verify version should always succeed
	deploymentTest.MockContainerEngine.EXPECT().VerifyVersion(ctx).Return(nil).Times(1)

	// Image gets pulled
	deploymentTest.MockContainerEngine.EXPECT().ImagePull(ctx, dockerImageName).Return(errors.New("image pull failed")).Times(1)

	// The image was downloaded before
	deploymentTest.MockContainerEngine.EXPECT().ImageList(ctx, dockerImageName).Return([]container.Image{{ID: dockerImageName}}, nil).Times(1)

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

	// Because no port is specified docker inspect will happen
	deploymentTest.MockContainerEngine.EXPECT().ContainerInspect(ctx, opts.LocalMongodHostname()).Return([]*container.InspectData{{
		Config: &container.InspectDataConfig{
			Labels: map[string]string{
				"version": "7.0.12",
			},
		},
		NetworkSettings: &container.NetworkSettings{
			Ports: map[string][]container.InspectDataHostPort{
				"27017/tcp": {
					container.InspectDataHostPort{
						HostIP:   "127.0.0.1",
						HostPort: "12345",
					},
				},
			},
		},
	}}, nil).Times(1)

	// Verify
	if err := opts.Run(ctx); err != nil {
		t.Fatalf("Run() unexpected error: %v", err)
	}
}

// Unhappy path. Image does not exist, image update fails. No containers exist.
func TestSetupOpts_LocalDev_UnhappyPathOfflinePull(t *testing.T) {
	ctrl := gomock.NewController(t)
	ctx := t.Context()
	deploymentTest := fixture.NewMockLocalDeploymentOpts(ctrl, deploymentName)
	buf := new(bytes.Buffer)

	opts := &SetupOpts{
		DeploymentOpts: *deploymentTest.Opts,
		OutputOpts: cli.OutputOpts{
			OutWriter: buf,
		},
		force: true,
	}

	// Container engine is fine
	deploymentTest.MockContainerEngine.EXPECT().Ready().Return(nil).Times(1)

	// Verify version should always succeed
	deploymentTest.MockContainerEngine.EXPECT().VerifyVersion(ctx).Return(nil).Times(1)

	// Image gets pulled
	deploymentTest.MockContainerEngine.EXPECT().ImagePull(ctx, dockerImageName).Return(errors.New("image pull failed")).Times(1)

	// No local dev container exists yet
	deploymentTest.MockContainerEngine.EXPECT().ContainerList(ctx, "mongodb-atlas-local=container").Return([]container.Container{}, nil).Times(1)

	// The image was downloaded before
	deploymentTest.MockContainerEngine.EXPECT().ImageList(ctx, dockerImageName).Return([]container.Image{}, nil).Times(1)

	// Container is removed
	deploymentTest.MockContainerEngine.EXPECT().ContainerRm(ctx, deploymentName).Return(nil).Times(1)

	// Verify
	if err := opts.Run(ctx); err == nil {
		t.Fatal("Run() unexpected success, should fail")
	}
}

// Happy path, image is already downloaded. Containers exist.
func TestSetupOpts_LocalDev_HappyPathEverythingAlreadyExists(t *testing.T) {
	ctrl := gomock.NewController(t)
	ctx := t.Context()
	deploymentTest := fixture.NewMockLocalDeploymentOpts(ctrl, deploymentName)
	buf := new(bytes.Buffer)

	opts := &SetupOpts{
		DeploymentOpts: *deploymentTest.Opts,
		OutputOpts: cli.OutputOpts{
			OutWriter: buf,
		},
		force: true,
	}

	// Container engine is fine
	deploymentTest.MockContainerEngine.EXPECT().Ready().Return(nil).Times(1)

	// Verify version should always succeed
	deploymentTest.MockContainerEngine.EXPECT().VerifyVersion(ctx).Return(nil).Times(1)

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
	ctx := t.Context()
	deploymentTest := fixture.NewMockLocalDeploymentOpts(ctrl, deploymentName)
	buf := new(bytes.Buffer)
	log.SetLevel(log.DebugLevel)

	opts := &SetupOpts{
		DeploymentOpts: *deploymentTest.Opts,
		OutputOpts: cli.OutputOpts{
			OutWriter: buf,
		},
		force: true,
	}

	// Container engine is fine
	deploymentTest.MockContainerEngine.EXPECT().Ready().Return(nil).Times(1)

	// Verify version should always succeed
	deploymentTest.MockContainerEngine.EXPECT().VerifyVersion(ctx).Return(nil).Times(1)

	// Image gets pulled (updated)
	deploymentTest.MockContainerEngine.EXPECT().ImagePull(ctx, dockerImageName).Return(nil).Times(1)

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

	// Docker logs
	deploymentTest.MockContainerEngine.EXPECT().ContainerLogs(ctx, deploymentName).Return([]string{"something", "went", "wrong"}, nil).Times(1)

	// Container is removed
	deploymentTest.MockContainerEngine.EXPECT().ContainerRm(ctx, deploymentName).Return(nil).Times(1)

	// Verify
	if err := opts.Run(ctx); err == nil {
		t.Fatal("Run() unexpected success, should fail")
	}
}

func TestValidateFlags_mdbVersions(t *testing.T) {
	testCases := []struct {
		name          string
		version       string
		expectedError error
	}{
		{name: "mdb70", version: mdb70, expectedError: nil},
		{name: "mdb80", version: mdb80, expectedError: nil},
		{name: "mdb8", version: mdb8, expectedError: nil},
		{name: "mdb7", version: mdb7, expectedError: nil},
		{name: "mdb82", version: "8.2", expectedError: errInvalidMongoDBVersion},
		{name: "invalid", version: "9.0", expectedError: errInvalidMongoDBVersion},
	}
	for _, testCase := range testCases {
		opts := &SetupOpts{}
		opts.MdbVersion = testCase.version
		err := opts.validateFlags()
		if testCase.expectedError != nil {
			assert.ErrorIs(t, err, testCase.expectedError)
		} else {
			assert.NoError(t, err)
		}
	}
}

func TestSetupOpts_MongodDockerImageName(t *testing.T) {
	testCases := []struct {
		name          string
		version       string
		expectedImage string
	}{
		{name: "mdb70", version: mdb70, expectedImage: "docker.io/mongodb/mongodb-atlas-local:7"},
		{name: "mdb80", version: mdb80, expectedImage: "docker.io/mongodb/mongodb-atlas-local:8"},
		{name: "mdb8", version: mdb8, expectedImage: "docker.io/mongodb/mongodb-atlas-local:8"},
		{name: "mdb7", version: mdb7, expectedImage: "docker.io/mongodb/mongodb-atlas-local:7"},
	}
	for _, testCase := range testCases {
		opts := &SetupOpts{}
		opts.MdbVersion = testCase.version
		assert.Equal(t, testCase.expectedImage, opts.MongodDockerImageName())
	}
}

func TestSetupOpts_isDiskSpaceError(t *testing.T) {
	testCases := []struct {
		name        string
		err         error
		expectedRes bool
	}{
		{
			name:        "no error",
			err:         nil,
			expectedRes: false,
		},
		{
			name:        "disk space error",
			err:         errors.New("docker: Error response from daemon: write /var/lib/docker/tmp/docker-builder123/Dockerfile: no space left on device"),
			expectedRes: true,
		},
		{
			name:        "non-disk space error",
			err:         errors.New("network timeout"),
			expectedRes: false,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			opts := &SetupOpts{}
			result := opts.isDiskSpaceError(testCase.err)
			assert.Equal(t, testCase.expectedRes, result, "Unexpected result for isDiskSpaceError")
		})
	}
}

func TestSetupOpts_downloadImage_HappyPath(t *testing.T) {
	ctrl := gomock.NewController(t)
	ctx := t.Context()
	deploymentTest := fixture.NewMockLocalDeploymentOpts(ctrl, deploymentName)

	opts := &SetupOpts{
		DeploymentOpts: *deploymentTest.Opts,
	}
	opts.MdbVersion = "8"

	deploymentTest.MockContainerEngine.EXPECT().
		ImagePull(ctx, dockerImageName).
		Return(nil).
		Times(1)

	err := opts.downloadImage(ctx, 1)
	require.NoError(t, err, "Expected no error on successful image download")
}

func TestSetupOpts_downloadImage_ErrorPulling_ImageExists(t *testing.T) {
	ctrl := gomock.NewController(t)
	ctx := t.Context()
	deploymentTest := fixture.NewMockLocalDeploymentOpts(ctrl, deploymentName)

	opts := &SetupOpts{
		DeploymentOpts: *deploymentTest.Opts,
	}
	opts.MdbVersion = "8"

	// Mock image pull failure with network error
	deploymentTest.MockContainerEngine.EXPECT().
		ImagePull(ctx, dockerImageName).
		Return(errors.New("network timeout")).
		Times(1)

	// Mock existing image found
	deploymentTest.MockContainerEngine.EXPECT().
		ImageList(ctx, dockerImageName).
		Return([]container.Image{{ID: dockerImageName}}, nil).
		Times(1)

	// Should return nil when image exists locally
	err := opts.downloadImage(ctx, 1)
	require.NoError(t, err, "Expected no error when image exists locally")
}

func TestSetupOpts_downloadImage_FailedToDownloadImageError(t *testing.T) {
	ctrl := gomock.NewController(t)
	ctx := t.Context()
	deploymentTest := fixture.NewMockLocalDeploymentOpts(ctrl, deploymentName)

	opts := &SetupOpts{
		DeploymentOpts: *deploymentTest.Opts,
	}
	opts.MdbVersion = "8"

	// Mock image pull failure with network error
	deploymentTest.MockContainerEngine.EXPECT().
		ImagePull(ctx, dockerImageName).
		Return(errors.New("network timeout")).
		Times(1)

	// Mock no existing image found
	deploymentTest.MockContainerEngine.EXPECT().
		ImageList(ctx, dockerImageName).
		Return([]container.Image{}, nil).
		Times(1)

	// Should return errFailedToDownloadImage when no image exists
	err := opts.downloadImage(ctx, 1)
	require.ErrorIs(t, err, errFailedToDownloadImage, "Expected errFailedToDownloadImage when no local image exists")
}

func TestSetupOpts_downloadImage_DiskSpaceError(t *testing.T) {
	ctrl := gomock.NewController(t)
	ctx := t.Context()
	deploymentTest := fixture.NewMockLocalDeploymentOpts(ctrl, deploymentName)

	opts := &SetupOpts{
		DeploymentOpts: *deploymentTest.Opts,
	}
	opts.MdbVersion = "8"

	// Mock image pull failure with disk space error
	deploymentTest.MockContainerEngine.EXPECT().
		ImagePull(ctx, dockerImageName).
		Return(errors.New("no space left on device")).
		Times(1)

	// Should return errInsufficientDiskSpace directly
	err := opts.downloadImage(ctx, 1)
	require.ErrorIs(t, err, errInsufficientDiskSpace, "Expected errInsufficientDiskSpace for disk space error")
}

func TestSetupOpts_LocalDev_DiskSpaceError_During_ImagePull(t *testing.T) {
	ctrl := gomock.NewController(t)
	ctx := t.Context()
	deploymentTest := fixture.NewMockLocalDeploymentOpts(ctrl, deploymentName)
	buf := new(bytes.Buffer)

	opts := &SetupOpts{
		DeploymentOpts: *deploymentTest.Opts,
		OutputOpts: cli.OutputOpts{
			OutWriter: buf,
		},
		force: true,
	}

	// Container engine is fine
	deploymentTest.MockContainerEngine.EXPECT().Ready().Return(nil).Times(1)

	// Verify version should always succeed
	deploymentTest.MockContainerEngine.EXPECT().VerifyVersion(ctx).Return(nil).Times(1)

	// No local dev container exists yet
	deploymentTest.MockContainerEngine.EXPECT().ContainerList(ctx, "mongodb-atlas-local=container").Return([]container.Container{}, nil).Times(1)

	// Image pull fails with disk space error
	deploymentTest.MockContainerEngine.EXPECT().ImagePull(ctx, dockerImageName).Return(errors.New("write /var/lib/docker: no space left on device")).Times(1)

	// Container is removed on failure
	deploymentTest.MockContainerEngine.EXPECT().ContainerRm(ctx, deploymentName).Return(nil).Times(1)

	// Verify that the run fails with disk space error
	err := opts.Run(ctx)
	require.Error(t, err, "Expected error due to disk space issue")
	assert.ErrorIs(t, err, errInsufficientDiskSpace, "Expected errInsufficientDiskSpace")
}

func TestSetupOpts_LocalDev_DiskSpaceError_During_ContainerRun(t *testing.T) {
	ctrl := gomock.NewController(t)
	ctx := t.Context()
	deploymentTest := fixture.NewMockLocalDeploymentOpts(ctrl, deploymentName)
	buf := new(bytes.Buffer)

	opts := &SetupOpts{
		DeploymentOpts: *deploymentTest.Opts,
		OutputOpts: cli.OutputOpts{
			OutWriter: buf,
		},
		force: true,
	}

	// Container engine is fine
	deploymentTest.MockContainerEngine.EXPECT().Ready().Return(nil).Times(1)

	// Verify version should always succeed
	deploymentTest.MockContainerEngine.EXPECT().VerifyVersion(ctx).Return(nil).Times(1)

	// Image gets pulled successfully
	deploymentTest.MockContainerEngine.EXPECT().ImagePull(ctx, dockerImageName).Return(nil).Times(1)

	// No local dev container exists yet
	deploymentTest.MockContainerEngine.EXPECT().ContainerList(ctx, "mongodb-atlas-local=container").Return([]container.Container{}, nil).Times(1)

	// Image health check succeeds
	deploymentTest.MockContainerEngine.EXPECT().ImageHealthCheck(ctx, dockerImageName).Return(&container.ImageHealthCheck{
		Test: []string{"/bin/some-path"},
	}, nil).Times(1)

	// Container run fails with disk space error
	deploymentTest.MockContainerEngine.EXPECT().ContainerRun(ctx, gomock.Any(), gomock.Any()).Return("", errors.New("docker: no space left on device")).Times(1)

	// Container is removed on failure
	deploymentTest.MockContainerEngine.EXPECT().ContainerRm(ctx, deploymentName).Return(nil).Times(1)

	// Verify that the run fails with disk space error
	err := opts.Run(ctx)
	require.Error(t, err, "Expected error due to disk space issue during container run")
	// The error should be wrapped but contain disk space information
	assert.Contains(t, err.Error(), "insufficient disk space", "Error message should mention disk space")
}
