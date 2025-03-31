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

//go:build e2e || e2eSnap || (atlas && search_nodes)

package e2e_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	atlasv2 "go.mongodb.org/atlas-sdk/v20250312001/admin"
)

const minSearchNodesMDBVersion = "6.0"

func TestSearchNodes(t *testing.T) {
	g := newAtlasE2ETestGenerator(t, withSnapshot())
	cliPath, err := AtlasCLIBin()
	req := require.New(t)
	req.NoError(err)

	g.generateProject("searchNodes")
	g.tier = tierM20
	g.mDBVer = minSearchNodesMDBVersion
	g.generateCluster()

	g.Run("Verify no search node setup yet", func(t *testing.T) { //nolint:thelper // g.Run replaces t.Run
		cmd := exec.Command(cliPath,
			clustersEntity,
			searchEntity,
			nodesEntity,
			"list",
			"--clusterName", g.clusterName,
			"--projectId", g.projectID,
			"-o=json",
		)

		resp, err := cmd.CombinedOutput()
		respStr := string(resp)

		require.NoError(t, err, respStr)
		require.Equal(t, "{}\n", respStr)
	})

	g.Run("Create search node", func(t *testing.T) { //nolint:thelper // g.Run replaces t.Run
		cmd := exec.Command(cliPath,
			clustersEntity,
			searchEntity,
			nodesEntity,
			"create",
			"--clusterName", g.clusterName,
			"--projectId", g.projectID,
			"--file", "data/search_nodes_spec.json",
			"-w",
			"-o=json",
		)

		cmd.Env = append(os.Environ(), "GOCOVERDIR="+os.Getenv("BINGOCOVERDIR"))
		resp, err := RunAndGetStdOut(cmd)
		resp = bytes.TrimLeft(resp, ".")

		require.NoError(t, err, resp)
		var searchNode atlasv2.ApiSearchDeploymentResponse
		require.NoError(t, json.Unmarshal(resp, &searchNode))

		assert.Equal(t, g.projectID, searchNode.GetGroupId())
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
			"--clusterName", g.clusterName,
			"--projectId", g.projectID,
			"-o=json",
		)

		cmd.Env = append(os.Environ(), "GOCOVERDIR="+os.Getenv("BINGOCOVERDIR"))
		resp, err := RunAndGetStdOut(cmd)
		resp = bytes.TrimLeft(resp, ".")

		require.NoError(t, err, resp)
		var searchNode atlasv2.ApiSearchDeploymentResponse
		require.NoError(t, json.Unmarshal(resp, &searchNode))

		assert.Equal(t, g.projectID, searchNode.GetGroupId())
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
			"--clusterName", g.clusterName,
			"--projectId", g.projectID,
			"--file", "data/search_nodes_spec_update.json",
			"-w",
			"-o=json",
		)

		cmd.Env = append(os.Environ(), "GOCOVERDIR="+os.Getenv("BINGOCOVERDIR"))
		resp, err := RunAndGetStdOut(cmd)
		resp = bytes.TrimLeft(resp, ".")

		require.NoError(t, err, resp)
		var searchNode atlasv2.ApiSearchDeploymentResponse
		require.NoError(t, json.Unmarshal(resp, &searchNode))

		assert.Equal(t, g.projectID, searchNode.GetGroupId())
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
			"--clusterName", g.clusterName,
			"--projectId", g.projectID,
			"-o=json",
		)

		cmd.Env = append(os.Environ(), "GOCOVERDIR="+os.Getenv("BINGOCOVERDIR"))
		resp, err := RunAndGetStdOut(cmd)
		resp = bytes.TrimLeft(resp, ".")

		require.NoError(t, err, resp)
		var searchNode atlasv2.ApiSearchDeploymentResponse
		require.NoError(t, json.Unmarshal(resp, &searchNode))

		assert.Equal(t, g.projectID, searchNode.GetGroupId())
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
			"--clusterName", g.clusterName,
			"--projectId", g.projectID,
			"--force",
			"-o=json",
		)

		cmd.Env = append(os.Environ(), "GOCOVERDIR="+os.Getenv("BINGOCOVERDIR"))
		resp, err := RunAndGetStdOut(cmd)
		respStr := strings.TrimLeft(string(resp), ".")

		require.NoError(t, err, respStr)

		expected := fmt.Sprintf("\"%s\"\n", g.clusterName)
		assert.Equal(t, expected, respStr)
	})
}
