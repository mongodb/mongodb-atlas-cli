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
//go:build e2e || (atlas && deployments)

package atlas_test

import (
	"bytes"
	"context"
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
	deploymentName = "test"
	createMessage  = "Your search index is being created"
)

func splitOutput(cmd *exec.Cmd) (string, string, error) {
	var o, e bytes.Buffer
	cmd.Stdout = &o
	cmd.Stderr = &e
	err := cmd.Run()
	return o.String(), e.String(), err
}

func TestDeployments(t *testing.T) {
	req := require.New(t)

	cliPath, err := e2e.AtlasCLIBin()
	req.NoError(err)

	t.Run("Setup", func(t *testing.T) {
		defer func(t *testing.T) {
			t.Helper()
			cmd := exec.Command(cliPath,
				deploymentEntity,
				"diagnostics",
				deploymentName,
				"-o",
				"json",
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
			"--debug",
		)

		cmd.Env = os.Environ()

		r, setupErr := cmd.CombinedOutput()
		req.NoError(setupErr, string(r))
	})

	t.Cleanup(func() {
		cmd := exec.Command(cliPath,
			deploymentEntity,
			"delete",
			"test",
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
		)

		cmd.Env = os.Environ()

		o, e, err := splitOutput(cmd)
		req.NoError(err, e)

		a := assert.New(t)
		a.Contains(o, deploymentName)
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
			"--deploymentName",
			deploymentName,
			"--db",
			databaseName,
			"--collection",
			collectionName,
		)

		cmd.Env = os.Environ()

		r, err := cmd.CombinedOutput()
		out := string(r)
		req.NoError(err, out)
		a := assert.New(t)
		a.Contains(out, createMessage)
	})

	var indexID string
	t.Run("Wait for search index", func(t *testing.T) {
		for {
			t.Log("Waiting for index...")
			cursor, err := myCol.Aggregate(ctx, mongo.Pipeline{
				{
					{Key: "$listSearchIndexes", Value: bson.D{}},
				},
			})
			req.NoError(err)
			var results []bson.M
			req.NoError(cursor.All(ctx, &results))
			if len(results) == 0 {
				continue // no index found
			}
			status, ok := results[0]["status"].(string)
			if !ok {
				continue // no status found
			}
			if status == "STEADY" {
				indexID, _ = results[0]["id"].(string)
				break
			}
		}
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
		req.Equal(1, len(results))
	})

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
			"--debug",
		)

		cmd.Env = os.Environ()

		var o, e bytes.Buffer
		cmd.Stdout = &o
		cmd.Stderr = &e
		err := cmd.Run()
		req.NoError(err, e.String())
		a := assert.New(t)
		a.Contains(o.String(), indexName)
	})
}
