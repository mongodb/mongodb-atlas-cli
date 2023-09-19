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

package instance

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/mongodb/mongodb-atlas-cli/internal/flag"
	"github.com/mongodb/mongodb-atlas-cli/internal/mocks"
	"github.com/mongodb/mongodb-atlas-cli/internal/store"
	"github.com/mongodb/mongodb-atlas-cli/internal/test"
	"go.mongodb.org/atlas/mongodbatlas"
)

func TestUpdateOpts_Run(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockStore := mocks.NewMockStreamsUpdater(ctrl)

	updateOpts := &UpdateOpts{
		store:    mockStore,
		name:     "Example Name",
		provider: "AWS",
		region:   "US-EAST-1",
	}

	expected := &store.StreamProcessorInstance{Name: updateOpts.name, GroupID: updateOpts.ProjectID, DataProcessRegion: mongodbatlas.DataProcessRegion{CloudProvider: "AWS", Region: "US-EAST-1"}}

	mockStore.
		EXPECT().
		UpdateStream(updateOpts.ProjectID, updateOpts.name, expected).
		Return(expected, nil).
		Times(1)

	if err := updateOpts.Run(); err != nil {
		t.Fatalf("Run() unexpected error: %v", err)
	}
}

func TestUpdateBuilder(t *testing.T) {
	test.CmdValidator(
		t,
		UpdateBuilder(),
		0,
		[]string{flag.Provider, flag.Region},
	)
}
