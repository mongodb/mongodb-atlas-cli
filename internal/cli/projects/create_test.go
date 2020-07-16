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

package projects

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/mongodb/mongocli/internal/mocks"
	"go.mongodb.org/atlas/mongodbatlas"
)

func TestCreate_Run(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockStore := mocks.NewMockProjectCreator(ctrl)

	defer ctrl.Finish()

	expected := &mongodbatlas.Project{}

	mockStore.
		EXPECT().
		CreateProject(gomock.Eq("ProjectBar"), gomock.Eq("5a0a1e7e0f2912c554080adc")).Return(expected, nil).
		Times(1)

	createOpts := &CreateOpts{
		store: mockStore,
		name:  "ProjectBar",
	}
	createOpts.OrgID = "5a0a1e7e0f2912c554080adc"
	err := createOpts.Run()
	if err != nil {
		t.Fatalf("Run() unexpected error: %v", err)
	}
}
