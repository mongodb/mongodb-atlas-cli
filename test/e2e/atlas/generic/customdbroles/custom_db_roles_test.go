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

package customdbroles

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"slices"
	"testing"

	"github.com/mongodb/mongodb-atlas-cli/atlascli/test/internal"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	atlasv2 "go.mongodb.org/atlas-sdk/v20250312007/admin"
)

const (
	customDBRoleEntity = "customDbRoles"
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
	if testing.Short() {
		t.Skip("skipping test in short mode")
	}

	g := internal.NewAtlasE2ETestGenerator(t, internal.WithSnapshot())
	n := g.MemoryRand("rand", 1000)

	roleName := fmt.Sprintf("role-%v", n)

	cliPath, err := internal.AtlasCLIBin()
	require.NoError(t, err)

	g.Run("Create", func(t *testing.T) { //nolint:thelper // g.Run replaces t.Run
		cmd := exec.Command(cliPath,
			customDBRoleEntity,
			"create",
			roleName,
			"--privilege", fmt.Sprintf("%s@db.collection,%s@db.collection2", createPrivilege, findPrivilege),
			"--inheritedRole", enableShardingInheritedRole,
			"-o=json",
			"-P",
			internal.ProfileName(),
		)
		cmd.Env = os.Environ()
		resp, err := internal.RunAndGetStdOut(cmd)
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

	g.Run("List", func(t *testing.T) { //nolint:thelper // g.Run replaces t.Run
		cmd := exec.Command(cliPath,
			customDBRoleEntity,
			"ls",
			"-o=json",
			"-P",
			internal.ProfileName())
		cmd.Env = os.Environ()
		resp, err := internal.RunAndGetStdOut(cmd)
		require.NoError(t, err, string(resp))

		var roles []atlasv2.UserCustomDBRole
		require.NoError(t, json.Unmarshal(resp, &roles))

		assert.NotEmpty(t, roles)
	})

	g.Run("Describe", func(t *testing.T) { //nolint:thelper // g.Run replaces t.Run
		cmd := exec.Command(cliPath,
			customDBRoleEntity,
			"describe",
			roleName,
			"-o=json",
			"-P",
			internal.ProfileName())
		cmd.Env = os.Environ()
		resp, err := internal.RunAndGetStdOut(cmd)
		require.NoError(t, err, string(resp))

		var role atlasv2.UserCustomDBRole
		require.NoError(t, json.Unmarshal(resp, &role))

		a := assert.New(t)
		a.Equal(roleName, role.RoleName)
		a.Len(role.GetActions(), 2)
		got := []string{role.GetActions()[0].Action, role.GetActions()[1].Action}
		slices.Sort(got)
		a.Equal(createPrivilege, got[1])
		a.Len(role.GetInheritedRoles(), 1)
		a.Equal(enableShardingRole, role.GetInheritedRoles()[0].Role)
	})

	g.Run("Update with append", func(t *testing.T) { //nolint:thelper // g.Run replaces t.Run
		cmd := exec.Command(cliPath,
			customDBRoleEntity,
			"update",
			roleName,
			"--inheritedRole", readInheritedRole,
			"--privilege", updatePrivilege,
			"--privilege", createPrivilege+"@db2.collection",
			"--append",
			"-o=json",
			"-P",
			internal.ProfileName())
		cmd.Env = os.Environ()
		resp, err := internal.RunAndGetStdOut(cmd)
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

	g.Run("Update", func(t *testing.T) { //nolint:thelper // g.Run replaces t.Run
		cmd := exec.Command(cliPath,
			customDBRoleEntity,
			"update",
			roleName,
			"--inheritedRole", enableShardingInheritedRole,
			"--privilege", updatePrivilege,
			"-o=json",
			"-P",
			internal.ProfileName())
		cmd.Env = os.Environ()
		resp, err := internal.RunAndGetStdOut(cmd)
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

	g.Run("Delete", func(t *testing.T) { //nolint:thelper // g.Run replaces t.Run
		cmd := exec.Command(cliPath,
			customDBRoleEntity,
			"delete",
			roleName,
			"--force",
			"-P",
			internal.ProfileName())
		cmd.Env = os.Environ()
		resp, err := internal.RunAndGetStdOut(cmd)
		require.NoError(t, err, string(resp))

		a := assert.New(t)
		expected := fmt.Sprintf("Custom database role '%s' deleted\n", roleName)
		a.Equal(expected, string(resp))
	})
}
