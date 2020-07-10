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

	"github.com/go-test/deep"
	atlas "go.mongodb.org/atlas/mongodbatlas"
)

func TestDatalake(t *testing.T) {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	const datalakeEntity = "datalake"
	const updateRegion = "VIRGINIA_USA"
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
			t.Errorf("unexpected error: %v", err)
		}

		if err := deep.Equal(datalake.Name, datalakeName); err != nil {
			t.Error(err)
		}
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
			t.Errorf("unexpected error: %v, resp: %v", err, string(resp))
		}

		datalake := atlas.DataLake{}
		err = json.Unmarshal(resp, &datalake)
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}

		if datalake.Name != datalakeName {
			t.Errorf("expected name %v, got %v", datalakeName, datalake.Name)
		}
	})

	t.Run("List", func(t *testing.T) {
		cmd := exec.Command(cliPath, atlasEntity, clustersEntity, "ls")
		cmd.Env = os.Environ()
		resp, err := cmd.CombinedOutput()

		if err != nil {
			t.Errorf("unexpected error: %v, resp: %v", err, string(resp))
		}
	})

	t.Run("Delete", func(t *testing.T) {
		cmd := exec.Command(cliPath, atlasEntity, datalakeName, "delete", datalakeName)
		cmd.Env = os.Environ()

		err := cmd.Run()
		if err != nil {
			t.Errorf("unexpceted error: %v", err)
		}
	})

	t.Run("Update", func(t *testing.T) {
		cmd := exec.Command(cliPath,
			atlasEntity,
			datalakeEntity,
			"update",
			datalakeName,
			"--region",
			updateRegion)
		cmd.Env = os.Environ()
		resp, err := cmd.CombinedOutput()

		if err != nil {
			t.Errorf("unexpected error: %v, resp: %v", err, string(resp))
		}

		datalake := atlas.DataLake{}
		err = json.Unmarshal(resp, &datalake)
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}

		if err := deep.Equal(datalake.DataProcessRegion.Region, updateRegion); err != nil {
			t.Error(err)
		}
	})
}
