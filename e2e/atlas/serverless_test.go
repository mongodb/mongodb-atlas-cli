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
// +build e2e atlas,serverless

package atlas_test

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"testing"

	"github.com/mongodb/mongocli/e2e"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.mongodb.org/atlas/mongodbatlas"
)

func TestServerless(t *testing.T) {
	cliPath, err := e2e.Bin()
	req := require.New(t)
	req.NoError(err)

	clusterName, err := RandClusterName()
	req.NoError(err)

	t.Run("Create", func(t *testing.T) {
		cmd := exec.Command(cliPath,
			atlasEntity,
			serverlessEntity,
			"create",
			clusterName,
			"--region=US_EAST_1",
			"--provider=AWS",
			"-o=json")
		cmd.Env = os.Environ()
		resp, err := cmd.CombinedOutput()
		req.NoError(err, string(resp))

		var cluster *mongodbatlas.Cluster
		err = json.Unmarshal(resp, &cluster)
		req.NoError(err)

		t.Helper()
		a := assert.New(t)
		a.Equal(clusterName, cluster.Name)
	})

	t.Run("Watch", func(t *testing.T) {
		cmd := exec.Command(cliPath,
			atlasEntity,
			serverlessEntity,
			"watch",
			clusterName,
		)
		cmd.Env = os.Environ()
		resp, err := cmd.CombinedOutput()
		req.NoError(err, string(resp))

		a := assert.New(t)
		a.Contains(string(resp), "Instance available")
	})

	t.Run("List", func(t *testing.T) {
		cmd := exec.Command(cliPath,
			atlasEntity,
			serverlessEntity,
			"ls",
			"-o=json")
		cmd.Env = os.Environ()
		resp, err := cmd.CombinedOutput()
		req.NoError(err, string(resp))

		var clusters mongodbatlas.ClustersResponse
		err = json.Unmarshal(resp, &clusters)
		req.NoError(err)

		a := assert.New(t)
		a.NotEmpty(clusters.Results)
	})

	t.Run("Describe", func(t *testing.T) {
		cmd := exec.Command(cliPath,
			atlasEntity,
			serverlessEntity,
			"describe",
			clusterName,
			"-o=json")
		cmd.Env = os.Environ()
		resp, err := cmd.CombinedOutput()
		req.NoError(err, string(resp))

		var cluster mongodbatlas.Cluster
		err = json.Unmarshal(resp, &cluster)
		req.NoError(err)

		a := assert.New(t)
		a.Equal(clusterName, cluster.Name)
	})

	t.Run("Delete", func(t *testing.T) {
		cmd := exec.Command(
			cliPath,
			atlasEntity,
			serverlessEntity,
			"delete",
			clusterName,
			"--force")
		cmd.Env = os.Environ()
		resp, err := cmd.CombinedOutput()
		req.NoError(err, string(resp))

		expected := fmt.Sprintf("Serverless instance '%s' deleted\n", clusterName)
		a := assert.New(t)
		a.Equal(expected, string(resp))
	})
}
