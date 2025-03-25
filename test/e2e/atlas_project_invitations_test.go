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
//go:build e2e || (iam && !atlas)

package e2e_test

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.mongodb.org/atlas-sdk/v20250312001/admin"
)

func TestAtlasProjectInvitations(t *testing.T) {
	cliPath, err := AtlasCLIBin()
	require.NoError(t, err)

	g := newAtlasE2ETestGenerator(t)
	g.generateProject("invitations")

	var invitationID string
	n, err := RandInt(1000)
	require.NoError(t, err)

	emailProject := fmt.Sprintf("test-%v@mongodb.com", n)

	t.Run("Invite", func(t *testing.T) {
		cmd := exec.Command(cliPath,
			projectsEntity,
			invitationsEntity,
			"invite",
			emailProject,
			"--role",
			"GROUP_READ_ONLY",
			"--projectId",
			g.projectID,
			"-o=json")
		cmd.Env = os.Environ()
		resp, err := RunAndGetStdOut(cmd)
		require.NoError(t, err, string(resp))
		a := assert.New(t)

		var invitation admin.GroupInvitation
		require.NoError(t, json.Unmarshal(resp, &invitation))
		a.Equal(emailProject, invitation.GetUsername())
		require.NotEmpty(t, invitation.GetId())
		invitationID = invitation.GetId()
	})

	t.Run("List", func(t *testing.T) {
		cmd := exec.Command(cliPath,
			projectsEntity,
			invitationsEntity,
			"ls",
			"--projectId",
			g.projectID,
			"-o=json")
		cmd.Env = os.Environ()
		resp, err := RunAndGetStdOut(cmd)
		require.NoError(t, err, string(resp))
		a := assert.New(t)

		var invitations []admin.GroupInvitation
		require.NoError(t, json.Unmarshal(resp, &invitations))
		a.NotEmpty(invitations)
	})

	t.Run("Describe", func(t *testing.T) {
		cmd := exec.Command(cliPath,
			projectsEntity,
			invitationsEntity,
			"get",
			invitationID,
			"--projectId",
			g.projectID,
			"-o=json")
		cmd.Env = os.Environ()
		resp, err := RunAndGetStdOut(cmd)
		require.NoError(t, err, string(resp))
		a := assert.New(t)

		var invitation admin.GroupInvitation
		require.NoError(t, json.Unmarshal(resp, &invitation))
		a.Equal(invitationID, invitation.GetId())
		a.Equal([]string{"GROUP_READ_ONLY"}, invitation.GetRoles())
	})

	t.Run("Update by email", func(t *testing.T) {
		cmd := exec.Command(cliPath,
			projectsEntity,
			invitationsEntity,
			"update",
			"--email",
			emailProject,
			"--role",
			roleName1,
			"--role",
			roleName2,
			"--projectId",
			g.projectID,
			"-o=json")
		cmd.Env = os.Environ()
		resp, err := RunAndGetStdOut(cmd)
		require.NoError(t, err, string(resp))
		a := assert.New(t)

		var invitation admin.GroupInvitation
		require.NoError(t, json.Unmarshal(resp, &invitation))
		a.Equal(emailProject, invitation.GetUsername())
		a.ElementsMatch([]string{roleName1, roleName2}, invitation.GetRoles())
	})

	t.Run("Update by ID", func(t *testing.T) {
		cmd := exec.Command(cliPath,
			projectsEntity,
			invitationsEntity,
			"update",
			invitationID,
			"--role",
			roleName1,
			"--role",
			roleName2,
			"--projectId",
			g.projectID,
			"-o=json")
		cmd.Env = os.Environ()
		resp, err := RunAndGetStdOut(cmd)
		require.NoError(t, err, string(resp))
		a := assert.New(t)

		var invitation admin.GroupInvitation
		require.NoError(t, json.Unmarshal(resp, &invitation))
		a.Equal(emailProject, invitation.GetUsername())
		a.ElementsMatch([]string{roleName1, roleName2}, invitation.GetRoles())
	})

	t.Run("Delete", func(t *testing.T) {
		cmd := exec.Command(cliPath,
			projectsEntity,
			invitationsEntity,
			"delete",
			invitationID,
			"--force",
			"--projectId",
			g.projectID)
		cmd.Env = os.Environ()
		resp, err := RunAndGetStdOut(cmd)
		a := assert.New(t)
		require.NoError(t, err, string(resp))
		expected := fmt.Sprintf("Invitation '%s' deleted\n", invitationID)
		a.Equal(expected, string(resp))
	})
}
