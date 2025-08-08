// Copyright 2022 MongoDB Inc
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

//go:build unit

package auth

import (
	"bytes"
	"testing"

	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/config"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_whoOpts_Run(t *testing.T) {
	buf := new(bytes.Buffer)
	opts := &whoOpts{
		OutWriter:   buf,
		authSubject: "test@test.com",
		authType:    "account",
	}
	require.NoError(t, opts.Run())
	assert.Equal(t, "Logged in as test@test.com account\n", buf.String())
}

func Test_authTypeAndSubject(t *testing.T) {
	tests := []struct {
		name            string
		setupConfig     func()
		expectedType    string
		expectedSubject string
		expectedError   error
	}{
		{
			name: "API Keys authentication",
			setupConfig: func() {
				config.SetAuthType(config.APIKeys)
				config.SetPublicAPIKey("test-public-key")
			},
			expectedType:    "key",
			expectedSubject: "test-public-key",
			expectedError:   nil,
		},
		{
			name: "Service Account authentication",
			setupConfig: func() {
				config.SetAuthType(config.ServiceAccount)
				config.SetClientID("test-client-id")
			},
			expectedType:    "service account",
			expectedSubject: "test-client-id",
			expectedError:   nil,
		},
		{
			name: "User Account authentication",
			setupConfig: func() {
				config.SetAuthType(config.UserAccount)
				config.SetAccessToken("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiJ0ZXN0QGV4YW1wbGUuY29tIiwibmFtZSI6IlRlc3QgVXNlciIsImlhdCI6MTUxNjIzOTAyMn0.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c")
			},
			expectedType:    "account",
			expectedSubject: "test@example.com",
			expectedError:   nil,
		},
		{
			name: "NoAuth authentication",
			setupConfig: func() {
				config.SetAuthType(config.NoAuth)
			},
			expectedType:    "",
			expectedSubject: "",
			expectedError:   ErrUnauthenticated,
		},
		{
			name: "Empty authentication type",
			setupConfig: func() {
				config.SetAuthType(config.AuthMechanism(""))
			},
			expectedType:    "",
			expectedSubject: "",
			expectedError:   ErrUnauthenticated,
		},
		{
			name: "Unknown authentication type",
			setupConfig: func() {
				config.SetAuthType(config.AuthMechanism("unknown"))
			},
			expectedType:    "",
			expectedSubject: "",
			expectedError:   ErrUnauthenticated,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Save original config state
			originalAuthType := config.AuthType()
			originalPublicKey := config.PublicAPIKey()
			originalClientID := config.ClientID()
			originalAccessToken := config.AccessToken()

			defer func() {
				config.SetAuthType(originalAuthType)
				config.SetPublicAPIKey(originalPublicKey)
				config.SetClientID(originalClientID)
				config.SetAccessToken(originalAccessToken)
			}()

			// Clear config state
			config.SetAuthType(config.AuthMechanism(""))
			config.SetPublicAPIKey("")
			config.SetClientID("")
			config.SetAccessToken("")

			// Setup test configuration
			tt.setupConfig()

			// Execute the function
			authType, authSubject, err := authTypeAndSubject()

			// Verify results
			assert.Equal(t, tt.expectedType, authType)
			assert.Equal(t, tt.expectedSubject, authSubject)

			if tt.expectedError != nil {
				assert.ErrorIs(t, err, tt.expectedError)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func Test_authTypeAndSubject_EdgeCases(t *testing.T) {
	t.Run("API Keys with empty public key", func(t *testing.T) {
		// Save original config state
		originalAuthType := config.AuthType()
		originalPublicKey := config.PublicAPIKey()

		defer func() {
			config.SetAuthType(originalAuthType)
			config.SetPublicAPIKey(originalPublicKey)
		}()

		config.SetAuthType(config.APIKeys)
		config.SetPublicAPIKey("")

		authType, authSubject, err := authTypeAndSubject()

		assert.Equal(t, "key", authType)
		assert.Equal(t, "", authSubject)
		assert.NoError(t, err)
	})

	t.Run("Service Account with empty client ID", func(t *testing.T) {
		// Save original config state
		originalAuthType := config.AuthType()
		originalClientID := config.ClientID()

		defer func() {
			config.SetAuthType(originalAuthType)
			config.SetClientID(originalClientID)
		}()

		config.SetAuthType(config.ServiceAccount)
		config.SetClientID("")

		authType, authSubject, err := authTypeAndSubject()

		assert.Equal(t, "service account", authType)
		assert.Equal(t, "", authSubject)
		assert.NoError(t, err)
	})

	t.Run("User Account with malformed access token", func(t *testing.T) {
		// Save original config state
		originalAuthType := config.AuthType()
		originalAccessToken := config.AccessToken()

		defer func() {
			config.SetAuthType(originalAuthType)
			config.SetAccessToken(originalAccessToken)
		}()

		config.SetAuthType(config.UserAccount)
		config.SetAccessToken("invalid-token")

		authType, authSubject, err := authTypeAndSubject()

		assert.Equal(t, "account", authType)
		assert.Equal(t, "", authSubject) // AccessTokenSubject will return empty for invalid token
		assert.NoError(t, err)
	})
}
