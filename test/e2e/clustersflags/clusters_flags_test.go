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
//go:build e2e || e2eSnap || (atlas && clusters && flags)

package clustersflags

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"testing"

	"github.com/Masterminds/semver/v3"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/test/internal"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	atlasClustersPinned "go.mongodb.org/atlas-sdk/v20240530005/admin"
	atlasv2 "go.mongodb.org/atlas-sdk/v20250312005/admin"
)

const (
	clustersEntity = "clusters"
	diskSizeGB40   = "40"
	diskSizeGB30   = "30"

	// Cluster settings.
	e2eClusterProvider = "AWS"
)

const writeConcern = "majority"

func TestClustersFlags(t *testing.T) {
	g := internal.NewAtlasE2ETestGenerator(t, internal.WithSnapshot())
	g.GenerateProject("clustersFlags")

	cliPath, err := internal.AtlasCLIBin()
	req := require.New(t)
	req.NoError(err)

	clusterName := g.Memory("clusterName", internal.Must(internal.RandClusterName())).(string)

	tier := internal.E2eTier()
	region, err := g.NewAvailableRegion(tier, e2eClusterProvider)
	req.NoError(err)

	mdbVersion, err := internal.MongoDBMajorVersion()
	req.NoError(err)

	previousMdbVersion, err := getPreviousMajorVersion(mdbVersion)
	req.NoError(err)

	g.Run("Create", func(t *testing.T) { //nolint:thelper // g.Run replaces t.Run
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
			"--projectId", g.ProjectID,
			"--tag", "env=test", "-w",
			"-o=json")
		cmd.Env = os.Environ()
		resp, err := internal.RunAndGetStdOut(cmd)
		require.NoError(t, err, string(resp))

		var cluster *atlasClustersPinned.AdvancedClusterDescription
		require.NoError(t, json.Unmarshal(resp, &cluster))

		internal.EnsureCluster(t, cluster, clusterName, previousMdbVersion, 30, true)
	})

	g.Run("Load Sample Data", func(t *testing.T) { //nolint:thelper // g.Run replaces t.Run
		cmd := exec.Command(cliPath,
			clustersEntity,
			"sampleData",
			"load",
			clusterName,
			"--projectId", g.ProjectID,
			"-o=json")
		cmd.Env = os.Environ()
		resp, err := internal.RunAndGetStdOut(cmd)
		require.NoError(t, err, string(resp))

		var job *atlasv2.SampleDatasetStatus
		require.NoError(t, json.Unmarshal(resp, &job))
		assert.Equal(t, clusterName, job.GetClusterName())
	})

	g.Run("List", func(t *testing.T) { //nolint:thelper // g.Run replaces t.Run
		cmd := exec.Command(cliPath,
			clustersEntity,
			"ls",
			"--projectId", g.ProjectID,
			"-o=json")
		cmd.Env = os.Environ()
		resp, err := internal.RunAndGetStdOut(cmd)
		require.NoError(t, err, string(resp))

		var clusters atlasClustersPinned.PaginatedAdvancedClusterDescription
		require.NoError(t, json.Unmarshal(resp, &clusters))
		assert.NotEmpty(t, clusters.Results)
	})

	g.Run("Describe", func(t *testing.T) { //nolint:thelper // g.Run replaces t.Run
		cmd := exec.Command(cliPath,
			clustersEntity,
			"describe",
			clusterName,
			"--projectId", g.ProjectID,
			"-o=json")
		cmd.Env = os.Environ()
		resp, err := internal.RunAndGetStdOut(cmd)
		require.NoError(t, err, string(resp))

		var cluster atlasClustersPinned.AdvancedClusterDescription
		require.NoError(t, json.Unmarshal(resp, &cluster))
		assert.Equal(t, clusterName, cluster.GetName())
	})

	g.Run("Describe Connection String", func(t *testing.T) { //nolint:thelper // g.Run replaces t.Run
		cmd := exec.Command(cliPath,
			clustersEntity,
			"cs",
			"describe",
			clusterName,
			"--projectId", g.ProjectID,
			"-o=json")
		cmd.Env = os.Environ()
		resp, err := internal.RunAndGetStdOut(cmd)
		require.NoError(t, err, string(resp))

		var connectionString atlasv2.ClusterConnectionStrings
		require.NoError(t, json.Unmarshal(resp, &connectionString))

		a := assert.New(t)
		a.NotEmpty(connectionString.GetStandard())
		a.NotEmpty(connectionString.GetStandardSrv())
	})

	g.Run("Update Advanced Configuration Settings", func(t *testing.T) { //nolint:thelper // g.Run replaces t.Run
		cmd := exec.Command(cliPath,
			clustersEntity,
			"advancedSettings",
			"update",
			clusterName,
			"--writeConcern",
			writeConcern,
			"--projectId", g.ProjectID,
			"-o=json")
		cmd.Env = os.Environ()
		resp, err := internal.RunAndGetStdOut(cmd)
		require.NoError(t, err, string(resp))
	})

	g.Run("Describe Advanced Configuration Settings", func(t *testing.T) { //nolint:thelper // g.Run replaces t.Run
		cmd := exec.Command(cliPath,
			clustersEntity,
			"advancedSettings",
			"describe",
			clusterName,
			"--projectId", g.ProjectID,
			"-o=json")
		cmd.Env = os.Environ()
		resp, err := internal.RunAndGetStdOut(cmd)
		require.NoError(t, err, string(resp))

		var config atlasClustersPinned.ClusterDescriptionProcessArgs
		require.NoError(t, json.Unmarshal(resp, &config))

		a := assert.New(t)
		a.NotEmpty(config.GetMinimumEnabledTlsProtocol())
		a.Equal(writeConcern, config.GetDefaultWriteConcern())
	})

	g.Run("Create Rolling Index", func(t *testing.T) { //nolint:thelper // g.Run replaces t.Run
		cmd := exec.Command(cliPath,
			clustersEntity,
			"indexes",
			"create",
			"--clusterName", clusterName,
			"--db=tes",
			"--collection=tes",
			"--key=name:1",
			"--projectId", g.ProjectID,
		)
		cmd.Env = os.Environ()
		resp, err := internal.RunAndGetStdOut(cmd)
		require.NoError(t, err, string(resp))
	})

	g.Run("Fail Delete for Termination Protection enabled", func(t *testing.T) { //nolint:thelper // g.Run replaces t.Run
		cmd := exec.Command(cliPath,
			clustersEntity,
			"delete",
			clusterName,
			"--force",
			"--projectId", g.ProjectID)

		cmd.Env = os.Environ()
		resp, err := internal.RunAndGetStdOut(cmd)
		require.Error(t, err, string(resp))
	})

	g.Run("Update", func(t *testing.T) { //nolint:thelper // g.Run replaces t.Run
		cmd := exec.Command(cliPath,
			clustersEntity,
			"update",
			clusterName,
			"--diskSizeGB", diskSizeGB40,
			"--mdbVersion", mdbVersion,
			"--disableTerminationProtection",
			"--projectId", g.ProjectID,
			"-o=json")
		cmd.Env = os.Environ()
		resp, err := internal.RunAndGetStdOut(cmd)
		require.NoError(t, err, string(resp))

		var cluster atlasClustersPinned.AdvancedClusterDescription
		require.NoError(t, json.Unmarshal(resp, &cluster))

		internal.EnsureCluster(t, &cluster, clusterName, mdbVersion, 40, false)
	})

	g.Run("Delete", func(t *testing.T) { //nolint:thelper // g.Run replaces t.Run
		cmd := exec.Command(cliPath, clustersEntity, "delete", clusterName, "--projectId", g.ProjectID, "--force", "-w")
		cmd.Env = os.Environ()
		resp, err := internal.RunAndGetStdOut(cmd)
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
