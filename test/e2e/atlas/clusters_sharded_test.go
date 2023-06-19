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
//go:build e2e || (atlas && clusters && sharded)

package atlas_test

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"testing"

	"github.com/mongodb/mongodb-atlas-cli/test/e2e"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	atlasv2 "go.mongodb.org/atlas-sdk/admin"
)

func TestShardedCluster(t *testing.T) {
	g := newAtlasE2ETestGenerator(t)
	g.generateProject("shardedClusters")

	cliPath, err := e2e.AtlasCLIBin()
	a := assert.New(t)
	req := require.New(t)
	req.NoError(err)

	shardedClusterName, err := RandClusterName()
	req.NoError(err)

	tier := e2eTier()
	region, err := g.newAvailableRegion(tier, e2eClusterProvider)
	req.NoError(err)

	t.Run("Create sharded cluster", func(t *testing.T) {
		cmd := exec.Command(cliPath,
			clustersEntity,
			"create",
			shardedClusterName,
			"--region", region,
			"--type=SHARDED",
			"--shards=2",
			"--members=3",
			"--tier", tier,
			"--provider", e2eClusterProvider,
			"--mdbVersion", e2eMDBVer,
			"--diskSizeGB", diskSizeGB30,
			"--projectId", g.projectID,
			"-o=json")

		cmd.Env = os.Environ()
		resp, err := cmd.CombinedOutput()
		req.NoError(err, string(resp))

		var cluster atlasv2.AdvancedClusterDescription
		err = json.Unmarshal(resp, &cluster)
		req.NoError(err)

		ensureCluster(t, &cluster, shardedClusterName, e2eMDBVer, 30, false)
	})

	t.Run("Delete sharded cluster", func(t *testing.T) {
		cmd := exec.Command(cliPath, clustersEntity, "delete", shardedClusterName, "--projectId", g.projectID, "--force")
		cmd.Env = os.Environ()
		resp, err := cmd.CombinedOutput()
		req.NoError(err, string(resp))

		expected := fmt.Sprintf("Deleting cluster '%s'", shardedClusterName)
		a.Equal(expected, string(resp))
	})

	t.Run("Watch deletion", func(t *testing.T) {
		cmd := exec.Command(cliPath,
			clustersEntity,
			"watch",
			shardedClusterName,
			"--projectId", g.projectID,
		)
		cmd.Env = os.Environ()
		// this command will fail with 404 once the cluster is deleted
		// we just need to wait for this to close the project
		resp, _ := cmd.CombinedOutput()
		t.Log(string(resp))
	})
}
