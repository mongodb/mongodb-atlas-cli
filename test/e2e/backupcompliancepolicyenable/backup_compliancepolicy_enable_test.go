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

//go:build e2e || e2eSnap || (atlas && backup && compliancepolicy)

package backupcompliancepolicyenable

import (
	"encoding/json"
	"os"
	"os/exec"
	"testing"

	"github.com/mongodb/mongodb-atlas-cli/atlascli/test/internal"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	atlasv2 "go.mongodb.org/atlas-sdk/v20250312005/admin"
)

const (
	backupsEntity          = "backups"
	compliancePolicyEntity = "compliancepolicy"

	// Backup compliance policy constants.
	authorizedUserFirstName = "firstname"
	authorizedUserLastName  = "lastname"
	authorizedEmail         = "firstname.lastname@example.com"
)

func TestBackupCompliancePolicyEnable(t *testing.T) {
	g := internal.NewAtlasE2ETestGenerator(t, internal.WithSnapshot())
	cliPath, err := internal.AtlasCLIBin()
	r := require.New(t)
	r.NoError(err)

	g.GenerateProject("enable-compliance-policy")

	cmd := exec.Command(cliPath,
		backupsEntity,
		compliancePolicyEntity,
		"enable",
		"--projectId",
		g.ProjectID,
		"--authorizedUserFirstName",
		authorizedUserFirstName,
		"--authorizedUserLastName",
		authorizedUserLastName,
		"--authorizedEmail",
		authorizedEmail,
		"-o=json",
		"--force",
	)
	cmd.Env = os.Environ()
	resp, outputErr := internal.RunAndGetStdOut(cmd)
	r.NoError(outputErr, string(resp))
	var result atlasv2.DataProtectionSettings20231001
	r.NoError(json.Unmarshal(resp, &result), string(resp))

	assert.Equal(t, authorizedEmail, result.GetAuthorizedEmail())
}
