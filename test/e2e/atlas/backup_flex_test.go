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
//go:build e2e || (atlas && backup && flex)

package atlas_test

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"testing"

	"github.com/mongodb/mongodb-atlas-cli/atlascli/test/e2e"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	atlasv2 "go.mongodb.org/atlas-sdk/v20241113004/admin"
)

// Note that the FlexClusters are only available in the 5efda6aea3f2ed2e7dd6ce05 (Atlas CLI E2E Project)
// They will be fully enabled in https://jira.mongodb.org/browse/CLOUDP-291186. We will be able to move these e2e tests
// to create their project once the ticket is completed.
func TestFlexBackup(t *testing.T) {
	cliPath, err := e2e.AtlasCLIBin()
	require.NoError(t, err)

	g := newAtlasE2ETestGenerator(t)
	g.projectID = os.Getenv("MCLI_PROJECT_ID")
	g.generateFlexCluster()

	clusterName := os.Getenv("E2E_FLEX_INSTANCE_NAME")
	require.NotEmpty(t, clusterName)

	var snapshotID string
	t.Run("Snapshot List", func(t *testing.T) {
		cmd := exec.Command(cliPath,
			backupsEntity,
			snapshotsEntity,
			"list",
			clusterName,
			"-o=json")
		cmd.Env = os.Environ()
		resp, err := e2e.RunAndGetStdOut(cmd)
		require.NoError(t, err, string(resp))

		var r atlasv2.PaginatedApiAtlasFlexBackupSnapshot20241113
		require.NoError(t, json.Unmarshal(resp, &r), string(resp))
		assert.NotEmpty(t, r)
		snapshotID = r.GetResults()[0].GetId()
		t.Log("snapshotID", snapshotID, "snapshotStatus", *r.GetResults()[0].Status)
		require.NotEmpty(t, snapshotID)
	})

	t.Run("Snapshot Describe", func(t *testing.T) {
		cmd := exec.Command(cliPath,
			backupsEntity,
			snapshotsEntity,
			"describe",
			snapshotID,
			"--clusterName",
			clusterName,
			"-o=json")
		cmd.Env = os.Environ()
		resp, err := e2e.RunAndGetStdOut(cmd)

		require.NoError(t, err, string(resp))
		var result atlasv2.FlexBackupSnapshot20241113
		require.NoError(t, json.Unmarshal(resp, &result), string(resp))
		assert.Equal(t, snapshotID, result.GetId())
	})

	t.Run("Snapshot Watch", func(t *testing.T) {
		cmd := exec.Command(cliPath,
			backupsEntity,
			snapshotsEntity,
			"watch",
			snapshotID,
			"--clusterName",
			clusterName)
		cmd.Env = os.Environ()
		resp, err := e2e.RunAndGetStdOut(cmd)

		require.NoError(t, err, string(resp))
	})

	var restoreJobID string

	t.Run("Restores Create", func(t *testing.T) {
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
			g.clusterName,
			"--targetProjectId",
			g.projectID,
			"-o=json")
		cmd.Env = os.Environ()
		resp, err := e2e.RunAndGetStdOut(cmd)

		require.NoError(t, err, string(resp))
		var result atlasv2.FlexBackupRestoreJob20241113
		require.NoError(t, json.Unmarshal(resp, &result), string(resp))
		restoreJobID = result.GetId()
		t.Log("snapshotID", restoreJobID)
		require.NotEmpty(t, restoreJobID)
	})

	t.Run("Restores Watch", func(t *testing.T) {
		cmd := exec.Command(cliPath,
			backupsEntity,
			restoresEntity,
			"watch",
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
			backupsEntity,
			restoresEntity,
			"list",
			clusterName,
			"-o=json")
		cmd.Env = os.Environ()
		resp, err := e2e.RunAndGetStdOut(cmd)
		require.NoError(t, err, string(resp))

		var result atlasv2.PaginatedApiAtlasFlexBackupRestoreJob20241113
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
			"-o=json")
		cmd.Env = os.Environ()
		resp, err := e2e.RunAndGetStdOut(cmd)
		require.NoError(t, err, string(resp))
		var result atlasv2.FlexBackupRestoreJob20241113
		require.NoError(t, json.Unmarshal(resp, &result))
		assert.NotEmpty(t, result)
	})

	t.Run("Delete flex cluster", func(t *testing.T) {
		cmd := exec.Command(cliPath,
			clustersEntity,
			"delete",
			g.clusterName,
			"--force",
			"--watch")
		cmd.Env = os.Environ()
		resp, err := e2e.RunAndGetStdOut(cmd)
		require.NoError(t, err, string(resp))

		expected := fmt.Sprintf("Deleting cluster '%s'Cluster deleted\n", g.clusterName)
		assert.Equal(t, expected, string(resp))
	})
}
