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
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli/deployments/test/fixture"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli/search"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/container"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/flag"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/mocks"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/pointer"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/test"
	"github.com/stretchr/testify/assert"
	atlasv2 "go.mongodb.org/atlas-sdk/v20240530002/admin"
)

func TestList_RunLocal(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockMongodbClient := mocks.NewMockMongoDBClient(ctrl)
	mockDB := mocks.NewMockDatabase(ctrl)
	ctx := context.Background()

	const (
		expectedLocalDeployment = "localDeployment1"
		expectedDB              = "db1"
		expectedCollection      = "col1"
		expectedName            = "test"
		expectedID              = "1"
		expectedStatus          = "STEADY"
		expectedType            = "search"
	)

	buf := new(bytes.Buffer)
	deploymentTest := fixture.NewMockLocalDeploymentOpts(ctrl, expectedLocalDeployment)

	opts := &ListOpts{
		DeploymentOpts: *deploymentTest.Opts,
		OutputOpts: cli.OutputOpts{
			OutWriter: buf,
			Template:  listTemplate,
		},
		mongodbClient: mockMongodbClient,
		IndexOpts: search.IndexOpts{
			DBName:     expectedDB,
			Collection: expectedCollection,
		},
	}

	deploymentTest.LocalMockFlow(ctx)

	deploymentTest.MockContainerEngine.
		EXPECT().
		ContainerInspect(ctx, expectedLocalDeployment).
		Return([]*container.InspectData{
			{
				Name: expectedLocalDeployment,
				Config: &container.InspectDataConfig{
					Labels: map[string]string{
						"version": "7.0.1",
					},
				},
				HostConfig: &container.InspectDataHostConfig{
					PortBindings: map[string][]container.InspectDataHostPort{
						"27017/tcp": {
							{
								HostIP:   "127.0.0.1",
								HostPort: "27017",
							},
						},
					},
				},
			},
		}, nil).
		Times(1)
	mockMongodbClient.
		EXPECT().
		Disconnect().
		Times(1)

	mockMongodbClient.
		EXPECT().
		Connect("mongodb://localhost:27017/?directConnection=true", int64(10)).
		Return(nil).
		Times(1)

	mockMongodbClient.
		EXPECT().
		Database(expectedDB).
		Return(mockDB).
		Times(1)

	expected := []*atlasv2.ClusterSearchIndex{
		{
			Name:           expectedName,
			IndexID:        pointer.Get(expectedID),
			CollectionName: expectedCollection,
			Database:       expectedDB,
			Status:         pointer.Get(expectedStatus),
			Type:           pointer.Get(expectedType),
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

	assert.Equal(t, fmt.Sprintf(`ID    NAME   DATABASE   COLLECTION   STATUS   TYPE
%s     %s   %s        %s         %s   %s
`, expectedID, expectedName, expectedDB, expectedCollection, expectedStatus, expectedType), buf.String())
	t.Log(buf.String())
}

func TestList_RunAtlas(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockMongodbClient := mocks.NewMockMongoDBClient(ctrl)
	mockStore := mocks.NewMockSearchIndexLister(ctrl)
	ctx := context.Background()

	var (
		expectedLocalDeployment = "localDeployment1"
		expectedDB              = "db1"
		expectedCollection      = "col1"
		expectedName            = "test"
		expectedID              = "1"
		expectedProjectID       = "1"
	)

	buf := new(bytes.Buffer)
	deploymentTest := fixture.NewMockAtlasDeploymentOpts(ctrl, expectedLocalDeployment)

	opts := &ListOpts{
		DeploymentOpts: *deploymentTest.Opts,
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

	deploymentTest.CommonAtlasMocks(expectedProjectID)

	mockStore.
		EXPECT().
		SearchIndexes(opts.ProjectID, opts.DeploymentName, opts.DBName, opts.Collection).
		Return([]atlasv2.ClusterSearchIndex{
			{
				Name:           expectedName,
				Database:       expectedDB,
				CollectionName: expectedCollection,
				IndexID:        pointer.Get(expectedID),
			},
		}, nil).
		Times(1)

	expected := []*atlasv2.ClusterSearchIndex{
		{
			Name:           expectedName,
			IndexID:        pointer.Get(expectedID),
			CollectionName: expectedCollection,
			Database:       expectedDB,
		},
	}

	test.VerifyOutputTemplate(t, listTemplate, expected)
	if err := opts.Run(ctx); err != nil {
		t.Fatalf("Run() unexpected error: %v", err)
	}
}

func TestList_PostRun(t *testing.T) {
	ctrl := gomock.NewController(t)
	deploymentTest := fixture.NewMockLocalDeploymentOpts(ctrl, "localDeployment")
	buf := new(bytes.Buffer)

	opts := &ListOpts{
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

func TestListBuilder(t *testing.T) {
	test.CmdValidator(
		t,
		ListBuilder(),
		0,
		[]string{flag.DeploymentName, flag.TypeFlag, flag.ProjectID, flag.Database, flag.Collection},
	)
}
