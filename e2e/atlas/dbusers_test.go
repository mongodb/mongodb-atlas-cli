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
// +build e2e

package atlas_test

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"path/filepath"
	"testing"
	"time"

	"github.com/mongodb/go-client-mongodb-atlas/mongodbatlas"
)

const (
	roleReadWrite = "readWrite"
)

func TestAtlasDBUsers(t *testing.T) {
	cliPath, err := filepath.Abs("../../bin/mongocli")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	_, err = os.Stat(cliPath)

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	atlasEntity := "atlas"
	dbusersEntity := "dbusers"
	username := fmt.Sprintf("user-%v", r.Uint32())

	t.Run("Create", func(t *testing.T) {
		cmd := exec.Command(cliPath,
			atlasEntity,
			dbusersEntity,
			"create",
			"atlasAdmin",
			"--username", username,
			"--password=passW0rd")
		cmd.Env = os.Environ()
		resp, err := cmd.CombinedOutput()

		if err != nil {
			t.Fatalf("unexpected error: %v, resp: %v", err, string(resp))
		}

		user := mongodbatlas.DatabaseUser{}
		err = json.Unmarshal(resp, &user)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		if user.Username != username {
			t.Errorf("got=%#v\nwant=%#v\n", user.Username, username)
		}
	})

	t.Run("List", func(t *testing.T) {
		cmd := exec.Command(cliPath, atlasEntity, dbusersEntity, "ls")
		cmd.Env = os.Environ()
		resp, err := cmd.CombinedOutput()

		if err != nil {
			t.Fatalf("unexpected error: %v, resp: %v", err, string(resp))
		}
	})

	t.Run("Update", func(t *testing.T) {
		cmd := exec.Command(cliPath,
			atlasEntity,
			dbusersEntity,
			"update",
			username,
			"--role",
			roleReadWrite)
		cmd.Env = os.Environ()
		resp, err := cmd.CombinedOutput()

		if err != nil {
			t.Fatalf("unexpected error: %v, resp: %v", err, string(resp))
		}

		user := mongodbatlas.DatabaseUser{}
		err = json.Unmarshal(resp, &user)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		if user.Username != username {
			t.Errorf("got=%#v\nwant=%#v\n", user.Username, username)
		}

		if len(user.Roles) != 1 {
			t.Errorf("len(user.Roles) got=%#v\nwant=%#v\n", len(user.Roles), 1)
		}

		if user.Roles[0].DatabaseName != "admin" {
			t.Errorf("got=%#v\nwant=%#v\n", "admin", user.Roles[0].DatabaseName)
		}

		if user.Roles[0].RoleName != roleReadWrite {
			t.Errorf("got=%#v\nwant=%#v\n", roleReadWrite, user.Roles[0].RoleName)
		}

	})

	t.Run("Delete", func(t *testing.T) {
		cmd := exec.Command(cliPath, atlasEntity, dbusersEntity, "delete", username, "--force", "--authDB", "admin")
		cmd.Env = os.Environ()
		resp, err := cmd.CombinedOutput()

		if err != nil {
			t.Fatalf("unexpected error: %v, resp: %v", err, string(resp))
		}

		expected := fmt.Sprintf("DB user '%s' deleted\n", username)
		if string(resp) != expected {
			t.Errorf("got=%#v\nwant=%#v\n", string(resp), expected)
		}
	})

}
