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
//go:build e2e || (atlas && streams && file)

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
	cliPath, err := e2e.AtlasCLIBin()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	instanceName, err := RandEntityWithRevision("instance")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	connectionName, err := RandEntityWithRevision("connection")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	t.Run("List all streams in the e2e project", func(t *testing.T) {
		cmd := exec.Command(cliPath,
			"streams",
			"instance",
			"list",
			"-o=json",
		)

		cmd.Env = os.Environ()
		resp, err := cmd.CombinedOutput()
		a := assert.New(t)
		if a.NoError(err, string(resp)) {
			var instances atlasv2.PaginatedApiStreamsTenant
			err := json.Unmarshal(resp, &instances)
			a.NoError(err)
			// We expect a default instance always
			a.Len(instances.Results, 1)
			a.Equal(*instances.Results[0].Name, defaultInstanceName)
		}
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
		)

		cmd.Env = os.Environ()
		resp, err := cmd.CombinedOutput()
		a := assert.New(t)
		if a.NoError(err, string(resp)) {
			var instance atlasv2.StreamsTenant
			err := json.Unmarshal(resp, &instance)
			a.NoError(err)
			a.Equal(*instance.Name, instanceName)
		}
	})

	t.Run("Describing a streams instance", func(t *testing.T) {
		cmd := exec.Command(cliPath,
			"streams",
			"instance",
			"describe",
			instanceName,
			"-o=json",
		)

		cmd.Env = os.Environ()
		resp, err := cmd.CombinedOutput()
		a := assert.New(t)
		if a.NoError(err, string(resp)) {
			var instance atlasv2.StreamsTenant
			err := json.Unmarshal(resp, &instance)
			a.NoError(err)
			a.Equal(*instance.Name, instanceName)
			a.Equal(instance.DataProcessRegion.CloudProvider, "AWS")
			a.Equal(instance.DataProcessRegion.Region, "VIRGINIA_USA")
		}
	})

	t.Run("Updating a streams instance", func(t *testing.T) {
		cmd := exec.Command(cliPath,
			"streams",
			"instance",
			"update",
			"-r",
			"VIRGINIA_USA",
			instanceName,
			"-o=json",
		)

		cmd.Env = os.Environ()
		resp, err := cmd.CombinedOutput()
		a := assert.New(t)
		if a.NoError(err, string(resp)) {
			var instance atlasv2.StreamsTenant
			err := json.Unmarshal(resp, &instance)
			a.NoError(err)
			a.Equal(*instance.Name, instanceName)
			a.Equal(instance.DataProcessRegion.CloudProvider, "AWS")
			a.Equal(instance.DataProcessRegion.Region, "VIRGINIA_USA")
		}
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
		)

		cmd.Env = os.Environ()
		resp, err := cmd.CombinedOutput()
		a := assert.New(t)
		if a.NoError(err, string(resp)) {
			var connection atlasv2.StreamsConnection
			err := json.Unmarshal(resp, &connection)
			a.NoError(err)
			a.Equal(*connection.Name, connectionName)
		}
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
		)

		cmd.Env = os.Environ()
		resp, err := cmd.CombinedOutput()
		a := assert.New(t)
		if a.NoError(err, string(resp)) {
			var connection atlasv2.StreamsConnection
			err := json.Unmarshal(resp, &connection)
			a.NoError(err)

			a.Equal(connection.Name, connectionName)
			a.Equal(connection.Type, "Kafka")
			a.Equal(connection.BootstrapServers, "example.com:8080,fraud.example.com:8000")
		}
	})

	t.Run("Listing streams connections", func(t *testing.T) {
		cmd := exec.Command(cliPath,
			"streams",
			"connection",
			"list",
			instanceName,
			"-o=json",
		)

		cmd.Env = os.Environ()
		resp, err := cmd.CombinedOutput()
		a := assert.New(t)
		if a.NoError(err, string(resp)) {
			var connections []atlasv2.StreamsConnection
			err := json.Unmarshal(resp, &connections)
			a.NoError(err)

			a.Len(connections, 1)
			a.Equal(connections[0].Name, connectionName)
			a.Equal(connections[0].Type, "Kafka")
			a.Equal(connections[0].BootstrapServers, "example.com:8080,fraud.example.com:8000")
		}
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
		)

		cmd.Env = os.Environ()
		resp, err := cmd.CombinedOutput()
		a := assert.New(t)
		if a.NoError(err, string(resp)) {
			var connection atlasv2.StreamsConnection
			err := json.Unmarshal(resp, &connection)
			a.NoError(err)
			a.Equal(*connection.Name, connectionName)

			a.Equal(*connection.Security.Protocol, "SSL")
		}
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
		)

		cmd.Env = os.Environ()
		resp, err := cmd.CombinedOutput()

		req := require.New(t)
		req.NoError(err, string(resp))

		expected := fmt.Sprintf("'%s' deleted\n", connectionName)
		a := assert.New(t)
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
		)

		cmd.Env = os.Environ()
		resp, err := cmd.CombinedOutput()

		req := require.New(t)
		req.NoError(err, string(resp))

		expected := fmt.Sprintf("Atlas Streams processor instance '%s' deleted\n", instanceName)
		a := assert.New(t)
		a.Equal(expected, string(resp))
	})
}
