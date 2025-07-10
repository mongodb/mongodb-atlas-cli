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
//go:build e2e || e2eSnap || brew

package brew

import (
	"os"
	"os/exec"
	"strings"
	"testing"

	"github.com/mongodb/mongodb-atlas-cli/atlascli/test/internal"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const (
	profileString = "PROFILE NAME"
	errorMessage  = "this action requires authentication"
)

func TestAtlasCLIConfig(t *testing.T) {
	_ = internal.TempConfigFolder(t)

	cliPath, err := internal.AtlasCLIBin()
	require.NoError(t, err)

	t.Setenv("MONGODB_ATLAS_ORG_ID", "")
	t.Setenv("MONGODB_ATLAS_PROJECT_ID", "")
	t.Setenv("MONGODB_ATLAS_PUBLIC_API_KEY", "")
	t.Setenv("MONGODB_ATLAS_PRIVATE_API_KEY", "")
	t.Setenv("MONGODB_ATLAS_OPS_MANAGER_URL", "")
	t.Setenv("MONGODB_ATLAS_SERVICE", "")

	t.Run("config ls", func(t *testing.T) {
		cmd := exec.Command(cliPath, "config", "ls")
		cmd.Env = os.Environ()
		resp, err := internal.RunAndGetStdOut(cmd)
		require.NoError(t, err, string(resp))
		got := strings.TrimSpace(string(resp))
		assert.Equal(t, profileString, got)
	})

	t.Run("projects ls", func(t *testing.T) {
		cmd := exec.Command(cliPath, "projects", "ls")
		cmd.Env = os.Environ()
		resp, err := cmd.CombinedOutput()
		got := string(resp)
		require.Error(t, err, got)
		assert.Contains(t, got, errorMessage)
	})

	t.Run("help", func(t *testing.T) {
		cmd := exec.Command(cliPath, "help")
		cmd.Env = os.Environ()
		resp, err := internal.RunAndGetStdOut(cmd)
		require.NoError(t, err, string(resp))
	})
}
