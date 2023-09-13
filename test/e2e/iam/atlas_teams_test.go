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
	atlasv2 "go.mongodb.org/atlas-sdk/v20230201007/admin"
)

func TestAtlasTeams(t *testing.T) {
	cliPath, err := e2e.AtlasCLIBin()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	n, err := e2e.RandInt(1000)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	teamName := fmt.Sprintf("teams%v", n)
	var teamID string

	t.Run("Create", func(t *testing.T) {
		username, _, err := OrgNUser(0)
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
		resp, err := cmd.CombinedOutput()

		a := assert.New(t)
		a.NoError(err, string(resp))

		var team atlasv2.Team
		if err := json.Unmarshal(resp, &team); a.NoError(err) {
			a.Equal(teamName, team.Name)
			teamID = *team.Id
		}
	})

	t.Run("Describe By Id", func(t *testing.T) {
		cmd := exec.Command(cliPath,
			teamsEntity,
			"describe",
			"--id",
			teamID,
			"-o=json")
		cmd.Env = os.Environ()
		resp, err := cmd.CombinedOutput()

		a := assert.New(t)
		a.NoError(err, string(resp))

		var team atlasv2.TeamResponse
		if err := json.Unmarshal(resp, &team); a.NoError(err) {
			a.Equal(teamID, *team.Id)
		}
	})

	t.Run("Describe By Name", func(t *testing.T) {
		cmd := exec.Command(cliPath,
			teamsEntity,
			"describe",
			"--name",
			teamName,
			"-o=json")
		cmd.Env = os.Environ()
		resp, err := cmd.CombinedOutput()

		a := assert.New(t)
		a.NoError(err, string(resp))

		var team atlasv2.TeamResponse
		if err := json.Unmarshal(resp, &team); a.NoError(err) {
			a.Equal(teamName, *team.Name)
		}
	})

	t.Run("List", func(t *testing.T) {
		cmd := exec.Command(cliPath,
			teamsEntity,
			"ls",
			"-o=json")
		cmd.Env = os.Environ()
		resp, err := cmd.CombinedOutput()

		a := assert.New(t)
		a.NoError(err, string(resp))

		var teams atlasv2.PaginatedTeam
		if err := json.Unmarshal(resp, &teams); a.NoError(err) {
			a.NotEmpty(t, teams.Results)
		}
	})

	t.Run("List Compact", func(t *testing.T) {
		cmd := exec.Command(cliPath,
			teamsEntity,
			"ls",
			"-c",
			"-o=json")
		cmd.Env = os.Environ()
		resp, err := cmd.CombinedOutput()

		a := assert.New(t)
		a.NoError(err, string(resp))

		var teams []atlasv2.TeamResponse
		if err := json.Unmarshal(resp, &teams); a.NoError(err) {
			a.NotEmpty(t, teams)
		}
	})

	t.Run("Delete", func(t *testing.T) {
		cmd := exec.Command(cliPath,
			teamsEntity,
			"delete",
			teamID,
			"--force")
		cmd.Env = os.Environ()
		resp, err := cmd.CombinedOutput()
		a := assert.New(t)
		if a.NoError(err, string(resp)) {
			expected := fmt.Sprintf("Team '%s' deleted\n", teamID)
			a.Equal(expected, string(resp))
		}
	})
}
