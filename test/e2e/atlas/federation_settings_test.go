// Copyright 2023 MongoDB Inc
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
	"os"
	"os/exec"
	"testing"

	"github.com/mongodb/mongodb-atlas-cli/atlascli/test/e2e"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	atlasv2 "go.mongodb.org/atlas-sdk/v20231115010/admin"
)

func TestFederationSettings(t *testing.T) {
	_ = newAtlasE2ETestGenerator(t)
	req := require.New(t)

	cliPath, err := e2e.AtlasCLIBin()
	req.NoError(err)

	var orgID string
	var set bool
	if orgID, set = os.LookupEnv("MCLI_ORG_ID"); !set {
		t.Skip("MCLI_ORG_ID must be set")
	}

	t.Run("Describing an org federation settings", func(t *testing.T) {
		cmd := exec.Command(cliPath,
			"federatedAuthentication",
			"federationSettings",
			"describe",
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
	})
}
