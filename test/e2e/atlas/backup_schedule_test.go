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
//go:build e2e || (atlas && backup)

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
	atlas "go.mongodb.org/atlas/mongodbatlas"
)

func TestSchedules(t *testing.T) {
	cliPath, err := e2e.AtlasCLIBin()
	r := require.New(t)
	r.NoError(err)

	clusterName, err := RandClusterName()
	r.NoError(err)
	fmt.Println(clusterName)

	var policy *atlas.CloudProviderSnapshotBackupPolicy

	t.Run("Create cluster", func(t *testing.T) {
		cmd := exec.Command(cliPath,
			clustersEntity,
			"create",
			clusterName,
			"--backup",
			"--tier", tierM10,
			"--region=US_EAST_1",
			"--provider", e2eClusterProvider,
			"--mdbVersion", e2eSharedMDBVer,
			"-o=json")
		cmd.Env = os.Environ()
		resp, err := cmd.CombinedOutput()
		r.NoError(err, string(resp))

		var cluster *atlas.Cluster
		err = json.Unmarshal(resp, &cluster)
		r.NoError(err)

		ensureSharedCluster(t, cluster, clusterName, e2eSharedMDBVer, tierM10, 10, false)
	})

	t.Run("Watch create cluster", func(t *testing.T) {
		cmd := exec.Command(cliPath,
			clustersEntity,
			"watch",
			clusterName,
		)
		cmd.Env = os.Environ()
		resp, _ := cmd.CombinedOutput()
		t.Log(string(resp))
	})

	t.Run("Describe", func(t *testing.T) {
		cmd := exec.Command(cliPath,
			backupsEntity,
			"schedule",
			"describe",
			clusterName,
			"-o=json",
		)
		cmd.Env = os.Environ()
		resp, err := cmd.CombinedOutput()
		r.NoError(err)

		err = json.Unmarshal(resp, &policy)
		r.NoError(err)

		r.Equal(clusterName, policy.ClusterName)
	})

	t.Run("Update", func(t *testing.T) {
		cmd := exec.Command(cliPath,
			backupsEntity,
			"schedule",
			"update",
			clusterName,
			"-o=json",
		)
		cmd.Env = os.Environ()
		resp, err := cmd.CombinedOutput()
		r.NoError(err)

		err = json.Unmarshal(resp, &policy)
		r.NoError(err)

		r.Equal(clusterName, policy.ClusterName)
	})

	t.Run("Delete cluster", func(t *testing.T) {
		cmd := exec.Command(cliPath,
			clustersEntity,
			"delete",
			clusterName,
			"--force",
		)
		cmd.Env = os.Environ()
		resp, err := cmd.CombinedOutput()
		r.NoError(err, string(resp))

		expected := fmt.Sprintf("Cluster '%s' deleted\n", clusterName)
		a := assert.New(t)
		a.Equal(expected, string(resp))
	})

	t.Run("Watch delete cluster", func(t *testing.T) {
		cmd := exec.Command(cliPath,
			clustersEntity,
			"watch",
			clusterName,
		)
		cmd.Env = os.Environ()
		resp, _ := cmd.CombinedOutput()
		t.Log(string(resp))
	})
}
