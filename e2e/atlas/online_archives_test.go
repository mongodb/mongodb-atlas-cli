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
// +build e2e atlas,onlinearchive

package atlas_test

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"testing"

	"github.com/mongodb/mongocli/e2e"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/atlas/mongodbatlas"
)

func TestOnlineArchives(t *testing.T) {
	const onlineArchiveEntity = "onlineArchives"

	clusterName, err := deployCluster()
	if err != nil {
		t.Fatalf("failed to deploy a cluster: %v", err)
	}

	defer func() {
		if e := deleteCluster(clusterName); e != nil {
			t.Errorf("error deleting test cluster: %v", e)
		}
	}()

	cliPath, err := e2e.Bin()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	var archiveID string
	t.Run("Create", func(t *testing.T) {
		const dbName = "test"
		cmd := exec.Command(cliPath,
			atlasEntity,
			clustersEntity,
			onlineArchiveEntity,
			"create",
			"--clusterName="+clusterName,
			"--db="+dbName,
			"--collection=test",
			"--dateField=test",
			"--archiveAfter=3",
			"--partition=test:date",
			"-o=json")

		cmd.Env = os.Environ()
		resp, err := cmd.CombinedOutput()
		if err != nil {
			t.Fatalf("unexpected error: %v, resp: %v", err, string(resp))
		}

		var archive mongodbatlas.OnlineArchive
		if err = json.Unmarshal(resp, &archive); err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		assert.Equal(t, dbName, archive.DBName)
		archiveID = archive.ID
	})

	if archiveID == "" {
		t.Fatal("Failed to create archive")
	}

	t.Run("Describe", func(t *testing.T) {
		cmd := exec.Command(cliPath,
			atlasEntity,
			clustersEntity,
			onlineArchiveEntity,
			"describe",
			archiveID,
			"--clusterName="+clusterName,
			"-o=json")

		cmd.Env = os.Environ()
		resp, err := cmd.CombinedOutput()
		if err != nil {
			t.Fatalf("unexpected error: %v, resp: %v", err, string(resp))
		}

		var archive mongodbatlas.OnlineArchive
		if err = json.Unmarshal(resp, &archive); err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		assert.Equal(t, archiveID, archive.ID)
	})

	t.Run("list", func(t *testing.T) {
		cmd := exec.Command(cliPath,
			atlasEntity,
			clustersEntity,
			onlineArchiveEntity,
			"list",
			"--clusterName="+clusterName,
			"-o=json")

		cmd.Env = os.Environ()
		resp, err := cmd.CombinedOutput()
		if err != nil {
			t.Fatalf("unexpected error: %v, resp: %v", err, string(resp))
		}
		var archives []*mongodbatlas.OnlineArchive
		if err = json.Unmarshal(resp, &archives); err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		assert.NotEmpty(t, archives)
	})

	t.Run("Update", func(t *testing.T) {
		const expireAfterDays = float64(4)
		cmd := exec.Command(cliPath,
			atlasEntity,
			clustersEntity,
			onlineArchiveEntity,
			"update",
			archiveID,
			"--clusterName="+clusterName,
			"--archiveAfter="+fmt.Sprintf("%.0f", expireAfterDays),
			"-o=json")

		cmd.Env = os.Environ()
		resp, err := cmd.CombinedOutput()
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		var archive mongodbatlas.OnlineArchive
		if err = json.Unmarshal(resp, &archive); err != nil {
			t.Fatalf("unexpected error: %v, resp: %v", err, string(resp))
		}
		archiveID = archive.ID
		assert.Equal(t, expireAfterDays, archive.Criteria.ExpireAfterDays)
	})

	t.Run("Delete", func(t *testing.T) {
		cmd := exec.Command(cliPath,
			atlasEntity,
			clustersEntity,
			onlineArchiveEntity,
			"rm",
			archiveID,
			"--clusterName="+clusterName,
			"--force")

		cmd.Env = os.Environ()
		resp, err := cmd.CombinedOutput()
		if err != nil {
			t.Fatalf("unexpected error: %v, resp: %v", err, string(resp))
		}
		expected := fmt.Sprintf("Archive '%s' deleted\n", archiveID)
		assert.Equal(t, string(resp), expected)
	})
}
