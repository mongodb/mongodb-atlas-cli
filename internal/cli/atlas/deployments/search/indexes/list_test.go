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
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/mongodb/mongodb-atlas-cli/internal/cli"
	"github.com/mongodb/mongodb-atlas-cli/internal/cli/atlas/deployments/options"
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
	mockStore := mocks.NewMockSearchIndexDescriber(ctrl)
	ctx := context.Background()

	const (
		expectedIndexName       = "idx1"
		expectedLocalDeployment = "localDeployment1"
		expectedDB              = "db1"
		expectedCollection      = "col1"
	)

	buf := new(bytes.Buffer)
	opts := &DescribeOpts{
		DeploymentOpts: options.DeploymentOpts{
			PodmanClient:   mockPodman,
			DeploymentName: expectedLocalDeployment,
		},
		OutputOpts: cli.OutputOpts{
			OutWriter: buf,
			Template:  describeTemplate,
		},
		mongodbClient: mockMongodbClient,
		indexID:       "test",
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

	expected := &mongodbclient.SearchIndexDefinition{
		Name:       "name",
		ID:         "test",
		Collection: "coll",
		Database:   "db",
	}

	mockMongodbClient.
		EXPECT().
		SearchIndex(ctx, "test").
		Return(expected, nil).
		Times(1)

	if err := opts.Run(ctx); err != nil {
		t.Fatalf("Run() unexpected error: %v", err)
	}

	assert.Equal(t, `ID     NAME   DATABASE   COLLECTION
test   name   db         coll
`, buf.String())
	t.Log(buf.String())
}

func TestList_RunAtlas(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockPodman := mocks.NewMockClient(ctrl)
	mockMongodbClient := mocks.NewMockMongoDBClient(ctrl)
	mockStore := mocks.NewMockSearchIndexDescriber(ctrl)
	ctx := context.Background()

	const (
		expectedIndexName       = "idx1"
		expectedLocalDeployment = "localDeployment1"
		expectedDB              = "db1"
		expectedCollection      = "col1"
	)

	buf := new(bytes.Buffer)
	opts := &DescribeOpts{
		DeploymentOpts: options.DeploymentOpts{
			PodmanClient:   mockPodman,
			DeploymentName: expectedLocalDeployment,
		},
		OutputOpts: cli.OutputOpts{
			OutWriter: buf,
			Template:  describeTemplate,
		},
		mongodbClient: mockMongodbClient,
		indexID:       "test",
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
		SearchIndex(ctx, "test").
		Return(nil, mongodbclient.ErrSearchIndexNotFound).
		Times(1)

	mockStore.
		EXPECT().
		SearchIndex(opts.ProjectID, opts.DeploymentName, opts.indexID).
		Return(&atlasv2.ClusterSearchIndex{
			Name:           "name",
			Database:       "db",
			CollectionName: "coll",
			IndexID:        pointer.GetStringPointerIfNotEmpty("test"),
		}, nil).
		Times(1)

	expected := &mongodbclient.SearchIndexDefinition{
		Name:       "name",
		ID:         "test",
		Collection: "coll",
		Database:   "db",
	}

	test.VerifyOutputTemplate(t, describeTemplate, expected)
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
