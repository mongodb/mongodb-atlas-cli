// Copyright 2022 MongoDB Inc
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
//go:build e2e || e2eSnap || (atlas && backup && exports && buckets)

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

func TestExportBuckets(t *testing.T) {
	g := newAtlasE2ETestGenerator(t, withSnapshot())
	cliPath, err := AtlasCLIBin()
	r := require.New(t)
	r.NoError(err)

	const cloudProvider = "AWS"
	iamRoleID := os.Getenv("E2E_CLOUD_ROLE_ID")
	bucketName := os.Getenv("E2E_TEST_BUCKET")
	r.NotEmpty(iamRoleID)
	r.NotEmpty(bucketName)
	var bucketID string

	g.Run("Create", func(t *testing.T) { //nolint:thelper // g.Run replaces t.Run
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
		cmd.Env = append(os.Environ(), "GOCOVERDIR="+os.Getenv("BINGOCOVERDIR"))
		resp, err := RunAndGetStdOut(cmd)

		r.NoError(err, string(resp))

		var exportBucket atlasv2.DiskBackupSnapshotExportBucketResponse
		r.NoError(json.Unmarshal(resp, &exportBucket))
		assert.Equal(t, bucketName, exportBucket.GetBucketName())
		bucketID = exportBucket.GetId()
	})

	g.Run("List", func(t *testing.T) { //nolint:thelper // g.Run replaces t.Run
		cmd := exec.Command(cliPath,
			backupsEntity,
			exportsEntity,
			bucketsEntity,
			"list",
			"-o=json")
		cmd.Env = append(os.Environ(), "GOCOVERDIR="+os.Getenv("BINGOCOVERDIR"))
		resp, err := RunAndGetStdOut(cmd)
		r.NoError(err, string(resp))
		var buckets atlasv2.PaginatedBackupSnapshotExportBuckets
		r.NoError(json.Unmarshal(resp, &buckets))
		assert.NotEmpty(t, buckets)
	})

	g.Run("Describe", func(t *testing.T) { //nolint:thelper // g.Run replaces t.Run
		cmd := exec.Command(cliPath,
			backupsEntity,
			exportsEntity,
			bucketsEntity,
			"describe",
			"--bucketId",
			bucketID,
			"-o=json")
		cmd.Env = append(os.Environ(), "GOCOVERDIR="+os.Getenv("BINGOCOVERDIR"))
		resp, err := RunAndGetStdOut(cmd)
		r.NoError(err, string(resp))
		var exportBucket atlasv2.DiskBackupSnapshotExportBucketResponse
		r.NoError(json.Unmarshal(resp, &exportBucket))
		assert.Equal(t, bucketName, exportBucket.GetBucketName())
	})

	g.Run("Delete", func(t *testing.T) { //nolint:thelper // g.Run replaces t.Run
		cmd := exec.Command(cliPath,
			backupsEntity,
			exportsEntity,
			bucketsEntity,
			"delete",
			"--bucketId",
			bucketID,
			"--force")
		cmd.Env = append(os.Environ(), "GOCOVERDIR="+os.Getenv("BINGOCOVERDIR"))
		resp, err := RunAndGetStdOut(cmd)

		require.NoError(t, err, string(resp))
	})
}
