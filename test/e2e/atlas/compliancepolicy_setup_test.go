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

package atlas_test

import (
	"encoding/json"
	"os"
	"os/exec"
	"testing"

	"github.com/mongodb/mongodb-atlas-cli/test/e2e"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	atlasv2 "go.mongodb.org/atlas-sdk/v20230201004/admin"
)

//nolint:thelper
func testCompliancePolicySetup(t *testing.T, projectID string) {
	cliPath, err := e2e.AtlasCLIBin()
	r := require.New(t)
	r.NoError(err)

	scheduledPolicyItem := atlasv2.DiskBackupApiPolicyItem{
		FrequencyInterval: 1,
		FrequencyType:     "daily",
		RetentionUnit:     "days",
		RetentionValue:    1,
	}

	email := authorizedEmail

	policy := &atlasv2.DataProtectionSettings{
		ScheduledPolicyItems: []atlasv2.DiskBackupApiPolicyItem{scheduledPolicyItem},
		ProjectId:            &projectID,
		AuthorizedEmail:      &email,
	}
	path := "./compliancepolicy.json"

	createJSONFile(t, policy, path)

	t.Run("setup happy flow", func(t *testing.T) {
		cmd := exec.Command(cliPath,
			backupsEntity,
			compliancepolicyEntity,
			"setup",
			"--projectId",
			projectID,
			"-o=json",
			"--force",
			"--file",
			path,
			"--watch", // avoiding HTTP 400 Bad Request "CANNOT_UPDATE_BACKUP_COMPLIANCE_POLICY_SETTINGS_WITH_PENDING_ACTION".
		)

		cmd.Env = os.Environ()
		resp, outputErr := cmd.CombinedOutput()

		trimmedResponse := removeDotsFromWatching(resp)

		r.NoError(outputErr, string(resp))
		a := assert.New(t)

		var result atlasv2.DataProtectionSettings
		require.NoError(t, json.Unmarshal(trimmedResponse, &result), trimmedResponse)
		a.Len(result.GetScheduledPolicyItems(), 1)
		a.Equal(result.GetAuthorizedEmail(), authorizedEmail)
	})
}
