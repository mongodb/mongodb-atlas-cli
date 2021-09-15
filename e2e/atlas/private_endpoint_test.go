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

//go:build e2e || (atlas && networking)
// +build e2e atlas,networking

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
	atlas "go.mongodb.org/atlas/mongodbatlas"
)

var regionsAWS = []string{
	"us-east-1",
	"us-east-2",
	"us-west-1",
	"us-west-2",
	"ca-central-1",
	"sa-east-1",
	"eu-north-1",
	"eu-west-1",
	"eu-west-2",
	"eu-west-3",
	"eu-central-1",
	"me-south-1",
	"ap-northeast-1",
	"ap-northeast-2",
	"ap-south-1",
	"ap-southeast-1",
	"ap-southeast-2",
	"ap-east-1",
}

func TestPrivateEndpointsAWS(t *testing.T) {
	n, err := e2e.RandInt(int64(len(regionsAWS)))
	require.NoError(t, err)

	cliPath, err := e2e.Bin()
	require.NoError(t, err)

	region := regionsAWS[n.Int64()]
	var id string

	projectName, err := RandProjectNameWithPrefix("private-endpoint-aws")
	require.NoError(t, err)
	projectID, err := createProject(projectName)
	require.NoError(t, err)

	defer func() {
		if e := deleteProject(projectID); e != nil {
			t.Errorf("error deleting project: %v", e)
		}
	}()

	t.Run("Create", func(t *testing.T) {
		cmd := exec.Command(cliPath,
			atlasEntity,
			privateEndpointsEntity,
			awsEntity,
			"create",
			"--region="+region,
			"--projectId",
			projectID,
			"-o=json")
		cmd.Env = os.Environ()

		a := assert.New(t)
		if resp, err := cmd.CombinedOutput(); a.NoError(err, string(resp)) {
			var r atlas.PrivateEndpointConnection
			if err = json.Unmarshal(resp, &r); a.NoError(err) {
				id = r.ID
			}
		}
	})
	require.NotEmpty(t, id)

	t.Run("Watch", func(t *testing.T) {
		cmd := exec.Command(cliPath,
			atlasEntity,
			privateEndpointsEntity,
			awsEntity,
			"watch",
			id,
			"--projectId",
			projectID)
		cmd.Env = os.Environ()

		resp, err := cmd.CombinedOutput()
		assert.NoError(t, err, string(resp))
	})

	t.Run("Describe", func(t *testing.T) {
		cmd := exec.Command(cliPath,
			atlasEntity,
			privateEndpointsEntity,
			awsEntity,
			"describe",
			id,
			"--projectId",
			projectID,
			"-o=json")
		cmd.Env = os.Environ()
		resp, err := cmd.CombinedOutput()
		require.NoError(t, err, string(resp))
		a := assert.New(t)
		var r atlas.PrivateEndpointConnection
		if err = json.Unmarshal(resp, &r); a.NoError(err) {
			a.Equal(id, r.ID)
		}
	})

	t.Run("List", func(t *testing.T) {
		cmd := exec.Command(cliPath,
			atlasEntity,
			privateEndpointsEntity,
			awsEntity,
			"ls",
			"--projectId",
			projectID,
			"-o=json")
		cmd.Env = os.Environ()
		resp, err := cmd.CombinedOutput()

		a := assert.New(t)
		a.NoError(err, string(resp))
		var r []atlas.PrivateEndpointConnection
		if err = json.Unmarshal(resp, &r); a.NoError(err) {
			a.NotEmpty(r)
		}
	})

	t.Run("Delete", func(t *testing.T) {
		cmd := exec.Command(cliPath,
			atlasEntity,
			privateEndpointsEntity,
			awsEntity,
			"delete",
			id,
			"--projectId",
			projectID,
			"--force")
		cmd.Env = os.Environ()

		resp, err := cmd.CombinedOutput()
		a := assert.New(t)
		a.NoError(err, string(resp))
		expected := fmt.Sprintf("Private endpoint '%s' deleted\n", id)
		a.Equal(expected, string(resp))
	})

	t.Run("Watch", func(t *testing.T) {
		cmd := exec.Command(cliPath,
			atlasEntity,
			privateEndpointsEntity,
			awsEntity,
			"watch",
			id,
			"--projectId",
			projectID)
		cmd.Env = os.Environ()

		resp, err := cmd.CombinedOutput()
		a := assert.New(t)
		// We expect a 404 error once the private endpoint has been completely deleted
		a.Error(err)
		a.Contains(string(resp), "404")
	})
}

var regionsAzure = []string{
	"US_EAST_2",
	"EUROPE_NORTH",
	"US_WEST_2",
	"ASIA_SOUTH_EAST",
}

func TestPrivateEndpointsAzure(t *testing.T) {
	n, err := e2e.RandInt(int64(len(regionsAzure)))
	require.NoError(t, err)

	cliPath, err := e2e.Bin()
	require.NoError(t, err)

	region := regionsAzure[n.Int64()]
	var id string

	projectName, err := RandProjectNameWithPrefix("private-endpoint-azure")
	require.NoError(t, err)
	projectID, err := createProject(projectName)
	require.NoError(t, err)

	defer func() {
		if e := deleteProject(projectID); e != nil {
			t.Errorf("error deleting project: %v", e)
		}
	}()

	t.Run("Create", func(t *testing.T) {
		cmd := exec.Command(cliPath,
			atlasEntity,
			privateEndpointsEntity,
			azureEntity,
			"create",
			"--region="+region,
			"--projectId",
			projectID,
			"-o=json")
		cmd.Env = os.Environ()

		a := assert.New(t)
		if resp, err := cmd.CombinedOutput(); a.NoError(err, string(resp)) {
			var r atlas.PrivateEndpointConnection
			if err = json.Unmarshal(resp, &r); a.NoError(err) {
				id = r.ID
			}
		}
	})
	if id == "" {
		assert.FailNow(t, "Failed to create private endpoint")
	}

	t.Run("Watch", func(t *testing.T) {
		cmd := exec.Command(cliPath,
			atlasEntity,
			privateEndpointsEntity,
			azureEntity,
			"watch",
			id,
			"--projectId",
			projectID)
		cmd.Env = os.Environ()

		_, err := cmd.CombinedOutput()
		a := assert.New(t)
		a.NoError(err)
	})

	t.Run("Describe", func(t *testing.T) {
		cmd := exec.Command(cliPath,
			atlasEntity,
			privateEndpointsEntity,
			azureEntity,
			"describe",
			id,
			"--projectId",
			projectID,
			"-o=json")
		cmd.Env = os.Environ()
		resp, err := cmd.CombinedOutput()
		a := assert.New(t)
		a.NoError(err, string(resp))
		var r atlas.PrivateEndpointConnection
		if err = json.Unmarshal(resp, &r); a.NoError(err) {
			a.Equal(id, r.ID)
		}
	})

	t.Run("List", func(t *testing.T) {
		cmd := exec.Command(cliPath,
			atlasEntity,
			privateEndpointsEntity,
			azureEntity,
			"ls",
			"--projectId",
			projectID,
			"-o=json")
		cmd.Env = os.Environ()
		resp, err := cmd.CombinedOutput()

		a := assert.New(t)
		a.NoError(err, string(resp))
		var r []atlas.PrivateEndpointConnection
		if err = json.Unmarshal(resp, &r); a.NoError(err) {
			a.NotEmpty(r)
		}
	})

	t.Run("Delete", func(t *testing.T) {
		cmd := exec.Command(cliPath,
			atlasEntity,
			privateEndpointsEntity,
			azureEntity,
			"delete",
			id,
			"--force",
			"--projectId",
			projectID)
		cmd.Env = os.Environ()

		resp, err := cmd.CombinedOutput()
		a := assert.New(t)
		a.NoError(err, string(resp))
		expected := fmt.Sprintf("Private endpoint '%s' deleted\n", id)
		a.Equal(expected, string(resp))
	})

	t.Run("Watch", func(t *testing.T) {
		cmd := exec.Command(cliPath,
			atlasEntity,
			privateEndpointsEntity,
			azureEntity,
			"watch",
			id,
			"--projectId",
			projectID)
		cmd.Env = os.Environ()

		resp, err := cmd.CombinedOutput()
		a := assert.New(t)
		// We expect a 404 error once the private endpoint has been completely deleted
		a.Error(err)
		a.Contains(string(resp), "404")
	})
}

func TestRegionalizedPrivateEndpointsSettings(t *testing.T) {
	cliPath, err := e2e.Bin()
	require.NoError(t, err)

	projectName, err := RandProjectNameWithPrefix("regionalized-private-endpoint-setting")
	require.NoError(t, err)
	projectID, err := createProject(projectName)
	require.NoError(t, err)

	defer func() {
		if e := deleteProject(projectID); e != nil {
			t.Errorf("error deleting project: %v", e)
		}
	}()

	t.Run("Enable regionalized private endpoint setting", func(t *testing.T) {
		cmd := exec.Command(cliPath,
			atlasEntity,
			privateEndpointsEntity,
			regionalModeEntity,
			"enable",
			"--projectId",
			projectID)
		cmd.Env = os.Environ()

		resp, err := cmd.CombinedOutput()
		a := assert.New(t)
		a.NoError(err, string(resp))
		a.Equal("Regionalized private endpoint setting enabled.\n", string(resp))
	})

	t.Run("Disable regionalized private endpoint setting", func(t *testing.T) {
		cmd := exec.Command(cliPath,
			atlasEntity,
			privateEndpointsEntity,
			regionalModeEntity,
			"disable",
			"--projectId",
			projectID)
		cmd.Env = os.Environ()

		resp, err := cmd.CombinedOutput()
		a := assert.New(t)
		a.NoError(err, string(resp))
		a.Equal("Regionalized private endpoint setting disabled.\n", string(resp))
	})

	t.Run("Get regionalized private endpoint setting", func(t *testing.T) {
		cmd := exec.Command(cliPath,
			atlasEntity,
			privateEndpointsEntity,
			regionalModeEntity,
			"get",
			"--projectId",
			projectID,
			"-o=json")
		cmd.Env = os.Environ()

		resp, err := cmd.CombinedOutput()
		a := assert.New(t)
		a.NoError(err, string(resp))
		var r atlas.RegionalizedPrivateEndpointSetting
		if err = json.Unmarshal(resp, &r); a.NoError(err) {
			a.False(r.Enabled)
		}
	})
}
