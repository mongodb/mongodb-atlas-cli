// Copyright 2020 MongoDB Inc
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

package migration

import (
	"os"
	"os/exec"
	"path"
	"testing"

	"github.com/mongodb/mongodb-atlas-cli/atlascli/test/internal"
	"github.com/pelletier/go-toml"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/zalando/go-keyring"
)

func TestConfigMigration(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping test in short mode")
	}

	// Create a temporary config folder
	dir := internal.TempConfigFolder(t)
	t.Cleanup(func() {
		os.RemoveAll(dir)
	})

	// Initialize the keychain
	require.NoError(t, internal.InitKeychain(t))

	// Write the old config format that needs to be migrated
	configPath := path.Join(dir, "config.toml")
	err := os.WriteFile(configPath, []byte(`[e2e-migration]
  org_id = "test_id"
  public_api_key = "test_pub"
  private_api_key = "test_priv"
  service = "cloud"
`), 0600)
	require.NoError(t, err)

	// Get the CLI path
	cliPath, err := internal.AtlasCLIBin()
	require.NoError(t, err)

	// Run the CLI, any command will trigger the migration
	cmd := exec.Command(cliPath)
	cmd.Env = os.Environ()

	// Get the output
	outputBytes, err := cmd.CombinedOutput()
	output := string(outputBytes)
	if err != nil {
		t.Fatalf("failed to run command: %v, output: %s", err, output)
	}

	// Ensure we're not falling back to insecure storage
	assert.NotContains(t, output, "Warning: Secure storage is not available, falling back to insecure storage")

	// Read the config file and check that it contains the expected migrated structure
	configContent, err := os.ReadFile(configPath)
	if err != nil {
		t.Fatalf("failed to read config file: %v", err)
	}

	// Parse the config file as TOML
	var config map[string]any
	require.NoError(t, toml.Unmarshal(configContent, &config))

	// Get the profile that we expect to be migrated
	migrationProfile, ok := config["e2e-migration"].(map[string]any)
	if !ok {
		t.Fatalf("e2e-migration not found in config")
	}

	// Check that the profile was migrated correctly
	// Secrets should be removed from the profile
	assert.Equal(t, "test_id", migrationProfile["org_id"])
	assert.Nil(t, migrationProfile["public_api_key"])
	assert.Nil(t, migrationProfile["private_api_key"])
	assert.Equal(t, "cloud", migrationProfile["service"])

	// Check that the secrets are stored in the keychain
	publicAPIKey, err := keyring.Get("atlascli_e2e-migration", "public_api_key")
	require.NoError(t, err)
	assert.Equal(t, "test_pub", publicAPIKey)

	privateAPIKey, err := keyring.Get("atlascli_e2e-migration", "private_api_key")
	require.NoError(t, err)
	assert.Equal(t, "test_priv", privateAPIKey)
}
