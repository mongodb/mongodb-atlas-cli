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
	"math/rand"
	"os"
	"os/exec"
	"testing"
	"time"

	"github.com/mongodb/mongocli/e2e"
	"go.mongodb.org/atlas/mongodbatlas"
)

func TestWhitelist(t *testing.T) {
	const whitelistEntity = "whitelist"
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	entry := fmt.Sprintf("192.168.0.%d", r.Int63n(255))

	cliPath, err := e2e.Bin()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	t.Run("Create Forever", func(t *testing.T) {
		cmd := exec.Command(cliPath,
			atlasEntity,
			whitelistEntity,
			"create",
			entry,
			"--comment=test")
		cmd.Env = os.Environ()
		resp, err := cmd.CombinedOutput()

		if err != nil {
			t.Fatalf("unexpected error: %v, resp: %v", err, string(resp))
		}

		var entries []*mongodbatlas.ProjectIPWhitelist
		if err := json.Unmarshal(resp, &entries); err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		found := false
		for i := range entries {
			if entries[i].IPAddress == entry {
				found = true
				break
			}
		}
		if !found {
			t.Errorf("entry=%#v not found in %#v\n", entry, entries)
		}
	})

	t.Run("List", func(t *testing.T) {
		cmd := exec.Command(cliPath, atlasEntity, whitelistEntity, "ls")
		cmd.Env = os.Environ()
		resp, err := cmd.CombinedOutput()

		if err != nil {
			t.Fatalf("unexpected error: %v, resp: %v", err, string(resp))
		}
	})

	t.Run("Describe", func(t *testing.T) {
		cmd := exec.Command(cliPath, atlasEntity, whitelistEntity, "describe", entry)
		cmd.Env = os.Environ()
		resp, err := cmd.CombinedOutput()

		if err != nil {
			t.Fatalf("unexpected error: %v, resp: %v", err, string(resp))
		}
	})

	t.Run("Delete", func(t *testing.T) {
		cmd := exec.Command(cliPath, atlasEntity, whitelistEntity, "delete", entry, "--force")
		cmd.Env = os.Environ()
		resp, err := cmd.CombinedOutput()

		if err != nil {
			t.Fatalf("unexpected error: %v, resp: %v", err, string(resp))
		}

		expected := fmt.Sprintf("Project whitelist entry '%s' deleted\n", entry)
		if string(resp) != expected {
			t.Errorf("got=%#v\nwant=%#v\n", string(resp), expected)
		}
	})

	t.Run("Create Delete After", func(t *testing.T) {
		cmd := exec.Command(cliPath,
			atlasEntity,
			whitelistEntity,
			"create",
			entry,
			"--deleteAfter="+time.Now().Add(time.Minute*time.Duration(5)).Format(time.RFC3339),
			"--comment=test")
		cmd.Env = os.Environ()
		resp, err := cmd.CombinedOutput()

		if err != nil {
			t.Fatalf("unexpected error: %v, resp: %v", err, string(resp))
		}

		var entries []*mongodbatlas.ProjectIPWhitelist
		if err := json.Unmarshal(resp, &entries); err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		found := false
		for i := range entries {
			if entries[i].IPAddress == entry {
				found = true
				break
			}
		}
		if !found {
			t.Errorf("entry=%#v not found in %#v\n", entry, entries)
		}
	})
}
