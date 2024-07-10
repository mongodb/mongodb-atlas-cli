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

//go:build e2e || (iam && atlas)

package iam_test

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"testing"

	"github.com/mongodb/mongodb-atlas-cli/atlascli/test/e2e"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	atlasv2 "go.mongodb.org/atlas-sdk/v20240530002/admin"
)

func TestAtlasOrgAPIKeyAccessList(t *testing.T) {
	cliPath, er := e2e.AtlasCLIBin()
	require.NoError(t, er)

	apiKeyID, e := createOrgAPIKey()
	require.NoError(t, e)

	t.Cleanup(func() {
		require.NoError(t, deleteOrgAPIKey(apiKeyID))
	})

	n, err := e2e.RandInt(255)
	require.NoError(t, err)

	entry := fmt.Sprintf("192.168.0.%d", n)

	t.Run("Create", func(t *testing.T) {
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
		resp, err := e2e.RunAndGetStdOut(cmd)
		require.NoError(t, err, string(resp))
		var key atlasv2.PaginatedApiUserAccessListResponse
		require.NoError(t, json.Unmarshal(resp, &key))
		assert.NotEmpty(t, key.Results)
	})

	t.Run("List", func(t *testing.T) {
		cmd := exec.Command(cliPath,
			orgEntity,
			apiKeysEntity,
			apiKeyAccessListEntity,
			"list",
			apiKeyID,
			"-o=json")
		cmd.Env = os.Environ()
		resp, err := e2e.RunAndGetStdOut(cmd)
		require.NoError(t, err, string(resp))
		var key atlasv2.PaginatedApiUserAccessListResponse
		require.NoError(t, json.Unmarshal(resp, &key))
		assert.NotEmpty(t, key.Results)
	})

	t.Run("Delete", func(t *testing.T) {
		deleteAtlasAccessListEntry(t, cliPath, entry, apiKeyID)
	})

	t.Run("Create Current IP", func(t *testing.T) {
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
		resp, err := e2e.RunAndGetStdOut(cmd)
		require.NoError(t, err, string(resp))
		var key atlasv2.PaginatedApiUserAccessListResponse
		require.NoError(t, json.Unmarshal(resp, &key))
		assert.NotEmpty(t, key.Results)
		entry = *key.GetResults()[0].IpAddress
	})

	t.Run("Delete", func(t *testing.T) {
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
	resp, err := e2e.RunAndGetStdOut(cmd)
	require.NoError(t, err, string(resp))
	expected := fmt.Sprintf("Access list entry '%s' deleted\n", entry)
	assert.Equal(t, expected, string(resp))
}
