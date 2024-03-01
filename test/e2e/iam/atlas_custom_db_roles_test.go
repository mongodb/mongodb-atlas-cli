// Copyright 2024 MongoDB Inc
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
	"github.com/stretchr/testify/require"
	atlasv2 "go.mongodb.org/atlas-sdk/v20231115007/admin"
)

func TestAtlasCustomDbRoles(t *testing.T) {
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

	role := "testDbRole"
	clusterRes := &[]atlasv2.DatabasePermittedNamespaceResource{{Cluster: true}}

	t.Run("Create", func(t *testing.T) {
		cmd := exec.Command(cliPath,
			customDBRoles,
			"create",
			role,
			"--inheritedRole",
			"read@mydb",
			"--privilege",
			"LIST_DATABASES",
			"--projectId",
			projectID,
			"-o=json")
		cmd.Env = os.Environ()
		resp, err := cmd.CombinedOutput()
		require.NoError(t, err, string(resp))
	})

	t.Run("Describe", func(t *testing.T) {
		cmd := exec.Command(cliPath,
			customDBRoles,
			"describe",
			role,
			"--projectId",
			projectID,
			"-o=json")
		cmd.Env = os.Environ()
		resp, err := cmd.CombinedOutput()
		require.NoError(t, err, string(resp))

		var customDBRole atlasv2.UserCustomDBRole
		require.NoError(t, json.Unmarshal(resp, &customDBRole), string(resp))
		expected := atlasv2.UserCustomDBRole{
			RoleName: role,
			Actions: &[]atlasv2.DatabasePrivilegeAction{
				{Action: "LIST_DATABASES", Resources: clusterRes},
			},
			InheritedRoles: &[]atlasv2.DatabaseInheritedRole{
				{Db: "mydb", Role: "read"},
			},
		}
		require.Equal(t, expected, customDBRole)
	})

	t.Run("Update", func(t *testing.T) {
		cmd := exec.Command(cliPath,
			customDBRoles,
			"update",
			role,
			"--inheritedRole",
			"readWrite@mydb",
			"--privilege",
			"GET_SHARD_MAP",
			"--privilege",
			"SHARDING_STATE",
			"--append",
			"--projectId",
			projectID,
			"-o=json")
		cmd.Env = os.Environ()
		resp, err := cmd.CombinedOutput()
		require.NoError(t, err, string(resp))
	})

	t.Run("List", func(t *testing.T) {
		cmd := exec.Command(cliPath,
			customDBRoles,
			"list",
			"--projectId",
			projectID,
			"-o=json")
		cmd.Env = os.Environ()
		resp, err := cmd.CombinedOutput()
		require.NoError(t, err, string(resp))

		var customDBRoles []atlasv2.UserCustomDBRole
		require.NoError(t, json.Unmarshal(resp, &customDBRoles), string(resp))
		require.Len(t, customDBRoles, 1)

		expectedActions := []atlasv2.DatabasePrivilegeAction{
			{Action: "GET_SHARD_MAP", Resources: clusterRes},
			{Action: "SHARDING_STATE", Resources: clusterRes},
			{Action: "LIST_DATABASES", Resources: clusterRes},
		}
		require.ElementsMatch(t, expectedActions, customDBRoles[0].GetActions())

		expectedRoles := []atlasv2.DatabaseInheritedRole{
			{Db: "mydb", Role: "readWrite"},
			{Db: "mydb", Role: "read"},
		}
		require.ElementsMatch(t, expectedRoles, customDBRoles[0].GetInheritedRoles())
	})

	t.Run("Delete", func(t *testing.T) {
		cmd := exec.Command(cliPath,
			customDBRoles,
			"delete",
			role,
			"--force",
			"--projectId",
			projectID)
		cmd.Env = os.Environ()
		resp, err := cmd.CombinedOutput()
		require.NoError(t, err, string(resp))
	})
}
