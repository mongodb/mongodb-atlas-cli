// Copyright 2023 MongoDB Inc
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
//go:build e2e || (atlas && backup && serverless)

package atlas_test

import (
	"encoding/json"
	"os"
	"os/exec"
	"testing"

	"github.com/mongodb/mongodb-atlas-cli/atlascli/test/e2e"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	atlasv2 "go.mongodb.org/atlas-sdk/v20240530002/admin"
)

func TestServerlessBackup(t *testing.T) {
	cliPath, err := e2e.AtlasCLIBin()
	require.NoError(t, err)

	clusterName := os.Getenv("E2E_SERVERLESS_INSTANCE_NAME")
	require.NotEmpty(t, clusterName)

	var snapshotID string

	g := newAtlasE2ETestGenerator(t)
	g.generateProject("serverlessBackup")
	g.generateServerlessCluster()

	t.Run("Snapshot List", func(t *testing.T) {
		cmd := exec.Command(cliPath,
			serverlessEntity,
			backupsEntity,
			snapshotsEntity,
			"list",
			clusterName,
			"-o=json")
		cmd.Env = os.Environ()
		resp, err := e2e.RunAndGetStdOut(cmd)
		require.NoError(t, err, string(resp))

		var r atlasv2.PaginatedApiAtlasServerlessBackupSnapshot
		require.NoError(t, json.Unmarshal(resp, &r), string(resp))
		assert.NotEmpty(t, r)
		snapshotID = r.GetResults()[0].GetId()
		t.Log("snapshotID", snapshotID)
		require.NotEmpty(t, snapshotID)
	})

	t.Run("Snapshot Describe", func(t *testing.T) {
		cmd := exec.Command(cliPath,
			serverlessEntity,
			backupsEntity,
			snapshotsEntity,
			"describe",
			"--snapshotId",
			snapshotID,
			"--clusterName",
			clusterName,
			"-o=json")
		cmd.Env = os.Environ()
		resp, err := e2e.RunAndGetStdOut(cmd)

		require.NoError(t, err, string(resp))
		var result atlasv2.ServerlessBackupSnapshot
		require.NoError(t, json.Unmarshal(resp, &result), string(resp))
		assert.Equal(t, snapshotID, result.GetId())
	})

	var restoreJobID string

	t.Run("Restores Create", func(t *testing.T) {
		cmd := exec.Command(cliPath,
			serverlessEntity,
			backupsEntity,
			restoresEntity,
			"create",
			"--deliveryType",
			"automated",
			"--clusterName",
			clusterName,
			"--snapshotId",
			snapshotID,
			"--targetClusterName",
			g.serverlessName,
			"--targetProjectId",
			g.projectID,
			"-o=json")
		cmd.Env = os.Environ()
		resp, err := e2e.RunAndGetStdOut(cmd)

		require.NoError(t, err, string(resp))
		var result atlasv2.ServerlessBackupRestoreJob
		require.NoError(t, json.Unmarshal(resp, &result), string(resp))
		restoreJobID = result.GetId()
		t.Log("snapshotID", restoreJobID)
		require.NotEmpty(t, restoreJobID)
	})

	t.Run("Restores Watch", func(t *testing.T) {
		cmd := exec.Command(cliPath,
			serverlessEntity,
			backupsEntity,
			restoresEntity,
			"watch",
			"--restoreJobId",
			restoreJobID,
			"--clusterName",
			clusterName,
			"-o=json")
		cmd.Env = os.Environ()
		resp, err := e2e.RunAndGetStdOut(cmd)

		require.NoError(t, err, string(resp))
	})

	t.Run("Restores List", func(t *testing.T) {
		cmd := exec.Command(cliPath,
			serverlessEntity,
			backupsEntity,
			restoresEntity,
			"list",
			clusterName,
			"-o=json")
		cmd.Env = os.Environ()
		resp, err := e2e.RunAndGetStdOut(cmd)
		require.NoError(t, err, string(resp))

		var result atlasv2.PaginatedApiAtlasServerlessBackupRestoreJob
		require.NoError(t, json.Unmarshal(resp, &result), string(resp))
		assert.NotEmpty(t, result)
	})

	t.Run("Restores Describe", func(t *testing.T) {
		cmd := exec.Command(cliPath,
			serverlessEntity,
			backupsEntity,
			restoresEntity,
			"describe",
			"--restoreJobId",
			restoreJobID,
			"--clusterName",
			clusterName,
			"-o=json")
		cmd.Env = os.Environ()
		resp, err := e2e.RunAndGetStdOut(cmd)
		require.NoError(t, err, string(resp))
		var result atlasv2.ServerlessBackupRestoreJob
		require.NoError(t, json.Unmarshal(resp, &result))
		assert.NotEmpty(t, result)
	})
}
