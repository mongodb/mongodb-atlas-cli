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
//go:build e2e || (atlas && deployments && local && auth)

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
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func TestDeploymentsLocalWithAuth(t *testing.T) {
	const (
		deploymentName = "test-auth"
		dbUsername     = "admin"
		dbUserPassword = "testpwd"
	)

	cliPath, err := e2e.AtlasCLIBin()
	req := require.New(t)
	req.NoError(err)

	t.Run("Setup", func(t *testing.T) {
		t.Cleanup(func() {
			cmd := exec.Command(cliPath,
				deploymentEntity,
				"diagnostics",
				deploymentName,
			)

			cmd.Env = os.Environ()

			r, errDiag := e2e.RunAndGetStdOut(cmd)
			t.Log("Diagnostics")
			t.Log(errDiag, string(r))
		})

		cmd := exec.Command(cliPath,
			deploymentEntity,
			"setup",
			deploymentName,
			"--type",
			"local",
			"--username",
			dbUsername,
			"--password",
			dbUserPassword,
			"--bindIpAll",
			"--force",
		)

		cmd.Env = os.Environ()

		r, setupErr := e2e.RunAndGetStdOut(cmd)
		require.NoError(t, setupErr, string(r))
	})

	t.Cleanup(func() {
		cmd := exec.Command(cliPath,
			deploymentEntity,
			"delete",
			deploymentName,
			"--type",
			"local",
			"--force",
		)

		cmd.Env = os.Environ()

		r, delErr := e2e.RunAndGetStdOut(cmd)
		require.NoError(t, delErr, string(r))
	})

	t.Run("List deployments", func(t *testing.T) {
		cmd := exec.Command(cliPath,
			deploymentEntity,
			"list",
			"--type",
			"local",
		)

		cmd.Env = os.Environ()

		o, e, err := splitOutput(cmd)
		require.NoError(t, err, e)

		outputLines := strings.Split(o, "\n")
		assert.Regexp(t, `NAME\s+TYPE\s+MDB VER\s+STATE`, outputLines[0])

		cols := strings.Fields(outputLines[1])
		assert.Equal(t, deploymentName, cols[0])
		assert.Equal(t, "LOCAL", cols[1])
		assert.Contains(t, cols[2], "7.0.")
		assert.Equal(t, "IDLE", cols[3])
	})

	ctx := context.Background()
	var client *mongo.Client
	var myDB *mongo.Database
	var myCol *mongo.Collection

	t.Run("Connect to database", func(t *testing.T) {
		cmd := exec.Command(cliPath,
			deploymentEntity,
			"connect",
			deploymentName,
			"--type",
			"local",
			"--username",
			dbUsername,
			"--password",
			dbUserPassword,
			"--connectWith",
			"connectionString",
		)

		cmd.Env = os.Environ()

		r, err := e2e.RunAndGetStdOut(cmd)
		require.NoError(t, err, string(r))

		connectionString := strings.TrimSpace(string(r))
		client, err = mongo.Connect(ctx, options.Client().ApplyURI(connectionString))
		req.NoError(err)
		myDB = client.Database(databaseName)
		myCol = myDB.Collection(collectionName)
	})

	t.Cleanup(func() {
		require.NoError(t, client.Disconnect(ctx))
	})

	t.Run("Seed database", func(t *testing.T) {
		ids, err := myCol.InsertMany(ctx, []any{
			bson.M{
				"name": "test1",
			}, bson.M{
				"name": "test2",
			},
		})
		req.NoError(err)
		t.Log(ids)
	})

	t.Run("Create Search Index", func(t *testing.T) {
		cmd := exec.Command(cliPath,
			deploymentEntity,
			searchEntity,
			indexEntity,
			"create",
			searchIndexName,
			"--type",
			"local",
			"--deploymentName",
			deploymentName,
			"--db",
			databaseName,
			"--collection",
			collectionName,
			"--username",
			dbUsername,
			"--password",
			dbUserPassword,
			"--type",
			"LOCAL",
			"-w",
		)

		cmd.Env = os.Environ()

		r, err := e2e.RunAndGetStdOut(cmd)
		out := string(r)
		req.NoError(err, out)
		assert.Contains(t, out, "Search index created with ID:")
	})

	var indexID string

	t.Run("Index List", func(t *testing.T) {
		cmd := exec.Command(cliPath,
			deploymentEntity,
			searchEntity,
			indexEntity,
			"ls",
			"--deploymentName",
			deploymentName,
			"--db",
			databaseName,
			"--collection",
			collectionName,
			"--type",
			"LOCAL",
			"--username",
			dbUsername,
			"--password",
			dbUserPassword,
		)

		cmd.Env = os.Environ()
		o, e, err := splitOutput(cmd)
		req.NoError(err, e)
		assert.Contains(t, o, searchIndexName)

		lines := strings.Split(o, "\n")
		cols := strings.Fields(lines[1])
		indexID = cols[0]
	})

	t.Run("Describe search index", func(t *testing.T) {
		cmd := exec.Command(cliPath,
			deploymentEntity,
			searchEntity,
			indexEntity,
			"describe",
			indexID,
			"--deploymentName",
			deploymentName,
			"--type",
			"LOCAL",
			"--username",
			dbUsername,
			"--password",
			dbUserPassword,
		)

		cmd.Env = os.Environ()

		r, err := e2e.RunAndGetStdOut(cmd)
		require.NoError(t, err, string(r))
	})

	t.Run("Test Search Index", func(t *testing.T) {
		c, err := myCol.Aggregate(ctx, bson.A{
			bson.M{
				"$search": bson.M{
					"index": searchIndexName,
					"text": bson.M{
						"query": "test1",
						"path":  "name",
					},
				},
			},
		})
		require.NoError(t, err)
		var results []bson.M
		require.NoError(t, c.All(ctx, &results))
		assert.Len(t, results, 1)
	})

	t.Run("Delete Index", func(t *testing.T) {
		cmd := exec.Command(cliPath,
			deploymentEntity,
			searchEntity,
			indexEntity,
			"rm",
			indexID,
			"--deploymentName",
			deploymentName,
			"--force",
			"--type",
			"LOCAL",
			"--username",
			dbUsername,
			"--password",
			dbUserPassword,
			"--debug",
		)

		cmd.Env = os.Environ()

		var o, e bytes.Buffer
		cmd.Stdout = &o
		cmd.Stderr = &e
		require.NoError(t, cmd.Run(), e.String())
		assert.Contains(t, o.String(), fmt.Sprintf("Index '%s' deleted", indexID))
	})

	t.Run("Pause Deployment", func(t *testing.T) {
		cmd := exec.Command(cliPath,
			deploymentEntity,
			"pause",
			deploymentName,
			"--type",
			"local",
			"--debug",
		)

		cmd.Env = os.Environ()

		r, err := cmd.CombinedOutput()
		out := string(r)
		require.NoError(t, err, out)
		assert.Contains(t, out, fmt.Sprintf("Pausing deployment '%s'", deploymentName))
	})

	t.Run("Start Deployment", func(t *testing.T) {
		cmd := exec.Command(cliPath,
			deploymentEntity,
			"start",
			deploymentName,
			"--type",
			"local",
			"--debug",
		)
		cmd.Env = os.Environ()
		r, err := cmd.CombinedOutput()
		out := string(r)
		require.NoError(t, err, out)
		assert.Contains(t, out, fmt.Sprintf("Starting deployment '%s'", deploymentName))
	})
}
