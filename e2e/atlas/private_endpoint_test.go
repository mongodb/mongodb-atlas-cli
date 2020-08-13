// Copyright 2020 MongoDB Inc
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
// +build e2e atlas,generic

package atlas_test

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
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

func TestPrivateEndpoints(t *testing.T) {
	n, err := e2e.RandInt(int64(len(regions)))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	cliPath, err := e2e.Bin()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	region := regions[n.Int64()]
	var id string

	t.Run("Create", func(t *testing.T) {
		cmd := exec.Command(cliPath,
			atlasEntity,
			privateEndpointsEntity,
			"create",
			"--region="+region,
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
		assert.FailNow(t, "Failed to create alert private endpoint")
	}

	t.Run("Describe", func(t *testing.T) {
		cmd := exec.Command(cliPath,
			atlasEntity,
			privateEndpointsEntity,
			"describe",
			id,
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
			"ls",
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
			"delete",
			id,
			"--force")
		cmd.Env = os.Environ()

		resp, err := cmd.CombinedOutput()
		a := assert.New(t)
		a.NoError(err, string(resp))
		expected := fmt.Sprintf("Private endpoint '%s' deleted\n", id)
		a.Equal(expected, string(resp))
	})
}
