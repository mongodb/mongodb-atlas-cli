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

//go:build e2e || (atlas && backup && restores)

package e2e_test

import (
	"encoding/json"
	"os"
	"os/exec"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	atlasv2 "go.mongodb.org/atlas-sdk/v20250312001/admin"
)

func TestRestores(t *testing.T) {
	cliPath, err := AtlasCLIBin()
	require.NoError(t, err)

	var snapshotID, restoreJobID string

	g := newAtlasE2ETestGenerator(t, withSnapshot(), withBackup())
	g.generateProjectAndCluster("backupRestores")
	require.NotEmpty(t, g.clusterName)

	projectID := g.projectID
	clusterName := g.clusterName

	g.projectID = ""
	g.clusterName = ""

	g.generateProjectAndCluster("backupRestores2")
	require.NotEmpty(t, g.clusterName)

	projectID2 := g.projectID
	clusterName2 := g.clusterName

	t.Run("Create snapshot", func(t *testing.T) {
		cmd := exec.Command(cliPath,
			backupsEntity,
			snapshotsEntity,
			"create",
			clusterName,
			"--desc",
			"test-snapshot",
			"--projectId",
			projectID,
			"-o=json")
		cmd.Env = os.Environ()
		resp, err := RunAndGetStdOut(cmd)

		require.NoError(t, err, string(resp))

		var snapshot atlasv2.DiskBackupSnapshot
		require.NoError(t, json.Unmarshal(resp, &snapshot))
		assert.Equal(t, "test-snapshot", snapshot.GetDescription())
		snapshotID = snapshot.GetId()
	})

	t.Run("Watch snapshot creation", func(t *testing.T) {
		cmd := exec.Command(cliPath,
			backupsEntity,
			snapshotsEntity,
			"watch",
			snapshotID,
			"--clusterName",
			clusterName,
			"--projectId",
			projectID)
		cmd.Env = os.Environ()
		resp, _ := RunAndGetStdOut(cmd)
		t.Log(string(resp))
	})

	t.Run("Restores Create - Automated", func(t *testing.T) {
		cmd := exec.Command(cliPath,
			backupsEntity,
			restoresEntity,
			"start",
			"automated",
			"--clusterName",
			clusterName,
			"--snapshotId",
			snapshotID,
			"--projectId",
			projectID,
			"--targetProjectId",
			projectID2,
			"--targetClusterName",
			clusterName2,
			"-o=json")
		cmd.Env = os.Environ()
		resp, err := RunAndGetStdOut(cmd)

		require.NoError(t, err, string(resp))
		var result atlasv2.DiskBackupSnapshotRestoreJob
		require.NoError(t, json.Unmarshal(resp, &result))
		restoreJobID = result.GetId()
	})

	t.Run("Restores Watch", func(t *testing.T) {
		cmd := exec.Command(cliPath,
			backupsEntity,
			restoresEntity,
			"watch",
			restoreJobID,
			"--clusterName",
			clusterName,
			"--projectId",
			projectID,
			"-o=json")
		cmd.Env = os.Environ()
		resp, err := RunAndGetStdOut(cmd)

		require.NoError(t, err, string(resp))
	})

	t.Run("Restores List", func(t *testing.T) {
		cmd := exec.Command(cliPath,
			backupsEntity,
			restoresEntity,
			"list",
			clusterName,
			"--projectId",
			projectID,
			"-o=json")
		cmd.Env = os.Environ()
		resp, err := RunAndGetStdOut(cmd)

		require.NoError(t, err, string(resp))
		var result atlasv2.PaginatedCloudBackupRestoreJob
		require.NoError(t, json.Unmarshal(resp, &result), string(resp))
		assert.NotEmpty(t, result)
	})

	t.Run("Restores Describe", func(t *testing.T) {
		cmd := exec.Command(cliPath,
			backupsEntity,
			restoresEntity,
			"describe",
			restoreJobID,
			"--clusterName",
			clusterName,
			"--projectId",
			projectID,
			"-o=json")
		cmd.Env = os.Environ()
		resp, err := RunAndGetStdOut(cmd)

		require.NoError(t, err, string(resp))

		var result atlasv2.DiskBackupSnapshotRestoreJob
		require.NoError(t, json.Unmarshal(resp, &result), string(resp))
		assert.NotEmpty(t, result)
	})

	t.Run("Restores Create - Download", func(t *testing.T) {
		cmd := exec.Command(cliPath,
			backupsEntity,
			restoresEntity,
			"start",
			"download",
			"--clusterName",
			clusterName,
			"--snapshotId",
			snapshotID,
			"--projectId",
			projectID,
			"-o=json")
		cmd.Env = os.Environ()
		resp, err := RunAndGetStdOut(cmd)

		require.NoError(t, err, string(resp))
		var result atlasv2.DiskBackupSnapshotRestoreJob
		require.NoError(t, json.Unmarshal(resp, &result))
		restoreJobID = result.GetId()
	})

	t.Run("Restores Watch - Download", func(t *testing.T) {
		cmd := exec.Command(cliPath,
			backupsEntity,
			restoresEntity,
			"watch",
			restoreJobID,
			"--clusterName",
			clusterName,
			"--projectId",
			projectID,
			"-o=json")
		cmd.Env = os.Environ()
		resp, err := RunAndGetStdOut(cmd)

		require.NoError(t, err, string(resp))
	})

	t.Run("Delete snapshot", func(t *testing.T) {
		cmd := exec.Command(cliPath,
			backupsEntity,
			snapshotsEntity,
			"delete",
			snapshotID,
			"--clusterName",
			clusterName,
			"--projectId",
			projectID,
			"--force")
		cmd.Env = os.Environ()
		resp, err := RunAndGetStdOut(cmd)
		require.NoError(t, err, string(resp))
	})

	if skipCleanup() {
		return
	}

	t.Run("Watch snapshot deletion", func(t *testing.T) {
		cmd := exec.Command(cliPath,
			backupsEntity,
			snapshotsEntity,
			"watch",
			snapshotID,
			"--clusterName",
			clusterName,
			"--projectId",
			projectID)
		cmd.Env = os.Environ()
		resp, _ := RunAndGetStdOut(cmd)
		t.Log(string(resp))
	})
}
