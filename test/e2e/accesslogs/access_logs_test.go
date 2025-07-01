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
//go:build e2e || e2eSnap || (atlas && logs)

package accesslogs

import (
	"encoding/json"
	"os"
	"os/exec"
	"testing"

	"github.com/mongodb/mongodb-atlas-cli/atlascli/test/internal"
	"github.com/stretchr/testify/require"
	atlasv2 "go.mongodb.org/atlas-sdk/v20250312005/admin"
)

const (
	accessLogsEntity = "accessLogs"
)

func TestAccessLogs(t *testing.T) {
	g := internal.NewAtlasE2ETestGenerator(t, internal.WithSnapshot())
	req := require.New(t)

	g.GenerateProjectAndCluster("accessLogs")

	h, err := g.GetHostname()
	req.NoError(err)

	cliPath, err := internal.AtlasCLIBin()
	req.NoError(err)

	g.Run("List by clusterName", func(t *testing.T) { //nolint:thelper // g.Run replaces t.Run
		cmd := exec.Command(cliPath,
			accessLogsEntity,
			"ls",
			"--clusterName", g.ClusterName,
			"--projectId", g.ProjectID,
			"-o=json")
		cmd.Env = os.Environ()
		resp, err := internal.RunAndGetStdOut(cmd)
		require.NoError(t, err, string(resp))
		var entries *atlasv2.MongoDBAccessLogsList
		require.NoError(t, json.Unmarshal(resp, &entries))
	})

	g.Run("List by hostname", func(t *testing.T) { //nolint:thelper // g.Run replaces t.Run
		cmd := exec.Command(cliPath,
			accessLogsEntity,
			"ls",
			"--hostname", h,
			"--projectId", g.ProjectID,
			"-o=json")
		cmd.Env = os.Environ()
		resp, err := internal.RunAndGetStdOut(cmd)
		require.NoError(t, err, string(resp))
		var entries *atlasv2.MongoDBAccessLogsList
		require.NoError(t, json.Unmarshal(resp, &entries))
	})
}
