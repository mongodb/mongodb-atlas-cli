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

package datafederationprivateendpoint

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"testing"

	"github.com/mongodb/mongodb-atlas-cli/atlascli/test/internal"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	atlasv2 "go.mongodb.org/atlas-sdk/v20250312008/admin"
)

const (
	datafederationEntity   = "datafederation"
	privateEndpointsEntity = "privateendpoints"
)

func TestDataFederationPrivateEndpointsAWS(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping test in short mode")
	}

	g := internal.NewAtlasE2ETestGenerator(t, internal.WithSnapshot())
	g.GenerateProject("dataFederationPrivateEndpointsAWS")

	cliPath, err := internal.AtlasCLIBin()
	require.NoError(t, err)

	n := g.MemoryRand("rand", int64(8000))
	vpcID := fmt.Sprintf("vpce-0fcd9d80bbafe%d", 1000+n.Int64())

	g.Run("Create", func(t *testing.T) { //nolint:thelper // g.Run replaces t.Run
		cmd := exec.Command(cliPath,
			datafederationEntity,
			privateEndpointsEntity,
			"create",
			vpcID,
			"--comment",
			"comment",
			"--projectId",
			g.ProjectID,
			"-o=json",
			"-P",
			internal.ProfileName(),
		)
		cmd.Env = os.Environ()

		a := assert.New(t)
		resp, err := internal.RunAndGetStdOut(cmd)
		require.NoError(t, err, string(resp))
		var r atlasv2.PaginatedPrivateNetworkEndpointIdEntry
		require.NoError(t, json.Unmarshal(resp, &r))
		a.NotEmpty(r.Results)
		a.Equal(r.GetResults()[0].GetEndpointId(), vpcID)
	})

	g.Run("Describe", func(t *testing.T) { //nolint:thelper // g.Run replaces t.Run
		cmd := exec.Command(cliPath,
			datafederationEntity,
			privateEndpointsEntity,
			"describe",
			vpcID,
			"--projectId",
			g.ProjectID,
			"-o=json",
			"-P",
			internal.ProfileName(),
		)
		cmd.Env = os.Environ()
		resp, err := internal.RunAndGetStdOut(cmd)
		require.NoError(t, err, string(resp))
		a := assert.New(t)
		var r atlasv2.PrivateNetworkEndpointIdEntry
		require.NoError(t, json.Unmarshal(resp, &r))
		a.Equal(vpcID, r.GetEndpointId())
	})

	g.Run("List", func(t *testing.T) { //nolint:thelper // g.Run replaces t.Run
		cmd := exec.Command(cliPath,
			datafederationEntity,
			privateEndpointsEntity,
			"ls",
			"--projectId",
			g.ProjectID,
			"-o=json",
			"-P",
			internal.ProfileName(),
		)
		cmd.Env = os.Environ()
		resp, err := internal.RunAndGetStdOut(cmd)

		a := assert.New(t)
		require.NoError(t, err, string(resp))
		var r atlasv2.PaginatedPrivateNetworkEndpointIdEntry
		require.NoError(t, json.Unmarshal(resp, &r))
		a.NotEmpty(r)
	})

	g.Run("Delete", func(t *testing.T) { //nolint:thelper // g.Run replaces t.Run
		cmd := exec.Command(cliPath,
			datafederationEntity,
			privateEndpointsEntity,
			"delete",
			vpcID,
			"--projectId",
			g.ProjectID,
			"--force",
			"-P",
			internal.ProfileName(),
		)
		cmd.Env = os.Environ()

		resp, err := internal.RunAndGetStdOut(cmd)
		a := assert.New(t)
		require.NoError(t, err, string(resp))
		expected := fmt.Sprintf("'%s' deleted\n", vpcID)
		a.Equal(expected, string(resp))
	})
}
