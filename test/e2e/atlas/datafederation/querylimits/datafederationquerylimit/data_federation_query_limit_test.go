// Copyright 2022 MongoDB Inc
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

package datafederationquerylimit

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"testing"

	"github.com/mongodb/mongodb-atlas-cli/atlascli/test/internal"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	atlasv2 "go.mongodb.org/atlas-sdk/v20250312007/admin"
)

const (
	datafederationEntity = "datafederation"
	queryLimitsEntity    = "querylimits"
)

func TestDataFederationQueryLimit(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping test in short mode")
	}

	g := internal.NewAtlasE2ETestGenerator(t, internal.WithSnapshot())
	cliPath, err := internal.AtlasCLIBin()
	r := require.New(t)
	r.NoError(err)

	n := g.MemoryRand("rand", 1000)
	dataFederationName := fmt.Sprintf("e2e-data-federation-%v", n)
	testBucket, err := internal.TestBucketName()
	r.NoError(err)
	roleID, err := internal.CloudRoleID()
	r.NoError(err)

	limitName := "bytesProcessed.query"

	g.Run("Create Data Federation", func(t *testing.T) { //nolint:thelper // g.Run replaces t.Run
		cmd := exec.Command(cliPath,
			datafederationEntity,
			"create",
			dataFederationName,
			"--awsRoleId",
			roleID,
			"--awsTestS3Bucket",
			testBucket,
			"-o=json",
			"-P",
			internal.ProfileName(),
		)
		cmd.Env = os.Environ()
		resp, err := internal.RunAndGetStdOut(cmd)

		r.NoError(err, string(resp))

		a := assert.New(t)
		var dataLake atlasv2.DataLakeTenant
		require.NoError(t, json.Unmarshal(resp, &dataLake))
		a.Equal(dataFederationName, dataLake.GetName())
	})

	g.Run("Create", func(t *testing.T) { //nolint:thelper // g.Run replaces t.Run
		cmd := exec.Command(cliPath,
			datafederationEntity,
			queryLimitsEntity,
			"create",
			limitName,
			"--value",
			"118000000000",
			"--dataFederation",
			dataFederationName,
			"--overrunPolicy",
			"BLOCK",
			"-o=json",
			"-P",
			internal.ProfileName(),
		)
		cmd.Env = os.Environ()

		a := assert.New(t)
		resp, err := internal.RunAndGetStdOut(cmd)
		require.NoError(t, err, string(resp))
		var r atlasv2.DataFederationTenantQueryLimit
		require.NoError(t, json.Unmarshal(resp, &r))
		a.Equal(dataFederationName, *r.TenantName)
	})

	g.Run("Describe", func(t *testing.T) { //nolint:thelper // g.Run replaces t.Run
		cmd := exec.Command(cliPath,
			datafederationEntity,
			queryLimitsEntity,
			"describe",
			limitName,
			"--dataFederation",
			dataFederationName,
			"-o=json",
			"-P",
			internal.ProfileName(),
		)
		cmd.Env = os.Environ()
		resp, err := internal.RunAndGetStdOut(cmd)
		require.NoError(t, err, string(resp))
		a := assert.New(t)
		var r atlasv2.DataFederationTenantQueryLimit
		require.NoError(t, json.Unmarshal(resp, &r))
		a.Equal(limitName, r.Name)
	})

	g.Run("List", func(t *testing.T) { //nolint:thelper // g.Run replaces t.Run
		cmd := exec.Command(cliPath,
			datafederationEntity,
			queryLimitsEntity,
			"ls",
			"--dataFederation",
			dataFederationName,
			"-o=json",
			"-P",
			internal.ProfileName(),
		)
		cmd.Env = os.Environ()
		resp, err := internal.RunAndGetStdOut(cmd)

		a := assert.New(t)
		require.NoError(t, err, string(resp))
		var r []atlasv2.DataFederationTenantQueryLimit
		require.NoError(t, json.Unmarshal(resp, &r))
		a.NotEmpty(r)
	})

	g.Run("Delete", func(t *testing.T) { //nolint:thelper // g.Run replaces t.Run
		cmd := exec.Command(cliPath,
			datafederationEntity,
			queryLimitsEntity,
			"delete",
			limitName,
			"--dataFederation",
			dataFederationName,
			"--force",
			"-P",
			internal.ProfileName(),
		)
		cmd.Env = os.Environ()

		resp, err := internal.RunAndGetStdOut(cmd)
		a := assert.New(t)
		require.NoError(t, err, string(resp))
		expected := fmt.Sprintf("'%s' deleted\n", limitName)
		a.Equal(expected, string(resp))
	})

	g.Run("Delete Data Federation", func(t *testing.T) { //nolint:thelper // g.Run replaces t.Run
		cmd := exec.Command(cliPath,
			datafederationEntity,
			"delete",
			dataFederationName,
			"--force",
			"-P",
			internal.ProfileName(),
		)
		cmd.Env = os.Environ()

		resp, err := internal.RunAndGetStdOut(cmd)
		r.NoError(err, string(resp))

		expected := fmt.Sprintf("'%s' deleted\n", dataFederationName)
		assert.Equal(t, expected, string(resp))
	})
}
