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

	"github.com/andreangiolillo/mongocli-test/test/e2e"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	collectionNameAtlas = "myCol"
	databaseNameAtlas   = "myDB"
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

	var connectionString string
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
			"--skipSampleData",
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
		req.NoError(err, e.String())

		connectionString = strings.TrimSpace(o.String())
		connectionString = strings.Replace(connectionString, "Your connection string: ", "", 1)
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

		r, err := cmd.CombinedOutput()
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
		req.NoError(err)
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
		req.NoError(err, string(resp))
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
		resp, err := cmd.CombinedOutput()
		req.NoError(err, string(resp))
		a := assert.New(t)
		a.Contains(string(resp), fmt.Sprintf("Starting deployment '%s'", clusterName))
	})
	require.NoError(t, watchCluster(g.projectID, clusterName))

	t.Run("Create Index", func(t *testing.T) {
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
		)
		cmd.Env = os.Environ()

		r, err := cmd.CombinedOutput()
		out := string(r)
		req.NoError(err, out)
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
		resp, err := cmd.CombinedOutput()
		req.NoError(err, string(resp))

		expected := fmt.Sprintf("Deployment '" + clusterName + "' deleted\n<nil>\n")
		assert.Equal(t, expected, string(resp))
	})
}
