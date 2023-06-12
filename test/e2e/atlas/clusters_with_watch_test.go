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
//go:build e2e || (atlas && clusters && watch)

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
	atlasv2 "go.mongodb.org/atlas-sdk/admin"
)

const writeConcern = "majority"

func TestClustersWatch(t *testing.T) {
	g := newAtlasE2ETestGenerator(t)
	g.generateProject("clustersWatch")

	cliPath, err := e2e.AtlasCLIBin()
	req := require.New(t)
	req.NoError(err)

	clusterName, err := RandClusterName()
	req.NoError(err)

	tier := e2eTier()
	region, err := g.newAvailableRegion(tier, e2eClusterProvider)
	req.NoError(err)

	t.Run("Create", func(t *testing.T) {
		cmd := exec.Command(cliPath,
			clustersEntity,
			"create",
			clusterName,
			"--region", region,
			"--members=3",
			"--tier", tier,
			"--provider", e2eClusterProvider,
			"--mdbVersion", e2eMDBVer,
			"--diskSizeGB", diskSizeGB30,
			"--projectId", g.projectID,
			"--tag", "env=test", "-w",
			"-o=json")
		cmd.Env = os.Environ()
		resp, err := cmd.CombinedOutput()
		t.Log(resp, err.Error())
	})

	t.Run("Describe", func(t *testing.T) {
		cmd := exec.Command(cliPath,
			clustersEntity,
			"describe",
			clusterName,
			"--projectId", g.projectID,
			"-o=json")
		cmd.Env = os.Environ()
		resp, err := cmd.CombinedOutput()
		req.NoError(err, string(resp))

		var cluster *atlasv2.ClusterDescriptionV15
		err = json.Unmarshal(resp, &cluster)
		req.NoError(err)

		ensureCluster(t, cluster, clusterName, e2eMDBVer, 30, true)
	})

	t.Run("Delete", func(t *testing.T) {
		cmd := exec.Command(cliPath, clustersEntity, "delete", clusterName, "--projectId", g.projectID, "--force", "-w")
		cmd.Env = os.Environ()
		resp, err := cmd.CombinedOutput()
		req.NoError(err)

		expected := fmt.Sprintf("Deleting cluster '%s'", clusterName)
		a := assert.New(t)
		a.Contains(string(resp), expected)
	})
}
