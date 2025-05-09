// Copyright 2021 MongoDB Inc
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
//go:build e2e || e2eSnap || (atlas && livemigrations)

package livemigrations

import (
	"os"
	"os/exec"
	"testing"

	"github.com/mongodb/mongodb-atlas-cli/atlascli/test/internal"
	"github.com/stretchr/testify/require"
)

const (
	liveMigrationsEntity = "liveMigrations"
)

func TestLinkToken(t *testing.T) {
	g := internal.NewAtlasE2ETestGenerator(t, internal.WithSnapshot())
	cliPath, err := internal.AtlasCLIBin()
	r := require.New(t)
	r.NoError(err)

	// Cleanup, do a delete in case a token already exists from another run of the test
	t.Logf("Removing link-tokens from previous/parallel test runs...")
	deleteCleanup := exec.Command(cliPath,
		liveMigrationsEntity,
		"link",
		"delete",
		"--force")
	deleteCleanup.Env = os.Environ()
	if err := deleteCleanup.Run(); err == nil {
		t.Logf("Warning: Deleted link-token.")
	}
	t.Logf("Cleanup complete.")

	g.Run("Create", func(t *testing.T) { //nolint:thelper // g.Run replaces t.Run
		cmd := exec.Command(cliPath,
			liveMigrationsEntity,
			"link",
			"create",
			"--accessListIp",
			"1.2.3.4,5.6.7.8")
		cmd.Env = os.Environ()
		resp, err := internal.RunAndGetStdOut(cmd)

		require.NoError(t, err, string(resp))
	})

	g.Run("Delete", func(t *testing.T) { //nolint:thelper // g.Run replaces t.Run
		cmd := exec.Command(cliPath,
			liveMigrationsEntity,
			"link",
			"delete",
			"--force")
		cmd.Env = os.Environ()
		resp, err := internal.RunAndGetStdOut(cmd)

		require.NoError(t, err, string(resp))
	})
}
