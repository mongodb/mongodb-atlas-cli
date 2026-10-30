// Copyright 2025 MongoDB Inc
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
	"context"
	"encoding/base64"
	"errors"
	"strconv"
	"testing"
	"time"

	"github.com/mongodb/atlas-cli-core/config"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	atlasauth "go.mongodb.org/atlas/auth"
	"go.uber.org/mock/gomock"
)

func Test_tokenOpts_Run_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockConfig := NewMockTokenConfig(ctrl)
	buf := new(bytes.Buffer)

	opts := &tokenOpts{
		config: mockConfig,
	}
	opts.OutWriter = buf
	opts.Output = "template" // Set output format to template

	mockConfig.EXPECT().AccessToken().Return("test-access-token").Times(1)

	err := opts.Run()
	require.NoError(t, err)
	assert.Equal(t, "test-access-token\n", buf.String())
}

func Test_tokenOpts_Run_NoToken(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockConfig := NewMockTokenConfig(ctrl)
	buf := new(bytes.Buffer)

	opts := &tokenOpts{
		config: mockConfig,
	}
	opts.OutWriter = buf

	mockConfig.EXPECT().AccessToken().Return("").Times(1)
	mockConfig.EXPECT().Name().Return("test-profile").Times(1)

	err := opts.Run()
	require.Error(t, err)
	assert.Contains(t, err.Error(), "no access token found for profile test-profile")
}

func TestTokenBuilder_PreRunE_DefaultConfig(t *testing.T) {
	cmd := TokenBuilder()

	// Test that PreRunE uses config.Default() when no profile in context
	err := cmd.PreRunE(cmd, []string{})

	// Should not error - just sets up the default config
	require.NoError(t, err)
}

func TestTokenBuilder_RefreshesExpiredTokenBeforePrinting(t *testing.T) {
	profile := setDefaultProfile(t, testAccessTokenWithExpiry(time.Now().Add(-time.Hour)), "old-refresh-token")
	ctrl := gomock.NewController(t)
	flow := mocks.NewMockRefresher(ctrl)
	flow.EXPECT().RefreshToken(gomock.Any(), "old-refresh-token").Return(&atlasauth.Token{
		AccessToken:  "new-access-token",
		RefreshToken: "new-refresh-token",
	}, nil, nil)

	refresher := &cli.RefresherOpts{}
	refresher.WithFlow(flow)
	cmd := newTokenBuilder(refresher)
	cmd.SetContext(config.WithProfile(t.Context(), profile))
	cmd.SilenceErrors = true
	cmd.SilenceUsage = true
	buf := new(bytes.Buffer)
	cmd.SetOut(buf)

	require.NoError(t, cmd.Execute())
	assert.Equal(t, "new-access-token\n", buf.String())
	assert.Equal(t, "new-access-token", profile.AccessToken())
	assert.Equal(t, "new-refresh-token", profile.RefreshToken())
}

func TestTokenBuilder_UsesContextProfileAndDoesNotRefreshValidToken(t *testing.T) {
	setDefaultProfile(t, testAccessTokenWithExpiry(time.Now().Add(-time.Hour)), "default-refresh-token")
	accessToken := testAccessTokenWithExpiry(time.Now().Add(time.Hour))
	profile := config.NewProfile("context-profile", config.NewInMemoryStore())
	profile.SetAccessToken(accessToken)
	ctrl := gomock.NewController(t)
	flow := mocks.NewMockRefresher(ctrl)

	refresher := &cli.RefresherOpts{}
	refresher.WithFlow(flow)
	cmd := newTokenBuilder(refresher)
	cmd.SetContext(config.WithProfile(t.Context(), profile))
	cmd.SilenceErrors = true
	cmd.SilenceUsage = true
	buf := new(bytes.Buffer)
	cmd.SetOut(buf)

	require.NoError(t, cmd.Execute())
	assert.Equal(t, accessToken+"\n", buf.String())
}

func TestTokenBuilder_ExpiredTokenWithoutRefreshTokenDoesNotPrint(t *testing.T) {
	profile := setDefaultProfile(t, testAccessTokenWithExpiry(time.Now().Add(-time.Hour)), "")
	ctrl := gomock.NewController(t)
	flow := mocks.NewMockRefresher(ctrl)

	refresher := &cli.RefresherOpts{}
	refresher.WithFlow(flow)
	cmd := newTokenBuilder(refresher)
	cmd.SetContext(config.WithProfile(t.Context(), profile))
	cmd.SilenceErrors = true
	cmd.SilenceUsage = true
	buf := new(bytes.Buffer)
	cmd.SetOut(buf)

	require.ErrorContains(t, cmd.Execute(), "expired and no refresh token is available")
	assert.Empty(t, buf.String())
}

func TestTokenBuilder_RefreshErrorDoesNotPrintExpiredToken(t *testing.T) {
	profile := setDefaultProfile(t, testAccessTokenWithExpiry(time.Now().Add(-time.Hour)), "refresh-token")
	ctrl := gomock.NewController(t)
	flow := mocks.NewMockRefresher(ctrl)
	refreshErr := errors.New("refresh failed")
	flow.EXPECT().RefreshToken(gomock.Any(), "refresh-token").Return(nil, nil, refreshErr)

	refresher := &cli.RefresherOpts{}
	refresher.WithFlow(flow)
	cmd := newTokenBuilder(refresher)
	cmd.SetContext(config.WithProfile(t.Context(), profile))
	cmd.SilenceErrors = true
	cmd.SilenceUsage = true
	buf := new(bytes.Buffer)
	cmd.SetOut(buf)

	require.ErrorIs(t, cmd.Execute(), refreshErr)
	assert.Empty(t, buf.String())
}

func setDefaultProfile(t *testing.T, accessToken, refreshToken string) *config.Profile {
	t.Helper()

	previousProfile := config.Default()
	profile := config.NewProfile("test-profile", config.NewInMemoryStore())
	profile.SetAccessToken(accessToken)
	profile.SetRefreshToken(refreshToken)
	config.SetDefaultProfile(profile)
	t.Cleanup(func() {
		config.SetDefaultProfile(previousProfile)
	})

	previousTokenRefreshed := cli.TokenRefreshed
	cli.TokenRefreshed = false
	t.Cleanup(func() {
		cli.TokenRefreshed = previousTokenRefreshed
	})

	return profile
}

func testAccessTokenWithExpiry(expiry time.Time) string {
	header := base64.RawURLEncoding.EncodeToString([]byte(`{"alg":"HS256","typ":"JWT"}`))
	claims := `{"exp":` + strconv.FormatInt(expiry.Unix(), 10) + `}`
	payload := base64.RawURLEncoding.EncodeToString([]byte(claims))
	signature := base64.RawURLEncoding.EncodeToString([]byte("signature"))
	return header + "." + payload + "." + signature
}

func TestTokenBuilder_UserAccountRefreshesThroughAuthServer(t *testing.T) {
	tests := []struct {
		name       string
		refreshErr error
	}{
		{name: "refreshed"},
		{name: "refresh fails", refreshErr: errors.New("refresh failed")},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			profile := setDefaultProfile(t, testAccessTokenWithExpiry(time.Now().Add(-time.Hour)), "refresh-token")
			profile.SetAuthType(config.UserAccount)
			ctrl := gomock.NewController(t)
			// The device flow must never receive Authorization Code tokens.
			flow := mocks.NewMockRefresher(ctrl)

			calls := 0
			refresher := &cli.RefresherOpts{}
			refresher.WithFlow(flow)
			refresher.WithUserAccountRefresher(func(_ context.Context, p *config.Profile) error {
				calls++
				require.Same(t, profile, p)
				if tt.refreshErr != nil {
					return tt.refreshErr
				}
				p.SetAccessToken("new-access-token")
				return nil
			})
			cmd := newTokenBuilder(refresher)
			cmd.SetContext(config.WithProfile(t.Context(), profile))
			cmd.SilenceErrors = true
			cmd.SilenceUsage = true
			buf := new(bytes.Buffer)
			cmd.SetOut(buf)

			err := cmd.Execute()
			require.Equal(t, 1, calls)
			if tt.refreshErr != nil {
				require.ErrorIs(t, err, tt.refreshErr)
				assert.Empty(t, buf.String())
				return
			}
			require.NoError(t, err)
			assert.Equal(t, "new-access-token\n", buf.String())
		})
	}
}

func TestTokenBuilder_UserAccountLegacyRefreshesThroughDeviceFlow(t *testing.T) {
	profile := setDefaultProfile(t, testAccessTokenWithExpiry(time.Now().Add(-time.Hour)), "old-refresh-token")
	profile.SetAuthType(config.UserAccountLegacy)
	ctrl := gomock.NewController(t)
	flow := mocks.NewMockRefresher(ctrl)
	flow.EXPECT().RefreshToken(gomock.Any(), "old-refresh-token").Return(&atlasauth.Token{
		AccessToken:  "new-access-token",
		RefreshToken: "new-refresh-token",
	}, nil, nil)

	refresher := &cli.RefresherOpts{}
	refresher.WithFlow(flow)
	refresher.WithUserAccountRefresher(func(context.Context, *config.Profile) error {
		t.Fatal("legacy profiles must not use the Authorization Server refresher")
		return nil
	})
	cmd := newTokenBuilder(refresher)
	cmd.SetContext(config.WithProfile(t.Context(), profile))
	cmd.SilenceErrors = true
	cmd.SilenceUsage = true
	buf := new(bytes.Buffer)
	cmd.SetOut(buf)

	require.NoError(t, cmd.Execute())
	assert.Equal(t, "new-access-token\n", buf.String())
}
