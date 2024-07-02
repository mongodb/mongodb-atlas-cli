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
	"fmt"
	"os"
	"os/exec"
	"testing"

	"github.com/Masterminds/semver/v3"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/test/e2e"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	atlasv2 "go.mongodb.org/atlas-sdk/v20240530002/admin"
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

	mdbVersion, err := MongoDBMajorVersion()
	req.NoError(err)

	previousMdbVersion, err := getPreviousMajorVersion(mdbVersion)
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
			"--mdbVersion", previousMdbVersion,
			"--diskSizeGB", diskSizeGB30,
			"--enableTerminationProtection",
			"--projectId", g.projectID,
			"--tag", "env=test", "-w",
			"-o=json")
		cmd.Env = os.Environ()
		resp, err := e2e.RunAndGetStdOut(cmd)
		require.NoError(t, err, string(resp))

		var cluster *atlasv2.AdvancedClusterDescription
		require.NoError(t, json.Unmarshal(resp, &cluster))

		ensureCluster(t, cluster, clusterName, previousMdbVersion, 30, true)
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
		resp, err := e2e.RunAndGetStdOut(cmd)
		require.NoError(t, err, string(resp))

		var job *atlasv2.SampleDatasetStatus
		require.NoError(t, json.Unmarshal(resp, &job))
		assert.Equal(t, clusterName, job.GetClusterName())
	})

	t.Run("List", func(t *testing.T) {
		cmd := exec.Command(cliPath,
			clustersEntity,
			"ls",
			"--projectId", g.projectID,
			"-o=json")
		cmd.Env = os.Environ()
		resp, err := e2e.RunAndGetStdOut(cmd)
		require.NoError(t, err, string(resp))

		var clusters atlasv2.PaginatedAdvancedClusterDescription
		require.NoError(t, json.Unmarshal(resp, &clusters))
		assert.NotEmpty(t, clusters.Results)
	})

	t.Run("Describe", func(t *testing.T) {
		cmd := exec.Command(cliPath,
			clustersEntity,
			"describe",
			clusterName,
			"--projectId", g.projectID,
			"-o=json")
		cmd.Env = os.Environ()
		resp, err := e2e.RunAndGetStdOut(cmd)
		require.NoError(t, err, string(resp))

		var cluster atlasv2.AdvancedClusterDescription
		require.NoError(t, json.Unmarshal(resp, &cluster))
		assert.Equal(t, clusterName, cluster.GetName())
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
		resp, err := e2e.RunAndGetStdOut(cmd)
		require.NoError(t, err, string(resp))

		var connectionString atlasv2.ClusterConnectionStrings
		require.NoError(t, json.Unmarshal(resp, &connectionString))

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
		resp, err := e2e.RunAndGetStdOut(cmd)
		require.NoError(t, err, string(resp))
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
		resp, err := e2e.RunAndGetStdOut(cmd)
		require.NoError(t, err, string(resp))

		var config atlasv2.ClusterDescriptionProcessArgs
		require.NoError(t, json.Unmarshal(resp, &config))

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
		resp, err := e2e.RunAndGetStdOut(cmd)
		require.NoError(t, err, string(resp))
	})

	t.Run("Fail Delete for Termination Protection enabled", func(t *testing.T) {
		cmd := exec.Command(cliPath,
			clustersEntity,
			"delete",
			clusterName,
			"--force",
			"--projectId", g.projectID)

		cmd.Env = os.Environ()
		resp, err := e2e.RunAndGetStdOut(cmd)
		require.Error(t, err, string(resp))
	})

	t.Run("Update", func(t *testing.T) {
		cmd := exec.Command(cliPath,
			clustersEntity,
			"update",
			clusterName,
			"--diskSizeGB", diskSizeGB40,
			"--mdbVersion", mdbVersion,
			"--disableTerminationProtection",
			"--projectId", g.projectID,
			"-o=json")
		cmd.Env = os.Environ()
		resp, err := e2e.RunAndGetStdOut(cmd)
		require.NoError(t, err, string(resp))

		var cluster atlasv2.AdvancedClusterDescription
		require.NoError(t, json.Unmarshal(resp, &cluster))

		ensureCluster(t, &cluster, clusterName, mdbVersion, 40, false)
	})

	t.Run("Delete", func(t *testing.T) {
		cmd := exec.Command(cliPath, clustersEntity, "delete", clusterName, "--projectId", g.projectID, "--force", "-w")
		cmd.Env = os.Environ()
		resp, err := e2e.RunAndGetStdOut(cmd)
		require.NoError(t, err, string(resp))

		expected := "Cluster deleted"
		assert.Contains(t, string(resp), expected)
	})
}

func getPreviousMajorVersion(version string) (string, error) {
	v, err := semver.NewVersion(version)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%d.%d", v.Major()-1, v.Minor()), nil
}
