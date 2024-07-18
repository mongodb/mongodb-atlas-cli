//go:build unit

package indexes

import (
	"bytes"
	"context"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli/deployments/test/fixture"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli/search"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/container"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/mocks"
	"github.com/stretchr/testify/assert"
	atlasv2 "go.mongodb.org/atlas-sdk/v20240530002/admin"
)

func TestUpdate_RunLocal(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockMongodbClient := mocks.NewMockMongoDBClient(ctrl)
	mockDB := mocks.NewMockDatabase(ctrl)
	ctx := context.Background()

	testDeployments := fixture.NewMockLocalDeploymentOpts(ctrl, expectedLocalDeployment)

	buf := new(bytes.Buffer)
	opts := &UpdateOpts{
		DeploymentOpts: *testDeployments.Opts,
		IndexOpts: search.IndexOpts{
			Name:       expectedIndexName,
			DBName:     expectedDB,
			Collection: expectedCollection,
			Dynamic:    true,
			Analyzer:   "unit.test.analyzer",
		},
		OutputOpts: cli.OutputOpts{
			OutWriter: buf,
			Template:  updateTemplate,
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

	existingIndex := &atlasv2.ClusterSearchIndex{
		IndexID:        &indexID,
		Name:           expectedIndexName,
		Database:       expectedDB,
		CollectionName: expectedCollection,
	}

	updatedIndex := &atlasv2.ClusterSearchIndex{
		IndexID:        &indexID,
		Name:           expectedIndexName,
		Database:       expectedDB,
		CollectionName: expectedCollection,
		Analyzer:       &opts.Analyzer,
		SearchAnalyzer: &opts.SearchAnalyzer,
		Mappings: &atlasv2.ApiAtlasFTSMappings{
			Dynamic: &opts.Dynamic,
		},
	}

	mockDB.
		EXPECT().
		SearchIndexByName(ctx, expectedIndexName, expectedCollection).
		Return(existingIndex, nil).
		Times(1)

	mockDB.
		EXPECT().
		UpdateSearchIndex(ctx, expectedCollection, gomock.Any()).
		Return(updatedIndex, nil).
		Times(1)

	if err := opts.Run(ctx); err != nil {
		t.Fatalf("Run() unexpected error: %v", err)
	}
	if err := opts.PostRun(ctx); err != nil {
		t.Fatalf("PostRun() unexpected error: %v", err)
	}

	assert.Equal(t, `Search index updated with ID: 6509bc5080b2f007e6a2a0ce
`, buf.String())
	t.Log(buf.String())
}

func TestUpdate_RunAtlas(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockIndexStore := mocks.NewMockSearchIndexUpdater(ctrl)
	ctx := context.Background()
	buf := new(bytes.Buffer)

	deploymentTest := fixture.NewMockAtlasDeploymentOpts(ctrl, expectedLocalDeployment)

	opts := &UpdateOpts{
		DeploymentOpts: *deploymentTest.Opts,
		IndexOpts: search.IndexOpts{
			Name:       expectedIndexName,
			DBName:     expectedDB,
			Collection: expectedCollection,
		},
		OutputOpts: cli.OutputOpts{
			OutWriter: buf,
			Template:  updateTemplate,
		},
		GlobalOpts: cli.GlobalOpts{
			ProjectID: "projectID",
		},
		store: mockIndexStore,
	}

	updatedIndex := atlasv2.ClusterSearchIndex{
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
		SearchIndexes(opts.ProjectID, opts.DeploymentName, opts.DBName, opts.Collection).
		Times(1).
		Return([]atlasv2.ClusterSearchIndex{updatedIndex}, nil)

	mockIndexStore.
		EXPECT().
		UpdateSearchIndexes(opts.ProjectID, opts.DeploymentName, *updatedIndex.IndexID, gomock.Any()).
		Times(1).
		Return(&updatedIndex, nil)

	if err := opts.Run(ctx); err != nil {
		t.Fatalf("Run() unexpected error: %v", err)
	}

	if err := opts.PostRun(ctx); err != nil {
		t.Fatalf("PostRun() unexpected error: %v", err)
	}

	assert.Equal(t, `Search index updated with ID: 6509bc5080b2f007e6a2a0ce
`, buf.String())
	t.Log(buf.String())
}
