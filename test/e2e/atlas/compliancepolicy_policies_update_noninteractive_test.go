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
func testPoliciesUpdate(t *testing.T, g *atlasE2ETestGenerator) {
	cliPath, err := e2e.AtlasCLIBin()
	r := require.New(t)
	r.NoError(err)

	t.Run("policies update happy flow", func(t *testing.T) {
		expected := atlasv2.DiskBackupApiPolicyItem{
			FrequencyInterval: 1,
			FrequencyType:     "ondemand",
			RetentionUnit:     "days",
			RetentionValue:    1,
		}

		path := "./ondemanditem.json"

		createJSONFile(t, expected, path)

		cmd := exec.Command(cliPath,
			backupsEntity,
			compliancepolicyEntity,
			"policies",
			"update",
			"--file",
			path,
			"-o=json",
			"--projectId",
			g.projectID,
			"--watch",
		)
		cmd.Env = os.Environ()
		resp, outputErr := cmd.CombinedOutput()

		r.NoError(outputErr, string(resp))
		a := assert.New(t)

		trimmedResponse := removeDotsFromWatching(resp)

		var result atlasv2.DataProtectionSettings
		require.NoError(t, json.Unmarshal(trimmedResponse, &result), string(trimmedResponse))

		actual := result.GetOnDemandPolicyItem()

		a.Equal(expected.GetFrequencyType(), actual.GetFrequencyType())
	})

	t.Run("policies update ID not found", func(t *testing.T) {
		expected := atlasv2.DiskBackupApiPolicyItem{
			Id: atlasv2.PtrString("non-existent ID"),
		}

		path := "./fakeItem.json"

		createJSONFile(t, expected, path)

		cmd := exec.Command(cliPath,
			backupsEntity,
			compliancepolicyEntity,
			"policies",
			"update",
			"--file",
			path,
			"-o=json",
			"--projectId",
			g.projectID,
			"--watch",
		)
		cmd.Env = os.Environ()
		resp, outputErr := cmd.CombinedOutput()

		r.Error(outputErr, string(resp))
	})
}
