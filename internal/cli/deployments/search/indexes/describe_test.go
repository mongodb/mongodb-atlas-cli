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
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli/deployments/test/fixture"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/container"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/flag"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/mocks"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/pointer"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/test"
	"github.com/stretchr/testify/assert"
	atlasv2 "go.mongodb.org/atlas-sdk/v20240530002/admin"
)

func TestDescribe_RunLocal(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockMongodbClient := mocks.NewMockMongoDBClient(ctrl)
	mockStore := mocks.NewMockSearchIndexDescriber(ctrl)
	ctx := context.Background()

	const (
		expectedLocalDeployment = "localDeployment1"
		expectedStatus          = "STEADY"
		expectedType            = "search"
	)

	deploymentTest := fixture.NewMockLocalDeploymentOpts(ctrl, expectedLocalDeployment)

	buf := new(bytes.Buffer)
	opts := &DescribeOpts{
		DeploymentOpts: *deploymentTest.Opts,
		OutputOpts: cli.OutputOpts{
			OutWriter: buf,
			Template:  describeTemplate,
		},
		mongodbClient: mockMongodbClient,
		indexID:       "test",
		store:         mockStore,
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

	expected := &atlasv2.ClusterSearchIndex{
		Name:           "name",
		IndexID:        pointer.Get("test"),
		CollectionName: "coll",
		Database:       "db",
		Status:         pointer.Get(expectedStatus),
		Type:           pointer.Get(expectedType),
	}

	mockMongodbClient.
		EXPECT().
		SearchIndex("test").
		Return(expected, nil).
		Times(1)

	if err := opts.Run(ctx); err != nil {
		t.Fatalf("Run() unexpected error: %v", err)
	}

	assert.Equal(t, `ID     NAME   DATABASE   COLLECTION   STATUS   TYPE
test   name   db         coll         STEADY   search
`, buf.String())
	t.Log(buf.String())
}

func TestDescribe_RunAtlas(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockMongodbClient := mocks.NewMockMongoDBClient(ctrl)
	mockStore := mocks.NewMockSearchIndexDescriber(ctrl)
	ctx := context.Background()

	var (
		expectedLocalDeployment = "localDeployment1"
		name                    = "name"
		database                = "db"
		collectionName          = "coll"
	)

	deploymentTest := fixture.NewMockAtlasDeploymentOpts(ctrl, expectedLocalDeployment)

	buf := new(bytes.Buffer)
	opts := &DescribeOpts{
		DeploymentOpts: *deploymentTest.Opts,
		OutputOpts: cli.OutputOpts{
			OutWriter: buf,
			Template:  describeTemplate,
		},
		GlobalOpts: cli.GlobalOpts{
			ProjectID: "ProjectID",
		},
		mongodbClient: mockMongodbClient,
		indexID:       "test",
		store:         mockStore,
	}

	deploymentTest.CommonAtlasMocks(opts.ProjectID)

	mockStore.
		EXPECT().
		SearchIndex(opts.ProjectID, opts.DeploymentName, opts.indexID).
		Return(&atlasv2.ClusterSearchIndex{
			Name:           name,
			Database:       database,
			CollectionName: collectionName,
			IndexID:        pointer.Get("test"),
		}, nil).
		Times(1)

	expected := &atlasv2.ClusterSearchIndex{
		Name:           name,
		Database:       database,
		CollectionName: collectionName,
		IndexID:        pointer.Get("test"),
	}

	test.VerifyOutputTemplate(t, describeTemplate, expected)
	if err := opts.Run(ctx); err != nil {
		t.Fatalf("Run() unexpected error: %v", err)
	}
}

func TestDescribeOpts_PostRun(t *testing.T) {
	ctrl := gomock.NewController(t)
	deploymentTest := fixture.NewMockLocalDeploymentOpts(ctrl, "localDeployment")
	buf := new(bytes.Buffer)

	opts := &DescribeOpts{
		DeploymentOpts: *deploymentTest.Opts,
		OutputOpts: cli.OutputOpts{
			OutWriter: buf,
		},
	}

	deploymentTest.
		MockDeploymentTelemetry.
		EXPECT().
		AppendDeploymentType().
		Times(1)

	opts.PostRun()
}

func TestDescribeBuilder(t *testing.T) {
	test.CmdValidator(
		t,
		DescribeBuilder(),
		0,
		[]string{flag.DeploymentName, flag.TypeFlag, flag.ProjectID},
	)
}
