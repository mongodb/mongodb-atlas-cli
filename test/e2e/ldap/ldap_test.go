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
//go:build e2e || e2eSnap || (atlas && ldap)

package ldap

import (
	"bytes"
	"encoding/json"
	"os"
	"os/exec"
	"testing"

	"github.com/mongodb/mongodb-atlas-cli/atlascli/test/internal"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	atlasv2 "go.mongodb.org/atlas-sdk/v20250312004/admin"
)

const (
	securityEntity = "security"
	ldapEntity     = "ldap"
)

const (
	pending          = "PENDING"
	ldapHostname     = "localhost"
	ldapPort         = "19657"
	ldapBindPassword = "admin"
)

func TestLDAPWithFlags(t *testing.T) {
	g := internal.NewAtlasE2ETestGenerator(t, internal.WithSnapshot())
	g.GenerateProjectAndCluster("ldap")

	cliPath, err := internal.AtlasCLIBin()
	require.NoError(t, err)

	var requestID string
	g.Run("Verify", func(t *testing.T) { //nolint:thelper // g.Run replaces t.Run
		cmd := exec.Command(cliPath,
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
			ldapBindPassword,
			"--projectId", g.ProjectID,
			"-o",
			"json")

		requestID = testLDAPVerifyCmd(t, cmd)
	})

	require.NotEmpty(t, requestID)

	g.Run("Watch", func(t *testing.T) { //nolint:thelper // g.Run replaces t.Run
		cmd := exec.Command(cliPath,
			securityEntity,
			ldapEntity,
			"verify",
			"status",
			"watch",
			requestID,
			"--projectId", g.ProjectID,
		)
		cmd.Env = os.Environ()
		resp, err := internal.RunAndGetStdOut(cmd)
		require.NoError(t, err, string(resp))
		assert.Contains(t, string(resp), "LDAP Configuration request completed.")
	})

	g.Run("Get Status", func(t *testing.T) { //nolint:thelper // g.Run replaces t.Run
		cmd := exec.Command(cliPath,
			securityEntity,
			ldapEntity,
			"verify",
			"status",
			requestID,
			"--projectId", g.ProjectID,
			"-o",
			"json",
		)
		cmd.Env = os.Environ()
		resp, err := internal.RunAndGetStdOut(cmd)
		require.NoError(t, err, string(resp))

		a := assert.New(t)
		var configuration atlasv2.LDAPVerifyConnectivityJobRequest
		require.NoError(t, json.Unmarshal(resp, &configuration))
		a.Equal(requestID, *configuration.RequestId)
	})

	g.Run("Save", func(t *testing.T) { //nolint:thelper // g.Run replaces t.Run
		cmd := exec.Command(cliPath,
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
			ldapBindPassword,
			"--mappingMatch",
			"(.+)@ENGINEERING.EXAMPLE.COM",
			"--mappingSubstitution",
			"cn={0},ou=engineering,dc=example,dc=com",
			"--projectId", g.ProjectID,
			"-o",
			"json",
		)

		testLDAPSaveCmd(t, cmd)
	})

	g.Run("Get", func(t *testing.T) { //nolint:thelper // g.Run replaces t.Run
		cmd := exec.Command(cliPath,
			securityEntity,
			ldapEntity,
			"get",
			"--projectId", g.ProjectID,
			"-o",
			"json")
		cmd.Env = os.Environ()
		resp, err := internal.RunAndGetStdOut(cmd)
		require.NoError(t, err, string(resp))

		a := assert.New(t)
		var configuration atlasv2.UserSecurity
		require.NoError(t, json.Unmarshal(resp, &configuration))
		a.Equal(ldapHostname, *configuration.Ldap.Hostname)
	})

	g.Run("Delete", func(t *testing.T) { //nolint:thelper // g.Run replaces t.Run
		testLDAPDelete(t, cliPath, g.ProjectID)
	})
}

func TestLDAPWithStdin(t *testing.T) {
	g := internal.NewAtlasE2ETestGenerator(t, internal.WithSnapshot())
	g.GenerateProjectAndCluster("ldap")

	cliPath, err := internal.AtlasCLIBin()
	require.NoError(t, err)

	var requestID string

	g.Run("Verify", func(t *testing.T) { //nolint:thelper // g.Run replaces t.Run
		cmd := exec.Command(cliPath,
			securityEntity,
			ldapEntity,
			"verify",
			"--hostname",
			ldapHostname,
			"--port",
			ldapPort,
			"--bindUsername",
			"cn=admin,dc=example,dc=org",
			"--projectId", g.ProjectID,
			"-o",
			"json")

		passwordStdin := bytes.NewBuffer([]byte(ldapBindPassword))
		cmd.Stdin = passwordStdin

		requestID = testLDAPVerifyCmd(t, cmd)
	})

	require.NotEmpty(t, requestID)

	g.Run("Save", func(t *testing.T) { //nolint:thelper // g.Run replaces t.Run
		cmd := exec.Command(cliPath,
			securityEntity,
			ldapEntity,
			"save",
			"--hostname",
			ldapHostname,
			"--port",
			ldapPort,
			"--bindUsername",
			"cn=admin,dc=example,dc=org",
			"--mappingMatch",
			"(.+)@ENGINEERING.EXAMPLE.COM",
			"--mappingSubstitution",
			"cn={0},ou=engineering,dc=example,dc=com",
			"--projectId", g.ProjectID,
			"-o",
			"json",
		)

		passwordStdin := bytes.NewBuffer([]byte(ldapBindPassword))
		cmd.Stdin = passwordStdin

		testLDAPSaveCmd(t, cmd)
	})

	g.Run("Delete", func(t *testing.T) { //nolint:thelper // g.Run replaces t.Run
		testLDAPDelete(t, cliPath, g.ProjectID)
	})
}

func testLDAPDelete(t *testing.T, cliPath, projectID string) {
	t.Helper()

	cmd := exec.Command(cliPath,
		securityEntity,
		ldapEntity,
		"delete",
		"--projectId", projectID,
		"--force")
	cmd.Env = os.Environ()
	resp, err := internal.RunAndGetStdOut(cmd)
	require.NoError(t, err, string(resp))
	assert.Contains(t, string(resp), "LDAP configuration userToDNMapping deleted")
}

func testLDAPVerifyCmd(t *testing.T, cmd *exec.Cmd) string {
	t.Helper()

	cmd.Env = os.Environ()
	resp, err := internal.RunAndGetStdOut(cmd)
	require.NoError(t, err, string(resp))

	a := assert.New(t)
	var configuration atlasv2.LDAPVerifyConnectivityJobRequest
	require.NoError(t, json.Unmarshal(resp, &configuration))
	a.Equal(pending, *configuration.Status)
	return *configuration.RequestId
}

func testLDAPSaveCmd(t *testing.T, cmd *exec.Cmd) {
	t.Helper()

	cmd.Env = os.Environ()
	resp, err := internal.RunAndGetStdOut(cmd)
	require.NoError(t, err, string(resp))

	a := assert.New(t)
	var configuration atlasv2.UserSecurity
	require.NoError(t, json.Unmarshal(resp, &configuration))
	a.Equal(ldapHostname, *configuration.Ldap.Hostname)
}
