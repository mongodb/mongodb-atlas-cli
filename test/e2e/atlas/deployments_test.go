// Copyright 2021 MongoDB Inc
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
	"time"

	"github.com/mongodb/mongodb-atlas-cli/internal/mongosh"
	"github.com/mongodb/mongodb-atlas-cli/test/e2e"
	"github.com/stretchr/testify/require"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func TestDeployments(t *testing.T) {
	req := require.New(t)

	cliPath, err := e2e.AtlasCLIBin()
	req.NoError(err)

	var connectionString string

	t.Run("Setup", func(t *testing.T) {
		cmd := exec.Command(cliPath,
			"deployments",
			"setup",
			"test",
			"--type",
			"local",
			"--force",
		)

		cmd.Env = os.Environ()

		var o, e bytes.Buffer
		cmd.Stdout = &o
		cmd.Stderr = &e
		err = cmd.Run()
		req.NoError(err, e.String())

		connectionString = strings.TrimSpace(o.String())
		connectionString = strings.Replace(connectionString, "Connection string: ", "", 1)

		time.Sleep(40 * time.Second) // takes 30 seconds for mongot to be available
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
		myDB = client.Database("myDB")
		myCol = myDB.Collection("myCol")
	})

	t.Cleanup(func() {
		_ = client.Disconnect(ctx)
	})

	t.Run("Seed database", func(t *testing.T) {
		_, err = myCol.InsertMany(ctx, []interface{}{
			bson.M{
				"name": "test1",
			}, bson.M{
				"name": "test2",
			},
		})
		req.NoError(err)
	})

	t.Run("Create Search Index", func(t *testing.T) {
		err = mongosh.Exec(false, connectionString, "--eval", "use 'myDB'; db.myCol.createSearchIndex('default', {'mappings': {'dynamic': true}});")
		req.NoError(err)
	})

	t.Run("Test search index", func(t *testing.T) {
		c, err := myCol.Aggregate(ctx, bson.A{
			bson.M{
				"$search": bson.M{
					"text": bson.M{
						"query": "test1",
						"path":  "name",
					},
				},
			},
		})
		req.NoError(err)
		defer c.Close(ctx)
		var results []bson.D
		err = c.All(ctx, &results)
		req.NoError(err)
		req.Equal(1, len(results))
	})
}
