// Copyright 2021 MongoDB Inc
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

package clustersm0

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"testing"

	"github.com/mongodb/mongodb-atlas-cli/atlascli/test/internal"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	atlasClustersPinned "go.mongodb.org/atlas-sdk/v20240530005/admin"
)

const (
	clustersEntity = "clusters"
	tierM0         = "M0"

	// Cluster settings.
	e2eClusterProvider = "AWS"
)

func TestClustersM0Flags(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping test in short mode")
	}

	g := internal.NewAtlasE2ETestGenerator(t, internal.WithSnapshot())
	g.GenerateProject("clustersM0")

	cliPath, err := internal.AtlasCLIBin()
	req := require.New(t)
	req.NoError(err)

	clusterName := g.Memory("clusterName", internal.Must(internal.RandClusterName())).(string)

	g.Run("Create", func(t *testing.T) { //nolint:thelper // g.Run replaces t.Run
		cmd := exec.Command(cliPath,
			clustersEntity,
			"create",
			clusterName,
			"--region", "US_EAST_1",
			"--members=3",
			"--tier", tierM0,
			"--provider", e2eClusterProvider,
			"--projectId", g.ProjectID,
			"-o=json",
			"-P",
			internal.ProfileName())
		cmd.Env = os.Environ()
		resp, err := internal.RunAndGetStdOut(cmd)
		req.NoError(err, string(resp))

		var cluster *atlasClustersPinned.AdvancedClusterDescription
		err = json.Unmarshal(resp, &cluster)
		req.NoError(err)

		internal.EnsureCluster(t, cluster, clusterName, "8.0", 0.5, false)
	})

	g.Run("Watch", func(t *testing.T) { //nolint:thelper // g.Run replaces t.Run
		cmd := exec.Command(cliPath,
			clustersEntity,
			"watch",
			clusterName,
			"--projectId", g.ProjectID,
			"-P",
			internal.ProfileName(),
		)
		cmd.Env = os.Environ()
		resp, err := internal.RunAndGetStdOut(cmd)
		req.NoError(err, string(resp))

		a := assert.New(t)
		a.Contains(string(resp), "Cluster available")
	})

	g.Run("Describe", func(t *testing.T) { //nolint:thelper // g.Run replaces t.Run
		cmd := exec.Command(cliPath,
			clustersEntity,
			"describe",
			clusterName,
			"--projectId", g.ProjectID,
			"-o=json",
			"-P",
			internal.ProfileName(),
		)
		cmd.Env = os.Environ()
		resp, err := internal.RunAndGetStdOut(cmd)
		req.NoError(err, string(resp))

		var cluster atlasClustersPinned.AdvancedClusterDescription
		err = json.Unmarshal(resp, &cluster)
		req.NoError(err)

		a := assert.New(t)
		a.Equal(clusterName, cluster.GetName())
	})

	g.Run("Delete", func(t *testing.T) { //nolint:thelper // g.Run replaces t.Run
		cmd := exec.Command(cliPath,
			clustersEntity,
			"delete",
			clusterName,
			"--force",
			"--projectId", g.ProjectID,
			"-P",
			internal.ProfileName(),
		)
		cmd.Env = os.Environ()
		resp, err := internal.RunAndGetStdOut(cmd)
		req.NoError(err, string(resp))

		expected := fmt.Sprintf("Deleting cluster '%s'", clusterName)
		a := assert.New(t)
		a.Equal(expected, string(resp))
	})

	if internal.SkipCleanup() {
		return
	}

	g.Run("Watch", func(t *testing.T) { //nolint:thelper // g.Run replaces t.Run
		cmd := exec.Command(cliPath,
			clustersEntity,
			"watch",
			clusterName,
			"--projectId", g.ProjectID,
			"-P",
			internal.ProfileName(),
		)
		cmd.Env = os.Environ()
		// this command will fail with 404 once the cluster is deleted
		// we just need to wait for this to close the project
		resp, _ := internal.RunAndGetStdOut(cmd)
		t.Log(string(resp))
	})
}
