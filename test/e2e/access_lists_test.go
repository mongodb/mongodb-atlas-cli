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
//go:build e2e || (atlas && generic)

package e2e_test

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	atlasv2 "go.mongodb.org/atlas-sdk/v20250312001/admin"
)

func TestAccessList(t *testing.T) {
	g := newAtlasE2ETestGenerator(t, withSnapshot())
	g.generateProject("accessList")

	n := g.memoryRand("rand", 255)
	req := require.New(t)

	entry := fmt.Sprintf("192.168.0.%d", n)
	currentIPEntry := ""

	cliPath, err := AtlasCLIBin()
	req.NoError(err)

	g.Run("Create Forever", func(t *testing.T) { //nolint:thelper // g.Run replaces t.Run
		cmd := exec.Command(cliPath,
			accessListEntity,
			"create",
			entry,
			"--comment=test",
			"--projectId",
			g.projectID,
			"-o=json")
		cmd.Env = os.Environ()
		resp, err := RunAndGetStdOut(cmd)
		require.NoError(t, err)

		var entries *atlasv2.PaginatedNetworkAccess
		require.NoError(t, json.Unmarshal(resp, &entries))

		found := false
		for i := range entries.GetResults() {
			if entries.GetResults()[i].GetIpAddress() == entry {
				found = true
				break
			}
		}

		assert.True(t, found)
	})

	g.Run("List", func(t *testing.T) { //nolint:thelper // g.Run replaces t.Run
		cmd := exec.Command(cliPath,
			accessListEntity,
			"ls",
			"--projectId",
			g.projectID,
			"-o=json")
		cmd.Env = os.Environ()
		resp, err := RunAndGetStdOut(cmd)
		require.NoError(t, err, string(resp))
		var entries *atlasv2.PaginatedNetworkAccess
		require.NoError(t, json.Unmarshal(resp, &entries))
	})

	g.Run("Describe", func(t *testing.T) { //nolint:thelper // g.Run replaces t.Run
		cmd := exec.Command(cliPath,
			accessListEntity,
			"describe",
			entry,
			"--projectId",
			g.projectID,
			"-o=json")
		cmd.Env = os.Environ()
		resp, err := RunAndGetStdOut(cmd)
		require.NoError(t, err, string(resp))
		var entry *atlasv2.NetworkPermissionEntry
		require.NoError(t, json.Unmarshal(resp, &entry))
	})

	g.Run("Delete", func(t *testing.T) { //nolint:thelper // g.Run replaces t.Run
		cmd := exec.Command(cliPath,
			accessListEntity,
			"delete",
			entry,
			"--projectId",
			g.projectID,
			"--force")
		cmd.Env = os.Environ()
		resp, err := RunAndGetStdOut(cmd)
		require.NoError(t, err, string(resp))

		expected := fmt.Sprintf("Project access list entry '%s' deleted\n", entry)
		assert.Equal(t, expected, string(resp))
	})

	g.Run("Create Delete After", func(t *testing.T) { //nolint:thelper // g.Run replaces t.Run
		cmd := exec.Command(cliPath,
			accessListEntity,
			"create",
			entry,
			"--deleteAfter="+time.Now().Add(time.Minute*time.Duration(5)).Format(time.RFC3339),
			"--comment=test",
			"--projectId",
			g.projectID,
			"-o=json")
		cmd.Env = os.Environ()
		resp, err := RunAndGetStdOut(cmd)
		require.NoError(t, err, string(resp))

		var entries *atlasv2.PaginatedNetworkAccess
		require.NoError(t, json.Unmarshal(resp, &entries))

		found := false
		for i := range entries.GetResults() {
			if entries.GetResults()[i].GetIpAddress() == entry {
				found = true
				break
			}
		}
		assert.True(t, found)
	})

	g.Run("Delete", func(t *testing.T) { //nolint:thelper // g.Run replaces t.Run
		cmd := exec.Command(cliPath,
			accessListEntity,
			"delete",
			entry,
			"--projectId",
			g.projectID,
			"--force")
		cmd.Env = os.Environ()
		resp, err := RunAndGetStdOut(cmd)
		require.NoError(t, err, string(resp))

		expected := fmt.Sprintf("Project access list entry '%s' deleted\n", entry)
		assert.Equal(t, expected, string(resp))
	})

	g.Run("Create with CurrentIp", func(t *testing.T) { //nolint:thelper // g.Run replaces t.Run
		cmd := exec.Command(cliPath,
			accessListEntity,
			"create",
			"--currentIp",
			"--comment=test",
			"--projectId",
			g.projectID,
			"-o=json")
		cmd.Env = os.Environ()
		resp, err := RunAndGetStdOut(cmd)
		require.NoError(t, err, string(resp))

		var entries *atlasv2.PaginatedNetworkAccess
		require.NoError(t, json.Unmarshal(resp, &entries))

		a := assert.New(t)
		a.NotEmpty(entries.GetResults())
		a.Len(entries.GetResults(), 1)

		currentIPEntry = entries.GetResults()[0].GetIpAddress()
	})

	g.Run("Delete", func(t *testing.T) { //nolint:thelper // g.Run replaces t.Run
		cmd := exec.Command(cliPath,
			accessListEntity,
			"delete",
			currentIPEntry,
			"--projectId",
			g.projectID,
			"--force")
		cmd.Env = os.Environ()
		resp, err := RunAndGetStdOut(cmd)
		require.NoError(t, err, string(resp))

		expected := fmt.Sprintf("Project access list entry '%s' deleted\n", currentIPEntry)
		assert.Equal(t, expected, string(resp))
	})
}
