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
//go:build e2e || (atlas && deployments && local)

package atlas_test

import (
	"bytes"
	"context"
	"fmt"
	"os"
	"os/exec"
	"strings"
	"testing"

	"github.com/mongodb/mongodb-atlas-cli/test/e2e"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	collectionName = "myCol"
	databaseName   = "myDB"
	indexName      = "indexTest"
)

func splitOutput(cmd *exec.Cmd) (string, string, error) {
	var o, e bytes.Buffer
	cmd.Stdout = &o
	cmd.Stderr = &e
	err := cmd.Run()
	return o.String(), e.String(), err
}

func TestDeploymentsLocal(t *testing.T) {
	const deploymentName = "test"

	g := newAtlasE2ETestGenerator(t)
	g.generateProject("deployments")

	cliPath, err := e2e.AtlasCLIBin()
	req := require.New(t)
	req.NoError(err)

	t.Run("Setup", func(t *testing.T) {
		defer func(t *testing.T) {
			t.Helper()
			cmd := exec.Command(cliPath,
				deploymentEntity,
				"diagnostics",
				deploymentName,
			)

			cmd.Env = os.Environ()

			r, errDiag := cmd.CombinedOutput()
			t.Log("Diagnostics")
			t.Log(errDiag, string(r))
		}(t)

		cmd := exec.Command(cliPath,
			deploymentEntity,
			"setup",
			deploymentName,
			"--type",
			"local",
			"--force",
		)

		cmd.Env = os.Environ()

		r, setupErr := cmd.CombinedOutput()
		req.NoError(setupErr, string(r))
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

		r, delErr := cmd.CombinedOutput()
		req.NoError(delErr, string(r))
	})

	t.Run("List deployments", func(t *testing.T) {
		cmd := exec.Command(cliPath,
			deploymentEntity,
			"list",
			"--projectId",
			g.projectID,
		)

		cmd.Env = os.Environ()

		o, e, err := splitOutput(cmd)
		req.NoError(err, e)

		outputLines := strings.Split(o, "\n")
		req.Equal(`NAME   TYPE    MDB VER   STATE`, outputLines[0])

		cols := strings.Fields(outputLines[1])
		req.Equal(deploymentName, cols[0])
		req.Equal("LOCAL", cols[1])
		req.Contains(cols[2], "7.0.")
		req.Equal("IDLE", cols[3])
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
			"--connectWith",
			"connectionString",
		)

		cmd.Env = os.Environ()

		r, err := cmd.CombinedOutput()
		req.NoError(err, string(r))

		connectionString := strings.TrimSpace(string(r))
		client, err = mongo.Connect(ctx, options.Client().ApplyURI(connectionString))
		req.NoError(err)
		myDB = client.Database(databaseName)
		myCol = myDB.Collection(collectionName)
	})

	t.Cleanup(func() {
		_ = client.Disconnect(ctx)
	})

	t.Run("Seed database", func(t *testing.T) {
		ids, err := myCol.InsertMany(ctx, []interface{}{
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
			indexName,
			"--type",
			"local",
			"--deploymentName",
			deploymentName,
			"--db",
			databaseName,
			"--collection",
			collectionName,
			"--type=LOCAL",
			"-w",
		)

		cmd.Env = os.Environ()

		r, err := cmd.CombinedOutput()
		out := string(r)
		req.NoError(err, out)
		a := assert.New(t)
		a.Contains(out, "Search index created with ID:")
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
			"--type=LOCAL",
		)

		cmd.Env = os.Environ()

		o, e, err := splitOutput(cmd)
		req.NoError(err, e)
		a := assert.New(t)
		a.Contains(o, indexName)

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
			"--type=LOCAL",
		)

		cmd.Env = os.Environ()

		r, err := cmd.CombinedOutput()
		req.NoError(err, string(r))
	})

	t.Run("Test Search Index", func(t *testing.T) {
		c, err := myCol.Aggregate(ctx, bson.A{
			bson.M{
				"$search": bson.M{
					"index": indexName,
					"text": bson.M{
						"query": "test1",
						"path":  "name",
					},
				},
			},
		})
		req.NoError(err)
		var results []bson.M
		err = c.All(ctx, &results)
		req.NoError(err)
		req.Len(results, 1)
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
			"--type=LOCAL",
			"--debug",
		)

		cmd.Env = os.Environ()

		var o, e bytes.Buffer
		cmd.Stdout = &o
		cmd.Stderr = &e
		err := cmd.Run()
		req.NoError(err, e.String())
		a := assert.New(t)
		a.Contains(o.String(), fmt.Sprintf("Index '%s' deleted", indexID))
	})

	t.Run("Pause Deployment", func(t *testing.T) {
		cmd := exec.Command(cliPath,
			deploymentEntity,
			"pause",
			deploymentName,
			"--type=LOCAL",
			"--debug",
		)

		cmd.Env = os.Environ()

		r, err := cmd.CombinedOutput()
		out := string(r)
		req.NoError(err, out)
		a := assert.New(t)
		a.Contains(out, fmt.Sprintf("Pausing deployment '%s'", deploymentName))
	})

	t.Run("Start Deployment", func(t *testing.T) {
		cmd := exec.Command(cliPath,
			deploymentEntity,
			"start",
			deploymentName,
			"--type=LOCAL",
			"--debug",
		)

		cmd.Env = os.Environ()

		r, err := cmd.CombinedOutput()
		out := string(r)
		req.NoError(err, out)
		a := assert.New(t)
		a.Contains(out, fmt.Sprintf("Starting deployment '%s'", deploymentName))
	})
}

func TestDeploymentsLocalWithAuth(t *testing.T) {
	const (
		deploymentName = "test-auth"
		dbUsername     = "admin"
		dbUserPassword = "testpwd"
	)

	g := newAtlasE2ETestGenerator(t)
	g.generateProject("auth_deployment")

	cliPath, err := e2e.AtlasCLIBin()
	req := require.New(t)
	req.NoError(err)

	t.Run("Setup", func(t *testing.T) {
		defer func(t *testing.T) {
			t.Helper()
			cmd := exec.Command(cliPath,
				deploymentEntity,
				"diagnostics",
				deploymentName,
			)

			cmd.Env = os.Environ()

			r, errDiag := cmd.CombinedOutput()
			t.Log("Diagnostics")
			t.Log(errDiag, string(r))
		}(t)

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

		r, setupErr := cmd.CombinedOutput()
		req.NoError(setupErr, string(r))
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

		r, delErr := cmd.CombinedOutput()
		req.NoError(delErr, string(r))
	})

	t.Run("List deployments", func(t *testing.T) {
		cmd := exec.Command(cliPath,
			deploymentEntity,
			"list",
			"--projectId",
			g.projectID,
		)

		cmd.Env = os.Environ()

		o, e, err := splitOutput(cmd)
		req.NoError(err, e)

		outputLines := strings.Split(o, "\n")
		req.Equal(`NAME        TYPE    MDB VER   STATE`, outputLines[0])

		cols := strings.Fields(outputLines[1])
		req.Equal(deploymentName, cols[0])
		req.Equal("LOCAL", cols[1])
		req.Contains(cols[2], "7.0.")
		req.Equal("IDLE", cols[3])
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

		r, err := cmd.CombinedOutput()
		req.NoError(err, string(r))

		connectionString := strings.TrimSpace(string(r))
		client, err = mongo.Connect(ctx, options.Client().ApplyURI(connectionString))
		req.NoError(err)
		myDB = client.Database(databaseName)
		myCol = myDB.Collection(collectionName)
	})

	t.Cleanup(func() {
		_ = client.Disconnect(ctx)
	})

	t.Run("Seed database", func(t *testing.T) {
		ids, err := myCol.InsertMany(ctx, []interface{}{
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
			indexName,
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

		r, err := cmd.CombinedOutput()
		out := string(r)
		req.NoError(err, out)
		a := assert.New(t)
		a.Contains(out, "Search index created with ID:")
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
		a := assert.New(t)
		a.Contains(o, indexName)

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

		r, err := cmd.CombinedOutput()
		req.NoError(err, string(r))
	})

	t.Run("Test Search Index", func(t *testing.T) {
		c, err := myCol.Aggregate(ctx, bson.A{
			bson.M{
				"$search": bson.M{
					"index": indexName,
					"text": bson.M{
						"query": "test1",
						"path":  "name",
					},
				},
			},
		})
		req.NoError(err)
		var results []bson.M
		err = c.All(ctx, &results)
		req.NoError(err)
		req.Len(results, 1)
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
		err := cmd.Run()
		req.NoError(err, e.String())
		a := assert.New(t)
		a.Contains(o.String(), fmt.Sprintf("Index '%s' deleted", indexID))
	})

	t.Run("Pause Deployment", func(t *testing.T) {
		cmd := exec.Command(cliPath,
			deploymentEntity,
			"pause",
			deploymentName,
			"--type=LOCAL",
			"--debug",
		)

		cmd.Env = os.Environ()

		r, err := cmd.CombinedOutput()
		out := string(r)
		req.NoError(err, out)
		a := assert.New(t)
		a.Contains(out, fmt.Sprintf("Pausing deployment '%s'", deploymentName))
	})

	t.Run("Start Deployment", func(t *testing.T) {
		cmd := exec.Command(cliPath,
			deploymentEntity,
			"start",
			deploymentName,
			"--type=LOCAL",
			"--debug",
		)

		cmd.Env = os.Environ()

		r, err := cmd.CombinedOutput()
		out := string(r)
		req.NoError(err, out)
		a := assert.New(t)
		a.Contains(out, fmt.Sprintf("Starting deployment '%s'", deploymentName))
	})
}
