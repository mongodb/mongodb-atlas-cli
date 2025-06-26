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

package invitations

import (
	"testing"

	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli"
	atlasv2 "go.mongodb.org/atlas-sdk/v20250312004/admin"
	"go.uber.org/mock/gomock"
)

func TestUpdate_Run(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockStore := NewMockProjectInvitationUpdater(ctrl)

	expected := &atlasv2.GroupInvitation{}

	updateOpts := &UpdateOpts{
		roles:       []string{"test"},
		store:       mockStore,
		ProjectOpts: cli.ProjectOpts{ProjectID: "1"},
	}

	mockStore.
		EXPECT().
		UpdateProjectInvitation(updateOpts.ConfigProjectID(), updateOpts.invitationID, updateOpts.newInvitation()).
		Return(expected, nil).
		Times(1)

	if err := updateOpts.Run(); err != nil {
		t.Fatalf("Run() unexpected error: %v", err)
	}
}
