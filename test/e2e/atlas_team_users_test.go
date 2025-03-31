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

//go:build e2e || e2eSnap || (iam && atlas)

package e2e_test

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	atlasv2 "go.mongodb.org/atlas-sdk/v20250312001/admin"
)

func TestAtlasTeamUsers(t *testing.T) {
	g := newAtlasE2ETestGenerator(t, withSnapshot())
	cliPath, err := AtlasCLIBin()
	require.NoError(t, err)

	n := g.memoryRand("rand", 1000)

	teamName := fmt.Sprintf("teams%v", n)
	teamID, err := createTeam(teamName)
	require.NoError(t, err)
	defer func() {
		if e := deleteTeam(teamID); e != nil {
			t.Errorf("error deleting project: %v", e)
		}
	}()

	username, userID, err := OrgNUser(1)
	require.NoError(t, err)

	g.Run("Add", func(t *testing.T) { //nolint:thelper // g.Run replaces t.Run
		cmd := exec.Command(cliPath,
			teamsEntity,
			usersEntity,
			"add",
			userID,
			"--teamId",
			teamID,
			"-o=json",
		)
		cmd.Env = append(os.Environ(), "GOCOVERDIR="+os.Getenv("BINGOCOVERDIR"))
		resp, err := RunAndGetStdOut(cmd)
		require.NoError(t, err, string(resp))
		a := assert.New(t)

		var users atlasv2.PaginatedApiAppUser
		require.NoError(t, json.Unmarshal(resp, &users))
		found := false
		for _, user := range users.GetResults() {
			if user.Username == username {
				found = true
				break
			}
		}
		a.True(found)
	})

	g.Run("List", func(t *testing.T) { //nolint:thelper // g.Run replaces t.Run
		cmd := exec.Command(cliPath,
			teamsEntity,
			usersEntity,
			"ls",
			"--teamId",
			teamID,
			"-o=json")
		cmd.Env = append(os.Environ(), "GOCOVERDIR="+os.Getenv("BINGOCOVERDIR"))
		resp, err := RunAndGetStdOut(cmd)
		require.NoError(t, err, string(resp))
		a := assert.New(t)
		var teams atlasv2.PaginatedOrgGroup
		require.NoError(t, json.Unmarshal(resp, &teams))
		a.NotEmpty(teams.Results)
	})

	g.Run("List Compact", func(t *testing.T) { //nolint:thelper // g.Run replaces t.Run
		cmd := exec.Command(cliPath,
			teamsEntity,
			usersEntity,
			"ls",
			"-c",
			"--teamId",
			teamID,
			"-o=json")
		cmd.Env = append(os.Environ(), "GOCOVERDIR="+os.Getenv("BINGOCOVERDIR"))
		resp, err := RunAndGetStdOut(cmd)
		require.NoError(t, err, string(resp))
		a := assert.New(t)
		var teams []atlasv2.PaginatedOrgGroup
		require.NoError(t, json.Unmarshal(resp, &teams))
		a.NotEmpty(teams)
	})

	g.Run("Delete", func(t *testing.T) { //nolint:thelper // g.Run replaces t.Run
		cmd := exec.Command(cliPath,
			teamsEntity,
			usersEntity,
			"delete",
			userID,
			"--teamId",
			teamID,
			"--force")
		cmd.Env = append(os.Environ(), "GOCOVERDIR="+os.Getenv("BINGOCOVERDIR"))
		resp, err := RunAndGetStdOut(cmd)
		require.NoError(t, err, string(resp))

		a := assert.New(t)
		require.NoError(t, err, string(resp))
		expected := fmt.Sprintf("User '%s' deleted from the team\n", userID)
		a.Equal(expected, string(resp))
	})
}
