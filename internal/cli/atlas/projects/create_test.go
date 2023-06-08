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

//go:build unit

package projects

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/mongodb/mongodb-atlas-cli/internal/flag"
	mocks "github.com/mongodb/mongodb-atlas-cli/internal/mocks/atlas"
	"github.com/mongodb/mongodb-atlas-cli/internal/pointer"
	"github.com/mongodb/mongodb-atlas-cli/internal/test"
	atlasv2 "go.mongodb.org/atlas-sdk/admin"
)

func TestCreate_Run(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockStore := mocks.NewMockProjectCreator(ctrl)

	opts := CreateOpts{}
	expected := &atlasv2.Group{}

	createOpts := &CreateOpts{
		store: mockStore,
		name:  "ProjectBar",
	}
	createOpts.OrgID = "5a0a1e7e0f2912c554080adc"
	params := &atlasv2.CreateProjectApiParams{
		ProjectOwnerId: &opts.projectOwnerID,
		Group:          pointer.Get(createOpts.newCreateProjectGroup()),
	}

	mockStore.
		EXPECT().
		CreateProject(params).Return(expected, nil).
		Times(1)

	if err := createOpts.Run(); err != nil {
		t.Fatalf("Run() unexpected error: %v", err)
	}
}

func TestCreateBuilder(t *testing.T) {
	test.CmdValidator(
		t,
		CreateBuilder(),
		0,
		[]string{flag.OrgID, flag.OwnerID, flag.GovCloudRegionsOnly, flag.WithoutDefaultAlertSettings},
	)
}
