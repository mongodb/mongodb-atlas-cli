// Copyright 2020 MongoDB Inc
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

package clusters

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/mongodb/mongodb-atlas-cli/internal/flag"
	"github.com/mongodb/mongodb-atlas-cli/internal/mocks"
	"github.com/mongodb/mongodb-atlas-cli/internal/test"
	atlasv2 "go.mongodb.org/atlas-sdk/admin"
)

func TestDescribe_Run(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockStore := mocks.NewMockAtlasClusterDescriber(ctrl)

	expected := &atlasv2.AdvancedClusterDescription{}

	describeOpts := &DescribeOpts{
		name:  "test",
		store: mockStore,
	}

	mockStore.
		EXPECT().
		AtlasCluster(describeOpts.ProjectID, describeOpts.name).
		Return(expected, nil).
		Times(1)

	if err := describeOpts.Run(); err != nil {
		t.Fatalf("Run() unexpected error: %v", err)
	}
}

func TestDescribeBuilder(t *testing.T) {
	test.CmdValidator(
		t,
		DescribeBuilder(),
		0,
		[]string{flag.Output, flag.ProjectID},
	)
}
