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
	"github.com/mongodb/mongodb-atlas-cli/internal/cli"
	"github.com/mongodb/mongodb-atlas-cli/internal/cli/atlas/deployments/options"
	"github.com/mongodb/mongodb-atlas-cli/internal/config"
	"github.com/mongodb/mongodb-atlas-cli/internal/flag"
	"github.com/mongodb/mongodb-atlas-cli/internal/log"
	"github.com/mongodb/mongodb-atlas-cli/internal/mocks"
	"github.com/mongodb/mongodb-atlas-cli/internal/podman"
	"github.com/mongodb/mongodb-atlas-cli/internal/pointer"
	"github.com/mongodb/mongodb-atlas-cli/internal/test"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/atlas-sdk/v20230201008/admin"
	"go.mongodb.org/atlas/mongodbatlas"
)

func TestList_Run(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockStore := mocks.NewMockClusterLister(ctrl)
	mockCredentialsGetter := mocks.NewMockCredentialsGetter(ctrl)
	mockProfileReader := mocks.NewMockProfileReader(ctrl)
	mockPodman := mocks.NewMockClient(ctrl)
	ctx := context.Background()

	expectedAtlasClusters := &admin.PaginatedAdvancedClusterDescription{
		Results: []admin.AdvancedClusterDescription{
			{
				Name:           pointer.Get("atlasCluster2"),
				Id:             pointer.Get("123"),
				MongoDBVersion: pointer.Get("7.0.0"),
				StateName:      pointer.Get("IDLE"),
			},
			{
				Name:           pointer.Get("atlasCluster1"),
				Id:             pointer.Get("123"),
				MongoDBVersion: pointer.Get("7.0.0"),
				StateName:      pointer.Get("IDLE"),
			},
		},
	}

	expectedLocalDeployments := []*podman.Container{
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
		store:  mockStore,
		config: mockProfileReader,
		DeploymentOpts: options.DeploymentOpts{
			PodmanClient: mockPodman,
			CredStore:    mockCredentialsGetter,
		},
		GlobalOpts: cli.GlobalOpts{
			ProjectID: "64f670f0bf789926667dad1a",
		},
		OutputOpts: cli.OutputOpts{
			Template:  listTemplate,
			OutWriter: buf,
		},
	}

	mockStore.
		EXPECT().
		ProjectClusters(listOpts.ProjectID,
			&mongodbatlas.ListOptions{
				PageNum:      cli.DefaultPage,
				ItemsPerPage: MaxItemsPerPage,
			},
		).
		Return(expectedAtlasClusters, nil).
		Times(1)

	mockCredentialsGetter.
		EXPECT().
		AuthType().
		Return(config.OAuth).
		Times(1)

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

func TestListBuilder(t *testing.T) {
	test.CmdValidator(
		t,
		ListBuilder(),
		0,
		[]string{flag.ProjectID},
	)
}

func TestWarnMissingPodman(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockStore := mocks.NewMockClusterLister(ctrl)
	mockCredentialsGetter := mocks.NewMockCredentialsGetter(ctrl)
	mockProfileReader := mocks.NewMockProfileReader(ctrl)
	mockPodman := mocks.NewMockClient(ctrl)
	ctx := context.Background()

	buf := new(bytes.Buffer)
	log.SetWriter(buf)

	listOpts := &ListOpts{
		store:  mockStore,
		config: mockProfileReader,
		DeploymentOpts: options.DeploymentOpts{
			PodmanClient: mockPodman,
			CredStore:    mockCredentialsGetter,
		},
		GlobalOpts: cli.GlobalOpts{
			ProjectID: "64f670f0bf789926667dad1a",
		},
		OutputOpts: cli.OutputOpts{
			Template:  listTemplate,
			OutWriter: buf,
		},
	}

	emptyClusters := &admin.PaginatedAdvancedClusterDescription{
		Results: []admin.AdvancedClusterDescription{},
	}

	mockStore.
		EXPECT().
		ProjectClusters(listOpts.ProjectID,
			&mongodbatlas.ListOptions{
				PageNum:      cli.DefaultPage,
				ItemsPerPage: MaxItemsPerPage,
			},
		).
		Return(emptyClusters, nil).
		Times(1)

	mockCredentialsGetter.
		EXPECT().
		AuthType().
		Return(config.OAuth).
		Times(2)

	mockPodman.
		EXPECT().
		Ready(ctx).
		Return(podman.ErrPodmanNotFound).
		Times(1)

	if err := listOpts.Run(ctx); err != nil {
		t.Fatalf("Run() unexpected error: %v", err)
	}

	assert.Equal(t, `NAME   TYPE   MDB VER   STATE
`, buf.String())
	t.Log(buf.String())

	if err := listOpts.PostRun(ctx); err != nil {
		t.Fatalf("PostRun() unexpected error: %v", err)
	}
	assert.Equal(t, `NAME   TYPE   MDB VER   STATE
To get output for both local and Atlas clusters, install Podman.
`, buf.String())
	t.Log(buf.String())
}
