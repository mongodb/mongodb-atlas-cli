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

package users

import (
	"testing"

	atlasv2 "go.mongodb.org/atlas-sdk/v20250312012/admin"
	"go.uber.org/mock/gomock"
)

func TestAdd_Run(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockStore := NewMockTeamAdder(ctrl)

	var expected *atlasv2.PaginatedApiAppUser

	listOpts := &AddOpts{
		store: mockStore,
	}

	mockStore.
		EXPECT().
		AddUsersToTeam(listOpts.OrgID, listOpts.teamID, listOpts.newUsers()).
		Return(expected, nil).
		Times(1)

	if err := listOpts.Run(); err != nil {
		t.Fatalf("Run() unexpected error: %v", err)
	}
}
