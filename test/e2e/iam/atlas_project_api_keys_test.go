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

func TestAtlasProjectAPIKeys(t *testing.T) {
	cliPath, err := e2e.AtlasCLIBin()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	var ID string

	// This test must run first to grab the ID of the project to later describe
	t.Run("Create", func(t *testing.T) {
		const desc = "e2e-test"
		cmd := exec.Command(cliPath,
			projectsEntity,
			apiKeysEntity,
			"create",
			"--desc",
			desc,
			"--role=GROUP_READ_ONLY",
			"-o=json")
		cmd.Env = os.Environ()
		resp, err := e2e.RunAndGetStdOut(cmd)
		a := assert.New(t)
		require.NoError(t, err, string(resp))
		var key atlasv2.ApiKeyUserDetails
		require.NoError(t, json.Unmarshal(resp, &key))
		a.Equal(desc, *key.Desc)
		ID = *key.Id
	})
	require.NotEmpty(t, ID)

	defer func() {
		if e := deleteOrgAPIKey(ID); e != nil {
			t.Errorf("error deleting test apikey: %v", e)
		}
	}()

	t.Run("Assign", func(t *testing.T) {
		cmd := exec.Command(cliPath,
			projectsEntity,
			apiKeysEntity,
			"assign",
			ID,
			"--role=GROUP_DATA_ACCESS_READ_ONLY",
			"-o=json")
		cmd.Env = os.Environ()
		resp, err := e2e.RunAndGetStdOut(cmd)
		require.NoError(t, err, string(resp))
	})

	t.Run("List", func(t *testing.T) {
		cmd := exec.Command(cliPath,
			projectsEntity,
			apiKeysEntity,
			"ls",
			"-o=json")
		cmd.Env = os.Environ()
		resp, err := e2e.RunAndGetStdOut(cmd)

		if err != nil {
			t.Fatalf("unexpected error: %v, resp: %v", err, string(resp))
		}
		var keys atlasv2.PaginatedApiApiUser
		if err := json.Unmarshal(resp, &keys); err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		assert.NotEmpty(t, keys.Results)
	})

	t.Run("List Compact", func(t *testing.T) {
		cmd := exec.Command(cliPath,
			projectsEntity,
			apiKeysEntity,
			"ls",
			"-c",
			"-o=json")
		cmd.Env = os.Environ()
		resp, err := e2e.RunAndGetStdOut(cmd)

		if err != nil {
			t.Fatalf("unexpected error: %v, resp: %v", err, string(resp))
		}
		var keys []atlasv2.ApiKeyUserDetails
		if err := json.Unmarshal(resp, &keys); err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		assert.NotEmpty(t, keys)
	})

	t.Run("Delete", func(t *testing.T) {
		cmd := exec.Command(cliPath,
			projectsEntity,
			apiKeysEntity,
			"rm",
			ID,
			"--force")
		cmd.Env = os.Environ()
		resp, err := e2e.RunAndGetStdOut(cmd)
		require.NoError(t, err, string(resp))
		expected := fmt.Sprintf("API Key '%s' deleted\n", ID)
		assert.Equal(t, expected, string(resp))
	})
}
