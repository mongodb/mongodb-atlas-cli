// Copyright 2024 MongoDB Inc
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

//go:build e2e || (atlas && plugin && uninstall)

package e2e_test

import (
	"os/exec"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPluginUninstall(t *testing.T) {
	_ = tempConfigFolder(t)

	g := newAtlasE2ETestGenerator(t, withSnapshot())
	cliPath, err := AtlasCLIBin()
	require.NoError(t, err)

	runPluginUninstallTest(g, cliPath, "Uninstall Successful with repository values", false, examplePluginRepository)
	runPluginUninstallTest(g, cliPath, "Uninstall Successful with plugin name", false, examplePluginName)
	runPluginUninstallTest(g, cliPath, "Plugin could not be found", true, "invalid plugin")
}

func runPluginUninstallTest(g *atlasE2ETestGenerator, cliPath string, testName string, requireError bool, pluginValue string) {
	g.t.Helper()
	installExamplePlugin(g.t, cliPath, "latest")
	g.Run(testName, func(t *testing.T) { //nolint:thelper // g.Run replaces t.Run
		cmd := exec.Command(cliPath,
			"plugin",
			"uninstall",
			pluginValue)
		resp, err := RunAndGetStdOut(cmd)
		if requireError {
			require.Error(t, err, string(resp))
		} else {
			require.NoError(t, err, string(resp))
		}
	})
	deleteAllPlugins(g.t)
}
