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

package indexes

import (
	"bytes"
	"context"
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli/deployments/test/fixture"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli/search"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/container"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/flag"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/mocks"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/mongodbclient"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/pointer"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/test"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	atlasv2 "go.mongodb.org/atlas-sdk/v20240530002/admin"
)

var indexID = "6509bc5080b2f007e6a2a0ce"

const (
	expectedIndexName       = "idx1"
	expectedLocalDeployment = "localDeployment1"
	expectedDB              = "db1"
	expectedCollection      = "col1"
)

func TestCreate_RunLocal(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockMongodbClient := mocks.NewMockMongoDBClient(ctrl)
	mockDB := mocks.NewMockDatabase(ctrl)
	ctx := context.Background()

	testDeployments := fixture.NewMockLocalDeploymentOpts(ctrl, expectedLocalDeployment)

	buf := new(bytes.Buffer)
	opts := &CreateOpts{
		DeploymentOpts: *testDeployments.Opts,
		IndexOpts: search.IndexOpts{
			Name:       expectedIndexName,
			DBName:     expectedDB,
			Collection: expectedCollection,
		},
		OutputOpts: cli.OutputOpts{
			OutWriter: buf,
			Template:  createTemplate,
		},
		mongodbClient: mockMongodbClient,
	}

	testDeployments.LocalMockFlow(ctx)

	testDeployments.MockContainerEngine.
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
		Connect("mongodb://localhost:27017/?directConnection=true", int64(10)).
		Return(nil).
		Times(1)
	mockMongodbClient.
		EXPECT().
		Disconnect().
		Times(1)
	mockMongodbClient.
		EXPECT().
		Database(expectedDB).
		Return(mockDB).
		Times(1)

	index := &atlasv2.ClusterSearchIndex{
		Analyzer:       &opts.Analyzer,
		CollectionName: opts.Collection,
		Database:       opts.DBName,
		Mappings: &atlasv2.ApiAtlasFTSMappings{
			Dynamic: &opts.Dynamic,
			Fields:  nil,
		},
		Name:           opts.Name,
		SearchAnalyzer: &opts.SearchAnalyzer,
		Type:           pointer.Get(search.DefaultType),
	}

	indexWithID := &atlasv2.ClusterSearchIndex{
		Analyzer:       &opts.Analyzer,
		CollectionName: opts.Collection,
		Database:       opts.DBName,
		Mappings: &atlasv2.ApiAtlasFTSMappings{
			Dynamic: &opts.Dynamic,
			Fields:  nil,
		},
		Name:           opts.Name,
		SearchAnalyzer: &opts.SearchAnalyzer,
		IndexID:        &indexID,
		Type:           pointer.Get(search.DefaultType),
	}

	mockDB.
		EXPECT().
		SearchIndexByName(ctx, index.Name, index.CollectionName).
		Return(nil, mongodbclient.ErrSearchIndexNotFound).
		Times(1)

	mockDB.
		EXPECT().
		CreateSearchIndex(ctx, expectedCollection, gomock.Any()).
		Return(indexWithID, nil).
		Times(1)

	if err := opts.Run(ctx); err != nil {
		t.Fatalf("Run() unexpected error: %v", err)
	}
	if err := opts.PostRun(ctx); err != nil {
		t.Fatalf("PostRun() unexpected error: %v", err)
	}

	assert.Equal(t, `Search index created with ID: 6509bc5080b2f007e6a2a0ce
`, buf.String())
	t.Log(buf.String())
}

func TestCreate_Duplicated(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockMongodbClient := mocks.NewMockMongoDBClient(ctrl)
	mockDB := mocks.NewMockDatabase(ctrl)
	ctx := context.Background()

	testDeployments := fixture.NewMockLocalDeploymentOpts(ctrl, expectedLocalDeployment)

	buf := new(bytes.Buffer)
	opts := &CreateOpts{
		DeploymentOpts: *testDeployments.Opts,
		IndexOpts: search.IndexOpts{
			Name:       expectedIndexName,
			DBName:     expectedDB,
			Collection: expectedCollection,
		},
		OutputOpts: cli.OutputOpts{
			OutWriter: buf,
			Template:  createTemplate,
		},
		mongodbClient: mockMongodbClient,
	}

	testDeployments.LocalMockFlow(ctx)

	testDeployments.MockContainerEngine.
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
		Connect("mongodb://localhost:27017/?directConnection=true", int64(10)).
		Return(nil).
		Times(1)
	mockMongodbClient.
		EXPECT().
		Disconnect().
		Times(1)
	mockMongodbClient.
		EXPECT().
		Database(expectedDB).
		Return(mockDB).
		Times(1)

	index := &atlasv2.ClusterSearchIndex{
		Analyzer:       &opts.Analyzer,
		CollectionName: opts.Collection,
		Database:       opts.DBName,
		Mappings: &atlasv2.ApiAtlasFTSMappings{
			Dynamic: &opts.Dynamic,
			Fields:  nil,
		},
		Name:           opts.Name,
		SearchAnalyzer: &opts.SearchAnalyzer,
	}

	indexWithID := &atlasv2.ClusterSearchIndex{
		Analyzer:       &opts.Analyzer,
		CollectionName: opts.Collection,
		Database:       opts.DBName,
		Mappings: &atlasv2.ApiAtlasFTSMappings{
			Dynamic: &opts.Dynamic,
			Fields:  nil,
		},
		Name:           opts.Name,
		SearchAnalyzer: &opts.SearchAnalyzer,
		IndexID:        &indexID,
	}

	mockDB.
		EXPECT().
		SearchIndexByName(ctx, index.Name, index.CollectionName).
		Return(indexWithID, nil).
		Times(1)
	if err := opts.Run(ctx); err == nil || !errors.Is(err, ErrSearchIndexDuplicated) {
		t.Fatalf("Run() unexpected error: %v", err)
	}
}

func TestCreate_RunAtlas(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockIndexStore := mocks.NewMockSearchIndexCreator(ctrl)
	ctx := context.Background()
	buf := new(bytes.Buffer)

	deploymentTest := fixture.NewMockAtlasDeploymentOpts(ctrl, expectedLocalDeployment)

	opts := &CreateOpts{
		DeploymentOpts: *deploymentTest.Opts,
		IndexOpts: search.IndexOpts{
			Name:       expectedIndexName,
			DBName:     expectedDB,
			Collection: expectedCollection,
		},
		OutputOpts: cli.OutputOpts{
			OutWriter: buf,
			Template:  createTemplate,
		},
		GlobalOpts: cli.GlobalOpts{
			ProjectID: "projectID",
		},
		store: mockIndexStore,
	}

	index, err := opts.NewSearchIndex()
	require.NoError(t, err)

	indexWithID := &atlasv2.ClusterSearchIndex{
		CollectionName: opts.Collection,
		Database:       opts.DBName,
		Analyzer:       &opts.Analyzer,
		Mappings: &atlasv2.ApiAtlasFTSMappings{
			Dynamic: &opts.Dynamic,
			Fields:  nil,
		},
		SearchAnalyzer: &opts.SearchAnalyzer,
		Name:           opts.Name,
		IndexID:        &indexID,
	}

	deploymentTest.CommonAtlasMocks(opts.ProjectID)

	mockIndexStore.
		EXPECT().
		CreateSearchIndexes(opts.ProjectID, opts.DeploymentName, index).
		Times(1).
		Return(indexWithID, nil)

	if err := opts.Run(ctx); err != nil {
		t.Fatalf("Run() unexpected error: %v", err)
	}

	if err := opts.PostRun(ctx); err != nil {
		t.Fatalf("PostRun() unexpected error: %v", err)
	}

	assert.Equal(t, `Search index created with ID: 6509bc5080b2f007e6a2a0ce
`, buf.String())
	t.Log(buf.String())
}

func TestCreateBuilder(t *testing.T) {
	test.CmdValidator(
		t,
		CreateBuilder(),
		0,
		[]string{flag.DeploymentName, flag.Database, flag.Collection, flag.File},
	)
}
