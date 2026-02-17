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

package clustersflexfile

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"testing"

	"github.com/mongodb/mongodb-atlas-cli/atlascli/test/internal"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.mongodb.org/atlas-sdk/v20250312014/admin"
)

const (
	clustersEntity = "clusters"
)

// Note that the FlexClusters are only available in the 5efda6aea3f2ed2e7dd6ce05 (Atlas CLI E2E Project)
// They will be fully enabled in https://jira.mongodb.org/browse/CLOUDP-291186. We will be able to move these e2e tests
// to create their project once the ticket is completed.
func TestFlexClustersFile(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping test in short mode")
	}

	g := internal.NewAtlasE2ETestGenerator(t, internal.WithSnapshot())
	cliPath, err := internal.AtlasCLIBin()
	req := require.New(t)
	req.NoError(err)

	clusterFileName := g.Memory("clusterFileName", internal.Must(internal.RandClusterName())).(string)

	g.Run("Create Flex Cluster via file", func(t *testing.T) { //nolint:thelper // g.Run replaces t.Run
		cmd := exec.Command(cliPath,
			clustersEntity,
			"create",
			clusterFileName,
			"--file", "testdata/create_flex_cluster_test.json",
			"-o=json",
			"-P",
			internal.ProfileName())

		cmd.Env = os.Environ()
		resp, err := internal.RunAndGetStdOut(cmd)
		req.NoError(err, string(resp))

		var cluster admin.FlexClusterDescription20241113
		req.NoError(json.Unmarshal(resp, &cluster))
		internal.EnsureFlexCluster(t, &cluster, clusterFileName, 5, false)
	})

	g.Run("Delete Flex Cluster - created via file", func(t *testing.T) { //nolint:thelper // g.Run replaces t.Run
		cmd := exec.Command(
			cliPath,
			clustersEntity,
			"delete",
			clusterFileName,
			"--watch",
			"--force",
			"-P",
			internal.ProfileName())

		cmd.Env = os.Environ()
		resp, err := internal.RunAndGetStdOut(cmd)
		req.NoError(err, string(resp))

		expected := fmt.Sprintf("Deleting cluster '%s'Cluster deleted\n", clusterFileName)
		assert.Equal(t, expected, string(resp))
	})
}
