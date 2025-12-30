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

package accesslists

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"testing"
	"time"

	"github.com/mongodb/mongodb-atlas-cli/atlascli/test/internal"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	atlasv2 "go.mongodb.org/atlas-sdk/v20250312011/admin"
)

const (
	accessListEntity = "accessList"
)

func TestAccessList(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping test in short mode")
	}

	g := internal.NewAtlasE2ETestGenerator(t, internal.WithSnapshot())
	g.GenerateProject("accessList")

	n := g.MemoryRand("rand", 255)
	req := require.New(t)

	entry := fmt.Sprintf("192.168.0.%d", n)
	currentIPEntry := ""

	cliPath, err := internal.AtlasCLIBin()
	req.NoError(err)

	g.Run("Create Forever", func(t *testing.T) { //nolint:thelper // g.Run replaces t.Run
		cmd := exec.Command(cliPath,
			accessListEntity,
			"create",
			entry,
			"--comment=test",
			"--projectId",
			g.ProjectID,
			"-o=json",
			"-P",
			internal.ProfileName(),
		)
		cmd.Env = os.Environ()
		resp, err := internal.RunAndGetStdOut(cmd)
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
			g.ProjectID,
			"-o=json",
			"-P",
			internal.ProfileName(),
		)
		cmd.Env = os.Environ()
		resp, err := internal.RunAndGetStdOut(cmd)
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
			g.ProjectID,
			"-o=json",
			"-P",
			internal.ProfileName(),
		)
		cmd.Env = os.Environ()
		resp, err := internal.RunAndGetStdOut(cmd)
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
			g.ProjectID,
			"--force",
			"-P",
			internal.ProfileName(),
		)
		cmd.Env = os.Environ()
		resp, err := internal.RunAndGetStdOut(cmd)
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
			g.ProjectID,
			"-o=json",
			"-P",
			internal.ProfileName(),
		)
		cmd.Env = os.Environ()
		resp, err := internal.RunAndGetStdOut(cmd)
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
			g.ProjectID,
			"--force",
			"-P",
			internal.ProfileName(),
		)
		cmd.Env = os.Environ()
		resp, err := internal.RunAndGetStdOut(cmd)
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
			g.ProjectID,
			"-o=json",
			"-P",
			internal.ProfileName(),
		)
		cmd.Env = os.Environ()
		resp, err := internal.RunAndGetStdOut(cmd)
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
			g.ProjectID,
			"--force",
			"-P",
			internal.ProfileName(),
		)
		cmd.Env = os.Environ()
		resp, err := internal.RunAndGetStdOut(cmd)
		require.NoError(t, err, string(resp))

		expected := fmt.Sprintf("Project access list entry '%s' deleted\n", currentIPEntry)
		assert.Equal(t, expected, string(resp))
	})
}
