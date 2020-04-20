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

package cli

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/mongodb/mongocli/internal/mocks"
)

func TestAtlasClustersIndexesCreate_Run(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockStore := mocks.NewMockIndexCreator(ctrl)

	defer ctrl.Finish()

	createOpts := &atlasClustersIndexesCreateOpts{
		name:        "ProjectBar",
		clusterName: "US",
		db:          "test",
		collection:  "test",
		keys:        []string{"name:1"},
		store:       mockStore,
	}

	index, _ := createOpts.newIndex()
	mockStore.
		EXPECT().
		CreateIndex(createOpts.projectID, createOpts.clusterName, index).
		Return(nil).
		Times(1)

	err := createOpts.Run()
	if err != nil {
		t.Fatalf("Run() unexpected error: %v", err)
	}
}
