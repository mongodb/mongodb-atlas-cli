// Copyright 2024 MongoDB Inc
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

package atlas_test

import (
	"os"
	"os/exec"
	"testing"

	"github.com/mongodb/mongodb-atlas-cli/atlascli/test/e2e"
	"github.com/stretchr/testify/require"
)

func validateProfile(t *testing.T, cliPath string, profile string, profileValid bool) {
	t.Helper()

	// Setup the command
	cmd := exec.Command(cliPath,
		authEntity,
		whoami,
		"--profile", profile)

	cmd.Env = os.Environ()

	// Execute the command
	resp, err := e2e.RunAndGetStdOut(cmd)

	// We only care about errors that happened due to something other than a non-zero exit code
	if err != nil {
		require.ErrorContains(t, err, "exit status")
	}

	response := string(resp)

	if profileValid {
		require.NotContains(t, response, "Profile should not contain '.'")
	} else {
		require.NotContains(t, response, "not logged in", "Logged in as")
	}
}

func TestProfile(t *testing.T) {
	cliPath, err := e2e.AtlasCLIBin()
	require.NoError(t, err)

	t.Run("profile name valid", func(t *testing.T) {
		validateProfile(t, cliPath, "default", true)
		validateProfile(t, cliPath, "default-123", true)
		validateProfile(t, cliPath, "default-test", true)
	})

	t.Run("profile name invalid", func(t *testing.T) {
		validateProfile(t, cliPath, "d.efault", false)
		validateProfile(t, cliPath, "default.123", false)
		validateProfile(t, cliPath, "default.test", false)
	})
}
