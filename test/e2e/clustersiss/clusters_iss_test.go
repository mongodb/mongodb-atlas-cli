// Copyright 2024 MongoDB Inc
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
//go:build e2e || e2eSnap || (atlas && clusters && iss)

package clustersiss

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"testing"

	"github.com/mongodb/mongodb-atlas-cli/atlascli/test/internal"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.mongodb.org/atlas-sdk/v20250312003/admin"
)

const (
	clustersEntity              = "clusters"
	diskSizeGB30                = "30"
	independentShardScalingFlag = "independentShardScaling"
	clusterWideScalingFlag      = "clusterWideScaling"

	// Cluster settings.
	e2eClusterProvider = "AWS"
)

func TestIndependendShardScalingCluster(t *testing.T) {
	g := internal.NewAtlasE2ETestGenerator(t, internal.WithSnapshot())
	cliPath, err := internal.AtlasCLIBin()
	req := require.New(t)
	req.NoError(err)

	g.GenerateProject("clustersIss")

	issClusterName := g.Memory("issClusterName", internal.Must(internal.RandClusterName())).(string)

	tier := internal.E2eTier()
	region, err := g.NewAvailableRegion(tier, e2eClusterProvider)
	req.NoError(err)

	mdbVersion, err := internal.MongoDBMajorVersion()
	req.NoError(err)

	g.Run("Create ISS cluster", func(t *testing.T) { //nolint:thelper // g.Run replaces t.Run
		cmd := exec.Command(cliPath,
			clustersEntity,
			"create",
			issClusterName,
			"--region", region,
			"--type=SHARDED",
			"--shards=2",
			"--members=3",
			"--tier", tier,
			"--provider", e2eClusterProvider,
			"--mdbVersion", mdbVersion,
			"--diskSizeGB", diskSizeGB30,
			"--autoScalingMode", independentShardScalingFlag,
			"--watch",
			"-o=json")

		cmd.Env = os.Environ()
		resp, err := internal.RunAndGetStdOut(cmd)
		req.NoError(err, string(resp))

		var cluster admin.ClusterDescription20240805
		req.NoError(json.Unmarshal(resp, &cluster))

		internal.EnsureClusterLatest(t, &cluster, issClusterName, mdbVersion, 30, false)
		assert.Len(t, cluster.GetReplicationSpecs(), 2)
		assert.Equal(t, "SHARDED", cluster.GetClusterType())
	})

	g.Run("Get ISS cluster", func(t *testing.T) { //nolint:thelper // g.Run replaces t.Run
		cmd := exec.Command(cliPath,
			clustersEntity,
			"get",
			issClusterName,
			"-o=json")

		cmd.Env = os.Environ()
		resp, err := internal.RunAndGetStdOut(cmd)
		req.NoError(err, string(resp))

		var cluster admin.ClusterDescription20240805
		req.NoError(json.Unmarshal(resp, &cluster))
		assert.Equal(t, "SHARDED", cluster.GetClusterType())
		assert.Len(t, cluster.GetReplicationSpecs(), 2)
	})

	g.Run("Get ISS cluster autoScalingMode", func(t *testing.T) { //nolint:thelper // g.Run replaces t.Run
		cmd := exec.Command(cliPath,
			clustersEntity,
			"autoScalingConfig",
			issClusterName,
			"-o=json")

		cmd.Env = os.Environ()
		resp, err := internal.RunAndGetStdOut(cmd)
		req.NoError(err, string(resp))

		var config admin.ClusterDescriptionAutoScalingModeConfiguration
		req.NoError(json.Unmarshal(resp, &config))
		assert.Equal(t, "INDEPENDENT_SHARD_SCALING", config.GetAutoScalingMode())
	})

	g.Run("Pause ISS cluster with the wrong flag", func(t *testing.T) { //nolint:thelper // g.Run replaces t.Run
		cmd := exec.Command(cliPath,
			clustersEntity,
			"pause",
			issClusterName,
			"--autoScalingMode", clusterWideScalingFlag,
		)
		cmd.Env = os.Environ()
		resp, err := internal.RunAndGetStdOut(cmd)
		req.NoError(err, string(resp))

		expected := fmt.Sprintf("Pausing cluster '%s'Cluster paused\n", issClusterName)
		assert.Equal(t, expected, string(resp))
	})

	g.Run("Get ISS cluster and check autoScalingMode is still independentShardScaling", func(t *testing.T) { //nolint:thelper // g.Run replaces t.Run
		cmd := exec.Command(cliPath,
			clustersEntity,
			"get",
			issClusterName,
			"--autoScalingMode", independentShardScalingFlag,
			"-o=json")

		cmd.Env = os.Environ()
		resp, err := internal.RunAndGetStdOut(cmd)
		req.NoError(err, string(resp))

		var cluster admin.ClusterDescription20240805
		req.NoError(json.Unmarshal(resp, &cluster))

		internal.EnsureClusterLatest(t, &cluster, issClusterName, mdbVersion, 30, false)
		assert.Len(t, cluster.GetReplicationSpecs(), 2)
	})

	g.Run("List ISS cluster", func(t *testing.T) { //nolint:thelper // g.Run replaces t.Run
		cmd := exec.Command(cliPath,
			clustersEntity,
			"list",
			"--autoScalingMode", independentShardScalingFlag,
			"-o=json")
		cmd.Env = os.Environ()
		resp, err := internal.RunAndGetStdOut(cmd)
		req.NoError(err, string(resp))

		var clusters admin.PaginatedClusterDescription20240805
		req.NoError(json.Unmarshal(resp, &clusters))

		assert.Positive(t, clusters.GetTotalCount())
		assert.NotEmpty(t, clusters.Results)
	})

	g.Run("Start ISS cluster", func(t *testing.T) { //nolint:thelper // g.Run replaces t.Run
		cmd := exec.Command(cliPath,
			clustersEntity,
			"start",
			issClusterName,
			"--autoScalingMode", independentShardScalingFlag,
		)
		cmd.Env = os.Environ()
		resp, err := internal.RunAndGetStdOut(cmd)
		req.NoError(err, string(resp))

		expected := fmt.Sprintf("Starting cluster '%s'Cluster started\n", issClusterName)
		assert.Equal(t, expected, string(resp))
	})

	g.Run("Delete ISS cluster", func(t *testing.T) { //nolint:thelper // g.Run replaces t.Run
		cmd := exec.Command(cliPath,
			clustersEntity,
			"delete",
			issClusterName,
			"--force",
			"--watch")
		cmd.Env = os.Environ()
		resp, err := internal.RunAndGetStdOut(cmd)
		req.NoError(err, string(resp))

		expected := fmt.Sprintf("Deleting cluster '%s'Cluster deleted\n", issClusterName)
		assert.Equal(t, expected, string(resp))
	})
}
