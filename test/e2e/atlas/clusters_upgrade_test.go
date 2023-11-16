// Copyright 2022 MongoDB Inc
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
//go:build e2e || (atlas && clusters && upgrade)

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
	atlasv2 "go.mongodb.org/atlas-sdk/v20231115002/admin"
)

func TestSharedClusterUpgrade(t *testing.T) {
	g := newAtlasE2ETestGenerator(t)
	g.generateProject("clustersUpgrade")

	cliPath, err := e2e.AtlasCLIBin()
	req := require.New(t)
	req.NoError(err)

	clusterName, err := RandClusterName()
	req.NoError(err)

	region, err := g.newAvailableRegion(tierM2, e2eClusterProvider)
	req.NoError(err)

	t.Run("Create", func(t *testing.T) {
		cmd := exec.Command(cliPath,
			clustersEntity,
			"create",
			clusterName,
			"--region", region,
			"--tier", tierM2,
			"--provider", e2eClusterProvider,
			"--mdbVersion", e2eSharedMDBVer,
			"--enableTerminationProtection",
			"--projectId", g.projectID,
			"-o=json")
		cmd.Env = os.Environ()
		resp, err := cmd.CombinedOutput()
		req.NoError(err, string(resp))

		var cluster *atlasv2.LegacyAtlasCluster
		err = json.Unmarshal(resp, &cluster)
		req.NoError(err)

		ensureSharedCluster(t, cluster, clusterName, tierM2, 2, true)
	})

	t.Run("Watch create", func(t *testing.T) {
		cmd := exec.Command(cliPath,
			clustersEntity,
			"watch",
			clusterName,
			"--projectId", g.projectID,
		)
		cmd.Env = os.Environ()
		resp, _ := cmd.CombinedOutput()
		t.Log(string(resp))
	})

	t.Run("Fail Delete for Termination Protection", func(t *testing.T) {
		cmd := exec.Command(cliPath,
			clustersEntity,
			"delete",
			clusterName,
			"--force",
			"--projectId", g.projectID)

		cmd.Env = os.Environ()
		resp, err := cmd.CombinedOutput()
		req.Error(err, string(resp))
	})

	t.Run("Upgrade", func(t *testing.T) {
		cmd := exec.Command(cliPath,
			clustersEntity,
			"upgrade",
			clusterName,
			"--tier", tierM10,
			"--diskSizeGB", diskSizeGB40,
			"--mdbVersion=6.0",
			"--disableTerminationProtection",
			"--projectId", g.projectID,
			"--tag", "env=e2e",
			"-o=json")
		cmd.Env = os.Environ()
		resp, err := cmd.CombinedOutput()
		req.NoError(err, string(resp))
	})

	t.Run("Watch upgrade", func(t *testing.T) {
		cmd := exec.Command(cliPath,
			clustersEntity,
			"watch",
			clusterName,
			"--projectId", g.projectID,
		)
		cmd.Env = os.Environ()
		resp, _ := cmd.CombinedOutput()
		t.Log(string(resp))
	})

	t.Run("Ensure upgrade", func(t *testing.T) {
		cmd := exec.Command(cliPath,
			clustersEntity,
			"get",
			clusterName,
			"--projectId", g.projectID,
			"-o=json")
		cmd.Env = os.Environ()
		resp, err := cmd.CombinedOutput()
		req.NoError(err, string(resp))

		var clusterResponse *atlasv2.AdvancedClusterDescription
		err = json.Unmarshal(resp, &clusterResponse)
		req.NoError(err)

		ensureCluster(t, clusterResponse, clusterName, "6.0", 40, false)
	})

	t.Run("Delete", func(t *testing.T) {
		cmd := exec.Command(cliPath,
			clustersEntity,
			"delete",
			clusterName,
			"--projectId", g.projectID,
			"--force",
		)
		cmd.Env = os.Environ()
		resp, err := cmd.CombinedOutput()
		req.NoError(err, string(resp))

		expected := fmt.Sprintf("Deleting cluster '%s'", clusterName)
		a := assert.New(t)
		a.Equal(expected, string(resp))
	})

	t.Run("Watch deletion", func(t *testing.T) {
		cmd := exec.Command(cliPath,
			clustersEntity,
			"watch",
			clusterName,
			"--projectId", g.projectID,
		)
		cmd.Env = os.Environ()
		resp, _ := cmd.CombinedOutput()
		t.Log(string(resp))
	})
}
