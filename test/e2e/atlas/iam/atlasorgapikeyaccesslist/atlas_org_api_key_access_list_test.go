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

package atlasorgapikeyaccesslist

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"testing"

	"github.com/mongodb/mongodb-atlas-cli/atlascli/test/internal"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	atlasv2 "go.mongodb.org/atlas-sdk/v20250312005/admin"
)

const (
	orgEntity              = "org"
	apiKeysEntity          = "apikeys"
	apiKeyAccessListEntity = "accessLists"
)

var errNoAPIKey = errors.New("the apiKey ID is empty")

func TestAtlasOrgAPIKeyAccessList(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping test in short mode")
	}

	g := internal.NewAtlasE2ETestGenerator(t, internal.WithSnapshot())
	cliPath, er := internal.AtlasCLIBin()
	require.NoError(t, er)

	apiKeyID, e := createOrgAPIKey()
	require.NoError(t, e)

	t.Cleanup(func() {
		require.NoError(t, internal.DeleteOrgAPIKey(apiKeyID))
	})

	n := g.MemoryRand("rand", 255)
	entry := fmt.Sprintf("192.168.0.%d", n)

	g.Run("Create", func(t *testing.T) { //nolint:thelper // g.Run replaces t.Run
		cmd := exec.Command(cliPath,
			orgEntity,
			apiKeysEntity,
			apiKeyAccessListEntity,
			"create",
			"--apiKey",
			apiKeyID,
			"--ip",
			entry,
			"-o=json")
		cmd.Env = os.Environ()
		resp, err := internal.RunAndGetStdOut(cmd)
		require.NoError(t, err, string(resp))
		var key atlasv2.PaginatedApiUserAccessListResponse
		require.NoError(t, json.Unmarshal(resp, &key))
		assert.NotEmpty(t, key.Results)
	})

	g.Run("List", func(t *testing.T) { //nolint:thelper // g.Run replaces t.Run
		cmd := exec.Command(cliPath,
			orgEntity,
			apiKeysEntity,
			apiKeyAccessListEntity,
			"list",
			apiKeyID,
			"-o=json")
		cmd.Env = os.Environ()
		resp, err := internal.RunAndGetStdOut(cmd)
		require.NoError(t, err, string(resp))
		var key atlasv2.PaginatedApiUserAccessListResponse
		require.NoError(t, json.Unmarshal(resp, &key))
		assert.NotEmpty(t, key.Results)
	})

	g.Run("Delete", func(t *testing.T) { //nolint:thelper // g.Run replaces t.Run
		deleteAtlasAccessListEntry(t, cliPath, entry, apiKeyID)
	})

	g.Run("Create Current IP", func(t *testing.T) { //nolint:thelper // g.Run replaces t.Run
		cmd := exec.Command(cliPath,
			orgEntity,
			apiKeysEntity,
			apiKeyAccessListEntity,
			"create",
			"--apiKey",
			apiKeyID,
			"--currentIp",
			"-o=json")
		cmd.Env = os.Environ()
		resp, err := internal.RunAndGetStdOut(cmd)
		require.NoError(t, err, string(resp))
		var key atlasv2.PaginatedApiUserAccessListResponse
		require.NoError(t, json.Unmarshal(resp, &key))
		assert.NotEmpty(t, key.Results)
		entry = *key.GetResults()[0].IpAddress
	})

	g.Run("Delete", func(t *testing.T) { //nolint:thelper // g.Run replaces t.Run
		deleteAtlasAccessListEntry(t, cliPath, entry, apiKeyID)
	})
}

func deleteAtlasAccessListEntry(t *testing.T, cliPath, entry, apiKeyID string) {
	t.Helper()
	cmd := exec.Command(cliPath,
		orgEntity,
		apiKeysEntity,
		apiKeyAccessListEntity,
		"rm",
		entry,
		"--apiKey",
		apiKeyID,
		"--force")
	cmd.Env = os.Environ()
	resp, err := internal.RunAndGetStdOut(cmd)
	require.NoError(t, err, string(resp))
	expected := fmt.Sprintf("Access list entry '%s' deleted\n", entry)
	assert.Equal(t, expected, string(resp))
}

func createOrgAPIKey() (string, error) {
	cliPath, err := internal.AtlasCLIBin()
	if err != nil {
		return "", err
	}

	cmd := exec.Command(cliPath,
		orgEntity,
		apiKeysEntity,
		"create",
		"--desc=e2e-test-helper",
		"--role=ORG_READ_ONLY",
		"-o=json")
	cmd.Env = os.Environ()
	resp, err := internal.RunAndGetStdOut(cmd)

	if err != nil {
		return "", fmt.Errorf("%w: %s", err, string(resp))
	}

	var key atlasv2.ApiKeyUserDetails
	if err := json.Unmarshal(resp, &key); err != nil {
		return "", err
	}

	if key.GetId() != "" {
		return key.GetId(), nil
	}

	return "", errNoAPIKey
}
