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
	"fmt"
	"github.com/mongodb/mongocli/internal/cli/atlas/dbusers"
	"math/rand"
	"os"
	"os/exec"
	"testing"
	"time"

	"github.com/mongodb/mongocli/e2e"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/atlas/mongodbatlas"
)

func TestDBUserCerts(t *testing.T) {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	username := fmt.Sprintf("user-%v", r.Uint32())

	cliPath, err := e2e.Bin()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	t.Run("Create DBUser", func(t *testing.T) {
		cmd := exec.Command(cliPath,
			atlasEntity,
			dbusersEntity,
			"create",
			"atlasAdmin",
			"--username", username,
			"--x509Type", dbusers.X509TypeManaged)
		cmd.Env = os.Environ()
		resp, err := cmd.CombinedOutput()

		if err != nil {
			t.Fatalf("unexpected error: %v, resp: %v", err, string(resp))
		}

		var user mongodbatlas.DatabaseUser
		if err := json.Unmarshal(resp, &user); err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		assert.Equal(t, username, user.Username)
	})

	t.Run("Create", func(t *testing.T) {
		cmd := exec.Command(cliPath,
			atlasEntity,
			dbusersEntity,
			certsEntity,
			"create",
			"--username", username)
		cmd.Env = os.Environ()
		resp, err := cmd.CombinedOutput()

		if err != nil {
			t.Errorf("unexpected error: %v, resp: %v", err, string(resp))
		}

		var user mongodbatlas.UserCertificate
		if err := json.Unmarshal(resp, &user); err != nil {
			t.Errorf("unexpected error: %v", err)
		}
		assert.Equal(t, username, user.Username)
	})


	t.Run("List", func(t *testing.T) {
		cmd := exec.Command(cliPath, atlasEntity, dbusersEntity, certsEntity, "list", "--username", username)
		cmd.Env = os.Environ()
		resp, err := cmd.CombinedOutput()

		if err != nil {
			t.Fatalf("unexpected error: %v, resp: %v", err, string(resp))
		}

		var users []mongodbatlas.UserCertificate
		if err := json.Unmarshal(resp, &users); err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if len(users) == 0 {
			t.Fatalf("expected len(users) > 0, got 0")
		}
	})

	t.Run("Delete User", func(t *testing.T) {
		cmd := exec.Command(cliPath, atlasEntity, dbusersEntity, "delete", username, "--force", "--authDB", "admin")
		cmd.Env = os.Environ()
		resp, err := cmd.CombinedOutput()

		if err != nil {
			t.Fatalf("unexpected error: %v, resp: %v", err, string(resp))
		}

		expected := fmt.Sprintf("DB user '%s' deleted\n", username)
		assert.Equal(t, expected, string(resp))
	})
}
