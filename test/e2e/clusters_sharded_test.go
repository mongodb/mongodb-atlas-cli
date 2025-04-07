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

package e2e

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"testing"

	"github.com/mongodb/mongodb-atlas-cli/atlascli/test/internal"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	atlasClustersPinned "go.mongodb.org/atlas-sdk/v20240530005/admin"
)

func TestShardedCluster(t *testing.T) {
	g := internal.NewAtlasE2ETestGenerator(t, internal.WithSnapshot())
	g.GenerateProject("shardedClusters")

	cliPath, err := internal.AtlasCLIBin()
	req := require.New(t)
	req.NoError(err)

	shardedClusterName := g.Memory("shardedClusterName", internal.Must(internal.RandClusterName())).(string)

	tier := internal.E2eTier()
	region, err := g.NewAvailableRegion(tier, e2eClusterProvider)
	req.NoError(err)

	mdbVersion, err := internal.MongoDBMajorVersion()
	req.NoError(err)

	g.Run("Create sharded cluster", func(t *testing.T) { //nolint:thelper // g.Run replaces t.Run
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
			"--projectId", g.ProjectID,
			"-o=json")

		cmd.Env = os.Environ()
		resp, err := internal.RunAndGetStdOut(cmd)
		req.NoError(err, string(resp))

		var cluster atlasClustersPinned.AdvancedClusterDescription
		req.NoError(json.Unmarshal(resp, &cluster))

		internal.EnsureCluster(t, &cluster, shardedClusterName, mdbVersion, 30, false)
	})

	g.Run("Delete sharded cluster", func(t *testing.T) { //nolint:thelper // g.Run replaces t.Run
		cmd := exec.Command(cliPath, clustersEntity, "delete", shardedClusterName, "--projectId", g.ProjectID, "--force")
		cmd.Env = os.Environ()
		resp, err := internal.RunAndGetStdOut(cmd)
		req.NoError(err, string(resp))

		expected := fmt.Sprintf("Deleting cluster '%s'", shardedClusterName)
		assert.Equal(t, expected, string(resp))
	})

	if internal.SkipCleanup() {
		return
	}

	g.Run("Watch deletion", func(t *testing.T) { //nolint:thelper // g.Run replaces t.Run
		cmd := exec.Command(cliPath,
			clustersEntity,
			"watch",
			shardedClusterName,
			"--projectId", g.ProjectID,
		)
		cmd.Env = os.Environ()
		// this command will fail with 404 once the cluster is deleted
		// we just need to wait for this to close the project
		resp, _ := internal.RunAndGetStdOut(cmd)
		t.Log(string(resp))
	})
}
