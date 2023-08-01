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

//go:build e2e || (atlas && vinc)

package atlas_test

import (
	"encoding/json"
	"os"
	"os/exec"
	"testing"

	"github.com/mongodb/mongodb-atlas-cli/test/e2e"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	atlasv2 "go.mongodb.org/atlas-sdk/v20230201003/admin"
)

func TestCompliancePolicy(t *testing.T) {
	cliPath, err := e2e.AtlasCLIBin()
	r := require.New(t)
	r.NoError(err)

	g := newAtlasE2ETestGeneratorWithBackup(t)
	g.generateProject("compliancePolicy")

	t.Run("Compliance Policy Describe", func(t *testing.T) {
		cmd := exec.Command(cliPath,
			backupsEntity,
			compliancepolicyEntity,
			"describe",
			"--projectId",
			g.projectID,
			"-o=json")
		cmd.Env = os.Environ()
		resp, err := cmd.CombinedOutput()

		r.NoError(err, string(resp))

		a := assert.New(t)
		var result atlasv2.DataProtectionSettings
		err = json.Unmarshal(resp, &result)
		a.NoError(err, string(resp))
		// Will be changed after implementing enable/setup.
		// a.NotEmpty(result)
	})
}
