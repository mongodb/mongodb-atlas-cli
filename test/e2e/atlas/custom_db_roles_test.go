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
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"testing"

	"github.com/mongodb/mongodb-atlas-cli/atlascli/test/e2e"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	atlasv2 "go.mongodb.org/atlas-sdk/v20240530002/admin"
)

const (
	createPrivilege             = "UPDATE"
	findPrivilege               = "FIND"
	updatePrivilege             = "LIST_SESSIONS"
	enableShardingRole          = "enableSharding"
	enableShardingInheritedRole = "enableSharding@admin"
	readRole                    = "read"
	readInheritedRole           = "read@mydb"
)

func TestDBRoles(t *testing.T) {
	n, err := e2e.RandInt(1000)
	require.NoError(t, err)

	roleName := fmt.Sprintf("role-%v", n)

	cliPath, err := e2e.AtlasCLIBin()
	require.NoError(t, err)

	t.Run("Create", func(t *testing.T) {
		cmd := exec.Command(cliPath,
			customDBRoleEntity,
			"create",
			roleName,
			"--privilege", fmt.Sprintf("%s@db.collection,%s@db.collection2", createPrivilege, findPrivilege),
			"--inheritedRole", enableShardingInheritedRole,
			"-o=json",
		)
		cmd.Env = os.Environ()
		resp, err := e2e.RunAndGetStdOut(cmd)
		require.NoError(t, err, string(resp))

		var role atlasv2.UserCustomDBRole
		require.NoError(t, json.Unmarshal(resp, &role))

		a := assert.New(t)
		a.Equal(roleName, role.RoleName)
		a.Len(role.GetActions(), 2)
		a.ElementsMatch(
			[]string{role.GetActions()[0].Action, role.GetActions()[1].Action},
			[]string{createPrivilege, findPrivilege})
		a.Len(role.GetInheritedRoles(), 1)
		a.Equal(enableShardingRole, role.GetInheritedRoles()[0].Role)
	})

	t.Run("List", func(t *testing.T) {
		cmd := exec.Command(cliPath,
			customDBRoleEntity,
			"ls",
			"-o=json")
		cmd.Env = os.Environ()
		resp, err := e2e.RunAndGetStdOut(cmd)
		require.NoError(t, err, string(resp))

		var roles []atlasv2.UserCustomDBRole
		require.NoError(t, json.Unmarshal(resp, &roles))

		assert.NotEmpty(t, roles)
	})

	t.Run("Describe", func(t *testing.T) {
		cmd := exec.Command(cliPath,
			customDBRoleEntity,
			"describe",
			roleName,
			"-o=json")
		cmd.Env = os.Environ()
		resp, err := e2e.RunAndGetStdOut(cmd)
		require.NoError(t, err, string(resp))

		var role atlasv2.UserCustomDBRole
		require.NoError(t, json.Unmarshal(resp, &role))

		a := assert.New(t)
		a.Equal(roleName, role.RoleName)
		a.Len(role.GetActions(), 2)
		a.Equal(createPrivilege, role.GetActions()[0].Action)
		a.Len(role.GetInheritedRoles(), 1)
		a.Equal(enableShardingRole, role.GetInheritedRoles()[0].Role)
	})

	t.Run("Update with append", func(t *testing.T) {
		cmd := exec.Command(cliPath,
			customDBRoleEntity,
			"update",
			roleName,
			"--inheritedRole", readInheritedRole,
			"--privilege", updatePrivilege,
			"--privilege", createPrivilege+"@db2.collection",
			"--append",
			"-o=json")
		cmd.Env = os.Environ()
		resp, err := e2e.RunAndGetStdOut(cmd)
		require.NoError(t, err, string(resp))

		var role atlasv2.UserCustomDBRole
		require.NoError(t, json.Unmarshal(resp, &role))

		a := assert.New(t)
		a.Equal(roleName, role.RoleName)
		a.Len(role.GetActions(), 3)
		a.ElementsMatch(
			[]string{role.GetActions()[0].Action, role.GetActions()[1].Action, role.GetActions()[2].Action},
			[]string{updatePrivilege, createPrivilege, findPrivilege})
		a.ElementsMatch(
			[]string{enableShardingRole, readRole},
			[]string{role.GetInheritedRoles()[0].Role, role.GetInheritedRoles()[1].Role})
	})

	t.Run("Update", func(t *testing.T) {
		cmd := exec.Command(cliPath,
			customDBRoleEntity,
			"update",
			roleName,
			"--inheritedRole", enableShardingInheritedRole,
			"--privilege", updatePrivilege,
			"-o=json")
		cmd.Env = os.Environ()
		resp, err := e2e.RunAndGetStdOut(cmd)
		require.NoError(t, err, string(resp))

		var role atlasv2.UserCustomDBRole
		require.NoError(t, json.Unmarshal(resp, &role))

		a := assert.New(t)
		a.Equal(roleName, role.RoleName)
		a.Len(role.GetActions(), 1)
		a.Equal(updatePrivilege, role.GetActions()[0].Action)
		a.Len(role.GetInheritedRoles(), 1)
		a.Equal(enableShardingRole, role.GetInheritedRoles()[0].Role)
	})

	t.Run("Delete", func(t *testing.T) {
		cmd := exec.Command(cliPath,
			customDBRoleEntity,
			"delete",
			roleName,
			"--force")
		cmd.Env = os.Environ()
		resp, err := e2e.RunAndGetStdOut(cmd)
		require.NoError(t, err, string(resp))

		a := assert.New(t)
		expected := fmt.Sprintf("Custom database role '%s' deleted\n", roleName)
		a.Equal(expected, string(resp))
	})
}
