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
// +build e2e atlas,ldap

package atlas_test

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"testing"

	"github.com/mongodb/mongocli/e2e"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/atlas/mongodbatlas"
)

const (
	PENDING  = "PENDING"
	HOSTNAME = "2.tcp.ngrok.io"
)

func TestLDAP(t *testing.T) {
	n, err := e2e.RandInt(255)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	cliPath, err := e2e.Bin()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	projectName := fmt.Sprintf("e2e-integration-proj-%v", n)
	projectID, err := createProject(projectName)

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	clusterName, err := deployClusterFromProject(projectID)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	defer func() {
		if e := deleteProject(projectID); e != nil {
			t.Errorf("error deleting project: %v", e)
		}
		if e := deleteClusterForProject(clusterName, projectID); e != nil {
			t.Errorf("error deleting test cluster: %v", e)
		}
	}()

	var requestID string
	t.Run("Verify", func(t *testing.T) {
		cmd := exec.Command(cliPath,
			atlasEntity,
			securityEntity,
			ldapEntity,
			"verify",
			"--hostname",
			HOSTNAME,
			"--port",
			"19657",
			"--bindUsername",
			"cn=admin,dc=example,dc=org",
			"--binPassword",
			"admin",
			"--projectId",
			projectID,
			"-o",
			"json")
		cmd.Env = os.Environ()
		resp, err := cmd.CombinedOutput()

		a := assert.New(t)
		a.NoError(err, string(resp))

		var configuration mongodbatlas.LDAPConfiguration
		if err := json.Unmarshal(resp, &configuration); a.NoError(err) {
			a.Equal(PENDING, configuration.Status)
			requestID = configuration.RequestID
		}
	})

	t.Run("Watch", func(t *testing.T) {
		cmd := exec.Command(cliPath,
			atlasEntity,
			securityEntity,
			ldapEntity,
			"verify",
			"status",
			"watch",
			requestID,
			"--projectId",
			projectID)
		cmd.Env = os.Environ()
		resp, err := cmd.CombinedOutput()

		a := assert.New(t)
		a.NoError(err, string(resp))
		a.Contains(string(resp), "LDAP Configuration request completed.")
	})

	t.Run("Get Status", func(t *testing.T) {
		cmd := exec.Command(cliPath,
			atlasEntity,
			securityEntity,
			ldapEntity,
			"verify",
			"status",
			requestID,
			"--projectId",
			projectID,
			"-o",
			"json")
		cmd.Env = os.Environ()
		resp, err := cmd.CombinedOutput()

		a := assert.New(t)
		a.NoError(err, string(resp))

		var configuration mongodbatlas.LDAPConfiguration
		if err := json.Unmarshal(resp, &configuration); a.NoError(err) {
			a.Equal(requestID, configuration.RequestID)
		}
	})

	t.Run("Save", func(t *testing.T) {
		cmd := exec.Command(cliPath,
			atlasEntity,
			securityEntity,
			ldapEntity,
			"save",
			"--hostname",
			HOSTNAME,
			"--port",
			"19657",
			"--bindUsername",
			"cn=admin,dc=example,dc=org",
			"--binPassword",
			"admin",
			"--projectId",
			projectID,
			"-o",
			"json")
		cmd.Env = os.Environ()
		resp, err := cmd.CombinedOutput()

		a := assert.New(t)
		a.NoError(err, string(resp))

		var configuration mongodbatlas.LDAPConfiguration
		if err := json.Unmarshal(resp, &configuration); a.NoError(err) {
			a.Equal(HOSTNAME, configuration.LDAP.Hostname)
		}
	})

	t.Run("Get", func(t *testing.T) {
		cmd := exec.Command(cliPath,
			atlasEntity,
			securityEntity,
			ldapEntity,
			"get",
			"--projectId",
			projectID,
			"-o",
			"json")
		cmd.Env = os.Environ()
		resp, err := cmd.CombinedOutput()

		a := assert.New(t)
		a.NoError(err, string(resp))

		var configuration mongodbatlas.LDAPConfiguration
		if err := json.Unmarshal(resp, &configuration); a.NoError(err) {
			a.Equal(HOSTNAME, configuration.LDAP.Hostname)
			requestID = configuration.RequestID
		}
	})

	t.Run("Delete", func(t *testing.T) {
		cmd := exec.Command(cliPath,
			atlasEntity,
			securityEntity,
			ldapEntity,
			"delete",
			"--force",
			"--projectId",
			projectID)
		cmd.Env = os.Environ()
		resp, err := cmd.CombinedOutput()

		a := assert.New(t)
		a.NoError(err, string(resp))

		expected := fmt.Sprintf("LDAP configuration userToDNMapping deleted from project'%s'\n", projectID)
		a.Equal(expected, string(resp))
	})
}
