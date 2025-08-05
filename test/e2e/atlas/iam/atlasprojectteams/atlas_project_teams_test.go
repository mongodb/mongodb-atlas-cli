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

package atlasprojectteams

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"testing"

	"github.com/mongodb/mongodb-atlas-cli/atlascli/test/internal"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	atlasv2 "go.mongodb.org/atlas-sdk/v20250312005/admin"
)

const (
	projectsEntity = "projects"
	teamsEntity    = "teams"

	// Roles constants.
	roleName1   = "GROUP_READ_ONLY"
	roleName2   = "GROUP_DATA_ACCESS_READ_ONLY"
	roleNameOrg = "ORG_READ_ONLY"
)

func TestAtlasProjectTeams(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping test in short mode")
	}

	g := internal.NewAtlasE2ETestGenerator(t, internal.WithSnapshot())
	cliPath, err := internal.AtlasCLIBin()
	require.NoError(t, err)

	g.GenerateProject("teams")

	n := g.MemoryRand("rand", 1000)
	teamName := fmt.Sprintf("e2e-teams-%v", n)
	teamID, err := internal.CreateTeam(teamName)
	require.NoError(t, err)
	t.Cleanup(func() {
		e := internal.DeleteTeam(teamID)
		require.NoError(t, e)
	})

	g.Run("Add", func(t *testing.T) { //nolint:thelper // g.Run replaces t.Run
		cmd := exec.Command(cliPath,
			projectsEntity,
			teamsEntity,
			"add",
			teamID,
			"--role",
			"GROUP_READ_ONLY",
			"--projectId",
			g.ProjectID,
			"-o=json")
		cmd.Env = os.Environ()
		resp, err := internal.RunAndGetStdOut(cmd)
		a := assert.New(t)
		require.NoError(t, err, string(resp))

		var teams atlasv2.PaginatedTeamRole
		require.NoError(t, json.Unmarshal(resp, &teams))
		found := false
		for _, team := range teams.GetResults() {
			if team.GetTeamId() == teamID {
				found = true
				break
			}
		}
		a.True(found)
	})

	g.Run("Update", func(t *testing.T) { //nolint:thelper // g.Run replaces t.Run
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
			g.ProjectID,
			"-o=json")
		cmd.Env = os.Environ()
		resp, err := internal.RunAndGetStdOut(cmd)
		a := assert.New(t)
		require.NoError(t, err, string(resp))

		var roles atlasv2.PaginatedTeamRole
		require.NoError(t, json.Unmarshal(resp, &roles))
		a.Len(roles.GetResults(), 1)

		role := roles.GetResults()[0]
		a.Equal(teamID, role.GetTeamId())
		a.Len(role.GetRoleNames(), 2)
		a.ElementsMatch([]string{roleName1, roleName2}, role.GetRoleNames())
	})

	g.Run("List", func(t *testing.T) { //nolint:thelper // g.Run replaces t.Run
		cmd := exec.Command(cliPath,
			projectsEntity,
			teamsEntity,
			"ls",
			"--projectId",
			g.ProjectID,
			"-o=json")
		cmd.Env = os.Environ()
		resp, err := internal.RunAndGetStdOut(cmd)
		a := assert.New(t)
		require.NoError(t, err, string(resp))

		var teams atlasv2.PaginatedTeamRole
		require.NoError(t, json.Unmarshal(resp, &teams))
		a.NotEmpty(teams.GetResults())
	})

	g.Run("Delete", func(t *testing.T) { //nolint:thelper // g.Run replaces t.Run
		cmd := exec.Command(cliPath,
			projectsEntity,
			teamsEntity,
			"delete",
			teamID,
			"--force",
			"--projectId",
			g.ProjectID)
		cmd.Env = os.Environ()
		resp, err := internal.RunAndGetStdOut(cmd)
		a := assert.New(t)
		require.NoError(t, err, string(resp))
		expected := fmt.Sprintf("Team '%s' deleted\n", teamID)
		a.Equal(expected, string(resp))
	})
}
