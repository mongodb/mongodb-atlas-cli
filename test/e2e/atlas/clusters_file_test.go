// Copyright 2020 MongoDB Inc
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
//go:build e2e || (atlas && clusters && file)

package atlas_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"testing"
	"text/template"

	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/config"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/test/e2e"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	atlasv2 "go.mongodb.org/atlas-sdk/v20240530002/admin"
)

func TestClustersFile(t *testing.T) {
	g := newAtlasE2ETestGenerator(t)
	g.generateProject("clustersFile")

	cliPath, err := e2e.AtlasCLIBin()
	req := require.New(t)
	req.NoError(err)

	clusterFileName, err := RandClusterName()
	req.NoError(err)

	mdbVersion, err := MongoDBMajorVersion()
	req.NoError(err)

	clusterFile, err := generateClusterFile(mdbVersion)
	req.NoError(err)

	t.Cleanup(func() {
		os.Remove(clusterFile)
	})

	t.Run("Create via file", func(t *testing.T) {
		cmd := exec.Command(cliPath,
			clustersEntity,
			"create",
			clusterFileName,
			"--file", clusterFile,
			"--projectId", g.projectID,
			"-o=json")
		cmd.Env = os.Environ()
		resp, err := e2e.RunAndGetStdOut(cmd)
		req.NoError(err, string(resp))

		var cluster atlasv2.AdvancedClusterDescription
		req.NoError(json.Unmarshal(resp, &cluster))

		ensureCluster(t, &cluster, clusterFileName, mdbVersion, 30, false)
	})

	t.Run("Watch", func(t *testing.T) {
		cmd := exec.Command(cliPath,
			clustersEntity,
			"watch",
			"--projectId", g.projectID,
			clusterFileName,
		)
		cmd.Env = os.Environ()
		resp, err := e2e.RunAndGetStdOut(cmd)
		req.NoError(err, string(resp))
		assert.Contains(t, string(resp), "Cluster available")
	})

	t.Run("Create Partial Index", func(t *testing.T) {
		cmd := exec.Command(cliPath,
			clustersEntity,
			"indexes",
			"create",
			"--clusterName", clusterFileName,
			"--file=data/create_partial_index.json",
			"--projectId", g.projectID,
		)
		cmd.Env = os.Environ()
		resp, err := e2e.RunAndGetStdOut(cmd)
		require.NoError(t, err, string(resp))
	})

	t.Run("Create Sparse Index", func(t *testing.T) {
		cmd := exec.Command(cliPath,
			clustersEntity,
			"indexes",
			"create",
			"--clusterName", clusterFileName,
			"--file=data/create_sparse_index.json",
			"--projectId", g.projectID,
		)
		cmd.Env = os.Environ()
		resp, err := e2e.RunAndGetStdOut(cmd)
		require.NoError(t, err, string(resp))
	})

	t.Run("Create 2dspere Index", func(t *testing.T) {
		cmd := exec.Command(cliPath,
			clustersEntity,
			"indexes",
			"create",
			"--clusterName", clusterFileName,
			"--file=data/create_2dspere_index.json",
			"--projectId", g.projectID,
		)
		cmd.Env = os.Environ()
		resp, err := e2e.RunAndGetStdOut(cmd)
		require.NoError(t, err, string(resp))
	})

	t.Run("Update via file", func(t *testing.T) {
		cmd := exec.Command(cliPath,
			clustersEntity,
			"update",
			clusterFileName,
			"--file=data/update_cluster_test.json",
			"--projectId", g.projectID,
			"-o=json")

		cmd.Env = os.Environ()
		resp, err := e2e.RunAndGetStdOut(cmd)
		req.NoError(err, string(resp))

		var cluster atlasv2.AdvancedClusterDescription
		req.NoError(json.Unmarshal(resp, &cluster))
		t.Logf("%v\n", cluster)
		ensureCluster(t, &cluster, clusterFileName, mdbVersion, 40, false)
		assert.Empty(t, cluster.GetTags())
	})

	t.Run("Delete file creation", func(t *testing.T) {
		cmd := exec.Command(
			cliPath,
			clustersEntity,
			"delete",
			clusterFileName,
			"--projectId", g.projectID,
			"--force")
		cmd.Env = os.Environ()
		resp, err := e2e.RunAndGetStdOut(cmd)
		req.NoError(err, string(resp))

		expected := fmt.Sprintf("Deleting cluster '%s'", clusterFileName)
		assert.Equal(t, expected, string(resp))
	})

	t.Run("Watch deletion", func(t *testing.T) {
		cmd := exec.Command(cliPath,
			clustersEntity,
			"watch",
			clusterFileName,
			"--projectId", g.projectID,
		)
		cmd.Env = os.Environ()
		// this command will fail with 404 once the cluster is deleted
		// we just need to wait for this to close the project
		resp, _ := e2e.RunAndGetStdOut(cmd)
		t.Log(string(resp))
	})
}

func generateClusterFile(mdbVersion string) (string, error) {
	data := struct {
		MongoDBMajorVersion string
	}{
		MongoDBMajorVersion: mdbVersion,
	}

	templateFile := "data/create_cluster_test.json"
	if service := os.Getenv("MCLI_SERVICE"); service == config.CloudGovService {
		templateFile = "data/create_cluster_gov_test.json"
	}

	tmpl, err := template.ParseFiles(templateFile)
	if err != nil {
		return "", err
	}

	var tempBuffer bytes.Buffer
	if err = tmpl.Execute(&tempBuffer, data); err != nil {
		return "", err
	}

	const clusterFile = "data/create_cluster.json"
	file, err := os.Create(clusterFile)
	if err != nil {
		return "", err
	}
	defer file.Close()

	_, err = tempBuffer.WriteTo(file)
	return clusterFile, err
}
