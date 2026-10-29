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
	"bytes"
	"encoding/base64"
	"encoding/json"
	"os"
	"testing"

	"github.com/mongodb/atlas-cli-core/config"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/log"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const (
	oktaIssuer       = "https://auth.mongodb.com/oauth2/default"
	authServerIssuer = "https://authorize.mongodb.com"
)

func TestWarnOnTokenAuthTypeMismatch(t *testing.T) {
	tests := []struct {
		name         string
		authType     config.AuthMechanism
		service      string
		envIssuer    string
		storedIssuer string
		accessToken  string
		refreshToken string
		want         string
	}{
		{
			name:         "no auth type",
			accessToken:  issuerToken(t, oktaIssuer),
			refreshToken: "refresh",
			want:         missingAuthTypeWarning,
		},
		{
			name:         "user_account with device-code token",
			authType:     config.UserAccount,
			accessToken:  issuerToken(t, oktaIssuer),
			refreshToken: "refresh",
			want:         deviceCodeTokenWarning,
		},
		{
			name:         "user_account with auth server token",
			authType:     config.UserAccount,
			accessToken:  issuerToken(t, authServerIssuer),
			refreshToken: "refresh",
		},
		{
			name:         "user_account with custom auth server",
			authType:     config.UserAccount,
			envIssuer:    "https://authorize-dev.mongodb.com",
			accessToken:  issuerToken(t, "https://authorize-dev.mongodb.com"),
			refreshToken: "refresh",
		},
		{
			name:         "user_account with production token against custom auth server",
			authType:     config.UserAccount,
			envIssuer:    "https://authorize-dev.mongodb.com",
			accessToken:  issuerToken(t, authServerIssuer),
			refreshToken: "refresh",
			want:         deviceCodeTokenWarning,
		},
		{
			name:         "gov user_account with gov token",
			authType:     config.UserAccount,
			service:      config.CloudGovService,
			accessToken:  issuerToken(t, "https://authorize.mongodbgov.com"),
			refreshToken: "refresh",
		},
		{
			name:         "gov user_account with device-code token",
			authType:     config.UserAccount,
			service:      config.CloudGovService,
			accessToken:  issuerToken(t, "https://auth.mongodbgov.com/oauth2/default"),
			refreshToken: "refresh",
			want:         deviceCodeTokenWarning,
		},
		{
			name:         "user_account with stored issuer and device-code token",
			authType:     config.UserAccount,
			storedIssuer: authServerIssuer,
			accessToken:  issuerToken(t, oktaIssuer),
			refreshToken: "refresh",
			want:         deviceCodeTokenWarning,
		},
		{
			name:         "legacy with device-code token",
			authType:     config.UserAccountLegacy,
			accessToken:  issuerToken(t, oktaIssuer),
			refreshToken: "refresh",
		},
		{
			name:         "legacy with auth server token",
			authType:     config.UserAccountLegacy,
			accessToken:  issuerToken(t, authServerIssuer),
			refreshToken: "refresh",
			want:         authCodeTokenWarning,
		},
		{
			name:         "legacy with custom auth server token",
			authType:     config.UserAccountLegacy,
			envIssuer:    "https://authorize-dev.mongodb.com",
			accessToken:  issuerToken(t, "https://authorize-dev.mongodb.com"),
			refreshToken: "refresh",
			want:         authCodeTokenWarning,
		},
		{
			name:         "legacy with stored issuer and auth server token",
			authType:     config.UserAccountLegacy,
			storedIssuer: "https://authorize-dev.mongodb.com",
			accessToken:  issuerToken(t, "https://authorize-dev.mongodb.com"),
			refreshToken: "refresh",
			want:         authCodeTokenWarning,
		},
		{
			name:         "gov legacy with gov auth server token",
			authType:     config.UserAccountLegacy,
			service:      config.CloudGovService,
			accessToken:  issuerToken(t, "https://authorize.mongodbgov.com"),
			refreshToken: "refresh",
			want:         authCodeTokenWarning,
		},
		{
			name:         "gov legacy with device-code token",
			authType:     config.UserAccountLegacy,
			service:      config.CloudGovService,
			accessToken:  issuerToken(t, "https://auth.mongodbgov.com/oauth2/default"),
			refreshToken: "refresh",
		},
		{
			name:         "legacy with non-JWT token",
			authType:     config.UserAccountLegacy,
			accessToken:  "not-a-jwt",
			refreshToken: "refresh",
		},
		{
			name:         "user_account with non-JWT token",
			authType:     config.UserAccount,
			accessToken:  "not-a-jwt",
			refreshToken: "refresh",
		},
		{
			name:         "user_account with token missing iss",
			authType:     config.UserAccount,
			accessToken:  issuerToken(t, ""),
			refreshToken: "refresh",
		},
		{
			name:        "access token only",
			accessToken: issuerToken(t, oktaIssuer),
		},
		{
			name:         "service account",
			authType:     config.ServiceAccount,
			accessToken:  issuerToken(t, oktaIssuer),
			refreshToken: "refresh",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Setenv("MONGODB_ATLAS_AUTH_SERVER_URL", tt.envIssuer)
			profile := config.NewProfile("test-profile", config.NewInMemoryStore())
			profile.SetAuthType(tt.authType)
			profile.SetService(tt.service)
			profile.SetAccessToken(tt.accessToken)
			profile.SetRefreshToken(tt.refreshToken)
			profile.SetAuthServerIssuer(tt.storedIssuer)

			buf := new(bytes.Buffer)
			log.SetWriter(buf)
			t.Cleanup(func() { log.SetWriter(os.Stderr) })

			WarnOnTokenAuthTypeMismatch(profile)

			assert.Equal(t, tt.want, buf.String())
		})
	}
}

func issuerToken(t *testing.T, issuer string) string {
	t.Helper()
	claims := map[string]any{"sub": "user@example.com"}
	if issuer != "" {
		claims["iss"] = issuer
	}
	payload, err := json.Marshal(claims)
	require.NoError(t, err)
	return base64.RawURLEncoding.EncodeToString([]byte(`{"alg":"none"}`)) + "." +
		base64.RawURLEncoding.EncodeToString(payload) + "."
}
