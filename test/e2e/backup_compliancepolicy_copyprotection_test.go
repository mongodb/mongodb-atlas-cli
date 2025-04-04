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

//go:build e2e || (atlas && backup && compliancepolicy)

package e2e

import (
	"encoding/json"
	"os"
	"os/exec"
	"testing"

	"github.com/mongodb/mongodb-atlas-cli/atlascli/test/internal"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	atlasv2 "go.mongodb.org/atlas-sdk/v20250312001/admin"
)

func TestBackupCompliancePolicyCopyProtection(t *testing.T) {
	g := internal.NewAtlasE2ETestGenerator(t, internal.WithSnapshot())
	cliPath, err := internal.AtlasCLIBin()
	r := require.New(t)
	r.NoError(err)

	g.GenerateProject("copyprotection-compliance-policy")
	r.NoError(internal.EnableCompliancePolicy(g.ProjectID))

	g.Run("enable", func(t *testing.T) { //nolint:thelper // g.Run replaces t.Run
		cmd := exec.Command(
			cliPath,
			backupsEntity,
			compliancePolicyEntity,
			"copyprotection",
			"enable",
			"-o=json",
			"--projectId",
			g.ProjectID,
			"--watch", // avoiding HTTP 400 Bad Request "CANNOT_UPDATE_BACKUP_COMPLIANCE_POLICY_SETTINGS_WITH_PENDING_ACTION".
		)
		cmd.Env = os.Environ()
		resp, outputErr := internal.RunAndGetStdOut(cmd)
		r.NoError(outputErr, string(resp))

		trimmedResponse := internal.RemoveDotsFromWatching(resp)

		var compliancepolicy atlasv2.DataProtectionSettings20231001
		r.NoError(json.Unmarshal(trimmedResponse, &compliancepolicy), string(trimmedResponse))

		assert.True(t, *compliancepolicy.CopyProtectionEnabled)
	})

	g.Run("disable", func(t *testing.T) { //nolint:thelper // g.Run replaces t.Run
		cmd := exec.Command(
			cliPath,
			backupsEntity,
			compliancePolicyEntity,
			"copyprotection",
			"disable",
			"-o=json",
			"--projectId",
			g.ProjectID,
		)
		cmd.Env = os.Environ()
		resp, outputErr := internal.RunAndGetStdOut(cmd)
		r.NoError(outputErr, string(resp))

		var compliancepolicy atlasv2.DataProtectionSettings20231001
		r.NoError(json.Unmarshal(resp, &compliancepolicy), string(resp))

		assert.False(t, *compliancepolicy.CopyProtectionEnabled)
	})
}
