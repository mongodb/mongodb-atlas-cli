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
	"path/filepath"
	"testing"

	"go.mongodb.org/atlas/mongodbatlas"
)

func TestOrgs(t *testing.T) {
	cliPath, err := filepath.Abs("../../bin/mongocli")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	_, err = os.Stat(cliPath)

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	const iamEntity = "iam"
	const orgEntity = "orgs"

	var orgID string

	// This test must run first to grab the ID of the org to later describe
	t.Run("List", func(t *testing.T) {
		cmd := exec.Command(cliPath, iamEntity, orgEntity, "ls")
		cmd.Env = os.Environ()
		resp, err := cmd.CombinedOutput()

		if err != nil {
			t.Fatalf("unexpected error: %v, resp: %v", err, string(resp))
		}
		var orgs mongodbatlas.Organizations
		err = json.Unmarshal(resp, &orgs)

		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		if len(orgs.Results) == 0 {
			t.Errorf("got=%#v\nwant>0\n", len(orgs.Results))
		}
		orgID = orgs.Results[0].ID
	})

	t.Run("Describe", func(t *testing.T) {
		cmd := exec.Command(cliPath, iamEntity, orgEntity, "describe", orgID)
		cmd.Env = os.Environ()
		resp, err := cmd.CombinedOutput()

		if err != nil {
			t.Fatalf("unexpected error: %v, resp: %v", err, string(resp))
		}
	})
}
