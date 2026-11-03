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

package clustersissfile

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"testing"
	"time"

	"github.com/mongodb/mongodb-atlas-cli/atlascli/test/internal"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.mongodb.org/atlas-sdk/v20250312027/admin"
)

const (
	clustersEntity = "clusters"
)

const (
	pauseTimeout       = 20 * time.Minute
	pauseRetryInterval = 30 * time.Second
)

func TestISSClustersFile(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping test in short mode")
	}

	g := internal.NewAtlasE2ETestGenerator(t)

	cliPath, err := internal.AtlasCLIBin()
	req := require.New(t)
	req.NoError(err)

	clusterIssFileName := internal.Must(internal.RandClusterName())

	g.Run("Create ISS Cluster via file", func(t *testing.T) { //nolint:thelper // g.Run replaces t.Run
		cmd := exec.Command(cliPath,
			clustersEntity,
			"create",
			clusterIssFileName,
			"--file", "testdata/create_iss_cluster_test.json",
			"-o=json",
			"-P",
			internal.ProfileName(),
		)

		cmd.Env = os.Environ()
		resp, err := internal.RunAndGetStdOut(cmd)
		req.NoError(err, string(resp))

		var cluster admin.ClusterDescription20240805
		req.NoError(json.Unmarshal(resp, &cluster))
		internal.EnsureClusterLatest(t, &cluster, clusterIssFileName, "8.0", 10, false)
	})

	g.Run("Get ISS cluster autoScalingMode", func(t *testing.T) { //nolint:thelper // g.Run replaces t.Run
		cmd := exec.Command(cliPath,
			clustersEntity,
			"autoScalingConfig",
			clusterIssFileName,
			"-o=json",
			"-P",
			internal.ProfileName(),
		)

		cmd.Env = os.Environ()
		resp, err := internal.RunAndGetStdOut(cmd)
		req.NoError(err, string(resp))

		var config admin.ClusterDescriptionAutoScalingModeConfiguration
		req.NoError(json.Unmarshal(resp, &config))
		assert.Equal(t, "INDEPENDENT_SHARD_SCALING", config.GetAutoScalingMode())
	})

	g.Run("Watch ISS cluster", func(_ *testing.T) {
		cmd := exec.Command(cliPath,
			clustersEntity,
			"watch",
			clusterIssFileName,
			"-P",
			internal.ProfileName(),
		)

		cmd.Env = os.Environ()
		resp, err := internal.RunAndGetStdOut(cmd)
		req.NoError(err, string(resp))
	})

	g.Run("Pause ISS cluster", func(t *testing.T) { //nolint:thelper // g.Run replaces t.Run
		// A freshly provisioned ISS cluster can still report replication lag on a secondary
		// after it reaches the IDLE state, and Atlas refuses to pause until that clears
		// (OPERATION_INVALID_MEMBER_REPLICATION_LAG). This has been observed to take around 10 minutes.
		var resp []byte
		require.EventuallyWithT(t, func(c *assert.CollectT) {
			cmd := exec.Command(cliPath,
				clustersEntity,
				"pause",
				clusterIssFileName,
				"--autoScalingMode",
				"independentShardScaling",
				"--output",
				"json",
				"-P",
				internal.ProfileName(),
			)
			cmd.Env = os.Environ()

			var out, stderr bytes.Buffer
			cmd.Stdout = &out
			cmd.Stderr = &stderr
			err := cmd.Run()
			if err != nil {
				// EventuallyWithT discards the errors it collects once the condition
				// eventually holds, so log every rejection to keep the wait diagnosable.
				t.Logf("pause was rejected, retrying in %s: %s", pauseRetryInterval, stderr.String())
				assert.NoError(c, err, stderr.String())

				return
			}
			resp = out.Bytes()
		}, pauseTimeout, pauseRetryInterval)

		var cluster admin.ClusterDescription20240805
		req.NoError(json.Unmarshal(resp, &cluster))
	})

	g.Run("Start ISS cluster", func(_ *testing.T) {
		cmd := exec.Command(cliPath,
			clustersEntity,
			"start",
			clusterIssFileName,
			"--autoScalingMode",
			"independentShardScaling",
			"--output",
			"json",
			"-P",
			internal.ProfileName(),
		)

		cmd.Env = os.Environ()
		resp, err := internal.RunAndGetStdOut(cmd)
		req.NoError(err, string(resp))

		var cluster admin.ClusterDescription20240805
		req.NoError(json.Unmarshal(resp, &cluster))
	})

	g.Run("Update ISS cluster with file", func(_ *testing.T) {
		cmd := exec.Command(cliPath,
			clustersEntity,
			"update",
			clusterIssFileName,
			"--file",
			"testdata/create_iss_cluster_test_update.json",
			"--output",
			"json",
			"-P",
			internal.ProfileName(),
		)

		cmd.Env = os.Environ()
		resp, err := internal.RunAndGetStdOut(cmd)
		req.NoError(err, string(resp))

		var cluster admin.ClusterDescription20240805
		req.NoError(json.Unmarshal(resp, &cluster))
	})

	g.Run("Delete ISS Cluster - created via file", func(_ *testing.T) {
		cmd := exec.Command(
			cliPath,
			clustersEntity,
			"delete",
			clusterIssFileName,
			"--watch",
			"--force",
			"-P",
			internal.ProfileName(),
		)

		cmd.Env = os.Environ()
		resp, err := internal.RunAndGetStdOut(cmd)
		req.NoError(err, string(resp))

		expected := fmt.Sprintf("Deleting cluster '%s'Cluster deleted\n", clusterIssFileName)
		assert.Equal(t, expected, string(resp))
	})
}
