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
// +build e2e atlas,generic

package atlas_test

import (
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
	scopeClusterDataLake = "Cluster0,Cluster1@CLUSTER"
	clusterName0         = "Cluster0"
	clusterName1         = "Cluster1"
	clusterType          = "CLUSTER"
)

func TestDBUsers(t *testing.T) {
	n, err := e2e.RandInt(1000)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	username := fmt.Sprintf("user-%v", n)

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
			"--password=passW0rd",
			"--scope", scopeClusterDataLake,
			"-o=json",
		)
		cmd.Env = os.Environ()
		resp, err := cmd.CombinedOutput()

		if err != nil {
			t.Fatalf("unexpected error: %v, resp: %v", err, string(resp))
		}

		var user mongodbatlas.DatabaseUser
		if err := json.Unmarshal(resp, &user); err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		assert.Equal(t, username, user.Username)
		assert.Len(t, user.Scopes, 2)
		assert.Equal(t, user.Scopes[0].Name, clusterName0)
		assert.Equal(t, user.Scopes[0].Type, clusterType)
		assert.Equal(t, user.Scopes[1].Name, clusterName1)
		assert.Equal(t, user.Scopes[0].Type, clusterType)
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
	})
}
