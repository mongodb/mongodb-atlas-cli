// Copyright 2023 MongoDB Inc
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

package auditing

import (
	"encoding/json"
	"os"
	"os/exec"
	"testing"

	"github.com/mongodb/mongodb-atlas-cli/atlascli/test/internal"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	atlasv2 "go.mongodb.org/atlas-sdk/v20250312014/admin"
)

const (
	auditingEntity = "auditing"
)

func TestAuditing(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping test in short mode")
	}

	g := internal.NewAtlasE2ETestGenerator(t, internal.WithSnapshot())
	g.GenerateProject("auditing")
	cliPath, err := internal.AtlasCLIBin()
	require.NoError(t, err)

	g.Run("Describe", func(t *testing.T) { //nolint:thelper // g.Run replaces t.Run
		cmd := exec.Command(cliPath,
			auditingEntity,
			"describe",
			"--projectId", g.ProjectID,
			"-o=json",
			"-P",
			internal.ProfileName())
		cmd.Env = os.Environ()
		resp, err := internal.RunAndGetStdOut(cmd)
		require.NoError(t, err, string(resp))
		var setting *atlasv2.AuditLog
		require.NoError(t, json.Unmarshal(resp, &setting), string(resp))
	})

	g.Run("Update", func(t *testing.T) { //nolint:thelper // g.Run replaces t.Run
		cmd := exec.Command(cliPath,
			auditingEntity,
			"update",
			"--projectId", g.ProjectID,
			"--enabled",
			"--auditAuthorizationSuccess",
			"--auditFilter", "{\"atype\": \"authenticate\"}",
			"-o=json",
			"-P",
			internal.ProfileName())
		cmd.Env = os.Environ()
		resp, err := internal.RunAndGetStdOut(cmd)
		require.NoError(t, err, string(resp))
		var setting *atlasv2.AuditLog
		require.NoError(t, json.Unmarshal(resp, &setting), string(resp))
		assert.True(t, *setting.Enabled)
		assert.True(t, *setting.AuditAuthorizationSuccess)
		assert.JSONEq(t, "{\"atype\": \"authenticate\"}", *setting.AuditFilter)
	})

	g.Run("Update via file", func(t *testing.T) { //nolint:thelper // g.Run replaces t.Run
		cmd := exec.Command(cliPath,
			auditingEntity,
			"update",
			"--projectId", g.ProjectID,
			"--enabled",
			"--auditAuthorizationSuccess",
			"-f", "testdata/update_auditing.json",
			"-o=json",
			"-P",
			internal.ProfileName())
		cmd.Env = os.Environ()
		resp, err := internal.RunAndGetStdOut(cmd)
		require.NoError(t, err, string(resp))
		var setting *atlasv2.AuditLog
		require.NoError(t, json.Unmarshal(resp, &setting), string(resp))
	})
}
