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

package connection

import (
	"bytes"
	"testing"

	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/test"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/atlas-sdk/v20250312005/admin"
	"go.uber.org/mock/gomock"
)

func TestListOpts_Run(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockStore := NewMockStreamsConnectionLister(ctrl)

	expected := admin.PaginatedApiStreamsConnection{
		Results: &[]admin.StreamsConnection{
			{
				Name:             admin.PtrString("ExampleConn"),
				Type:             admin.PtrString("Kafka"),
				BootstrapServers: admin.PtrString("example.com:8080"),
			},
			{
				Name:        admin.PtrString("ExampleConn2"),
				Type:        admin.PtrString("Cluster"),
				ClusterName: admin.PtrString("MyCluster"),
			},
		},
	}

	buf := new(bytes.Buffer)
	listOpts := &ListOpts{
		store: mockStore,
		OutputOpts: cli.OutputOpts{
			Template:  listTemplate,
			OutWriter: buf,
		},
	}

	mockStore.
		EXPECT().
		StreamsConnections(listOpts.ProjectID, "").
		Return(&expected, nil).
		Times(1)

	if err := listOpts.Run(); err != nil {
		t.Fatalf("Run() unexpected error: %v", err)
	}
	t.Log(buf.String())
	test.VerifyOutputTemplate(t, listTemplate, expected)
	assert.Equal(t, `NAME           TYPE      SERVERS
ExampleConn    Kafka     example.com:8080
ExampleConn2   Cluster   MyCluster
`, buf.String())
}
