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
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/flag"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/mocks"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/test"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/atlas-sdk/v20240530002/admin"
)

func TestDescribe_Run(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockStore := mocks.NewMockStreamsConnectionDescriber(ctrl)
	expected := admin.StreamsConnection{Name: admin.PtrString("id"), Type: admin.PtrString("Kafka"), BootstrapServers: admin.PtrString("example.com:8080")}

	buf := new(bytes.Buffer)
	describeOpts := &DescribeOpts{
		id:              "id",
		streamsInstance: "Foo",
		store:           mockStore,
		OutputOpts: cli.OutputOpts{
			Template:  describeTemplate,
			OutWriter: buf,
		},
	}

	mockStore.
		EXPECT().
		StreamConnection(describeOpts.ConfigProjectID(), describeOpts.streamsInstance, describeOpts.id).
		Return(&expected, nil).
		Times(1)

	if err := describeOpts.Run(); err != nil {
		t.Fatalf("Run() unexpected error: %v", err)
	}
	t.Log(buf.String())
	test.VerifyOutputTemplate(t, describeTemplate, expected)
	assert.Equal(t, `NAME   TYPE    SERVERS
id     Kafka   example.com:8080
`, buf.String())
}

func TestDescribeBuilder(t *testing.T) {
	test.CmdValidator(
		t,
		DescribeBuilder(),
		0,
		[]string{flag.ProjectID, flag.Output},
	)
}
