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
//go:build e2e || (atlas && backup && exports && buckets)

package atlas_test

import (
	"encoding/json"
	"os"
	"os/exec"
	"testing"

	"github.com/mongodb/mongodb-atlas-cli/test/e2e"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	atlasv2 "go.mongodb.org/atlas-sdk/v20231001002/admin"
)

func TestExportBuckets(t *testing.T) {
	cliPath, err := e2e.AtlasCLIBin()
	r := require.New(t)
	r.NoError(err)

	const cloudProvider = "AWS"
	iamRoleID := os.Getenv("E2E_CLOUD_ROLE_ID")
	bucketName := os.Getenv("E2E_TEST_BUCKET")
	r.NotEmpty(iamRoleID)
	r.NotEmpty(bucketName)
	var bucketID string

	t.Run("Create", func(t *testing.T) {
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
		r.NoError(json.Unmarshal(resp, &exportBucket))
		a.Equal(bucketName, exportBucket.GetBucketName())
		bucketID = exportBucket.GetId()
	})

	t.Run("List", func(t *testing.T) {
		cmd := exec.Command(cliPath,
			backupsEntity,
			exportsEntity,
			bucketsEntity,
			"list",
			"-o=json")
		cmd.Env = os.Environ()
		resp, err := cmd.CombinedOutput()
		r.NoError(err, string(resp))
		var buckets atlasv2.PaginatedBackupSnapshotExportBucket
		r.NoError(json.Unmarshal(resp, &buckets))
		assert.NotEmpty(t, buckets)
	})

	t.Run("Describe", func(t *testing.T) {
		cmd := exec.Command(cliPath,
			backupsEntity,
			exportsEntity,
			bucketsEntity,
			"describe",
			"--bucketId",
			bucketID,
			"-o=json")
		cmd.Env = os.Environ()
		resp, err := cmd.CombinedOutput()
		r.NoError(err, string(resp))
		var exportBucket atlasv2.DiskBackupSnapshotAWSExportBucket
		r.NoError(json.Unmarshal(resp, &exportBucket))
		assert.Equal(t, bucketName, exportBucket.GetBucketName())
	})

	t.Run("Delete", func(t *testing.T) {
		cmd := exec.Command(cliPath,
			backupsEntity,
			exportsEntity,
			bucketsEntity,
			"delete",
			"--bucketId",
			bucketID,
			"--force")
		cmd.Env = os.Environ()
		resp, err := cmd.CombinedOutput()

		r.NoError(err, string(resp))
	})
}
