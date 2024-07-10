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

package instance

import (
	"bytes"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/flag"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/mocks"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/test"
	atlasv2 "go.mongodb.org/atlas-sdk/v20240530002/admin"
)

func TestUpdateOpts_Run(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockStore := mocks.NewMockStreamsUpdater(ctrl)

	buf := new(bytes.Buffer)
	updateOpts := &UpdateOpts{
		store:    mockStore,
		name:     "Example Name",
		provider: "AWS",
		region:   "VIRGINIA_USA",
	}

	expected := &atlasv2.StreamsTenant{Name: &updateOpts.name, GroupId: &updateOpts.ProjectID, DataProcessRegion: &atlasv2.StreamsDataProcessRegion{CloudProvider: "AWS", Region: "VIRGINIA_USA"}}
	updateOpts.ProjectID = "update-project-id"

	mockStore.
		EXPECT().
		UpdateStream(updateOpts.ProjectID, updateOpts.name, expected.DataProcessRegion).
		Return(expected, nil).
		Times(1)

	if err := updateOpts.Run(); err != nil {
		t.Fatalf("Run() unexpected error: %v", err)
	}
	t.Log(buf.String())
	test.VerifyOutputTemplate(t, updateTemplate, expected)
}

func TestUpdateBuilder(t *testing.T) {
	test.CmdValidator(
		t,
		UpdateBuilder(),
		0,
		[]string{flag.Provider, flag.Region},
	)
}
