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

//go:build e2e || (atlas && interactive)

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
	"go.mongodb.org/atlas/mongodbatlas"
)

func TestQuickstart(t *testing.T) {
	g := newAtlasE2ETestGenerator(t)
	g.generateProject("quickstart")

	cliPath, err := e2e.AtlasCLIBin()
	req := require.New(t)
	req.NoError(err)

	clusterName, err := RandClusterName()
	req.NoError(err)

	dbUserUsername, err := RandUsername()
	req.NoError(err)

	t.Run("Run", func(t *testing.T) {
		cmd := exec.Command(cliPath,
			"quickstart",
			"--clusterName", clusterName,
			"--username", dbUserUsername,
			"--skipMongosh",
			"--skipSampleData",
			"--projectId", g.projectID,
			"--force")
		cmd.Env = os.Environ()
		resp, err := cmd.CombinedOutput()
		req.NoError(err, string(resp))
		assert.Contains(t, string(resp), "Cluster created.", string(resp))
	})

	t.Run("Watch Cluster", func(t *testing.T) {
		cmd := exec.Command(cliPath,
			clustersEntity,
			"watch",
			clusterName,
			"--projectId", g.projectID,
		)
		cmd.Env = os.Environ()
		resp, err := cmd.CombinedOutput()
		req.NoError(err, string(resp))
		assert.Contains(t, string(resp), "Cluster available")
	})

	t.Run("Describe DB User", func(t *testing.T) {
		cmd := exec.Command(cliPath,
			dbusersEntity,
			"describe",
			dbUserUsername,
			"-o=json",
			"--projectId", g.projectID,
		)
		cmd.Env = os.Environ()
		resp, err := cmd.CombinedOutput()
		require.NoError(t, err, string(resp))

		var user mongodbatlas.DatabaseUser
		require.NoError(t, json.Unmarshal(resp, &user), string(resp))
		assert.Equal(t, dbUserUsername, user.Username)
	})

	t.Run("Delete Cluster", func(t *testing.T) {
		cmd := exec.Command(cliPath,
			clustersEntity,
			"delete",
			clusterName,
			"--force",
			"--projectId", g.projectID,
		)
		cmd.Env = os.Environ()
		resp, err := cmd.CombinedOutput()
		req.NoError(err, string(resp))

		expected := fmt.Sprintf("Cluster '%s' deleted\n", clusterName)
		assert.Equal(t, expected, string(resp))
	})

	t.Run("Watch cluster deletion", func(t *testing.T) {
		cmd := exec.Command(cliPath,
			clustersEntity,
			"watch",
			clusterName,
			"--projectId", g.projectID,
		)
		cmd.Env = os.Environ()
		// this command will fail with 404 once the cluster is deleted
		// we just need to wait for this to close the project
		resp, _ := cmd.CombinedOutput()
		t.Log(string(resp))
	})
}
