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
	"os"
	"os/exec"
	"testing"

	"go.mongodb.org/atlas/mongodbatlas"
)

func TestOnlineArchives(t *testing.T) {
	const clustersEntity = "clusters"
	const onlineArchiveEntity = "onlineArchives"

	clusterName, err := deployCluster()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	cliPath, err := cli()
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
		_, err := cmd.CombinedOutput()

		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
	})

	var archiveID string
	t.Run("Create", func(t *testing.T) {
		cmd := exec.Command(cliPath,
			atlasEntity,
			clustersEntity,
			onlineArchiveEntity,
			"create",
			"--clusterName="+clusterName,
			"--db=test",
			"--collection=test",
			"--dateField=test",
			"--archiveAfter=3",
			"--partition=test:date")

		cmd.Env = os.Environ()
		resp, err := cmd.CombinedOutput()
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		var archive mongodbatlas.OnlineArchive
		err = json.Unmarshal(resp, &archive)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		archiveID = archive.ID
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
			t.Fatalf("unexpected error: %v", err)
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
		_, err := cmd.CombinedOutput()
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
	})

	if err := deleteCluster(clusterName); err != nil {
		t.Fatalf("unexpected error: %s", err)
	}
}
