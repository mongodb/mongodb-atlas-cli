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

//go:build e2e || e2eSnap || (atlas && onlinearchive)

package onlinearchives

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"testing"

	"github.com/mongodb/mongodb-atlas-cli/atlascli/test/internal"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	atlasv2 "go.mongodb.org/atlas-sdk/v20250312002/admin"
)

const (
	clustersEntity      = "clusters"
	onlineArchiveEntity = "onlineArchives"
)

func TestOnlineArchives(t *testing.T) {
	g := internal.NewAtlasE2ETestGenerator(t, internal.WithSnapshot())
	g.GenerateProjectAndCluster("onlineArchives")

	cliPath, err := internal.AtlasCLIBin()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	var archiveID string
	g.Run("Create", func(t *testing.T) { //nolint:thelper // g.Run replaces t.Run
		archiveID = createOnlineArchive(t, cliPath, g.ProjectID, g.ClusterName)
	})

	if archiveID == "" {
		t.Fatal("Failed to create archive")
	}

	g.Run("Describe", func(t *testing.T) { //nolint:thelper // g.Run replaces t.Run
		describeOnlineArchive(t, cliPath, g.ProjectID, g.ClusterName, archiveID)
	})

	g.Run("List", func(t *testing.T) { //nolint:thelper // g.Run replaces t.Run
		listOnlineArchives(t, cliPath, g.ProjectID, g.ClusterName)
	})

	g.Run("Pause", func(t *testing.T) { //nolint:thelper // g.Run replaces t.Run
		pauseOnlineArchive(t, cliPath, g.ProjectID, g.ClusterName, archiveID)
	})

	g.Run("Start", func(t *testing.T) { //nolint:thelper // g.Run replaces t.Run
		startOnlineArchive(t, cliPath, g.ProjectID, g.ClusterName, archiveID)
	})

	g.Run("Update", func(t *testing.T) { //nolint:thelper // g.Run replaces t.Run
		updateOnlineArchive(t, cliPath, g.ProjectID, g.ClusterName, archiveID)
	})

	g.Run("Delete", func(t *testing.T) { //nolint:thelper // g.Run replaces t.Run
		deleteOnlineArchive(t, cliPath, g.ProjectID, g.ClusterName, archiveID)
	})

	g.Run("Watch", func(t *testing.T) { //nolint:thelper // g.Run replaces t.Run
		watchOnlineArchive(t, cliPath, g.ProjectID, g.ClusterName, archiveID)
	})
}

func deleteOnlineArchive(t *testing.T, cliPath, projectID, clusterName, archiveID string) {
	t.Helper()
	cmd := exec.Command(cliPath,
		clustersEntity,
		onlineArchiveEntity,
		"rm",
		archiveID,
		"--clusterName", clusterName,
		"--projectId", projectID,
		"--force")

	cmd.Env = os.Environ()
	resp, err := internal.RunAndGetStdOut(cmd)
	require.NoError(t, err, string(resp))
	expected := fmt.Sprintf("Archive '%s' deleted\n", archiveID)
	assert.Equal(t, expected, string(resp))
}

func watchOnlineArchive(t *testing.T, cliPath, projectID, clusterName, archiveID string) {
	t.Helper()
	cmd := exec.Command(cliPath,
		clustersEntity,
		onlineArchiveEntity,
		"watch",
		archiveID,
		"--clusterName", clusterName,
		"--projectId", projectID,
	)
	cmd.Env = os.Environ()
	_ = cmd.Run()
}

func startOnlineArchive(t *testing.T, cliPath, projectID, clusterName, archiveID string) {
	t.Helper()
	cmd := exec.Command(cliPath,
		clustersEntity,
		onlineArchiveEntity,
		"start",
		archiveID,
		"--clusterName", clusterName,
		"--projectId", projectID,
		"-o=json")

	cmd.Env = os.Environ()
	resp, err := cmd.CombinedOutput()
	// online archive never reaches goal state as the db and collection must exist
	const expectedError = "ONLINE_ARCHIVE_CANNOT_MODIFY_FIELD"
	if err != nil && !strings.Contains(string(resp), expectedError) {
		t.Fatalf("unexpected error: %v, resp: %v", err, string(resp))
	}
}

func pauseOnlineArchive(t *testing.T, cliPath, projectID, clusterName, archiveID string) {
	t.Helper()
	cmd := exec.Command(cliPath,
		clustersEntity,
		onlineArchiveEntity,
		"pause",
		archiveID,
		"--clusterName", clusterName,
		"--projectId", projectID,
		"-o=json")

	cmd.Env = os.Environ()
	resp, err := cmd.CombinedOutput()
	// online archive never reaches goal state as the db and collection must exist
	const expectedError = "ONLINE_ARCHIVE_MUST_BE_ACTIVE_TO_PAUSE"
	if err != nil && !strings.Contains(string(resp), expectedError) {
		t.Fatalf("unexpected error: %v, resp: %v", err, string(resp))
	}
}

func updateOnlineArchive(t *testing.T, cliPath, projectID, clusterName, archiveID string) {
	t.Helper()
	const expireAfterDays = 4
	expireAfterDaysStr := strconv.Itoa(expireAfterDays)
	cmd := exec.Command(cliPath,
		clustersEntity,
		onlineArchiveEntity,
		"update",
		archiveID,
		"--clusterName", clusterName,
		"--projectId", projectID,
		"--archiveAfter", expireAfterDaysStr,
		"-o=json")

	cmd.Env = os.Environ()
	resp, err := internal.RunAndGetStdOut(cmd)
	if err != nil {
		t.Fatalf("unexpected error: %v, resp: %v", err, string(resp))
	}
	var archive atlasv2.BackupOnlineArchive
	if err = json.Unmarshal(resp, &archive); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	assert.Equal(t, expireAfterDays, archive.Criteria.GetExpireAfterDays())
}

func describeOnlineArchive(t *testing.T, cliPath, projectID, clusterName, archiveID string) {
	t.Helper()
	cmd := exec.Command(cliPath,
		clustersEntity,
		onlineArchiveEntity,
		"describe",
		archiveID,
		"--clusterName", clusterName,
		"--projectId", projectID,
		"-o=json")

	cmd.Env = os.Environ()
	resp, err := internal.RunAndGetStdOut(cmd)
	if err != nil {
		t.Fatalf("unexpected error: %v, resp: %v", err, string(resp))
	}

	var archive atlasv2.BackupOnlineArchive
	if err = json.Unmarshal(resp, &archive); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	assert.Equal(t, archiveID, archive.GetId())
}

func listOnlineArchives(t *testing.T, cliPath, projectID, clusterName string) {
	t.Helper()
	cmd := exec.Command(cliPath,
		clustersEntity,
		onlineArchiveEntity,
		"list",
		"--clusterName", clusterName,
		"--projectId", projectID,
		"-o=json")

	cmd.Env = os.Environ()
	resp, err := internal.RunAndGetStdOut(cmd)
	if err != nil {
		t.Fatalf("unexpected error: %v, resp: %v", err, string(resp))
	}
	var archives *atlasv2.PaginatedOnlineArchive
	if err = json.Unmarshal(resp, &archives); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	assert.NotEmpty(t, archives)
}

func createOnlineArchive(t *testing.T, cliPath, projectID, clusterName string) string {
	t.Helper()
	const dbName = "test"
	cmd := exec.Command(cliPath,
		clustersEntity,
		onlineArchiveEntity,
		"create",
		"--clusterName", clusterName,
		"--db", dbName,
		"--collection=test",
		"--dateField=test",
		"--archiveAfter=3",
		"--partition=test",
		"--projectId", projectID,
		"-o=json")

	cmd.Env = os.Environ()
	resp, err := internal.RunAndGetStdOut(cmd)
	if err != nil {
		t.Fatalf("unexpected error: %v, resp: %v", err, string(resp))
	}

	var archive atlasv2.BackupOnlineArchive
	if err = json.Unmarshal(resp, &archive); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	assert.Equal(t, dbName, archive.GetDbName())
	return archive.GetId()
}
