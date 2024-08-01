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

package atlas_test

import (
	"os/exec"
	"testing"

	"github.com/mongodb/mongodb-atlas-cli/atlascli/test/e2e"
	"github.com/stretchr/testify/require"
)

const (
	examplePluginName = "atlas-cli-plugin-example"
)

func TestPluginUninstall(t *testing.T) {
	cliPath, err := e2e.AtlasCLIBin()
	require.NoError(t, err)

	runPluginUninstallTest(t, cliPath, "Uninstall Successful with repository values", false, examplePluginRepository)
	runPluginUninstallTest(t, cliPath, "Uninstall Successful with plugin name", false, examplePluginName)
	runPluginUninstallTest(t, cliPath, "Plugin could not be found", true, "invalid plugin")
}

func runPluginUninstallTest(t *testing.T, cliPath string, testName string, requireError bool, pluginValue string) {
	t.Helper()
	installExamplePlugin(t, cliPath)
	t.Run(testName, func(t *testing.T) {
		cmd := exec.Command(cliPath,
			"plugin",
			"uninstall",
			pluginValue)
		resp, err := e2e.RunAndGetStdOut(cmd)
		if requireError {
			require.Error(t, err, string(resp))
		} else {
			require.NoError(t, err, string(resp))
		}
	})
	err := deleteAllPlugins()
	require.NoError(t, err)
}

func installExamplePlugin(t *testing.T, cliPath string) {
	t.Helper()
	cmd := exec.Command(cliPath,
		"plugin",
		"install",
		examplePluginRepository)
	resp, err := e2e.RunAndGetStdOut(cmd)
	require.NoError(t, err, string(resp))
}
