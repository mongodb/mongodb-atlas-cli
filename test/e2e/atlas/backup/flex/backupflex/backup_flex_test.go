// Copyright 2024 MongoDB Inc
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

package backupflex

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"testing"

	"github.com/mongodb/mongodb-atlas-cli/atlascli/test/internal"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	atlasv2 "go.mongodb.org/atlas-sdk/v20250312009/admin"
)

const (
	clustersEntity  = "clusters"
	backupsEntity   = "backups"
	snapshotsEntity = "snapshots"
	restoresEntity  = "restores"
)

// Note that the FlexClusters are only available in the 5efda6aea3f2ed2e7dd6ce05 (Atlas CLI E2E Project)
// They will be fully enabled in https://jira.mongodb.org/browse/CLOUDP-291186. We will be able to move these e2e tests
// to create their project once the ticket is completed.
func TestFlexBackup(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping test in short mode")
	}

	g := internal.NewAtlasE2ETestGenerator(t, internal.WithSnapshot())
	cliPath, err := internal.AtlasCLIBin()
	require.NoError(t, err)

	profile, err := internal.ProfileData()
	require.NoError(t, err)

	g.ProjectID = profile["project_id"]
	generateFlexCluster(t, g)

	clusterName, err := internal.FlexInstanceName()
	require.NoError(t, err)

	var snapshotID string
	g.Run("Snapshot List", func(t *testing.T) { //nolint:thelper // g.Run replaces t.Run
		cmd := exec.Command(cliPath,
			backupsEntity,
			snapshotsEntity,
			"list",
			clusterName,
			"-o=json",
			"-P",
			internal.ProfileName(),
		)
		cmd.Env = os.Environ()
		resp, err := internal.RunAndGetStdOut(cmd)
		require.NoError(t, err, string(resp))

		var r atlasv2.PaginatedApiAtlasFlexBackupSnapshot20241113
		require.NoError(t, json.Unmarshal(resp, &r), string(resp))
		assert.NotEmpty(t, r)
		snapshotID = r.GetResults()[0].GetId()
		t.Log("snapshotID", snapshotID, "snapshotStatus", *r.GetResults()[0].Status)
		require.NotEmpty(t, snapshotID)
	})

	g.Run("Snapshot Describe", func(t *testing.T) { //nolint:thelper // g.Run replaces t.Run
		cmd := exec.Command(cliPath,
			backupsEntity,
			snapshotsEntity,
			"describe",
			snapshotID,
			"--clusterName",
			clusterName,
			"-o=json",
			"-P",
			internal.ProfileName(),
		)
		cmd.Env = os.Environ()
		resp, err := internal.RunAndGetStdOut(cmd)

		require.NoError(t, err, string(resp))
		var result atlasv2.FlexBackupSnapshot20241113
		require.NoError(t, json.Unmarshal(resp, &result), string(resp))
		assert.Equal(t, snapshotID, result.GetId())
	})

	g.Run("Snapshot Watch", func(t *testing.T) { //nolint:thelper // g.Run replaces t.Run
		cmd := exec.Command(cliPath,
			backupsEntity,
			snapshotsEntity,
			"watch",
			snapshotID,
			"--clusterName",
			clusterName,
			"-P",
			internal.ProfileName(),
		)
		cmd.Env = os.Environ()
		resp, err := internal.RunAndGetStdOut(cmd)

		require.NoError(t, err, string(resp))
	})

	var restoreJobID string

	g.Run("Restores Create - Automated", func(t *testing.T) { //nolint:thelper // g.Run replaces t.Run
		cmd := exec.Command(cliPath,
			backupsEntity,
			restoresEntity,
			"start",
			"automated",
			"--clusterName",
			clusterName,
			"--snapshotId",
			snapshotID,
			"--targetClusterName",
			g.ClusterName,
			"--targetProjectId",
			g.ProjectID,
			"-o=json",
			"-P",
			internal.ProfileName(),
		)
		cmd.Env = os.Environ()
		resp, err := internal.RunAndGetStdOut(cmd)

		require.NoError(t, err, string(resp))
		var result atlasv2.FlexBackupRestoreJob20241113
		require.NoError(t, json.Unmarshal(resp, &result), string(resp))
		restoreJobID = result.GetId()
		t.Log("restoreJobId", restoreJobID)

		require.NotEmpty(t, restoreJobID)
	})

	g.Run("Restores Watch", func(t *testing.T) { //nolint:thelper // g.Run replaces t.Run
		cmd := exec.Command(cliPath,
			backupsEntity,
			restoresEntity,
			"watch",
			restoreJobID,
			"--clusterName",
			clusterName,
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
			clusterName,
			"-o=json",
			"-P",
			internal.ProfileName(),
		)
		cmd.Env = os.Environ()
		resp, err := internal.RunAndGetStdOut(cmd)
		require.NoError(t, err, string(resp))

		var result atlasv2.PaginatedApiAtlasFlexBackupRestoreJob20241113
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
			clusterName,
			"-o=json",
			"-P",
			internal.ProfileName(),
		)
		cmd.Env = os.Environ()
		resp, err := internal.RunAndGetStdOut(cmd)
		require.NoError(t, err, string(resp))
		var result atlasv2.FlexBackupRestoreJob20241113
		require.NoError(t, json.Unmarshal(resp, &result))
		assert.NotEmpty(t, result)
	})

	g.Run("Restores Create - Download", func(t *testing.T) { //nolint:thelper // g.Run replaces t.Run
		cmd := exec.Command(cliPath,
			backupsEntity,
			restoresEntity,
			"start",
			"download",
			"--clusterName",
			clusterName,
			"--snapshotId",
			snapshotID,
			"--targetClusterName",
			g.ClusterName,
			"--targetProjectId",
			g.ProjectID,
			"-o=json",
			"-P",
			internal.ProfileName(),
		)
		cmd.Env = os.Environ()
		resp, err := internal.RunAndGetStdOut(cmd)

		require.NoError(t, err, string(resp))
		var result atlasv2.FlexBackupRestoreJob20241113
		require.NoError(t, json.Unmarshal(resp, &result), string(resp))
		restoreJobID = result.GetId()
		t.Log("snapshotID", restoreJobID)
		require.NotEmpty(t, restoreJobID)
	})

	g.Run("Restores Watch - Download", func(t *testing.T) { //nolint:thelper // g.Run replaces t.Run
		cmd := exec.Command(cliPath,
			backupsEntity,
			restoresEntity,
			"watch",
			restoreJobID,
			"--clusterName",
			clusterName,
			"-o=json",
			"-P",
			internal.ProfileName(),
		)
		cmd.Env = os.Environ()
		resp, err := internal.RunAndGetStdOut(cmd)

		require.NoError(t, err, string(resp))
	})

	g.Run("Delete flex cluster", func(t *testing.T) { //nolint:thelper // g.Run replaces t.Run
		cmd := exec.Command(cliPath,
			clustersEntity,
			"delete",
			g.ClusterName,
			"--force",
			"--watch",
			"-P",
			internal.ProfileName(),
		)
		cmd.Env = os.Environ()
		resp, err := internal.RunAndGetStdOut(cmd)
		require.NoError(t, err, string(resp))

		expected := fmt.Sprintf("Deleting cluster '%s'Cluster deleted\n", g.ClusterName)
		assert.Equal(t, expected, string(resp))
	})
}

func generateFlexCluster(t *testing.T, g *internal.AtlasE2ETestGenerator) {
	t.Helper()

	if g.ProjectID == "" {
		t.Fatal("unexpected error: project must be generated")
	}

	g.ClusterName = g.Memory("generateFlexClusterName", internal.Must(internal.RandClusterName())).(string)

	err := deployFlexClusterForProject(g.ProjectID, g.ClusterName)
	if err != nil {
		t.Fatalf("unexpected error deploying flex cluster: %v", err)
	}
	t.Logf("flexClusterName=%s", g.ClusterName)

	if internal.SkipCleanup() {
		return
	}

	t.Cleanup(func() {
		_ = internal.DeleteClusterForProject(g.ProjectID, g.ClusterName)
	})
}

func deployFlexClusterForProject(projectID, clusterName string) error {
	cliPath, err := internal.AtlasCLIBin()
	if err != nil {
		return err
	}

	args := []string{
		clustersEntity,
		"create",
		clusterName,
		"--region", "US_EAST_1",
		"--provider", "AWS",
		"-P",
		internal.ProfileName(),
	}

	if projectID != "" {
		args = append(args, "--projectId", projectID)
	}

	create := exec.Command(cliPath, args...)
	create.Env = os.Environ()
	if resp, err := internal.RunAndGetStdOut(create); err != nil {
		return fmt.Errorf("error creating flex cluster (%s): %w - %s", clusterName, err, string(resp))
	}

	watchArgs := []string{
		clustersEntity,
		"watch",
		clusterName,
		"-P",
		internal.ProfileName(),
	}

	if projectID != "" {
		watchArgs = append(watchArgs, "--projectId", projectID)
	}

	watch := exec.Command(cliPath, watchArgs...)
	watch.Env = os.Environ()
	if resp, err := internal.RunAndGetStdOut(watch); err != nil {
		return fmt.Errorf("error watching cluster %w: %s", err, string(resp))
	}

	return nil
}
