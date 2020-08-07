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
	"fmt"
	"os"
	"os/exec"

	"github.com/mongodb/mongocli/e2e"
	"go.mongodb.org/atlas/mongodbatlas"
)

const (
	iamEntity             = "iam"
	orgEntity             = "orgs"
	projectEntity         = "projects"
	apiKeysEntity         = "apikeys"
	apiKeyWhitelistEntity = "whitelist"
	whitelistIP           = "93.144.26.147"
)

func createOrgAPIKey() (string, error) {
	cliPath, err := e2e.Bin()
	if err != nil {
		return "", err
	}

	cmd := exec.Command(cliPath, iamEntity,
		orgEntity,
		apiKeysEntity,
		"create",
		"--desc=e2e-test",
		"--role=ORG_READ_ONLY",
		"-o=json")
	cmd.Env = os.Environ()
	resp, err := cmd.CombinedOutput()

	var key mongodbatlas.APIKey
	if err := json.Unmarshal(resp, &key); err != nil {
		return "", err
	}

	if key.ID != "" {
		return key.ID, nil
	}

	return "", fmt.Errorf("the apiKey ID is empty")

}

func deleteOrgAPIKey(id string) error {
	cliPath, err := e2e.Bin()
	if err != nil {
		return err
	}
	cmd := exec.Command(cliPath,
		iamEntity,
		orgEntity,
		apiKeysEntity,
		"rm",
		id,
		"--force")
	cmd.Env = os.Environ()
	return cmd.Run()
}
