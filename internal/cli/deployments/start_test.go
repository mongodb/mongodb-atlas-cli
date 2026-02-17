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
	"fmt"
	"testing"

	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli/deployments/options"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli/deployments/test/fixture"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/pointer"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	atlasClustersPinned "go.mongodb.org/atlas-sdk/v20240530005/admin"
	atlasv2 "go.mongodb.org/atlas-sdk/v20250312014/admin"
	"go.uber.org/mock/gomock"
)

func TestStart_RunLocal_PausedContainers(t *testing.T) {
	ctrl := gomock.NewController(t)
	ctx := t.Context()

	deploymentsTest := fixture.NewMockLocalDeploymentOpts(ctrl, deploymentName)

	buf := new(bytes.Buffer)
	startOpts := &StartOpts{
		DeploymentOpts: *deploymentsTest.Opts,
		ProjectOpts: cli.ProjectOpts{
			ProjectID: projectID,
		},
		OutputOpts: cli.OutputOpts{
			Template:  startTemplate,
			OutWriter: buf,
		},
	}

	expected := deploymentsTest.MockContainerWithState("paused")
	deploymentsTest.LocalMockFlowWithMockContainer(ctx, expected)

	deploymentsTest.MockContainerEngine.
		EXPECT().
		ContainerUnpause(ctx, startOpts.LocalMongodHostname()).
		Return(nil).
		Times(1)

	require.NoError(t, startOpts.Run(ctx))
	assert.Equal(t, fmt.Sprintf("\nStarting deployment '%s'.\n", deploymentName), buf.String())
	t.Log(buf.String())
}

func TestStart_RunLocal_StoppedContainers(t *testing.T) {
	ctrl := gomock.NewController(t)
	ctx := t.Context()

	deploymentsTest := fixture.NewMockLocalDeploymentOpts(ctrl, deploymentName)

	buf := new(bytes.Buffer)
	startOpts := &StartOpts{
		DeploymentOpts: *deploymentsTest.Opts,
		ProjectOpts: cli.ProjectOpts{
			ProjectID: projectID,
		},
		OutputOpts: cli.OutputOpts{
			Template:  startTemplate,
			OutWriter: buf,
		},
	}

	expected := deploymentsTest.MockContainerWithState("exited")
	deploymentsTest.LocalMockFlowWithMockContainer(ctx, expected)

	deploymentsTest.MockContainerEngine.
		EXPECT().
		ContainerStart(ctx, startOpts.LocalMongodHostname()).
		Return(nil).
		Times(1)

	require.NoError(t, startOpts.Run(ctx))
	assert.Equal(t, fmt.Sprintf("\nStarting deployment '%s'.\n", deploymentName), buf.String())
	t.Log(buf.String())
}

func TestStart_RunAtlas_clusterWideScaling(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockStore := NewMockClusterStarter(ctrl)
	ctx := t.Context()
	const deploymentName = "atlasCluster1"

	deploymentsTest := fixture.NewMockAtlasDeploymentOpts(ctrl, deploymentName)

	buf := new(bytes.Buffer)
	opts := &StartOpts{
		store:          mockStore,
		DeploymentOpts: *deploymentsTest.Opts,
		ProjectOpts: cli.ProjectOpts{
			ProjectID: projectID,
		},
		OutputOpts: cli.OutputOpts{
			Template:  startTemplate,
			OutWriter: buf,
		},
	}

	deploymentsTest.CommonAtlasMocksWithState(projectID, "STOPPED")

	mockStore.
		EXPECT().
		GetClusterAutoScalingConfig(projectID, deploymentName).
		Return(
			&atlasv2.ClusterDescriptionAutoScalingModeConfiguration{
				AutoScalingMode: pointer.Get(string(options.ClusterWideScalingResponse)),
			}, nil).
		Times(1)

	deploymentsTest.
		MockDeploymentTelemetry.
		EXPECT().
		AppendClusterWideScalingMode().
		Times(1)

	mockStore.
		EXPECT().
		StartCluster(projectID, deploymentName).
		Return(
			&atlasClustersPinned.AdvancedClusterDescription{
				Name: pointer.Get(deploymentName),
			}, nil).
		Times(1)

	require.NoError(t, opts.Run(ctx))
	assert.Equal(t, fmt.Sprintf("\nStarting deployment '%s'.\n", deploymentName), buf.String())
	t.Log(buf.String())
}

func TestStart_RunAtlas_independentShardScaling(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockStore := NewMockClusterStarter(ctrl)
	ctx := t.Context()
	const deploymentName = "atlasCluster1"

	deploymentsTest := fixture.NewMockAtlasDeploymentOpts(ctrl, deploymentName)

	buf := new(bytes.Buffer)
	opts := &StartOpts{
		store:          mockStore,
		DeploymentOpts: *deploymentsTest.Opts,
		ProjectOpts: cli.ProjectOpts{
			ProjectID: projectID,
		},
		OutputOpts: cli.OutputOpts{
			Template:  startTemplate,
			OutWriter: buf,
		},
	}

	deploymentsTest.CommonAtlasMocksWithState(projectID, "STOPPED")

	mockStore.
		EXPECT().
		GetClusterAutoScalingConfig(projectID, deploymentName).
		Return(
			&atlasv2.ClusterDescriptionAutoScalingModeConfiguration{
				AutoScalingMode: pointer.Get(string(options.IndependentShardScalingResponse)),
			}, nil).
		Times(1)

	deploymentsTest.
		MockDeploymentTelemetry.
		EXPECT().
		AppendIndependentShardScalingMode().
		Times(1)

	mockStore.
		EXPECT().
		StartClusterLatest(projectID, deploymentName).
		Return(
			&atlasv2.ClusterDescription20240805{
				Name: pointer.Get(deploymentName),
			}, nil).
		Times(1)

	require.NoError(t, opts.Run(ctx))
	assert.Equal(t, fmt.Sprintf("\nStarting deployment '%s'.\n", deploymentName), buf.String())
	t.Log(buf.String())
}

func TestStartOpts_PostRun(t *testing.T) {
	ctrl := gomock.NewController(t)
	deploymentsTest := fixture.NewMockLocalDeploymentOpts(ctrl, "localDeployment")
	buf := new(bytes.Buffer)

	opts := &StartOpts{
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

	deploymentsTest.
		MockDeploymentTelemetry.
		EXPECT().
		AppendDeploymentUUID().
		Times(1)

	deploymentsTest.MockContainerEngine.EXPECT().Ready().Return(nil).Times(1)

	require.NoError(t, opts.PostRun())
}
