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

package deploymentsatlas

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"strings"
	"testing"

	"github.com/mongodb/mongodb-atlas-cli/atlascli/test/internal"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func TestDeploymentsAtlasISS(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping test in short mode")
	}

	if internal.TestRunMode() != internal.TestModeLive {
		t.Skip("skipping test in snapshot mode")
	}

	g := internal.NewAtlasE2ETestGenerator(t, internal.WithSnapshot())
	g.GenerateProject("setup")
	cliPath, err := internal.AtlasCLIBin()
	req := require.New(t)
	req.NoError(err)

	clusterName := g.Memory("clusterName", internal.Must(internal.RandClusterName())).(string)
	dbUserUsername := g.Memory("dbUserUsername", internal.Must(internal.RandUsername())).(string)

	dbUserPassword := dbUserUsername + "~PwD"

	var client *mongo.Client
	ctx := t.Context()

	g.Run("Setup", func(t *testing.T) { //nolint:thelper // g.Run replaces t.Run
		cmd := exec.Command(cliPath,
			deploymentEntity,
			"setup",
			clusterName,
			"--type",
			"atlas",
			"--tier",
			"M10",
			"--mdbVersion",
			"8.0",
			"--skipMongosh",
			"--force",
			"--debug",
			"--autoScalingMode",
			"independentShardScaling",
			"--projectId", g.ProjectID,
			"--username", dbUserUsername,
			"--password", dbUserPassword,
		)

		cmd.Env = os.Environ()

		var o, e bytes.Buffer
		cmd.Stdout = &o
		cmd.Stderr = &e
		err = cmd.Run()
		require.NoError(t, err, e.String())
	})
	require.NoError(t, internal.WatchCluster(g.ProjectID, clusterName))

	g.Run("Connect to database", func(t *testing.T) { //nolint:thelper // g.Run replaces t.Run
		cmd := exec.Command(cliPath,
			deploymentEntity,
			"connect",
			clusterName,
			"--type", "atlas",
			"--connectWith", "connectionString",
			"--projectId", g.ProjectID,
		)

		cmd.Env = os.Environ()

		r, err := internal.RunAndGetStdOut(cmd)
		req.NoError(err, string(r))

		connectionString := strings.TrimSpace(string(r))
		client, err = mongo.Connect(
			ctx,
			options.Client().
				ApplyURI(connectionString).
				SetAuth(options.Credential{
					AuthMechanism: "PLAIN",
					Username:      dbUserUsername,
					Password:      dbUserPassword,
				}),
		)
		require.NoError(t, err)
	})

	t.Cleanup(func() {
		require.NoError(t, client.Disconnect(ctx))
	})

	g.Run("Pause Cluster", func(t *testing.T) { //nolint:thelper // g.Run replaces t.Run
		cmd := exec.Command(cliPath,
			deploymentEntity,
			"pause",
			clusterName,
			"--type=ATLAS",
			"--projectId", g.ProjectID,
		)
		cmd.Env = os.Environ()
		resp, err := cmd.CombinedOutput()
		require.NoError(t, err, string(resp))
		assert.Contains(t, string(resp), fmt.Sprintf("Pausing deployment '%s'", clusterName))
	})

	g.Run("Start Cluster", func(t *testing.T) { //nolint:thelper // g.Run replaces t.Run
		cmd := exec.Command(cliPath,
			deploymentEntity,
			"start",
			clusterName,
			"--type=ATLAS",
			"--projectId", g.ProjectID,
		)
		cmd.Env = os.Environ()
		resp, err := internal.RunAndGetStdOut(cmd)
		require.NoError(t, err, string(resp))
		assert.Contains(t, string(resp), fmt.Sprintf("Starting deployment '%s'", clusterName))
	})
	require.NoError(t, internal.WatchCluster(g.ProjectID, clusterName))

	g.Run("Create Search Index", func(t *testing.T) { //nolint:thelper // g.Run replaces t.Run
		cmd := exec.Command(cliPath,
			deploymentEntity,
			searchEntity,
			indexEntity,
			"create",
			"testIndex",
			"--type",
			"atlas",
			"--projectId", g.ProjectID,
			"--deploymentName", clusterName,
			"--db",
			databaseNameAtlas,
			"--collection",
			collectionNameAtlas,
			"--watch",
		)
		cmd.Env = os.Environ()

		r, err := internal.RunAndGetStdOut(cmd)
		out := string(r)
		require.NoError(t, err, out)
		assert.Contains(t, out, "Search index created")
	})

	g.Run("Delete Cluster", func(t *testing.T) { //nolint:thelper // g.Run replaces t.Run
		cmd := exec.Command(cliPath,
			deploymentEntity,
			"delete",
			clusterName,
			"--type",
			"ATLAS",
			"--force",
			"--watch",
			"--watchTimeout", "300",
			"--projectId", g.ProjectID,
		)
		cmd.Env = os.Environ()
		resp, err := internal.RunAndGetStdOut(cmd)
		req.NoError(err, string(resp))

		expected := "Deployment '" + clusterName + "' deleted\n<nil>\n"
		assert.Equal(t, expected, string(resp))
	})
}
