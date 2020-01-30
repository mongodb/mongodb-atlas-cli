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

	"github.com/10gen/mcli/internal/fixtures"
	"github.com/10gen/mcli/internal/mocks"
	"github.com/golang/mock/gomock"
)

func TestAtlasClustersUpdate_Run(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockStore := mocks.NewMockClusterStore(ctrl)

	defer ctrl.Finish()

	expected := fixtures.Cluster()

	createOpts := &atlasClustersUpdateOpts{
		globalOpts:   newGlobalOpts(),
		name:         "ProjectBar",
		instanceSize: atlasM2,
		diskSizeGB:   10,
		mdbVersion:   currentMDBVersion,
		store:        mockStore,
	}

	mockStore.
		EXPECT().
		Cluster(createOpts.projectID, createOpts.name).
		Return(expected, nil).
		Times(1)

	createOpts.update(expected)

	mockStore.
		EXPECT().
		UpdateCluster(expected).Return(expected, nil).
		Times(1)

	err := createOpts.Run()
	if err != nil {
		t.Fatalf("Run() unexpected error: %v", err)
	}
}
