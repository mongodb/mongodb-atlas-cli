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
//go:build e2e || (atlas && backup && schedule)

package atlas_test

import (
	"encoding/json"
	"os"
	"os/exec"
	"testing"

	"github.com/mongodb/mongodb-atlas-cli/atlascli/test/e2e"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	atlasv2 "go.mongodb.org/atlas-sdk/v20240530002/admin"
)

func TestSchedule(t *testing.T) {
	cliPath, err := e2e.AtlasCLIBin()
	r := require.New(t)
	r.NoError(err)

	g := newAtlasE2ETestGenerator(t)
	g.enableBackup = true
	g.generateProjectAndCluster("backupSchedule")

	var policy *atlasv2.DiskBackupSnapshotSchedule

	t.Run("Describe", func(t *testing.T) {
		cmd := exec.Command(cliPath,
			backupsEntity,
			"schedule",
			"describe",
			g.clusterName,
			"--projectId",
			g.projectID,
			"-o=json",
		)
		cmd.Env = os.Environ()
		resp, err := e2e.RunAndGetStdOut(cmd)
		require.NoError(t, err)
		require.NoError(t, json.Unmarshal(resp, &policy))

		assert.Equal(t, g.clusterName, policy.GetClusterName())
	})

	t.Run("Update", func(t *testing.T) {
		cmd := exec.Command(cliPath,
			backupsEntity,
			"schedule",
			"update",
			"--clusterName",
			g.clusterName,
			"--useOrgAndGroupNamesInExportPrefix",
			"--projectId",
			g.projectID,
			"-o=json",
		)
		cmd.Env = os.Environ()
		resp, err := e2e.RunAndGetStdOut(cmd)
		require.NoError(t, err, string(resp))
	})

	t.Run("Delete", func(t *testing.T) {
		cmd := exec.Command(cliPath,
			backupsEntity,
			"schedule",
			"delete",
			g.clusterName,
			"--projectId",
			g.projectID,
			"--force",
		)
		cmd.Env = os.Environ()
		resp, err := e2e.RunAndGetStdOut(cmd)
		require.NoError(t, err, string(resp))
	})
}
