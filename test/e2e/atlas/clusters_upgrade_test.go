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
//go:build e2e || (atlas && clusters && upgrade)

package atlas_test

import (
	"encoding/json"
	"os"
	"os/exec"
	"testing"
	"time"

	"github.com/mongodb/mongodb-atlas-cli/test/e2e"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	atlasv2 "go.mongodb.org/atlas-sdk/v20231115002/admin"
)

func TestSharedClusterUpgrade(t *testing.T) {
	g := newAtlasE2ETestGenerator(t)
	g.generateProject("clustersUpgrade")
	g.tier = tierM0
	g.generateCluster()
	cliPath, err := e2e.AtlasCLIBin()
	require.NoError(t, err)

	t.Run("Upgrade to shared cluster", func(t *testing.T) {
		cmd := exec.Command(cliPath,
			clustersEntity,
			"upgrade",
			g.clusterName,
			"--tier", tierM2,
			"--diskSizeGB=2",
			"--projectId", g.projectID,
			"--tag", "env=e2e",
			"-o=json")
		cmd.Env = os.Environ()
		resp, err := cmd.CombinedOutput()
		require.NoError(t, err, string(resp))
	})
	require.NoError(t, watchCluster(g.projectID, g.clusterName))
	ensureClusterTier(t, cliPath, g.projectID, g.clusterName, tierM2)

	t.Run("Upgrade to dedicated tier", func(t *testing.T) {
		cmd := exec.Command(cliPath,
			clustersEntity,
			"upgrade",
			g.clusterName,
			"--tier", tierM10,
			"--diskSizeGB", diskSizeGB40,
			"--mdbVersion=6.0",
			"--projectId", g.projectID,
			"--tag", "env=e2e",
			"-o=json")
		cmd.Env = os.Environ()
		resp, err := cmd.CombinedOutput()
		require.NoError(t, err, string(resp))
	})
	require.NoError(t, watchCluster(g.projectID, g.clusterName))
	ensureClusterTier(t, cliPath, g.projectID, g.clusterName, tierM10)
}

func ensureClusterTier(t *testing.T, cliPath, projectID, clusterName, expectedTier string) {
	t.Helper()
	var result string
	backoff := 1
	for attempts := 1; attempts <= maxRetryAttempts; attempts++ {
		cmd := exec.Command(cliPath,
			clustersEntity,
			"get",
			clusterName,
			"--projectId", projectID,
			"-o=json")
		cmd.Env = os.Environ()
		resp, err := cmd.CombinedOutput()
		req := require.New(t)
		req.NoError(err, string(resp))
		var clusterResponse *atlasv2.AdvancedClusterDescription
		req.NoError(json.Unmarshal(resp, &clusterResponse), string(resp))
		req.NotEmpty(clusterResponse.GetReplicationSpecs())
		req.NotEmpty(clusterResponse.GetReplicationSpecs()[0].GetRegionConfigs())
		if expectedTier == clusterResponse.GetReplicationSpecs()[0].GetRegionConfigs()[0].ElectableSpecs.GetInstanceSize() {
			result = clusterResponse.GetReplicationSpecs()[0].GetRegionConfigs()[0].ElectableSpecs.GetInstanceSize()
			break
		}
		t.Logf("attempt=%d, sleeping for=%ds", attempts, backoff)
		time.Sleep(time.Duration(backoff) * time.Second)
		backoff *= 2
	}
	assert.Equal(t, expectedTier, result)
}
