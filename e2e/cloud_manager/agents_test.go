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
// +build e2e cloudmanager,remote

package cloud_manager_test

import (
	"encoding/json"
	"os"
	"os/exec"
	"testing"

	"github.com/mongodb/mongocli/e2e"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/ops-manager/opsmngr"
)

const (
	agentsEntity = "agents"
)

func TestAgents(t *testing.T) {
	cliPath, err := e2e.Bin()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	t.Run("List", func(t *testing.T) {
		cmd := exec.Command(cliPath,
			entity,
			agentsEntity,
			"list",
			"AUTOMATION",
		)

		cmd.Env = os.Environ()
		resp, err := cmd.CombinedOutput()
		a := assert.New(t)

		if a.NoError(err, string(resp)) {
			var servers *opsmngr.Agents
			err := json.Unmarshal(resp, &servers)
			a.NoError(err)
			a.NotZero(servers.TotalCount)
		}
	})
}
