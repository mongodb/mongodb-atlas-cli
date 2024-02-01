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

//go:build e2e || (iam && !om60 && !atlas)

package iam_test

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"testing"

	"github.com/andreaangiolillo/mongocli-test/test/e2e"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.mongodb.org/atlas/mongodbatlas"
)

func TestProjectTeams(t *testing.T) {
	cliPath, err := e2e.Bin()
	require.NoError(t, err)

	n, err := e2e.RandInt(1000)
	require.NoError(t, err)

	projectName := fmt.Sprintf("e2e-proj-%v", n)
	projectID, err := e2e.CreateProject(projectName)
	require.NoError(t, err)
	t.Cleanup(func() {
		e2e.DeleteProjectWithRetry(t, projectID)
	})

	teamName := fmt.Sprintf("e2e-teams-%v", n)
	teamID, err := createTeam(teamName)
	require.NoError(t, err)
	t.Cleanup(func() {
		e := deleteTeam(teamID)
		require.NoError(t, e)
	})

	t.Run("Add", func(t *testing.T) {
		cmd := exec.Command(cliPath,
			iamEntity,
			projectsEntity,
			teamsEntity,
			"add",
			teamID,
			"--role",
			"GROUP_READ_ONLY",
			"--projectId",
			projectID,
			"-o=json")
		cmd.Env = os.Environ()
		resp, err := cmd.CombinedOutput()
		a := assert.New(t)
		require.NoError(t, err, string(resp))

		var teams mongodbatlas.TeamsAssigned
		require.NoError(t, json.Unmarshal(resp, &teams))
		found := false
		for _, team := range teams.Results {
			if team.TeamID == teamID {
				found = true
				break
			}
		}
		a.True(found)
	})

	t.Run("Update", func(t *testing.T) {
		roleName1 := "GROUP_READ_ONLY"
		roleName2 := "GROUP_DATA_ACCESS_READ_ONLY"
		cmd := exec.Command(cliPath,
			iamEntity,
			projectsEntity,
			teamsEntity,
			"update",
			teamID,
			"--role",
			roleName1,
			"--role",
			roleName2,
			"--projectId",
			projectID,
			"-o=json")
		cmd.Env = os.Environ()
		resp, err := cmd.CombinedOutput()
		a := assert.New(t)
		require.NoError(t, err, string(resp))

		var roles []mongodbatlas.TeamRoles
		require.NoError(t, json.Unmarshal(resp, &roles))
		a.Len(roles, 1)

		role := roles[0]
		a.Equal(teamID, role.TeamID)
		a.Len(role.RoleNames, 2)
		a.ElementsMatch([]string{roleName1, roleName2}, role.RoleNames)
	})

	t.Run("List", func(t *testing.T) {
		cmd := exec.Command(cliPath,
			iamEntity,
			projectsEntity,
			teamsEntity,
			"ls",
			"--projectId",
			projectID,
			"-o=json")
		cmd.Env = os.Environ()
		resp, err := cmd.CombinedOutput()
		a := assert.New(t)
		require.NoError(t, err, string(resp))

		var teams mongodbatlas.TeamsAssigned
		require.NoError(t, json.Unmarshal(resp, &teams))
		a.NotEmpty(teams.Results)
	})

	t.Run("Delete", func(t *testing.T) {
		cmd := exec.Command(cliPath,
			iamEntity,
			projectsEntity,
			teamsEntity,
			"delete",
			teamID,
			"--force",
			"--projectId",
			projectID)
		cmd.Env = os.Environ()
		resp, err := cmd.CombinedOutput()
		a := assert.New(t)
		require.NoError(t, err, string(resp))
		expected := fmt.Sprintf("Team '%s' deleted\n", teamID)
		a.Equal(expected, string(resp))
	})
}
