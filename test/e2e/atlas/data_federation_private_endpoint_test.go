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
//go:build e2e || (atlas && datafederation && privatenetwork)

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

func TestDataFederationPrivateEndpointsAWS(t *testing.T) {
	g := newAtlasE2ETestGenerator(t)
	g.generateProject("dataFederationPrivateEndpointsAWS")

	cliPath, err := e2e.AtlasCLIBin()
	require.NoError(t, err)

	n, err := e2e.RandInt(int64(8000))
	require.NoError(t, err)
	vpcID := fmt.Sprintf("vpce-0fcd9d80bbafe%d", 1000+n.Int64())

	t.Run("Create", func(t *testing.T) {
		cmd := exec.Command(cliPath,
			datafederationEntity,
			privateEndpointsEntity,
			"create",
			vpcID,
			"--comment",
			"comment",
			"--projectId",
			g.projectID,
			"-o=json")
		cmd.Env = os.Environ()

		a := assert.New(t)
		resp, err := e2e.RunAndGetStdOut(cmd)
		require.NoError(t, err, string(resp))
		var r atlasv2.PaginatedPrivateNetworkEndpointIdEntry
		require.NoError(t, json.Unmarshal(resp, &r))
		a.NotEmpty(r.Results)
		a.Equal(r.GetResults()[0].GetEndpointId(), vpcID)
	})

	t.Run("Describe", func(t *testing.T) {
		cmd := exec.Command(cliPath,
			datafederationEntity,
			privateEndpointsEntity,
			"describe",
			vpcID,
			"--projectId",
			g.projectID,
			"-o=json")
		cmd.Env = os.Environ()
		resp, err := e2e.RunAndGetStdOut(cmd)
		require.NoError(t, err, string(resp))
		a := assert.New(t)
		var r atlasv2.PrivateNetworkEndpointIdEntry
		require.NoError(t, json.Unmarshal(resp, &r))
		a.Equal(vpcID, r.GetEndpointId())
	})

	t.Run("List", func(t *testing.T) {
		cmd := exec.Command(cliPath,
			datafederationEntity,
			privateEndpointsEntity,
			"ls",
			"--projectId",
			g.projectID,
			"-o=json")
		cmd.Env = os.Environ()
		resp, err := e2e.RunAndGetStdOut(cmd)

		a := assert.New(t)
		require.NoError(t, err, string(resp))
		var r atlasv2.PaginatedPrivateNetworkEndpointIdEntry
		require.NoError(t, json.Unmarshal(resp, &r))
		a.NotEmpty(r)
	})

	t.Run("Delete", func(t *testing.T) {
		cmd := exec.Command(cliPath,
			datafederationEntity,
			privateEndpointsEntity,
			"delete",
			vpcID,
			"--projectId",
			g.projectID,
			"--force")
		cmd.Env = os.Environ()

		resp, err := e2e.RunAndGetStdOut(cmd)
		a := assert.New(t)
		require.NoError(t, err, string(resp))
		expected := fmt.Sprintf("'%s' deleted\n", vpcID)
		a.Equal(expected, string(resp))
	})
}
