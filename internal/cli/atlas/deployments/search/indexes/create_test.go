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
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/mongodb/mongodb-atlas-cli/internal/cli"
	"github.com/mongodb/mongodb-atlas-cli/internal/cli/atlas/deployments/options"
	"github.com/mongodb/mongodb-atlas-cli/internal/cli/atlas/search"
	"github.com/mongodb/mongodb-atlas-cli/internal/flag"
	"github.com/mongodb/mongodb-atlas-cli/internal/mocks"
	"github.com/mongodb/mongodb-atlas-cli/internal/podman"
	"github.com/mongodb/mongodb-atlas-cli/internal/test"
	"github.com/stretchr/testify/assert"
	atlasv2 "go.mongodb.org/atlas-sdk/v20230201008/admin"
)

func TestCreate_Run(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockPodman := mocks.NewMockClient(ctrl)
	mockMongodbClient := mocks.NewMockMongoDBClient(ctrl)
	mockDB := mocks.NewMockDatabase(ctrl)
	ctx := context.Background()

	const (
		expectedIndexName       = "idx1"
		expectedLocalDeployment = "localDeployment1"
		expectedDB              = "db1"
		expectedCollection      = "col1"
	)

	buf := new(bytes.Buffer)
	opts := &CreateOpts{
		DeploymentOpts: options.DeploymentOpts{
			PodmanClient:   mockPodman,
			DeploymentName: expectedLocalDeployment,
		},
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
		Times(2)

	mockMongodbClient.
		EXPECT().
		Connect("mongodb://localhost:0/?directConnection=true", int64(10)).
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

	mockDB.
		EXPECT().
		CreateSearchIndex(ctx, expectedCollection, &atlasv2.ClusterSearchIndex{
			Analyzer:       &opts.Analyzer,
			CollectionName: opts.Collection,
			Database:       opts.DBName,
			Mappings: &atlasv2.ApiAtlasFTSMappings{
				Dynamic: &opts.Dynamic,
				Fields:  nil,
			},
			Name:           opts.Name,
			SearchAnalyzer: &opts.SearchAnalyzer,
		}).
		Return(nil).
		Times(1)

	if err := opts.Run(ctx); err != nil {
		t.Fatalf("Run() unexpected error: %v", err)
	}

	assert.Equal(t, `Your search index is being created
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
