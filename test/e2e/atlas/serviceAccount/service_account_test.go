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

package serviceaccount

import (
	"fmt"
	"os"
	"os/exec"
	"testing"
	"time"

	"github.com/mongodb/mongodb-atlas-cli/atlascli/test/internal"
	"github.com/stretchr/testify/require"
)

func TestCreateOrgServiceAccount(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping test in short mode")
	}

	mode, err := internal.TestRunMode()
	require.NoError(t, err)
	if mode != internal.TestModeLive {
		t.Skip("skipping test in snapshot mode")
	}

	cliPath, err := internal.AtlasCLIBin()
	require.NoError(t, err)

	// Create a new Service Account
	g := internal.NewAtlasE2ETestGenerator(t)
	name := "test-service-account"
	g.GenerateOrgServiceAccount(cliPath, name)

	// Set up a Service Account profile
	saProfileName := fmt.Sprintf("sa-test-%d", time.Now().Unix())
	err = setupServiceAccountProfile(t, cliPath, saProfileName, g.ClientID, g.ClientSecret)
	require.NoError(t, err)

	// Run some command with the Service Account profile
	cmd := exec.Command(cliPath, "projects", "ls", "-P", saProfileName)
	cmd.Env = os.Environ()

	resp, err := internal.RunAndGetStdOut(cmd)
	require.NoError(t, err, string(resp))
	t.Logf("Service account authentication successful: Command executed using service account profile. Command output: %s", string(resp))

	// Logout from the service account profile
	err = logoutServiceAccountProfile(t, cliPath, saProfileName)
	require.NoError(t, err)
}

func setupServiceAccountProfile(t *testing.T, cliPath, profileName, clientID, clientSecret string) error {
	t.Helper()

	// Set client ID
	cmd := exec.Command(cliPath, "config", "set", "client_id", clientID, "-P", profileName)
	cmd.Env = os.Environ()
	if _, err := internal.RunAndGetStdOut(cmd); err != nil {
		return fmt.Errorf("failed to set client_id: %w", err)
	}

	// Set client secret
	cmd = exec.Command(cliPath, "config", "set", "client_secret", clientSecret, "-P", profileName)
	cmd.Env = os.Environ()
	if _, err := internal.RunAndGetStdOut(cmd); err != nil {
		return fmt.Errorf("failed to set client_secret: %w", err)
	}

	// Set service
	cmd = exec.Command(cliPath, "config", "set", "service", "cloud", "-P", profileName)
	cmd.Env = os.Environ()
	if _, err := internal.RunAndGetStdOut(cmd); err != nil {
		return fmt.Errorf("failed to set service: %w", err)
	}

	// Set ops manager URL
	opsManagerURL := "https://cloud-dev.mongodb.com/"
	cmd = exec.Command(cliPath, "config", "set", "ops_manager_url", opsManagerURL, "-P", profileName)
	cmd.Env = os.Environ()
	if _, err := internal.RunAndGetStdOut(cmd); err != nil {
		return fmt.Errorf("failed to set ops_manager_url: %w", err)
	}

	return nil
}

func logoutServiceAccountProfile(t *testing.T, cliPath, profileName string) error {
	t.Helper()

	cmd := exec.Command(cliPath, "auth", "logout", "--force", "-P", profileName)
	cmd.Env = os.Environ()
	if _, err := internal.RunAndGetStdOut(cmd); err != nil {
		return fmt.Errorf("failed to logout service account: %w", err)
	}

	return nil
}
