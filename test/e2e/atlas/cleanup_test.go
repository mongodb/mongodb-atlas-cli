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

//go:build e2e || (atlas && cleanup)

package atlas_test

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"testing"

	"github.com/mongodb/mongodb-atlas-cli/test/e2e"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.mongodb.org/atlas/mongodbatlas"
)

func TestCleanup(t *testing.T) {
	req := require.New(t)
	cliPath, err := e2e.AtlasCLIBin()
	req.NoError(err)

	cmd := exec.Command(cliPath,
		projectEntity,
		"list",
		"-o=json")
	cmd.Env = os.Environ()
	resp, err := cmd.CombinedOutput()
	req.NoError(err, string(resp))

	var projects mongodbatlas.Projects
	err = json.Unmarshal(resp, &projects)
	req.NoError(err)

	for _, project := range projects.Results {
		projectID := project.ID
		t.Run("deleting project "+project.ID, func(t *testing.T) {
			t.Parallel()
			if projectID == os.Getenv("MCLI_PROJECT_ID") {
				t.Skip("skipping project", projectID)
			}
			deleteClustersForProject(t, cliPath, projectID)
			deleteProjectWithRetry(t, projectID)
		})
	}

	fmt.Println(projects)
}

func deleteClustersForProject(t *testing.T, cliPath, projectID string) {
	t.Helper()
	cmd := exec.Command(cliPath,
		clustersEntity,
		"list",
		"--projectId", projectID,
		"-o=json")
	cmd.Env = os.Environ()
	resp, err := cmd.CombinedOutput()
	require.NoError(t, err, string(resp))
	var clusters mongodbatlas.AdvancedClustersResponse
	require.NoError(t, json.Unmarshal(resp, &clusters), string(resp))

	for _, cluster := range clusters.Results {
		clusterName := cluster.Name
		t.Run("deleting cluster "+clusterName, func(t *testing.T) {
			t.Parallel()
			assert.NoError(t, deleteClusterForProject(projectID, clusterName))
		})
	}
}
