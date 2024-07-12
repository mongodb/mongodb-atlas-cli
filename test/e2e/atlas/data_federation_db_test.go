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
//go:build e2e || (atlas && datafederation && db)

package atlas_test

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"testing"
	"time"

	"github.com/mongodb/mongodb-atlas-cli/atlascli/test/e2e"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	atlasv2 "go.mongodb.org/atlas-sdk/v20240530002/admin"
)

func TestDataFederation(t *testing.T) {
	cliPath, err := e2e.AtlasCLIBin()
	r := require.New(t)
	r.NoError(err)

	n, err := e2e.RandInt(1000)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	dataFederationName := fmt.Sprintf("e2e-data-federation-%v", n)
	testBucket := os.Getenv("E2E_TEST_BUCKET")
	r.NotEmpty(testBucket)
	roleID := os.Getenv("E2E_CLOUD_ROLE_ID")
	r.NotEmpty(roleID)

	t.Run("Create", func(t *testing.T) {
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
		resp, err := e2e.RunAndGetStdOut(cmd)

		r.NoError(err, string(resp))

		var dataLake atlasv2.DataLakeTenant
		require.NoError(t, json.Unmarshal(resp, &dataLake))
		assert.Equal(t, dataFederationName, dataLake.GetName())
	})

	t.Run("Describe", func(t *testing.T) {
		cmd := exec.Command(cliPath,
			datafederationEntity,
			"describe",
			dataFederationName,
			"-o=json")
		cmd.Env = os.Environ()
		resp, err := e2e.RunAndGetStdOut(cmd)

		r.NoError(err, string(resp))
		var dataLake atlasv2.DataLakeTenant
		require.NoError(t, json.Unmarshal(resp, &dataLake))
		assert.Equal(t, dataFederationName, dataLake.GetName())
	})

	t.Run("List", func(t *testing.T) {
		cmd := exec.Command(cliPath,
			datafederationEntity,
			"ls",
			"-o=json")
		cmd.Env = os.Environ()
		resp, err := e2e.RunAndGetStdOut(cmd)
		r.NoError(err, string(resp))

		var r []atlasv2.DataLakeTenant
		require.NoError(t, json.Unmarshal(resp, &r))
		assert.NotEmpty(t, r)
	})

	t.Run("Update", func(t *testing.T) {
		const updateRegion = "VIRGINIA_USA"
		cmd := exec.Command(cliPath,
			datafederationEntity,
			"update",
			dataFederationName,
			"--region",
			updateRegion,
			"-o=json")
		cmd.Env = os.Environ()
		resp, err := e2e.RunAndGetStdOut(cmd)
		r.NoError(err, string(resp))

		var dataLake atlasv2.DataLakeTenant
		require.NoError(t, json.Unmarshal(resp, &dataLake))
		assert.Equal(t, updateRegion, dataLake.GetDataProcessRegion().Region)
	})

	t.Run("Log", func(t *testing.T) {
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

		resp, err := e2e.RunAndGetStdOut(cmd)
		require.NoError(t, err, string(resp))
	})

	t.Run("Delete", func(t *testing.T) {
		cmd := exec.Command(cliPath,
			datafederationEntity,
			"delete",
			dataFederationName,
			"--force")
		cmd.Env = os.Environ()

		resp, err := e2e.RunAndGetStdOut(cmd)
		r.NoError(err, string(resp))

		expected := fmt.Sprintf("'%s' deleted\n", dataFederationName)
		assert.Equal(t, expected, string(resp))
	})
}
