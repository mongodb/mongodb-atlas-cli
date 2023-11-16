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
//go:build e2e || (atlas && clusters && flags)

package atlas_test

import (
	"encoding/json"
	"os"
	"os/exec"
	"testing"

	"github.com/mongodb/mongodb-atlas-cli/test/e2e"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	atlasv2 "go.mongodb.org/atlas-sdk/v20231115002/admin"
)

const writeConcern = "majority"

func TestClustersFlags(t *testing.T) {
	g := newAtlasE2ETestGenerator(t)
	g.generateProject("clustersFlags")

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
			"--enableTerminationProtection",
			"--projectId", g.projectID,
			"--tag", "env=test", "-w",
			"-o=json")
		cmd.Env = os.Environ()
		resp, err := cmd.CombinedOutput()
		req.NoError(err, string(resp))

		var cluster *atlasv2.AdvancedClusterDescription
		err = json.Unmarshal(resp, &cluster)
		req.NoError(err)

		ensureCluster(t, cluster, clusterName, e2eMDBVer, 30, true)
	})

	t.Run("Load Sample Data", func(t *testing.T) {
		cmd := exec.Command(cliPath,
			clustersEntity,
			"sampleData",
			"load",
			clusterName,
			"--projectId", g.projectID,
			"-o=json")
		cmd.Env = os.Environ()
		resp, err := cmd.CombinedOutput()
		req.NoError(err, string(resp))

		var job *atlasv2.SampleDatasetStatus
		err = json.Unmarshal(resp, &job)
		req.NoError(err)

		a := assert.New(t)
		a.Equal(clusterName, job.GetClusterName())
	})

	t.Run("List", func(t *testing.T) {
		cmd := exec.Command(cliPath,
			clustersEntity,
			"ls",
			"--projectId", g.projectID,
			"-o=json")
		cmd.Env = os.Environ()
		resp, err := cmd.CombinedOutput()
		req.NoError(err, string(resp))

		var clusters atlasv2.PaginatedAdvancedClusterDescription
		err = json.Unmarshal(resp, &clusters)
		req.NoError(err)

		a := assert.New(t)
		a.NotEmpty(clusters.Results)
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

		var cluster atlasv2.AdvancedClusterDescription
		err = json.Unmarshal(resp, &cluster)
		req.NoError(err)

		a := assert.New(t)
		a.Equal(clusterName, cluster.GetName())
	})

	t.Run("Describe Connection String", func(t *testing.T) {
		cmd := exec.Command(cliPath,
			clustersEntity,
			"cs",
			"describe",
			clusterName,
			"--projectId", g.projectID,
			"-o=json")
		cmd.Env = os.Environ()
		resp, err := cmd.CombinedOutput()
		req.NoError(err, string(resp))

		var connectionString atlasv2.ClusterConnectionStrings
		err = json.Unmarshal(resp, &connectionString)
		req.NoError(err)

		a := assert.New(t)
		a.NotEmpty(connectionString.GetStandard())
		a.NotEmpty(connectionString.GetStandardSrv())
	})

	t.Run("Update Advanced Configuration Settings", func(t *testing.T) {
		cmd := exec.Command(cliPath,
			clustersEntity,
			"advancedSettings",
			"update",
			clusterName,
			"--writeConcern",
			writeConcern,
			"--projectId", g.projectID,
			"-o=json")
		cmd.Env = os.Environ()
		resp, err := cmd.CombinedOutput()
		req.NoError(err, string(resp))
	})

	t.Run("Describe Advanced Configuration Settings", func(t *testing.T) {
		cmd := exec.Command(cliPath,
			clustersEntity,
			"advancedSettings",
			"describe",
			clusterName,
			"--projectId", g.projectID,
			"-o=json")
		cmd.Env = os.Environ()
		resp, err := cmd.CombinedOutput()
		req.NoError(err, string(resp))

		var config atlasv2.ClusterDescriptionProcessArgs
		err = json.Unmarshal(resp, &config)
		req.NoError(err)

		a := assert.New(t)
		a.NotEmpty(config.GetMinimumEnabledTlsProtocol())
		a.Equal(writeConcern, config.GetDefaultWriteConcern())
	})

	t.Run("Create Rolling Index", func(t *testing.T) {
		cmd := exec.Command(cliPath,
			clustersEntity,
			"indexes",
			"create",
			"--clusterName", clusterName,
			"--db=tes",
			"--collection=tes",
			"--key=name:1",
			"--projectId", g.projectID,
		)
		cmd.Env = os.Environ()
		resp, err := cmd.CombinedOutput()
		req.NoError(err, string(resp))
	})

	t.Run("Fail Delete for Termination Protection enabled", func(t *testing.T) {
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

	t.Run("Update", func(t *testing.T) {
		cmd := exec.Command(cliPath,
			clustersEntity,
			"update",
			clusterName,
			"--diskSizeGB", diskSizeGB40,
			"--mdbVersion=5.0",
			"--disableTerminationProtection",
			"--projectId", g.projectID,
			"-o=json")
		cmd.Env = os.Environ()
		resp, err := cmd.CombinedOutput()
		req.NoError(err, string(resp))

		var cluster atlasv2.AdvancedClusterDescription
		err = json.Unmarshal(resp, &cluster)
		req.NoError(err)

		ensureCluster(t, &cluster, clusterName, "5.0", 40, false)
	})

	t.Run("Delete", func(t *testing.T) {
		cmd := exec.Command(cliPath, clustersEntity, "delete", clusterName, "--projectId", g.projectID, "--force", "-w")
		cmd.Env = os.Environ()
		resp, err := cmd.CombinedOutput()
		req.NoError(err)

		expected := "Cluster deleted"
		a := assert.New(t)
		a.Contains(string(resp), expected)
	})
}
