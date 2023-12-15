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

	"github.com/mongodb/mongodb-atlas-cli/test/e2e"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	atlasv2 "go.mongodb.org/atlas-sdk/v20231115002/admin"
)

func TestSharedClusterUpgradeToSharedTier(t *testing.T) {
	g := newAtlasE2ETestGenerator(t)
	g.generateProject("clustersUpgrade")
	g.tier = tierM0
	g.generateCluster()
	cliPath, err := e2e.AtlasCLIBin()
	req := require.New(t)
	req.NoError(err)

	t.Run("Upgrade", func(t *testing.T) {
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
		req.NoError(err, string(resp))
	})

	t.Run("Watch upgrade", func(t *testing.T) {
		cmd := exec.Command(cliPath,
			clustersEntity,
			"watch",
			g.clusterName,
			"--projectId", g.projectID,
		)
		cmd.Env = os.Environ()
		resp, _ := cmd.CombinedOutput()
		t.Log(string(resp))
	})

	t.Run("Ensure upgrade", func(t *testing.T) {
		cmd := exec.Command(cliPath,
			clustersEntity,
			"get",
			g.clusterName,
			"--projectId", g.projectID,
			"-o=json")
		cmd.Env = os.Environ()
		resp, err := cmd.CombinedOutput()
		req.NoError(err, string(resp))
		t.Log(string(resp))
		var clusterResponse *atlasv2.AdvancedClusterDescription
		req.NoError(json.Unmarshal(resp, &clusterResponse), string(resp))
		req.NotEmpty(clusterResponse.GetReplicationSpecs())
		req.NotEmpty(clusterResponse.GetReplicationSpecs()[0].GetRegionConfigs())
		assert.Equal(t, tierM2, clusterResponse.GetReplicationSpecs()[0].GetRegionConfigs()[0].ElectableSpecs.GetInstanceSize())
	})
}

func TestSharedClusterUpgradeToDedicatedTier(t *testing.T) {
	g := newAtlasE2ETestGenerator(t)
	g.generateProject("clustersUpgrade")
	g.tier = tierM2
	g.generateCluster()
	cliPath, err := e2e.AtlasCLIBin()
	req := require.New(t)
	req.NoError(err)

	t.Run("Upgrade", func(t *testing.T) {
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
		req.NoError(err, string(resp))
	})

	t.Run("Watch upgrade", func(t *testing.T) {
		cmd := exec.Command(cliPath,
			clustersEntity,
			"watch",
			g.clusterName,
			"--projectId", g.projectID,
		)
		cmd.Env = os.Environ()
		resp, _ := cmd.CombinedOutput()
		t.Log(string(resp))
	})

	t.Run("Ensure upgrade", func(t *testing.T) {
		cmd := exec.Command(cliPath,
			clustersEntity,
			"get",
			g.clusterName,
			"--projectId", g.projectID,
			"-o=json")
		cmd.Env = os.Environ()
		resp, err := cmd.CombinedOutput()
		req.NoError(err, string(resp))

		var clusterResponse *atlasv2.AdvancedClusterDescription
		req.NoError(json.Unmarshal(resp, &clusterResponse), string(resp))
		req.NotEmpty(clusterResponse.GetReplicationSpecs())
		req.NotEmpty(clusterResponse.GetReplicationSpecs()[0].GetRegionConfigs())
		assert.Equal(t, tierM2, clusterResponse.GetReplicationSpecs()[0].GetRegionConfigs()[0].ElectableSpecs.GetInstanceSize())
	})
}
