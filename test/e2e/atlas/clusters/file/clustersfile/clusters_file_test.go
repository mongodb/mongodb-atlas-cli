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

package clustersfile

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"testing"
	"text/template"

	"github.com/mongodb/mongodb-atlas-cli/atlascli/test/internal"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	atlasClustersPinned "go.mongodb.org/atlas-sdk/v20240530005/admin"
)

const (
	clustersEntity = "clusters"
	indexEntity    = "index"
)

func TestClustersFile(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping test in short mode")
	}

	g := internal.NewAtlasE2ETestGenerator(t, internal.WithSnapshot())
	g.GenerateProject("clustersFile")

	cliPath, err := internal.AtlasCLIBin()
	req := require.New(t)
	req.NoError(err)

	clusterFileName := g.Memory("clusterFileName", internal.Must(internal.RandClusterName())).(string)

	mdbVersion, err := internal.MongoDBMajorVersion()
	req.NoError(err)

	clusterFile, err := generateClusterFile(mdbVersion)
	req.NoError(err)

	t.Cleanup(func() {
		os.Remove(clusterFile)
	})

	g.Run("Create via file", func(t *testing.T) { //nolint:thelper // g.Run replaces t.Run
		cmd := exec.Command(cliPath,
			clustersEntity,
			"create",
			clusterFileName,
			"--file", clusterFile,
			"--projectId", g.ProjectID,
			"-o=json")
		cmd.Env = os.Environ()
		resp, err := internal.RunAndGetStdOut(cmd)
		req.NoError(err, string(resp))

		var cluster atlasClustersPinned.AdvancedClusterDescription
		req.NoError(json.Unmarshal(resp, &cluster))

		internal.EnsureCluster(t, &cluster, clusterFileName, mdbVersion, 30, false)
	})

	g.Run("Watch", func(t *testing.T) { //nolint:thelper // g.Run replaces t.Run
		cmd := exec.Command(cliPath,
			clustersEntity,
			"watch",
			"--projectId", g.ProjectID,
			clusterFileName,
		)
		cmd.Env = os.Environ()
		resp, err := internal.RunAndGetStdOut(cmd)
		req.NoError(err, string(resp))
		assert.Contains(t, string(resp), "Cluster available")
	})

	g.Run("Create Partial Index", func(t *testing.T) { //nolint:thelper // g.Run replaces t.Run
		t.Skip("Skipping as part of CLOUDP-272716 until fix is made for flaky test in CLOUDP-280777.")
		cmd := exec.Command(cliPath,
			clustersEntity,
			"indexes",
			"create",
			"--clusterName", clusterFileName,
			"--file=testdata/create_partial_index.json",
			"--projectId", g.ProjectID,
		)
		cmd.Env = os.Environ()
		resp, err := internal.RunAndGetStdOut(cmd)
		require.NoError(t, err, string(resp))
	})

	g.Run("Create Sparse Index", func(t *testing.T) { //nolint:thelper // g.Run replaces t.Run
		t.Skip("Skipping as part of CLOUDP-272716 until fix is made for flaky test in CLOUDP-280777.")
		cmd := exec.Command(cliPath,
			clustersEntity,
			"indexes",
			"create",
			"--clusterName", clusterFileName,
			"--file=testdata/create_sparse_index.json",
			"--projectId", g.ProjectID,
		)
		cmd.Env = os.Environ()
		resp, err := internal.RunAndGetStdOut(cmd)
		require.NoError(t, err, string(resp))
	})

	g.Run("Create 2dspere Index", func(t *testing.T) { //nolint:thelper // g.Run replaces t.Run
		t.Skip("Skipping as part of CLOUDP-272716 until fix is made for flaky test in CLOUDP-280777.")
		cmd := exec.Command(cliPath,
			clustersEntity,
			"indexes",
			"create",
			"--clusterName", clusterFileName,
			"--file=testdata/create_2dspere_index.json",
			"--projectId", g.ProjectID,
		)
		cmd.Env = os.Environ()
		resp, err := internal.RunAndGetStdOut(cmd)
		require.NoError(t, err, string(resp))
	})

	g.Run("Create index with unknown fields", func(t *testing.T) { //nolint:thelper // g.Run replaces t.Run
		var stdErr bytes.Buffer

		cmd := exec.Command(cliPath,
			clustersEntity,
			indexEntity,
			"create",
			"--clusterName", clusterFileName,
			"--file=testdata/create_index_test-unknown-fields.json",
			"--projectId", g.ProjectID,
		)

		cmd.Env = os.Environ()
		cmd.Stderr = &stdErr
		_ = cmd.Run()
		assert.Contains(t, stdErr.String(), `json: unknown field "unique"`)
	})

	g.Run("Update via file", func(t *testing.T) { //nolint:thelper // g.Run replaces t.Run
		cmd := exec.Command(cliPath,
			clustersEntity,
			"update",
			clusterFileName,
			"--file=testdata/update_cluster_test.json",
			"--projectId", g.ProjectID,
			"-o=json")

		cmd.Env = os.Environ()
		resp, err := internal.RunAndGetStdOut(cmd)
		req.NoError(err, string(resp))

		var cluster atlasClustersPinned.AdvancedClusterDescription
		req.NoError(json.Unmarshal(resp, &cluster))
		t.Logf("%v\n", cluster)
		internal.EnsureCluster(t, &cluster, clusterFileName, mdbVersion, 40, false)
		assert.Empty(t, cluster.GetTags())
	})

	g.Run("Delete file creation", func(t *testing.T) { //nolint:thelper // g.Run replaces t.Run
		cmd := exec.Command(
			cliPath,
			clustersEntity,
			"delete",
			clusterFileName,
			"--projectId", g.ProjectID,
			"--force")
		cmd.Env = os.Environ()
		resp, err := internal.RunAndGetStdOut(cmd)
		req.NoError(err, string(resp))

		expected := fmt.Sprintf("Deleting cluster '%s'", clusterFileName)
		assert.Equal(t, expected, string(resp))
	})

	if internal.SkipCleanup() {
		return
	}

	g.Run("Watch deletion", func(t *testing.T) { //nolint:thelper // g.Run replaces t.Run
		cmd := exec.Command(cliPath,
			clustersEntity,
			"watch",
			clusterFileName,
			"--projectId", g.ProjectID,
		)
		cmd.Env = os.Environ()
		// this command will fail with 404 once the cluster is deleted
		// we just need to wait for this to close the project
		resp, _ := internal.RunAndGetStdOut(cmd)
		t.Log(string(resp))
	})
}

func generateClusterFile(mdbVersion string) (string, error) {
	data := struct {
		MongoDBMajorVersion string
	}{
		MongoDBMajorVersion: mdbVersion,
	}

	templateFile := "testdata/create_cluster_test.json"

	if internal.IsGov() {
		templateFile = "testdata/create_cluster_gov_test.json"
	}

	tmpl, err := template.ParseFiles(templateFile)
	if err != nil {
		return "", err
	}

	var tempBuffer bytes.Buffer
	if err = tmpl.Execute(&tempBuffer, data); err != nil {
		return "", err
	}

	const clusterFile = "testdata/create_cluster.json"
	file, err := os.Create(clusterFile)
	if err != nil {
		return "", err
	}
	defer file.Close()

	_, err = tempBuffer.WriteTo(file)
	return clusterFile, err
}
