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

	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/config"
	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
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

	mockConfig.
		EXPECT().
		SetAccessToken("").
		Times(1)
	mockConfig.
		EXPECT().
		SetRefreshToken("").
		Times(1)

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

	mockConfig.
		EXPECT().
		SetPublicAPIKey("").
		Times(1)
	mockConfig.
		EXPECT().
		SetPrivateAPIKey("").
		Times(1)
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
	}
	ctx := t.Context()
	mockConfig.
		EXPECT().
		AuthType().
		Return(config.ServiceAccount).
		Times(1)

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

	mockConfig.
		EXPECT().
		SetAccessToken("").
		Times(1)
	mockConfig.
		EXPECT().
		SetRefreshToken("").
		Times(1)
	mockConfig.
		EXPECT().
		SetProjectID("").
		Times(1)
	mockConfig.
		EXPECT().
		SetOrgID("").
		Times(1)
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

	mockConfig.
		EXPECT().
		SetPublicAPIKey("").
		Times(1)
	mockConfig.
		EXPECT().
		SetPrivateAPIKey("").
		Times(1)
	mockConfig.
		EXPECT().
		SetProjectID("").
		Times(1)
	mockConfig.
		EXPECT().
		SetOrgID("").
		Times(1)
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
	}
	ctx := t.Context()
	mockConfig.
		EXPECT().
		AuthType().
		Return(config.ServiceAccount).
		Times(1)

	mockConfig.
		EXPECT().
		SetProjectID("").
		Times(1)
	mockConfig.
		EXPECT().
		SetOrgID("").
		Times(1)
	mockConfig.
		EXPECT().
		Save().
		Return(nil).
		Times(1)

	require.NoError(t, opts.Run(ctx))
}

func Test_LogoutBuilder_PreRunE(t *testing.T) {
	t.Run("successful prerun", func(t *testing.T) {
		cmd := LogoutBuilder()

		// Create a test command context
		testCmd := &cobra.Command{}
		buf := new(bytes.Buffer)
		testCmd.SetOut(buf)

		// Execute PreRunE
		err := cmd.PreRunE(testCmd, []string{})

		// Should not return an error
		assert.NoError(t, err)
	})
}

func Test_logoutOpts_initFlow(t *testing.T) {
	t.Run("successful flow initialization", func(t *testing.T) {
		opts := &logoutOpts{}

		err := opts.initFlow()

		// Should not return an error under normal conditions
		assert.NoError(t, err)
		assert.NotNil(t, opts.flow)
	})
}

func Test_LogoutBuilder_RunE_ErrorHandling(t *testing.T) {
	t.Run("no refresh token error for user account", func(t *testing.T) {
		// Save original config state
		originalRefreshToken := config.RefreshToken()
		originalAuthType := config.AuthType()
		defer func() {
			config.SetRefreshToken(originalRefreshToken)
			config.SetAuthType(originalAuthType)
		}()

		// Set up UserAccount auth type but clear refresh token to trigger error
		config.SetAuthType(config.UserAccount)
		config.SetRefreshToken("")

		cmd := LogoutBuilder()

		// Create a test command context
		testCmd := &cobra.Command{}
		buf := new(bytes.Buffer)
		testCmd.SetOut(buf)

		// Execute PreRunE first
		err := cmd.PreRunE(testCmd, []string{})
		assert.NoError(t, err)

		// Execute RunE - should return ErrUnauthenticated
		err = cmd.RunE(testCmd, []string{})
		assert.ErrorIs(t, err, ErrUnauthenticated)
	})

	t.Run("api keys flow validates properly", func(t *testing.T) {
		// Save original config state
		originalAuthType := config.AuthType()
		originalPublicKey := config.PublicAPIKey()
		defer func() {
			config.SetAuthType(originalAuthType)
			config.SetPublicAPIKey(originalPublicKey)
		}()

		// Set up API key configuration
		config.SetAuthType(config.APIKeys)
		config.SetPublicAPIKey("test-public-key")

		cmd := LogoutBuilder()

		// Create a test command context
		testCmd := &cobra.Command{}
		buf := new(bytes.Buffer)
		testCmd.SetOut(buf)

		// Execute PreRunE first
		err := cmd.PreRunE(testCmd, []string{})
		assert.NoError(t, err)

		// For API keys, the RunE should work without refresh token
		// Note: This would normally prompt for confirmation, but we're just testing structure
		assert.NotNil(t, cmd.RunE)
	})
}

func Test_LogoutBuilder_Integration(t *testing.T) {
	t.Run("command structure validation", func(t *testing.T) {
		cmd := LogoutBuilder()

		// Verify command metadata
		assert.Equal(t, "logout", cmd.Use)
		assert.Contains(t, cmd.Short, "Log out")
		assert.NotEmpty(t, cmd.Example)

		// Verify command functions are set
		assert.NotNil(t, cmd.PreRunE)
		assert.NotNil(t, cmd.RunE)

		// Verify flags are configured
		assert.True(t, cmd.Flags().Lookup("force") != nil)
		assert.True(t, cmd.Flags().Lookup("keep") != nil)

		// Verify the keep flag is hidden
		keepFlag := cmd.Flags().Lookup("keep")
		assert.NotNil(t, keepFlag)
		assert.True(t, keepFlag.Hidden)
	})
}
