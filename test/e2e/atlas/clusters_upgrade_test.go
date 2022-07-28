// Copyright 2022 MongoDB Inc
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
//go:build e2e || (atlas && clusters && upgrade)

package atlas_test

import (
	"encoding/json"
	"fmt"
	"os"
	"testing"

	"github.com/mongodb/mongodb-atlas-cli/test/e2e"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.mongodb.org/atlas/mongodbatlas"
	exec "golang.org/x/sys/execabs"
)

func TestSharedClusterUpgrade(t *testing.T) {
	g := newAtlasE2ETestGenerator(t)
	g.generateProject("clustersUpgrade")

	cliPath, err := e2e.AtlasCLIBin()
	req := require.New(t)
	req.NoError(err)

	clusterName, err := RandClusterName()
	req.NoError(err)

	region, err := g.newAvailableRegion(tierM2, e2eClusterProvider)
	req.NoError(err)

	t.Run("Create", func(t *testing.T) {
		cmd := exec.Command(cliPath,
			clustersEntity,
			"create",
			clusterName,
			"--region", region,
			"--tier", tierM2,
			"--provider", e2eClusterProvider,
			"--mdbVersion", e2eSharedMDBVer,
			"--projectId", g.projectID,
			"-o=json")
		cmd.Env = os.Environ()
		resp, err := cmd.CombinedOutput()
		req.NoError(err, string(resp))

		var cluster *mongodbatlas.Cluster
		err = json.Unmarshal(resp, &cluster)
		req.NoError(err)

		ensureSharedCluster(t, cluster, clusterName, e2eSharedMDBVer, tierM2, 2)
	})

	t.Run("Watch create", func(t *testing.T) {
		cmd := exec.Command(cliPath,
			clustersEntity,
			"watch",
			clusterName,
			"--projectId", g.projectID,
		)
		cmd.Env = os.Environ()
		resp, _ := cmd.CombinedOutput()
		t.Log(string(resp))
	})

	t.Run("Upgrade", func(t *testing.T) {
		cmd := exec.Command(cliPath,
			clustersEntity,
			"upgrade",
			clusterName,
			"--tier", tierM10,
			"--diskSizeGB", diskSizeGB40,
			"--mdbVersion=6.0",
			"--projectId", g.projectID,
			"-o=json")
		cmd.Env = os.Environ()
		resp, err := cmd.CombinedOutput()
		req.NoError(err, string(resp))
	})

	t.Run("Watch upgrade", func(t *testing.T) {
		cmd := exec.Command(cliPath,
			clustersEntity,
			"watch",
			clusterName,
			"--projectId", g.projectID,
		)
		cmd.Env = os.Environ()
		resp, _ := cmd.CombinedOutput()
		t.Log(string(resp))
	})

	t.Run("Ensure upgrade", func(t *testing.T) {
		cmd := exec.Command(cliPath,
			clustersEntity,
			"ls",
			"--projectId", g.projectID,
			"-o=json")
		cmd.Env = os.Environ()
		resp, err := cmd.CombinedOutput()
		req.NoError(err, string(resp))

		var clusterResponse mongodbatlas.ClustersResponse
		err = json.Unmarshal(resp, &clusterResponse)
		req.NoError(err)

		ensureSharedCluster(t, clusterResponse.Results[0], clusterName, "6.0", tierM10, 40)
	})

	t.Run("Delete", func(t *testing.T) {
		cmd := exec.Command(cliPath,
			clustersEntity,
			"delete",
			clusterName,
			"--projectId", g.projectID,
			"--force",
		)
		cmd.Env = os.Environ()
		resp, err := cmd.CombinedOutput()
		req.NoError(err, string(resp))

		expected := fmt.Sprintf("Cluster '%s' deleted\n", clusterName)
		a := assert.New(t)
		a.Equal(expected, string(resp))
	})

	t.Run("Watch deletion", func(t *testing.T) {
		cmd := exec.Command(cliPath,
			clustersEntity,
			"watch",
			clusterName,
			"--projectId", g.projectID,
		)
		cmd.Env = os.Environ()
		resp, _ := cmd.CombinedOutput()
		t.Log(string(resp))
	})
}
