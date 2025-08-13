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

//go:build e2e || e2eSnap || (atlas && interactive)

package setupfailure

import (
	"os"
	"os/exec"
	"testing"

	"github.com/mongodb/mongodb-atlas-cli/atlascli/test/internal"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const (
	setupEntity = "setup"
)

func TestSetupFailureFlow(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping test in short mode")
	}

	g := internal.NewAtlasE2ETestGenerator(t, internal.WithSnapshot(), internal.WithSnapshotSkip(internal.SkipSimilarSnapshots))
	g.GenerateProject("setup")
	cliPath, err := internal.AtlasCLIBin()
	req := require.New(t)
	req.NoError(err)

	g.Run("Invalid Public Key", func(t *testing.T) { //nolint:thelper // g.Run replaces t.Run
		t.Setenv("MONGODB_ATLAS_PUBLIC_API_KEY", "invalid_public_key")
		cmd := exec.Command(cliPath,
			setupEntity,
			"--skipMongosh",
			"--skipSampleData",
			"--force")
		cmd.Env = os.Environ()
		resp, err := cmd.CombinedOutput()
		req.Error(err)
		assert.Contains(t, string(resp), "this action requires authentication")
	})

	g.Run("Invalid Private Key", func(t *testing.T) { //nolint:thelper // g.Run replaces t.Run
		t.Setenv("MONGODB_ATLAS_PRIVATE_API_KEY", "invalid_private_key")
		cmd := exec.Command(cliPath,
			setupEntity,
			"--skipMongosh",
			"--skipSampleData",
			"--force")
		cmd.Env = os.Environ()
		resp, err := cmd.CombinedOutput()
		req.Error(err)
		assert.Contains(t, string(resp), "this action requires authentication")
	})

	g.Run("Invalid Project ID", func(t *testing.T) { //nolint:thelper // g.Run replaces t.Run
		// The invalid ProjectID should be 24 characters long, otherwise
		// an early error will be thrown about incorrect length.
		invalidProjectID := "111111111111111111111111"
		cmd := exec.Command(cliPath,
			setupEntity,
			"--skipMongosh",
			"--skipSampleData",
			"--projectId", invalidProjectID,
			"--force")
		cmd.Env = os.Environ()
		resp, err := cmd.CombinedOutput()
		req.Error(err)
		assert.Contains(t, string(resp), "GROUP_NOT_FOUND", "Expected GROUP_NOT_FOUND (invalid Project ID) error")
	})
}
