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
	atlasv2 "go.mongodb.org/atlas-sdk/admin"
	"go.mongodb.org/atlas/mongodbatlas"
)

func TestSetup(t *testing.T) {
	g := newAtlasE2ETestGenerator(t)
	g.generateProject("setup")
	cliPath, err := e2e.AtlasCLIBin()
	req := require.New(t)
	req.NoError(err)

	clusterName, err := RandClusterName()
	req.NoError(err)

	dbUserUsername, err := RandUsername()
	req.NoError(err)

	tagKey := "env"
	tagValue := "e2etest"
	arbitraryAccessListIP := "21.150.105.221"

	t.Run("Run", func(t *testing.T) {
		cmd := exec.Command(cliPath,
			setupEntity,
			"--clusterName", clusterName,
			"--username", dbUserUsername,
			"--skipMongosh",
			"--skipSampleData",
			"--projectId", g.projectID,
			"--tag", tagKey+"="+tagValue,
			"--accessListIp", arbitraryAccessListIP,
			"--force")
		cmd.Env = os.Environ()
		resp, err := cmd.CombinedOutput()
		req.NoError(err, string(resp))
		assert.Contains(t, string(resp), "Cluster created.", string(resp))
	})

	t.Run("Check accessListIp was correctly added", func(t *testing.T) {
		cmd := exec.Command(cliPath,
			accessListEntity,
			"ls",
			"--projectId",
			g.projectID,
			"-o=json")
		cmd.Env = os.Environ()
		resp, err := cmd.CombinedOutput()
		req.NoError(err)

		var entries *atlasv2.PaginatedNetworkAccess
		err = json.Unmarshal(resp, &entries)
		req.NoError(err)
		req.Len(entries.Results, 1, "Expected 1 IP in list of IP's")
		req.Contains(entries.Results[0].GetIpAddress(), arbitraryAccessListIP, "IP from list does not match added IP")
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

	t.Run("Describe Cluster", func(t *testing.T) {
		cmd := exec.Command(cliPath,
			clustersEntity,
			"describe",
			clusterName,
			"-o=json",
			"--projectId", g.projectID,
		)
		cmd.Env = os.Environ()
		resp, err := cmd.CombinedOutput()
		require.NoError(t, err, string(resp))

		var cluster atlasv2.AdvancedClusterDescription
		require.NoError(t, json.Unmarshal(resp, &cluster), string(resp))
		assert.Equal(t, clusterName, *cluster.Name)

		assert.Len(t, cluster.Tags, 1)
		assert.Equal(t, tagKey, *cluster.Tags[0].Key)
		assert.Equal(t, tagValue, *cluster.Tags[0].Value)
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

		expected := fmt.Sprintf("Deleting cluster '%s'", clusterName)
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
