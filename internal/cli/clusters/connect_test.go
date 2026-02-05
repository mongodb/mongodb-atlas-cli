// Copyright 2026 MongoDB Inc
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

package clusters

import (
	"bytes"
	"testing"

	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/pointer"
	"github.com/stretchr/testify/assert"
	atlasClustersPinned "go.mongodb.org/atlas-sdk/v20240530005/admin"
	"go.uber.org/mock/gomock"
)

const expectedAtlasCluster = "atlasCluster1"

func TestRun_ConnectAtlas(t *testing.T) {
	ctrl := gomock.NewController(t)
	ctx := t.Context()
	buf := new(bytes.Buffer)

	mockStore := NewMockConnectClusterStore(ctrl)

	connectOpts := &ConnectOpts{
		OutputOpts: cli.OutputOpts{
			OutWriter: buf,
		},
		ProjectOpts: cli.ProjectOpts{
			ProjectID: "projectID",
		},
		name:                 expectedAtlasCluster,
		ConnectWith:          "connectionString",
		ConnectionStringType: "standard",
		store:                mockStore,
	}

	expectedCluster := &atlasClustersPinned.AdvancedClusterDescription{
		Name:           pointer.Get(expectedAtlasCluster),
		Id:             pointer.Get("123"),
		MongoDBVersion: pointer.Get("7.0.0"),
		StateName:      pointer.Get("IDLE"),
		Paused:         pointer.Get(false),
		ConnectionStrings: &atlasClustersPinned.ClusterConnectionStrings{
			StandardSrv: pointer.Get("mongodb://localhost:27017/?directConnection=true"),
		},
	}

	mockStore.
		EXPECT().
		AtlasCluster(connectOpts.ProjectID, expectedAtlasCluster).
		Return(expectedCluster, nil).
		Times(1)

	if err := Run(ctx, connectOpts); err != nil {
		t.Fatalf("Run() unexpected error: %v", err)
	}

	assert.Equal(t, `mongodb://localhost:27017/?directConnection=true
`, buf.String())
}
