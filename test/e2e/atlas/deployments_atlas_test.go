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
	"os"
	"os/exec"
	"testing"

	"github.com/mongodb/mongodb-atlas-cli/atlascli/test/e2e"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
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
		require.NoError(t, err, e.String())
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

		r, err := cmd.CombinedOutput()
		out := string(r)
		require.NoError(t, err, out)
		assert.Contains(t, out, "Search index created")
	})
}
