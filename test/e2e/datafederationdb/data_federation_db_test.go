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

//go:build e2e || e2eSnap || (atlas && datafederation && db)

package datafederationdb

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"testing"
	"time"

	"github.com/mongodb/mongodb-atlas-cli/atlascli/test/internal"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	atlasv2 "go.mongodb.org/atlas-sdk/v20250312004/admin"
)

const (
	datafederationEntity = "datafederation"
)

func TestDataFederation(t *testing.T) {
	g := internal.NewAtlasE2ETestGenerator(t, internal.WithSnapshot())
	cliPath, err := internal.AtlasCLIBin()
	r := require.New(t)
	r.NoError(err)

	n := g.MemoryRand("rand", 1000)

	dataFederationName := fmt.Sprintf("e2e-data-federation-%v", n)
	testBucket := os.Getenv("E2E_TEST_BUCKET")
	r.NotEmpty(testBucket)
	roleID := os.Getenv("E2E_CLOUD_ROLE_ID")
	r.NotEmpty(roleID)

	g.Run("Create", func(t *testing.T) { //nolint:thelper // g.Run replaces t.Run
		cmd := exec.Command(cliPath,
			datafederationEntity,
			"create",
			dataFederationName,
			"--awsRoleId",
			roleID,
			"--awsTestS3Bucket",
			testBucket,
			"-o=json")
		cmd.Env = os.Environ()
		resp, err := internal.RunAndGetStdOut(cmd)

		r.NoError(err, string(resp))

		var dataLake atlasv2.DataLakeTenant
		require.NoError(t, json.Unmarshal(resp, &dataLake))
		assert.Equal(t, dataFederationName, dataLake.GetName())
	})

	t.Cleanup(func() {
		t.Log("cleaning up created datafederation")
		cmd := exec.Command(cliPath,
			datafederationEntity,
			"delete",
			dataFederationName,
			"--force")
		cmd.Env = os.Environ()

		// this command will only succeed in case one of the tests after this one fails
		// not printing the output, because it might cause confusion
		_ = cmd.Run()
		t.Log("finished cleaning up created datafederation")
	})

	g.Run("Describe", func(t *testing.T) { //nolint:thelper // g.Run replaces t.Run
		cmd := exec.Command(cliPath,
			datafederationEntity,
			"describe",
			dataFederationName,
			"-o=json")
		cmd.Env = os.Environ()
		resp, err := internal.RunAndGetStdOut(cmd)

		r.NoError(err, string(resp))
		var dataLake atlasv2.DataLakeTenant
		require.NoError(t, json.Unmarshal(resp, &dataLake))
		assert.Equal(t, dataFederationName, dataLake.GetName())
	})

	g.Run("List", func(t *testing.T) { //nolint:thelper // g.Run replaces t.Run
		cmd := exec.Command(cliPath,
			datafederationEntity,
			"ls",
			"-o=json")
		cmd.Env = os.Environ()
		resp, err := internal.RunAndGetStdOut(cmd)
		r.NoError(err, string(resp))

		var r []atlasv2.DataLakeTenant
		require.NoError(t, json.Unmarshal(resp, &r))
		assert.NotEmpty(t, r)
	})

	g.Run("Update", func(t *testing.T) { //nolint:thelper // g.Run replaces t.Run
		const updateRegion = "VIRGINIA_USA"
		cmd := exec.Command(cliPath,
			datafederationEntity,
			"update",
			dataFederationName,
			"--region",
			updateRegion,
			"-o=json")
		cmd.Env = os.Environ()
		resp, err := internal.RunAndGetStdOut(cmd)
		r.NoError(err, string(resp))

		var dataLake atlasv2.DataLakeTenant
		require.NoError(t, json.Unmarshal(resp, &dataLake))
		assert.Equal(t, updateRegion, dataLake.GetDataProcessRegion().Region)
	})

	g.Run("Log", func(t *testing.T) { //nolint:thelper // g.Run replaces t.Run
		cmd := exec.Command(cliPath,
			datafederationEntity,
			"logs",
			dataFederationName,
			"--out",
			"-",
			"--start",
			strconv.FormatInt(time.Now().Add(-10*time.Second).Unix(), 10),
			"--end",
			strconv.FormatInt(time.Now().Unix(), 10),
			"--force")
		cmd.Env = os.Environ()

		resp, err := internal.RunAndGetStdOut(cmd)
		require.NoError(t, err, string(resp))
	})

	g.Run("Download Logs", func(t *testing.T) { //nolint:thelper // g.Run replaces t.Run
		cmd := exec.Command(cliPath,
			datafederationEntity,
			"logs",
			dataFederationName,
			"--out",
			"testLogFile")
		cmd.Env = os.Environ()

		resp, err := internal.RunAndGetStdOut(cmd)
		require.NoError(t, err, string(resp))
	})

	g.Run("Delete", func(t *testing.T) { //nolint:thelper // g.Run replaces t.Run
		cmd := exec.Command(cliPath,
			datafederationEntity,
			"delete",
			dataFederationName,
			"--force")
		cmd.Env = os.Environ()

		resp, err := internal.RunAndGetStdOut(cmd)
		r.NoError(err, string(resp))

		expected := fmt.Sprintf("'%s' deleted\n", dataFederationName)
		assert.Equal(t, expected, string(resp))
	})
}
