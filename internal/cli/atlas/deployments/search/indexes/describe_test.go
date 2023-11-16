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

	"github.com/containers/podman/v4/libpod/define"
	"github.com/golang/mock/gomock"
	"github.com/mongodb/mongodb-atlas-cli/internal/cli"
	"github.com/mongodb/mongodb-atlas-cli/internal/cli/atlas/deployments/options"
	"github.com/mongodb/mongodb-atlas-cli/internal/cli/atlas/deployments/test/fixture"
	"github.com/mongodb/mongodb-atlas-cli/internal/flag"
	"github.com/mongodb/mongodb-atlas-cli/internal/mocks"
	"github.com/mongodb/mongodb-atlas-cli/internal/pointer"
	"github.com/mongodb/mongodb-atlas-cli/internal/test"
	"github.com/stretchr/testify/assert"
	atlasv2 "go.mongodb.org/atlas-sdk/v20231115001/admin"
)

func TestDescribe_RunLocal(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockMongodbClient := mocks.NewMockMongoDBClient(ctrl)
	mockStore := mocks.NewMockSearchIndexDescriber(ctrl)
	ctx := context.Background()

	const (
		expectedIndexName       = "idx1"
		expectedLocalDeployment = "localDeployment1"
		expectedDB              = "db1"
		expectedCollection      = "col1"
		expectedStatus          = "STEADY"
	)

	deploymentTest := fixture.NewMockLocalDeploymentOpts(ctrl, expectedLocalDeployment)
	mockPodman := deploymentTest.MockPodman

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

	mockPodman.
		EXPECT().
		ContainerInspect(ctx, options.MongodHostnamePrefix+"-"+expectedLocalDeployment).
		Return([]*define.InspectContainerData{
			{
				Name: options.MongodHostnamePrefix + "-" + expectedLocalDeployment,
				Config: &define.InspectContainerConfig{
					Labels: map[string]string{
						"version": "7.0.1",
					},
				},
				HostConfig: &define.InspectContainerHostConfig{
					PortBindings: map[string][]define.InspectHostPort{
						"27017/tcp": {
							{
								HostIP:   "127.0.0.1",
								HostPort: "27017",
							},
						},
					},
				},
				Mounts: []define.InspectMount{
					{
						Name: opts.DeploymentOpts.LocalMongodDataVolume(),
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
		IndexID:        pointer.GetStringPointerIfNotEmpty("test"),
		CollectionName: "coll",
		Database:       "db",
		Status:         pointer.GetStringPointerIfNotEmpty(expectedStatus),
	}

	mockMongodbClient.
		EXPECT().
		SearchIndex("test").
		Return(expected, nil).
		Times(1)

	if err := opts.Run(ctx); err != nil {
		t.Fatalf("Run() unexpected error: %v", err)
	}

	assert.Equal(t, `ID     NAME   DATABASE   COLLECTION   STATUS
test   name   db         coll         STEADY
`, buf.String())
	t.Log(buf.String())
}

func TestDescribe_RunAtlas(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockMongodbClient := mocks.NewMockMongoDBClient(ctrl)
	mockStore := mocks.NewMockSearchIndexDescriber(ctrl)
	ctx := context.Background()

	const (
		expectedIndexName       = "idx1"
		expectedLocalDeployment = "localDeployment1"
		expectedDB              = "db1"
		expectedCollection      = "col1"
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
			Name:           "name",
			Database:       "db",
			CollectionName: "coll",
			IndexID:        pointer.GetStringPointerIfNotEmpty("test"),
		}, nil).
		Times(1)

	expected := &atlasv2.ClusterSearchIndex{
		Name:           "name",
		IndexID:        pointer.GetStringPointerIfNotEmpty("test"),
		CollectionName: "coll",
		Database:       "db",
	}

	test.VerifyOutputTemplate(t, describeTemplate, expected)
	if err := opts.Run(ctx); err != nil {
		t.Fatalf("Run() unexpected error: %v", err)
	}
}

func TestDescribeBuilder(t *testing.T) {
	test.CmdValidator(
		t,
		DescribeBuilder(),
		0,
		[]string{flag.DeploymentName, flag.TypeFlag, flag.ProjectID},
	)
}
