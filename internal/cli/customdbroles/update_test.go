// Copyright 2021 MongoDB Inc
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

package customdbroles

import (
	"testing"

	atlasv2 "go.mongodb.org/atlas-sdk/v20250312005/admin"
	"go.uber.org/mock/gomock"
)

func TestUpdateOpts_Run(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockStore := NewMockDatabaseRoleUpdater(ctrl)

	expected := &atlasv2.UserCustomDBRole{}

	updateOpts := &UpdateOpts{
		store: mockStore,
	}

	t.Run("default", func(t *testing.T) {
		mockStore.
			EXPECT().
			UpdateDatabaseRole(updateOpts.ConfigProjectID(), updateOpts.roleName, updateOpts.newCustomDBRole(expected)).Return(expected, nil).
			Times(1)

		if err := updateOpts.Run(); err != nil {
			t.Fatalf("Run() unexpected error: %v", err)
		}
	})
	t.Run("with append", func(t *testing.T) {
		updateOpts.append = true
		mockStore.
			EXPECT().
			DatabaseRole(updateOpts.ConfigProjectID(), updateOpts.roleName).Return(expected, nil).
			Times(1)
		mockStore.
			EXPECT().
			UpdateDatabaseRole(updateOpts.ConfigProjectID(), updateOpts.roleName, updateOpts.newCustomDBRole(expected)).Return(expected, nil).
			Times(1)

		if err := updateOpts.Run(); err != nil {
			t.Fatalf("Run() unexpected error: %v", err)
		}
	})
}
