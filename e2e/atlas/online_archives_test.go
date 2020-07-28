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
	"strconv"
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

	t.Run("list", func(t *testing.T) {
		cmd := exec.Command(cliPath,
			atlasEntity,
			clustersEntity,
			onlineArchiveEntity,
			"list",
			"--clusterName="+clusterName)

		cmd.Env = os.Environ()
		resp, err := cmd.CombinedOutput()

		if err != nil {
			t.Fatalf("unexpected error: %v, resp: %v", err, string(resp))
		}
	})

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
			"--partition=test:date")

		cmd.Env = os.Environ()
		resp, err := cmd.CombinedOutput()
		if err != nil {
			t.Fatalf("unexpected error: %v, resp: %v", err, string(resp))
		}

		var archive mongodbatlas.OnlineArchive
		err = json.Unmarshal(resp, &archive)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		archiveID = archive.ID
		if archive.DBName != dbName {
			t.Errorf("got=%#v\nwant=%#v\n", archive.DBName, dbName)
		}
	})

	t.Run("Describe", func(t *testing.T) {
		cmd := exec.Command(cliPath,
			atlasEntity,
			clustersEntity,
			onlineArchiveEntity,
			"describe",
			archiveID,
			"--clusterName="+clusterName)

		cmd.Env = os.Environ()
		resp, err := cmd.CombinedOutput()
		if err != nil {
			t.Fatalf("unexpected error: %v, resp: %v", err, string(resp))
		}

		var archive mongodbatlas.OnlineArchive
		err = json.Unmarshal(resp, &archive)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if archiveID != archive.ID {
			t.Errorf("expected: %s, got: %s", archiveID, archive.ID)
		}
	})

	t.Run("Update", func(t *testing.T) {
		const expireAfterDays = 4
		cmd := exec.Command(cliPath,
			atlasEntity,
			clustersEntity,
			onlineArchiveEntity,
			"update",
			archiveID,
			"--clusterName="+clusterName,
			"--archiveAfter="+strconv.Itoa(expireAfterDays))

		cmd.Env = os.Environ()
		resp, err := cmd.CombinedOutput()
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		var archive mongodbatlas.OnlineArchive
		err = json.Unmarshal(resp, &archive)
		if err != nil {
			t.Fatalf("unexpected error: %v, resp: %v", err, string(resp))
		}
		archiveID = archive.ID
		if archive.Criteria.ExpireAfterDays != expireAfterDays {
			t.Errorf("got=%#v\nwant=%#v\n", archive.Criteria.ExpireAfterDays, expireAfterDays)
		}
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
