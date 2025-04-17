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

package atlasteams

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"testing"

	"github.com/mongodb/mongodb-atlas-cli/atlascli/test/internal"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	atlasv2 "go.mongodb.org/atlas-sdk/v20250312001/admin"
)

const (
	teamsEntity = "teams"
)

func TestAtlasTeams(t *testing.T) {
	g := internal.NewAtlasE2ETestGenerator(t, internal.WithSnapshot())
	cliPath, err := internal.AtlasCLIBin()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	n := g.MemoryRand("rand", 1000)

	teamName := fmt.Sprintf("teams%v", n)
	var teamID string

	g.Run("Create", func(t *testing.T) { //nolint:thelper // g.Run replaces t.Run
		username, _, err := internal.OrgNUser(0)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		cmd := exec.Command(cliPath,
			teamsEntity,
			"create",
			teamName,
			"--username",
			username,
			"-o=json")
		cmd.Env = os.Environ()
		resp, err := internal.RunAndGetStdOut(cmd)

		a := assert.New(t)
		require.NoError(t, err, string(resp))

		var team atlasv2.Team
		require.NoError(t, json.Unmarshal(resp, &team))
		a.Equal(teamName, team.Name)
		teamID = team.GetId()
	})
	require.NotEmpty(t, teamID)

	g.Run("Describe By ID", func(t *testing.T) { //nolint:thelper // g.Run replaces t.Run
		cmd := exec.Command(cliPath,
			teamsEntity,
			"describe",
			"--id",
			teamID,
			"-o=json")
		cmd.Env = os.Environ()
		resp, err := internal.RunAndGetStdOut(cmd)
		require.NoError(t, err, string(resp))

		var team atlasv2.TeamResponse
		require.NoError(t, json.Unmarshal(resp, &team))
		assert.Equal(t, teamID, team.GetId())
	})

	g.Run("Describe By Name", func(t *testing.T) { //nolint:thelper // g.Run replaces t.Run
		cmd := exec.Command(cliPath,
			teamsEntity,
			"describe",
			"--name",
			teamName,
			"-o=json")
		cmd.Env = os.Environ()
		resp, err := internal.RunAndGetStdOut(cmd)
		require.NoError(t, err, string(resp))
		var team atlasv2.TeamResponse
		require.NoError(t, json.Unmarshal(resp, &team))
		assert.Equal(t, teamName, team.GetName())
	})

	g.Run("Rename", func(t *testing.T) { //nolint:thelper // g.Run replaces t.Run
		teamName += "_renamed"
		cmd := exec.Command(cliPath,
			teamsEntity,
			"rename",
			teamName,
			"--teamId",
			teamID,
			"-o=json")
		cmd.Env = os.Environ()
		resp, err := internal.RunAndGetStdOut(cmd)
		require.NoError(t, err, string(resp))

		var team atlasv2.TeamResponse
		require.NoError(t, json.Unmarshal(resp, &team))
		assert.Equal(t, teamID, team.GetId())
		assert.Equal(t, teamName, team.GetName())
	})

	g.Run("List", func(t *testing.T) { //nolint:thelper // g.Run replaces t.Run
		cmd := exec.Command(cliPath,
			teamsEntity,
			"ls",
			"-o=json")
		cmd.Env = os.Environ()
		resp, err := internal.RunAndGetStdOut(cmd)
		require.NoError(t, err, string(resp))

		var teams atlasv2.PaginatedTeam
		require.NoError(t, json.Unmarshal(resp, &teams))
		assert.NotEmpty(t, teams.Results)
	})

	g.Run("List Compact", func(t *testing.T) { //nolint:thelper // g.Run replaces t.Run
		cmd := exec.Command(cliPath,
			teamsEntity,
			"ls",
			"-c",
			"-o=json")
		cmd.Env = os.Environ()
		resp, err := internal.RunAndGetStdOut(cmd)
		require.NoError(t, err, string(resp))

		var teams []atlasv2.TeamResponse
		require.NoError(t, json.Unmarshal(resp, &teams))
		assert.NotEmpty(t, teams)
	})

	g.Run("Delete", func(t *testing.T) { //nolint:thelper // g.Run replaces t.Run
		cmd := exec.Command(cliPath,
			teamsEntity,
			"delete",
			teamID,
			"--force")
		cmd.Env = os.Environ()
		resp, err := internal.RunAndGetStdOut(cmd)
		require.NoError(t, err, string(resp))
		expected := fmt.Sprintf("Team '%s' deleted\n", teamID)
		assert.Equal(t, expected, string(resp))
	})
}
