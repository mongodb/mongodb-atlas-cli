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
	atlasv2 "go.mongodb.org/atlas-sdk/v20231115010/admin"
)

func TestIdentityProviders(t *testing.T) {
	_ = newAtlasE2ETestGenerator(t)
	req := require.New(t)

	cliPath, err := e2e.AtlasCLIBin()
	req.NoError(err)

	var federationSettingsID *string
	var oidcIdentityProviderID *string
	var orgID string
	var set bool
	if orgID, set = os.LookupEnv("MCLI_ORG_ID"); !set {
		t.Skip("MCLI_ORG_ID must be set")
	}

	t.Run("Describe an org federation settings", func(t *testing.T) {
		cmd := exec.Command(cliPath,
			federatedAuthenticationEntity,
			federationSettingsEntity,
			"describe",
			"--orgId",
			orgID,
			"-o=json",
		)

		cmd.Env = os.Environ()
		resp, err := cmd.CombinedOutput()
		req.NoError(err, string(resp))

		var settings atlasv2.OrgFederationSettings
		req.NoError(json.Unmarshal(resp, &settings))

		a := assert.New(t)
		a.NotEmpty(settings.Id)
		a.NotEmpty(settings.IdentityProviderStatus)
		federationSettingsID = settings.Id
	})

	t.Run("List OIDC IdPs WORKFORCE", func(_ *testing.T) {
		cmd := exec.Command(cliPath,
			federatedAuthenticationEntity,
			identityProviderEntity,
			"list",
			"--federationSettingsId",
			*federationSettingsID,
			"--protocol",
			"OIDC",
			"--idpType",
			"WORKFORCE",
			"-o=json",
		)

		fmt.Println("Printing federationSettingsID")
		fmt.Println(*federationSettingsID)

		cmd.Env = os.Environ()
		resp, err := cmd.CombinedOutput()
		req.NoError(err, string(resp))

		var settings atlasv2.FederationIdentityProvider
		req.NoError(json.Unmarshal(resp, &settings))
	})

	t.Run("List OIDC IdPs WORKLOAD", func(_ *testing.T) {
		cmd := exec.Command(cliPath,
			federatedAuthenticationEntity,
			identityProviderEntity,
			"list",
			"--federationSettingsId",
			*federationSettingsID,
			"--protocol",
			"OIDC",
			"--idpType",
			"WORKLOAD",
			"-o=json",
		)

		cmd.Env = os.Environ()
		resp, err := cmd.CombinedOutput()
		req.NoError(err, string(resp))

		var settings atlasv2.FederationIdentityProvider
		req.NoError(json.Unmarshal(resp, &settings))
	})

	t.Run("List SAML IdPs", func(_ *testing.T) {
		cmd := exec.Command(cliPath,
			federatedAuthenticationEntity,
			identityProviderEntity,
			"list",
			"--federationSettingsId",
			*federationSettingsID,
			"--protocol",
			"SAML",
			"-o=json",
		)

		cmd.Env = os.Environ()
		resp, err := cmd.CombinedOutput()
		req.NoError(err, string(resp))

		var settings atlasv2.FederationIdentityProvider
		req.NoError(json.Unmarshal(resp, &settings))
	})

	t.Run("Create OIDC IdP WORKLOAD", func(t *testing.T) {
		cmd := exec.Command(cliPath,
			federatedAuthenticationEntity,
			identityProviderEntity,
			"create",
			"oidc",
			"cliTestProvider",
			"--federationSettingsId",
			*federationSettingsID,
			"--audience",
			"AtlasCLIAudience",
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
		resp, err := cmd.CombinedOutput()
		req.NoError(err, string(resp))

		var provider atlasv2.FederationIdentityProvider
		req.NoError(json.Unmarshal(resp, &provider))

		a := assert.New(t)
		a.NotEmpty(provider.Id)
		oidcIdentityProviderID = &provider.Id
	})

	t.Run("Describe an OIDC identity provider of type WORKFORCE", func(t *testing.T) {
		cmd := exec.Command(cliPath,
			federatedAuthenticationEntity,
			identityProviderEntity,
			"describe",
			*oidcIdentityProviderID,
			"--federationSettingsId",
			*federationSettingsID,
			"-o=json",
		)

		cmd.Env = os.Environ()
		resp, err := cmd.CombinedOutput()
		req.NoError(err, string(resp))

		var settings atlasv2.FederationIdentityProvider
		req.NoError(json.Unmarshal(resp, &settings))

		a := assert.New(t)
		a.NotEmpty(settings.Id)
	})

	t.Run("Delete an OIDC identity provider of type WORKFORCE", func(t *testing.T) {
		cmd := exec.Command(cliPath,
			federatedAuthenticationEntity,
			identityProviderEntity,
			"delete",
			*oidcIdentityProviderID,
			"--federationSettingsId",
			*federationSettingsID,
			"-o=json",
		)

		cmd.Env = os.Environ()
		resp, err := cmd.CombinedOutput()
		req.NoError(err, string(resp))

		var settings atlasv2.FederationIdentityProvider
		req.NoError(json.Unmarshal(resp, &settings))

		a := assert.New(t)
		a.NotEmpty(settings.Id)
	})

	t.Run("Create OIDC IdP WORKFORCE", func(t *testing.T) {
		cmd := exec.Command(cliPath,
			federatedAuthenticationEntity,
			identityProviderEntity,
			"create",
			"oidc",
			"cliTestProvider",
			"--federationSettingsId",
			*federationSettingsID,
			"--audience",
			"AtlasCLIAudience",
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
		resp, err := cmd.CombinedOutput()
		req.NoError(err, string(resp))

		var provider atlasv2.FederationIdentityProvider
		req.NoError(json.Unmarshal(resp, &provider))

		a := assert.New(t)
		a.NotEmpty(provider.Id)
		oidcIdentityProviderID = &provider.Id
	})

	t.Run("Describe an OIDC identity provider of type WORKFORCE", func(t *testing.T) {
		cmd := exec.Command(cliPath,
			federatedAuthenticationEntity,
			identityProviderEntity,
			"describe",
			*oidcIdentityProviderID,
			"--federationSettingsId",
			*federationSettingsID,
			"-o=json",
		)

		cmd.Env = os.Environ()
		resp, err := cmd.CombinedOutput()
		req.NoError(err, string(resp))

		var settings atlasv2.FederationIdentityProvider
		req.NoError(json.Unmarshal(resp, &settings))

		a := assert.New(t)
		a.NotEmpty(settings.Id)
	})

	t.Run("Delete an OIDC identity provider of type WORKFORCE", func(t *testing.T) {
		cmd := exec.Command(cliPath,
			federatedAuthenticationEntity,
			identityProviderEntity,
			"delete",
			*oidcIdentityProviderID,
			"--federationSettingsId",
			*federationSettingsID,
			"-o=json",
		)

		cmd.Env = os.Environ()
		resp, err := cmd.CombinedOutput()
		req.NoError(err, string(resp))

		var settings atlasv2.FederationIdentityProvider
		req.NoError(json.Unmarshal(resp, &settings))

		a := assert.New(t)
		a.NotEmpty(settings.Id)
	})
}
