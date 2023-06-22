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
	atlasv2 "go.mongodb.org/atlas-sdk/admin"
)

func TestAtlasUsers(t *testing.T) {
	cliPath, err := e2e.AtlasCLIBin()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	var username string
	var userID string
	var orgID string

	t.Run("List", func(t *testing.T) {
		cmd := exec.Command(cliPath,
			projectsEntity,
			usersEntity,
			"list",
			"-o=json")
		cmd.Env = os.Environ()
		resp, err := cmd.CombinedOutput()

		if err != nil {
			t.Fatalf("unexpected error: %v, resp: %v", err, string(resp))
		}

		var users atlasv2.PaginatedApiAppUser
		if err := json.Unmarshal(resp, &users); err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		if len(users.Results) == 0 {
			t.Fatalf("expected len(users) > 0, got %v", len(users.Results))
		}

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

		if err != nil {
			t.Fatalf("unexpected error: %v, resp: %v", err, string(resp))
		}

		var user atlasv2.CloudAppUser

		if err := json.Unmarshal(resp, &user); err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		assert.Equal(t, username, user.GetUsername())
		orgID = user.GetRoles()[0].GetOrgId()
	})

	t.Run("Describe by id", func(t *testing.T) {
		cmd := exec.Command(cliPath,
			usersEntity,
			"describe",
			"--id",
			userID,
			"-o=json")
		cmd.Env = os.Environ()
		resp, err := cmd.CombinedOutput()

		if err != nil {
			t.Fatalf("unexpected error: %v, resp: %v", err, string(resp))
		}

		var user atlasv2.CloudAppUser

		if err := json.Unmarshal(resp, &user); err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		assert.Equal(t, userID, user.GetId())
	})

	t.Run("Invite", func(t *testing.T) {
		n, err := e2e.RandInt(1000)
		assert.NoError(t, err)
		emailUser := fmt.Sprintf("test-%v@moongodb.com", n)
		cmd := exec.Command(cliPath,
			usersEntity,
			"invite",
			"--username", emailUser,
			"--password", "passW0rd",
			"--country", "US",
			"--email", emailUser,
			"--firstName", "TestFirstName",
			"--lastName", "TestLastName",
			"--orgRole", orgID+":ORG_MEMBER",
			"-o=json")
		cmd.Env = os.Environ()
		resp, err := cmd.CombinedOutput()
		assert.NoError(t, err, string(resp))

		var user atlasv2.CloudAppUser
		if err := json.Unmarshal(resp, &user); assert.NoError(t, err) {
			assert.Equal(t, emailUser, user.GetUsername())
			assert.NotEmpty(t, user.GetId())
			assert.Empty(t, user.GetRoles()) // This is returned empty until the invite is accepted
		}
	})
}
