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
	"errors"
	"testing"

	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli/deployments/options"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli/deployments/test/fixture"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/config"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/container"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/mocks"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/pointer"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/store"
	"github.com/stretchr/testify/assert"
	atlasv2 "go.mongodb.org/atlas-sdk/v20250312005/admin"
	"go.uber.org/mock/gomock"
)

func TestList_Run(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockStore := fixture.NewMockClusterLister(ctrl)
	mockCredentialsGetter := mocks.NewMockCredentialsGetter(ctrl)
	mockContainerEngine := mocks.NewMockEngine(ctrl)
	ctx := t.Context()

	cli.TokenRefreshed = true
	t.Cleanup(func() {
		cli.TokenRefreshed = false
	})

	expectedAtlasClusters := &atlasv2.PaginatedClusterDescription20240805{
		Results: &[]atlasv2.ClusterDescription20240805{
			{
				Name:           pointer.Get("atlasCluster2"),
				Id:             pointer.Get("123"),
				MongoDBVersion: pointer.Get("7.0.0"),
				StateName:      pointer.Get("IDLE"),
				Paused:         pointer.Get(false),
			},
			{
				Name:           pointer.Get("atlasCluster1"),
				Id:             pointer.Get("123"),
				MongoDBVersion: pointer.Get("7.0.0"),
				StateName:      pointer.Get("IDLE"),
				Paused:         pointer.Get(false),
			},
		},
	}

	expectedLocalDeployments := []container.Container{
		{
			Names:  []string{"localTest2"},
			State:  "running",
			Labels: map[string]string{"version": "6.0.9"},
		},
		{
			Names:  []string{"localTest1"},
			State:  "running",
			Labels: map[string]string{"version": "7.0.0"},
		},
	}

	buf := new(bytes.Buffer)
	listOpts := &ListOpts{
		DeploymentOpts: options.DeploymentOpts{
			ContainerEngine:       mockContainerEngine,
			CredStore:             mockCredentialsGetter,
			AtlasClusterListStore: mockStore,
		},
		ProjectOpts: cli.ProjectOpts{
			ProjectID: "64f670f0bf789926667dad1a",
		},
		OutputOpts: cli.OutputOpts{
			Template:  listTemplate,
			OutWriter: buf,
		},
	}

	mockStore.
		EXPECT().
		LatestProjectClusters(listOpts.ProjectID,
			&store.ListOptions{
				PageNum:      cli.DefaultPage,
				ItemsPerPage: options.MaxItemsPerPage,
			},
		).
		Return(expectedAtlasClusters, nil).
		Times(1)

	mockCredentialsGetter.
		EXPECT().
		AuthType().
		Return(config.OAuth).
		Times(2)

	mockContainerEngine.
		EXPECT().
		Ready().
		Return(nil).
		Times(1)

	// Verify version should always succeed
	mockContainerEngine.EXPECT().VerifyVersion(ctx).Return(nil).Times(1)

	mockContainerEngine.
		EXPECT().
		ContainerList(ctx, options.ContainerFilter).
		Return(expectedLocalDeployments, nil).
		Times(1)

	if err := listOpts.Run(ctx); err != nil {
		t.Fatalf("Run() unexpected error: %v", err)
	}

	assert.Equal(t, `NAME            TYPE    MDB VER   STATE
atlasCluster2   ATLAS   7.0.0     IDLE
atlasCluster1   ATLAS   7.0.0     IDLE
localTest1      LOCAL   7.0.0     IDLE
localTest2      LOCAL   6.0.9     IDLE
`, buf.String())
	t.Log(buf.String())
}

func TestList_Run_NoLocal(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockStore := fixture.NewMockClusterLister(ctrl)
	mockCredentialsGetter := mocks.NewMockCredentialsGetter(ctrl)

	mockContainerEngine := mocks.NewMockEngine(ctrl)
	ctx := t.Context()

	cli.TokenRefreshed = true
	t.Cleanup(func() {
		cli.TokenRefreshed = false
	})

	expectedAtlasClusters := &atlasv2.PaginatedClusterDescription20240805{
		Results: &[]atlasv2.ClusterDescription20240805{
			{
				Name:           pointer.Get("atlasCluster2"),
				Id:             pointer.Get("123"),
				MongoDBVersion: pointer.Get("7.0.0"),
				StateName:      pointer.Get("IDLE"),
				Paused:         pointer.Get(false),
			},
			{
				Name:           pointer.Get("atlasCluster1"),
				Id:             pointer.Get("123"),
				MongoDBVersion: pointer.Get("7.0.0"),
				StateName:      pointer.Get("IDLE"),
				Paused:         pointer.Get(false),
			},
		},
	}

	buf := new(bytes.Buffer)
	listOpts := &ListOpts{
		DeploymentOpts: options.DeploymentOpts{
			ContainerEngine:       mockContainerEngine,
			CredStore:             mockCredentialsGetter,
			AtlasClusterListStore: mockStore,
		},
		ProjectOpts: cli.ProjectOpts{
			ProjectID: "64f670f0bf789926667dad1a",
		},
		OutputOpts: cli.OutputOpts{
			Template:  listTemplate,
			OutWriter: buf,
		},
	}

	mockStore.
		EXPECT().
		LatestProjectClusters(listOpts.ProjectID,
			&store.ListOptions{
				PageNum:      cli.DefaultPage,
				ItemsPerPage: options.MaxItemsPerPage,
			},
		).
		Return(expectedAtlasClusters, nil).
		Times(1)

	mockCredentialsGetter.
		EXPECT().
		AuthType().
		Return(config.OAuth).
		Times(2)

	mockContainerEngine.
		EXPECT().
		Ready().
		Return(nil).
		Times(1)

	// Verify version should always succeed
	mockContainerEngine.EXPECT().VerifyVersion(ctx).Return(nil).Times(1)

	mockContainerEngine.
		EXPECT().
		ContainerList(ctx, options.ContainerFilter).
		Return(nil, errors.New("this is an error")).
		Times(1)

	if err := listOpts.Run(ctx); err != nil {
		t.Fatalf("Run() unexpected error: %v", err)
	}

	assert.Equal(t, `NAME            TYPE    MDB VER   STATE
atlasCluster2   ATLAS   7.0.0     IDLE
atlasCluster1   ATLAS   7.0.0     IDLE
`, buf.String())
	t.Log(buf.String())
}

func TestList_Run_NoAtlas(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockStore := fixture.NewMockClusterLister(ctrl)
	mockCredentialsGetter := mocks.NewMockCredentialsGetter(ctrl)
	mockContainerEngine := mocks.NewMockEngine(ctrl)
	ctx := t.Context()

	cli.TokenRefreshed = true
	t.Cleanup(func() {
		cli.TokenRefreshed = false
	})

	expectedAtlasClusters := &atlasv2.PaginatedClusterDescription20240805{
		Results: &[]atlasv2.ClusterDescription20240805{
			{
				Name:           pointer.Get("atlasCluster2"),
				Id:             pointer.Get("123"),
				MongoDBVersion: pointer.Get("7.0.0"),
				StateName:      pointer.Get("IDLE"),
				Paused:         pointer.Get(false),
			},
			{
				Name:           pointer.Get("atlasCluster1"),
				Id:             pointer.Get("123"),
				MongoDBVersion: pointer.Get("7.0.0"),
				StateName:      pointer.Get("IDLE"),
				Paused:         pointer.Get(false),
			},
		},
	}

	buf := new(bytes.Buffer)
	listOpts := &ListOpts{
		DeploymentOpts: options.DeploymentOpts{
			ContainerEngine:       mockContainerEngine,
			CredStore:             mockCredentialsGetter,
			AtlasClusterListStore: mockStore,
		},
		ProjectOpts: cli.ProjectOpts{
			ProjectID: "64f670f0bf789926667dad1a",
		},
		OutputOpts: cli.OutputOpts{
			Template:  listTemplate,
			OutWriter: buf,
		},
	}

	mockStore.
		EXPECT().
		LatestProjectClusters(listOpts.ProjectID,
			&store.ListOptions{
				PageNum:      cli.DefaultPage,
				ItemsPerPage: options.MaxItemsPerPage,
			},
		).
		Return(expectedAtlasClusters, nil).
		Times(1)

	mockCredentialsGetter.
		EXPECT().
		AuthType().
		Return(config.OAuth).
		Times(2)

	mockContainerEngine.
		EXPECT().
		Ready().
		Return(nil).
		Times(1)

	// Verify version should always succeed
	mockContainerEngine.EXPECT().VerifyVersion(ctx).Return(nil).Times(1)

	mockContainerEngine.
		EXPECT().
		ContainerList(ctx, options.ContainerFilter).
		Return(nil, errors.New("new error test")).
		Times(1)

	if err := listOpts.Run(ctx); err != nil {
		t.Fatalf("Run() unexpected error: %v", err)
	}

	assert.Equal(t, `NAME            TYPE    MDB VER   STATE
atlasCluster2   ATLAS   7.0.0     IDLE
atlasCluster1   ATLAS   7.0.0     IDLE
`, buf.String())
	t.Log(buf.String())
}

func TestListOpts_PostRun(t *testing.T) {
	ctrl := gomock.NewController(t)
	buf := new(bytes.Buffer)

	mockStore := fixture.NewMockClusterLister(ctrl)
	mockCredentialsGetter := mocks.NewMockCredentialsGetter(ctrl)

	deploymentsTest := fixture.NewMockLocalDeploymentOpts(ctrl, "localDeployment")
	deploymentsTest.Opts.CredStore = mockCredentialsGetter
	deploymentsTest.Opts.AtlasClusterListStore = mockStore

	listOpts := &ListOpts{
		DeploymentOpts: *deploymentsTest.Opts,
		ProjectOpts: cli.ProjectOpts{
			ProjectID: "64f670f0bf789926667dad1a",
		},
		OutputOpts: cli.OutputOpts{
			Template:  listTemplate,
			OutWriter: buf,
		},
	}

	mockCredentialsGetter.
		EXPECT().
		AuthType().
		Return(config.OAuth).
		Times(1)

	deploymentsTest.MockDeploymentTelemetry.
		EXPECT().
		AppendDeploymentType().
		Times(1)

	deploymentsTest.MockContainerEngine.EXPECT().Ready().Return(nil).Times(1)

	if err := listOpts.PostRun(); err != nil {
		t.Fatalf("PostRun() unexpected error: %v", err)
	}
}
