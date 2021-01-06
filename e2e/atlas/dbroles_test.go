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
	"go.mongodb.org/atlas/mongodbatlas"
)

const (
	find           = "FIND"
	update         = "UPDATE"
	inheritedRole  = "enableSharding@admin"
	enableSharding = "enableSharding"
)

func TestDBRoles(t *testing.T) {
	n, err := e2e.RandInt(1000)
	assert.NoError(t, err)

	roleName := fmt.Sprintf("role-%v", n)

	cliPath, err := e2e.Bin()
	assert.NoError(t, err)

	t.Run("Create", func(t *testing.T) {
		cmd := exec.Command(cliPath,
			atlasEntity,
			dbrolesEntity,
			"create",
			"--roleName", roleName,
			"--action", find,
			"--db=db",
			"--inheritedRole", inheritedRole,
			"-o=json",
		)
		cmd.Env = os.Environ()
		resp, err := cmd.CombinedOutput()
		assert.NoError(t, err)

		var role mongodbatlas.CustomDBRole
		if err := json.Unmarshal(resp, &role); err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		assert.Equal(t, roleName, role.RoleName)
		assert.Len(t, role.Actions, 1)
		assert.Equal(t, find, role.Actions[0].Action)
		assert.Len(t, role.InheritedRoles, 1)
		assert.Equal(t, enableSharding, role.InheritedRoles[0].Role)
	})

	t.Run("List", func(t *testing.T) {
		cmd := exec.Command(cliPath,
			atlasEntity,
			dbrolesEntity,
			"ls",
			"-o=json")
		cmd.Env = os.Environ()
		resp, err := cmd.CombinedOutput()
		assert.NoError(t, err)

		var roles []mongodbatlas.CustomDBRole
		if err := json.Unmarshal(resp, &roles); err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if len(*roles) == 0 {
			t.Fatalf("expected len(roles) > 0, got 0")
		}
	})

	t.Run("Describe", func(t *testing.T) {
		cmd := exec.Command(cliPath,
			atlasEntity,
			dbrolesEntity,
			"describe",
			roleName,
			"-o=json")
		cmd.Env = os.Environ()
		resp, err := cmd.CombinedOutput()
		assert.NoError(t, err)

		var role mongodbatlas.CustomDBRole
		if err := json.Unmarshal(resp, &role); err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		assert.Equal(t, roleName, role.RoleName)
		assert.Len(t, role.Actions, 1)
		assert.Equal(t, find, role.Actions[0].Action)
		assert.Len(t, role.InheritedRoles, 1)
		assert.Equal(t, enableSharding, role.InheritedRoles[0].Role)
	})

	t.Run("Update with append", func(t *testing.T) {
		cmd := exec.Command(cliPath,
			atlasEntity,
			dbrolesEntity,
			"update",
			roleName,
			"--action",
			update,
			"--db=db",
			"--append",
			"-o=json")
		cmd.Env = os.Environ()
		resp, err := cmd.CombinedOutput()
		assert.NoError(t, err)

		var role mongodbatlas.CustomDBRole
		if err := json.Unmarshal(resp, &role); err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		assert.Equal(t, roleName, role.RoleName)
		assert.Len(t, role.Actions, 2)
		assert.Equal(t, find, role.Actions[0].Action)
		assert.Equal(t, update, role.Actions[1].Action)
		assert.Len(t, role.InheritedRoles, 1)
		assert.Equal(t, enableSharding, role.InheritedRoles[0].Role)
	})

	t.Run("Update", func(t *testing.T) {
		cmd := exec.Command(cliPath,
			atlasEntity,
			dbrolesEntity,
			"update",
			roleName,
			"--action",
			update,
			"--db=db",
			"-o=json")
		cmd.Env = os.Environ()
		resp, err := cmd.CombinedOutput()
		assert.NoError(t, err)

		var role mongodbatlas.CustomDBRole
		if err := json.Unmarshal(resp, &role); err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		assert.Equal(t, roleName, role.RoleName)
		assert.Len(t, role.Actions, 1)
		assert.Equal(t, update, role.Actions[1].Action)
	})

	t.Run("Delete", func(t *testing.T) {
		cmd := exec.Command(cliPath,
			atlasEntity,
			dbrolesEntity,
			"delete",
			roleName,
			"--force")
		cmd.Env = os.Environ()
		resp, err := cmd.CombinedOutput()

		if err != nil {
			t.Fatalf("unexpected error: %v, resp: %v", err, string(resp))
		}

		expected := fmt.Sprintf("Custom Database role '%s' deleted\n", roleName)
		assert.Equal(t, expected, string(resp))
	})
}
