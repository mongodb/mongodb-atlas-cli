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

package invitations

import (
	"encoding/json"
	"testing"

	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/pointer"
	"github.com/spf13/afero"
	"github.com/stretchr/testify/require"
	"go.mongodb.org/atlas-sdk/v20250312006/admin"
	"go.uber.org/mock/gomock"
)

func TestCreate_Run(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockStore := NewMockOrganizationInviter(ctrl)

	expected := &admin.OrganizationInvitation{}
	opts := &InviteOpts{
		store:    mockStore,
		username: "test",
	}

	request, err := opts.newInvitation()
	require.NoError(t, err)

	mockStore.
		EXPECT().
		InviteUser(opts.ConfigOrgID(), request).Return(expected, nil).
		Times(1)

	require.NoError(t, opts.Run())
}

func TestInvite_Run_WithFile(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockStore := NewMockOrganizationInviter(ctrl)
	fs := afero.NewMemMapFs()

	testFile := "invitation.json"
	_, _ = fs.Create(testFile)
	invitation := &admin.OrganizationInvitationRequest{
		Username: pointer.Get("test-user@mongodb.com"),
		Roles:    pointer.Get([]string{"ORG_READ_ONLY"}),
		TeamIds:  pointer.Get([]string{"5f71e5255afec75a3d0f96dc"}),
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

	opts := &InviteOpts{
		store:    mockStore,
		fs:       fs,
		filename: testFile,
	}

	mockStore.
		EXPECT().
		InviteUser(opts.ConfigOrgID(), invitation).Return(expectedResult, nil).
		Times(1)

	require.NoError(t, opts.Run())
}
