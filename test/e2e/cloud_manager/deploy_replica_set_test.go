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

//go:build e2e || (remote && replica && (cloudmanager || om60))

package cloud_manager_test

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"testing"

	"github.com/andreaangiolillo/mongocli-test/internal/convert"
	"github.com/andreaangiolillo/mongocli-test/test/e2e"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestDeployReplicaSet(t *testing.T) {
	cliPath, err := e2e.Bin()
	require.NoError(t, err)

	const testFile = "om-new-cluster.json"

	n, err := e2e.RandInt(1000)
	require.NoError(t, err)
	clusterName := fmt.Sprintf("e2e-cluster-%v", n)

	hostname, err := automationServerHostname(cliPath)
	require.NoError(t, err)
	require.NoError(
		t,
		generateRSConfig(testFile, hostname, clusterName, testedMDBVersion, testedMDBFCV),
	)

	t.Run("Apply", func(t *testing.T) {
		cmd := exec.Command(cliPath,
			entity,
			clustersEntity,
			"apply",
			"-f",
			testFile,
		)

		cmd.Env = os.Environ()
		resp, err := cmd.CombinedOutput()
		require.NoError(t, err, string(resp))
	})

	t.Run("Watch", watchAutomation(cliPath))

	t.Run("List", func(t *testing.T) {
		cmd := exec.Command(cliPath,
			entity,
			clustersEntity,
			"ls",
			"-o=json",
		)

		cmd.Env = os.Environ()
		resp, err := cmd.CombinedOutput()
		require.NoError(t, err, string(resp))
		var clusters []*convert.ClusterConfig
		require.NoError(t, json.Unmarshal(resp, &clusters), string(resp))
		assert.NotEmpty(t, clusters)
	})

	t.Run("List No Output", func(t *testing.T) {
		cmd := exec.Command(cliPath,
			entity,
			clustersEntity,
			"ls",
			"-o=",
			"-o=json",
		)

		cmd.Env = os.Environ()
		resp, err := cmd.CombinedOutput()
		require.NoError(t, err, string(resp))
	})

	t.Run("Describe", func(t *testing.T) {
		cmd := exec.Command(cliPath,
			entity,
			clustersEntity,
			"describe",
			clusterName,
			"-o=json",
		)

		cmd.Env = os.Environ()
		resp, err := cmd.CombinedOutput()
		require.NoError(t, err, string(resp))

		var cluster convert.ClusterConfig
		require.NoError(t, json.Unmarshal(resp, &cluster), string(resp))
		assert.Equal(t, cluster.Name, clusterName)
	})

	t.Run("Update", func(t *testing.T) {
		require.NoError(t, generateRSConfigUpdate(testFile))
		cmd := exec.Command(cliPath,
			entity,
			clustersEntity,
			"update",
			clusterName,
			"-f",
			testFile,
		)

		cmd.Env = os.Environ()
		resp, err := cmd.CombinedOutput()
		require.NoError(t, err, string(resp))
	})

	t.Run("Watch", watchAutomation(cliPath))

	t.Run("Reclaim free space", func(t *testing.T) {
		cmd := exec.Command(cliPath,
			entity,
			clustersEntity,
			"reclaimFreeSpace",
			clusterName,
			"--force",
		)

		cmd.Env = os.Environ()
		resp, err := cmd.CombinedOutput()
		require.NoError(t, err, string(resp))
	})

	t.Run("Watch", watchAutomation(cliPath))

	t.Run("Restart", func(t *testing.T) {
		cmd := exec.Command(cliPath,
			entity,
			clustersEntity,
			"restart",
			clusterName,
			"--force",
		)

		cmd.Env = os.Environ()
		resp, err := cmd.CombinedOutput()
		require.NoError(t, err, string(resp))
	})

	t.Run("Watch", watchAutomation(cliPath))

	t.Run("Shutdown", func(t *testing.T) {
		cmd := exec.Command(cliPath,
			entity,
			clustersEntity,
			"shutdown",
			clusterName,
			"--force",
		)

		cmd.Env = os.Environ()
		resp, err := cmd.CombinedOutput()
		require.NoError(t, err, string(resp))
	})

	t.Run("Watch", watchAutomation(cliPath))

	t.Run("Unmanage", func(t *testing.T) {
		cmd := exec.Command(cliPath,
			entity,
			clustersEntity,
			"unmanage",
			clusterName,
			"--force",
		)

		cmd.Env = os.Environ()
		resp, err := cmd.CombinedOutput()
		require.NoError(t, err, string(resp))
	})

	t.Run("Watch", watchAutomation(cliPath))

	t.Run("Stop Monitoring", func(t *testing.T) {
		ids, err := hostIDs(cliPath)
		require.NoError(t, err)
		for _, h := range ids {
			cmd := exec.Command(cliPath,
				entity,
				monitoringEntity,
				"rm",
				h,
				"--force",
			)

			cmd.Env = os.Environ()
			resp, err := cmd.CombinedOutput()
			require.NoError(t, err, string(resp))
		}
	})
}

func TestDeployAndDeleteReplicaSet(t *testing.T) {
	cliPath, err := e2e.Bin()
	require.NoError(t, err)

	const testFile = "om-new-cluster.json"

	n, err := e2e.RandInt(1000)
	require.NoError(t, err)
	clusterName := fmt.Sprintf("e2e-cluster-%v", n)
	hostname, err := automationServerHostname(cliPath)
	require.NoError(t, err)
	require.NoError(
		t,
		generateRSConfig(testFile, hostname, clusterName, testedMDBVersion, testedMDBFCV),
	)
	t.Run("Apply", func(t *testing.T) {
		cmd := exec.Command(cliPath,
			entity,
			clustersEntity,
			"apply",
			"-f",
			testFile,
		)

		cmd.Env = os.Environ()
		resp, err := cmd.CombinedOutput()
		require.NoError(t, err, string(resp))
	})

	t.Run("Watch", watchAutomation(cliPath))

	t.Run("Delete Cluster", func(t *testing.T) {
		cmd := exec.Command(cliPath,
			entity,
			clustersEntity,
			"delete",
			clusterName,
			"--force",
		)

		cmd.Env = os.Environ()
		resp, err := cmd.CombinedOutput()
		require.NoError(t, err, string(resp))
	})
}
