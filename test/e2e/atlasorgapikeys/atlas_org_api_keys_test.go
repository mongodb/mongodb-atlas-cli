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
//go:build e2e || e2eSnap || (iam && atlas)

package atlasorgapikeys

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"testing"

	"github.com/mongodb/mongodb-atlas-cli/atlascli/test/internal"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	atlasv2 "go.mongodb.org/atlas-sdk/v20250312004/admin"
)

const (
	orgEntity     = "org"
	apiKeysEntity = "apikeys"
)

func TestAtlasOrgAPIKeys(t *testing.T) {
	g := internal.NewAtlasE2ETestGenerator(t, internal.WithSnapshot())
	cliPath, err := internal.AtlasCLIBin()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	var ID string

	// This test must run first to grab the ID of the org to later describe
	g.Run("Create", func(t *testing.T) { //nolint:thelper // g.Run replaces t.Run
		desc := "e2e-test-atlas-org"
		cmd := exec.Command(cliPath,
			orgEntity,
			apiKeysEntity,
			"create",
			"--desc",
			desc,
			"--role=ORG_READ_ONLY",
			"-o=json")
		cmd.Env = os.Environ()
		resp, err := internal.RunAndGetStdOut(cmd)
		require.NoError(t, err, string(resp))
		var key atlasv2.ApiKeyUserDetails
		require.NoError(t, json.Unmarshal(resp, &key))
		assert.Equal(t, desc, *key.Desc)
		ID = *key.Id
	})
	require.NotEmpty(t, ID)

	g.Run("List", func(t *testing.T) { //nolint:thelper // g.Run replaces t.Run
		cmd := exec.Command(cliPath,
			orgEntity,
			apiKeysEntity,
			"ls",
			"-o=json")
		cmd.Env = os.Environ()
		resp, err := internal.RunAndGetStdOut(cmd)
		require.NoError(t, err, string(resp))
		var keys atlasv2.PaginatedApiApiUser
		require.NoError(t, json.Unmarshal(resp, &keys))
		assert.NotEmpty(t, keys.Results)
	})

	g.Run("List Compact", func(t *testing.T) { //nolint:thelper // g.Run replaces t.Run
		cmd := exec.Command(cliPath,
			orgEntity,
			apiKeysEntity,
			"ls",
			"-c",
			"-o=json")
		cmd.Env = os.Environ()
		resp, err := internal.RunAndGetStdOut(cmd)
		require.NoError(t, err, string(resp))

		var keys []atlasv2.ApiKeyUserDetails
		require.NoError(t, json.Unmarshal(resp, &keys))
		assert.NotEmpty(t, keys)
	})

	g.Run("Update", func(t *testing.T) { //nolint:thelper // g.Run replaces t.Run
		const newDesc = "e2e-test-atlas-org-updated"
		cmd := exec.Command(cliPath,
			orgEntity,
			apiKeysEntity,
			"updates",
			ID,
			"--desc",
			newDesc,
			"--role=ORG_READ_ONLY",
			"-o=json")
		cmd.Env = os.Environ()
		resp, err := internal.RunAndGetStdOut(cmd)
		require.NoError(t, err, string(resp))
		var key atlasv2.ApiKeyUserDetails
		require.NoError(t, json.Unmarshal(resp, &key))
		assert.Equal(t, newDesc, *key.Desc)
	})

	g.Run("Describe", func(t *testing.T) { //nolint:thelper // g.Run replaces t.Run
		cmd := exec.Command(cliPath,
			orgEntity,
			apiKeysEntity,
			"describe",
			ID,
			"-o=json")
		cmd.Env = os.Environ()
		resp, err := internal.RunAndGetStdOut(cmd)
		require.NoError(t, err, string(resp))
		var key atlasv2.ApiKeyUserDetails
		require.NoError(t, json.Unmarshal(resp, &key))
		assert.Equal(t, ID, *key.Id)
	})

	g.Run("Delete", func(t *testing.T) { //nolint:thelper // g.Run replaces t.Run
		cmd := exec.Command(cliPath,
			orgEntity,
			apiKeysEntity,
			"rm",
			ID,
			"--force")
		cmd.Env = os.Environ()
		resp, err := internal.RunAndGetStdOut(cmd)
		require.NoError(t, err, string(resp))
		expected := fmt.Sprintf("API Key '%s' deleted\n", ID)
		assert.Equal(t, expected, string(resp))
	})
}
