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

package dbusers

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/mongodb/mongodb-atlas-cli/internal/mocks"
	"go.mongodb.org/atlas-sdk/v20230201008/admin"
)

func TestDBUserUpdate_Run(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockStore := mocks.NewMockDatabaseUserUpdater(ctrl)

	expected := &admin.CloudDatabaseUser{}

	updateOpts := &UpdateOpts{
		currentUsername: "test4",
		password:        "US",
		roles:           []string{"admin@admin"},
		store:           mockStore,
	}

	dbUser := admin.CloudDatabaseUser{}
	updateOpts.update(&dbUser)

	params := &admin.UpdateDatabaseUserApiParams{
		GroupId:           dbUser.GroupId,
		DatabaseName:      dbUser.DatabaseName,
		Username:          updateOpts.currentUsername,
		CloudDatabaseUser: &dbUser,
	}

	mockStore.
		EXPECT().
		UpdateDatabaseUser(params).
		Return(expected, nil).
		Times(1)

	if err := updateOpts.Run(); err != nil {
		t.Fatalf("Run() unexpected error: %v", err)
	}
}
