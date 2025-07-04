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
//go:build e2e || e2eSnap || (atlas && streams_with_cluster)

package streamswithclusters

import (
	"bytes"
	"encoding/json"
	"os"
	"os/exec"
	"testing"
	"text/template"

	"github.com/mongodb/mongodb-atlas-cli/atlascli/test/internal"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	atlasv2 "go.mongodb.org/atlas-sdk/v20250312005/admin"
)

func TestStreamsWithClusters(t *testing.T) {
	g := internal.NewAtlasE2ETestGenerator(t, internal.WithSnapshot())
	if internal.IsGov() {
		t.Skip("Skipping Streams integration test, Streams processing is not enabled in cloudgov")
	}

	g.GenerateProject("atlasStreams")
	req := require.New(t)

	cliPath, err := internal.AtlasCLIBin()
	req.NoError(err)

	instanceName := g.Memory("instanceName", internal.Must(internal.RandEntityWithRevision("instance"))).(string)

	g.GenerateCluster()

	g.Run("Creating a streams instance", func(t *testing.T) { //nolint:thelper // g.Run replaces t.Run
		cmd := exec.Command(cliPath,
			"streams",
			"instance",
			"create",
			"--provider",
			"AWS",
			"-r",
			"VIRGINIA_USA",
			"--tier",
			"SP30",
			instanceName,
			"-o=json",
			"--projectId",
			g.ProjectID,
		)

		cmd.Env = os.Environ()
		resp, err := internal.RunAndGetStdOut(cmd)
		req.NoError(err, string(resp))

		var instance atlasv2.StreamsTenant
		req.NoError(json.Unmarshal(resp, &instance))

		assert.Equal(t, instance.GetName(), instanceName)
	})

	g.Run("Create a streams connection with an atlas cluster", func(t *testing.T) { //nolint:thelper // g.Run replaces t.Run
		configFile, err := generateAtlasConnectionConfigFile(g.ClusterName)
		req.NoError(err)

		t.Cleanup(func() {
			req.NoError(os.Remove(configFile))
		})

		connectionName := "ClusterConn"

		streamsCmd := exec.Command(cliPath,
			"streams",
			"connection",
			"create",
			connectionName,
			"-f",
			configFile,
			"-i",
			instanceName,
			"-o=json",
			"--projectId",
			g.ProjectID,
		)

		streamsCmd.Env = os.Environ()
		streamsResp, streamsErr := internal.RunAndGetStdOut(streamsCmd)
		req.NoError(streamsErr, string(streamsResp))

		var connection atlasv2.StreamsConnection
		req.NoError(json.Unmarshal(streamsResp, &connection))

		// Assert on config from create_streams_connection_atlas_test.json
		a := assert.New(t)
		a.Equal(connectionName, *connection.Name)
		a.Equal("atlasAdmin", *connection.DbRoleToExecute.Role)
	})

	t.Run("Delete the streams instance", func(_ *testing.T) {
		cmd := exec.Command(cliPath,
			"streams",
			"instance",
			"delete",
			"--force",
			instanceName,
			"--projectId",
			g.ProjectID,
		)

		cmd.Env = os.Environ()
		resp, err := internal.RunAndGetStdOut(cmd)
		req.NoError(err, string(resp))
	})
}

func generateAtlasConnectionConfigFile(clusterName string) (string, error) {
	data := struct {
		ClusterName string
	}{
		ClusterName: clusterName,
	}

	templateFile := "testdata/create_streams_connection_atlas_test.json"
	tmpl, err := template.ParseFiles(templateFile)
	if err != nil {
		return "", err
	}

	var tempBuffer bytes.Buffer
	if err = tmpl.Execute(&tempBuffer, data); err != nil {
		return "", err
	}

	const configFile = "testdata/create_streams_connection_atlas.json"
	file, err := os.Create(configFile)
	if err != nil {
		return "", err
	}
	defer file.Close()

	_, err = tempBuffer.WriteTo(file)
	return configFile, err
}
