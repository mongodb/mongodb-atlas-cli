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

	"github.com/mongodb/mongodb-atlas-cli/atlascli/test/e2e"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	atlasv2 "go.mongodb.org/atlas-sdk/v20240530002/admin"
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
		resp, err := e2e.RunAndGetStdOut(cmd)
		require.NoError(t, err, string(resp))
		require.NoError(t, watchCluster(g.projectID, g.clusterName))
		cluster := fetchCluster(t, cliPath, g.projectID, g.clusterName)
		ensureClusterTier(t, cluster, tierM2)
		assert.Contains(t, cluster.GetTags(), atlasv2.ResourceTag{Key: "env", Value: "e2e"})
	})

	t.Run("Upgrade to dedicated tier", func(t *testing.T) {
		cmd := exec.Command(cliPath,
			clustersEntity,
			"upgrade",
			g.clusterName,
			"--tier", tierM10,
			"--diskSizeGB", diskSizeGB40,
			"--projectId", g.projectID,
			"--tag", "env=e2e",
			"-o=json")
		cmd.Env = os.Environ()
		resp, err := e2e.RunAndGetStdOut(cmd)
		require.NoError(t, err, string(resp))
		require.NoError(t, watchCluster(g.projectID, g.clusterName))
		cluster := fetchCluster(t, cliPath, g.projectID, g.clusterName)
		ensureClusterTier(t, cluster, tierM10)
		assert.InDelta(t, 40, cluster.GetDiskSizeGB(), 0.01)
		assert.Contains(t, cluster.GetTags(), atlasv2.ResourceTag{Key: "env", Value: "e2e"})
	})
}

func fetchCluster(t *testing.T, cliPath, projectID, clusterName string) *atlasv2.AdvancedClusterDescription {
	t.Helper()
	cmd := exec.Command(cliPath,
		clustersEntity,
		"get",
		clusterName,
		"--projectId", projectID,
		"-o=json")
	cmd.Env = os.Environ()
	resp, err := e2e.RunAndGetStdOut(cmd)
	req := require.New(t)
	req.NoError(err, string(resp))
	var c *atlasv2.AdvancedClusterDescription
	req.NoError(json.Unmarshal(resp, &c), string(resp))
	return c
}

func ensureClusterTier(t *testing.T, cluster *atlasv2.AdvancedClusterDescription, expected string) {
	t.Helper()
	req := require.New(t)
	req.NotEmpty(cluster.GetReplicationSpecs())
	req.NotEmpty(cluster.GetReplicationSpecs()[0].GetRegionConfigs())
	assert.Equal(t, expected, cluster.GetReplicationSpecs()[0].GetRegionConfigs()[0].ElectableSpecs.GetInstanceSize())
}
