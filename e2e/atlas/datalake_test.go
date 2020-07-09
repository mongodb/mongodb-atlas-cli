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
// +build e2e atlas,datalake

package atlas_test

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"testing"
	"time"

	"github.com/magiconair/properties/assert"

	atlas "go.mongodb.org/atlas/mongodbatlas"
)

func TestDatalake(t *testing.T) {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	const datalakeEntity = "datalake"
	datalakeName := fmt.Sprintf("e2e-data-lake-%v", r.Uint32())

	cliPath, err := cli()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	t.Run("Create", func(t *testing.T) {
		cmd := exec.Command(cliPath,
			atlasEntity,
			datalakeEntity,
			"create",
			datalakeName)
		cmd.Env = os.Environ()
		resp, err := cmd.CombinedOutput()

		if err != nil {
			t.Fatalf("unexpected error: %v, resp: %v", err, string(resp))
		}

		datalake := atlas.DataLake{}
		err = json.Unmarshal(resp, &datalake)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		assert.Equal(t, datalake.Name, datalakeName)
	})

	t.Run("Describe", func(t *testing.T) {
		cmd := exec.Command(cliPath,
			atlasEntity,
			datalakeEntity,
			"describe",
			datalakeName)
		cmd.Env = os.Environ()
		resp, err := cmd.CombinedOutput()

		if err != nil {
			t.Fatalf("unexpected error: %v, resp: %v", err, string(resp))
		}

		datalake := atlas.DataLake{}
		err = json.Unmarshal(resp, &datalake)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		assert.Equal(t, datalake.Name, datalakeName)
	})

	t.Run("List", func(t *testing.T) {
		cmd := exec.Command(cliPath, atlasEntity, clustersEntity, "ls")
		cmd.Env = os.Environ()
		resp, err := cmd.CombinedOutput()

		if err != nil {
			t.Fatalf("unexpected error: %v, resp: %v", err, string(resp))
		}
	})

	t.Run("Delete", func(t *testing.T) {
		cmd := exec.Command(cliPath, atlasEntity, datalakeName, "delete", datalakeName)
		cmd.Env = os.Environ()

		resp, err := cmd.CombinedOutput()
		if err != nil {
			t.Fatalf("unexpected error: %v, resp: %v", err, string(resp))
		}

		var state *os.ProcessState
		state, err = cmd.Process.Wait()
		if err != nil {
			t.Fatalf("unexpceted error: %v", err)
		}

		assert.Equal(t, state.Success(), true)
	})
}
