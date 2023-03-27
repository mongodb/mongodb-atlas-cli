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

package atlas_test

import (
	"encoding/json"
	"os"
	"os/exec"
	"testing"

	"github.com/mongodb/mongodb-atlas-cli/test/e2e"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	atlas "go.mongodb.org/atlas/mongodbatlas"
)

func TestRestores(t *testing.T) {
	cliPath, err := e2e.AtlasCLIBin()
	r := require.New(t)
	r.NoError(err)

	var snapshotID, restoreJobID string

	g := newAtlasE2ETestGenerator(t)
	g.enableBackup = true
	g.generateProjectAndCluster("backupRestores")

	t.Run("Create snapshot", func(t *testing.T) {
		cmd := exec.Command(cliPath,
			backupsEntity,
			snapshotsEntity,
			"create",
			g.clusterName,
			"--desc",
			"test-snapshot",
			"--projectId",
			g.projectID,
			"-o=json")
		cmd.Env = os.Environ()
		resp, err := cmd.CombinedOutput()

		r.NoError(err, string(resp))

		a := assert.New(t)
		var snapshot atlas.CloudProviderSnapshot
		if err = json.Unmarshal(resp, &snapshot); a.NoError(err) {
			a.Equal("test-snapshot", snapshot.Description)
		}
		snapshotID = snapshot.ID
	})

	t.Run("Watch snapshot creation", func(t *testing.T) {
		cmd := exec.Command(cliPath,
			backupsEntity,
			snapshotsEntity,
			"watch",
			snapshotID,
			"--clusterName",
			g.clusterName,
			"--projectId",
			g.projectID)
		cmd.Env = os.Environ()
		resp, _ := cmd.CombinedOutput()
		t.Log(string(resp))
	})

	t.Run("Restores Create", func(t *testing.T) {
		cmd := exec.Command(cliPath,
			serverlessEntity,
			backupsEntity,
			restoresEntity,
			"start",
			"download",
			"--clusterName",
			g.clusterName,
			"--snapshotId",
			snapshotID,
			"--projectId",
			g.projectID,
			"-o=json")
		cmd.Env = os.Environ()
		resp, err := cmd.CombinedOutput()

		r.NoError(err, string(resp))
		a := assert.New(t)
		var result atlas.CloudProviderSnapshotRestoreJob
		if err = json.Unmarshal(resp, &result); a.NoError(err) {
			restoreJobID = result.ID
		}
	})

	t.Run("Restores Watch", func(t *testing.T) {
		cmd := exec.Command(cliPath,
			backupsEntity,
			restoresEntity,
			"watch",
			"--restoreJobId",
			restoreJobID,
			"--clusterName",
			g.clusterName,
			"--projectId",
			g.projectID,
			"-o=json")
		cmd.Env = os.Environ()
		resp, err := cmd.CombinedOutput()

		r.NoError(err, string(resp))
	})

	t.Run("Restores List", func(t *testing.T) {
		cmd := exec.Command(cliPath,
			backupsEntity,
			restoresEntity,
			"list",
			g.clusterName,
			"--projectId",
			g.projectID,
			"-o=json")
		cmd.Env = os.Environ()
		resp, err := cmd.CombinedOutput()

		r.NoError(err, string(resp))

		a := assert.New(t)
		var result atlas.CloudProviderSnapshotRestoreJobs
		if err = json.Unmarshal(resp, &result); a.NoError(err) {
			a.NotEmpty(result)
		}
	})

	t.Run("Restores Describe", func(t *testing.T) {
		cmd := exec.Command(cliPath,
			backupsEntity,
			restoresEntity,
			"describe",
			"--restoreJobId",
			restoreJobID,
			"--clusterName",
			g.clusterName,
			"--projectId",
			g.projectID,
			"-o=json")
		cmd.Env = os.Environ()
		resp, err := cmd.CombinedOutput()

		r.NoError(err, string(resp))

		a := assert.New(t)
		var result atlas.CloudProviderSnapshotRestoreJob
		if err = json.Unmarshal(resp, &result); a.NoError(err) {
			a.NotEmpty(result)
		}
	})

	t.Run("Delete snapshot", func(t *testing.T) {
		cmd := exec.Command(cliPath,
			backupsEntity,
			snapshotsEntity,
			"delete",
			snapshotID,
			"--clusterName",
			g.clusterName,
			"--projectId",
			g.projectID,
			"--force")
		cmd.Env = os.Environ()
		resp, err := cmd.CombinedOutput()
		r.NoError(err, string(resp))
	})

	t.Run("Watch snapshot deletion", func(t *testing.T) {
		cmd := exec.Command(cliPath,
			backupsEntity,
			snapshotsEntity,
			"watch",
			snapshotID,
			"--clusterName",
			g.clusterName,
			"--projectId",
			g.projectID)
		cmd.Env = os.Environ()
		resp, _ := cmd.CombinedOutput()
		t.Log(string(resp))
	})
}
