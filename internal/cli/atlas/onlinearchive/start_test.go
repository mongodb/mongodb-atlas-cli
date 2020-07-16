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

// +build unit

package onlinearchive

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/mongodb/mongocli/internal/mocks"
	"go.mongodb.org/atlas/mongodbatlas"
)

func TestStart_Run(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockStore := mocks.NewMockOnlineArchiveUpdater(ctrl)

	defer ctrl.Finish()

	updateOpts := &StartOpts{
		id:    "1",
		store: mockStore,
	}

	paused := false
	expected := &mongodbatlas.OnlineArchive{
		ID:     updateOpts.id,
		Paused: &paused,
	}

	mockStore.
		EXPECT().
		UpdateOnlineArchive(updateOpts.ConfigProjectID(), updateOpts.clusterName, expected).
		Return(expected, nil).
		Times(1)

	err := updateOpts.Run()
	if err != nil {
		t.Fatalf("Run() unexpected error: %v", err)
	}
}
