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
// +build e2e cloudmanager,clusters

package cloud_manager_test

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"testing"
	"time"

	"github.com/mongodb/mongocli/e2e"
	"github.com/mongodb/mongocli/internal/convert"
)

func TestDeployReplicaSet(t *testing.T) {
	cliPath, err := e2e.Bin()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	const clustersEntity = "clusters"
	const testFile = "om-new-cluster.json"

	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	clusterName := fmt.Sprintf("e2e-cluster-%v", r.Uint32())

	hostname, err := automatedServer(cliPath)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if err := generateConfig(testFile, hostname, clusterName, "4.2.0", "4.2"); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	t.Run("Apply", func(t *testing.T) {
		cmd := exec.Command(cliPath,
			entity,
			clustersEntity,
			"apply",
			"-f="+testFile,
		)

		cmd.Env = os.Environ()
		resp, err := cmd.CombinedOutput()

		if err != nil {
			t.Fatalf("unexpected error: %v, resp: %v\n", err, string(resp))
		}
	})

	t.Run("Watch", func(t *testing.T) {
		cmd := exec.Command(cliPath,
			entity,
			"automation",
			"watch",
		)

		cmd.Env = os.Environ()
		resp, err := cmd.CombinedOutput()

		if err != nil {
			t.Fatalf("unexpected error: %v, resp: %v\n", err, string(resp))
		}
	})

	t.Run("List", func(t *testing.T) {
		cmd := exec.Command(cliPath,
			entity,
			clustersEntity,
			"ls",
		)

		cmd.Env = os.Environ()
		resp, err := cmd.CombinedOutput()

		if err != nil {
			t.Fatalf("unexpected error: %v, resp: %v\n", err, string(resp))
		}
		var clusters []*convert.ClusterConfig
		if err := json.Unmarshal(resp, &clusters); err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		if len(clusters) == 0 {
			t.Errorf("expected len(clusters) > 0, got 0\n")
		}
	})

	t.Run("Describe", func(t *testing.T) {
		cmd := exec.Command(cliPath,
			entity,
			clustersEntity,
			"describe",
			clusterName,
		)

		cmd.Env = os.Environ()
		resp, err := cmd.CombinedOutput()

		if err != nil {
			t.Fatalf("unexpected error: %v, resp: %v\n", err, string(resp))
		}
		var cluster convert.ClusterConfig
		if err := json.Unmarshal(resp, &cluster); err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		if cluster.Name != clusterName {
			t.Errorf("expected %s, got %s\n", clusterName, cluster.Name)
		}
	})
}
