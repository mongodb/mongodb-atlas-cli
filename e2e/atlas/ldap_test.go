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

//go:build e2e || (atlas && ldap)
// +build e2e atlas,ldap

package atlas_test

import (
	"encoding/json"
	"os"
	"os/exec"
	"testing"
	"time"

	"github.com/mongodb/mongocli/e2e"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.mongodb.org/atlas/mongodbatlas"
)

const (
	pending  = "PENDING"
	hostname = "localhost"
	ldapPort = "19657"
)

func TestLDAP(t *testing.T) {
	cliPath, err := e2e.Bin()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	clusterName, err := deployCluster()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	defer func() {
		if e := deleteCluster(clusterName); e != nil {
			t.Errorf("error deleting test cluster: %v", e)
		}
	}()
	time.Sleep(2 * time.Minute) // wait for the cluster to be fully available
	var requestID string
	t.Run("Verify", func(t *testing.T) {
		cmd := exec.Command(cliPath,
			atlasEntity,
			securityEntity,
			ldapEntity,
			"verify",
			"--hostname",
			hostname,
			"--port",
			ldapPort,
			"--bindUsername",
			"cn=admin,dc=example,dc=org",
			"--bindPassword",
			"admin",
			"-o",
			"json")
		cmd.Env = os.Environ()
		resp, err := cmd.CombinedOutput()
		require.NoError(t, err, string(resp))

		a := assert.New(t)
		var configuration mongodbatlas.LDAPConfiguration
		if err := json.Unmarshal(resp, &configuration); a.NoError(err) {
			a.Equal(pending, configuration.Status)
			requestID = configuration.RequestID
		}
	})

	require.NotEmpty(t, requestID)

	t.Run("Watch", func(t *testing.T) {
		cmd := exec.Command(cliPath,
			atlasEntity,
			securityEntity,
			ldapEntity,
			"verify",
			"status",
			"watch",
			requestID)
		cmd.Env = os.Environ()
		resp, err := cmd.CombinedOutput()
		require.NoError(t, err, string(resp))
		assert.Contains(t, string(resp), "LDAP Configuration request completed.")
	})

	t.Run("Get Status", func(t *testing.T) {
		cmd := exec.Command(cliPath,
			atlasEntity,
			securityEntity,
			ldapEntity,
			"verify",
			"status",
			requestID,
			"-o",
			"json")
		cmd.Env = os.Environ()
		resp, err := cmd.CombinedOutput()
		require.NoError(t, err, string(resp))

		a := assert.New(t)
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
			hostname,
			"--port",
			ldapPort,
			"--bindUsername",
			"cn=admin,dc=example,dc=org",
			"--bindPassword",
			"admin",
			"--mappingMatch",
			"(.+)@ENGINEERING.EXAMPLE.COM",
			"--mappingSubstitution",
			"cn={0},ou=engineering,dc=example,dc=com",
			"-o",
			"json")
		cmd.Env = os.Environ()
		resp, err := cmd.CombinedOutput()
		require.NoError(t, err, string(resp))

		a := assert.New(t)
		var configuration mongodbatlas.LDAPConfiguration
		if err := json.Unmarshal(resp, &configuration); a.NoError(err) {
			a.Equal(hostname, configuration.LDAP.Hostname)
		}
	})

	t.Run("Get", func(t *testing.T) {
		cmd := exec.Command(cliPath,
			atlasEntity,
			securityEntity,
			ldapEntity,
			"get",
			"-o",
			"json")
		cmd.Env = os.Environ()
		resp, err := cmd.CombinedOutput()
		require.NoError(t, err, string(resp))

		a := assert.New(t)
		var configuration mongodbatlas.LDAPConfiguration
		if err := json.Unmarshal(resp, &configuration); a.NoError(err) {
			a.Equal(hostname, configuration.LDAP.Hostname)
			requestID = configuration.RequestID
		}
	})

	t.Run("Delete", func(t *testing.T) {
		cmd := exec.Command(cliPath,
			atlasEntity,
			securityEntity,
			ldapEntity,
			"delete",
			"--force")
		cmd.Env = os.Environ()
		resp, err := cmd.CombinedOutput()
		require.NoError(t, err, string(resp))
		assert.Contains(t, string(resp), "LDAP configuration userToDNMapping deleted")
	})
}
