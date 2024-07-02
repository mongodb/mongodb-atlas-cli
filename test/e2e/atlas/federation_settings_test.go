// Copyright 2024 MongoDB Inc
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
//go:build e2e || (iam && atlas)

package atlas_test

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"testing"

	"github.com/mongodb/mongodb-atlas-cli/atlascli/test/e2e"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	atlasv2 "go.mongodb.org/atlas-sdk/v20240530002/admin"
)

func TestIdentityProviders(t *testing.T) {
	req := require.New(t)

	cliPath, err := e2e.AtlasCLIBin()
	req.NoError(err)

	var federationSettingsID string
	var oidcWorkloadIdpID string
	var oidcIWorkforceIdpID string

	t.Run("Describe an org federation settings", func(t *testing.T) {
		cmd := exec.Command(cliPath,
			federatedAuthenticationEntity,
			federationSettingsEntity,
			"describe",
			"-o=json",
		)

		cmd.Env = os.Environ()
		resp, err := e2e.RunAndGetStdOut(cmd)
		req.NoError(err, string(resp))

		var settings atlasv2.OrgFederationSettings
		req.NoError(json.Unmarshal(resp, &settings))

		a := assert.New(t)
		a.NotEmpty(settings.GetId())
		a.NotEmpty(settings.GetIdentityProviderStatus())
		federationSettingsID = settings.GetId()
	})

	t.Run("Create OIDC IdP WORKLOAD", func(t *testing.T) {
		idpName, err := RandIdentityProviderName()
		req.NoError(err)

		cmd := exec.Command(cliPath,
			federatedAuthenticationEntity,
			federationSettingsEntity,
			identityProviderEntity,
			"create",
			"oidc",
			idpName,
			"--federationSettingsId",
			federationSettingsID,
			"--audience",
			idpName, // using random as audience also should be unique
			"--authorizationType",
			"GROUP",
			"--desc",
			"CLI TEST Provider",
			"--groupsClaim",
			"groups",
			"--idpType",
			"WORKLOAD",
			"--issuerUri",
			"https://accounts.google.com",
			"--userClaim",
			"user",
			"-o=json",
		)

		cmd.Env = os.Environ()
		resp, err := e2e.RunAndGetStdOut(cmd)
		req.NoError(err, string(resp))

		var provider atlasv2.FederationIdentityProvider
		req.NoError(json.Unmarshal(resp, &provider))

		assert.NotEmpty(t, provider.GetId())
		oidcWorkloadIdpID = provider.GetId()
	})

	t.Run("Connect OIDC IdP WORKLOAD", func(t *testing.T) {
		cmd := exec.Command(cliPath,
			federatedAuthenticationEntity,
			federationSettingsEntity,
			connectedOrgsConfigsEntity,
			"connect",
			"--identityProviderId",
			oidcWorkloadIdpID,
			"--federationSettingsId",
			federationSettingsID,
			"-o=json",
		)

		cmd.Env = os.Environ()
		resp, err := e2e.RunAndGetStdOut(cmd)
		req.NoError(err, string(resp))

		var config atlasv2.ConnectedOrgConfig
		req.NoError(json.Unmarshal(resp, &config))

		assert.NotEmpty(t, config.DataAccessIdentityProviderIds)
		assert.Contains(t, config.GetDataAccessIdentityProviderIds(), oidcWorkloadIdpID)
	})

	t.Run("Create OIDC IdP WORKFORCE", func(t *testing.T) {
		idpName, err := RandIdentityProviderName()
		fmt.Println(idpName)
		req.NoError(err)

		cmd := exec.Command(cliPath,
			federatedAuthenticationEntity,
			federationSettingsEntity,
			identityProviderEntity,
			"create",
			"oidc",
			idpName,
			"--federationSettingsId",
			federationSettingsID,
			"--audience",
			idpName, // using random as audience also should be unique
			"--authorizationType",
			"GROUP",
			"--clientId",
			"cliClients",
			"--desc",
			"CLI TEST Provider",
			"--groupsClaim",
			"groups",
			"--idpType",
			"WORKFORCE",
			"--issuerUri",
			"https://accounts.google.com",
			"--userClaim",
			"user",
			"--associatedDomain",
			"iam-test-domain-dev.com",
			"-o=json",
		)

		cmd.Env = os.Environ()
		resp, err := e2e.RunAndGetStdOut(cmd)
		req.NoError(err, string(resp))

		var provider atlasv2.FederationIdentityProvider
		req.NoError(json.Unmarshal(resp, &provider))

		assert.NotEmpty(t, provider.GetId())
		oidcIWorkforceIdpID = provider.Id
	})

	t.Run("Describe OIDC IdP WORKFORCE", func(t *testing.T) {
		cmd := exec.Command(cliPath,
			federatedAuthenticationEntity,
			federationSettingsEntity,
			identityProviderEntity,
			"describe",
			oidcIWorkforceIdpID,
			"--federationSettingsId",
			federationSettingsID,
			"-o=json",
		)

		cmd.Env = os.Environ()
		resp, err := e2e.RunAndGetStdOut(cmd)
		req.NoError(err, string(resp))

		var provider atlasv2.FederationIdentityProvider
		req.NoError(json.Unmarshal(resp, &provider))

		assert.NotEmpty(t, provider.GetId())
	})

	t.Run("Connect OIDC IdP WORKFORCE", func(t *testing.T) {
		cmd := exec.Command(cliPath,
			federatedAuthenticationEntity,
			federationSettingsEntity,
			connectedOrgsConfigsEntity,
			"connect",
			"--identityProviderId",
			oidcIWorkforceIdpID,
			"--federationSettingsId",
			federationSettingsID,
			"-o=json",
		)

		cmd.Env = os.Environ()
		resp, err := e2e.RunAndGetStdOut(cmd)
		req.NoError(err, string(resp))

		var config atlasv2.ConnectedOrgConfig
		req.NoError(json.Unmarshal(resp, &config))

		assert.NotEmpty(t, config.DataAccessIdentityProviderIds)
		assert.Contains(t, config.GetDataAccessIdentityProviderIds(), oidcIWorkforceIdpID)
	})

	t.Run("Describe connectedOrgsConfig", func(t *testing.T) {
		cmd := exec.Command(cliPath,
			federatedAuthenticationEntity,
			federationSettingsEntity,
			connectedOrgsConfigsEntity,
			"describe",
			"--federationSettingsId",
			federationSettingsID,
			"-o=json",
		)

		cmd.Env = os.Environ()
		resp, err := e2e.RunAndGetStdOut(cmd)
		req.NoError(err, string(resp))

		var config atlasv2.ConnectedOrgConfig
		req.NoError(json.Unmarshal(resp, &config))

		assert.Contains(t, config.GetDataAccessIdentityProviderIds(), oidcIWorkforceIdpID)
		assert.Contains(t, config.GetDataAccessIdentityProviderIds(), oidcWorkloadIdpID)
	})

	t.Run("Disconnect OIDC IdP WORKLOAD", func(t *testing.T) {
		cmd := exec.Command(cliPath,
			federatedAuthenticationEntity,
			federationSettingsEntity,
			connectedOrgsConfigsEntity,
			"disconnect",
			"--identityProviderId",
			oidcWorkloadIdpID,
			"--federationSettingsId",
			federationSettingsID,
			"-o=json",
		)

		cmd.Env = os.Environ()
		resp, err := e2e.RunAndGetStdOut(cmd)
		req.NoError(err, string(resp))

		var config atlasv2.ConnectedOrgConfig
		req.NoError(json.Unmarshal(resp, &config))

		assert.NotContains(t, config.GetDataAccessIdentityProviderIds(), oidcWorkloadIdpID)
	})

	t.Run("Disconnect OIDC IdP WORKFORCE", func(t *testing.T) {
		cmd := exec.Command(cliPath,
			federatedAuthenticationEntity,
			federationSettingsEntity,
			connectedOrgsConfigsEntity,
			"disconnect",
			"--identityProviderId",
			oidcIWorkforceIdpID,
			"--federationSettingsId",
			federationSettingsID,
			"-o=json",
		)

		cmd.Env = os.Environ()
		resp, err := e2e.RunAndGetStdOut(cmd)
		req.NoError(err, string(resp))

		var config atlasv2.ConnectedOrgConfig
		req.NoError(json.Unmarshal(resp, &config))

		assert.NotContains(t, config.GetDataAccessIdentityProviderIds(), oidcIWorkforceIdpID)
	})

	t.Run("Update connected org config", func(t *testing.T) {
		cmd := exec.Command(cliPath,
			federatedAuthenticationEntity,
			federationSettingsEntity,
			connectedOrgsConfigsEntity,
			"update",
			"--federationSettingsId",
			federationSettingsID,
			"--file",
			"data/connected_org_config.json",
			"-o=json",
		)

		cmd.Env = os.Environ()
		resp, err := e2e.RunAndGetStdOut(cmd)
		req.NoError(err, string(resp))

		var config atlasv2.ConnectedOrgConfig
		req.NoError(json.Unmarshal(resp, &config))

		assert.NotEmpty(t, config.GetDomainAllowList())
	})

	t.Run("Update connected org config back", func(t *testing.T) {
		cmd := exec.Command(cliPath,
			federatedAuthenticationEntity,
			federationSettingsEntity,
			connectedOrgsConfigsEntity,
			"update",
			"--federationSettingsId",
			federationSettingsID,
			"--file",
			"data/connected_org_config_update.json",
			"-o=json",
		)

		cmd.Env = os.Environ()
		resp, err := e2e.RunAndGetStdOut(cmd)
		req.NoError(err, string(resp))

		var config atlasv2.ConnectedOrgConfig
		req.NoError(json.Unmarshal(resp, &config))

		assert.Empty(t, config.GetDomainAllowList())
	})

	t.Run("List OIDC IdPs WORKFORCE", func(_ *testing.T) {
		cmd := exec.Command(cliPath,
			federatedAuthenticationEntity,
			federationSettingsEntity,
			identityProviderEntity,
			"list",
			"--federationSettingsId",
			federationSettingsID,
			"--protocol",
			"OIDC",
			"--idpType",
			"WORKFORCE",
			"-o=json",
		)

		cmd.Env = os.Environ()
		resp, err := e2e.RunAndGetStdOut(cmd)
		req.NoError(err, string(resp))

		var provider atlasv2.FederationIdentityProvider
		req.NoError(json.Unmarshal(resp, &provider))
	})

	t.Run("List OIDC IdPs WORKLOAD", func(_ *testing.T) {
		cmd := exec.Command(cliPath,
			federatedAuthenticationEntity,
			federationSettingsEntity,
			identityProviderEntity,
			"list",
			"--federationSettingsId",
			federationSettingsID,
			"--protocol",
			"OIDC",
			"--idpType",
			"WORKLOAD",
			"-o=json",
		)

		cmd.Env = os.Environ()
		resp, err := e2e.RunAndGetStdOut(cmd)
		req.NoError(err, string(resp))

		var provider atlasv2.FederationIdentityProvider
		req.NoError(json.Unmarshal(resp, &provider))
	})

	t.Run("List SAML IdPs", func(_ *testing.T) {
		cmd := exec.Command(cliPath,
			federatedAuthenticationEntity,
			federationSettingsEntity,
			identityProviderEntity,
			"list",
			"--federationSettingsId",
			federationSettingsID,
			"--protocol",
			"SAML",
			"-o=json",
		)

		cmd.Env = os.Environ()
		resp, err := e2e.RunAndGetStdOut(cmd)
		req.NoError(err, string(resp))

		var provider atlasv2.FederationIdentityProvider
		req.NoError(json.Unmarshal(resp, &provider))
	})

	t.Run("List connectedOrgsConfig", func(t *testing.T) {
		cmd := exec.Command(cliPath,
			federatedAuthenticationEntity,
			federationSettingsEntity,
			connectedOrgsConfigsEntity,
			"list",
			"--federationSettingsId",
			federationSettingsID,
			"-o=json",
		)

		cmd.Env = os.Environ()
		resp, err := e2e.RunAndGetStdOut(cmd)
		req.NoError(err, string(resp))

		var config atlasv2.PaginatedConnectedOrgConfigs
		req.NoError(json.Unmarshal(resp, &config))

		assert.NotEmpty(t, config.GetResults())
		assert.Len(t, config.GetResults(), 1)
		assert.NotContains(t, config.GetResults()[0].GetDataAccessIdentityProviderIds(), oidcIWorkforceIdpID)
		assert.NotContains(t, config.GetResults()[0].GetDataAccessIdentityProviderIds(), oidcWorkloadIdpID)
	})

	t.Run("Describe OIDC IdP WORKFORCE", func(t *testing.T) {
		cmd := exec.Command(cliPath,
			federatedAuthenticationEntity,
			federationSettingsEntity,
			identityProviderEntity,
			"describe",
			oidcIWorkforceIdpID,
			"--federationSettingsId",
			federationSettingsID,
			"-o=json",
		)

		cmd.Env = os.Environ()
		resp, err := e2e.RunAndGetStdOut(cmd)
		req.NoError(err, string(resp))

		var provider atlasv2.FederationIdentityProvider
		req.NoError(json.Unmarshal(resp, &provider))
		assert.NotEmpty(t, provider.GetId())
	})

	t.Run("Revoke JWK from OIDC IdP WORKFORCE", func(_ *testing.T) {
		cmd := exec.Command(cliPath,
			federatedAuthenticationEntity,
			federationSettingsEntity,
			identityProviderEntity,
			"revokeJwk",
			oidcIWorkforceIdpID,
			"--federationSettingsId",
			federationSettingsID,
			"-o=json",
		)

		cmd.Env = os.Environ()
		resp, err := e2e.RunAndGetStdOut(cmd)
		req.NoError(err, string(resp))
	})

	t.Run("Delete OIDC IdP WORKFORCE", func(_ *testing.T) {
		cmd := exec.Command(cliPath,
			federatedAuthenticationEntity,
			federationSettingsEntity,
			identityProviderEntity,
			"delete",
			oidcIWorkforceIdpID,
			"--federationSettingsId",
			federationSettingsID,
			"--force",
			"-o=json",
		)

		cmd.Env = os.Environ()
		resp, err := e2e.RunAndGetStdOut(cmd)
		req.NoError(err, string(resp))
	})

	t.Run("Revoke JWK from OIDC IdP WORKLOAD", func(_ *testing.T) {
		cmd := exec.Command(cliPath,
			federatedAuthenticationEntity,
			federationSettingsEntity,
			identityProviderEntity,
			"revokeJwk",
			oidcWorkloadIdpID,
			"--federationSettingsId",
			federationSettingsID,
			"-o=json",
		)

		cmd.Env = os.Environ()
		resp, err := e2e.RunAndGetStdOut(cmd)
		req.NoError(err, string(resp))
	})

	t.Run("Delete OIDC IdP WORKLOAD", func(_ *testing.T) {
		cmd := exec.Command(cliPath,
			federatedAuthenticationEntity,
			federationSettingsEntity,
			identityProviderEntity,
			"delete",
			oidcWorkloadIdpID,
			"--federationSettingsId",
			federationSettingsID,
			"--force",
			"-o=json",
		)

		cmd.Env = os.Environ()
		resp, err := e2e.RunAndGetStdOut(cmd)
		req.NoError(err, string(resp))
	})
}
