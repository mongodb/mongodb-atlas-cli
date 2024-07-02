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
//go:build e2e || (iam && atlas)

package iam_test

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"testing"

	"github.com/mongodb/mongodb-atlas-cli/atlascli/test/e2e"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.mongodb.org/atlas-sdk/v20240530002/admin"
)

func TestAtlasOrgInvitations(t *testing.T) {
	cliPath, err := e2e.AtlasCLIBin()
	require.NoError(t, err)

	n, err := e2e.RandInt(1000)
	require.NoError(t, err)

	emailOrg := fmt.Sprintf("test-%v@mongodb.com", n)
	var orgInvitationID string

	t.Run("Invite", func(t *testing.T) {
		cmd := exec.Command(cliPath,
			orgEntity,
			invitationsEntity,
			"invite",
			emailOrg,
			"--role",
			"ORG_MEMBER",
			"-o=json")
		cmd.Env = os.Environ()
		resp, err := e2e.RunAndGetStdOut(cmd)
		a := assert.New(t)
		require.NoError(t, err, string(resp))

		var invitation admin.OrganizationInvitation
		require.NoError(t, json.Unmarshal(resp, &invitation))
		a.Equal(emailOrg, invitation.GetUsername())
		require.NotEmpty(t, invitation.GetId())
		orgInvitationID = invitation.GetId()
	})

	t.Run("List", func(t *testing.T) {
		cmd := exec.Command(cliPath,
			orgEntity,
			invitationsEntity,
			"ls",
			"-o=json")
		cmd.Env = os.Environ()
		resp, err := e2e.RunAndGetStdOut(cmd)
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
			"-o=json")
		cmd.Env = os.Environ()
		resp, err := e2e.RunAndGetStdOut(cmd)
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
			"-o=json")
		cmd.Env = os.Environ()
		resp, err := e2e.RunAndGetStdOut(cmd)
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
			"-o=json")
		cmd.Env = os.Environ()
		resp, err := e2e.RunAndGetStdOut(cmd)
		require.NoError(t, err, string(resp))

		a := assert.New(t)
		var invitation admin.OrganizationInvitation
		require.NoError(t, json.Unmarshal(resp, &invitation))
		a.Equal(emailOrg, invitation.GetUsername())
		a.ElementsMatch([]string{roleNameOrg}, invitation.GetRoles())
	})

	t.Run("Delete", func(t *testing.T) {
		cmd := exec.Command(cliPath,
			orgEntity,
			invitationsEntity,
			"delete",
			orgInvitationID,
			"--force")
		cmd.Env = os.Environ()
		resp, err := e2e.RunAndGetStdOut(cmd)
		a := assert.New(t)
		require.NoError(t, err, string(resp))
		expected := fmt.Sprintf("Invitation '%s' deleted\n", orgInvitationID)
		a.Equal(expected, string(resp))
	})
}
