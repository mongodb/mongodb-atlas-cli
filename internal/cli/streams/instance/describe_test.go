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

package instance

import (
	"bytes"
	"testing"

	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/test"
	"github.com/stretchr/testify/assert"
	atlasv2 "go.mongodb.org/atlas-sdk/v20250312005/admin"
	"go.uber.org/mock/gomock"
)

func TestDescribeOpts_Run(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockStore := NewMockStreamsDescriber(ctrl)

	buf := new(bytes.Buffer)
	describeOpts := &DescribeOpts{
		store: mockStore,
		name:  "Example Name",
		OutputOpts: cli.OutputOpts{
			Template:  describeTemplate,
			OutWriter: buf,
		},
	}

	id := "1"
	name := "ExampleInstance"
	expected := &atlasv2.StreamsTenant{Id: &id, Name: &name}
	expected.DataProcessRegion = atlasv2.NewStreamsDataProcessRegion("AWS", "US_EAST_1")

	mockStore.
		EXPECT().
		AtlasStream(describeOpts.ConfigProjectID(), describeOpts.name).
		Return(expected, nil).
		Times(1)

	if err := describeOpts.Run(); err != nil {
		t.Fatalf("Run() unexpected error: %v", err)
	}
	t.Log(buf.String())
	test.VerifyOutputTemplate(t, describeTemplate, expected)
	assert.Equal(t, `ID    NAME              CLOUD   REGION
1     ExampleInstance   AWS     US_EAST_1
`, buf.String())
}
