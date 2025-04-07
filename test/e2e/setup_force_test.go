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

package e2e

import (
	"encoding/json"
	"os"
	"os/exec"
	"testing"

	"github.com/mongodb/mongodb-atlas-cli/atlascli/test/internal"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	atlasClustersPinned "go.mongodb.org/atlas-sdk/v20240530005/admin"
	atlasv2 "go.mongodb.org/atlas-sdk/v20250312001/admin"
)

func TestSetup(t *testing.T) {
	g := internal.NewAtlasE2ETestGenerator(t, internal.WithSnapshot(), internal.WithSnapshotSkip(internal.SkipSimilarSnapshots))
	g.GenerateProject("setup")
	cliPath, err := internal.AtlasCLIBin()
	req := require.New(t)
	req.NoError(err)

	clusterName := g.Memory("clusterName", internal.Must(internal.RandClusterName())).(string)

	dbUserUsername := g.Memory("dbUserUsername", internal.Must(internal.RandClusterName())).(string)

	tagKey := "env"
	tagValue := "e2etest"
	arbitraryAccessListIP := "21.150.105.221"

	g.Run("Run", func(t *testing.T) { //nolint:thelper // g.Run replaces t.Run
		cmd := exec.Command(cliPath,
			setupEntity,
			"--clusterName", clusterName,
			"--username", dbUserUsername,
			"--skipMongosh",
			"--skipSampleData",
			"--projectId", g.ProjectID,
			"--tag", tagKey+"="+tagValue,
			"--accessListIp", arbitraryAccessListIP,
			"--force")
		cmd.Env = os.Environ()
		resp, err := internal.RunAndGetStdOut(cmd)
		require.NoError(t, err, string(resp))
		assert.Contains(t, string(resp), "Cluster created.", string(resp))
	})
	t.Cleanup(func() {
		require.NoError(t, internal.DeleteClusterForProject(g.ProjectID, clusterName))
	})
	g.Run("Check accessListIp was correctly added", func(t *testing.T) { //nolint:thelper // g.Run replaces t.Run
		cmd := exec.Command(cliPath,
			accessListEntity,
			"ls",
			"--projectId",
			g.ProjectID,
			"-o=json")
		cmd.Env = os.Environ()
		resp, err := internal.RunAndGetStdOut(cmd)
		require.NoError(t, err, string(resp))

		var entries *atlasv2.PaginatedNetworkAccess
		require.NoError(t, json.Unmarshal(resp, &entries))

		assert.Len(t, entries.GetResults(), 1, "Expected 1 IP in list of IP's")
		assert.Contains(t, entries.GetResults()[0].GetIpAddress(), arbitraryAccessListIP, "IP from list does not match added IP")
	})

	require.NoError(t, internal.WatchCluster(g.ProjectID, clusterName))

	g.Run("Describe DB User", func(t *testing.T) { //nolint:thelper // g.Run replaces t.Run
		cmd := exec.Command(cliPath,
			dbusersEntity,
			"describe",
			dbUserUsername,
			"-o=json",
			"--projectId", g.ProjectID,
		)
		cmd.Env = os.Environ()
		resp, err := internal.RunAndGetStdOut(cmd)
		require.NoError(t, err, string(resp))

		var user atlasv2.CloudDatabaseUser
		require.NoError(t, json.Unmarshal(resp, &user), string(resp))
		assert.Equal(t, dbUserUsername, user.GetUsername())
	})

	g.Run("Describe Cluster", func(t *testing.T) { //nolint:thelper // g.Run replaces t.Run
		cmd := exec.Command(cliPath,
			clustersEntity,
			"describe",
			clusterName,
			"-o=json",
			"--projectId", g.ProjectID,
		)
		cmd.Env = os.Environ()
		resp, err := internal.RunAndGetStdOut(cmd)
		require.NoError(t, err, string(resp))

		var cluster atlasClustersPinned.AdvancedClusterDescription
		require.NoError(t, json.Unmarshal(resp, &cluster), string(resp))
		assert.Equal(t, clusterName, *cluster.Name)

		assert.Len(t, cluster.GetTags(), 1)
		assert.Equal(t, tagKey, cluster.GetTags()[0].Key)
		assert.Equal(t, tagValue, cluster.GetTags()[0].Value)
	})
}
