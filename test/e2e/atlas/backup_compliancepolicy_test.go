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

func TestCompliancePolicy_enable(t *testing.T) {
	cliPath, err := e2e.AtlasCLIBin()
	r := require.New(t)
	r.NoError(err)

	g := newAtlasE2ETestGeneratorWithBackup(t)
	g.generateProject("compliancePolicy")

	authorizedEmail := "firstname.lastname@example.com"

	t.Run("enable happy flow", func(t *testing.T) {
		cmd := exec.Command(cliPath,
			backupsEntity,
			compliancepolicyEntity,
			"enable",
			"--projectId",
			g.projectID,
			"--authorizedEmail",
			authorizedEmail,
			"-o=json",
		)
		cmd.Env = os.Environ()
		resp, outputErr := cmd.CombinedOutput()
		r.NoError(outputErr, string(resp))
		var result atlasv2.DataProtectionSettings
		err = json.Unmarshal(resp, &result)
		a := assert.New(t)

		a.Equal(result.GetAuthorizedEmail(), authorizedEmail)
	})
}
func TestCompliancePolicy(t *testing.T) {
	cliPath, err := e2e.AtlasCLIBin()
	r := require.New(t)
	r.NoError(err)

	g := newAtlasE2ETestGeneratorWithBackup(t)
	g.generateProject("compliancePolicy")

	scheduledPolicyItem := atlasv2.DiskBackupApiPolicyItem{
		FrequencyInterval: 1,
		FrequencyType:     "daily",
		RetentionUnit:     "days",
		RetentionValue:    1,
	}

	authorizedEmail := "firstname.lastname@example.com"

	policy := &atlasv2.DataProtectionSettings{
		ScheduledPolicyItems: []atlasv2.DiskBackupApiPolicyItem{scheduledPolicyItem},
		ProjectId:            &g.projectID,
		AuthorizedEmail:      &authorizedEmail,
	}
	path := "./compliancepolicy.json"

	createCompliancePolicyJSONFile(t, policy, path)

	t.Run("setup happy flow", func(t *testing.T) {
		cmd := exec.Command(cliPath,
			backupsEntity,
			compliancepolicyEntity,
			"setup",
			"--projectId",
			g.projectID,
			"-o=json",
			"--force",
			"--file",
			path,
		)
		cmd.Env = os.Environ()
		resp, outputErr := cmd.CombinedOutput()

		r.NoError(outputErr, string(resp))
	})

	t.Run("describe", func(t *testing.T) {
		cmd := exec.Command(cliPath,
			backupsEntity,
			compliancepolicyEntity,
			"describe",
			"--projectId",
			g.projectID,
			"-o=json")
		cmd.Env = os.Environ()
		resp, outputErr := cmd.CombinedOutput()

		r.NoError(outputErr, string(resp))

		a := assert.New(t)
		var result atlasv2.DataProtectionSettings
		err = json.Unmarshal(resp, &result)
		a.NoError(err, string(resp))
		// Will be changed after implementing enable/setup.
		// a.NotEmpty(result) Ticket to enforce this: CLOUDP-193023
	})

	t.Run("copyprotection invalid argument", func(t *testing.T) {
		invalidArgument := "invalid"
		cmd := exec.Command(
			cliPath,
			backupsEntity,
			compliancepolicyEntity,
			"copyprotection",
			invalidArgument)
		cmd.Env = os.Environ()
		_, err = cmd.CombinedOutput()
		r.Error(err)
	})

	t.Run("copyprotection invalid nr of arguments", func(t *testing.T) {
		cmd := exec.Command(
			cliPath,
			backupsEntity,
			compliancepolicyEntity,
			"copyprotection")
		cmd.Env = os.Environ()
		_, err = cmd.CombinedOutput()
		r.Error(err)
	})
}

// createCompliancePolicyJSONFile creates a new JSON file at the specified path with the specified policy
// and also registers its deletion on test cleanup.
func createCompliancePolicyJSONFile(t *testing.T, policy *atlasv2.DataProtectionSettings, path string) {
	t.Helper()

	jsonData, err := json.Marshal(policy)
	if err != nil {
		t.Errorf("Error marshaling to JSON: %v", err)
		return
	}

	err = os.WriteFile(path, jsonData, 0600)
	if err != nil {
		t.Errorf("Error writing JSON to file: %v", err)
		return
	}

	t.Cleanup(func() {
		deleteFile(t, path)
	})
}

func deleteFile(t *testing.T, path string) {
	t.Helper()
	if err := os.Remove(path); err != nil {
		t.Errorf("Error deleting file: %v", err)
	}
}
