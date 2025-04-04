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

package e2e

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"testing"

	"github.com/mongodb/mongodb-atlas-cli/atlascli/test/internal"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	atlasClustersPinned "go.mongodb.org/atlas-sdk/v20240530005/admin"

	atlasv2 "go.mongodb.org/atlas-sdk/v20250312001/admin"
)

func TestExportJobs(t *testing.T) {
	g := internal.NewAtlasE2ETestGenerator(t, internal.WithSnapshot())
	cliPath, err := internal.AtlasCLIBin()
	r := require.New(t)
	r.NoError(err)

	clusterName := g.Memory("clusterName", internal.Must(internal.RandClusterName())).(string)
	fmt.Println(clusterName)

	mdbVersion, err := internal.MongoDBMajorVersion()
	r.NoError(err)

	const cloudProvider = "AWS"
	iamRoleID := os.Getenv("E2E_CLOUD_ROLE_ID")
	bucketName := os.Getenv("E2E_TEST_BUCKET")
	r.NotEmpty(iamRoleID)
	r.NotEmpty(bucketName)
	var bucketID string
	var exportJobID string
	var snapshotID string

	g.Run("Create cluster", func(t *testing.T) { //nolint:thelper // g.Run replaces t.Run
		cmd := exec.Command(cliPath,
			clustersEntity,
			"create",
			clusterName,
			"--backup",
			"--tier", tierM10,
			"--region=US_EAST_1",
			"--provider", e2eClusterProvider,
			"--mdbVersion", mdbVersion,
			"-o=json")
		cmd.Env = os.Environ()
		resp, err := internal.RunAndGetStdOut(cmd)
		r.NoError(err, string(resp))

		var cluster *atlasClustersPinned.AdvancedClusterDescription
		r.NoError(json.Unmarshal(resp, &cluster))
		internal.EnsureCluster(t, cluster, clusterName, mdbVersion, 10, false)
	})
	t.Cleanup(func() {
		require.NoError(t, internal.DeleteClusterForProject("", clusterName))
	})
	require.NoError(t, internal.WatchCluster("", clusterName))

	g.Run("Create bucket", func(t *testing.T) { //nolint:thelper // g.Run replaces t.Run
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
		resp, err := internal.RunAndGetStdOut(cmd)
		r.NoError(err, string(resp))
		var exportBucket atlasv2.DiskBackupSnapshotExportBucketResponse
		r.NoError(json.Unmarshal(resp, &exportBucket))
		assert.Equal(t, bucketName, exportBucket.GetBucketName())
		bucketID = exportBucket.GetId()
	})

	g.Run("Create snapshot", func(t *testing.T) { //nolint:thelper // g.Run replaces t.Run
		cmd := exec.Command(cliPath,
			backupsEntity,
			snapshotsEntity,
			"create",
			clusterName,
			"--desc",
			"test-snapshot",
			"-o=json")
		cmd.Env = os.Environ()
		resp, err := internal.RunAndGetStdOut(cmd)

		r.NoError(err, string(resp))

		a := assert.New(t)
		var snapshot atlasv2.DiskBackupSnapshot
		r.NoError(json.Unmarshal(resp, &snapshot))
		a.Equal("test-snapshot", snapshot.GetDescription())
		snapshotID = snapshot.GetId()
	})

	g.Run("Watch snapshot creation", func(t *testing.T) { //nolint:thelper // g.Run replaces t.Run
		cmd := exec.Command(cliPath,
			backupsEntity,
			snapshotsEntity,
			"watch",
			snapshotID,
			"--clusterName",
			clusterName)
		cmd.Env = os.Environ()
		resp, _ := internal.RunAndGetStdOut(cmd)
		t.Log(string(resp))
	})

	g.Run("Create job", func(t *testing.T) { //nolint:thelper // g.Run replaces t.Run
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
		resp, err := internal.RunAndGetStdOut(cmd)

		r.NoError(err, string(resp))
		var job atlasv2.DiskBackupExportJob
		r.NoError(json.Unmarshal(resp, &job))
		assert.Equal(t, job.GetExportBucketId(), bucketID)
		exportJobID = job.GetId()
	})

	g.Run("Watch create job", func(t *testing.T) { //nolint:thelper // g.Run replaces t.Run
		cmd := exec.Command(cliPath,
			backupsEntity,
			exportsEntity,
			jobsEntity,
			"watch",
			exportJobID,
			"--clusterName",
			clusterName)
		cmd.Env = os.Environ()
		resp, _ := internal.RunAndGetStdOut(cmd)
		t.Log(string(resp))
	})

	g.Run("Describe job", func(t *testing.T) { //nolint:thelper // g.Run replaces t.Run
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
		resp, err := internal.RunAndGetStdOut(cmd)

		r.NoError(err, string(resp))

		var job atlasv2.DiskBackupExportJob
		r.NoError(json.Unmarshal(resp, &job))
		assert.Equal(t, job.GetExportBucketId(), bucketID)
	})

	g.Run("List jobs", func(t *testing.T) { //nolint:thelper // g.Run replaces t.Run
		cmd := exec.Command(cliPath,
			backupsEntity,
			exportsEntity,
			jobsEntity,
			"ls",
			clusterName,
			"-o=json")
		cmd.Env = os.Environ()
		resp, err := internal.RunAndGetStdOut(cmd)
		r.NoError(err, string(resp))

		var jobs atlasv2.PaginatedApiAtlasDiskBackupExportJob
		r.NoError(json.Unmarshal(resp, &jobs))
		assert.NotEmpty(t, jobs)
	})

	g.Run("Delete snapshot", func(t *testing.T) { //nolint:thelper // g.Run replaces t.Run
		cmd := exec.Command(cliPath,
			backupsEntity,
			snapshotsEntity,
			"delete",
			snapshotID,
			"--clusterName",
			clusterName,
			"--force")
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
			clusterName)
		cmd.Env = os.Environ()
		resp, _ := internal.RunAndGetStdOut(cmd)
		t.Log(string(resp))
	})
}
