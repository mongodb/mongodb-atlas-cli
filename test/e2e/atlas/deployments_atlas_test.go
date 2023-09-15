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
	"go.mongodb.org/mongo-driver/mongo"
)

func TestDeployments(t *testing.T) {
	g := newAtlasE2ETestGenerator(t)
	g.generateProject("setup")
	cliPath, err := e2e.AtlasCLIBin()
	req := require.New(t)
	req.NoError(err)

	clusterName, err := RandClusterName()
	req.NoError(err)

	tagKey := "env"
	tagValue := "e2etestlocal"

	var connectionString string

	t.Run("Setup", func(t *testing.T) {
		cmd := exec.Command(cliPath,
			"deployments",
			"setup",
			clusterName,
			"--type",
			"atlas",
			"--force",
			"--skipMongosh",
			"--skipSampleData",
			"--tag", tagKey+"="+tagValue,
			"--debug",
			"--projectId", g.projectID,
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

	t.Run("Watch Cluster", func(t *testing.T) {
		cmd := exec.Command(cliPath,
			clustersEntity,
			"watch",
			clusterName,
			"--projectId", g.projectID,
		)
		cmd.Env = os.Environ()
		resp, err := cmd.CombinedOutput()
		req.NoError(err, string(resp))
		assert.Contains(t, string(resp), "Cluster available")
	})

	// TODO: Update with deployments delete CLOUDP-199629
	t.Cleanup(func() {
		cmd := exec.Command(cliPath,
			clustersEntity,
			"delete",
			clusterName,
			"--projectId", g.projectID,
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

	// TODO: Add support for connect CLOUDP-199422
	t.Cleanup(func() {
		_ = client.Disconnect(ctx)
	})
}
