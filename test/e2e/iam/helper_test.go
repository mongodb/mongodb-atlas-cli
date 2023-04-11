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
//go:build e2e || iam

package iam_test

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"strconv"

	"github.com/mongodb/mongodb-atlas-cli/test/e2e"
	"go.mongodb.org/atlas/mongodbatlas"
)

const (
	iamEntity              = "iam"
	orgEntity              = "orgs"
	apiKeysEntity          = "apikeys"
	apiKeyAccessListEntity = "accessLists"
	usersEntity            = "users"
	projectsEntity         = "projects"
	teamsEntity            = "teams"
	invitationsEntity      = "invitations"
)

const (
	roleName1   = "GROUP_READ_ONLY"
	roleName2   = "GROUP_DATA_ACCESS_READ_ONLY"
	roleNameOrg = "ORG_READ_ONLY"
)

var errNoAPIKey = errors.New("the apiKey ID is empty")

func createOrgAPIKey() (string, error) {
	cliPath, err := e2e.Bin()
	if err != nil {
		return "", err
	}

	cmd := exec.Command(cliPath, iamEntity,
		orgEntity,
		apiKeysEntity,
		"create",
		"--desc=e2e-test-helper",
		"--role=ORG_READ_ONLY",
		"-o=json")
	cmd.Env = os.Environ()
	resp, err := cmd.CombinedOutput()

	if err != nil {
		return "", fmt.Errorf("%w: %s", err, string(resp))
	}

	var key mongodbatlas.APIKey
	if err := json.Unmarshal(resp, &key); err != nil {
		return "", err
	}

	if key.ID != "" {
		return key.ID, nil
	}

	return "", errNoAPIKey
}

func deleteOrgAPIKey(id string) error {
	cliPath, err := e2e.Bin()
	if err != nil {
		return err
	}
	cmd := exec.Command(cliPath,
		iamEntity,
		orgEntity,
		apiKeysEntity,
		"rm",
		id,
		"--force")
	cmd.Env = os.Environ()
	return cmd.Run()
}

func createTeam(teamName string) (string, error) {
	cliPath, err := e2e.Bin()
	if err != nil {
		return "", err
	}
	username, _, err := OrgNUser(0)

	if err != nil {
		return "", err
	}
	cmd := exec.Command(cliPath,
		iamEntity,
		teamsEntity,
		"create",
		teamName,
		"--username",
		username,
		"-o=json")
	cmd.Env = os.Environ()
	resp, err := cmd.CombinedOutput()
	if err != nil {
		return "", fmt.Errorf("%w: %s", err, string(resp))
	}

	var team mongodbatlas.Team
	if err := json.Unmarshal(resp, &team); err != nil {
		return "", err
	}

	return team.ID, nil
}

func deleteTeam(teamID string) error {
	cliPath, err := e2e.Bin()
	if err != nil {
		return err
	}
	cmd := exec.Command(cliPath,
		iamEntity,
		teamsEntity,
		"delete",
		teamID,
		"--force")
	cmd.Env = os.Environ()
	return cmd.Run()
}

var errInvalidIndex = errors.New("invalid index")

// OrgNUser returns the user at the position userIndex.
// We need to pass the userIndex because the command iam teams users add would not work
// if the user is already in the team.
func OrgNUser(n int) (username, userID string, err error) {
	cliPath, err := e2e.Bin()
	if err != nil {
		return "", "", err
	}
	cmd := exec.Command(cliPath,
		iamEntity,
		orgEntity,
		usersEntity,
		"list",
		"--limit",
		strconv.Itoa(n+1),
		"-o=json")
	cmd.Env = os.Environ()
	resp, err := cmd.CombinedOutput()
	if err != nil {
		return "", "", fmt.Errorf("error loading org users: %w (%s)", err, string(resp))
	}

	var users mongodbatlas.AtlasUsersResponse
	if err := json.Unmarshal(resp, &users); err != nil {
		return "", "", err
	}

	if len(users.Results) <= n {
		return "", "", fmt.Errorf("%w: %d for %d users", errInvalidIndex, n, len(users.Results))
	}

	return users.Results[n].Username, users.Results[n].ID, nil
}
