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

//go:build e2e || e2eSnap || (atlas && autogeneration)

package autogeneratedcommands

import (
	"encoding/json"
	"os"
	"os/exec"
	"testing"

	"github.com/mongodb/mongodb-atlas-cli/atlascli/test/internal"
	"github.com/stretchr/testify/require"
)

func TestAutogeneratedCommands(t *testing.T) {
	g := internal.NewAtlasE2ETestGenerator(t, internal.WithSnapshot())
	g.MDBVer = "8.0"
	g.GenerateProjectAndCluster("AutogeneratedCommands")

	cliPath, err := internal.AtlasCLIBin()
	req := require.New(t)
	req.NoError(err)

	t.Run("getCluster_RequiredFieldNotSet", func(_ *testing.T) {
		cmd := exec.Command(cliPath, "api", "clusters", "getCluster")
		cmd.Env = os.Environ()
		_, err := internal.RunAndGetStdOutAndErr(cmd)

		var exitErr = new(exec.ExitError)
		req.ErrorAs(err, &exitErr, "should fail because --clusterName is not set")
		req.ErrorContains(err, "required flag(s) \"clusterName\" not set")
	})

	type Cluster struct {
		GroupID string `json:"groupId"`
		Name    string `json:"name"`
	}

	t.Run("getCluster_OK", func(_ *testing.T) {
		cmd := exec.Command(cliPath, "api", "clusters", "getCluster", "--groupId", g.ProjectID, "--clusterName", g.ClusterName)
		cmd.Env = os.Environ()
		stdOut, stdErr, err := internal.RunAndGetSeparateStdOutAndErr(cmd)

		req.NoError(err)
		var c Cluster
		req.NoError(json.Unmarshal(stdOut, &c))
		req.Equal(g.ClusterName, c.Name)
		req.Equal(g.ProjectID, c.GroupID)

		// verify that we print the rinder for user to pin versions
		// we're checking the first and the second half of the message because we don't want to fix this test every time clusters releases a new version
		errorStr := string(stdErr)
		req.Contains(errorStr, "warning: using default API version '")
		req.Contains(errorStr, "'; consider pinning a version to ensure consisentcy when updating the CLI\n")
	})

	t.Run("getCluster_valid_version_OK", func(_ *testing.T) {
		cmd := exec.Command(cliPath, "api", "clusters", "getCluster", "--groupId", g.ProjectID, "--clusterName", g.ClusterName, "--version", "2024-08-05")
		cmd.Env = os.Environ()
		stdOut, stdErr, err := internal.RunAndGetSeparateStdOutAndErr(cmd)

		req.NoError(err)
		var c Cluster
		req.NoError(json.Unmarshal(stdOut, &c))
		req.Equal(g.ClusterName, c.Name)
		req.Equal(g.ProjectID, c.GroupID)

		// verify that we don't print the reminder for users to pin versions
		// we're checking the first and the second half of the message because we don't want to fix this test every time clusters releases a new version
		errorStr := string(stdErr)
		req.NotContains(errorStr, "warning: using default API version '")
		req.NotContains(errorStr, "'; consider pinning a version to ensure consisentcy when updating the CLI\n")
	})

	t.Run("getCluster_valid_version_OK_using_projectID_alias", func(_ *testing.T) {
		cmd := exec.Command(cliPath, "api", "clusters", "getCluster", "--projectId", g.ProjectID, "--clusterName", g.ClusterName, "--version", "2024-08-05")
		cmd.Env = os.Environ()
		stdOut, stdErr, err := internal.RunAndGetSeparateStdOutAndErr(cmd)

		req.NoError(err)
		var c Cluster
		req.NoError(json.Unmarshal(stdOut, &c))
		req.Equal(g.ClusterName, c.Name)
		req.Equal(g.ProjectID, c.GroupID)

		// verify that we don't print the reminder for users to pin versions
		// we're checking the first and the second half of the message because we don't want to fix this test every time clusters releases a new version
		errorStr := string(stdErr)
		req.NotContains(errorStr, "warning: using default API version '")
		req.NotContains(errorStr, "'; consider pinning a version to ensure consisentcy when updating the CLI\n")
	})

	t.Run("getCluster_invalid_version_warn", func(_ *testing.T) {
		cmd := exec.Command(cliPath, "api", "clusters", "getCluster", "--groupId", g.ProjectID, "--clusterName", g.ClusterName, "--version", "2020-01-01")
		cmd.Env = os.Environ()
		stdOut, stdErr, err := internal.RunAndGetSeparateStdOutAndErr(cmd)

		req.NoError(err)
		var c Cluster
		req.NoError(json.Unmarshal(stdOut, &c))
		req.Equal(g.ClusterName, c.Name)
		req.Equal(g.ProjectID, c.GroupID)

		req.Contains(string(stdErr), "warning: version '2020-01-01' is not supported for this endpoint, using default API version '")
		req.Contains(string(stdErr), "'; consider pinning a version to ensure consisentcy when updating the CLI\n")
	})
}
