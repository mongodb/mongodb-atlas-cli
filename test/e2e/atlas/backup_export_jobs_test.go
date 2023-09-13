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
//go:build e2e || (atlas && backup && exports && jobs)

package atlas_test

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"testing"

	"github.com/mongodb/mongodb-atlas-cli/test/e2e"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	atlasv2 "go.mongodb.org/atlas-sdk/v20230201007/admin"
	atlas "go.mongodb.org/atlas/mongodbatlas"
)

func TestExportJobs(t *testing.T) {
	cliPath, err := e2e.AtlasCLIBin()
	r := require.New(t)
	r.NoError(err)

	clusterName, err := RandClusterName()
	fmt.Println(clusterName)
	r.NoError(err)

	const cloudProvider = "AWS"
	iamRoleID := os.Getenv("E2E_CLOUD_ROLE_ID")
	bucketName := os.Getenv("E2E_TEST_BUCKET")
	r.NotEmpty(iamRoleID)
	r.NotEmpty(bucketName)
	var bucketID string
	var exportJobID string
	var snapshotID string

	t.Run("Create cluster", func(t *testing.T) {
		cmd := exec.Command(cliPath,
			clustersEntity,
			"create",
			clusterName,
			"--backup",
			"--tier", tierM10,
			"--region=US_EAST_1",
			"--provider", e2eClusterProvider,
			"--mdbVersion", e2eSharedMDBVer,
			"-o=json")
		cmd.Env = os.Environ()
		resp, err := cmd.CombinedOutput()
		r.NoError(err, string(resp))

		var cluster *atlas.Cluster
		err = json.Unmarshal(resp, &cluster)
		r.NoError(err)

		ensureSharedCluster(t, cluster, clusterName, tierM10, 10, false)
	})

	t.Run("Watch create cluster", func(t *testing.T) {
		cmd := exec.Command(cliPath,
			clustersEntity,
			"watch",
			clusterName,
		)
		cmd.Env = os.Environ()
		resp, _ := cmd.CombinedOutput()
		t.Log(string(resp))
	})

	t.Run("Create bucket", func(t *testing.T) {
		cmd := exec.Command(cliPath,
			backupsEntity,
			exportsEntity,
			bucketsEntity,
			"create",
			bucketName,
			"--cloudProvider",
			cloudProvider,
			"--iamRoleId",
			iamRoleID,
			"-o=json")
		cmd.Env = os.Environ()
		resp, err := cmd.CombinedOutput()

		r.NoError(err, string(resp))

		a := assert.New(t)
		var exportBucket atlasv2.DiskBackupSnapshotAWSExportBucket
		if err = json.Unmarshal(resp, &exportBucket); a.NoError(err) {
			a.Equal(bucketName, exportBucket.GetBucketName())
		}
		bucketID = exportBucket.GetId()
	})

	t.Run("Create snapshot", func(t *testing.T) {
		cmd := exec.Command(cliPath,
			backupsEntity,
			snapshotsEntity,
			"create",
			clusterName,
			"--desc",
			"test-snapshot",
			"-o=json")
		cmd.Env = os.Environ()
		resp, err := cmd.CombinedOutput()

		r.NoError(err, string(resp))

		a := assert.New(t)
		var snapshot atlasv2.DiskBackupSnapshot
		if err = json.Unmarshal(resp, &snapshot); a.NoError(err) {
			a.Equal("test-snapshot", snapshot.GetDescription())
		}
		snapshotID = snapshot.GetId()
	})

	t.Run("Watch snapshot creation", func(t *testing.T) {
		cmd := exec.Command(cliPath,
			backupsEntity,
			snapshotsEntity,
			"watch",
			snapshotID,
			"--clusterName",
			clusterName)
		cmd.Env = os.Environ()
		resp, _ := cmd.CombinedOutput()
		t.Log(string(resp))
	})

	t.Run("Create job", func(t *testing.T) {
		cmd := exec.Command(cliPath,
			backupsEntity,
			exportsEntity,
			jobsEntity,
			"create",
			"--bucketId",
			bucketID,
			"--clusterName",
			clusterName,
			"--snapshotId",
			snapshotID,
			"-o=json")
		cmd.Env = os.Environ()
		resp, err := cmd.CombinedOutput()

		r.NoError(err, string(resp))

		a := assert.New(t)
		var job atlasv2.DiskBackupExportJob
		if err = json.Unmarshal(resp, &job); a.NoError(err) {
			a.Equal(job.GetExportBucketId(), bucketID)

			exportJobID = job.GetId()
		}
	})

	t.Run("Watch create job", func(t *testing.T) {
		cmd := exec.Command(cliPath,
			backupsEntity,
			exportsEntity,
			jobsEntity,
			"watch",
			exportJobID,
			"--clusterName",
			clusterName)
		cmd.Env = os.Environ()
		resp, _ := cmd.CombinedOutput()
		t.Log(string(resp))
	})

	t.Run("Describe job", func(t *testing.T) {
		cmd := exec.Command(cliPath,
			backupsEntity,
			exportsEntity,
			jobsEntity,
			"describe",
			"--clusterName",
			clusterName,
			"--exportId",
			exportJobID,
			"-o=json")
		cmd.Env = os.Environ()
		resp, err := cmd.CombinedOutput()

		r.NoError(err, string(resp))

		a := assert.New(t)
		var job atlasv2.DiskBackupExportJob
		if err = json.Unmarshal(resp, &job); a.NoError(err) {
			a.Equal(job.GetExportBucketId(), bucketID)
		}
	})

	t.Run("List jobs", func(t *testing.T) {
		cmd := exec.Command(cliPath,
			backupsEntity,
			exportsEntity,
			jobsEntity,
			"ls",
			clusterName,
			"-o=json")
		cmd.Env = os.Environ()
		resp, err := cmd.CombinedOutput()
		r.NoError(err, string(resp))

		var r atlasv2.PaginatedApiAtlasDiskBackupExportJob
		a := assert.New(t)
		if err = json.Unmarshal(resp, &r); a.NoError(err) {
			a.NotEmpty(r)
		}
	})

	t.Run("Delete snapshot", func(t *testing.T) {
		cmd := exec.Command(cliPath,
			backupsEntity,
			snapshotsEntity,
			"delete",
			snapshotID,
			"--clusterName",
			clusterName,
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
			clusterName)
		cmd.Env = os.Environ()
		resp, _ := cmd.CombinedOutput()
		t.Log(string(resp))
	})

	t.Run("Delete cluster", func(t *testing.T) {
		cmd := exec.Command(cliPath,
			clustersEntity,
			"delete",
			clusterName,
			"--force",
		)
		cmd.Env = os.Environ()
		resp, err := cmd.CombinedOutput()
		r.NoError(err, string(resp))

		expected := fmt.Sprintf("Deleting cluster '%s'", clusterName)
		a := assert.New(t)
		a.Equal(expected, string(resp))
	})

	t.Run("Watch delete cluster", func(t *testing.T) {
		cmd := exec.Command(cliPath,
			clustersEntity,
			"watch",
			clusterName,
		)
		cmd.Env = os.Environ()
		resp, _ := cmd.CombinedOutput()
		t.Log(string(resp))
	})
}
