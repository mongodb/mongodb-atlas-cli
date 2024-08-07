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

//go:build e2e || (atlas && plugin && install)

package atlas_test

import (
	"fmt"
	"os"
	"os/exec"
	"path"
	"testing"

	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/plugin"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/test/e2e"
	"github.com/stretchr/testify/require"
)

func generateTestPluginDirectory(directoryName string) (string, error) {
	defaultPluginDir, err := plugin.GetDefaultPluginDirectory()
	if err != nil {
		return "", err
	}
	directoryPath := path.Join(defaultPluginDir, directoryName)
	err = os.MkdirAll(directoryPath, 0755)
	if err != nil {
		return "", fmt.Errorf("error creating directory: %w", err)
	}
	return directoryPath, nil
}

func generateTestPlugin(directoryName string, binaryName string, manifestContent string) error {
	directoryPath, err := generateTestPluginDirectory(directoryName)

	if err != nil {
		return err
	}

	// Write manifest.yml file
	manifestFile, err := os.Create(path.Join(directoryPath, "/manifest.yml"))
	if err != nil {
		return fmt.Errorf("error creating manifest.yml: %w", err)
	}
	defer manifestFile.Close()

	_, err = manifestFile.WriteString(manifestContent)
	if err != nil {
		return fmt.Errorf("error writing to manifest.yml: %w", err)
	}

	// Create empty binary file
	binaryFile, err := os.Create(path.Join(directoryPath, binaryName))
	if err != nil {
		return fmt.Errorf("error creating binary file: %w", err)
	}
	defer binaryFile.Close()

	return nil
}

func TestPluginInstall(t *testing.T) {
	cliPath, err := e2e.AtlasCLIBin()
	require.NoError(t, err)

	runPluginInstallTest(t, cliPath, "Invalid version for plugin", true, examplePluginRepository+"@2.3.4.5.6")
	runPluginInstallTest(t, cliPath, "Plugin version does not exist", true, examplePluginRepository+"@300.200.100")
	runPluginInstallTest(t, cliPath, "Repository Values invalid", true, "invalid-repository")
	runPluginInstallTest(t, cliPath, "Plugin does not exist", true, "github-repo/does-not-exist")
	runPluginInstallTest(t, cliPath, "Install Successful", false, examplePluginRepository)
	runPluginInstallTest(t, cliPath, "Plugin already installed", true, examplePluginRepository)

	deleteAllPlugins(t)

	err = generateTestPlugin("testplugin", "binary", `name: testplugin
description: description
version: 1.2.3
binary: binary
commands:
    example:
        description: command with same name as plugin command`)
	require.NoError(t, err)
	runPluginInstallTest(t, cliPath, "Plugin with same command already installed", true, examplePluginRepository)

	deleteAllPlugins(t)

	err = generateTestPlugin("testplugin", "binary", `name: atlas-cli-plugin-example
description: description
version: 1.2.3
binary: binary
commands:
    testplugin:
        description: this is the a test command`)
	require.NoError(t, err)
	runPluginInstallTest(t, cliPath, "Plugin with same name already installed", true, examplePluginRepository)
}

func runPluginInstallTest(t *testing.T, cliPath string, testName string, requireError bool, pluginValue string) {
	t.Helper()
	t.Run(testName, func(t *testing.T) {
		cmd := exec.Command(cliPath,
			"plugin",
			"install",
			pluginValue)
		resp, err := e2e.RunAndGetStdOut(cmd)
		if requireError {
			require.Error(t, err, string(resp))
		} else {
			require.NoError(t, err, string(resp))
		}
	})
}
