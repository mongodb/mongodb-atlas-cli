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

func TestDeployments(t *testing.T) {
	req := require.New(t)

	cliPath, err := e2e.AtlasCLIBin()
	req.NoError(err)
	var connectionString string

	t.Run("Setup", func(t *testing.T) {
		defer func(t *testing.T) {
			t.Helper()
			cmd := exec.Command(cliPath,
				"deployments",
				"diagnostics",
				"test",
				"-o",
				"json",
			)

			cmd.Env = os.Environ()

			r, errDiag := cmd.CombinedOutput()
			t.Log("Diagnostics")
			t.Log(errDiag, string(r))
		}(t)

		cmd := exec.Command(cliPath,
			"deployments",
			"setup",
			deploymentName,
			"--type",
			"local",
			"--force",
			"--debug",
		)

		cmd.Env = os.Environ()

		var o, e bytes.Buffer
		cmd.Stdout = &o
		cmd.Stderr = &e
		err = cmd.Run()
		req.NoError(err, e.String())

		connectionString = strings.TrimSpace(o.String())
		connectionString = strings.Replace(connectionString, "Connection string: ", "", 1)
	})

	t.Cleanup(func() {
		cmd := exec.Command(cliPath,
			"deployments",
			"delete",
			"test",
			"--force",
		)

		cmd.Env = os.Environ()

		r, delErr := cmd.CombinedOutput()
		req.NoError(delErr, string(r))
	})

	ctx := context.Background()
	var client *mongo.Client
	var myDB *mongo.Database
	var myCol *mongo.Collection

	t.Run("Connect to database", func(t *testing.T) {
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
		result := myDB.RunCommand(ctx, bson.D{
			{
				Key:   "createSearchIndexes",
				Value: "myCol",
			},
			{
				Key: "indexes",
				Value: []bson.D{
					{
						{
							Key:   "name",
							Value: "test",
						},
						{
							Key: "definition",
							Value: bson.D{
								{
									Key:   "name",
									Value: "test",
								},
								{
									Key: "mappings",
									Value: bson.D{
										{
											Key:   "dynamic",
											Value: true,
										},
									},
								},
							},
						},
					},
				},
			},
		})
		req.NoError(result.Err())
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
				break
			}
		}
	})

	t.Run("Test Search Index", func(t *testing.T) {
		c, err := myCol.Aggregate(ctx, bson.A{
			bson.M{
				"$search": bson.M{
					"index": "test",
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

	t.Run("Index Create", func(t *testing.T) {
		t.Helper()
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
			"--debug",
		)

		cmd.Env = os.Environ()

		var o, e bytes.Buffer
		cmd.Stdout = &o
		cmd.Stderr = &e
		err = cmd.Run()
		req.NoError(err, e.String())

		a := assert.New(t)
		a.Contains(o.String(), createMessage)
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
		err = cmd.Run()
		req.NoError(err, e.String())
		a := assert.New(t)
		a.Contains(o.String(), indexName)
	})
}
