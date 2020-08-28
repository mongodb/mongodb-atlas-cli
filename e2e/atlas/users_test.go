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
// +build e2e atlas,generic

package atlas_test

import (
	"encoding/json"
	"os"
	"os/exec"
	"testing"

	"github.com/mongodb/mongocli/e2e"
	"go.mongodb.org/atlas/mongodbatlas"
)

func TestUsers(t *testing.T) {
	cliPath, err := e2e.Bin()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	t.Run("List", func(t *testing.T) {
		cmd := exec.Command(cliPath,
			atlasEntity,
			iamEntity,
			projectsEntity,
			usersEntity,
			"list",
			"-o=json")
		cmd.Env = os.Environ()
		resp, err := cmd.CombinedOutput()

		if err != nil {
			t.Fatalf("unexpected error: %v, resp: %v", err, string(resp))
		}

		var users []mongodbatlas.AtlasUser
		if err := json.Unmarshal(resp, &users); err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if len(users) > 0 {
			t.Fatalf("expected len(users) == 0, got %v", len(users))
		}
	})

	t.Run("Describe", func(t *testing.T) {
		cmd := exec.Command(cliPath,
			atlasEntity,
			iamEntity,
			usersEntity,
			"describe",
			"nonExistentUser",
			"-o=json")
		cmd.Env = os.Environ()
		resp, err := cmd.CombinedOutput()

		if err != nil {
			t.Fatalf("unexpected error: %v, resp: %v", err, string(resp))
		}

		var user mongodbatlas.AtlasUser

		if err := json.Unmarshal(resp, &user); err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
	})
}
