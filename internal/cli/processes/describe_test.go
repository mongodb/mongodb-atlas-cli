// Copyright 2022 MongoDB Inc
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

package processes

import (
	"testing"

	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/test"
	atlasv2 "go.mongodb.org/atlas-sdk/v20250312005/admin"
	"go.uber.org/mock/gomock"
)

func TestDescribeOpts_Run(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockStore := NewMockProcessDescriber(ctrl)

	expected := atlasv2.ApiHostViewAtlas{}

	opts := &DescribeOpts{
		store: mockStore,
		ProjectOpts: cli.ProjectOpts{
			ProjectID: "projectID",
		},
		host: "host",
		port: 1000,
	}
	params := atlasv2.GetAtlasProcessApiParams{
		ProcessId: "host:1000",
		GroupId:   "projectID",
	}

	mockStore.
		EXPECT().
		Process(&params).
		Return(&expected, nil).
		Times(1)

	if err := opts.Run(); err != nil {
		t.Fatalf("Run() unexpected error: %v", err)
	}
	test.VerifyOutputTemplate(t, describeTemplate, expected)
}
