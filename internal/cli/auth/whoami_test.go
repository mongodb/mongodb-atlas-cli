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

package auth

import (
	"bytes"
	"encoding/base64"
	"strconv"
	"testing"
	"time"

	"github.com/mongodb/atlas-cli-core/config"
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

func Test_authTypeAndSubject_tokenExpiry(t *testing.T) {
	const expiry = "2026-10-28T12:00:00Z"
	tests := []struct {
		authType   config.AuthMechanism
		wantExpiry string
	}{
		{authType: config.UserAccount, wantExpiry: expiry},
		{authType: config.UserAccountLegacy, wantExpiry: ""},
	}
	for _, tt := range tests {
		t.Run(string(tt.authType), func(t *testing.T) {
			previous := config.Default()
			profile := config.NewProfile("test-profile", config.NewInMemoryStore())
			profile.SetAuthType(tt.authType)
			profile.SetAccessToken(testAccessTokenWithSubject(t, "user@example.com", time.Now().Add(time.Hour)))
			profile.SetRefreshToken("refresh-token")
			profile.SetTokenExpiry(expiry)
			config.SetDefaultProfile(profile)
			t.Cleanup(func() { config.SetDefaultProfile(previous) })

			authType, subject, tokenExpiry, err := authTypeAndSubject()
			require.NoError(t, err)
			assert.Equal(t, "account", authType)
			assert.Equal(t, "user@example.com", subject)
			assert.Equal(t, tt.wantExpiry, tokenExpiry)
		})
	}
}

func Test_whoOpts_Run_tokenExpiry(t *testing.T) {
	buf := new(bytes.Buffer)
	const expiry = "2026-10-28T12:00:00Z"
	opts := &whoOpts{
		OutWriter:   buf,
		authSubject: "test@test.com",
		authType:    "account",
	}
	opts.tokenExpiry = expiry
	require.NoError(t, opts.Run())
	assert.Equal(t, "Logged in as test@test.com account (token expires "+expiry+")\n", buf.String())
}

func testAccessTokenWithSubject(t *testing.T, subject string, expiry time.Time) string {
	t.Helper()
	header := base64.RawURLEncoding.EncodeToString([]byte(`{"alg":"HS256","typ":"JWT"}`))
	claims := `{"sub":"` + subject + `","exp":` + strconv.FormatInt(expiry.Unix(), 10) + `}`
	payload := base64.RawURLEncoding.EncodeToString([]byte(claims))
	signature := base64.RawURLEncoding.EncodeToString([]byte("signature"))
	return header + "." + payload + "." + signature
}
