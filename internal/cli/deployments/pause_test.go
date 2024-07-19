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
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli/deployments/test/fixture"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/flag"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/mocks"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/pointer"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/test"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/atlas-sdk/v20240530002/admin"
)

const (
	deploymentName = "localTest2"
	projectID      = "64f670f0bf789926667dad1a"
)

func TestPause_RunLocal(t *testing.T) {
	ctrl := gomock.NewController(t)
	deploymentTest := fixture.NewMockLocalDeploymentOpts(ctrl, deploymentName)
	ctx := context.Background()

	buf := new(bytes.Buffer)
	pauseOpts := &PauseOpts{
		DeploymentOpts: *deploymentTest.Opts,
		OutputOpts: cli.OutputOpts{
			Template:  pauseTemplate,
			OutWriter: buf,
		},
	}

	deploymentTest.LocalMockFlow(ctx)

	deploymentTest.MockContainerEngine.
		EXPECT().
		ContainerStop(ctx, pauseOpts.LocalMongodHostname()).
		Return(nil).
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
	ctx := context.Background()

	deploymentTest := fixture.NewMockAtlasDeploymentOpts(ctrl, deploymentName)

	buf := new(bytes.Buffer)
	listOpts := &PauseOpts{
		store:          mockStore,
		DeploymentOpts: *deploymentTest.Opts,
		GlobalOpts: cli.GlobalOpts{
			ProjectID: projectID,
		},
		OutputOpts: cli.OutputOpts{
			Template:  pauseTemplate,
			OutWriter: buf,
		},
	}

	deploymentTest.CommonAtlasMocks(projectID)

	mockStore.
		EXPECT().
		PauseCluster(projectID, deploymentName).
		Return(
			&admin.AdvancedClusterDescription{
				Name: pointer.Get(deploymentName),
			}, nil).
		Times(1)

	if err := listOpts.Run(ctx); err != nil {
		t.Fatalf("Run() unexpected error: %v", err)
	}

	assert.Equal(t, fmt.Sprintf("Pausing deployment '%s'.\n", deploymentName), buf.String())
	t.Log(buf.String())
}

func TestPauseOpts_PostRun(t *testing.T) {
	ctrl := gomock.NewController(t)
	deploymentTest := fixture.NewMockLocalDeploymentOpts(ctrl, deploymentName)
	buf := new(bytes.Buffer)

	opts := &PauseOpts{
		DeploymentOpts: *deploymentTest.Opts,
		OutputOpts: cli.OutputOpts{
			OutWriter: buf,
		},
	}

	deploymentTest.MockDeploymentTelemetry.
		EXPECT().
		AppendDeploymentType().
		Times(1)

	deploymentTest.MockContainerEngine.EXPECT().Ready().Return(nil).Times(1)

	if err := opts.PostRun(); err != nil {
		t.Fatalf("PostRun() unexpected error: %v", err)
	}
}

func TestPauseBuilder(t *testing.T) {
	test.CmdValidator(
		t,
		PauseBuilder(),
		0,
		[]string{flag.ProjectID, flag.TypeFlag},
	)
}
