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

	"github.com/mongodb/mongodb-atlas-cli/atlascli/test/e2e"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	atlasv2 "go.mongodb.org/atlas-sdk/v20240530002/admin"
)

func TestShardedCluster(t *testing.T) {
	g := newAtlasE2ETestGenerator(t)
	g.generateProject("shardedClusters")

	cliPath, err := e2e.AtlasCLIBin()
	req := require.New(t)
	req.NoError(err)

	shardedClusterName, err := RandClusterName()
	req.NoError(err)

	tier := e2eTier()
	region, err := g.newAvailableRegion(tier, e2eClusterProvider)
	req.NoError(err)

	mdbVersion, err := MongoDBMajorVersion()
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
			"--mdbVersion", mdbVersion,
			"--diskSizeGB", diskSizeGB30,
			"--projectId", g.projectID,
			"-o=json")

		cmd.Env = os.Environ()
		resp, err := e2e.RunAndGetStdOut(cmd)
		req.NoError(err, string(resp))

		var cluster atlasv2.AdvancedClusterDescription
		req.NoError(json.Unmarshal(resp, &cluster))

		ensureCluster(t, &cluster, shardedClusterName, mdbVersion, 30, false)
	})

	t.Run("Delete sharded cluster", func(t *testing.T) {
		cmd := exec.Command(cliPath, clustersEntity, "delete", shardedClusterName, "--projectId", g.projectID, "--force")
		cmd.Env = os.Environ()
		resp, err := e2e.RunAndGetStdOut(cmd)
		req.NoError(err, string(resp))

		expected := fmt.Sprintf("Deleting cluster '%s'", shardedClusterName)
		assert.Equal(t, expected, string(resp))
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
		resp, _ := e2e.RunAndGetStdOut(cmd)
		t.Log(string(resp))
	})
}
