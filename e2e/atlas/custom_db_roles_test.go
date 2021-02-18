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
// +build e2e atlas,generic

package atlas_test

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"testing"

	"github.com/mongodb/mongocli/e2e"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.mongodb.org/atlas/mongodbatlas"
)

const (
	createPrivilege = "LIST_SESSIONS"
	updatePrivilege = "UPDATE"
	inheritedRole   = "enableSharding@admin"
	enableSharding  = "enableSharding"
)

func TestDBRoles(t *testing.T) {
	n, err := e2e.RandInt(1000)
	require.NoError(t, err)

	roleName := fmt.Sprintf("role-%v", n)

	cliPath, err := e2e.Bin()
	require.NoError(t, err)

	t.Run("Create", func(t *testing.T) {
		cmd := exec.Command(cliPath,
			atlasEntity,
			customDBRoleEntity,
			"create",
			roleName,
			"--privilege", createPrivilege,
			"--inheritedRole", inheritedRole,
			"-o=json",
		)
		cmd.Env = os.Environ()
		resp, err := cmd.CombinedOutput()
		require.NoError(t, err, string(resp))

		var role mongodbatlas.CustomDBRole
		require.NoError(t, json.Unmarshal(resp, &role))

		a := assert.New(t)
		a.Equal(roleName, role.RoleName)
		a.Len(role.Actions, 1)
		a.Equal(createPrivilege, role.Actions[0].Action)
		a.Len(role.InheritedRoles, 1)
		a.Equal(enableSharding, role.InheritedRoles[0].Role)
	})

	t.Run("List", func(t *testing.T) {
		cmd := exec.Command(cliPath,
			atlasEntity,
			customDBRoleEntity,
			"ls",
			"-o=json")
		cmd.Env = os.Environ()
		resp, err := cmd.CombinedOutput()
		require.NoError(t, err, string(resp))

		var roles []mongodbatlas.CustomDBRole
		require.NoError(t, json.Unmarshal(resp, &roles))

		assert.NotEmpty(t, roles)
	})

	t.Run("Describe", func(t *testing.T) {
		cmd := exec.Command(cliPath,
			atlasEntity,
			customDBRoleEntity,
			"describe",
			roleName,
			"-o=json")
		cmd.Env = os.Environ()
		resp, err := cmd.CombinedOutput()
		require.NoError(t, err, string(resp))

		var role mongodbatlas.CustomDBRole
		require.NoError(t, json.Unmarshal(resp, &role))

		a := assert.New(t)
		a.Equal(roleName, role.RoleName)
		a.Len(role.Actions, 1)
		a.Equal(createPrivilege, role.Actions[0].Action)
		a.Len(role.InheritedRoles, 1)
		a.Equal(enableSharding, role.InheritedRoles[0].Role)
	})

	t.Run("Update with append", func(t *testing.T) {
		cmd := exec.Command(cliPath,
			atlasEntity,
			customDBRoleEntity,
			"update",
			roleName,
			"--privilege", fmt.Sprintf("%s@db", updatePrivilege),
			"--privilege", fmt.Sprintf("%s@db.collection", updatePrivilege),
			"--append",
			"-o=json")
		cmd.Env = os.Environ()
		resp, err := cmd.CombinedOutput()
		require.NoError(t, err, string(resp))

		var role mongodbatlas.CustomDBRole
		require.NoError(t, json.Unmarshal(resp, &role))

		a := assert.New(t)
		a.Equal(roleName, role.RoleName)
		a.Len(role.Actions, 2)
		a.Equal(updatePrivilege, role.Actions[0].Action)
		a.Equal(createPrivilege, role.Actions[1].Action)
		a.Len(role.InheritedRoles, 1)
		a.Equal(enableSharding, role.InheritedRoles[0].Role)
	})

	t.Run("Update", func(t *testing.T) {
		cmd := exec.Command(cliPath,
			atlasEntity,
			customDBRoleEntity,
			"update",
			roleName,
			"--privilege", fmt.Sprintf("%s@db", updatePrivilege),
			"-o=json")
		cmd.Env = os.Environ()
		resp, err := cmd.CombinedOutput()
		require.NoError(t, err, string(resp))

		var role mongodbatlas.CustomDBRole
		require.NoError(t, json.Unmarshal(resp, &role))

		a := assert.New(t)
		a.Equal(roleName, role.RoleName)
		a.Len(role.Actions, 1)
		a.Equal(updatePrivilege, role.Actions[0].Action)
	})

	t.Run("Delete", func(t *testing.T) {
		cmd := exec.Command(cliPath,
			atlasEntity,
			customDBRoleEntity,
			"delete",
			roleName,
			"--force")
		cmd.Env = os.Environ()
		resp, err := cmd.CombinedOutput()
		require.NoError(t, err, string(resp))

		a := assert.New(t)
		expected := fmt.Sprintf("Custom database role '%s' deleted\n", roleName)
		a.Equal(expected, string(resp))
	})
}
