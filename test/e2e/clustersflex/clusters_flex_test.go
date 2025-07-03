// Copyright 2024 MongoDB Inc
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
//go:build e2e || e2eSnap || (atlas && clusters && flex)

package clustersflex

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"testing"

	"github.com/mongodb/mongodb-atlas-cli/atlascli/test/internal"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.mongodb.org/atlas-sdk/v20250312005/admin"
)

const (
	clustersEntity = "clusters"

	// Cluster settings.
	e2eClusterProvider = "AWS"
)

// Note that the FlexClusters are only available in the 5efda6aea3f2ed2e7dd6ce05 (Atlas CLI E2E Project)
// They will be fully enabled in https://jira.mongodb.org/browse/CLOUDP-291186. We will be able to move these e2e tests
// to create their project once the ticket is completed.
func TestFlexCluster(t *testing.T) {
	g := internal.NewAtlasE2ETestGenerator(t, internal.WithSnapshot())
	cliPath, err := internal.AtlasCLIBin()
	req := require.New(t)
	req.NoError(err)

	flexClusterName := g.Memory("flexClusterName", internal.Must(internal.RandClusterName())).(string)

	g.Run("Create flex cluster", func(t *testing.T) { //nolint:thelper // g.Run replaces t.Run
		cmd := exec.Command(cliPath,
			clustersEntity,
			"create",
			flexClusterName,
			"--region=US_EAST_1",
			"--tier=FLEX",
			"--provider", e2eClusterProvider,
			"-o=json")

		cmd.Env = os.Environ()
		resp, err := internal.RunAndGetStdOut(cmd)
		req.NoError(err, string(resp))

		var cluster admin.FlexClusterDescription20241113
		req.NoError(json.Unmarshal(resp, &cluster))

		internal.EnsureFlexCluster(t, &cluster, flexClusterName, 5, false)
	})

	g.Run("Get flex cluster", func(t *testing.T) { //nolint:thelper // g.Run replaces t.Run
		cmd := exec.Command(cliPath,
			clustersEntity,
			"get",
			flexClusterName,
			"-o=json")

		cmd.Env = os.Environ()
		resp, err := internal.RunAndGetStdOut(cmd)
		req.NoError(err, string(resp))

		var cluster admin.FlexClusterDescription20241113
		req.NoError(json.Unmarshal(resp, &cluster))

		internal.EnsureFlexCluster(t, &cluster, flexClusterName, 5, false)
	})

	g.Run("List flex cluster", func(t *testing.T) { //nolint:thelper // g.Run replaces t.Run
		cmd := exec.Command(cliPath,
			clustersEntity,
			"list",
			"--tier=FLEX",
			"-o=json")
		cmd.Env = os.Environ()
		resp, err := internal.RunAndGetStdOut(cmd)
		req.NoError(err, string(resp))

		var clusters admin.PaginatedFlexClusters20241113
		req.NoError(json.Unmarshal(resp, &clusters))

		assert.Positive(t, clusters.GetTotalCount())
		assert.NotEmpty(t, clusters.Results)
	})

	g.Run("Delete flex cluster", func(t *testing.T) { //nolint:thelper // g.Run replaces t.Run
		cmd := exec.Command(cliPath,
			clustersEntity,
			"delete",
			flexClusterName,
			"--force",
			"--watch")
		cmd.Env = os.Environ()
		resp, err := internal.RunAndGetStdOut(cmd)
		req.NoError(err, string(resp))

		expected := fmt.Sprintf("Deleting cluster '%s'Cluster deleted\n", flexClusterName)
		assert.Equal(t, expected, string(resp))
	})
}
