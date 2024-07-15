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
//go:build e2e || brew

package brew_test

import (
	"os"
	"os/exec"
	"strings"
	"testing"

	"github.com/mongodb/mongodb-atlas-cli/atlascli/test/e2e"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const (
	profileString = "PROFILE NAME"
	errorMessage  = "Error: this action requires authentication"
)

func TestAtlasCLIConfig(t *testing.T) {
	cliPath, err := e2e.AtlasCLIBin()
	require.NoError(t, err)
	tempDirEnv := "XDG_CONFIG_HOME=" + os.TempDir() // make sure no config.toml is detected

	t.Run("config ls", func(t *testing.T) {
		cmd := exec.Command(cliPath, "config", "ls")
		cmd.Env = append(os.Environ(), tempDirEnv)
		resp, err := e2e.RunAndGetStdOut(cmd)
		require.NoError(t, err, string(resp))
		got := strings.TrimSpace(string(resp))
		assert.Equal(t, profileString, got)
	})

	t.Run("projects ls", func(t *testing.T) {
		cmd := exec.Command(cliPath, "projects", "ls")
		cmd.Env = append(os.Environ(), tempDirEnv)
		resp, err := cmd.CombinedOutput()
		got := string(resp)
		require.Error(t, err, got)
		assert.Contains(t, got, errorMessage)
	})

	t.Run("help", func(t *testing.T) {
		cmd := exec.Command(cliPath, "help")
		cmd.Env = append(os.Environ(), tempDirEnv)
		resp, err := e2e.RunAndGetStdOut(cmd)
		require.NoError(t, err, string(resp))
	})
}
