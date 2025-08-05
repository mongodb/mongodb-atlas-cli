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

package streams

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"testing"
	"time"

	"github.com/mongodb/mongodb-atlas-cli/atlascli/test/internal"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	atlasv2 "go.mongodb.org/atlas-sdk/v20250312005/admin"
)

func TestStreams(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping test in short mode")
	}

	g := internal.NewAtlasE2ETestGenerator(t, internal.WithSnapshot())
	if internal.IsGov() {
		t.Skip("Skipping Streams integration test, Streams processing is not enabled in cloudgov")
	}

	g.GenerateProject("atlasStreams")
	req := require.New(t)

	cliPath, err := internal.AtlasCLIBin()
	req.NoError(err)

	instanceName := g.Memory("instanceName", internal.Must(internal.RandEntityWithRevision("instance"))).(string)

	connectionName := g.Memory("connectionName", internal.Must(internal.RandEntityWithRevision("connection"))).(string)

	g.Run("List all streams in the e2e project", func(t *testing.T) { //nolint:thelper // g.Run replaces t.Run
		cmd := exec.Command(cliPath,
			"streams",
			"instance",
			"list",
			"-o=json",
			"--projectId",
			g.ProjectID,
		)

		cmd.Env = os.Environ()
		resp, err := internal.RunAndGetStdOut(cmd)
		req.NoError(err, string(resp))

		var instances atlasv2.PaginatedApiStreamsTenant
		req.NoError(json.Unmarshal(resp, &instances))

		// These instances don't have a default instance, since the projects are instantiated automatically
		assert.Empty(t, instances.Results, "A new project should have no instances")
	})

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

	g.Run("Downloading streams instance logs instance", func(t *testing.T) { //nolint:thelper // g.Run replaces t.Run
		cmd := exec.Command(cliPath,
			"streams",
			"instance",
			"download",
			instanceName,
			"--out",
			"-",
			"--start",
			strconv.FormatInt(time.Now().Add(-10*time.Second).Unix(), 10),
			"--end",
			strconv.FormatInt(time.Now().Unix(), 10),
			"--force",
			"--projectId",
			g.ProjectID,
		)
		cmd.Env = os.Environ()

		resp, err := internal.RunAndGetStdOut(cmd)
		require.NoError(t, err, string(resp))
	})

	g.Run("List all streams in the e2e project after creating", func(t *testing.T) { //nolint:thelper // g.Run replaces t.Run
		cmd := exec.Command(cliPath,
			"streams",
			"instance",
			"list",
			"-o=json",
			"--projectId",
			g.ProjectID,
		)

		cmd.Env = os.Environ()
		resp, err := internal.RunAndGetStdOut(cmd)
		require.NoError(t, err, string(resp))

		var instances atlasv2.PaginatedApiStreamsTenant
		require.NoError(t, json.Unmarshal(resp, &instances))

		a := assert.New(t)
		a.Len(instances.GetResults(), 1)
		a.Equal(*instances.GetResults()[0].Name, instanceName)
		a.Equal("AWS", instances.GetResults()[0].DataProcessRegion.CloudProvider)
		a.Equal("VIRGINIA_USA", instances.GetResults()[0].DataProcessRegion.Region)
	})

	g.Run("Describing a streams instance", func(t *testing.T) { //nolint:thelper // g.Run replaces t.Run
		cmd := exec.Command(cliPath,
			"streams",
			"instance",
			"describe",
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

		a := assert.New(t)
		a.Equal(instanceName, *instance.Name)
		a.Equal("AWS", instance.DataProcessRegion.CloudProvider)
		a.Equal("VIRGINIA_USA", instance.DataProcessRegion.Region)
	})

	g.Run("Updating a streams instance", func(t *testing.T) { //nolint:thelper // g.Run replaces t.Run
		// Have to set the same values as with create, as streams currently only supports one region
		cmd := exec.Command(cliPath,
			"streams",
			"instance",
			"update",
			"--provider",
			"AWS",
			"-r",
			"VIRGINIA_USA",
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

		a := assert.New(t)
		a.Equal(*instance.Name, instanceName)
		a.Equal("AWS", instance.DataProcessRegion.CloudProvider)
		a.Equal("VIRGINIA_USA", instance.DataProcessRegion.Region)
	})

	// Endpoints
	var endpointID string
	g.Run("Creating a streams privateLink endpoint", func(t *testing.T) { //nolint:thelper // g.Run replaces t.Run
		streamsCmd := exec.Command(cliPath,
			"streams",
			"privateLink",
			"create",
			"-f",
			"testdata/create_streams_privateLink_test.json",
			"-o=json",
			"--projectId",
			g.ProjectID,
		)

		streamsCmd.Env = os.Environ()
		streamsResp, err := internal.RunAndGetStdOut(streamsCmd)
		req.NoError(err, string(streamsResp))

		var privateLinkEndpoint atlasv2.StreamsPrivateLinkConnection
		req.NoError(json.Unmarshal(streamsResp, &privateLinkEndpoint))

		a := assert.New(t)
		a.Equal("Azure", privateLinkEndpoint.GetProvider())
		a.Equal("US_EAST_2", privateLinkEndpoint.GetRegion())
		a.Equal("/subscriptions/fd01adff-b37e-4693-8497-83ecf183a145/resourceGroups/test-rg/providers/Microsoft.EventHub/namespaces/test-namespace", privateLinkEndpoint.GetServiceEndpointId())
		a.Equal("test-namespace.servicebus.windows.net", privateLinkEndpoint.GetDnsDomain())

		// Assign the endpoint ID so that it can be used in subsequent tests
		endpointID = privateLinkEndpoint.GetId()
	})

	g.Run("Describing a streams privateLink endpoint", func(t *testing.T) { //nolint:thelper // g.Run replaces t.Run
		streamsCmd := exec.Command(cliPath,
			"streams",
			"privateLink",
			"describe",
			endpointID,
			"-o=json",
			"--projectId",
			g.ProjectID,
		)

		streamsCmd.Env = os.Environ()
		streamsResp, err := internal.RunAndGetStdOut(streamsCmd)
		req.NoError(err, string(streamsResp))

		var privateLinkEndpoint atlasv2.StreamsPrivateLinkConnection
		req.NoError(json.Unmarshal(streamsResp, &privateLinkEndpoint))

		a := assert.New(t)
		a.Equal(endpointID, privateLinkEndpoint.GetId())
		a.Equal("AZURE", privateLinkEndpoint.GetProvider())
		a.Equal("US_EAST_2", privateLinkEndpoint.GetRegion())
		a.Equal("/subscriptions/fd01adff-b37e-4693-8497-83ecf183a145/resourceGroups/test-rg/providers/Microsoft.EventHub/namespaces/test-namespace", privateLinkEndpoint.GetServiceEndpointId())
		a.Equal("test-namespace.servicebus.windows.net", privateLinkEndpoint.GetDnsDomain())
	})

	g.Run("List all streams privateLink endpoints", func(t *testing.T) { //nolint:thelper // g.Run replaces t.Run
		streamsCmd := exec.Command(cliPath,
			"streams",
			"privateLinks",
			"list",
			"-o=json",
			"--projectId",
			g.ProjectID,
		)

		streamsCmd.Env = os.Environ()
		streamsResp, err := internal.RunAndGetStdOut(streamsCmd)
		req.NoError(err, string(streamsResp))

		var privateLinkEndpoints atlasv2.PaginatedApiStreamsPrivateLink
		req.NoError(json.Unmarshal(streamsResp, &privateLinkEndpoints))

		a := assert.New(t)
		a.Len(privateLinkEndpoints.GetResults(), 1)
		a.Equal("AZURE", privateLinkEndpoints.GetResults()[0].GetProvider())
		a.Equal("US_EAST_2", privateLinkEndpoints.GetResults()[0].GetRegion())
		a.Equal("/subscriptions/fd01adff-b37e-4693-8497-83ecf183a145/resourceGroups/test-rg/providers/Microsoft.EventHub/namespaces/test-namespace", privateLinkEndpoints.GetResults()[0].GetServiceEndpointId())
		a.Equal("test-namespace.servicebus.windows.net", privateLinkEndpoints.GetResults()[0].GetDnsDomain())
	})

	g.Run("Deleting a streams privateLink endpoint", func(t *testing.T) { //nolint:thelper // g.Run replaces t.Run
		streamsCmd := exec.Command(cliPath,
			"streams",
			"privateLink",
			"delete",
			endpointID,
			"--force",
			"--projectId",
			g.ProjectID,
		)

		streamsCmd.Env = os.Environ()
		resp, err := internal.RunAndGetStdOut(streamsCmd)
		require.NoError(t, err, string(resp))

		expected := fmt.Sprintf("Atlas Stream Processing PrivateLink endpoint '%s' deleted.\n", endpointID)
		assert.Equal(t, expected, string(resp))
	})

	// Connections
	g.Run("Creating a streams connection", func(t *testing.T) { //nolint:thelper // g.Run replaces t.Run
		cmd := exec.Command(cliPath,
			"streams",
			"connection",
			"create",
			connectionName,
			"-f",
			"testdata/create_streams_connection_test.json",
			"-i",
			instanceName,
			"-o=json",
			"--projectId",
			g.ProjectID,
		)

		cmd.Env = os.Environ()
		resp, err := internal.RunAndGetStdOut(cmd)
		req.NoError(err, string(resp))

		var connection atlasv2.StreamsConnection
		req.NoError(json.Unmarshal(resp, &connection))

		assert.Equal(t, connection.GetName(), connectionName)
	})

	g.Run("Describing a streams connection", func(t *testing.T) { //nolint:thelper // g.Run replaces t.Run
		cmd := exec.Command(cliPath,
			"streams",
			"connection",
			"describe",
			connectionName,
			"-i",
			instanceName,
			"-o=json",
			"--projectId",
			g.ProjectID,
		)

		cmd.Env = os.Environ()
		resp, err := internal.RunAndGetStdOut(cmd)
		require.NoError(t, err, string(resp))

		var connection atlasv2.StreamsConnection
		require.NoError(t, json.Unmarshal(resp, &connection))

		a := assert.New(t)
		a.Equal(connectionName, *connection.Name)
		a.Equal("Kafka", *connection.Type)
		a.Equal("example.com:8080,fraud.example.com:8000", *connection.BootstrapServers)
	})

	g.Run("Listing streams connections", func(t *testing.T) { //nolint:thelper // g.Run replaces t.Run
		cmd := exec.Command(cliPath,
			"streams",
			"connection",
			"list",
			"--instance",
			instanceName,
			"-o=json",
			"--projectId",
			g.ProjectID,
		)

		cmd.Env = os.Environ()
		resp, err := internal.RunAndGetStdOut(cmd)
		req.NoError(err, string(resp))

		var response atlasv2.PaginatedApiStreamsConnection
		req.NoError(json.Unmarshal(resp, &response))

		connections := response.GetResults()
		a := assert.New(t)
		a.Len(connections, 1)

		expected := []struct {
			Name, Type, BootstrapServers string
		}{
			{Name: connectionName, Type: "Kafka", BootstrapServers: "example.com:8080,fraud.example.com:8000"},
		}
		got := []struct {
			Name, Type, BootstrapServers string
		}{
			{Name: connections[0].GetName(), Type: connections[0].GetType(), BootstrapServers: connections[0].GetBootstrapServers()},
		}

		a.ElementsMatch(expected, got)
	})

	g.Run("Updating a streams connection", func(t *testing.T) { //nolint:thelper // g.Run replaces t.Run
		cmd := exec.Command(cliPath,
			"streams",
			"connection",
			"update",
			connectionName,
			"-f",
			"testdata/update_streams_connection_test.json",
			"-i",
			instanceName,
			"-o=json",
			"--projectId",
			g.ProjectID,
		)

		cmd.Env = os.Environ()
		resp, err := internal.RunAndGetStdOut(cmd)
		req.NoError(err, string(resp))

		var connection atlasv2.StreamsConnection
		req.NoError(json.Unmarshal(resp, &connection))
		a := assert.New(t)
		a.Equal(*connection.Name, connectionName)
		a.Equal("SSL", connection.Security.GetProtocol())
	})

	g.Run("Deleting a streams connection", func(t *testing.T) { //nolint:thelper // g.Run replaces t.Run
		cmd := exec.Command(cliPath,
			"streams",
			"connection",
			"delete",
			"-i",
			instanceName,
			"--force",
			connectionName,
			"--projectId",
			g.ProjectID,
		)

		cmd.Env = os.Environ()
		resp, err := internal.RunAndGetStdOut(cmd)
		require.NoError(t, err, string(resp))

		expected := fmt.Sprintf("Atlas Stream Processing connection '%s' deleted\n", connectionName)
		assert.Equal(t, expected, string(resp))
	})

	// Runs last after the connection work

	g.Run("Deleting a streams instance", func(t *testing.T) { //nolint:thelper // g.Run replaces t.Run
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
		require.NoError(t, err, string(resp))

		expected := fmt.Sprintf("Atlas Streams processor instance '%s' deleted\n", instanceName)
		assert.Equal(t, expected, string(resp))
	})
}
