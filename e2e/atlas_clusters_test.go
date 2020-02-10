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
// +build e2e

package e2e

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"testing"

	"github.com/mongodb/go-client-mongodb-atlas/mongodbatlas"
)

func TestAtlasClusters(t *testing.T) {
	cliPath, err := filepath.Abs("../bin/mcli")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	_, err = os.Stat(cliPath)

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	atlasEntity := "atlas"
	clustersEntity := "clusters"
	clusterName := "e2e-cluster"

	t.Run("Create", func(t *testing.T) {
		cmd := exec.Command(cliPath, atlasEntity, clustersEntity, "create", clusterName, "--region=US_EAST_1", "--members=3", "--instanceSize=M5", "--provider=AWS", "--mdbVersion=4.2")
		cmd.Env = os.Environ()
		resp, err := cmd.CombinedOutput()

		if err != nil {
			t.Fatalf("unexpected error: %v, resp: %v", err, string(resp))
		}

		cluster := mongodbatlas.Cluster{}
		err = json.Unmarshal(resp, &cluster)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		if cluster.Name != clusterName {
			t.Errorf("got=%#v\nwant=%#v\n", cluster.Name, clusterName)
		}
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
	})

	t.Run("Delete", func(t *testing.T) {
		cmd := exec.Command(cliPath, atlasEntity, clustersEntity, "delete", clusterName, "--force")
		cmd.Env = os.Environ()
		resp, err := cmd.CombinedOutput()

		if err != nil {
			t.Fatalf("unexpected error: %v, resp: %v", err, string(resp))
		}

		if string(resp) != fmt.Sprintf("Cluster '%s' deleted\n", clusterName) {
			t.Errorf("got=%#v\nwant=%#v\n", string(resp), fmt.Sprintf("Cluster '%s' deleted\n", clusterName))
		}
	})
}
