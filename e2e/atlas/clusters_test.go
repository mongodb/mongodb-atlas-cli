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
// +build e2e,atlas

package atlas_test

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"path/filepath"
	"testing"
	"time"

	"github.com/mongodb/go-client-mongodb-atlas/mongodbatlas"
)

func TestAtlasClusters(t *testing.T) {
	cliPath, err := filepath.Abs("../../bin/mongocli")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	_, err = os.Stat(cliPath)

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	atlasEntity := "atlas"
	clustersEntity := "clusters"
	clusterName := fmt.Sprintf("e2e-cluster-%v", r.Uint32())

	t.Run("Create via params", func(t *testing.T) {
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
			"--diskSizeGB=10")
		cmd.Env = os.Environ()
		resp, err := cmd.CombinedOutput()

		if err != nil {
			t.Fatalf("unexpected error: %v, resp: %v", err, string(resp))
		}

		cluster := new(mongodbatlas.Cluster)
		err = json.Unmarshal(resp, cluster)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		ensureCluster(t, cluster, clusterName, "4.0", 10)
	})

	t.Run("List", func(t *testing.T) {
		cmd := exec.Command(cliPath, atlasEntity, clustersEntity, "ls")
		cmd.Env = os.Environ()
		resp, err := cmd.CombinedOutput()

		if err != nil {
			t.Fatalf("unexpected error: %v, resp: %v", err, string(resp))
		}
	})

	t.Run("Describe", func(t *testing.T) {
		cmd := exec.Command(cliPath, atlasEntity, clustersEntity, "describe", clusterName)
		cmd.Env = os.Environ()
		resp, err := cmd.CombinedOutput()

		if err != nil {
			t.Fatalf("unexpected error: %v, resp: %v", err, string(resp))
		}

		cluster := new(mongodbatlas.Cluster)
		err = json.Unmarshal(resp, cluster)
		if err != nil {
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
			"--mdbVersion=4.2")
		cmd.Env = os.Environ()
		resp, err := cmd.CombinedOutput()

		if err != nil {
			t.Fatalf("unexpected error: %v, resp: %v", err, string(resp))
		}

		cluster := new(mongodbatlas.Cluster)
		err = json.Unmarshal(resp, cluster)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		ensureCluster(t, cluster, clusterName, "4.2", 20)
	})

	// TODO: this fails as the cluster is not healthy we may need to re think how we test this
	//t.Run("Create Index", func(t *testing.T) {
	//	cmd := exec.Command(cliPath,
	//		atlasEntity,
	//		clustersEntity,
	//		"indexes",
	//		"create",
	//		"--clusterName="+clusterName,
	//		"--db=tes",
	//		"--collection=tes",
	//		"--key=name:1")
	//	cmd.Env = os.Environ()
	//	resp, err := cmd.CombinedOutput()
	//
	//	if err != nil {
	//		t.Fatalf("unexpected error: %v, resp: %v", err, string(resp))
	//	}
	//})

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

	clusterFileName := fmt.Sprintf("e2e-cluster-%v", r.Uint32())
	t.Run("Create via file", func(t *testing.T) {
		cmd := exec.Command(cliPath,
			atlasEntity,
			clustersEntity,
			"create",
			clusterFileName,
			"--file=create_cluster_test.json")
		cmd.Env = os.Environ()
		resp, err := cmd.CombinedOutput()

		if err != nil {
			t.Fatalf("unexpected error: %v, resp: %v", err, string(resp))
		}

		cluster := new(mongodbatlas.Cluster)
		err = json.Unmarshal(resp, cluster)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		ensureCluster(t, cluster, clusterFileName, "4.2", 10)
	})

	t.Run("Update via file", func(t *testing.T) {
		cmd := exec.Command(cliPath,
			atlasEntity,
			clustersEntity,
			"update",
			clusterFileName,
			"--file=update_cluster_test.json")
		cmd.Env = os.Environ()
		resp, err := cmd.CombinedOutput()

		if err != nil {
			t.Fatalf("unexpected error: %v, resp: %v", err, string(resp))
		}

		cluster := new(mongodbatlas.Cluster)
		err = json.Unmarshal(resp, cluster)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		ensureCluster(t, cluster, clusterFileName, "4.2", 25)
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

func ensureCluster(t *testing.T, cluster *mongodbatlas.Cluster, clusterName string, version string, diskSizeGB float64) {
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
