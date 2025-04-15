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

package backupcompliancepolicypitrestore

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

func TestBackupCompliancePolicyPointInTimeRestore(t *testing.T) {
	g := internal.NewAtlasE2ETestGenerator(t, internal.WithSnapshot())
	cliPath, err := internal.AtlasCLIBin()
	r := require.New(t)
	r.NoError(err)

	g.GenerateProject("compliance-policy-pointintimerestore")
	initialItem := atlasv2.BackupComplianceScheduledPolicyItem{
		FrequencyInterval: 1,
		FrequencyType:     "hourly",
		RetentionUnit:     "days",
		RetentionValue:    1,
	}
	compliancePolicy := atlasv2.DataProtectionSettings20231001{
		ScheduledPolicyItems: &[]atlasv2.BackupComplianceScheduledPolicyItem{initialItem},
	}
	res, err := internal.SetupCompliancePolicy(t, g.ProjectID, &compliancePolicy)
	r.NoError(err)
	assert.False(t, res.GetPitEnabled())
	assert.Zero(t, res.GetRestoreWindowDays())

	g.Run("enable", func(t *testing.T) { //nolint:thelper // g.Run replaces t.Run
		cmd := exec.Command(
			cliPath,
			backupsEntity,
			compliancePolicyEntity,
			"pointintimerestore",
			"enable",
			"-o=json",
			"--projectId",
			g.ProjectID,
			"--restoreWindowDays",
			"1",
		)
		cmd.Env = os.Environ()
		resp, outputErr := internal.RunAndGetStdOut(cmd)
		r.NoError(outputErr, string(resp))

		var compliancepolicy atlasv2.DataProtectionSettings20231001
		r.NoError(json.Unmarshal(resp, &compliancepolicy), string(resp))

		assert.True(t, compliancepolicy.GetPitEnabled())
		assert.Equal(t, 1, compliancepolicy.GetRestoreWindowDays())
	})
}
