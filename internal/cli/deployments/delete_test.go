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
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/flag"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/mocks"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/test"
)

func TestDelete_Run_Atlas(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockAtlasStore := mocks.NewMockClusterDeleter(ctrl)
	ctx := context.Background()

	deploymentsTest := fixture.NewMockAtlasDeploymentOpts(ctrl, "atlasDeployment")

	buf := new(bytes.Buffer)
	opts := &DeleteOpts{
		atlasStore:     mockAtlasStore,
		DeploymentOpts: *deploymentsTest.Opts,
		GlobalOpts: cli.GlobalOpts{
			ProjectID: "64f670f0bf789926667dad1a",
		},
		OutputOpts: cli.OutputOpts{
			OutWriter: buf,
		},
		DeleteOpts: cli.NewDeleteOpts(deleteSuccessMessage, deleteFailMessage),
	}
	opts.Confirm = true

	deploymentsTest.CommonAtlasMocks(opts.ProjectID)

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
	ctx := context.Background()
	buf := new(bytes.Buffer)

	deploymentsTest := fixture.NewMockLocalDeploymentOpts(ctrl, "testDeployment")

	opts := &DeleteOpts{
		DeploymentOpts: *deploymentsTest.Opts,
		GlobalOpts: cli.GlobalOpts{
			ProjectID: "64f670f0bf789926667dad1a",
		},
		OutputOpts: cli.OutputOpts{
			OutWriter: buf,
		},
		DeleteOpts: cli.NewDeleteOpts(deleteSuccessMessage, deleteFailMessage),
	}
	opts.Confirm = true

	deploymentsTest.LocalMockFlow(ctx)

	deploymentsTest.
		MockContainerEngine.
		EXPECT().
		ContainerRm(ctx, opts.DeploymentName).
		Return(nil).
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

func TestDeleteOpts_PostRun(t *testing.T) {
	ctrl := gomock.NewController(t)
	deploymentsTest := fixture.NewMockLocalDeploymentOpts(ctrl, "localDeployment")
	buf := new(bytes.Buffer)

	opts := &DeleteOpts{
		DeploymentOpts: *deploymentsTest.Opts,
		GlobalOpts: cli.GlobalOpts{
			ProjectID: "64f670f0bf789926667dad1a",
		},
		OutputOpts: cli.OutputOpts{
			OutWriter: buf,
		},
		DeleteOpts: cli.NewDeleteOpts(deleteSuccessMessage, deleteFailMessage),
	}

	deploymentsTest.MockDeploymentTelemetry.EXPECT().AppendDeploymentType().Times(1)

	if err := opts.PostRun(); err != nil {
		t.Fatalf("PostRun() unexpected error: %v", err)
	}
}
