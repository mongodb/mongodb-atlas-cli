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
// +build e2e cloudmanager,generic

package cloud_manager_test

import (
	"encoding/json"
	"github.com/mongodb/mongocli/e2e"
	"go.mongodb.org/ops-manager/opsmngr"
	"os"
	"os/exec"
	"testing"
)

func TestServers(t *testing.T) {
	cliPath, err := e2e.Bin()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	const serversEntity = "servers"

	t.Run("List", func(t *testing.T) {
		cmd := exec.Command(cliPath,
			entity,
			serversEntity,
			"list",
		)

		cmd.Env = os.Environ()
		resp, err := cmd.CombinedOutput()

		if err != nil {
			t.Fatalf("unexpected error: %v, resp: %v\n", err, string(resp))
		}

		var servers *opsmngr.Agents
		if err = json.Unmarshal(resp, &servers); err != nil {
			t.Fatalf("unexpected error: %v\n", err)
		}
		if servers.TotalCount != 1 {
			t.Errorf("expected one server, got=%d\n", servers.TotalCount)
		}
	})
}
