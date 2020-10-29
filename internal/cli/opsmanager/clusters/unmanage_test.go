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

package clusters

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/mongodb/mongocli/internal/cli"
	"github.com/mongodb/mongocli/internal/fixture"
	"github.com/mongodb/mongocli/internal/mocks"
)

func TestUnmanage_Run(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockStore := mocks.NewMockAutomationPatcher(ctrl)
	defer ctrl.Finish()

	expected := fixture.AutomationConfig()

	deleteOpts := &UnmanageOpts{
		store: mockStore,
		DeleteOpts: &cli.DeleteOpts{
			Confirm: true,
			Entry:   "myReplicaSet",
		},
	}

	mockStore.
		EXPECT().
		GetAutomationConfig(deleteOpts.ProjectID).
		Return(expected, nil).
		Times(1)

	mockStore.
		EXPECT().
		UpdateAutomationConfig(deleteOpts.ProjectID, expected).
		Return(nil).
		Times(1)

	err := deleteOpts.Run()
	if err != nil {
		t.Fatalf("Run() unexpected error: %v", err)
	}
}
