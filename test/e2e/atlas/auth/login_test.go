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
	"bufio"
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/http/cookiejar"
	"net/url"
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

	authErrorCodeNone = "NONE"

	authLoginUsernameEnv           = "ATLAS_AUTH_LOGIN_E2E_USERNAME"
	authLoginPasswordEnv           = "ATLAS_AUTH_LOGIN_E2E_PASSWORD"
	authLoginRequireCredentialsEnv = "ATLAS_AUTH_LOGIN_E2E_REQUIRE_CREDENTIALS"
	authLoginProfile               = "auth-code-login"
	authCodeCallbackPath           = "/atlas-cli/callback"
)

var (
	authLoginUsername string
	authLoginPassword string
)

func TestMain(m *testing.M) {
	authLoginUsername = os.Getenv(authLoginUsernameEnv)
	authLoginPassword = os.Getenv(authLoginPasswordEnv)
	_ = os.Unsetenv(authLoginUsernameEnv)
	_ = os.Unsetenv(authLoginPasswordEnv)
	os.Exit(m.Run())
}

// isolateLoginEnv makes commands authenticate only with what login saved to the profile.
func isolateLoginEnv(t *testing.T) {
	t.Helper()

	// atlas-cli-core reads only MCLI_* variables when any is set.
	for _, entry := range os.Environ() {
		require.False(t, strings.HasPrefix(entry, "MCLI_"), "unset MCLI_* environment variables before running this test")
		key, _, _ := strings.Cut(entry, "=")
		if strings.HasPrefix(key, "MONGODB_ATLAS_AUTH_") {
			t.Setenv(key, "")
		}
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
	t.Setenv("MONGODB_ATLAS_ACCESS_TOKEN", "")
	t.Setenv("MONGODB_ATLAS_REFRESH_TOKEN", "")
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

type cloudDevAuthSession struct {
	client    *http.Client
	baseURL   string
	csrfToken string
	csrfTime  string
}

type authEndpointResponse struct {
	Status      string `json:"status"`
	ErrorCode   string `json:"errorCode"`
	CsrfToken   string `json:"csrfToken"`
	CsrfTime    string `json:"csrfTime"`
	PID         string `json:"pid"`
	RedirectURI string `json:"redirect_uri"`
}

func newCloudDevAuthSession(t *testing.T, baseURL, username, password string) *cloudDevAuthSession {
	t.Helper()

	jar, err := cookiejar.New(nil)
	require.NoError(t, err)
	client := &http.Client{
		Jar:     jar,
		Timeout: loginTimeout,
		CheckRedirect: func(_ *http.Request, _ []*http.Request) error {
			return http.ErrUseLastResponse
		},
	}
	endpoint, err := url.JoinPath(baseURL, "test/utils/user/v1/auth")
	require.NoError(t, err)
	requestBody, err := json.Marshal(map[string]string{
		"username": username,
		"password": password,
	})
	require.NoError(t, err)
	request, err := http.NewRequestWithContext(t.Context(), http.MethodPost, endpoint, bytes.NewReader(requestBody))
	require.NoError(t, err)
	request.Header.Set("Content-Type", "application/json")

	response, err := client.Do(request)
	require.NoError(t, err)
	defer func() { _ = response.Body.Close() }()
	require.Equal(t, http.StatusOK, response.StatusCode, "test-utils auth returned an unexpected HTTP status")

	var login authEndpointResponse
	require.NoError(t, json.NewDecoder(io.LimitReader(response.Body, 1<<20)).Decode(&login))
	require.Equal(t, authErrorCodeNone, login.ErrorCode, "test-utils auth failed with status %q", login.Status)
	require.NotEmpty(t, login.CsrfToken)
	require.NotEmpty(t, login.CsrfTime)

	return &cloudDevAuthSession{
		client:    client,
		baseURL:   baseURL,
		csrfToken: login.CsrfToken,
		csrfTime:  login.CsrfTime,
	}
}

func (s *cloudDevAuthSession) post(ctx context.Context, path string, query url.Values) (*http.Response, error) {
	endpoint, err := url.JoinPath(s.baseURL, path)
	if err != nil {
		return nil, err
	}
	if encoded := query.Encode(); encoded != "" {
		endpoint += "?" + encoded
	}
	request, err := http.NewRequestWithContext(ctx, http.MethodPost, endpoint, nil)
	if err != nil {
		return nil, errors.New("could not build cloud-dev auth request")
	}
	request.Header.Set("X-CSRF-Token", s.csrfToken)
	request.Header.Set("X-CSRF-Time", s.csrfTime)
	response, err := s.client.Do(request)
	if err != nil {
		return nil, errors.New("cloud-dev auth request failed")
	}
	return response, nil
}

func (s *cloudDevAuthSession) authorize(ctx context.Context, t *testing.T, authURL *url.URL) (string, string) {
	t.Helper()
	require.Equal(t, "cloud-dev.mongodb.com", authURL.Hostname(), "authorization endpoint is not cloud-dev")
	require.Equal(t, "/oauth/authorize", authURL.Path, "unexpected cloud-dev authorization endpoint")
	query := authURL.Query()
	state := query.Get("state")
	require.NotEmpty(t, state, "authorization URL has no state parameter")
	require.NotEmpty(t, query.Get("code_challenge"), "authorization URL has no PKCE challenge")
	require.Equal(t, "S256", query.Get("code_challenge_method"))

	response, err := s.post(ctx, "oauth/authorize", query)
	require.NoError(t, err)
	defer func() { _ = response.Body.Close() }()
	require.Equal(t, http.StatusOK, response.StatusCode, "preauthorization returned an unexpected HTTP status")

	var result authEndpointResponse
	require.NoError(t, json.NewDecoder(io.LimitReader(response.Body, 1<<20)).Decode(&result))
	require.True(t, result.ErrorCode == "" || result.ErrorCode == authErrorCodeNone, "preauthorization failed with error code %q", result.ErrorCode)
	require.NotEmpty(t, result.PID, "preauthorization returned no pid")
	return result.PID, state
}

func (s *cloudDevAuthSession) allowConsent(ctx context.Context, t *testing.T, pid string) string {
	t.Helper()
	response, err := s.post(ctx, "oauth/consent", url.Values{
		"consent": {"ALLOW"},
		"pid":     {pid},
	})
	require.NoError(t, err)
	defer func() { _ = response.Body.Close() }()
	require.Equal(t, http.StatusOK, response.StatusCode, "consent returned an unexpected HTTP status")

	var result authEndpointResponse
	require.NoError(t, json.NewDecoder(io.LimitReader(response.Body, 1<<20)).Decode(&result))
	require.True(t, result.ErrorCode == "" || result.ErrorCode == authErrorCodeNone, "consent failed with error code %q", result.ErrorCode)
	require.NotEmpty(t, result.RedirectURI, "consent returned no callback URL")
	return result.RedirectURI
}

func completeAuthCodeCallback(ctx context.Context, t *testing.T, callbackURL, requestedRedirectURI, expectedState string) {
	t.Helper()
	parsed, err := url.Parse(callbackURL)
	if err != nil {
		t.Fatal("consent returned an invalid callback URL")
	}
	requestedRedirect, err := url.Parse(requestedRedirectURI)
	if err != nil {
		t.Fatal("login requested an invalid redirect URI")
	}
	if requestedRedirect.Scheme != "http" || requestedRedirect.Hostname() != "127.0.0.1" || requestedRedirect.Path != authCodeCallbackPath || requestedRedirect.User != nil {
		t.Fatal("login requested an invalid loopback callback URI")
	}
	if parsed.Scheme != requestedRedirect.Scheme || parsed.Host != requestedRedirect.Host || parsed.Path != requestedRedirect.Path || parsed.User != nil {
		t.Fatal("consent callback URI does not match the requested redirect URI")
	}
	if parsed.Query().Get("state") != expectedState {
		t.Fatal("callback state did not match")
	}
	require.NotEmpty(t, parsed.Query().Get("code"))

	request, err := http.NewRequestWithContext(ctx, http.MethodGet, callbackURL, nil)
	if err != nil {
		t.Fatal("could not build loopback callback request")
	}
	client := &http.Client{
		Timeout: loginTimeout,
		CheckRedirect: func(_ *http.Request, _ []*http.Request) error {
			return http.ErrUseLastResponse
		},
	}
	response, err := client.Do(request)
	if err != nil {
		t.Fatal("loopback callback request failed")
	}
	defer func() { _ = response.Body.Close() }()
	require.GreaterOrEqual(t, response.StatusCode, http.StatusOK)
	require.Less(t, response.StatusCode, http.StatusBadRequest, "callback returned an error status")
	_, _ = io.Copy(io.Discard, response.Body)
}

func TestLoginUserAccountAuthCode(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping test in short mode")
	}

	username := strings.TrimSpace(authLoginUsername)
	password := authLoginPassword
	if username == "" || password == "" {
		if os.Getenv(authLoginRequireCredentialsEnv) == "true" {
			t.Fatal("Evergreen auth E2E requires username and password expansions")
		}
		t.Skipf("set Evergreen expansions for %s and %s to run the cloud-dev User Account E2E", authLoginUsernameEnv, authLoginPasswordEnv)
	}

	seed := readProfileSeed(t)
	opsManagerURL, err := url.Parse(seed.opsManagerURL)
	require.NoError(t, err)
	require.Equal(t, "https", opsManagerURL.Scheme, "this E2E requires HTTPS")
	require.Equal(t, "cloud-dev.mongodb.com", opsManagerURL.Hostname(), "this E2E only runs against cloud-dev")

	cliPath, err := internal.AtlasCLIBin()
	require.NoError(t, err)

	isolateLoginEnv(t)
	seedLoginEnv(t, seed)
	session := newCloudDevAuthSession(t, seed.opsManagerURL, username, password)
	logoutNeeded := false
	logout := func() ([]byte, error) {
		logoutCtx, cancel := context.WithTimeout(context.Background(), loginTimeout)
		defer cancel()
		cmd := exec.CommandContext(logoutCtx, cliPath, authEntity, "logout", "--force", "-P", authLoginProfile)
		cmd.Env = os.Environ()
		return cmd.CombinedOutput()
	}
	t.Cleanup(func() {
		if logoutNeeded {
			_, _ = logout()
		}
	})

	ctx, cancel := context.WithTimeout(t.Context(), loginTimeout)
	defer cancel()
	cmd := exec.CommandContext(ctx, cliPath,
		authEntity, loginEntity,
		"--authType", "UserAccount",
		"--output", "plaintext",
		"--projectId", seed.projectID,
		"--noBrowser",
		"-P", authLoginProfile,
	)
	cmd.Env = os.Environ()
	stdout, err := cmd.StdoutPipe()
	require.NoError(t, err)
	var stderr bytes.Buffer
	cmd.Stderr = &stderr
	logoutNeeded = true
	require.NoError(t, cmd.Start())
	defer func() {
		if cmd.Process != nil && cmd.ProcessState == nil {
			_ = cmd.Process.Kill()
			_ = cmd.Wait()
		}
	}()

	const authURLPrefix = "To continue, go to "
	scanner := bufio.NewScanner(stdout)
	var cliOutput strings.Builder
	var authURL string
	for scanner.Scan() {
		line := scanner.Text()
		if index := strings.Index(line, authURLPrefix); index >= 0 {
			authURL = strings.TrimSpace(line[index+len(authURLPrefix):])
			_, _ = cliOutput.WriteString(line[:index] + authURLPrefix + "[redacted]\n")
			break
		}
		_, _ = cliOutput.WriteString(line + "\n")
	}
	require.NoError(t, scanner.Err(), "reading login output")
	require.NotEmpty(t, authURL, "login did not print an authorization URL; output: %s", cliOutput.String())
	parsedAuthURL, err := url.Parse(authURL)
	if err != nil {
		t.Fatal("login printed an invalid authorization URL")
	}
	requestedRedirectURI := parsedAuthURL.Query().Get("redirect_uri")
	require.NotEmpty(t, requestedRedirectURI, "authorization URL has no redirect URI")

	pid, state := session.authorize(ctx, t, parsedAuthURL)
	callbackURL := session.allowConsent(ctx, t, pid)
	completeAuthCodeCallback(ctx, t, callbackURL, requestedRedirectURI, state)

	for scanner.Scan() {
		_, _ = cliOutput.WriteString(scanner.Text() + "\n")
	}
	require.NoError(t, scanner.Err(), "reading login output")
	loginErr := cmd.Wait()
	require.NoError(t, ctx.Err(), "login timed out; output: %s; stderr: %s", cliOutput.String(), stderr.String())
	require.NoError(t, loginErr, "login failed; output: %s; stderr: %s", cliOutput.String(), stderr.String())
	assert.Contains(t, cliOutput.String(), "Successfully logged in as")

	t.Run("profile is saved", func(t *testing.T) {
		cmd := exec.Command(cliPath, "config", "describe", authLoginProfile, "-o", "json")
		cmd.Env = os.Environ()
		resp, err := internal.RunAndGetStdOut(cmd)
		require.NoError(t, err, "config describe failed")

		var saved map[string]string
		require.NoError(t, json.Unmarshal(resp, &saved), "config describe returned invalid JSON")
		assert.Equal(t, "user_account", saved["auth_type"])
		assert.Equal(t, "plaintext", saved["output"])
		assert.Equal(t, seed.orgID, saved["org_id"])
		assert.Equal(t, seed.projectID, saved["project_id"])
	})

	t.Run("profile authenticates", func(t *testing.T) {
		cmd := exec.Command(cliPath, "projects", "ls", "-P", authLoginProfile)
		cmd.Env = os.Environ()
		_, err := internal.RunAndGetStdOut(cmd)
		require.NoError(t, err, "projects ls failed with the saved User Account profile")
	})

	logoutOutput, err := logout()
	require.NoError(t, err, string(logoutOutput))
	assert.NotContains(t, string(logoutOutput), "unable to revoke user account token")
	logoutNeeded = false
}

func TestLoginMultipleProjectsNoFlags(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping test in short mode")
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
