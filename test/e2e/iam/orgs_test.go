// Copyright 2020 MongoDB Inc
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

//go:build e2e || (iam && !atlas)

package iam_test

import (
	"encoding/json"
	"os"
	"os/exec"
	"testing"

	"github.com/mongodb/mongodb-atlas-cli/test/e2e"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	atlasv2 "go.mongodb.org/atlas-sdk/v20231115004/admin"
)

func TestOrgs(t *testing.T) {
	cliPath, err := e2e.Bin()
	require.NoError(t, err)

	var orgID string

	// This test must run first to grab the ID of the org to later describe
	t.Run("List", func(t *testing.T) {
		cmd := exec.Command(cliPath,
			iamEntity,
			orgEntity,
			"ls",
			"-o=json")
		cmd.Env = os.Environ()
		resp, err := cmd.CombinedOutput()
		require.NoError(t, err, string(resp))
		var orgs atlasv2.PaginatedOrganization
		err = json.Unmarshal(resp, &orgs)
		require.NoError(t, err, string(resp))
		assert.NotEmpty(t, orgs.Results)
		orgID = *orgs.GetResults()[0].Id
		require.NotEmpty(t, orgID, "orgID not set, resp: %s", resp)
	})
	require.NotEmpty(t, orgID)

	t.Run("Describe", func(t *testing.T) {
		cmd := exec.Command(cliPath,
			iamEntity,
			orgEntity,
			"describe",
			orgID,
			"-o=json")
		cmd.Env = os.Environ()
		resp, err := cmd.CombinedOutput()
		require.NoError(t, err, string(resp))
	})

	t.Run("List Org Users", func(t *testing.T) {
		cmd := exec.Command(cliPath,
			iamEntity,
			orgEntity,
			usersEntity,
			"ls",
			"--orgId",
			orgID,
			"-o=json")
		cmd.Env = os.Environ()
		resp, err := cmd.CombinedOutput()
		require.NoError(t, err, string(resp))
	})
}
