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
// +build e2e atlas,networking

package atlas_test

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"strings"
	"testing"

	"github.com/mongodb/mongocli/e2e"
	"github.com/stretchr/testify/assert"
	atlas "go.mongodb.org/atlas/mongodbatlas"
)

var regions = []string{
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

func TestPrivateEndpointsDeprecated(t *testing.T) {
	n, err := e2e.RandInt(int64(len(regions)))
	a := assert.New(t)
	a.NoError(err)

	cliPath, err := e2e.Bin()
	a.NoError(err)

	region := regions[n.Int64()]
	var id string

	n, err = e2e.RandInt(1000)
	a.NoError(err)

	projectName := fmt.Sprintf("e2e-integration-private-endpoint-deprecated-%v", n)
	projectID, err := createProject(projectName)
	a.NoError(err)

	defer func() {
		if e := deleteProject(projectID); e != nil {
			t.Errorf("error deleting project: %v", e)
		}
	}()

	t.Run("Create", func(t *testing.T) {
		cmd := exec.Command(cliPath,
			atlasEntity,
			privateEndpointsEntity,
			"create",
			"--region",
			region,
			"--projectId",
			projectID,
			"-o=json")
		cmd.Env = os.Environ()
		resp, err := cmd.CombinedOutput()

		newResp := strings.ReplaceAll(string(resp), `Command "create" is deprecated, Please use mongocli atlas privateEndpoints aws create [--region region] [--projectId projectId]`, "")
		a := assert.New(t)
		if a.NoError(err, resp) {
			var r atlas.PrivateEndpointConnectionDeprecated
			if err = json.Unmarshal([]byte(newResp), &r); a.NoError(err) {
				id = r.ID
			}
		}
	})
	if id == "" {
		assert.FailNow(t, "Failed to create alert private endpoint")
	}

	t.Run("Watch", func(t *testing.T) {
		cmd := exec.Command(cliPath,
			atlasEntity,
			privateEndpointsEntity,
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
			"describe",
			id,
			"--projectId",
			projectID,
			"-o=json")
		cmd.Env = os.Environ()
		resp, err := cmd.CombinedOutput()
		a := assert.New(t)
		a.NoError(err, string(resp))
		newResp := strings.ReplaceAll(string(resp), `Command "describe" is deprecated, Please use mongocli atlas privateEndpoints aws describe <ID> [--projectId projectId]`, "")
		var r atlas.PrivateEndpointConnectionDeprecated
		if err = json.Unmarshal([]byte(newResp), &r); a.NoError(err) {
			a.Equal(id, r.ID)
		}
	})

	t.Run("List", func(t *testing.T) {
		cmd := exec.Command(cliPath,
			atlasEntity,
			privateEndpointsEntity,
			"ls",
			"--projectId",
			projectID,
			"-o=json")
		cmd.Env = os.Environ()
		resp, err := cmd.CombinedOutput()

		a := assert.New(t)
		a.NoError(err, string(resp))
		var r []atlas.PrivateEndpointConnectionDeprecated
		newResp := strings.ReplaceAll(string(resp), `Command "list" is deprecated, Please use mongocli atlas privateEndpoints aws list|ls [--projectId projectId]`, "")
		if err = json.Unmarshal([]byte(newResp), &r); a.NoError(err) {
			a.NotEmpty(r)
		}
	})

	t.Run("Delete", func(t *testing.T) {
		cmd := exec.Command(cliPath,
			atlasEntity,
			privateEndpointsEntity,
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
		a.Contains(string(resp), expected)
	})

	t.Run("Watch", func(t *testing.T) {
		cmd := exec.Command(cliPath,
			atlasEntity,
			privateEndpointsEntity,
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
