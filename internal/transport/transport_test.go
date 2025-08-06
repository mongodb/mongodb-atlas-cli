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

//go:build unit

package transport

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/config"
	"github.com/stretchr/testify/require"
	"go.mongodb.org/atlas/auth"
)

func TestNewAccessTokenTransport(t *testing.T) {
	mockToken := &auth.Token{
		AccessToken:  "mock-access-token",
		RefreshToken: "mock-refresh-token",
	}

	saveToken := func(_ *auth.Token) error { return nil }

	base := Default()
	accessTokenTransport, err := NewAccessTokenTransport(mockToken, base, saveToken)
	require.NoError(t, err)
	require.NotNil(t, accessTokenTransport)

	req := httptest.NewRequest(http.MethodGet, "http://example.com", nil)
	resp, err := accessTokenTransport.RoundTrip(req)
	require.NoError(t, err)
	require.NotNil(t, resp)

	authHeader := req.Header.Get("Authorization")
	expectedHeader := "Bearer " + mockToken.AccessToken
	require.Equal(t, expectedHeader, authHeader)
}

func TestNewServiceAccountTransport(t *testing.T) {
	// Mock the token endpoint since the actual endpoint requires a valid client ID and secret.
	tokenServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		if _, err := w.Write([]byte(`{"access_token":"mock-token","token_type":"bearer","expires_in":3600}`)); err != nil {
			t.Errorf("Failed to write response: %v", err)
		}
	}))
	defer tokenServer.Close()

	// Temporarily set OpsManagerURL to mock tokenServer URL
	originalURL := config.OpsManagerURL()
	config.SetOpsManagerURL(tokenServer.URL + "/")
	defer func() { config.SetOpsManagerURL(originalURL) }()

	clientID := "mock-client-id"
	clientSecret := "mock-client-secret" //nolint:gosec

	client := NewServiceAccountTransport(clientID, clientSecret)
	require.NotNil(t, client)

	// Create request to check authentication header
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if got := r.Header.Get("Authorization"); got != "Bearer mock-token" {
			t.Errorf("Expected Authorization header to be 'Bearer mock-token', but got: %v", got)
		}
		w.WriteHeader(http.StatusOK)
	}))
	defer server.Close()

	req := httptest.NewRequest(http.MethodGet, server.URL, nil)
	resp, err := tr.RoundTrip(req)
	require.NoError(t, err)
	require.NotNil(t, resp)
}
