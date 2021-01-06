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

package dbroles

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/mongodb/mongocli/internal/flag"
	"github.com/mongodb/mongocli/internal/mocks"
	"github.com/mongodb/mongocli/internal/test"
	"go.mongodb.org/atlas/mongodbatlas"
)

func TestUpdateOpts_Run(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockStore := mocks.NewMockDatabaseRoleUpdater(ctrl)
	defer ctrl.Finish()

	expected := &mongodbatlas.CustomDBRole{}

	updateOpts := &UpdateOpts{
		store: mockStore,
	}

	mockStore.
		EXPECT().
		UpdateDatabaseRole(updateOpts.ConfigProjectID(), updateOpts.roleName, updateOpts.newCustomDBRole(expected)).Return(expected, nil).
		Times(1)

	err := updateOpts.Run()
	if err != nil {
		t.Fatalf("Run() unexpected error: %v", err)
	}
}

func TestUpdateOptsWithAppend_Run(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockStore := mocks.NewMockDatabaseRoleUpdater(ctrl)
	defer ctrl.Finish()

	expected := &mongodbatlas.CustomDBRole{}

	updateOpts := &UpdateOpts{
		store:  mockStore,
		append: true,
	}

	mockStore.
		EXPECT().
		DatabaseRole(updateOpts.ConfigProjectID(), updateOpts.roleName).Return(expected, nil).
		Times(1)
	mockStore.
		EXPECT().
		UpdateDatabaseRole(updateOpts.ConfigProjectID(), updateOpts.roleName, updateOpts.newCustomDBRole(expected)).Return(expected, nil).
		Times(1)

	err := updateOpts.Run()
	if err != nil {
		t.Fatalf("Run() unexpected error: %v", err)
	}
}

func TestUpdateBuilder(t *testing.T) {
	test.CmdValidator(
		t,
		UpdateBuilder(),
		0,
		[]string{flag.ProjectID, flag.Output, flag.Database, flag.Action, flag.InheritedRole},
	)
}
