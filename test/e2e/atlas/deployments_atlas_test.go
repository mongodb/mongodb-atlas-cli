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
//go:build e2e || (atlas && deployments && atlasclusters)

package atlas_test

import (
	"bytes"
	"context"
	"fmt"
	"os"
	"os/exec"
	"strings"
	"testing"

	"github.com/mongodb/mongodb-atlas-cli/atlascli/test/e2e"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	collectionNameAtlas = "movies"
	databaseNameAtlas   = "sample_mflix"
)

func TestDeploymentsAtlas(t *testing.T) {
	g := newAtlasE2ETestGenerator(t)
	g.generateProject("setup")
	cliPath, err := e2e.AtlasCLIBin()
	req := require.New(t)
	req.NoError(err)

	clusterName, err := RandClusterName()
	req.NoError(err)

	dbUserUsername, err := RandUsername()
	req.NoError(err)

	dbUserPassword := dbUserUsername + "~PwD"

	var client *mongo.Client
	ctx := context.Background()

	t.Run("Setup", func(t *testing.T) {
		cmd := exec.Command(cliPath,
			deploymentEntity,
			"setup",
			clusterName,
			"--type",
			"atlas",
			"--tier",
			"M10",
			"--force",
			"--skipMongosh",
			"--debug",
			"--projectId", g.projectID,
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
	require.NoError(t, watchCluster(g.projectID, clusterName))

	t.Run("Connect to database", func(t *testing.T) {
		cmd := exec.Command(cliPath,
			deploymentEntity,
			"connect",
			clusterName,
			"--type", "atlas",
			"--connectWith", "connectionString",
			"--projectId", g.projectID,
		)

		cmd.Env = os.Environ()

		r, err := e2e.RunAndGetStdOut(cmd)
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

	t.Run("Pause Cluster", func(t *testing.T) {
		cmd := exec.Command(cliPath,
			deploymentEntity,
			"pause",
			clusterName,
			"--type=ATLAS",
			"--projectId", g.projectID,
		)
		cmd.Env = os.Environ()
		resp, err := cmd.CombinedOutput()
		require.NoError(t, err, string(resp))
		assert.Contains(t, string(resp), fmt.Sprintf("Pausing deployment '%s'", clusterName))
	})

	t.Run("Start Cluster", func(t *testing.T) {
		cmd := exec.Command(cliPath,
			deploymentEntity,
			"start",
			clusterName,
			"--type=ATLAS",
			"--projectId", g.projectID,
		)
		cmd.Env = os.Environ()
		resp, err := e2e.RunAndGetStdOut(cmd)
		require.NoError(t, err, string(resp))
		assert.Contains(t, string(resp), fmt.Sprintf("Starting deployment '%s'", clusterName))
	})
	require.NoError(t, watchCluster(g.projectID, clusterName))

	t.Run("Create Search Index", func(t *testing.T) {
		cmd := exec.Command(cliPath,
			deploymentEntity,
			searchEntity,
			indexEntity,
			"create",
			"testIndex",
			"--type",
			"atlas",
			"--projectId", g.projectID,
			"--deploymentName", clusterName,
			"--db",
			databaseNameAtlas,
			"--collection",
			collectionNameAtlas,
			"--watch",
		)
		cmd.Env = os.Environ()

		r, err := e2e.RunAndGetStdOut(cmd)
		out := string(r)
		require.NoError(t, err, out)
		assert.Contains(t, out, "Search index created")
	})

	t.Run("Delete Cluster", func(t *testing.T) {
		cmd := exec.Command(cliPath,
			deploymentEntity,
			"delete",
			clusterName,
			"--type",
			"ATLAS",
			"--force",
			"--watch",
			"--watchTimeout", "300",
			"--projectId", g.projectID,
		)
		cmd.Env = os.Environ()
		resp, err := e2e.RunAndGetStdOut(cmd)
		req.NoError(err, string(resp))

		expected := "Deployment '" + clusterName + "' deleted\n<nil>\n"
		assert.Equal(t, expected, string(resp))
	})
}
