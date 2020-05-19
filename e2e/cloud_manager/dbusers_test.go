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

package cloud_manager_test

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

func TestCloudManagerDBUsers(t *testing.T) {
	cliPath, err := filepath.Abs("../../bin/mongocli")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	_, err = os.Stat(cliPath)

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	entity := "cloud-manager"
	dbusersEntity := "dbusers"
	username := fmt.Sprintf("user-%v", r.Uint32())

	t.Run("Create", func(t *testing.T) {
		cmd := exec.Command(cliPath,
			entity,
			dbusersEntity,
			"create",
			"--username="+username,
			"--password=passW0rd",
			"--role=readWriteAnyDatabase",
			"--mechanisms=SCRAM-SHA-256 ")
		cmd.Env = os.Environ()
		resp, err := cmd.CombinedOutput()

		if err != nil {
			t.Fatalf("unexpected error: %v, resp: %v", err, string(resp))
		}
	})

	t.Run("List", func(t *testing.T) {
		cmd := exec.Command(cliPath, entity, dbusersEntity, "ls")
		cmd.Env = os.Environ()
		resp, err := cmd.CombinedOutput()

		if err != nil {
			t.Fatalf("unexpected error: %v, resp: %v", err, string(resp))
		}

		users := []mongodbatlas.DatabaseUser{}
		err = json.Unmarshal(resp, &users)

		if len(users) == 0 {
			t.Fatalf("expected len(users) > 0, got 0")
		}

	})

	t.Run("Delete", func(t *testing.T) {
		cmd := exec.Command(cliPath, entity, dbusersEntity, "delete", username, "--force", "--authDB", "admin")
		cmd.Env = os.Environ()
		resp, err := cmd.CombinedOutput()

		if err != nil {
			t.Fatalf("unexpected error: %v, resp: %v", err, string(resp))
		}

		expected := "Changes are being applied, please check https://cloud-dev.mongodb.com/v2/5ec2839e74c5aa25f02ff8ee#deployment/topology for status\n"
		if string(resp) != expected {
			t.Errorf("got=%#v\nwant=%#v\n", string(resp), expected)
		}
	})

}
