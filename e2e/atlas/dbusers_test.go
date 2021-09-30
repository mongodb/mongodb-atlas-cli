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
//go:build e2e || (atlas && generic)

package atlas_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"testing"
	"time"

	"github.com/mongodb/mongocli/e2e"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/atlas/mongodbatlas"
)

const (
	roleReadWrite        = "readWrite"
	userPassword         = "passW0rd"
	scopeClusterDataLake = "Cluster0,Cluster1:CLUSTER"
	clusterName0         = "Cluster0"
	clusterName1         = "Cluster1"
	clusterType          = "CLUSTER"
)

func generateUsername() (string, error) {
	n, err := e2e.RandInt(1000)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("user-%v", n), nil
}

func TestDBUsers(t *testing.T) {
	username, err := generateUsername()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	stdinUsername, err := generateUsername()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	cliPath, err := e2e.Bin()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	t.Run("Create", func(t *testing.T) {
		cmd := exec.Command(cliPath,
			atlasEntity,
			dbusersEntity,
			"create",
			"atlasAdmin",
			"--deleteAfter", time.Now().AddDate(0, 0, 1).Format(time.RFC3339),
			"--username", username,
			"--password", userPassword,
			"--scope", scopeClusterDataLake,
			"-o=json",
		)

		testCreatePasswordCmd(t, cmd, username)
	})

	t.Run("CreateWithPasswordFromStdin", func(t *testing.T) {
		cmd := exec.Command(cliPath,
			atlasEntity,
			dbusersEntity,
			"create",
			"atlasAdmin",
			"--deleteAfter", time.Now().AddDate(0, 0, 1).Format(time.RFC3339),
			"--username", stdinUsername,
			"--scope", scopeClusterDataLake,
			"-o=json",
		)

		passwordStdin := bytes.NewBuffer([]byte(fmt.Sprintf("%s", userPassword)))
		cmd.Stdin = passwordStdin

		testCreatePasswordCmd(t, cmd, stdinUsername)
	})

	t.Run("List", func(t *testing.T) {
		cmd := exec.Command(cliPath,
			atlasEntity,
			dbusersEntity,
			"ls",
			"-o=json")
		cmd.Env = os.Environ()
		resp, err := cmd.CombinedOutput()

		if err != nil {
			t.Fatalf("unexpected error: %v, resp: %v", err, string(resp))
		}

		var users []mongodbatlas.DatabaseUser
		if err := json.Unmarshal(resp, &users); err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if len(users) == 0 {
			t.Fatalf("expected len(users) > 0, got 0")
		}
	})

	t.Run("Describe", func(t *testing.T) {
		cmd := exec.Command(cliPath,
			atlasEntity,
			dbusersEntity,
			"describe",
			username,
			"-o=json")
		cmd.Env = os.Environ()
		resp, err := cmd.CombinedOutput()

		if err != nil {
			t.Fatalf("unexpected error: %v, resp: %v", err, string(resp))
		}

		var user mongodbatlas.DatabaseUser
		if err := json.Unmarshal(resp, &user); err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if user.Username != username {
			t.Fatalf("expected username to match %v, got %v", username, user.Username)
		}
	})

	t.Run("Update", func(t *testing.T) {
		cmd := exec.Command(cliPath,
			atlasEntity,
			dbusersEntity,
			"update",
			username,
			"--role",
			roleReadWrite,
			"--scope",
			clusterName0,
			"--password",
			userPassword,
			"-o=json")

		cmd.Env = os.Environ()
		resp, err := cmd.CombinedOutput()

		if err != nil {
			t.Fatalf("unexpected error: %v, resp: %v", err, string(resp))
		}

		var user mongodbatlas.DatabaseUser
		if err := json.Unmarshal(resp, &user); err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		a := assert.New(t)
		a.Equal(username, user.Username)
		if a.Len(user.Roles, 1) {
			a.Equal("admin", user.Roles[0].DatabaseName)
			a.Equal(roleReadWrite, user.Roles[0].RoleName)
		}

		a.Len(user.Scopes, 1)
		a.Equal(user.Scopes[0].Name, clusterName0)
		a.Equal(user.Scopes[0].Type, clusterType)
	})

	t.Run("Delete", func(t *testing.T) {
		testDeleteUser(t, cliPath, atlasEntity, dbusersEntity, username)
		testDeleteUser(t, cliPath, atlasEntity, dbusersEntity, stdinUsername)
	})
}

func testCreatePasswordCmd(t *testing.T, cmd *exec.Cmd, username string) {
	t.Helper()

	cmd.Env = os.Environ()

	resp, err := cmd.CombinedOutput()
	a := assert.New(t)
	a.NoError(err)

	var user mongodbatlas.DatabaseUser
	if err := json.Unmarshal(resp, &user); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	a.Equal(username, user.Username)
	if a.Len(user.Scopes, 2) {
		a.Equal(user.Scopes[0].Name, clusterName0)
		a.Equal(user.Scopes[0].Type, clusterType)
		a.Equal(user.Scopes[1].Name, clusterName1)
		a.Equal(user.Scopes[1].Type, clusterType)
	}
}

func testDeleteUser(t *testing.T, cliPath, atlasEntity, dbusersEntity, username string) {
	t.Helper()

	cmd := exec.Command(cliPath,
		atlasEntity,
		dbusersEntity,
		"delete",
		username,
		"--force",
		"--authDB",
		"admin")
	cmd.Env = os.Environ()
	resp, err := cmd.CombinedOutput()

	if err != nil {
		t.Fatalf("unexpected error: %v, resp: %v", err, string(resp))
	}

	expected := fmt.Sprintf("DB user '%s' deleted\n", username)
	assert.Equal(t, expected, string(resp))
}
