// Copyright 2020 MongoDB Inc
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

	"github.com/mongodb/mongodb-atlas-cli/test/e2e"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	atlasv2 "go.mongodb.org/atlas-sdk/v20231001002/admin"
)

func TestAtlasUsers(t *testing.T) {
	cliPath, err := e2e.AtlasCLIBin()
	require.NoError(t, err)
	var (
		username string
		userID   string
		orgID    string
	)

	t.Run("List", func(t *testing.T) {
		cmd := exec.Command(cliPath,
			projectsEntity,
			usersEntity,
			"list",
			"-o=json")
		cmd.Env = os.Environ()
		resp, err := cmd.CombinedOutput()
		require.NoError(t, err, string(resp))

		var users atlasv2.PaginatedApiAppUser
		require.NoError(t, json.Unmarshal(resp, &users), string(resp))
		require.NotEmpty(t, users.Results)
		username = users.Results[0].GetUsername()
		userID = users.Results[0].GetId()
	})

	t.Run("Describe by username", func(t *testing.T) {
		cmd := exec.Command(cliPath,
			usersEntity,
			"describe",
			"--username",
			username,
			"-o=json")
		cmd.Env = os.Environ()
		resp, err := cmd.CombinedOutput()
		require.NoError(t, err, string(resp))

		var user atlasv2.CloudAppUser
		require.NoError(t, json.Unmarshal(resp, &user), string(resp))
		assert.Equal(t, username, user.GetUsername())
		for i, item := range user.GetRoles() {
			if item.HasOrgId() {
				orgID = user.Roles[i].GetOrgId()
				break
			}
		}
		require.NotEmpty(t, orgID)
	})

	t.Run("Describe by ID", func(t *testing.T) {
		cmd := exec.Command(cliPath,
			usersEntity,
			"describe",
			"--id",
			userID,
			"-o=json")
		cmd.Env = os.Environ()
		resp, err := cmd.CombinedOutput()
		require.NoError(t, err, string(resp))

		var user atlasv2.CloudAppUser
		require.NoError(t, json.Unmarshal(resp, &user), string(resp))
		assert.Equal(t, userID, user.GetId())
	})

	t.Run("Invite", func(t *testing.T) {
		n, err := e2e.RandInt(10000)
		require.NoError(t, err)
		emailUser := fmt.Sprintf("test-%v@moongodb.com", n)
		t.Log("emailUser", emailUser, "orgID", orgID)
		cmd := exec.Command(cliPath,
			usersEntity,
			"invite",
			"--username", emailUser,
			"--password", fmt.Sprintf("~passW0rd%v~", n),
			"--country", "US",
			"--email", emailUser,
			"--firstName", "TestFirstName",
			"--lastName", "TestLastName",
			"--orgRole", orgID+":ORG_MEMBER",
			"-o=json")
		cmd.Env = os.Environ()
		resp, err := cmd.CombinedOutput()
		require.NoError(t, err, string(resp))

		var user atlasv2.CloudAppUser
		require.NoError(t, json.Unmarshal(resp, &user))
		assert.Equal(t, emailUser, user.GetUsername())
		assert.NotEmpty(t, user.GetId())
		assert.Empty(t, user.GetRoles()) // This is returned empty until the invite is accepted
	})
}
