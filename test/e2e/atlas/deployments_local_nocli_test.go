// Copyright 2024 MongoDB Inc
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
//go:build e2e || (atlas && deployments && local && nocli)

package atlas_test

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"
	"strings"
	"testing"

	opt "github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli/deployments/options" //nolint:importas //unique of this test
	"github.com/mongodb/mongodb-atlas-cli/atlascli/test/e2e"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func TestDeploymentsLocalWithNoCLI(t *testing.T) {
	const (
		deploymentName = "test-nocli"
		dbUsername     = "admin"
		dbUserPassword = "testpwd"
	)

	cliPath, err := e2e.AtlasCLIBin()
	req := require.New(t)
	req.NoError(err)

	bin := "docker"
	_, err = exec.LookPath("docker")
	if errors.Is(err, exec.ErrDot) {
		err = nil
	}
	if err != nil {
		bin = "podman"
	}

	image := os.Getenv("LOCALDEV_IMAGE")
	if image == "" {
		image = opt.LocalDevImage
	}

	t.Run("Pull", func(t *testing.T) {
		cmd := exec.Command(bin,
			"pull",
			image,
		)
		r, setupErr := e2e.RunAndGetStdOut(cmd)
		require.NoError(t, setupErr, string(r))
	})

	t.Run("Setup", func(t *testing.T) {
		cmd := exec.Command(bin,
			"run",
			"-d",
			"--name", deploymentName,
			"-P",
			"-e", "MONGODB_INITDB_ROOT_USERNAME="+dbUsername,
			"-e", "MONGODB_INITDB_ROOT_PASSWORD="+dbUserPassword,
			image,
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
		assert.Contains(t, cols[2], "8.0.")
		assert.Equal(t, "IDLE", cols[3])
	})

	ctx := t.Context()
	const localFile = "sampledata.archive"
	var connectionString string

	t.Run("Get connection string", func(t *testing.T) {
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

		connectionString = strings.TrimSpace(string(r))
	})

	t.Run("Download sample dataset", func(t *testing.T) {
		out, err := os.Create(localFile)
		require.NoError(t, err)
		defer out.Close()
		req, err := http.NewRequestWithContext(ctx, http.MethodGet, "https://atlas-education.s3.amazonaws.com/sampledata.archive", nil)
		require.NoError(t, err)
		resp, err := http.DefaultClient.Do(req)
		require.NoError(t, err)
		defer resp.Body.Close()
		_, err = io.Copy(out, resp.Body)
		require.NoError(t, err)
	})

	t.Cleanup(func() {
		os.Remove(localFile)
	})

	t.Run("Seed database", func(t *testing.T) {
		cmd := exec.Command("mongorestore",
			"--uri", connectionString,
			"--archive="+localFile,
		)

		cmd.Env = os.Environ()

		r, err := e2e.RunAndGetStdOut(cmd)
		require.NoError(t, err, string(r))
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

	t.Run("Create vectorSearch Index", func(t *testing.T) {
		cmd := exec.Command(cliPath,
			deploymentEntity,
			searchEntity,
			indexEntity,
			"create",
			"--deploymentName",
			deploymentName,
			"--username",
			dbUsername,
			"--password",
			dbUserPassword,
			"--type",
			"local",
			"--file",
			"data/sample_vector_search.json",
			"-w",
		)

		cmd.Env = os.Environ()

		r, err := e2e.RunAndGetStdOut(cmd)
		out := string(r)
		require.NoError(t, err, out)
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
		client, err := mongo.Connect(ctx, options.Client().ApplyURI(connectionString))
		req.NoError(err)
		t.Cleanup(func() {
			require.NoError(t, client.Disconnect(ctx))
		})
		c, err := client.Database(databaseName).Collection(collectionName).Aggregate(ctx, bson.A{
			bson.M{
				"$search": bson.M{
					"index": searchIndexName,
					"text": bson.M{
						"query": "baseball",
						"path":  "plot",
					},
				},
			},
			bson.M{
				"$limit": 5,
			},
			bson.M{
				"$project": bson.M{
					"_id":   0,
					"title": 1,
					"plot":  1,
				},
			},
		})
		require.NoError(t, err)
		var results []bson.M
		require.NoError(t, c.All(ctx, &results))
		assert.Len(t, results, 5)
	})

	t.Run("Test vectorSearch Index", func(t *testing.T) {
		b, err := os.ReadFile("data/sample_vector_search_pipeline.json")
		req.NoError(err)

		var pipeline []map[string]any
		err = json.Unmarshal(b, &pipeline)
		req.NoError(err)

		client, err := mongo.Connect(ctx, options.Client().ApplyURI(connectionString))
		req.NoError(err)
		t.Cleanup(func() {
			require.NoError(t, client.Disconnect(ctx))
		})
		c, err := client.Database(vectorSearchDB).Collection(vectorSearchCol).Aggregate(ctx, pipeline)
		req.NoError(err)
		var results []bson.M
		req.NoError(c.All(ctx, &results))
		t.Log(results)
		req.Len(results, 10)
		for _, v := range results {
			req.Greater(v["score"], float64(0))
		}
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
