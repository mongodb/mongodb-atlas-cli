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

package e2e

import (
	"os/exec"
	"testing"

	"github.com/mongodb/mongodb-atlas-cli/atlascli/test/internal"
	"github.com/stretchr/testify/require"
)

func TestPluginUninstall(t *testing.T) {
	_ = internal.TempConfigFolder(t)

	g := internal.NewAtlasE2ETestGenerator(t, internal.WithSnapshot())
	cliPath, err := internal.AtlasCLIBin()
	require.NoError(t, err)

	runPluginUninstallTest(t, g, cliPath, "Uninstall Successful with repository values", false, examplePluginRepository)
	runPluginUninstallTest(t, g, cliPath, "Uninstall Successful with plugin name", false, examplePluginName)
	runPluginUninstallTest(t, g, cliPath, "Plugin could not be found", true, "invalid plugin")
}

func runPluginUninstallTest(t *testing.T, g *internal.AtlasE2ETestGenerator, cliPath string, testName string, requireError bool, pluginValue string) {
	internal.InstallExamplePlugin(t, cliPath, "latest")
	g.Run(testName, func(t *testing.T) { //nolint:thelper // g.Run replaces t.Run
		cmd := exec.Command(cliPath,
			"plugin",
			"uninstall",
			pluginValue)
		resp, err := internal.RunAndGetStdOut(cmd)
		if requireError {
			require.Error(t, err, string(resp))
		} else {
			require.NoError(t, err, string(resp))
		}
	})
	internal.DeleteAllPlugins(t)
}
