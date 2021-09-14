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

//go:build e2e || (atlas && clusters && !m0)
// +build e2e atlas,clusters,!m0

package atlas_test

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"testing"

	"github.com/mongodb/mongocli/e2e"
	"github.com/mongodb/mongocli/internal/config"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.mongodb.org/atlas/mongodbatlas"
)

const (
	tierM30      = "M30"
	diskSizeGB40 = "40"
	diskSizeGB30 = "30"
)

func TestClustersFlags(t *testing.T) {
	cliPath, err := e2e.Bin()
	req := require.New(t)
	req.NoError(err)

	clusterName, err := RandClusterName()
	req.NoError(err)

	region, err := newAvailableRegion(tierM30, e2eClusterProvider)
	req.NoError(err)

	t.Run("Create", func(t *testing.T) {
		cmd := exec.Command(cliPath,
			atlasEntity,
			clustersEntity,
			"create",
			clusterName,
			"--region", region,
			"--members=3",
			"--tier", tierM30,
			"--provider", e2eClusterProvider,
			"--mdbVersion", e2eMDBVer,
			"--diskSizeGB", diskSizeGB30,
			"-o=json")
		cmd.Env = os.Environ()
		resp, err := cmd.CombinedOutput()
		req.NoError(err, string(resp))

		var cluster *mongodbatlas.AdvancedCluster
		err = json.Unmarshal(resp, &cluster)
		req.NoError(err)

		ensureCluster(t, cluster, clusterName, e2eMDBVer, 30)
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
		req.NoError(err, string(resp))

		a := assert.New(t)
		a.Contains(string(resp), "Cluster available")
	})

	t.Run("Load Sample Data", func(t *testing.T) {
		cmd := exec.Command(cliPath,
			atlasEntity,
			clustersEntity,
			"loadSampleData",
			clusterName,
			"-o=json")
		cmd.Env = os.Environ()
		resp, err := cmd.CombinedOutput()
		req.NoError(err, string(resp))

		var job *mongodbatlas.SampleDatasetJob
		err = json.Unmarshal(resp, &job)
		req.NoError(err)

		a := assert.New(t)
		a.Equal(clusterName, job.ClusterName)
	})

	t.Run("List", func(t *testing.T) {
		cmd := exec.Command(cliPath,
			atlasEntity,
			clustersEntity,
			"ls",
			"-o=json")
		cmd.Env = os.Environ()
		resp, err := cmd.CombinedOutput()
		req.NoError(err, string(resp))

		var clusters mongodbatlas.AdvancedClustersResponse
		err = json.Unmarshal(resp, &clusters)
		req.NoError(err)

		a := assert.New(t)
		a.NotEmpty(clusters.Results)
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
		req.NoError(err, string(resp))

		var cluster mongodbatlas.AdvancedCluster
		err = json.Unmarshal(resp, &cluster)
		req.NoError(err)

		a := assert.New(t)
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
		req.NoError(err, string(resp))

		var connectionString mongodbatlas.ConnectionStrings
		err = json.Unmarshal(resp, &connectionString)
		req.NoError(err)

		a := assert.New(t)
		a.NotEmpty(connectionString.Standard)
		a.NotEmpty(connectionString.StandardSrv)
	})

	t.Run("Create Rolling Index", func(t *testing.T) {
		cmd := exec.Command(cliPath,
			atlasEntity,
			clustersEntity,
			"indexes",
			"create",
			"--clusterName", clusterName,
			"--db=tes",
			"--collection=tes",
			"--key=name:1")
		cmd.Env = os.Environ()
		resp, err := cmd.CombinedOutput()
		req.NoError(err, string(resp))
	})

	t.Run("Update", func(t *testing.T) {
		cmd := exec.Command(cliPath,
			atlasEntity,
			clustersEntity,
			"update",
			clusterName,
			"--diskSizeGB", diskSizeGB40,
			"--mdbVersion=5.0",
			"-o=json")
		cmd.Env = os.Environ()
		resp, err := cmd.CombinedOutput()
		req.NoError(err, string(resp))

		var cluster mongodbatlas.AdvancedCluster
		err = json.Unmarshal(resp, &cluster)
		req.NoError(err)

		ensureCluster(t, &cluster, clusterName, "5.0", 40)
	})

	t.Run("Delete", func(t *testing.T) {
		cmd := exec.Command(cliPath, atlasEntity, clustersEntity, "delete", clusterName, "--force")
		cmd.Env = os.Environ()
		resp, err := cmd.CombinedOutput()
		req.NoError(err)

		expected := fmt.Sprintf("Cluster '%s' deleted\n", clusterName)
		a := assert.New(t)
		a.Equal(expected, string(resp))
	})
}

func TestClustersFile(t *testing.T) {
	cliPath, err := e2e.Bin()
	req := require.New(t)
	req.NoError(err)

	clusterFileName, err := RandClusterName()
	req.NoError(err)

	clusterFile := "create_cluster_test.json"
	if service := os.Getenv("MCLI_SERVICE"); service == config.CloudGovService {
		clusterFile = "create_cluster_gov_test.json"
	}

	t.Run("Create via file", func(t *testing.T) {
		cmd := exec.Command(cliPath,
			atlasEntity,
			clustersEntity,
			"create",
			clusterFileName,
			"--file", clusterFile,
			"-o=json")
		cmd.Env = os.Environ()
		resp, err := cmd.CombinedOutput()
		req.NoError(err, string(resp))

		var cluster mongodbatlas.AdvancedCluster
		err = json.Unmarshal(resp, &cluster)
		req.NoError(err)

		ensureCluster(t, &cluster, clusterFileName, e2eMDBVer, 30)
	})

	t.Run("Watch", func(t *testing.T) {
		cmd := exec.Command(cliPath,
			atlasEntity,
			clustersEntity,
			"watch",
			clusterFileName,
		)
		cmd.Env = os.Environ()
		resp, err := cmd.CombinedOutput()
		req.NoError(err, string(resp))

		a := assert.New(t)
		a.Contains(string(resp), "Cluster available")
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
		req.NoError(err, string(resp))

		var cluster mongodbatlas.AdvancedCluster
		err = json.Unmarshal(resp, &cluster)
		req.NoError(err)

		ensureCluster(t, &cluster, clusterFileName, e2eMDBVer, 40)
	})

	t.Run("Delete file creation", func(t *testing.T) {
		cmd := exec.Command(cliPath, atlasEntity, clustersEntity, "delete", clusterFileName, "--force")
		cmd.Env = os.Environ()
		resp, err := cmd.CombinedOutput()
		req.NoError(err, string(resp))

		expected := fmt.Sprintf("Cluster '%s' deleted\n", clusterFileName)
		a := assert.New(t)
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

	region, err := newAvailableRegion(tierM30, e2eClusterProvider)
	req.NoError(err)

	t.Run("Create sharded cluster", func(t *testing.T) {
		cmd := exec.Command(cliPath,
			atlasEntity,
			clustersEntity,
			"create",
			shardedClusterName,
			"--region", region,
			"--type=SHARDED",
			"--shards=2",
			"--members=3",
			"--tier", tierM30,
			"--provider", e2eClusterProvider,
			"--mdbVersion=4.2",
			"--diskSizeGB", diskSizeGB30,
			"-o=json")

		cmd.Env = os.Environ()
		resp, err := cmd.CombinedOutput()
		req.NoError(err, string(resp))

		var cluster mongodbatlas.AdvancedCluster
		err = json.Unmarshal(resp, &cluster)
		req.NoError(err)

		ensureCluster(t, &cluster, shardedClusterName, "4.2", 30)
	})

	t.Run("Delete sharded cluster", func(t *testing.T) {
		cmd := exec.Command(cliPath, atlasEntity, clustersEntity, "delete", shardedClusterName, "--force")
		cmd.Env = os.Environ()
		resp, err := cmd.CombinedOutput()
		req.NoError(err, string(resp))

		expected := fmt.Sprintf("Cluster '%s' deleted\n", shardedClusterName)
		a.Equal(expected, string(resp))
	})
}
