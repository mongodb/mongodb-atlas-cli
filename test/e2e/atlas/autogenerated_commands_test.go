// Copyright 2025 MongoDB Inc
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

//go:build e2e || (atlas && autogeneration)

package atlas_test

import (
	"encoding/json"
	"os"
	"os/exec"
	"testing"

	"github.com/mongodb/mongodb-atlas-cli/atlascli/test/e2e"
	"github.com/stretchr/testify/require"
)

func TestAutogeneratedCommands(t *testing.T) {
	g := newAtlasE2ETestGenerator(t)
	g.mDBVer = "8.0"
	g.generateProjectAndCluster("AutogeneratedCommands")

	cliPath, err := e2e.AtlasCLIBin()
	req := require.New(t)
	req.NoError(err)

	t.Run("getCluster_RequiredFieldNotSet", func(t *testing.T) {
		t.Parallel()
		cmd := exec.Command(cliPath, "api", "clusters", "getCluster")
		cmd.Env = os.Environ()
		_, err := e2e.RunAndGetStdOutAndErr(cmd)

		var exitErr = new(exec.ExitError)
		req.ErrorAs(err, &exitErr, "should fail because --clusterName is not set")
		req.ErrorContains(err, "required flag(s) \"clusterName\" not set")
	})

	type Cluster struct {
		GroupID string `json:"groupId"`
		Name    string `json:"name"`
	}

	t.Run("getCluster_OK", func(t *testing.T) {
		t.Parallel()
		cmd := exec.Command(cliPath, "api", "clusters", "getCluster", "--groupId", g.projectID, "--clusterName", g.clusterName)
		cmd.Env = os.Environ()
		resp, err := e2e.RunAndGetStdOut(cmd)

		req.NoError(err)
		var c Cluster
		req.NoError(json.Unmarshal(resp, &c))
		req.Equal(g.clusterName, c.Name)
		req.Equal(g.projectID, c.GroupID)
	})

	t.Run("getCluster_valid_version_OK", func(t *testing.T) {
		t.Parallel()
		cmd := exec.Command(cliPath, "api", "clusters", "getCluster", "--groupId", g.projectID, "--clusterName", g.clusterName, "--version", "2024-08-05")
		cmd.Env = os.Environ()
		resp, err := e2e.RunAndGetStdOut(cmd)

		req.NoError(err)
		var c Cluster
		req.NoError(json.Unmarshal(resp, &c))
		req.Equal(g.clusterName, c.Name)
		req.Equal(g.projectID, c.GroupID)
	})

	t.Run("getCluster_invalid_version_warn", func(t *testing.T) {
		t.Parallel()
		cmd := exec.Command(cliPath, "api", "clusters", "getCluster", "--groupId", g.projectID, "--clusterName", g.clusterName, "--version", "2020-01-01")
		cmd.Env = os.Environ()
		stdOut, stdErr, err := e2e.RunAndGetSeparateStdOutAndErr(cmd)

		req.NoError(err)
		var c Cluster
		req.NoError(json.Unmarshal(stdOut, &c))
		req.Equal(g.clusterName, c.Name)
		req.Equal(g.projectID, c.GroupID)

		req.Contains(string(stdErr), "warning: version '2020-01-01' is not supported for this endpoint, falling back to default version:")
	})
}
