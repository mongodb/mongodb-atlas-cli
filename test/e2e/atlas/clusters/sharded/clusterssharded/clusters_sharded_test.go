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

package clusterssharded

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
	atlasClustersPinned "go.mongodb.org/atlas-sdk/v20240530005/admin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	clustersEntity = "clusters"
	dbusersEntity  = "dbusers"
	diskSizeGB30   = "30"

	// Cluster settings.
	e2eClusterProvider = "AWS"
)

func TestShardedCluster(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping test in short mode")
	}

	g := internal.NewAtlasE2ETestGenerator(t, internal.WithSnapshot())
	g.GenerateProject("shardedClusters")

	cliPath, err := internal.AtlasCLIBin()
	req := require.New(t)
	req.NoError(err)

	shardedClusterName := g.Memory("shardedClusterName", internal.Must(internal.RandClusterName())).(string)
	dbUserUsername := g.Memory("dbUserUsername", internal.Must(internal.RandUsername())).(string)
	dbUserPassword := dbUserUsername + "~PassWord"

	var client *mongo.Client
	ctx := t.Context()

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
			"-o=json",
			"-P",
			internal.ProfileName(),
		)

		cmd.Env = os.Environ()
		resp, err := internal.RunAndGetStdOut(cmd)
		req.NoError(err, string(resp))

		var cluster atlasClustersPinned.AdvancedClusterDescription
		req.NoError(json.Unmarshal(resp, &cluster))

		internal.EnsureCluster(t, &cluster, shardedClusterName, mdbVersion, 30, false)

		cmd = exec.Command(cliPath,
			dbusersEntity,
			"create",
			"atlasAdmin",
			"--username", dbUserUsername,
			"--password", dbUserPassword,
			"--projectId", g.ProjectID,
			"-P",
			internal.ProfileName(),
		)
		cmd.Env = os.Environ()
		resp, err = internal.RunAndGetStdOut(cmd)
		req.NoError(err, string(resp))
	})
	require.NoError(t, internal.WatchCluster(g.ProjectID, shardedClusterName))

	g.Run("Connect to cluster", func(t *testing.T) { //nolint:thelper // g.Run replaces t.Run
		cmd := exec.Command(cliPath,
			clustersEntity,
			"connect",
			shardedClusterName,
			"--connectWith", "connectionString",
			"--projectId", g.ProjectID,
			"-P",
			internal.ProfileName(),
		)
		cmd.Env = os.Environ()
		r, err := internal.RunAndGetStdOut(cmd)
		req.NoError(err, string(r))

		connectionString := strings.TrimSpace(string(r))
		req.NotEmpty(connectionString, "connection string should not be empty")
		assert.Contains(t, connectionString, "mongodb", "connection string should contain mongodb URI")

		mode, err := internal.TestRunMode()
		req.NoError(err)

		if mode != internal.TestModeLive {
			t.Skip("skipping actual MongoDB connection in snapshot mode")
		}

		client, err = mongo.Connect(
			ctx,
			options.Client().
				ApplyURI(connectionString).
				SetAuth(options.Credential{
					AuthMechanism: "PLAIN",
					Username:      dbUserUsername,
					Password:      dbUserPassword,
				}),
		)
		require.NoError(t, err)
	})

	t.Cleanup(func() {
		if client != nil {
			_ = client.Disconnect(ctx)
		}
	})

	g.Run("Delete sharded cluster", func(t *testing.T) { //nolint:thelper // g.Run replaces t.Run
		cmd := exec.Command(cliPath,
			clustersEntity,
			"delete",
			shardedClusterName,
			"--projectId", g.ProjectID,
			"--force",
			"-P",
			internal.ProfileName(),
		)
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
			"-P",
			internal.ProfileName(),
		)
		cmd.Env = os.Environ()
		// this command will fail with 404 once the cluster is deleted
		// we just need to wait for this to close the project
		resp, _ := internal.RunAndGetStdOut(cmd)
		t.Log(string(resp))
	})
}
