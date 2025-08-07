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

package events

import (
	"encoding/json"
	"os"
	"os/exec"
	"testing"
	"time"

	"github.com/mongodb/mongodb-atlas-cli/atlascli/test/internal"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.mongodb.org/atlas-sdk/v20250312005/admin"
)

const (
	eventsEntity  = "events"
	projectEntity = "project"
	orgEntity     = "org"
)

func TestEvents(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping test in short mode")
	}

	g := internal.NewAtlasE2ETestGenerator(t, internal.WithSnapshot())
	cliPath, err := internal.AtlasCLIBin()
	require.NoError(t, err)
	g.Run("List Project Events", func(t *testing.T) { //nolint:thelper // g.Run replaces t.Run
		cmd := exec.Command(cliPath,
			eventsEntity,
			projectEntity,
			"list",
			"--omitCount",
			"-o=json",
			"-P",
			internal.ProfileName(),
		)

		cmd.Env = os.Environ()
		resp, err := internal.RunAndGetStdOut(cmd)
		require.NoError(t, err, string(resp))
		var events admin.GroupPaginatedEvent
		require.NoError(t, json.Unmarshal(resp, &events))
		assert.NotEmpty(t, events.GetResults())
	})

	g.Run("List Organization Events", func(t *testing.T) { //nolint:thelper // g.Run replaces t.Run
		cmd := exec.Command(cliPath,
			eventsEntity,
			orgEntity,
			"list",
			"--omitCount",
			"--minDate="+time.Now().Add(-time.Hour*time.Duration(24)).Format("2006-01-02T15:04:05-0700"),
			"-o=json",
			"-P",
			internal.ProfileName(),
		)

		cmd.Env = os.Environ()
		resp, err := internal.RunAndGetStdOut(cmd)
		require.NoError(t, err, string(resp))
		var events admin.OrgPaginatedEvent
		require.NoError(t, json.Unmarshal(resp, &events))
		assert.NotEmpty(t, events.GetResults())
	})
}
