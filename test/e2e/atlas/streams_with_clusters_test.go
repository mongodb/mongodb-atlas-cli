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
//go:build e2e || (atlas && streams_with_cluster)

package atlas_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"testing"
	"text/template"

	"github.com/mongodb/mongodb-atlas-cli/atlascli/test/e2e"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	atlasv2 "go.mongodb.org/atlas-sdk/v20241113004/admin"
)

func TestStreamsWithClusters(t *testing.T) {
	if IsGov() {
		t.Skip("Skipping Streams integration test, Streams processing is not enabled in cloudgov")
	}

	g := newAtlasE2ETestGenerator(t)
	g.generateProject("atlasStreams")
	req := require.New(t)

	cliPath, err := e2e.AtlasCLIBin()
	req.NoError(err)

	instanceName, err := RandEntityWithRevision("instance")
	req.NoError(err)

	g.generateCluster()

	t.Run("Creating a streams instance", func(t *testing.T) {
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
			g.projectID,
		)

		cmd.Env = os.Environ()
		resp, err := e2e.RunAndGetStdOut(cmd)
		req.NoError(err, string(resp))

		var instance atlasv2.StreamsTenant
		req.NoError(json.Unmarshal(resp, &instance))

		assert.Equal(t, instance.GetName(), instanceName)
	})

	t.Run("Create a streams connection with an atlas cluster", func(t *testing.T) {
		configFile, err := generateAtlasConnectionConfigFile(g.clusterName)
		req.NoError(err)

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
			g.projectID,
		)

		streamsCmd.Env = os.Environ()
		streamsResp, streamsErr := e2e.RunAndGetStdOut(streamsCmd)
		req.NoError(streamsErr, string(streamsResp))

		var connection atlasv2.StreamsConnection
		req.NoError(json.Unmarshal(streamsResp, &connection))

		// Assert on config from create_streams_connection_atlas_test.json
		a := assert.New(t)
		a.Equal(connectionName, *connection.Name)
		a.Equal("atlasAdmin", *connection.DbRoleToExecute.Role)
	})

	processorName := "ExampleSP"
	t.Run("Create streams connection to sample stream", func(t *testing.T) {
		streamsCmd := exec.Command(cliPath,
			"streams",
			"connection",
			"create",
			"-f",
			"data/create_streams_connection_sample.json",
			"-i",
			instanceName,
			"-o=json",
			"--projectId",
			g.projectID,
		)

		streamsCmd.Env = os.Environ()
		streamsResp, streamsErr := e2e.RunAndGetStdOut(streamsCmd)
		req.NoError(streamsErr, string(streamsResp))

		var connection atlasv2.StreamsConnection
		req.NoError(json.Unmarshal(streamsResp, &connection))

		// Assert on config from create_streams_connection_atlas_test.json
		a := assert.New(t)
		a.Equal("sample_stream_solar", *connection.Name)
	})

	t.Run("Create a stream processor with an atlas cluster sink", func(t *testing.T) {
		streamsCmd := exec.Command(cliPath,
			"streams",
			"processor",
			"create",
			processorName,
			"-f",
			"data/create_stream_processor_test.json",
			"-i",
			instanceName,
			"-o=json",
			"--projectId",
			g.projectID,
		)

		streamsCmd.Env = os.Environ()
		streamsResp, streamsErr := e2e.RunAndGetStdOut(streamsCmd)
		req.NoError(streamsErr, string(streamsResp))

		var processor atlasv2.StreamsProcessor
		req.NoError(json.Unmarshal(streamsResp, &processor))

		a := assert.New(t)
		a.Equal(processorName, *processor.Name)
	})

	t.Run("Describe a stream processor with an atlas cluster sink", func(t *testing.T) {
		streamsCmd := exec.Command(cliPath,
			"streams",
			"processor",
			"describe",
			processorName,
			"-i",
			instanceName,
			"--projectId",
			g.projectID,
		)

		streamsCmd.Env = os.Environ()
		streamsResp, streamsErr := e2e.RunAndGetStdOut(streamsCmd)
		req.NoError(streamsErr, string(streamsResp))

		var processor atlasv2.StreamsProcessorWithStats
		req.NoError(json.Unmarshal(streamsResp, &processor))

		a := assert.New(t)
		a.Equal(processorName, processor.Name)
	})

	t.Run("List stream processors", func(t *testing.T) {
		streamsCmd := exec.Command(cliPath,
			"streams",
			"processor",
			"list",
			"-i",
			instanceName,
			"--projectId",
			g.projectID,
		)

		streamsCmd.Env = os.Environ()
		streamsResp, streamsErr := e2e.RunAndGetStdOut(streamsCmd)
		req.NoError(streamsErr, string(streamsResp))

		var processors []atlasv2.StreamsProcessorWithStats
		req.NoError(json.Unmarshal(streamsResp, &processors))

		a := assert.New(t)
		a.Len(processors, 1)
		a.Equal(processorName, processors[0].Name)
	})

	t.Run("Start a stream processor with an atlas cluster sink", func(t *testing.T) {
		streamsCmd := exec.Command(cliPath,
			"streams",
			"processor",
			"start",
			processorName,
			"-i",
			instanceName,
			"--projectId",
			g.projectID,
		)

		streamsCmd.Env = os.Environ()
		streamsResp, streamsErr := e2e.RunAndGetStdOut(streamsCmd)
		req.NoError(streamsErr, string(streamsResp))

		a := assert.New(t)
		a.Equal("Successfully started Stream Processor\n", string(streamsResp))
	})

	t.Run("Stop a stream processor with an atlas cluster sink", func(t *testing.T) {
		streamsCmd := exec.Command(cliPath,
			"streams",
			"processor",
			"stop",
			processorName,
			"-i",
			instanceName,
			"--projectId",
			g.projectID,
		)

		streamsCmd.Env = os.Environ()
		streamsResp, streamsErr := e2e.RunAndGetStdOut(streamsCmd)
		req.NoError(streamsErr, string(streamsResp))

		a := assert.New(t)
		a.Equal("Successfully stopped Stream Processor\n", string(streamsResp))
	})

	t.Run("Delete a stream processor with an atlas cluster sink", func(t *testing.T) {
		streamsCmd := exec.Command(cliPath,
			"streams",
			"processor",
			"delete",
			processorName,
			"-i",
			instanceName,
			"--projectId",
			g.projectID,
			"--force",
		)

		streamsCmd.Env = os.Environ()
		streamsResp, streamsErr := e2e.RunAndGetStdOut(streamsCmd)
		req.NoError(streamsErr, string(streamsResp))

		a := assert.New(t)
		a.Equal(fmt.Sprintf("Atlas Stream Processor '%s' deleted\n", processorName), string(streamsResp))
	})
}

func generateAtlasConnectionConfigFile(clusterName string) (string, error) {
	data := struct {
		ClusterName string
	}{
		ClusterName: clusterName,
	}

	templateFile := "data/create_streams_connection_atlas_test.json"
	tmpl, err := template.ParseFiles(templateFile)
	if err != nil {
		return "", err
	}

	var tempBuffer bytes.Buffer
	if err = tmpl.Execute(&tempBuffer, data); err != nil {
		return "", err
	}

	const configFile = "data/create_streams_connection_atlas.json"
	file, err := os.Create(configFile)
	if err != nil {
		return "", err
	}
	defer file.Close()

	_, err = tempBuffer.WriteTo(file)
	return configFile, err
}
