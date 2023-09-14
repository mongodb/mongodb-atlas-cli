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
//go:build e2e || (atlas && generic)

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
	atlasv2 "go.mongodb.org/atlas-sdk/v20230201004/admin"
)

const (
	defaultInstanceName = "DefaultInstance"
)

func TestStreams(t *testing.T) {
	if IsGov() {
		t.Skip("Skipping Streams integration test, Streams processing is not enabled in cloudgov")
	}

	g := newAtlasE2ETestGenerator(t)
	g.generateProject("atlasStreams")

	a := assert.New(t)
	req := require.New(t)

	cliPath, err := e2e.AtlasCLIBin()
	req.NoError(err)

	instanceName, err := RandEntityWithRevision("instance")
	req.NoError(err)

	connectionName, err := RandEntityWithRevision("connection")
	req.NoError(err)

	t.Run("List all streams in the e2e project", func(t *testing.T) {
		cmd := exec.Command(cliPath,
			"streams",
			"instance",
			"list",
			"-o=json",
			"--projectId",
			g.projectID,
		)

		cmd.Env = os.Environ()
		resp, err := cmd.CombinedOutput()
		req.NoError(err)

		var instances atlasv2.PaginatedApiStreamsTenant
		err = json.Unmarshal(resp, &instances)
		a.NoError(err)
		// These instances don't have a default instance, since the projects are instantiated automatically
		a.Len(instances.Results, 0)
	})

	t.Run("Creating a streams instance", func(t *testing.T) {
		cmd := exec.Command(cliPath,
			"streams",
			"instance",
			"create",
			"--provider",
			"AWS",
			"-r",
			"VIRGINIA_USA",
			instanceName,
			"-o=json",
			"--projectId",
			g.projectID,
		)

		cmd.Env = os.Environ()
		resp, err := cmd.CombinedOutput()
		req.NoError(err)

		var instance atlasv2.StreamsTenant
		err = json.Unmarshal(resp, &instance)
		a.NoError(err)
		a.Equal(*instance.Name, instanceName)
	})

	t.Run("List all streams in the e2e project after creating", func(t *testing.T) {
		cmd := exec.Command(cliPath,
			"streams",
			"instance",
			"list",
			"-o=json",
			"--projectId",
			g.projectID,
		)

		cmd.Env = os.Environ()
		resp, err := cmd.CombinedOutput()
		req.NoError(err)

		var instances atlasv2.PaginatedApiStreamsTenant
		err = json.Unmarshal(resp, &instances)
		a.NoError(err)
		a.Len(instances.Results, 1)
		a.Equal(*instances.Results[0].Name, instanceName)
		a.Equal(instances.Results[0].DataProcessRegion.CloudProvider, "AWS")
		a.Equal(instances.Results[0].DataProcessRegion.Region, "VIRGINIA_USA")
	})

	t.Run("Describing a streams instance", func(t *testing.T) {
		cmd := exec.Command(cliPath,
			"streams",
			"instance",
			"describe",
			instanceName,
			"-o=json",
			"--projectId",
			g.projectID,
		)

		cmd.Env = os.Environ()
		resp, err := cmd.CombinedOutput()
		req.NoError(err)

		var instance atlasv2.StreamsTenant
		err = json.Unmarshal(resp, &instance)
		a.NoError(err)
		a.Equal(*instance.Name, instanceName)
		a.Equal(instance.DataProcessRegion.CloudProvider, "AWS")
		a.Equal(instance.DataProcessRegion.Region, "VIRGINIA_USA")
	})

	t.Run("Updating a streams instance", func(t *testing.T) {
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
			g.projectID,
		)

		cmd.Env = os.Environ()
		resp, err := cmd.CombinedOutput()
		req.NoError(err)

		var instance atlasv2.StreamsTenant
		err = json.Unmarshal(resp, &instance)
		a.NoError(err)
		a.Equal(*instance.Name, instanceName)
		a.Equal(instance.DataProcessRegion.CloudProvider, "AWS")
		a.Equal(instance.DataProcessRegion.Region, "VIRGINIA_USA")
	})

	// Connections
	t.Run("Creating a streams connection", func(t *testing.T) {
		cmd := exec.Command(cliPath,
			"streams",
			"connection",
			"create",
			connectionName,
			"-f",
			"create_streams_connection_test.json",
			"-i",
			instanceName,
			"-o=json",
			"--projectId",
			g.projectID,
		)

		cmd.Env = os.Environ()
		resp, err := cmd.CombinedOutput()
		req.NoError(err)

		var connection atlasv2.StreamsConnection
		err = json.Unmarshal(resp, &connection)
		a.NoError(err)
		a.Equal(*connection.Name, connectionName)
	})

	t.Run("Describing a streams connection", func(t *testing.T) {
		cmd := exec.Command(cliPath,
			"streams",
			"connection",
			"describe",
			connectionName,
			"-i",
			instanceName,
			"-o=json",
			"--projectId",
			g.projectID,
		)

		cmd.Env = os.Environ()
		resp, err := cmd.CombinedOutput()
		req.NoError(err)

		var connection atlasv2.StreamsConnection
		err = json.Unmarshal(resp, &connection)
		a.NoError(err)

		a.Equal(*connection.Name, connectionName)
		a.Equal(*connection.Type, "Kafka")
		a.Equal(*connection.BootstrapServers, "example.com:8080,fraud.example.com:8000")
	})

	t.Run("Listing streams connections", func(t *testing.T) {
		cmd := exec.Command(cliPath,
			"streams",
			"connection",
			"list",
			"--instance",
			instanceName,
			"-o=json",
			"--projectId",
			g.projectID,
		)

		cmd.Env = os.Environ()
		resp, err := cmd.CombinedOutput()
		req.NoError(err)

		var response atlasv2.PaginatedApiStreamsConnection
		err = json.Unmarshal(resp, &response)
		a.NoError(err)

		connections := response.Results
		a.Len(connections, 1)
		a.Equal(*connections[0].Name, connectionName)
		a.Equal(*connections[0].Type, "Kafka")
		a.Equal(*connections[0].BootstrapServers, "example.com:8080,fraud.example.com:8000")
	})

	t.Run("Updating a streams connection", func(t *testing.T) {
		cmd := exec.Command(cliPath,
			"streams",
			"connection",
			"update",
			connectionName,
			"-f",
			"update_streams_connection_test.json",
			"-i",
			instanceName,
			"-o=json",
			"--projectId",
			g.projectID,
		)

		cmd.Env = os.Environ()
		resp, err := cmd.CombinedOutput()
		req.NoError(err)

		var connection atlasv2.StreamsConnection
		err = json.Unmarshal(resp, &connection)
		a.NoError(err)
		a.Equal(*connection.Name, connectionName)
		a.Equal(*connection.Security.Protocol, "SSL")
	})

	t.Run("Deleting a streams connection", func(t *testing.T) {
		cmd := exec.Command(cliPath,
			"streams",
			"connection",
			"delete",
			"-i",
			instanceName,
			"--force",
			connectionName,
			"--projectId",
			g.projectID,
		)

		cmd.Env = os.Environ()
		resp, err := cmd.CombinedOutput()
		req.NoError(err)

		expected := fmt.Sprintf("'%s' deleted\n", connectionName)
		a.Equal(expected, string(resp))
	})

	// Runs last after the connection work

	t.Run("Deleting a streams instance", func(t *testing.T) {
		cmd := exec.Command(cliPath,
			"streams",
			"instance",
			"delete",
			"--force",
			instanceName,
			"--projectId",
			g.projectID,
		)

		cmd.Env = os.Environ()
		resp, err := cmd.CombinedOutput()
		req.NoError(err)

		expected := fmt.Sprintf("Atlas Streams processor instance '%s' deleted\n", instanceName)
		a.Equal(expected, string(resp))
	})

}
