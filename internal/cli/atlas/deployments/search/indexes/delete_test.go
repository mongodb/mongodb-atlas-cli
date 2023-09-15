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
	"github.com/mongodb/mongodb-atlas-cli/internal/cli"
	"github.com/mongodb/mongodb-atlas-cli/internal/cli/atlas/deployments/options"
	"github.com/mongodb/mongodb-atlas-cli/internal/flag"
	"github.com/mongodb/mongodb-atlas-cli/internal/mocks"
	"github.com/mongodb/mongodb-atlas-cli/internal/podman"
	"github.com/mongodb/mongodb-atlas-cli/internal/test"
)

func TestDelete_RunLocal(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockPodman := mocks.NewMockClient(ctrl)
	mockMongodbClient := mocks.NewMockMongoDBClient(ctrl)
	mockStore := mocks.NewMockSearchIndexDeleter(ctrl)
	ctx := context.Background()

	const (
		expectedIndexName       = "idx1"
		expectedLocalDeployment = "localDeployment1"
		expectedDB              = "db1"
		expectedCollection      = "col1"
		indexID                 = "1"
	)

	opts := &DeleteOpts{
		DeploymentOpts: options.DeploymentOpts{
			PodmanClient:   mockPodman,
			DeploymentName: expectedLocalDeployment,
		},
		DeleteOpts: &cli.DeleteOpts{
			Entry:   indexID,
			Confirm: true,
		},
		mongodbClient: mockMongodbClient,
		store:         mockStore,
	}

	mockPodman.
		EXPECT().
		ListContainers(ctx, options.MongodHostnamePrefix).
		Return([]*podman.Container{
			{
				Names:  []string{options.MongodHostnamePrefix + "-" + expectedLocalDeployment},
				State:  "running",
				Labels: map[string]string{"version": "6.0.9"},
			},
		}, nil).
		Times(1)
	mockMongodbClient.
		EXPECT().
		Disconnect(ctx).
		Times(1)

	mockMongodbClient.
		EXPECT().
		Connect(ctx, "mongodb://localhost:0/?directConnection=true", int64(10)).
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
	mockPodman := mocks.NewMockClient(ctrl)
	mockMongodbClient := mocks.NewMockMongoDBClient(ctrl)
	mockStore := mocks.NewMockSearchIndexDeleter(ctrl)
	ctx := context.Background()

	const (
		expectedIndexName       = "idx1"
		expectedLocalDeployment = "localDeployment1"
		expectedDB              = "db1"
		expectedCollection      = "col1"
		indexID                 = "1"
		projectID               = "1"
	)

	opts := &DeleteOpts{
		DeploymentOpts: options.DeploymentOpts{
			PodmanClient:   mockPodman,
			DeploymentName: expectedLocalDeployment,
		},
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

	mockPodman.
		EXPECT().
		ListContainers(ctx, options.MongodHostnamePrefix).
		Return([]*podman.Container{
			{
				Names:  []string{options.MongodHostnamePrefix + "-" + expectedLocalDeployment},
				State:  "running",
				Labels: map[string]string{"version": "6.0.9"},
			},
		}, nil).
		Times(1)

	mockMongodbClient.
		EXPECT().
		Connect(ctx, "mongodb://localhost:0/?directConnection=true", int64(10)).
		Return(options.ErrDeploymentNotFound).
		Times(1)

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
		[]string{flag.DeploymentName, flag.ProjectID, flag.Force},
	)
}
