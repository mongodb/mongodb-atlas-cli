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

package backuprestores

import (
	"encoding/json"
	"os"
	"os/exec"
	"testing"

	"github.com/mongodb/mongodb-atlas-cli/atlascli/test/internal"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	atlasv2 "go.mongodb.org/atlas-sdk/v20250312009/admin"
)

const (
	backupsEntity   = "backups"
	snapshotsEntity = "snapshots"
	restoresEntity  = "restores"
)

func TestRestores(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping test in short mode")
	}

	cliPath, err := internal.AtlasCLIBin()
	require.NoError(t, err)

	var snapshotID, restoreJobID string

	g := internal.NewAtlasE2ETestGenerator(t, internal.WithSnapshot(), internal.WithBackup())
	g.GenerateProjectAndCluster("backupRestores")
	require.NotEmpty(t, g.ClusterName)

	sourceProjectID := g.ProjectID
	sourceClusterName := g.ClusterName

	g.ProjectID = ""
	g.ClusterName = ""

	g.GenerateProjectAndCluster("backupRestores2")
	require.NotEmpty(t, g.ClusterName)

	targetProjectID := g.ProjectID
	targetClusterName := g.ClusterName

	t.Cleanup(func() {
		require.NoError(t, internal.DeleteClusterForProject(sourceProjectID, sourceClusterName))
		internal.DeleteProjectWithRetry(t, sourceProjectID)
		require.NoError(t, internal.DeleteClusterForProject(targetProjectID, targetClusterName))
		internal.DeleteProjectWithRetry(t, targetProjectID)
	})

	g.Run("Create snapshot", func(t *testing.T) { //nolint:thelper // g.Run replaces t.Run
		cmd := exec.Command(cliPath,
			backupsEntity,
			snapshotsEntity,
			"create",
			sourceClusterName,
			"--desc",
			"test-snapshot",
			"--projectId",
			sourceProjectID,
			"-o=json",
			"-P",
			internal.ProfileName(),
		)
		cmd.Env = os.Environ()
		resp, err := internal.RunAndGetStdOut(cmd)

		require.NoError(t, err, string(resp))

		var snapshot atlasv2.DiskBackupSnapshot
		require.NoError(t, json.Unmarshal(resp, &snapshot))
		assert.Equal(t, "test-snapshot", snapshot.GetDescription())
		snapshotID = snapshot.GetId()
	})

	g.Run("Watch snapshot creation", func(t *testing.T) { //nolint:thelper // g.Run replaces t.Run
		cmd := exec.Command(cliPath,
			backupsEntity,
			snapshotsEntity,
			"watch",
			snapshotID,
			"--clusterName",
			sourceClusterName,
			"--projectId",
			sourceProjectID,
			"-P",
			internal.ProfileName(),
		)
		cmd.Env = os.Environ()
		resp, _ := internal.RunAndGetStdOut(cmd)
		t.Log(string(resp))
	})

	g.Run("Restores Create - Automated", func(t *testing.T) { //nolint:thelper // g.Run replaces t.Run
		cmd := exec.Command(cliPath,
			backupsEntity,
			restoresEntity,
			"start",
			"automated",
			"--clusterName",
			sourceClusterName,
			"--snapshotId",
			snapshotID,
			"--projectId",
			sourceProjectID,
			"--targetProjectId",
			targetProjectID,
			"--targetClusterName",
			targetClusterName,
			"-o=json",
			"-P",
			internal.ProfileName(),
		)
		cmd.Env = os.Environ()
		resp, err := internal.RunAndGetStdOut(cmd)

		require.NoError(t, err, string(resp))
		var result atlasv2.DiskBackupSnapshotRestoreJob
		require.NoError(t, json.Unmarshal(resp, &result))
		restoreJobID = result.GetId()
	})

	g.Run("Restores Watch", func(t *testing.T) { //nolint:thelper // g.Run replaces t.Run
		cmd := exec.Command(cliPath,
			backupsEntity,
			restoresEntity,
			"watch",
			restoreJobID,
			"--clusterName",
			sourceClusterName,
			"--projectId",
			sourceProjectID,
			"-o=json",
			"-P",
			internal.ProfileName(),
		)
		cmd.Env = os.Environ()
		resp, err := internal.RunAndGetStdOut(cmd)

		require.NoError(t, err, string(resp))
	})

	g.Run("Restores List", func(t *testing.T) { //nolint:thelper // g.Run replaces t.Run
		cmd := exec.Command(cliPath,
			backupsEntity,
			restoresEntity,
			"list",
			sourceClusterName,
			"--projectId",
			sourceProjectID,
			"-o=json",
			"-P",
			internal.ProfileName(),
		)
		cmd.Env = os.Environ()
		resp, err := internal.RunAndGetStdOut(cmd)

		require.NoError(t, err, string(resp))
		var result atlasv2.PaginatedCloudBackupRestoreJob
		require.NoError(t, json.Unmarshal(resp, &result), string(resp))
		assert.NotEmpty(t, result)
	})

	g.Run("Restores Describe", func(t *testing.T) { //nolint:thelper // g.Run replaces t.Run
		cmd := exec.Command(cliPath,
			backupsEntity,
			restoresEntity,
			"describe",
			restoreJobID,
			"--clusterName",
			sourceClusterName,
			"--projectId",
			sourceProjectID,
			"-o=json",
			"-P",
			internal.ProfileName(),
		)
		cmd.Env = os.Environ()
		resp, err := internal.RunAndGetStdOut(cmd)

		require.NoError(t, err, string(resp))

		var result atlasv2.DiskBackupSnapshotRestoreJob
		require.NoError(t, json.Unmarshal(resp, &result), string(resp))
		assert.NotEmpty(t, result)
	})

	g.Run("Restores Create - Download", func(t *testing.T) { //nolint:thelper // g.Run replaces t.Run
		cmd := exec.Command(cliPath,
			backupsEntity,
			restoresEntity,
			"start",
			"download",
			"--clusterName",
			sourceClusterName,
			"--snapshotId",
			snapshotID,
			"--projectId",
			sourceProjectID,
			"-o=json",
			"-P",
			internal.ProfileName(),
		)
		cmd.Env = os.Environ()
		resp, err := internal.RunAndGetStdOut(cmd)

		require.NoError(t, err, string(resp))
		var result atlasv2.DiskBackupSnapshotRestoreJob
		require.NoError(t, json.Unmarshal(resp, &result))
		restoreJobID = result.GetId()
	})

	g.Run("Restores Watch - Download", func(t *testing.T) { //nolint:thelper // g.Run replaces t.Run
		cmd := exec.Command(cliPath,
			backupsEntity,
			restoresEntity,
			"watch",
			restoreJobID,
			"--clusterName",
			sourceClusterName,
			"--projectId",
			sourceProjectID,
			"-o=json",
			"-P",
			internal.ProfileName(),
		)
		cmd.Env = os.Environ()
		resp, err := internal.RunAndGetStdOut(cmd)

		require.NoError(t, err, string(resp))
	})

	g.Run("Delete snapshot", func(t *testing.T) { //nolint:thelper // g.Run replaces t.Run
		cmd := exec.Command(cliPath,
			backupsEntity,
			snapshotsEntity,
			"delete",
			snapshotID,
			"--clusterName",
			sourceClusterName,
			"--projectId",
			sourceProjectID,
			"--force",
			"-P",
			internal.ProfileName(),
		)
		cmd.Env = os.Environ()
		resp, err := internal.RunAndGetStdOut(cmd)
		require.NoError(t, err, string(resp))
	})

	if internal.SkipCleanup() {
		return
	}

	g.Run("Watch snapshot deletion", func(t *testing.T) { //nolint:thelper // g.Run replaces t.Run
		cmd := exec.Command(cliPath,
			backupsEntity,
			snapshotsEntity,
			"watch",
			snapshotID,
			"--clusterName",
			sourceClusterName,
			"--projectId",
			sourceProjectID,
			"-P",
			internal.ProfileName(),
		)
		cmd.Env = os.Environ()
		resp, _ := internal.RunAndGetStdOut(cmd)
		t.Log(string(resp))
	})
}
