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
	"testing"

	"github.com/mongodb/atlas-cli-core/config"
	"github.com/mongodb/atlas-cli-core/mocks"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
)

func Test_logoutOpts_Run_UserAccount(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockFlow := NewMockRevoker(ctrl)
	mockConfig := NewMockConfigDeleter(ctrl)

	buf := new(bytes.Buffer)

	opts := logoutOpts{
		OutWriter: buf,
		config:    mockConfig,
		flow:      mockFlow,
		DeleteOpts: &cli.DeleteOpts{
			Confirm: true,
		},
	}
	ctx := t.Context()

	mockConfig.
		EXPECT().
		AuthType().
		Return(config.UserAccount).
		Times(1)

	mockFlow.
		EXPECT().
		RevokeToken(ctx, gomock.Any(), gomock.Any()).
		Return(nil, nil).
		Times(1)

	mockTokenCleanUp(mockConfig)
	mockProjectAndOrgCleanUp(mockConfig)

	mockConfig.
		EXPECT().
		Delete().
		Return(nil).
		Times(1)
	require.NoError(t, opts.Run(ctx))
}

func Test_logoutOpts_Run_APIKeys(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockFlow := NewMockRevoker(ctrl)
	mockConfig := NewMockConfigDeleter(ctrl)

	buf := new(bytes.Buffer)

	opts := logoutOpts{
		OutWriter: buf,
		config:    mockConfig,
		flow:      mockFlow,
		DeleteOpts: &cli.DeleteOpts{
			Confirm: true,
		},
	}
	ctx := t.Context()
	mockConfig.
		EXPECT().
		AuthType().
		Return(config.APIKeys).
		Times(1)

	mockAPIKeysCleanUp(mockConfig)
	mockProjectAndOrgCleanUp(mockConfig)

	mockConfig.
		EXPECT().
		Delete().
		Return(nil).
		Times(1)
	require.NoError(t, opts.Run(ctx))
}

func Test_logoutOpts_Run_ServiceAccount(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockFlow := NewMockRevoker(ctrl)
	mockConfig := NewMockConfigDeleter(ctrl)

	buf := new(bytes.Buffer)

	opts := logoutOpts{
		OutWriter: buf,
		config:    mockConfig,
		flow:      mockFlow,
		DeleteOpts: &cli.DeleteOpts{
			Confirm: true,
		},
		revokeServiceAccountToken: func() error {
			return nil
		},
	}
	ctx := t.Context()
	mockConfig.
		EXPECT().
		AuthType().
		Return(config.ServiceAccount).
		Times(1)

	mockServiceAccountCleanUp(mockConfig)
	mockProjectAndOrgCleanUp(mockConfig)
	mockConfig.
		EXPECT().
		Delete().
		Return(nil).
		Times(1)

	require.NoError(t, opts.Run(ctx))
}

func Test_logoutOpts_Run_Keep_UserAccount(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockFlow := NewMockRevoker(ctrl)
	mockConfig := NewMockConfigDeleter(ctrl)

	buf := new(bytes.Buffer)

	opts := logoutOpts{
		OutWriter: buf,
		config:    mockConfig,
		flow:      mockFlow,
		DeleteOpts: &cli.DeleteOpts{
			Confirm: true,
		},
		keepConfig: true,
	}
	ctx := t.Context()
	mockConfig.
		EXPECT().
		AuthType().
		Return(config.UserAccount).
		Times(1)

	mockFlow.
		EXPECT().
		RevokeToken(ctx, gomock.Any(), gomock.Any()).
		Return(nil, nil).
		Times(1)

	mockTokenCleanUp(mockConfig)
	mockProjectAndOrgCleanUp(mockConfig)
	mockConfig.
		EXPECT().
		Save().
		Return(nil).
		Times(1)

	require.NoError(t, opts.Run(ctx))
}

func Test_logoutOpts_Run_Keep_APIKeys(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockFlow := NewMockRevoker(ctrl)
	mockConfig := NewMockConfigDeleter(ctrl)

	buf := new(bytes.Buffer)

	opts := logoutOpts{
		OutWriter: buf,
		config:    mockConfig,
		flow:      mockFlow,
		DeleteOpts: &cli.DeleteOpts{
			Confirm: true,
		},
		keepConfig: true,
	}
	ctx := t.Context()
	mockConfig.
		EXPECT().
		AuthType().
		Return(config.APIKeys).
		Times(1)

	mockAPIKeysCleanUp(mockConfig)
	mockProjectAndOrgCleanUp(mockConfig)
	mockConfig.
		EXPECT().
		Save().
		Return(nil).
		Times(1)

	require.NoError(t, opts.Run(ctx))
}

func Test_logoutOpts_Run_Keep_ServiceAccount(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockFlow := NewMockRevoker(ctrl)
	mockConfig := NewMockConfigDeleter(ctrl)

	buf := new(bytes.Buffer)

	opts := logoutOpts{
		OutWriter: buf,
		config:    mockConfig,
		flow:      mockFlow,
		DeleteOpts: &cli.DeleteOpts{
			Confirm: true,
		},
		keepConfig: true,
		revokeServiceAccountToken: func() error {
			return nil
		},
	}
	ctx := t.Context()
	mockConfig.
		EXPECT().
		AuthType().
		Return(config.ServiceAccount).
		Times(1)

	mockServiceAccountCleanUp(mockConfig)
	mockProjectAndOrgCleanUp(mockConfig)
	mockConfig.
		EXPECT().
		Save().
		Return(nil).
		Times(1)

	require.NoError(t, opts.Run(ctx))
}

func Test_logoutOpts_Run_NoAuth(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockFlow := NewMockRevoker(ctrl)
	mockConfig := NewMockConfigDeleter(ctrl)

	buf := new(bytes.Buffer)

	opts := logoutOpts{
		OutWriter: buf,
		config:    mockConfig,
		flow:      mockFlow,
		DeleteOpts: &cli.DeleteOpts{
			Confirm: true,
		},
		keepConfig: false,
	}
	ctx := t.Context()
	mockConfig.
		EXPECT().
		AuthType().
		Return(config.NoAuth).
		Times(1)

	mockAPIKeysCleanUp(mockConfig)
	mockTokenCleanUp(mockConfig)
	mockServiceAccountCleanUp(mockConfig)
	mockProjectAndOrgCleanUp(mockConfig)

	mockConfig.
		EXPECT().
		Delete().
		Return(nil).
		Times(1)

	require.NoError(t, opts.Run(ctx))
}

func mockAPIKeysCleanUp(mockConfig *MockConfigDeleter) {
	mockConfig.
		EXPECT().
		SetPublicAPIKey("").
		Times(1)
	mockConfig.
		EXPECT().
		SetPrivateAPIKey("").
		Times(1)
}

func mockTokenCleanUp(mockConfig *MockConfigDeleter) {
	mockConfig.
		EXPECT().
		SetRefreshToken("").
		Times(1)
	mockConfig.
		EXPECT().
		SetAccessToken("").
		Times(1)
}

func mockServiceAccountCleanUp(mockConfig *MockConfigDeleter) {
	mockConfig.
		EXPECT().
		SetClientID("").
		Times(1)
	mockConfig.
		EXPECT().
		SetClientSecret("").
		Times(1)
}

func mockProjectAndOrgCleanUp(mockConfig *MockConfigDeleter) {
	mockConfig.
		EXPECT().
		SetProjectID("").
		Times(1)
	mockConfig.
		EXPECT().
		SetOrgID("").
		Times(1)
}

func TestLogoutBuilder_PreRunE_DefaultConfig(t *testing.T) {
	cmd := LogoutBuilder()

	// Test that PreRunE uses config.Default() when no profile in context
	err := cmd.PreRunE(cmd, []string{})

	// Should not error - just sets up the default config
	require.NoError(t, err)
}

func TestLogoutBuilder_PreRunE_ProfileFromContext(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockStore := mocks.NewMockStore(ctrl)

	// Create a test profile
	testProfile := config.NewProfile("test-profile", mockStore)

	cmd := LogoutBuilder()

	// Add profile to context and execute the command with that context
	ctx := config.WithProfile(t.Context(), testProfile)
	cmd.SetContext(ctx)

	mockStore.EXPECT().
		GetHierarchicalValue("test-profile", gomock.Any()).
		Return("").
		AnyTimes()

	err := cmd.PreRunE(cmd, []string{})
	require.NoError(t, err)
}
