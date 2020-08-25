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
	"strings"
	"testing"

	"github.com/mongodb/mongocli/e2e"
	"go.mongodb.org/atlas/mongodbatlas"
)

func TestClustersFlags(t *testing.T) {
	cliPath, e := e2e.Bin()
	if e != nil {
		t.Fatalf("unexpected error: %v", e)
	}
	clusterName, e := RandClusterName()
	if e != nil {
		t.Fatalf("unexpected error: %v", e)
	}
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

		if err != nil {
			t.Fatalf("unexpected error: %v, resp: %v", err, string(resp))
		}

		var cluster *mongodbatlas.Cluster
		if err := json.Unmarshal(resp, &cluster); err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

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

		if err != nil {
			t.Fatalf("unexpected error: %v, resp: %v", err, string(resp))
		}

		if !strings.Contains(string(resp), "Cluster available") {
			t.Errorf("got=%#v\nwant=%#v\n", string(resp), "Cluster available at:")
		}
	})

	t.Run("List", func(t *testing.T) {
		cmd := exec.Command(cliPath,
			atlasEntity,
			clustersEntity,
			"ls",
			"-o=json")
		cmd.Env = os.Environ()
		resp, err := cmd.CombinedOutput()

		if err != nil {
			t.Fatalf("unexpected error: %v, resp: %v", err, string(resp))
		}
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

		if err != nil {
			t.Fatalf("unexpected error: %v, resp: %v", err, string(resp))
		}

		var cluster mongodbatlas.Cluster
		if err := json.Unmarshal(resp, &cluster); err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		if cluster.Name != clusterName {
			t.Errorf("got=%#v\nwant=%#v\n", cluster.Name, clusterName)
		}
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

		if err != nil {
			t.Fatalf("unexpected error: %v, resp: %v", err, string(resp))
		}

		var cluster mongodbatlas.Cluster
		if err := json.Unmarshal(resp, &cluster); err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		ensureCluster(t, &cluster, clusterName, "4.2", 20)
	})

	t.Run("Create Index", func(t *testing.T) {
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
		resp, err := cmd.CombinedOutput()

		if err != nil {
			t.Fatalf("unexpected error: %v, resp: %v", err, string(resp))
		}
	})

	t.Run("Delete", func(t *testing.T) {
		cmd := exec.Command(cliPath, atlasEntity, clustersEntity, "delete", clusterName, "--force")
		cmd.Env = os.Environ()
		resp, err := cmd.CombinedOutput()

		if err != nil {
			t.Fatalf("unexpected error: %v, resp: %v", err, string(resp))
		}

		expected := fmt.Sprintf("Cluster '%s' deleted\n", clusterName)
		if string(resp) != expected {
			t.Errorf("got=%#v\nwant=%#v\n", string(resp), expected)
		}
	})
}

func TestClustersFile(t *testing.T) {
	cliPath, e := e2e.Bin()
	if e != nil {
		t.Fatalf("unexpected error: %v", e)
	}
	clusterFileName, e := RandClusterName()
	if e != nil {
		t.Fatalf("unexpected error: %v", e)
	}
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

		if err != nil {
			t.Fatalf("unexpected error: %v, resp: %v", err, string(resp))
		}

		var cluster mongodbatlas.Cluster
		if err := json.Unmarshal(resp, &cluster); err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

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

		if err != nil {
			t.Fatalf("unexpected error: %v, resp: %v", err, string(resp))
		}

		var cluster mongodbatlas.Cluster
		if err := json.Unmarshal(resp, &cluster); err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		ensureCluster(t, &cluster, clusterFileName, "4.2", 25)
	})

	t.Run("Delete file creation", func(t *testing.T) {
		cmd := exec.Command(cliPath, atlasEntity, clustersEntity, "delete", clusterFileName, "--force")
		cmd.Env = os.Environ()
		resp, err := cmd.CombinedOutput()

		if err != nil {
			t.Fatalf("unexpected error: %v, resp: %v", err, string(resp))
		}

		expected := fmt.Sprintf("Cluster '%s' deleted\n", clusterFileName)
		if string(resp) != expected {
			t.Errorf("got=%#v\nwant=%#v\n", string(resp), expected)
		}
	})
}

func TestShardedCluster(t *testing.T) {
	cliPath, e := e2e.Bin()
	if e != nil {
		t.Fatalf("unexpected error: %v", e)
	}
	shardedClusterName, e := RandClusterName()
	if e != nil {
		t.Fatalf("unexpected error: %v", e)
	}
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

		if err != nil {
			t.Fatalf("unexpected error: %v, resp: %v", err, string(resp))
		}

		var cluster mongodbatlas.Cluster
		if err := json.Unmarshal(resp, &cluster); err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		ensureCluster(t, &cluster, shardedClusterName, "4.2", 10)
	})

	t.Run("Delete sharded cluster", func(t *testing.T) {
		cmd := exec.Command(cliPath, atlasEntity, clustersEntity, "delete", shardedClusterName, "--force")
		cmd.Env = os.Environ()
		resp, err := cmd.CombinedOutput()

		if err != nil {
			t.Fatalf("unexpected error: %v, resp: %v", err, string(resp))
		}

		expected := fmt.Sprintf("Cluster '%s' deleted\n", shardedClusterName)
		if string(resp) != expected {
			t.Errorf("got=%#v\nwant=%#v\n", string(resp), expected)
		}
	})
}

func ensureCluster(t *testing.T, cluster *mongodbatlas.Cluster, clusterName, version string, diskSizeGB float64) {
	if cluster.Name != clusterName {
		t.Errorf("Name, got=%s\nwant=%s\n", cluster.Name, clusterName)
	}
	if cluster.MongoDBMajorVersion != version {
		t.Errorf("MongoDBMajorVersion, got=%s\nwant=%s\n", cluster.MongoDBMajorVersion, version)
	}
	if *cluster.DiskSizeGB != diskSizeGB {
		t.Errorf("DiskSizeGB, got=%#v\nwant=%f\n", cluster.DiskSizeGB, diskSizeGB)
	}
}
