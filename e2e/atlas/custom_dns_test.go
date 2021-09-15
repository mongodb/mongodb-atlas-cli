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

//go:build e2e || (atlas && generic)
// +build e2e atlas,generic

package atlas_test

import (
	"encoding/json"
	"os"
	"os/exec"
	"testing"

	"github.com/mongodb/mongocli/e2e"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.mongodb.org/atlas/mongodbatlas"
)

func TestCustomDNS(t *testing.T) {
	cliPath, err := e2e.Bin()
	require.NoError(t, err)
	projectName, err := RandProjectNameWithPrefix("integration-custom-dns")
	require.NoError(t, err)
	projectID, err := createProject(projectName)
	require.NoError(t, err)

	defer func() {
		if e := deleteProject(projectID); e != nil {
			t.Errorf("error deleting project: %v", e)
		}
	}()

	t.Run("Enable", func(t *testing.T) {
		cmd := exec.Command(cliPath,
			atlasEntity,
			customDNSEntity,
			awsEntity,
			"enable",
			"--projectId",
			projectID,
			"-o=json")
		cmd.Env = os.Environ()
		resp, err := cmd.CombinedOutput()

		a := assert.New(t)
		a.NoError(err, string(resp))

		var dns mongodbatlas.AWSCustomDNSSetting
		if err := json.Unmarshal(resp, &dns); a.NoError(err) {
			a.True(dns.Enabled)
		}
	})

	t.Run("Describe", func(t *testing.T) {
		cmd := exec.Command(cliPath,
			atlasEntity,
			customDNSEntity,
			awsEntity,
			"describe",
			"--projectId",
			projectID,
			"-o=json")
		cmd.Env = os.Environ()
		resp, err := cmd.CombinedOutput()

		a := assert.New(t)
		a.NoError(err, string(resp))

		var dns mongodbatlas.AWSCustomDNSSetting
		if err := json.Unmarshal(resp, &dns); a.NoError(err) {
			a.True(dns.Enabled)
		}
	})

	t.Run("Disable", func(t *testing.T) {
		cmd := exec.Command(cliPath,
			atlasEntity,
			customDNSEntity,
			awsEntity,
			"disable",
			"--projectId",
			projectID,
			"-o=json")
		cmd.Env = os.Environ()
		resp, err := cmd.CombinedOutput()

		a := assert.New(t)
		a.NoError(err, string(resp))

		var dns mongodbatlas.AWSCustomDNSSetting
		if err := json.Unmarshal(resp, &dns); a.NoError(err) {
			a.False(dns.Enabled)
		}
	})
}
