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

package store

import (
	"context"
	"encoding/base64"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"sync"
	"testing"
	"time"

	"github.com/mongodb/atlas-cli-core/config"
	"github.com/stretchr/testify/require"
	atlasauth "go.mongodb.org/atlas/auth"
)

type auth struct {
	username     string
	password     string
	refreshToken string
	clientID     string
	clientSecret string
	accessToken  *atlasauth.Token
}

func (a auth) Token() (*atlasauth.Token, error) {
	return a.accessToken, nil
}

func (a auth) RefreshToken() string {
	return a.refreshToken
}

func (a auth) PublicAPIKey() string {
	return a.username
}

func (a auth) PrivateAPIKey() string {
	return a.password
}

func (a auth) ClientID() string {
	return a.clientID
}

func (a auth) ClientSecret() string {
	return a.clientSecret
}

func (a auth) AuthType() config.AuthMechanism {
	if a.username != "" {
		return config.APIKeys
	}
	if a.accessToken != nil {
		return config.UserAccount
	}
	if a.clientID != "" {
		return config.ServiceAccount
	}
	return ""
}

var _ CredentialsGetter = &auth{}

func TestService(t *testing.T) {
	c, err := New(Service(config.CloudService))
	if err != nil {
		t.Fatalf("New() unexpected error: %v", err)
	}

	if c.service != config.CloudService {
		t.Errorf("New() service = %s; expected %s", c.service, "cloud")
	}
}

func TestWithBaseURL(t *testing.T) {
	c, err := New(Service(config.CloudService), WithBaseURL("http://test"))
	if err != nil {
		t.Fatalf("New() unexpected error: %v", err)
	}

	if c.baseURL != "http://test" {
		t.Errorf("New() baseURL = %s; expected %s", c.baseURL, "http://test")
	}
}

type testConfig struct {
	url string
	auth
}

func (testConfig) OpsManagerCACertificate() string {
	return ""
}

func (testConfig) OpsManagerSkipVerify() string {
	return "false"
}

func (testConfig) Service() string {
	return config.CloudService
}

func (c testConfig) OpsManagerURL() string {
	return c.url
}

var _ AuthenticatedConfig = &testConfig{}

func TestWithAuthentication(t *testing.T) {
	tests := []struct {
		name           string
		setTestProfile func(p *config.Profile)
	}{
		{
			name: "api keys",
			setTestProfile: func(p *config.Profile) {
				p.SetAuthType(config.APIKeys)
				p.SetPublicAPIKey("test-key")
				p.SetPrivateAPIKey("test-secret")
			},
		},
		{
			name: "service account",
			setTestProfile: func(p *config.Profile) {
				p.SetAuthType(config.ServiceAccount)
				p.SetClientID("id")
				p.SetClientSecret("secret")
			},
		},
		{
			name: "user account",
			setTestProfile: func(p *config.Profile) {
				p.SetAuthType(config.UserAccount)
				// Token() requires both tokens and a parseable JWT.
				claims := fmt.Sprintf(`{"sub":"user","exp":%d}`, time.Now().Add(time.Hour).Unix())
				jwt := base64.RawURLEncoding.EncodeToString([]byte(`{"alg":"none"}`)) + "." +
					base64.RawURLEncoding.EncodeToString([]byte(claims)) + "."
				p.SetAccessToken(jwt)
				p.SetRefreshToken("refresh")
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Set up profile for testing
			profile := config.NewProfile("test", config.NewInMemoryStore())
			tt.setTestProfile(profile)
			config.SetDefaultProfile(profile)

			c, err := New(Service("cloud"), WithAuthentication())
			require.NoError(t, err)
			require.NotNil(t, c.httpClient)
			require.NotNil(t, c.httpClient.Transport)
			require.NotEqual(t, c.transport(), c.httpClient.Transport) // Check transport is not default
		})
	}
}

func TestWithAuthentication_UserAccountLegacyRefreshUsesPersistedClientID(t *testing.T) {
	t.Setenv("MCLI_CLIENT_ID", "")
	var (
		mu               sync.Mutex
		gotClientID      string
		gotRefreshToken  string
		gotAuthorization string
	)
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case "/api/private/unauth/account/device/token":
			if err := r.ParseForm(); err != nil {
				t.Errorf("ParseForm() error = %v", err)
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
			mu.Lock()
			gotClientID = r.Form.Get("client_id")
			gotRefreshToken = r.Form.Get("refresh_token")
			mu.Unlock()
			w.Header().Set("Content-Type", "application/json")
			_, _ = io.WriteString(w, `{"access_token":"refreshed-access","refresh_token":"refreshed-refresh","token_type":"Bearer","expires_in":3600}`)
		case "/resource":
			mu.Lock()
			gotAuthorization = r.Header.Get("Authorization")
			mu.Unlock()
			w.WriteHeader(http.StatusNoContent)
		default:
			t.Errorf("unexpected request path %q", r.URL.Path)
			http.NotFound(w, r)
		}
	}))
	defer server.Close()

	previous := config.Default()
	profile := config.NewProfile("legacy", config.NewInMemoryStore())
	profile.SetService(config.CloudService)
	profile.SetAuthType(config.UserAccountLegacy)
	profile.SetOpsManagerURL("https://cloud-dev.mongodb.com/")
	profile.Set(config.AccountURLField, server.URL+"/")
	profile.SetClientID("0oadn4hoajpzxeSEy357")
	claims := fmt.Sprintf(`{"sub":"user","exp":%d}`, time.Now().Add(-time.Hour).Unix())
	accessToken := base64.RawURLEncoding.EncodeToString([]byte(`{"alg":"none"}`)) + "." +
		base64.RawURLEncoding.EncodeToString([]byte(claims)) + "."
	profile.SetAccessToken(accessToken)
	profile.SetRefreshToken("old-refresh")
	config.SetDefaultProfile(profile)
	t.Cleanup(func() { config.SetDefaultProfile(previous) })

	s, err := New(AuthenticatedPreset(profile))
	require.NoError(t, err)
	request, err := http.NewRequestWithContext(t.Context(), http.MethodGet, server.URL+"/resource", nil)
	require.NoError(t, err)
	response, err := s.httpClient.Do(request) //nolint:gosec // G704: URL comes from httptest.NewServer, not user input
	require.NoError(t, err)
	require.NoError(t, response.Body.Close())

	mu.Lock()
	clientID, refreshToken, authorization := gotClientID, gotRefreshToken, gotAuthorization
	mu.Unlock()
	require.Equal(t, "0oadn4hoajpzxeSEy357", clientID)
	require.Equal(t, "old-refresh", refreshToken)
	require.Equal(t, "Bearer refreshed-access", authorization)
	require.Equal(t, "refreshed-access", profile.AccessToken())
	require.Equal(t, "refreshed-refresh", profile.RefreshToken())
}

func TestWithContext(t *testing.T) {
	c, err := New(Service(config.CloudService))
	if err != nil {
		t.Fatalf("New() unexpected error: %v", err)
	}

	if c.ctx != context.Background() {
		t.Errorf("New() got %v; expected %v", c.ctx, t.Context())
	}

	type myCustomType string
	var k, v myCustomType = "custom key", "custom value"

	ctx := context.WithValue(t.Context(), k, v)

	c, err = New(Service(config.CloudService), WithContext(ctx))
	if err != nil {
		t.Fatalf("New() unexpected error: %v", err)
	}

	if c.ctx != ctx {
		t.Errorf("New() got %v; expected %v", c.ctx, ctx)
	}
}
