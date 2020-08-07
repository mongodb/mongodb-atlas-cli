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
// +build e2e iam

package iam_test

import (
	"encoding/json"
	"os"
	"os/exec"
	"testing"

	"github.com/mongodb/mongocli/e2e"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/atlas/mongodbatlas"
)

func TestOrgAPIKeyWhitelist(t *testing.T) {
	cliPath, er := e2e.Bin()
	if er != nil {
		t.Fatalf("unexpected error: %v", er)
	}

	apiKeyID, e := createOrgAPIKey()
	if e != nil {
		t.Fatalf("unexpected error: %v", e)
	}

	t.Run("Create", func(t *testing.T) {
		cmd := exec.Command(cliPath, iamEntity,
			orgEntity,
			apiKeysEntity,
			apiKeyWhitelistEntity,
			"create",
			"--apiKey",
			apiKeyID,
			"--ip",
			whitelistIP,
			"-o=json")
		cmd.Env = os.Environ()
		resp, err := cmd.CombinedOutput()
		a := assert.New(t)
		if a.NoError(err, string(resp)) {
			var key mongodbatlas.WhitelistAPIKeys
			if err := json.Unmarshal(resp, &key); err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			a.NotEmpty(key.Results)
		}
	})

	t.Run("List", func(t *testing.T) {
		cmd := exec.Command(cliPath, iamEntity,
			orgEntity,
			apiKeysEntity,
			apiKeyWhitelistEntity,
			"list",
			apiKeyID,
			"-o=json")
		cmd.Env = os.Environ()
		resp, err := cmd.CombinedOutput()
		a := assert.New(t)
		if a.NoError(err, string(resp)) {
			var key mongodbatlas.WhitelistAPIKeys
			if err := json.Unmarshal(resp, &key); err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			a.NotEmpty(key.Results)
		}
	})

	t.Run("Delete", func(t *testing.T) {
		cmd := exec.Command(cliPath,
			iamEntity,
			orgEntity,
			apiKeysEntity,
			apiKeyWhitelistEntity,
			"rm",
			whitelistIP,
			"--apiKey",
			apiKeyID,
			"--force")
		cmd.Env = os.Environ()
		resp, err := cmd.CombinedOutput()
		assert.NoError(t, err, string(resp))
	})

	err := deleteOrgAPIKey(apiKeyID)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}
