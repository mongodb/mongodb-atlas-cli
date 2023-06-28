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

	"github.com/mongodb/mongodb-atlas-cli/test/e2e"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	atlasv2 "go.mongodb.org/atlas-sdk/admin"
)

func TestAtlasOrgAPIKeys(t *testing.T) {
	cliPath, err := e2e.AtlasCLIBin()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	var ID string

	// This test must run first to grab the ID of the org to later describe
	t.Run("Create", func(t *testing.T) {
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
		resp, err := cmd.CombinedOutput()
		a := assert.New(t)
		if a.NoError(err, string(resp)) {
			var key atlasv2.ApiKeyUserDetails
			if err := json.Unmarshal(resp, &key); err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			a.Equal(desc, *key.Desc)
			ID = *key.Id
		}
	})
	require.NotEmpty(t, ID)

	t.Run("List", func(t *testing.T) {
		cmd := exec.Command(cliPath,
			orgEntity,
			apiKeysEntity,
			"ls",
			"-o=json")
		cmd.Env = os.Environ()
		resp, err := cmd.CombinedOutput()
		require.NoError(t, err, string(resp))

		var keys atlasv2.PaginatedApiApiUser
		if err := json.Unmarshal(resp, &keys); err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		assert.NotEmpty(t, keys.Results)
	})

	t.Run("List Compact", func(t *testing.T) {
		cmd := exec.Command(cliPath,
			orgEntity,
			apiKeysEntity,
			"ls",
			"-c",
			"-o=json")
		cmd.Env = os.Environ()
		resp, err := cmd.CombinedOutput()
		require.NoError(t, err, string(resp))

		var keys []atlasv2.ApiKeyUserDetails
		if err := json.Unmarshal(resp, &keys); err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		assert.NotEmpty(t, keys)
	})

	t.Run("Update", func(t *testing.T) {
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
		resp, err := cmd.CombinedOutput()
		a := assert.New(t)
		if a.NoError(err, string(resp)) {
			var key atlasv2.ApiKeyUserDetails
			if err := json.Unmarshal(resp, &key); err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			a.Equal(newDesc, *key.Desc)
		}
	})

	t.Run("Describe", func(t *testing.T) {
		cmd := exec.Command(cliPath,
			orgEntity,
			apiKeysEntity,
			"describe",
			ID,
			"-o=json")
		cmd.Env = os.Environ()
		resp, err := cmd.CombinedOutput()

		a := assert.New(t)
		if a.NoError(err, string(resp)) {
			var key atlasv2.ApiKeyUserDetails
			if err := json.Unmarshal(resp, &key); err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			a.Equal(ID, *key.Id)
		}
	})

	t.Run("Delete", func(t *testing.T) {
		cmd := exec.Command(cliPath,
			orgEntity,
			apiKeysEntity,
			"rm",
			ID,
			"--force")
		cmd.Env = os.Environ()
		resp, err := cmd.CombinedOutput()

		a := assert.New(t)
		if a.NoError(err, string(resp)) {
			expected := fmt.Sprintf("API Key '%s' deleted\n", ID)
			a.Equal(expected, string(resp))
		}
	})
}
