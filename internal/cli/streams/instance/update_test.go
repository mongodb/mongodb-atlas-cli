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

	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/test"
	atlasv2 "go.mongodb.org/atlas-sdk/v20250312005/admin"
	"go.uber.org/mock/gomock"
)

const (
	updateTestProjectID = "update-project-id"
)

func TestUpdateOpts_Run(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockStore := NewMockStreamsUpdater(ctrl)
	t.Run("stream instances update --name --provider --region", func(t *testing.T) {
		buf := new(bytes.Buffer)
		updateOpts := &UpdateOpts{
			store:    mockStore,
			name:     "Example Name",
			provider: "AWS",
			region:   "VIRGINIA_USA",
		}

		expected := &atlasv2.StreamsTenant{Name: &updateOpts.name, GroupId: &updateOpts.ProjectID, DataProcessRegion: &atlasv2.StreamsDataProcessRegion{CloudProvider: "AWS", Region: "VIRGINIA_USA"}}
		updateOpts.ProjectID = updateTestProjectID

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
	})

	// Testing the parsing of flags but not passing into StreamConfig object
	t.Run("stream workspaces update --tier --defaultTier --maxTierSize", func(t *testing.T) {
		buf := new(bytes.Buffer)
		updateOpts := &UpdateOpts{
			store:       mockStore,
			name:        "Example Name",
			provider:    "AWS",
			region:      "VIRGINIA_USA",
			tier:        "SP30",
			defaultTier: "SP30",
			maxTierSize: "SP30",
		}

		expected := &atlasv2.StreamsTenant{Name: &updateOpts.name, GroupId: &updateOpts.ProjectID, DataProcessRegion: &atlasv2.StreamsDataProcessRegion{CloudProvider: "AWS", Region: "VIRGINIA_USA"}}
		updateOpts.ProjectID = updateTestProjectID

		mockStore.
			EXPECT().
			UpdateStream(updateOpts.ProjectID, updateOpts.name, expected.DataProcessRegion).
			Return(expected, nil).
			Times(1)

		if err := updateOpts.Run(); err != nil {
			t.Fatalf("Run() unexpected error: %v", err)
		}
		t.Log(buf.String())
		test.VerifyOutputTemplate(t, updateTemplateWorkspace, expected)
	})
}
