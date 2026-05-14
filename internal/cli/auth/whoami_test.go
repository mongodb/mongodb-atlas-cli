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
	"strings"
	"testing"
	"time"

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

func Test_whoOpts_RunUserDelegation(t *testing.T) {
	futureExpiry := time.Now().Add(time.Hour).Format(time.RFC3339)
	pastExpiry := time.Now().Add(-time.Hour).Format(time.RFC3339)

	tests := []struct {
		name         string
		tokenExpiry  string
		refreshToken string
		wantContains []string
	}{
		{
			name:         "valid expiry with refresh token",
			tokenExpiry:  futureExpiry,
			refreshToken: "rt",
			wantContains: []string{"expires in", "auto-refresh enabled"},
		},
		{
			name:         "expired token with refresh token",
			tokenExpiry:  pastExpiry,
			refreshToken: "rt",
			wantContains: []string{"token expired", "auto-refresh enabled"},
		},
		{
			name:         "no expiry with refresh token",
			tokenExpiry:  "",
			refreshToken: "rt",
			wantContains: []string{"auto-refresh enabled"},
		},
		{
			name:         "no expiry, no refresh token",
			tokenExpiry:  "",
			refreshToken: "",
			wantContains: []string{"Connected to MongoDB Atlas."},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			buf := new(bytes.Buffer)
			opts := &whoOpts{
				OutWriter:    buf,
				tokenExpiry:  tt.tokenExpiry,
				refreshToken: tt.refreshToken,
			}
			require.NoError(t, opts.RunUserDelegation())
			for _, want := range tt.wantContains {
				assert.True(t, strings.Contains(buf.String(), want),
					"output %q does not contain %q", buf.String(), want)
			}
		})
	}
}
