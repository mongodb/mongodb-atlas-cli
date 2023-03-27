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

	"github.com/mongodb/mongodb-atlas-cli/test/e2e"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	atlas "go.mongodb.org/atlas/mongodbatlas"
)

func TestServerlessBackup(t *testing.T) {
	cliPath, err := e2e.AtlasCLIBin()
	r := require.New(t)
	r.NoError(err)

	clusterName := os.Getenv("E2E_SERVERLESS_INSTANCE_NAME")
	r.NotEmpty(clusterName)

	var snapshotID, restoreJobID string

	g := newAtlasE2ETestGenerator(t)
	g.generateProjectAndCluster("serverlessBackup")

	t.Run("Snapshot List", func(t *testing.T) {
		cmd := exec.Command(cliPath,
			serverlessEntity,
			backupsEntity,
			snapshotsEntity,
			"list",
			clusterName,
			"-o=json")
		cmd.Env = os.Environ()
		resp, err := cmd.CombinedOutput()

		r.NoError(err, string(resp))

		var r atlas.CloudProviderSnapshots
		a := assert.New(t)
		if err = json.Unmarshal(resp, &r); a.NoError(err) {
			a.NotEmpty(r)

			snapshotID = r.Results[0].ID
		}
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
		resp, err := cmd.CombinedOutput()

		r.NoError(err, string(resp))

		a := assert.New(t)
		var result atlas.CloudProviderSnapshot
		if err = json.Unmarshal(resp, &result); a.NoError(err) {
			a.Equal(snapshotID, result.ID)
		}
	})

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
			g.clusterName,
			"--targetProjectId",
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
		resp, err := cmd.CombinedOutput()

		r.NoError(err, string(resp))
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
		resp, err := cmd.CombinedOutput()

		r.NoError(err, string(resp))

		a := assert.New(t)
		var result atlas.CloudProviderSnapshotRestoreJob
		if err = json.Unmarshal(resp, &result); a.NoError(err) {
			a.NotEmpty(result)
		}
	})
}
