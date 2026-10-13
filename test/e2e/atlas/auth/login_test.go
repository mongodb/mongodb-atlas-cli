// Copyright 2026 MongoDB Inc
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

package auth

import (
	"context"
	"encoding/json"
	"os"
	"os/exec"
	"runtime"
	"strings"
	"testing"
	"time"

	"github.com/mongodb/mongodb-atlas-cli/atlascli/test/internal"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const (
	authEntity  = "auth"
	loginEntity = "login"

	loginTimeout = 2 * time.Minute
)

// isolateLoginEnv makes commands authenticate only with what login saved to the profile.
func isolateLoginEnv(t *testing.T) {
	t.Helper()

	// atlas-cli-core reads only MCLI_* variables when any is set.
	for _, entry := range os.Environ() {
		require.False(t, strings.HasPrefix(entry, "MCLI_"), "unset MCLI_* environment variables before running this test")
	}

	_ = internal.TempConfigFolder(t)
	if runtime.GOOS == "darwin" {
		require.NoError(t, internal.InitKeychain(t))
	}

	t.Setenv("MONGODB_ATLAS_ORG_ID", "")
	t.Setenv("MONGODB_ATLAS_PROJECT_ID", "")
	t.Setenv("MONGODB_ATLAS_PUBLIC_API_KEY", "")
	t.Setenv("MONGODB_ATLAS_PRIVATE_API_KEY", "")
	t.Setenv("MONGODB_ATLAS_CLIENT_ID", "")
	t.Setenv("MONGODB_ATLAS_CLIENT_SECRET", "")
}

func runLogin(t *testing.T, cliPath string, args ...string) (string, error) {
	t.Helper()

	ctx, cancel := context.WithTimeout(t.Context(), loginTimeout)
	defer cancel()

	cmd := exec.CommandContext(ctx, cliPath, append([]string{authEntity, loginEntity}, args...)...)
	cmd.Env = os.Environ()
	resp, err := cmd.CombinedOutput()
	require.NoError(t, ctx.Err(), "login timed out, output:\n%s", string(resp))
	return string(resp), err
}

func TestLoginFlagValidation(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping test in short mode")
	}

	cliPath, err := internal.AtlasCLIBin()
	require.NoError(t, err)

	isolateLoginEnv(t)

	tests := []struct {
		name    string
		args    []string
		wantErr string
	}{
		{
			name:    "credentials without authType",
			args:    []string{"--clientId", "id", "--clientSecret", "secret", "--output", "plaintext"},
			wantErr: "--authType is required",
		},
		{
			name:    "authType without output",
			args:    []string{"--authType", "ServiceAccount", "--clientId", "id", "--clientSecret", "secret"},
			wantErr: "--output is required when --authType is set",
		},
		{
			name:    "service account without clientSecret",
			args:    []string{"--authType", "ServiceAccount", "--clientId", "id", "--output", "plaintext"},
			wantErr: "--clientId and --clientSecret are required when --authType is ServiceAccount",
		},
		{
			name:    "api keys without privateApiKey",
			args:    []string{"--authType", "APIKeys", "--publicApiKey", "pub", "--output", "plaintext"},
			wantErr: "--publicApiKey and --privateApiKey are required when --authType is APIKeys",
		},
		{
			name:    "invalid authType",
			args:    []string{"--authType", "Invalid", "--output", "plaintext"},
			wantErr: `invalid --authType "Invalid"`,
		},
		{
			name:    "invalid output",
			args:    []string{"--authType", "ServiceAccount", "--clientId", "id", "--clientSecret", "secret", "--output", "xml"},
			wantErr: `invalid --output "xml"`,
		},
		{
			name:    "invalid orgId",
			args:    []string{"--orgId", "nothex", "--authType", "ServiceAccount", "--clientId", "id", "--clientSecret", "secret", "--output", "plaintext"},
			wantErr: "invalid --orgId",
		},
		{
			name:    "invalid projectId",
			args:    []string{"--projectId", "nothex", "--authType", "ServiceAccount", "--clientId", "id", "--clientSecret", "secret", "--output", "plaintext"},
			wantErr: "invalid --projectId",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			out, err := runLogin(t, cliPath, tc.args...)
			require.Error(t, err, out)
			assert.Contains(t, out, tc.wantErr)
		})
	}
}

type profileSeed struct {
	opsManagerURL string
	orgID         string
	projectID     string
}

// readProfileSeed reads the e2e profile's backend settings; call it before isolateLoginEnv.
func readProfileSeed(t *testing.T) profileSeed {
	t.Helper()

	profileData, err := internal.ProfileData()
	require.NoError(t, err)

	seed := profileSeed{
		opsManagerURL: profileData["ops_manager_url"],
		orgID:         profileData["org_id"],
		projectID:     profileData["project_id"],
	}
	if seed.opsManagerURL == "" || seed.orgID == "" || seed.projectID == "" {
		t.Skipf("e2e profile %q lacks org_id, project_id or ops_manager_url", internal.ProfileName())
	}
	return seed
}

func seedLoginEnv(t *testing.T, seed profileSeed) {
	t.Helper()

	t.Setenv("MONGODB_ATLAS_OPS_MANAGER_URL", seed.opsManagerURL)
}

func TestLoginServiceAccount(t *testing.T) {
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

	g := internal.NewAtlasE2ETestGenerator(t)
	g.GenerateOrgServiceAccount(cliPath, "login-e2e")

	const profile = "sa-login"
	seed := readProfileSeed(t)

	isolateLoginEnv(t)
	seedLoginEnv(t, seed)

	out, err := runLogin(t, cliPath,
		"--authType", "ServiceAccount",
		"--clientId", g.ClientID,
		"--clientSecret", g.ClientSecret,
		"--output", "plaintext",
		"--orgId", seed.orgID,
		"--projectId", seed.projectID,
		"-P", profile,
	)
	require.NoError(t, err, out)

	t.Run("profile is saved", func(t *testing.T) {
		cmd := exec.Command(cliPath, "config", "describe", profile, "-o", "json")
		cmd.Env = os.Environ()
		resp, err := internal.RunAndGetStdOut(cmd)
		require.NoError(t, err, string(resp))

		var saved map[string]string
		require.NoError(t, json.Unmarshal(resp, &saved), string(resp))
		assert.Equal(t, "service_account", saved["auth_type"])
		assert.Equal(t, "plaintext", saved["output"])
		assert.NotEmpty(t, saved["client_id"])
		assert.NotEmpty(t, saved["client_secret"])
		assert.Equal(t, seed.orgID, saved["org_id"])
		assert.Equal(t, seed.projectID, saved["project_id"])
	})

	t.Run("profile authenticates", func(t *testing.T) {
		cmd := exec.Command(cliPath, "projects", "ls", "-P", profile)
		cmd.Env = os.Environ()
		resp, err := internal.RunAndGetStdOut(cmd)
		require.NoError(t, err, string(resp))
	})

	t.Run("logout", func(t *testing.T) {
		cmd := exec.Command(cliPath, authEntity, "logout", "--force", "-P", profile)
		cmd.Env = os.Environ()
		resp, err := internal.RunAndGetStdOut(cmd)
		require.NoError(t, err, string(resp))
	})
}

func TestLoginAPIKeys(t *testing.T) {
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

	const profile = "apikey-login"
	seed := readProfileSeed(t)

	// e2e-test-helper is swept by TestCleanup if this test's cleanup never runs.
	cmd := exec.Command(cliPath, "organizations", "apiKeys", "create",
		"--role", "ORG_OWNER",
		"--desc", "e2e-test-helper",
		"--orgId", seed.orgID,
		"-P", internal.ProfileName(),
		"-o", "json",
	)
	cmd.Env = os.Environ()
	resp, err := internal.RunAndGetStdOut(cmd)
	require.NoError(t, err, string(resp))

	var created struct {
		ID         string `json:"id"`
		PublicKey  string `json:"publicKey"`
		PrivateKey string `json:"privateKey"`
	}
	require.NoError(t, json.Unmarshal(resp, &created), string(resp))
	require.NotEmpty(t, created.ID, string(resp))
	t.Cleanup(func() {
		require.NoError(t, internal.DeleteOrgAPIKey(created.ID))
	})
	require.NotEmpty(t, created.PublicKey)
	require.NotEmpty(t, created.PrivateKey)

	isolateLoginEnv(t)
	seedLoginEnv(t, seed)

	out, err := runLogin(t, cliPath,
		"--authType", "APIKeys",
		"--publicApiKey", created.PublicKey,
		"--privateApiKey", created.PrivateKey,
		"--output", "plaintext",
		"--orgId", seed.orgID,
		"--projectId", seed.projectID,
		"-P", profile,
	)
	require.NoError(t, err, out)

	t.Run("profile is saved", func(t *testing.T) {
		cmd := exec.Command(cliPath, "config", "describe", profile, "-o", "json")
		cmd.Env = os.Environ()
		resp, err := internal.RunAndGetStdOut(cmd)
		require.NoError(t, err, string(resp))

		var saved map[string]string
		require.NoError(t, json.Unmarshal(resp, &saved), string(resp))
		assert.Equal(t, "api_keys", saved["auth_type"])
		assert.Equal(t, "plaintext", saved["output"])
		assert.NotEmpty(t, saved["public_api_key"])
		assert.NotEmpty(t, saved["private_api_key"])
		assert.Equal(t, seed.orgID, saved["org_id"])
		assert.Equal(t, seed.projectID, saved["project_id"])
	})

	t.Run("profile authenticates", func(t *testing.T) {
		cmd := exec.Command(cliPath, "projects", "ls", "-P", profile)
		cmd.Env = os.Environ()
		resp, err := internal.RunAndGetStdOut(cmd)
		require.NoError(t, err, string(resp))
	})

	t.Run("logout", func(t *testing.T) {
		cmd := exec.Command(cliPath, authEntity, "logout", "--force", "-P", profile)
		cmd.Env = os.Environ()
		resp, err := internal.RunAndGetStdOut(cmd)
		require.NoError(t, err, string(resp))
	})
}

func TestLoginMultipleProjectsNoFlags(t *testing.T) {
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

	g := internal.NewAtlasE2ETestGenerator(t)
	g.GenerateOrgServiceAccount(cliPath, "login-e2e")
	g.GenerateProject("login")

	const profile = "sa-login-multi"
	seed := readProfileSeed(t)

	isolateLoginEnv(t)
	seedLoginEnv(t, seed)

	out, err := runLogin(t, cliPath,
		"--authType", "ServiceAccount",
		"--clientId", g.ClientID,
		"--clientSecret", g.ClientSecret,
		"--output", "plaintext",
		"-P", profile,
	)
	require.NoError(t, err, out)
	assert.Contains(t, out, "No default project set")

	cmd := exec.Command(cliPath, "config", "describe", profile, "-o", "json")
	cmd.Env = os.Environ()
	resp, err := internal.RunAndGetStdOut(cmd)
	require.NoError(t, err, string(resp))

	var saved map[string]string
	require.NoError(t, json.Unmarshal(resp, &saved), string(resp))
	assert.Equal(t, seed.orgID, saved["org_id"])
	assert.Empty(t, saved["project_id"])
}

func TestLoginInaccessibleOrg(t *testing.T) {
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

	g := internal.NewAtlasE2ETestGenerator(t)
	g.GenerateOrgServiceAccount(cliPath, "login-e2e")

	seed := readProfileSeed(t)

	isolateLoginEnv(t)
	seedLoginEnv(t, seed)

	out, err := runLogin(t, cliPath,
		"--authType", "ServiceAccount",
		"--clientId", g.ClientID,
		"--clientSecret", g.ClientSecret,
		"--output", "plaintext",
		"--orgId", "000000000000000000000000",
		"-P", "sa-login-bad-org",
	)
	require.Error(t, err, out)
	assert.Contains(t, out, "organization is inaccessible")
}

func TestLoginInvalidCredentials(t *testing.T) {
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

	seed := readProfileSeed(t)

	isolateLoginEnv(t)
	seedLoginEnv(t, seed)

	tests := []struct {
		name string
		args []string
	}{
		{
			name: "service account",
			args: []string{
				"--authType", "ServiceAccount",
				"--clientId", "mdb_sa_id_000000000000000000000000",
				"--clientSecret", "mdb_sa_sk_000000000000000000000000000000000000000",
			},
		},
		{
			name: "api keys",
			args: []string{
				"--authType", "APIKeys",
				"--publicApiKey", "invalid",
				"--privateApiKey", "00000000-0000-0000-0000-000000000000",
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			out, err := runLogin(t, cliPath, append(tc.args, "--output", "plaintext", "-P", "bad-creds")...)
			require.Error(t, err, out)
			assert.Contains(t, out, "unauthorized")
			assert.NotContains(t, out, "Do you want to enter")
		})
	}
}
