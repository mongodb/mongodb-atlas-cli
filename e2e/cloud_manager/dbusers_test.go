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
// +build e2e cloudmanager,remote,replica

package cloud_manager_test

import (
	"crypto/rand"
	"encoding/json"
	"fmt"
	"math/big"
	"os"
	"os/exec"
	"strings"
	"testing"

	"github.com/mongodb/mongocli/e2e"
	"go.mongodb.org/ops-manager/opsmngr"
)

func TestDBUsers(t *testing.T) {
	cliPath, err := e2e.Bin()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	n, err := rand.Int(rand.Reader, big.NewInt(1000))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	const dbUsersEntity = "dbusers"
	username := fmt.Sprintf("user-%v", n)

	t.Run("Create", func(t *testing.T) {
		cmd := exec.Command(cliPath,
			entity,
			dbUsersEntity,
			"create",
			"--username="+username,
			"--password=passW0rd",
			"--role=readWriteAnyDatabase",
			"--mechanisms=SCRAM-SHA-256")
		cmd.Env = os.Environ()
		resp, err := cmd.CombinedOutput()

		if err != nil {
			t.Fatalf("unexpected error: %v, resp: %v", err, string(resp))
		}

		if !strings.Contains(string(resp), "Changes are being applied") {
			t.Errorf("got=%#v\nwant=%#v\n", string(resp), "Changes are being applied")
		}
	})

	t.Run("List", func(t *testing.T) {
		cmd := exec.Command(cliPath,
			entity,
			dbUsersEntity,
			"ls",
			"-o=json")
		cmd.Env = os.Environ()
		resp, err := cmd.CombinedOutput()

		if err != nil {
			t.Fatalf("unexpected error: %v, resp: %v", err, string(resp))
		}

		var users []opsmngr.MongoDBUser
		if err := json.Unmarshal(resp, &users); err != nil {
			t.Log(string(resp))
			t.Fatalf("unexpected error: %v", err)
		}

		if len(users) == 0 {
			t.Fatalf("expected len(users) > 0, got 0")
		}
	})

	t.Run("Delete", func(t *testing.T) {
		cmd := exec.Command(cliPath,
			entity,
			dbUsersEntity,
			"delete",
			username,
			"--force",
			"--authDB",
			"admin",
		)
		cmd.Env = os.Environ()
		resp, err := cmd.CombinedOutput()

		if err != nil {
			t.Fatalf("unexpected error: %v, resp: %v", err, string(resp))
		}

		if !strings.Contains(string(resp), "Changes are being applied") {
			t.Errorf("got=%#v\nwant=%#v\n", string(resp), "Changes are being applied")
		}
	})
}
