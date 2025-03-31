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

package e2e_test

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	atlasv2 "go.mongodb.org/atlas-sdk/v20250312001/admin"
)

func TestIntegrations(t *testing.T) {
	g := newAtlasE2ETestGenerator(t, withSnapshot())
	g.generateProject("integrations")

	n := g.memoryRand("rand", 255)
	key := "51c0ef87e9951c3e147accf0e12" + n.String()

	cliPath, err := AtlasCLIBin()
	require.NoError(t, err)

	g.Run("Create DATADOG", func(t *testing.T) { //nolint:thelper // g.Run replaces t.Run
		n := g.memoryRand("datadog_rand", 9)
		datadogKey := "000000000000000000000000000000" + n.String() + n.String()
		if IsGov() {
			t.Skip("Skipping DATADOG integration test, cloudgov does not have an available datadog region")
		}
		cmd := exec.Command(cliPath,
			integrationsEntity,
			"create",
			datadogEntity,
			"--apiKey",
			datadogKey,
			"--projectId",
			g.projectID,
			"-o=json")
		cmd.Env = os.Environ()
		resp, err := RunAndGetStdOut(cmd)

		a := assert.New(t)
		require.NoError(t, err, string(resp))
		var thirdPartyIntegrations atlasv2.PaginatedIntegration
		require.NoError(t, json.Unmarshal(resp, &thirdPartyIntegrations))
		a.True(integrationExists(datadogEntity, thirdPartyIntegrations))
	})

	g.Run("Create OPSGENIE", func(t *testing.T) { //nolint:thelper // g.Run replaces t.Run
		n := g.memoryRand("opsgenie_rand", 9)
		opsGenieKey := "00000000-aaaa-2222-bbbb-3333333333" + n.String() + n.String()
		cmd := exec.Command(cliPath,
			integrationsEntity,
			"create",
			opsGenieEntity,
			"--apiKey",
			opsGenieKey,
			"--projectId",
			g.projectID,
			"-o=json")
		cmd.Env = os.Environ()
		resp, err := RunAndGetStdOut(cmd)

		a := assert.New(t)
		require.NoError(t, err, string(resp))

		var thirdPartyIntegrations atlasv2.PaginatedIntegration
		require.NoError(t, json.Unmarshal(resp, &thirdPartyIntegrations))
		a.True(integrationExists(opsGenieEntity, thirdPartyIntegrations))
	})

	g.Run("Create PAGER_DUTY", func(t *testing.T) { //nolint:thelper // g.Run replaces t.Run
		n := g.memoryRand("pager_duty_rand", 9)
		pagerDutyKey := "000000000000000000000000000000" + n.String() + n.String()
		cmd := exec.Command(cliPath,
			integrationsEntity,
			"create",
			pagerDutyEntity,
			"--serviceKey",
			pagerDutyKey,
			"--projectId",
			g.projectID,
			"-o=json")
		cmd.Env = os.Environ()
		resp, err := RunAndGetStdOut(cmd)

		a := assert.New(t)
		require.NoError(t, err, string(resp))

		var thirdPartyIntegrations atlasv2.PaginatedIntegration
		require.NoError(t, json.Unmarshal(resp, &thirdPartyIntegrations))
		a.True(integrationExists(pagerDutyEntity, thirdPartyIntegrations))
	})

	g.Run("Create VICTOR_OPS", func(t *testing.T) { //nolint:thelper // g.Run replaces t.Run
		n := g.memoryRand("victor_ops_rand", 9)
		victorOpsKey := "fa07bbc8-eab2-4085-81af-daed47dc1c" + n.String() + n.String()
		cmd := exec.Command(cliPath,
			integrationsEntity,
			"create",
			victorOpsEntity,
			"--apiKey",
			victorOpsKey,
			"--routingKey",
			"test",
			"--projectId",
			g.projectID,
			"-o=json")
		cmd.Env = os.Environ()
		resp, err := RunAndGetStdOut(cmd)

		a := assert.New(t)
		require.NoError(t, err, string(resp))

		var thirdPartyIntegrations atlasv2.PaginatedIntegration
		require.NoError(t, json.Unmarshal(resp, &thirdPartyIntegrations))
		a.True(integrationExists(victorOpsEntity, thirdPartyIntegrations))
	})

	g.Run("Create WEBHOOK", func(t *testing.T) { //nolint:thelper // g.Run replaces t.Run
		cmd := exec.Command(cliPath,
			integrationsEntity,
			"create",
			webhookEntity,
			"--url",
			"https://example.com/"+key,
			"--secret",
			key,
			"--projectId",
			g.projectID,
			"-o=json")
		cmd.Env = os.Environ()
		resp, err := RunAndGetStdOut(cmd)

		a := assert.New(t)
		require.NoError(t, err, string(resp))

		var thirdPartyIntegrations atlasv2.PaginatedIntegration
		require.NoError(t, json.Unmarshal(resp, &thirdPartyIntegrations))
		a.True(integrationExists(webhookEntity, thirdPartyIntegrations))
	})

	g.Run("List", func(t *testing.T) { //nolint:thelper // g.Run replaces t.Run
		cmd := exec.Command(cliPath,
			integrationsEntity,
			"ls",
			"--projectId",
			g.projectID,
			"-o=json")
		cmd.Env = os.Environ()
		resp, err := RunAndGetStdOut(cmd)

		a := assert.New(t)
		require.NoError(t, err, string(resp))
		var thirdPartyIntegrations atlasv2.PaginatedIntegration
		require.NoError(t, json.Unmarshal(resp, &thirdPartyIntegrations))
		a.NotEmpty(thirdPartyIntegrations.Results)
	})

	g.Run("Describe", func(t *testing.T) { //nolint:thelper // g.Run replaces t.Run
		cmd := exec.Command(cliPath,
			integrationsEntity,
			"describe",
			webhookEntity,
			"--projectId",
			g.projectID,
			"-o=json")
		cmd.Env = os.Environ()
		resp, err := RunAndGetStdOut(cmd)

		a := assert.New(t)
		require.NoError(t, err, string(resp))
		var thirdPartyIntegration atlasv2.ThirdPartyIntegration
		require.NoError(t, json.Unmarshal(resp, &thirdPartyIntegration))
		a.Equal(webhookEntity, thirdPartyIntegration.GetType())
	})

	g.Run("Delete", func(t *testing.T) { //nolint:thelper // g.Run replaces t.Run
		cmd := exec.Command(cliPath,
			integrationsEntity,
			"delete",
			webhookEntity,
			"--force",
			"--projectId",
			g.projectID)
		cmd.Env = os.Environ()
		resp, err := RunAndGetStdOut(cmd)

		a := assert.New(t)
		require.NoError(t, err, string(resp))
		expected := fmt.Sprintf("Integration '%s' deleted\n", webhookEntity)
		a.Equal(expected, string(resp))
	})
}
