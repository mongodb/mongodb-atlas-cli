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
	atlasv2 "go.mongodb.org/atlas-sdk/v20230201008/admin"
)

func TestBackupCompliancePolicySetup(t *testing.T) {
	cliPath, err := e2e.AtlasCLIBin()
	r := require.New(t)
	r.NoError(err)

	g := newAtlasE2ETestGenerator(t)
	g.generateProject("setup-compliance-policy")

	scheduledPolicyItem := atlasv2.DiskBackupApiPolicyItem{
		FrequencyInterval: 1,
		FrequencyType:     "daily",
		RetentionUnit:     "days",
		RetentionValue:    1,
	}
	policy := &atlasv2.DataProtectionSettings{
		ScheduledPolicyItems: []atlasv2.DiskBackupApiPolicyItem{scheduledPolicyItem},
		ProjectId:            &g.projectID,
		AuthorizedEmail:      authorizedEmail,
	}
	path := "./compliancepolicy.json"

	createJSONFile(t, policy, path)

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

	var result atlasv2.DataProtectionSettings
	r.NoError(json.Unmarshal(resp, &result), resp)

	a := assert.New(t)
	a.Len(result.GetScheduledPolicyItems(), 1)
	a.Equal(authorizedEmail, result.GetAuthorizedEmail())
}
