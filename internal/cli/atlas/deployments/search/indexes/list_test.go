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
	"bytes"
	"context"
	"fmt"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/mongodb/mongodb-atlas-cli/internal/cli"
	"github.com/mongodb/mongodb-atlas-cli/internal/cli/atlas/deployments/options"
	"github.com/mongodb/mongodb-atlas-cli/internal/cli/atlas/search"
	"github.com/mongodb/mongodb-atlas-cli/internal/flag"
	"github.com/mongodb/mongodb-atlas-cli/internal/mocks"
	"github.com/mongodb/mongodb-atlas-cli/internal/mongodbclient"
	"github.com/mongodb/mongodb-atlas-cli/internal/podman"
	"github.com/mongodb/mongodb-atlas-cli/internal/pointer"
	"github.com/mongodb/mongodb-atlas-cli/internal/test"
	"github.com/stretchr/testify/assert"
	atlasv2 "go.mongodb.org/atlas-sdk/v20230201007/admin"
)

func TestList_RunLocal(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockPodman := mocks.NewMockClient(ctrl)
	mockMongodbClient := mocks.NewMockMongoDBClient(ctrl)
	mockDB := mocks.NewMockDatabase(ctrl)
	mockStore := mocks.NewMockSearchIndexLister(ctrl)
	ctx := context.Background()

	const (
		expectedLocalDeployment = "localDeployment1"
		expectedDB              = "db1"
		expectedCollection      = "col1"
		expectedName            = "test"
		expectedID              = "1"
	)

	buf := new(bytes.Buffer)
	opts := &ListOpts{
		DeploymentOpts: options.DeploymentOpts{
			PodmanClient:   mockPodman,
			DeploymentName: expectedLocalDeployment,
		},
		OutputOpts: cli.OutputOpts{
			OutWriter: buf,
			Template:  listTemplate,
		},
		mongodbClient: mockMongodbClient,
		IndexOpts: search.IndexOpts{
			DBName:     expectedDB,
			Collection: expectedCollection,
		},
		store: mockStore,
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
		Database(expectedDB).
		Return(mockDB).
		Times(1)

	expected := []*mongodbclient.SearchIndexDefinition{
		{
			Name:       expectedName,
			ID:         expectedID,
			Collection: expectedCollection,
			Database:   expectedDB,
		},
	}

	mockDB.
		EXPECT().
		SearchIndexes(ctx, expectedCollection).
		Return(expected, nil).
		Times(1)

	if err := opts.Run(ctx); err != nil {
		t.Fatalf("Run() unexpected error: %v", err)
	}

	assert.Equal(t, fmt.Sprintf(`ID    NAME   DATABASE   COLLECTION
%s     %s   %s        %s
`, expectedID, expectedName, expectedDB, expectedCollection), buf.String())
	t.Log(buf.String())
}

func TestList_RunAtlas(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockPodman := mocks.NewMockClient(ctrl)
	mockMongodbClient := mocks.NewMockMongoDBClient(ctrl)
	mockStore := mocks.NewMockSearchIndexLister(ctrl)
	ctx := context.Background()

	const (
		expectedLocalDeployment = "localDeployment1"
		expectedDB              = "db1"
		expectedCollection      = "col1"
		expectedName            = "test"
		expectedID              = "1"
		expectedProjectID       = "1"
	)

	buf := new(bytes.Buffer)
	opts := &ListOpts{
		DeploymentOpts: options.DeploymentOpts{
			PodmanClient:   mockPodman,
			DeploymentName: expectedLocalDeployment,
		},
		OutputOpts: cli.OutputOpts{
			OutWriter: buf,
			Template:  listTemplate,
		},
		mongodbClient: mockMongodbClient,
		IndexOpts: search.IndexOpts{
			DBName:     expectedDB,
			Collection: expectedCollection,
		},
		GlobalOpts: cli.GlobalOpts{
			ProjectID: expectedProjectID,
		},
		store: mockStore,
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
		SearchIndexes(opts.ProjectID, opts.DeploymentName, opts.DBName, opts.Collection).
		Return([]atlasv2.ClusterSearchIndex{
			{
				Name:           expectedName,
				Database:       expectedDB,
				CollectionName: expectedCollection,
				IndexID:        pointer.GetStringPointerIfNotEmpty(expectedID),
			},
		}, nil).
		Times(1)

	expected := []*mongodbclient.SearchIndexDefinition{
		{
			Name:       expectedName,
			ID:         expectedID,
			Collection: expectedCollection,
			Database:   expectedDB,
		},
	}

	test.VerifyOutputTemplate(t, listTemplate, expected)
	if err := opts.Run(ctx); err != nil {
		t.Fatalf("Run() unexpected error: %v", err)
	}
}

func TestListBuilder(t *testing.T) {
	test.CmdValidator(
		t,
		ListBuilder(),
		0,
		[]string{flag.DeploymentName, flag.ProjectID, flag.Database, flag.Collection},
	)
}
