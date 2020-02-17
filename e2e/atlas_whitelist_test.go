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

package e2e

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"testing"

	"github.com/mongodb/go-client-mongodb-atlas/mongodbatlas"
)

func TestAtlasWhitelist(t *testing.T) {
	cliPath, err := filepath.Abs("../bin/mcli")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	_, err = os.Stat(cliPath)

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	atlasEntity := "atlas"
	whitelistEntity := "whitelist"
	entry := "192.168.0.1"

	t.Run("Create", func(t *testing.T) {
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

		entries := make([]mongodbatlas.ProjectIPWhitelist, 1)
		err = json.Unmarshal(resp, &entries)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		if entries[0].IPAddress != entry {
			t.Errorf("got=%#v\nwant=%#v\n", entries[0].IPAddress, entry)
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
}
