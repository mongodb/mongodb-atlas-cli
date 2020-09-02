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

package teams

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/mongodb/mongocli/internal/mocks"
	"go.mongodb.org/atlas/mongodbatlas"
)

func TestDescribe_Run_ByID(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockStore := mocks.NewMockTeamDescriber(ctrl)
	defer ctrl.Finish()

	var expected *mongodbatlas.Team

	descOpts := &DescribeOpts{
		store: mockStore,
		id:    "id",
	}

	mockStore.
		EXPECT().
		TeamByID(descOpts.OrgID, descOpts.id).
		Return(expected, nil).
		Times(1)

	err := descOpts.Run()
	if err != nil {
		t.Fatalf("Run() unexpected error: %v", err)
	}
}

func TestDescribe_Run_ByName(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockStore := mocks.NewMockTeamDescriber(ctrl)
	defer ctrl.Finish()

	var expected *mongodbatlas.Team

	descOpts := &DescribeOpts{
		store: mockStore,
		name:  "test",
	}

	mockStore.
		EXPECT().
		TeamByName(descOpts.OrgID, descOpts.name).
		Return(expected, nil).
		Times(1)

	err := descOpts.Run()
	if err != nil {
		t.Fatalf("Run() unexpected error: %v", err)
	}
}
