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

//go:build e2e || (iam && !om50 && !om60 && atlas)

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
	atlasv2 "go.mongodb.org/atlas-sdk/admin"
)

func TestAtlasProjectTeams(t *testing.T) {
	cliPath, err := e2e.AtlasCLIBin()
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
		a.NoError(err, string(resp))

		var teams atlasv2.PaginatedTeamRole
		if err := json.Unmarshal(resp, &teams); a.NoError(err) {
			found := false
			for _, team := range teams.GetResults() {
				if team.GetTeamId() == teamID {
					found = true
					break
				}
			}
			a.True(found)
		}
	})

	t.Run("Update", func(t *testing.T) {
		cmd := exec.Command(cliPath,
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
		a.NoError(err, string(resp))

		var roles atlasv2.PaginatedTeamRole
		if err = json.Unmarshal(resp, &roles); a.NoError(err) {
			a.Len(roles.Results, 1)

			role := roles.Results[0]
			a.Equal(teamID, role.GetTeamId())
			a.Len(role.RoleNames, 2)
			a.ElementsMatch([]string{roleName1, roleName2}, role.RoleNames)
		}
	})

	t.Run("List", func(t *testing.T) {
		cmd := exec.Command(cliPath,
			projectsEntity,
			teamsEntity,
			"ls",
			"--projectId",
			projectID,
			"-o=json")
		cmd.Env = os.Environ()
		resp, err := cmd.CombinedOutput()
		a := assert.New(t)
		a.NoError(err, string(resp))

		var teams atlasv2.PaginatedTeamRole
		if err = json.Unmarshal(resp, &teams); a.NoError(err) {
			a.NotEmpty(teams.GetResults())
		}
	})

	t.Run("Delete", func(t *testing.T) {
		cmd := exec.Command(cliPath,
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
		if a.NoError(err, string(resp)) {
			expected := fmt.Sprintf("Team '%s' deleted\n", teamID)
			a.Equal(expected, string(resp))
		}
	})
}
