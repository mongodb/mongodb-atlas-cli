// Copyright 2025 MongoDB Inc
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

//go:build e2e || (atlas && plugin && kubernetes)

package e2e_test

import (
	"os"
	"os/exec"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestPluginKubernetes(t *testing.T) {
	_ = newAtlasE2ETestGenerator(t, withSnapshot())

	cliPath, err := AtlasCLIBin()
	require.NoError(t, err)

	t.Run("should install kubernetes plugin", func(t *testing.T) {
		removeFirstClassPlugin(t, "atlas-cli-plugin-kubernetes", cliPath)

		cmd := exec.Command(cliPath,
			"kubernetes")
		cmd.Env = os.Environ()
		resp, err := cmd.CombinedOutput()
		require.NoError(t, err, string(resp))
		assert.Contains(t, string(resp), "Plugin mongodb/atlas-cli-plugin-kubernetes successfully installed\n")
	})
}

func removeFirstClassPlugin(t *testing.T, name, cliPath string) {
	t.Helper()
	cmd := exec.Command(cliPath,
		"plugin",
		"uninstall",
		name)
	resp, err := cmd.CombinedOutput()
	if err != nil {
		require.Contains(t, string(resp), "Error: could not find plugin with name atlas-cli-plugin-kubernetes")
		return
	}
}
