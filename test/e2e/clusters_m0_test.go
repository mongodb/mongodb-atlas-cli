// Copyright 2021 MongoDB Inc
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
//go:build e2e || (atlas && clusters && m0)

package e2e_test

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	atlasClustersPinned "go.mongodb.org/atlas-sdk/v20240530005/admin"
)

func TestClustersM0Flags(t *testing.T) {
	g := newAtlasE2ETestGenerator(t)
	g.generateProject("clustersM0")

	cliPath, err := AtlasCLIBin()
	req := require.New(t)
	req.NoError(err)

	clusterName, err := RandClusterName()
	req.NoError(err)

	t.Run("Create", func(t *testing.T) {
		cmd := exec.Command(cliPath,
			clustersEntity,
			"create",
			clusterName,
			"--region", "US_EAST_1",
			"--members=3",
			"--tier", tierM0,
			"--provider", e2eClusterProvider,
			"--projectId", g.projectID,
			"-o=json")
		cmd.Env = os.Environ()
		resp, err := RunAndGetStdOut(cmd)
		req.NoError(err, string(resp))

		var cluster *atlasClustersPinned.AdvancedClusterDescription
		err = json.Unmarshal(resp, &cluster)
		req.NoError(err)

		ensureCluster(t, cluster, clusterName, "8.0", 0.5, false)
	})

	t.Run("Watch", func(t *testing.T) {
		cmd := exec.Command(cliPath,
			clustersEntity,
			"watch",
			clusterName,
			"--projectId", g.projectID,
		)
		cmd.Env = os.Environ()
		resp, err := RunAndGetStdOut(cmd)
		req.NoError(err, string(resp))

		a := assert.New(t)
		a.Contains(string(resp), "Cluster available")
	})

	t.Run("Describe", func(t *testing.T) {
		cmd := exec.Command(cliPath,
			clustersEntity,
			"describe",
			clusterName,
			"--projectId", g.projectID,
			"-o=json",
		)
		cmd.Env = os.Environ()
		resp, err := RunAndGetStdOut(cmd)
		req.NoError(err, string(resp))

		var cluster atlasClustersPinned.AdvancedClusterDescription
		err = json.Unmarshal(resp, &cluster)
		req.NoError(err)

		a := assert.New(t)
		a.Equal(clusterName, cluster.GetName())
	})

	t.Run("Delete", func(t *testing.T) {
		cmd := exec.Command(cliPath,
			clustersEntity,
			"delete",
			clusterName,
			"--force",
			"--projectId", g.projectID,
		)
		cmd.Env = os.Environ()
		resp, err := RunAndGetStdOut(cmd)
		req.NoError(err, string(resp))

		expected := fmt.Sprintf("Deleting cluster '%s'", clusterName)
		a := assert.New(t)
		a.Equal(expected, string(resp))
	})

	t.Run("Watch", func(t *testing.T) {
		cmd := exec.Command(cliPath,
			clustersEntity,
			"watch",
			clusterName,
			"--projectId", g.projectID,
		)
		cmd.Env = os.Environ()
		// this command will fail with 404 once the cluster is deleted
		// we just need to wait for this to close the project
		resp, _ := RunAndGetStdOut(cmd)
		t.Log(string(resp))
	})
}
