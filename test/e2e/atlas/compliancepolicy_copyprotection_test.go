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

func testCopyProtection(t *testing.T, g *atlasE2ETestGenerator) {
	cliPath, err := e2e.AtlasCLIBin()
	r := require.New(t)
	r.NoError(err)

	t.Run("copyprotection happy flow", func(t *testing.T) {
		cmd := exec.Command(
			cliPath,
			backupsEntity,
			compliancepolicyEntity,
			"copyprotection",
			"enable",
			"-o=json",
			"--projectId",
			g.projectID,
			"--watch",
		)
		cmd.Env = os.Environ()
		resp, outputErr := cmd.CombinedOutput()
		r.NoError(outputErr, string(resp))

		trimmedResponse := removeDotsFromWatching(resp)

		a := assert.New(t)

		var compliancepolicy atlasv2.DataProtectionSettings
		require.NoError(t, json.Unmarshal(trimmedResponse, &compliancepolicy), string(trimmedResponse))
		a.True(*compliancepolicy.CopyProtectionEnabled)
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
