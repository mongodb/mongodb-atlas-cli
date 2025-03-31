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
	"encoding/json"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/mocks"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/pointer"
	"github.com/spf13/afero"
	"github.com/stretchr/testify/require"
	"go.mongodb.org/atlas-sdk/v20250312001/admin"
)

func TestUpdate_Run(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockStore := mocks.NewMockOrganizationInvitationUpdater(ctrl)

	expected := &admin.OrganizationInvitation{}

	updateOpts := &UpdateOpts{
		roles:   []string{"test"},
		store:   mockStore,
		OrgOpts: cli.OrgOpts{OrgID: "1"},
	}

	request, err := updateOpts.newInvitation()
	require.NoError(t, err)

	mockStore.
		EXPECT().
		UpdateOrganizationInvitation(updateOpts.ConfigOrgID(), updateOpts.invitationID, request).
		Return(expected, nil).
		Times(1)

	if err := updateOpts.Run(); err != nil {
		t.Fatalf("Run() unexpected error: %v", err)
	}
}

func TestUpdate_Run_WithFile(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockStore := mocks.NewMockOrganizationInvitationUpdater(ctrl)
	fs := afero.NewMemMapFs()

	testFile := "update_invitation.json"
	invitationID := "6d39e6f9a16946a1abc390d4"
	_, _ = fs.Create(testFile)
	invitation := &admin.OrganizationInvitationRequest{
		Roles: pointer.Get([]string{"ORG_MEMBER"}),
		GroupRoleAssignments: pointer.Get([]admin.OrganizationInvitationGroupRoleAssignmentsRequest{
			{
				GroupId: pointer.Get("6c73999ae7966f00563911a4"),
				Roles:   pointer.Get([]string{"GROUP_CLUSTER_MANAGER"}),
			},
		}),
	}
	invitationJSON, err := json.Marshal(invitation)
	require.NoError(t, err)
	_ = afero.WriteFile(fs, testFile, invitationJSON, 0600)

	expectedResult := &admin.OrganizationInvitation{}

	opts := &UpdateOpts{
		store:        mockStore,
		fs:           fs,
		filename:     testFile,
		invitationID: invitationID,
		OrgOpts:      cli.OrgOpts{OrgID: "1"},
	}

	mockStore.
		EXPECT().
		UpdateOrganizationInvitation(opts.ConfigOrgID(), invitationID, invitation).Return(expectedResult, nil).
		Times(1)

	if err := opts.Run(); err != nil {
		t.Fatalf("Run() unexpected error: %v", err)
	}
}
