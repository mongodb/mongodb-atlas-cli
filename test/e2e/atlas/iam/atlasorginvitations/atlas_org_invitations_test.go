// Copyright 2021 MongoDB Inc
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

package atlasorginvitations

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"testing"

	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/pointer"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/test/internal"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.mongodb.org/atlas-sdk/v20250312027/admin"
)

const (
	orgEntity         = "org"
	invitationsEntity = "invitations"
)
const (
	roleNameOrg = "ORG_READ_ONLY"
)

func TestAtlasOrgInvitations(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping test in short mode")
	}
	cliPath, err := internal.AtlasCLIBin()
	require.NoError(t, err)

	n := internal.Must(internal.RandInt(1000))

	emailOrg := fmt.Sprintf("test-%v@mongodb.com", n)
	var orgInvitationID string
	var orgInvitationIDFile string // For the file-based invite test

	t.Run("Invite", func(t *testing.T) {
		cmd := exec.Command(cliPath,
			orgEntity,
			invitationsEntity,
			"invite",
			emailOrg,
			"--role",
			"ORG_MEMBER",
			"-o=json",
			"-P",
			internal.ProfileName(),
		)
		cmd.Env = os.Environ()
		resp, err := internal.RunAndGetStdOut(cmd)
		a := assert.New(t)
		require.NoError(t, err, string(resp))

		var invitation admin.OrganizationInvitation
		require.NoError(t, json.Unmarshal(resp, &invitation))
		a.Equal(emailOrg, invitation.GetUsername())
		a.Equal([]string{"ORG_MEMBER"}, invitation.GetRoles())
		require.NotEmpty(t, invitation.GetId())
		orgInvitationID = invitation.GetId()
	})

	t.Run("Invite with File", func(t *testing.T) {
		a := assert.New(t)
		// Create a unique email for this test
		nFile := internal.Must(internal.RandInt(1000))
		emailOrgFile := fmt.Sprintf("test-file-%v@mongodb.com", nFile)

		inviteData := admin.OrganizationInvitationRequest{
			Username: pointer.Get(emailOrgFile),
			Roles:    pointer.Get([]string{"ORG_READ_ONLY"}),
		}
		inviteFilename := fmt.Sprintf("%s/update-%s.json", t.TempDir(), nFile)
		internal.CreateJSONFile(t, inviteData, inviteFilename)

		cmd := exec.Command(cliPath,
			orgEntity,
			invitationsEntity,
			"invite",
			"--file", inviteFilename,
			"-o=json",
			"-P",
			internal.ProfileName(),
		)
		cmd.Env = os.Environ()
		resp, err := internal.RunAndGetStdOut(cmd)
		require.NoError(t, err, string(resp))

		var invitation admin.OrganizationInvitation
		require.NoError(t, json.Unmarshal(resp, &invitation))
		a.Equal(emailOrgFile, invitation.GetUsername())
		a.Equal(inviteData.GetRoles(), invitation.GetRoles())
		require.NotEmpty(t, invitation.GetId())
		orgInvitationIDFile = invitation.GetId() // Save ID for cleanup
	})

	t.Run("Invite with File", func(t *testing.T) {
		a := assert.New(t)
		// Create a unique email for this test
		nFile := internal.Must(internal.RandInt(1000))
		emailOrgFile := fmt.Sprintf("test-file-%v@mongodb.com", nFile)

		inviteData := admin.OrganizationInvitationRequest{
			Username: pointer.Get(emailOrgFile),
			Roles:    pointer.Get([]string{"ORG_READ_ONLY"}),
		}
		inviteFilename := fmt.Sprintf("%s/update-%s.json", t.TempDir(), nFile)
		internal.CreateJSONFile(t, inviteData, inviteFilename)

		cmd := exec.Command(cliPath,
			orgEntity,
			invitationsEntity,
			"invite",
			"--file", inviteFilename,
			"-o=json",
			"-P",
			internal.ProfileName(),
		)
		cmd.Env = os.Environ()
		resp, err := internal.RunAndGetStdOut(cmd)
		require.NoError(t, err, string(resp))

		var invitation admin.OrganizationInvitation
		require.NoError(t, json.Unmarshal(resp, &invitation))
		a.Equal(emailOrgFile, invitation.GetUsername())
		a.Equal(inviteData.GetRoles(), invitation.GetRoles())
		require.NotEmpty(t, invitation.GetId())
		orgInvitationIDFile = invitation.GetId() // Save ID for cleanup
	})

	t.Run("List", func(t *testing.T) {
		cmd := exec.Command(cliPath,
			orgEntity,
			invitationsEntity,
			"ls",
			"-o=json",
			"-P",
			internal.ProfileName(),
		)
		cmd.Env = os.Environ()
		resp, err := internal.RunAndGetStdOut(cmd)
		require.NoError(t, err, string(resp))

		a := assert.New(t)

		var invitations []admin.OrganizationInvitation
		require.NoError(t, json.Unmarshal(resp, &invitations))
		a.NotEmpty(invitations)
	})

	t.Run("Describe", func(t *testing.T) {
		cmd := exec.Command(cliPath,
			orgEntity,
			invitationsEntity,
			"get",
			orgInvitationID,
			"-o=json",
			"-P",
			internal.ProfileName(),
		)
		cmd.Env = os.Environ()
		resp, err := internal.RunAndGetStdOut(cmd)
		require.NoError(t, err, string(resp))

		a := assert.New(t)

		var invitation admin.OrganizationInvitation
		require.NoError(t, json.Unmarshal(resp, &invitation))
		a.Equal(orgInvitationID, invitation.GetId())
		a.Equal([]string{"ORG_MEMBER"}, invitation.GetRoles())
	})

	t.Run("Update by email", func(t *testing.T) {
		cmd := exec.Command(cliPath,
			orgEntity,
			invitationsEntity,
			"update",
			"--email",
			emailOrg,
			"--role",
			roleNameOrg,
			"-o=json",
			"-P",
			internal.ProfileName(),
		)
		cmd.Env = os.Environ()
		resp, err := internal.RunAndGetStdOut(cmd)
		require.NoError(t, err, string(resp))

		a := assert.New(t)

		var invitation admin.OrganizationInvitation
		require.NoError(t, json.Unmarshal(resp, &invitation))
		a.Equal(emailOrg, invitation.GetUsername())
		a.ElementsMatch([]string{roleNameOrg}, invitation.GetRoles())
	})

	t.Run("Update by ID", func(t *testing.T) {
		cmd := exec.Command(cliPath,
			orgEntity,
			invitationsEntity,
			"update",
			orgInvitationID,
			"--role",
			roleNameOrg,
			"-o=json",
			"-P",
			internal.ProfileName(),
		)
		cmd.Env = os.Environ()
		resp, err := internal.RunAndGetStdOut(cmd)
		require.NoError(t, err, string(resp))

		a := assert.New(t)
		var invitation admin.OrganizationInvitation
		require.NoError(t, json.Unmarshal(resp, &invitation))
		a.Equal(emailOrg, invitation.GetUsername())
		a.ElementsMatch([]string{roleNameOrg}, invitation.GetRoles())
	})

	const OrgGroupCreator = "ORG_GROUP_CREATOR"

	t.Run("Update with File", func(t *testing.T) {
		require.NotEmpty(t, orgInvitationID, "orgInvitationID must be set by Invite test")
		a := assert.New(t)

		nFile := internal.Must(internal.RandInt(1000))

		// Define the update data, including GroupRoleAssignments if desired
		updateRole := OrgGroupCreator
		updateData := admin.OrganizationInvitationRequest{
			Roles: pointer.Get([]string{updateRole}),
		}
		updateFilename := fmt.Sprintf("%s/update-%s.json", t.TempDir(), nFile)
		internal.CreateJSONFile(t, updateData, updateFilename)

		cmd := exec.Command(cliPath,
			orgEntity,
			invitationsEntity,
			"update",
			orgInvitationID, // Use ID from the original Invite test
			"--file", updateFilename,
			"-o=json",
			"-P",
			internal.ProfileName(),
		)
		cmd.Env = os.Environ()
		resp, err := internal.RunAndGetStdOut(cmd)
		require.NoError(t, err, string(resp))

		var invitation admin.OrganizationInvitation
		require.NoError(t, json.Unmarshal(resp, &invitation))
		a.Equal(orgInvitationID, invitation.GetId())
		a.ElementsMatch(updateData.GetRoles(), invitation.GetRoles()) // Check if roles were updated
		// Add assertions for GroupRoleAssignments if included in updateData
	})

	t.Run("Update with File", func(t *testing.T) {
		require.NotEmpty(t, orgInvitationID, "orgInvitationID must be set by Invite test")
		a := assert.New(t)

		nFile := internal.Must(internal.RandInt(1000))

		// Define the update data, including GroupRoleAssignments if desired
		updateRole := OrgGroupCreator
		updateData := admin.OrganizationInvitationRequest{
			Roles: pointer.Get([]string{updateRole}),
		}
		updateFilename := fmt.Sprintf("%s/update-%s.json", t.TempDir(), nFile)
		internal.CreateJSONFile(t, updateData, updateFilename)

		cmd := exec.Command(cliPath,
			orgEntity,
			invitationsEntity,
			"update",
			orgInvitationID, // Use ID from the original Invite test
			"--file", updateFilename,
			"-o=json",
			"-P",
			internal.ProfileName(),
		)
		cmd.Env = os.Environ()
		resp, err := internal.RunAndGetStdOut(cmd)
		require.NoError(t, err, string(resp))

		var invitation admin.OrganizationInvitation
		require.NoError(t, json.Unmarshal(resp, &invitation))
		a.Equal(orgInvitationID, invitation.GetId())
		a.ElementsMatch(updateData.GetRoles(), invitation.GetRoles()) // Check if roles were updated
		// Add assertions for GroupRoleAssignments if included in updateData
	})

	t.Run("Delete", func(t *testing.T) {
		cmd := exec.Command(cliPath,
			orgEntity,
			invitationsEntity,
			"delete",
			orgInvitationID,
			"--force",
			"-P",
			internal.ProfileName(),
		)
		cmd.Env = os.Environ()
		resp, err := internal.RunAndGetStdOut(cmd)
		a := assert.New(t)
		require.NoError(t, err, string(resp))
		expected := fmt.Sprintf("Invitation '%s' deleted\n", orgInvitationID)
		a.Equal(expected, string(resp))
	})

	t.Run("Delete Invitation from File Test", func(t *testing.T) {
		require.NotEmpty(t, orgInvitationIDFile, "orgInvitationIDFile must be set by Invite with File test")
		cmd := exec.Command(cliPath,
			orgEntity,
			invitationsEntity,
			"delete",
			orgInvitationIDFile,
			"--force",
			"-P",
			internal.ProfileName(),
		)
		cmd.Env = os.Environ()
		resp, err := internal.RunAndGetStdOut(cmd)
		a := assert.New(t)
		require.NoError(t, err, string(resp))
		expected := fmt.Sprintf("Invitation '%s' deleted\n", orgInvitationIDFile)
		a.Equal(expected, string(resp))
	})
}
