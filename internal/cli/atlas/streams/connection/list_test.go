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

package connection

import (
	"bytes"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/mongodb/mongodb-atlas-cli/internal/cli"
	"github.com/mongodb/mongodb-atlas-cli/internal/flag"
	"github.com/mongodb/mongodb-atlas-cli/internal/mocks"
	"github.com/mongodb/mongodb-atlas-cli/internal/store"
	"github.com/mongodb/mongodb-atlas-cli/internal/test"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/atlas-sdk/v20230201004/admin"
)

func TestListOpts_Run_display(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockStore := mocks.NewMockStreamsConnectionLister(ctrl)

	expected := store.StreamsConnectionList{
		PaginatedApiStreamsConnection: admin.PaginatedApiStreamsConnection{},
		Connections: []store.StreamsConnection{
			{StreamsConnection: admin.StreamsConnection{
				Name: admin.PtrString("ExampleConn"),
				Type: admin.PtrString("Kafka")},
				Instance: "Floop",
				Servers:  "example.com:8080"},
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
		Return(expected, nil).
		Times(1)

	if err := listOpts.Run(); err != nil {
		t.Fatalf("Run() unexpected error: %v", err)
	}
	t.Log(buf.String())
	test.VerifyOutputTemplate(t, listTemplate, expected)
	assert.Equal(t, `NAME          TYPE    INSTANCE   SERVERS
ExampleConn   Kafka   Floop      example.com:8080

`, buf.String())
}

func TestListOpts_Run_json(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockStore := mocks.NewMockStreamsConnectionLister(ctrl)

	expected := store.StreamsConnectionList{
		PaginatedApiStreamsConnection: admin.PaginatedApiStreamsConnection{
			Links: []admin.Link{},
			Results: []admin.StreamsConnection{
				{Name: admin.PtrString("Fraud"),
					Type: admin.PtrString("Kafka"),
					Authentication: &admin.StreamsKafkaAuthentication{
						Mechanism: admin.PtrString("SCRAM-256"),
						Username:  admin.PtrString("admin"),
					},
					BootstrapServers: admin.PtrString("another.example.com:8080"),
					Security: &admin.StreamsKafkaSecurity{
						Protocol: admin.PtrString("PLAINTEXT"),
					}},
			},
			TotalCount: admin.PtrInt(1)},
		Connections: []store.StreamsConnection{},
	}

	buf := new(bytes.Buffer)
	listOpts := &ListOpts{
		store: mockStore,
		OutputOpts: cli.OutputOpts{
			Template:  listTemplate,
			OutWriter: buf,
			Output:    "json",
		},
	}

	mockStore.
		EXPECT().
		StreamsConnections(listOpts.ProjectID, "").
		Return(expected, nil).
		Times(1)

	if err := listOpts.Run(); err != nil {
		t.Fatalf("Run() unexpected error: %v", err)
	}
	t.Log(buf.String())
	assert.Equal(t, `[
  {
    "name": "Fraud",
    "type": "Kafka",
    "authentication": {
      "mechanism": "SCRAM-256",
      "username": "admin"
    },
    "bootstrapServers": "another.example.com:8080",
    "security": {
      "protocol": "PLAINTEXT"
    }
  }
]
`, buf.String())
}

func TestListBuilder(t *testing.T) {
	test.CmdValidator(
		t,
		ListBuilder(),
		0,
		[]string{flag.ProjectID, flag.Output, flag.Instance},
	)
}
