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

const (
	pending      = "PENDING"
	ldapHostname = "localhost"
	ldapPort     = "19657"
)

func TestLDAP(t *testing.T) {
	g := newAtlasE2ETestGenerator(t)
	g.generateProjectAndCluster("ldap")

	cliPath, err := e2e.Bin()
	require.NoError(t, err)

	var requestID string
	t.Run("Verify", func(t *testing.T) {
		cmd := exec.Command(cliPath,
			atlasEntity,
			securityEntity,
			ldapEntity,
			"verify",
			"--hostname",
			ldapHostname,
			"--port",
			ldapPort,
			"--bindUsername",
			"cn=admin,dc=example,dc=org",
			"--bindPassword",
			"admin",
			"--projectId", g.projectID,
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
			requestID,
			"--projectId", g.projectID,
		)
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
			"--projectId", g.projectID,
			"-o",
			"json",
		)
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
			ldapHostname,
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
			"--projectId", g.projectID,
			"-o",
			"json",
		)
		cmd.Env = os.Environ()
		resp, err := cmd.CombinedOutput()
		require.NoError(t, err, string(resp))

		a := assert.New(t)
		var configuration mongodbatlas.LDAPConfiguration
		if err := json.Unmarshal(resp, &configuration); a.NoError(err) {
			a.Equal(ldapHostname, configuration.LDAP.Hostname)
		}
	})

	t.Run("Get", func(t *testing.T) {
		cmd := exec.Command(cliPath,
			atlasEntity,
			securityEntity,
			ldapEntity,
			"get",
			"--projectId", g.projectID,
			"-o",
			"json")
		cmd.Env = os.Environ()
		resp, err := cmd.CombinedOutput()
		require.NoError(t, err, string(resp))

		a := assert.New(t)
		var configuration mongodbatlas.LDAPConfiguration
		if err := json.Unmarshal(resp, &configuration); a.NoError(err) {
			a.Equal(ldapHostname, configuration.LDAP.Hostname)
			requestID = configuration.RequestID
		}
	})

	t.Run("Delete", func(t *testing.T) {
		cmd := exec.Command(cliPath,
			atlasEntity,
			securityEntity,
			ldapEntity,
			"delete",
			"--projectId", g.projectID,
			"--force")
		cmd.Env = os.Environ()
		resp, err := cmd.CombinedOutput()
		require.NoError(t, err, string(resp))
		assert.Contains(t, string(resp), "LDAP configuration userToDNMapping deleted")
	})
}
