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

package clustersupgrade

import (
	"encoding/json"
	"os"
	"os/exec"
	"testing"

	"github.com/mongodb/mongodb-atlas-cli/atlascli/test/internal"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	atlasClustersPinned "go.mongodb.org/atlas-sdk/v20240530005/admin"
)

const (
	clustersEntity = "clusters"
	tierM10        = "M10"
	tierM0         = "M0"
	diskSizeGB40   = "40"
)

func TestSharedClusterUpgrade(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping test in short mode")
	}

	g := internal.NewAtlasE2ETestGenerator(t, internal.WithSnapshot())
	g.GenerateProject("clustersUpgrade")
	g.Tier = tierM0
	g.GenerateCluster()
	cliPath, err := internal.AtlasCLIBin()
	require.NoError(t, err)

	g.Run("Upgrade to dedicated tier", func(t *testing.T) { //nolint:thelper // g.Run replaces t.Run
		cmd := exec.Command(cliPath,
			clustersEntity,
			"upgrade",
			g.ClusterName,
			"--tier", tierM10,
			"--diskSizeGB", diskSizeGB40,
			"--projectId", g.ProjectID,
			"--tag", "env=e2e",
			"-o=json",
			"-P",
			internal.ProfileName(),
		)
		cmd.Env = os.Environ()
		resp, err := internal.RunAndGetStdOut(cmd)
		require.NoError(t, err, string(resp))
		require.NoError(t, internal.WatchCluster(g.ProjectID, g.ClusterName))
		cluster := fetchCluster(t, cliPath, g.ProjectID, g.ClusterName)
		ensureClusterTier(t, cluster, tierM10)
		assert.InDelta(t, 40, cluster.GetDiskSizeGB(), 0.01)
		assert.Contains(t, cluster.GetTags(), atlasClustersPinned.ResourceTag{Key: "env", Value: "e2e"})
	})
}

func fetchCluster(t *testing.T, cliPath, projectID, clusterName string) *atlasClustersPinned.AdvancedClusterDescription {
	t.Helper()
	cmd := exec.Command(cliPath, //nolint:gosec // needed for e2e tests
		clustersEntity,
		"get",
		clusterName,
		"--projectId", projectID,
		"-o=json",
		"-P",
		internal.ProfileName(),
	)
	cmd.Env = os.Environ()
	resp, err := internal.RunAndGetStdOut(cmd)
	req := require.New(t)
	req.NoError(err, string(resp))
	var c *atlasClustersPinned.AdvancedClusterDescription
	req.NoError(json.Unmarshal(resp, &c), string(resp))
	return c
}

func ensureClusterTier(t *testing.T, cluster *atlasClustersPinned.AdvancedClusterDescription, expected string) {
	t.Helper()
	req := require.New(t)
	req.NotEmpty(cluster.GetReplicationSpecs())
	req.NotEmpty(cluster.GetReplicationSpecs()[0].GetRegionConfigs())
	assert.Equal(t, expected, cluster.GetReplicationSpecs()[0].GetRegionConfigs()[0].ElectableSpecs.GetInstanceSize())
}
