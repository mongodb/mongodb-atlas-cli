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

	"github.com/golang/mock/gomock"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/flag"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/mocks"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/test"
	"github.com/spf13/afero"
	"github.com/stretchr/testify/require"
	atlasv2 "go.mongodb.org/atlas-sdk/v20240530002/admin"
)

func TestUpdate_Run(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockStore := mocks.NewMockConnectionUpdater(ctrl)

	fs := afero.NewMemMapFs()

	fileContents := `
{
  "type": "Kafka",
  "bootstrapServers": "example.com:8080,fraud.example.com:8000",
  "security": {
    "protocol": "PLAINTEXT"
  },
  "authentication": {
    "mechanism": "SCRAM-256",
    "username": "admin",
    "password": "hunter2"
  },
  "configuration": {
    "auto.offset.reset": "earliest"
  }
}
`

	fileName := "connection.json"
	require.NoError(t, afero.WriteFile(fs, fileName, []byte(fileContents), 0600))

	buf := new(bytes.Buffer)
	updateOpts := &UpdateOpts{
		name:            "id",
		store:           mockStore,
		fs:              fs,
		filename:        fileName,
		streamsInstance: "Example Instance",
	}

	name := "Example Name"
	con := atlasv2.NewStreamsConnectionWithDefaults()
	con.Name = &name

	expected := atlasv2.NewStreamsConnection()
	expected.Name = &name

	mockStore.
		EXPECT().
		UpdateConnection(updateOpts.ConfigProjectID(), "Example Instance", updateOpts.name, gomock.Any()).Return(expected, nil).
		Times(1)

	if err := updateOpts.Run(); err != nil {
		t.Fatalf("Run() unexpected error: %v", err)
	}
	t.Log(buf.String())
	test.VerifyOutputTemplate(t, updateTemplate, expected)
}

func TestUpdateBuilder(t *testing.T) {
	test.CmdValidator(
		t,
		UpdateBuilder(),
		0,
		[]string{flag.ProjectID, flag.Output, flag.File, flag.Instance},
	)
}
