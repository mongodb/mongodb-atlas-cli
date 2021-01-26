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
// +build e2e atlas,clusters

package atlas_test

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"testing"

	"github.com/mongodb/mongocli/e2e"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.mongodb.org/atlas/mongodbatlas"
)

func TestClustersFlags(t *testing.T) {
	cliPath, err := e2e.Bin()
	a := assert.New(t)
	req := require.New(t)
	req.NoError(err)
	clusterName, err := RandClusterName()
	req.NoError(err)

	t.Run("Create", func(t *testing.T) {
		cmd := exec.Command(cliPath,
			atlasEntity,
			clustersEntity,
			"create",
			clusterName,
			"--region=US_EAST_1",
			"--members=3",
			"--tier=M10",
			"--provider=AWS",
			"--mdbVersion=4.0",
			"--diskSizeGB=10",
			"-o=json")
		cmd.Env = os.Environ()
		resp, err := cmd.CombinedOutput()

		req.NoError(err)

		var cluster *mongodbatlas.Cluster
		err = json.Unmarshal(resp, &cluster)
		a.NoError(err)

		ensureCluster(t, cluster, clusterName, "4.0", 10)
	})

	t.Run("Watch", func(t *testing.T) {
		cmd := exec.Command(cliPath,
			atlasEntity,
			clustersEntity,
			"watch",
			clusterName,
		)
		cmd.Env = os.Environ()
		resp, err := cmd.CombinedOutput()
		req.NoError(err)
		a.Contains(string(resp), "Cluster available")
	})

	t.Run("List", func(t *testing.T) {
		cmd := exec.Command(cliPath,
			atlasEntity,
			clustersEntity,
			"ls",
			"-o=json")
		cmd.Env = os.Environ()
		resp, err := cmd.CombinedOutput()
		req.NoError(err)

		var clusters []mongodbatlas.Cluster
		err = json.Unmarshal(resp, &clusters)
		req.NoError(err)
		a.NotEmpty(clusters)
	})

	t.Run("Describe", func(t *testing.T) {
		cmd := exec.Command(cliPath,
			atlasEntity,
			clustersEntity,
			"describe",
			clusterName,
			"-o=json")
		cmd.Env = os.Environ()
		resp, err := cmd.CombinedOutput()
		req.NoError(err)

		var cluster mongodbatlas.Cluster
		err = json.Unmarshal(resp, &cluster)
		req.NoError(err)
		a.Equal(clusterName, cluster.Name)
	})

	t.Run("Describe Connection String", func(t *testing.T) {
		cmd := exec.Command(cliPath,
			atlasEntity,
			clustersEntity,
			"cs",
			"describe",
			clusterName,
			"-o=json")
		cmd.Env = os.Environ()
		resp, err := cmd.CombinedOutput()
		req.NoError(err)

		var connectionString mongodbatlas.ConnectionStrings
		err = json.Unmarshal(resp, &connectionString)
		req.NoError(err)

		a.NotEmpty(connectionString.Standard)
		a.NotEmpty(connectionString.StandardSrv)
	})

	t.Run("Create Rolling Index", func(t *testing.T) {
		cmd := exec.Command(cliPath,
			atlasEntity,
			clustersEntity,
			"indexes",
			"create",
			"--clusterName="+clusterName,
			"--db=tes",
			"--collection=tes",
			"--key=name:1")
		cmd.Env = os.Environ()
		_, err := cmd.CombinedOutput()
		req.NoError(err)
	})

	t.Run("Update", func(t *testing.T) {
		cmd := exec.Command(cliPath,
			atlasEntity,
			clustersEntity,
			"update",
			clusterName,
			"--diskSizeGB=20",
			"--mdbVersion=4.2",
			"-o=json")
		cmd.Env = os.Environ()
		resp, err := cmd.CombinedOutput()
		req.NoError(err)

		var cluster mongodbatlas.Cluster
		err = json.Unmarshal(resp, &cluster)
		req.NoError(err)

		ensureCluster(t, &cluster, clusterName, "4.2", 20)
	})

	t.Run("Delete", func(t *testing.T) {
		cmd := exec.Command(cliPath, atlasEntity, clustersEntity, "delete", clusterName, "--force")
		cmd.Env = os.Environ()
		resp, err := cmd.CombinedOutput()
		req.NoError(err)

		expected := fmt.Sprintf("Cluster '%s' deleted\n", clusterName)
		a.Equal(expected, string(resp))
	})
}

func TestClustersFile(t *testing.T) {
	cliPath, err := e2e.Bin()
	a := assert.New(t)
	a.NoError(err)

	clusterFileName, err := RandClusterName()
	a.NoError(err)

	t.Run("Create via file", func(t *testing.T) {
		cmd := exec.Command(cliPath,
			atlasEntity,
			clustersEntity,
			"create",
			clusterFileName,
			"--file=create_cluster_test.json",
			"-o=json")
		cmd.Env = os.Environ()
		resp, err := cmd.CombinedOutput()
		a.NoError(err)

		var cluster mongodbatlas.Cluster
		err = json.Unmarshal(resp, &cluster)
		a.NoError(err)

		ensureCluster(t, &cluster, clusterFileName, "4.2", 10)
	})

	t.Run("Update via file", func(t *testing.T) {
		cmd := exec.Command(cliPath,
			atlasEntity,
			clustersEntity,
			"update",
			clusterFileName,
			"--file=update_cluster_test.json",
			"-o=json")
		cmd.Env = os.Environ()
		resp, err := cmd.CombinedOutput()
		a.NoError(err)

		var cluster mongodbatlas.Cluster
		err = json.Unmarshal(resp, &cluster)
		a.NoError(err)

		ensureCluster(t, &cluster, clusterFileName, "4.2", 25)
	})

	t.Run("Delete file creation", func(t *testing.T) {
		cmd := exec.Command(cliPath, atlasEntity, clustersEntity, "delete", clusterFileName, "--force")
		cmd.Env = os.Environ()
		resp, err := cmd.CombinedOutput()
		a.NoError(err)

		expected := fmt.Sprintf("Cluster '%s' deleted\n", clusterFileName)
		a.Equal(expected, string(resp))
	})
}

func TestShardedCluster(t *testing.T) {
	cliPath, err := e2e.Bin()
	a := assert.New(t)
	req := require.New(t)
	req.NoError(err)

	shardedClusterName, err := RandClusterName()
	req.NoError(err)

	t.Run("Create sharded cluster", func(t *testing.T) {
		cmd := exec.Command(cliPath,
			atlasEntity,
			clustersEntity,
			"create",
			shardedClusterName,
			"--region=US_EAST_1",
			"--type=SHARDED",
			"--shards=2",
			"--members=3",
			"--tier=M10",
			"--provider=AWS",
			"--mdbVersion=4.2",
			"--diskSizeGB=10",
			"-o=json")

		cmd.Env = os.Environ()
		resp, err := cmd.CombinedOutput()
		req.NoError(err)

		var cluster mongodbatlas.Cluster
		err = json.Unmarshal(resp, &cluster)
		req.NoError(err)

		ensureCluster(t, &cluster, shardedClusterName, "4.2", 10)
	})

	t.Run("Delete sharded cluster", func(t *testing.T) {
		cmd := exec.Command(cliPath, atlasEntity, clustersEntity, "delete", shardedClusterName, "--force")
		cmd.Env = os.Environ()
		resp, err := cmd.CombinedOutput()
		req.NoError(err)

		expected := fmt.Sprintf("Cluster '%s' deleted\n", shardedClusterName)
		a.Equal(expected, string(resp))
	})
}

func ensureCluster(t *testing.T, cluster *mongodbatlas.Cluster, clusterName, version string, diskSizeGB float64) {
	t.Helper()
	a := assert.New(t)
	a.Equal(clusterName, cluster.Name)
	a.Equal(version, cluster.MongoDBMajorVersion)
	a.Equal(diskSizeGB, *cluster.DiskSizeGB)
}
