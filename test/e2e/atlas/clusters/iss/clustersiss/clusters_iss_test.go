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

package clustersiss

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"strings"
	"testing"

	"github.com/mongodb/mongodb-atlas-cli/atlascli/test/internal"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.mongodb.org/atlas-sdk/v20250312012/admin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	clustersEntity              = "clusters"
	dbusersEntity               = "dbusers"
	diskSizeGB30                = "30"
	independentShardScalingFlag = "independentShardScaling"
	clusterWideScalingFlag      = "clusterWideScaling"

	// Cluster settings.
	e2eClusterProvider = "AWS"
)

func TestIndependendShardScalingCluster(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping test in short mode")
	}

	g := internal.NewAtlasE2ETestGenerator(t, internal.WithSnapshot())
	cliPath, err := internal.AtlasCLIBin()
	req := require.New(t)
	req.NoError(err)

	g.GenerateProject("clustersIss")

	issClusterName := g.Memory("issClusterName", internal.Must(internal.RandClusterName())).(string)

	issDbUserUsername := g.Memory("dbUserUsername", internal.Must(internal.RandUsername())).(string)

	issDbUserPassword := issDbUserUsername + "~PwD"

	var client *mongo.Client
	ctx := t.Context()

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
			"--projectId", g.ProjectID,
			"-o=json",
			"-P",
			internal.ProfileName(),
		)

		cmd.Env = os.Environ()
		resp, err := internal.RunAndGetStdOut(cmd)
		req.NoError(err, string(resp))

		var cluster admin.ClusterDescription20240805
		req.NoError(json.Unmarshal(resp, &cluster))

		internal.EnsureClusterLatest(t, &cluster, issClusterName, mdbVersion, 30, false)
		assert.Len(t, cluster.GetReplicationSpecs(), 2)
		assert.Equal(t, "SHARDED", cluster.GetClusterType())

		cmd = exec.Command(cliPath,
			dbusersEntity,
			"create",
			"atlasAdmin",
			"--username", issDbUserUsername,
			"--password", issDbUserPassword,
			"--projectId", g.ProjectID,
			"-P",
			internal.ProfileName(),
		)
		cmd.Env = os.Environ()
		resp, err = internal.RunAndGetStdOut(cmd)
		req.NoError(err, string(resp))
	})

	g.Run("Get ISS cluster", func(t *testing.T) { //nolint:thelper // g.Run replaces t.Run
		cmd := exec.Command(cliPath,
			clustersEntity,
			"get",
			issClusterName,
			"--autoScalingMode",
			independentShardScalingFlag,
			"--projectId", g.ProjectID,
			"-o=json",
			"-P",
			internal.ProfileName(),
		)

		cmd.Env = os.Environ()
		resp, err := internal.RunAndGetStdOut(cmd)
		req.NoError(err, string(resp))

		var cluster admin.ClusterDescription20240805
		req.NoError(json.Unmarshal(resp, &cluster))
		assert.Equal(t, "SHARDED", cluster.GetClusterType())
		assert.Len(t, cluster.GetReplicationSpecs(), 2)
	})

	g.Run("Pause ISS cluster with the wrong flag", func(_ *testing.T) {
		cmd := exec.Command(cliPath,
			clustersEntity,
			"pause",
			issClusterName,
			"--autoScalingMode",
			"clusterWideScaling",
			"--projectId", g.ProjectID,
			"--output",
			"json",
			"-P",
			internal.ProfileName(),
		)

		cmd.Env = os.Environ()
		resp, err := internal.RunAndGetStdOut(cmd)
		req.NoError(err, string(resp))

		var cluster admin.ClusterDescription20240805
		req.NoError(json.Unmarshal(resp, &cluster))
	})

	g.Run("Check autoScalingMode is independentShardScaling", func(_ *testing.T) {
		cmd := exec.Command(cliPath,
			clustersEntity,
			"autoScalingConfig",
			issClusterName,
			"--projectId", g.ProjectID,
			"-o=json",
			"-P",
			internal.ProfileName(),
		)

		cmd.Env = os.Environ()
		resp, err := internal.RunAndGetStdOut(cmd)
		req.NoError(err, string(resp))

		var config admin.ClusterDescriptionAutoScalingModeConfiguration
		req.NoError(json.Unmarshal(resp, &config))
		assert.Equal(t, "INDEPENDENT_SHARD_SCALING", config.GetAutoScalingMode())
	})

	g.Run("Start ISS cluster", func(_ *testing.T) {
		cmd := exec.Command(cliPath,
			clustersEntity,
			"start",
			issClusterName,
			"--autoScalingMode",
			"independentShardScaling",
			"--projectId", g.ProjectID,
			"--output",
			"json",
			"-P",
			internal.ProfileName(),
		)

		cmd.Env = os.Environ()
		resp, err := internal.RunAndGetStdOut(cmd)
		req.NoError(err, string(resp))

		var cluster admin.ClusterDescription20240805
		req.NoError(json.Unmarshal(resp, &cluster))
	})

	g.Run("Get ISS cluster autoScalingMode", func(t *testing.T) { //nolint:thelper // g.Run replaces t.Run
		cmd := exec.Command(cliPath,
			clustersEntity,
			"autoScalingConfig",
			issClusterName,
			"--projectId", g.ProjectID,
			"-o=json",
			"-P",
			internal.ProfileName(),
		)

		cmd.Env = os.Environ()
		resp, err := internal.RunAndGetStdOut(cmd)
		req.NoError(err, string(resp))

		var config admin.ClusterDescriptionAutoScalingModeConfiguration
		req.NoError(json.Unmarshal(resp, &config))
		assert.Equal(t, "INDEPENDENT_SHARD_SCALING", config.GetAutoScalingMode())
	})

	g.Run("List ISS cluster", func(t *testing.T) { //nolint:thelper // g.Run replaces t.Run
		cmd := exec.Command(cliPath,
			clustersEntity,
			"list",
			"--autoScalingMode", independentShardScalingFlag,
			"--projectId", g.ProjectID,
			"-o=json",
			"-P",
			internal.ProfileName(),
		)
		cmd.Env = os.Environ()
		resp, err := internal.RunAndGetStdOut(cmd)
		req.NoError(err, string(resp))

		var clusters admin.PaginatedClusterDescription20240805
		req.NoError(json.Unmarshal(resp, &clusters))

		assert.Positive(t, clusters.GetTotalCount())
		assert.NotEmpty(t, clusters.Results)
	})

	g.Run("Connect to ISS cluster", func(t *testing.T) { //nolint:thelper // g.Run replaces t.Run
		require.NoError(t, internal.WatchCluster(g.ProjectID, issClusterName), "cluster must be IDLE before connect")

		cmd := exec.Command(cliPath,
			clustersEntity,
			"connect",
			issClusterName,
			"--connectWith", "connectionString",
			"--autoScalingMode", independentShardScalingFlag,
			"--projectId", g.ProjectID,
			"-P",
			internal.ProfileName(),
		)

		cmd.Env = os.Environ()
		resp, err := internal.RunAndGetStdOut(cmd)
		req.NoError(err, string(resp))

		connectionString := strings.TrimSpace(string(resp))
		req.NotEmpty(connectionString, "connection string should not be empty")
		assert.Contains(t, connectionString, "mongodb", "connection string should contain mongodb URI")

		client, err = mongo.Connect(
			ctx,
			options.Client().
				ApplyURI(connectionString).
				SetAuth(options.Credential{
					AuthMechanism: "PLAIN",
					Username:      issDbUserUsername,
					Password:      issDbUserPassword,
				}),
		)
		require.NoError(t, err)
		t.Cleanup(func() {
			_ = client.Disconnect(ctx)
		})
	})

	g.Run("Delete ISS cluster", func(t *testing.T) { //nolint:thelper // g.Run replaces t.Run
		cmd := exec.Command(cliPath,
			clustersEntity,
			"delete",
			issClusterName,
			"--force",
			"--watch",
			"--projectId", g.ProjectID,
			"-P",
			internal.ProfileName(),
		)
		cmd.Env = os.Environ()
		resp, err := internal.RunAndGetStdOut(cmd)
		req.NoError(err, string(resp))

		expected := fmt.Sprintf("Deleting cluster '%s'Cluster deleted\n", issClusterName)
		assert.Equal(t, expected, string(resp))
	})
}
