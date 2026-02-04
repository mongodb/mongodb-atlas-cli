// Copyright 2024 MongoDB Inc
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package searchnodes

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"strings"
	"testing"

	"github.com/mongodb/mongodb-atlas-cli/atlascli/test/internal"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	atlasv2 "go.mongodb.org/atlas-sdk/v20250312013/admin"
)

const (
	clustersEntity = "clusters"
	searchEntity   = "search"
	nodesEntity    = "nodes"
	tierM20        = "M20"
)

const minSearchNodesMDBVersion = "7.0"

func TestSearchNodes(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping test in short mode")
	}

	g := internal.NewAtlasE2ETestGenerator(t, internal.WithSnapshot())
	cliPath, err := internal.AtlasCLIBin()
	req := require.New(t)
	req.NoError(err)

	g.GenerateProject("searchNodes")
	g.Tier = tierM20
	g.MDBVer = minSearchNodesMDBVersion
	g.GenerateCluster()

	g.Run("Verify no search node setup yet", func(t *testing.T) { //nolint:thelper // g.Run replaces t.Run
		cmd := exec.Command(cliPath,
			clustersEntity,
			searchEntity,
			nodesEntity,
			"list",
			"--clusterName", g.ClusterName,
			"--projectId", g.ProjectID,
			"-o=json",
			"-P",
			internal.ProfileName(),
		)

		stdout, stderr, err := internal.RunAndGetSeparateStdOutAndErr(cmd)
		require.NoError(t, err, string(stderr))
		require.Equal(t, "{}\n", string(stdout))
	})

	g.Run("Create search node", func(t *testing.T) { //nolint:thelper // g.Run replaces t.Run
		cmd := exec.Command(cliPath,
			clustersEntity,
			searchEntity,
			nodesEntity,
			"create",
			"--clusterName", g.ClusterName,
			"--projectId", g.ProjectID,
			"--file", "testdata/search_nodes_spec.json",
			"-w",
			"-o=json",
			"-P",
			internal.ProfileName(),
		)

		cmd.Env = os.Environ()
		resp, err := internal.RunAndGetStdOut(cmd)
		resp = bytes.TrimLeft(resp, ".")

		require.NoError(t, err, resp)
		var searchNode atlasv2.ApiSearchDeploymentResponse
		require.NoError(t, json.Unmarshal(resp, &searchNode))

		assert.Equal(t, g.ProjectID, searchNode.GetGroupId())
		assert.Equal(t, []atlasv2.ApiSearchDeploymentSpec{
			{
				InstanceSize: "S30_LOWCPU_NVME",
				NodeCount:    2,
			},
		}, searchNode.GetSpecs())
	})

	g.Run("List + verify created", func(t *testing.T) { //nolint:thelper // g.Run replaces t.Run
		cmd := exec.Command(cliPath,
			clustersEntity,
			searchEntity,
			nodesEntity,
			"list",
			"--clusterName", g.ClusterName,
			"--projectId", g.ProjectID,
			"-o=json",
			"-P",
			internal.ProfileName(),
		)

		cmd.Env = os.Environ()
		resp, err := internal.RunAndGetStdOut(cmd)
		resp = bytes.TrimLeft(resp, ".")

		require.NoError(t, err, resp)
		var searchNode atlasv2.ApiSearchDeploymentResponse
		require.NoError(t, json.Unmarshal(resp, &searchNode))

		assert.Equal(t, g.ProjectID, searchNode.GetGroupId())
		assert.Equal(t, []atlasv2.ApiSearchDeploymentSpec{
			{
				InstanceSize: "S30_LOWCPU_NVME",
				NodeCount:    2,
			},
		}, searchNode.GetSpecs())
	})

	g.Run("Update search node", func(t *testing.T) { //nolint:thelper // g.Run replaces t.Run
		cmd := exec.Command(cliPath,
			clustersEntity,
			searchEntity,
			nodesEntity,
			"update",
			"--clusterName", g.ClusterName,
			"--projectId", g.ProjectID,
			"--file", "testdata/search_nodes_spec_update.json",
			"-w",
			"-o=json",
			"-P",
			internal.ProfileName(),
		)

		cmd.Env = os.Environ()
		resp, err := internal.RunAndGetStdOut(cmd)
		resp = bytes.TrimLeft(resp, ".")

		require.NoError(t, err, resp)
		var searchNode atlasv2.ApiSearchDeploymentResponse
		require.NoError(t, json.Unmarshal(resp, &searchNode))

		assert.Equal(t, g.ProjectID, searchNode.GetGroupId())
		assert.Equal(t, []atlasv2.ApiSearchDeploymentSpec{
			{
				InstanceSize: "S20_HIGHCPU_NVME",
				NodeCount:    3,
			},
		}, searchNode.GetSpecs())
	})

	g.Run("List + verify updated", func(t *testing.T) { //nolint:thelper // g.Run replaces t.Run
		cmd := exec.Command(cliPath,
			clustersEntity,
			searchEntity,
			nodesEntity,
			"list",
			"--clusterName", g.ClusterName,
			"--projectId", g.ProjectID,
			"-o=json",
			"-P",
			internal.ProfileName(),
		)

		cmd.Env = os.Environ()
		resp, err := internal.RunAndGetStdOut(cmd)
		resp = bytes.TrimLeft(resp, ".")

		require.NoError(t, err, resp)
		var searchNode atlasv2.ApiSearchDeploymentResponse
		require.NoError(t, json.Unmarshal(resp, &searchNode))

		assert.Equal(t, g.ProjectID, searchNode.GetGroupId())
		assert.Equal(t, []atlasv2.ApiSearchDeploymentSpec{
			{
				InstanceSize: "S20_HIGHCPU_NVME",
				NodeCount:    3,
			},
		}, searchNode.GetSpecs())
	})

	g.Run("Delete search nodes", func(t *testing.T) { //nolint:thelper // g.Run replaces t.Run
		cmd := exec.Command(cliPath,
			clustersEntity,
			searchEntity,
			nodesEntity,
			"delete",
			"--clusterName", g.ClusterName,
			"--projectId", g.ProjectID,
			"--force",
			"-o=json",
			"-P",
			internal.ProfileName(),
		)

		cmd.Env = os.Environ()
		resp, err := internal.RunAndGetStdOut(cmd)
		respStr := strings.TrimLeft(string(resp), ".")

		require.NoError(t, err, respStr)

		expected := fmt.Sprintf("\"%s\"\n", g.ClusterName)
		assert.Equal(t, expected, respStr)
	})
}
