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
//go:build e2e || (atlas && generic)

package atlas_test

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"testing"

	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli/dbusers"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/test/e2e"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	atlasv2 "go.mongodb.org/atlas-sdk/v20240530002/admin"
)

func TestDBUserCerts(t *testing.T) {
	n, err := e2e.RandInt(1000)
	require.NoError(t, err)
	username := fmt.Sprintf("user%v", n)

	cliPath, err := e2e.AtlasCLIBin()
	require.NoError(t, err)
	t.Run("Create DBUser", func(t *testing.T) {
		cmd := exec.Command(cliPath,
			dbusersEntity,
			"create",
			"atlasAdmin",
			"--username",
			username,
			"--x509Type",
			dbusers.X509TypeManaged,
			"-o=json")
		cmd.Env = os.Environ()
		resp, err := e2e.RunAndGetStdOut(cmd)
		require.NoError(t, err, string(resp))
		var user atlasv2.CloudDatabaseUser
		require.NoError(t, json.Unmarshal(resp, &user), string(resp))
		assert.Equal(t, username, user.Username)
	})

	t.Run("Create", func(t *testing.T) {
		cmd := exec.Command(cliPath,
			dbusersEntity,
			certsEntity,
			"create",
			"--username", username,
			"-o=json")
		cmd.Env = os.Environ()
		resp, err := e2e.RunAndGetStdOut(cmd)
		require.NoError(t, err, string(resp))
	})

	t.Run("List", func(t *testing.T) {
		cmd := exec.Command(cliPath,
			dbusersEntity,
			certsEntity,
			"list",
			username,
			"-o=json")
		cmd.Env = os.Environ()
		resp, err := e2e.RunAndGetStdOut(cmd)
		require.NoError(t, err, string(resp))

		var users atlasv2.PaginatedUserCert
		require.NoError(t, json.Unmarshal(resp, &users), string(resp))
		assert.NotEmpty(t, users.Results)
	})

	t.Run("Delete User", func(t *testing.T) {
		cmd := exec.Command(cliPath,
			dbusersEntity,
			"delete",
			username,
			"--force",
			"--authDB",
			"$external")
		cmd.Env = os.Environ()
		resp, err := e2e.RunAndGetStdOut(cmd)
		require.NoError(t, err, string(resp))

		expected := fmt.Sprintf("DB user '%s' deleted\n", username)
		assert.Equal(t, expected, string(resp))
	})
}
