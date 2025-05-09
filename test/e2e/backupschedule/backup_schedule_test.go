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
//go:build e2e || e2eSnap || (atlas && backup && schedule)

package backupschedule

import (
	"encoding/json"
	"os"
	"os/exec"
	"testing"

	"github.com/mongodb/mongodb-atlas-cli/atlascli/test/internal"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	atlasClustersPinned "go.mongodb.org/atlas-sdk/v20240530005/admin"
)

const (
	backupsEntity = "backups"
)

func TestSchedule(t *testing.T) {
	g := internal.NewAtlasE2ETestGenerator(t, internal.WithSnapshot(), internal.WithBackup())
	cliPath, err := internal.AtlasCLIBin()
	r := require.New(t)
	r.NoError(err)

	g.GenerateProjectAndCluster("backupSchedule")

	var policy *atlasClustersPinned.DiskBackupSnapshotSchedule

	g.Run("Describe", func(t *testing.T) { //nolint:thelper // g.Run replaces t.Run
		cmd := exec.Command(cliPath,
			backupsEntity,
			"schedule",
			"describe",
			g.ClusterName,
			"--projectId",
			g.ProjectID,
			"-o=json",
		)
		cmd.Env = os.Environ()
		resp, err := internal.RunAndGetStdOut(cmd)
		require.NoError(t, err)
		require.NoError(t, json.Unmarshal(resp, &policy))

		assert.Equal(t, g.ClusterName, policy.GetClusterName())
	})

	g.Run("Update", func(t *testing.T) { //nolint:thelper // g.Run replaces t.Run
		cmd := exec.Command(cliPath,
			backupsEntity,
			"schedule",
			"update",
			"--clusterName",
			g.ClusterName,
			"--useOrgAndGroupNamesInExportPrefix",
			"--projectId",
			g.ProjectID,
			"-o=json",
		)
		cmd.Env = os.Environ()
		resp, err := internal.RunAndGetStdOut(cmd)
		require.NoError(t, err, string(resp))
	})

	g.Run("Delete", func(t *testing.T) { //nolint:thelper // g.Run replaces t.Run
		cmd := exec.Command(cliPath,
			backupsEntity,
			"schedule",
			"delete",
			g.ClusterName,
			"--projectId",
			g.ProjectID,
			"--force",
		)
		cmd.Env = os.Environ()
		resp, err := internal.RunAndGetStdOut(cmd)
		require.NoError(t, err, string(resp))
	})
}
