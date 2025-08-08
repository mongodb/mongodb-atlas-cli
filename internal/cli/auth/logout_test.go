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

	mockApiKeysCleanUp(mockConfig)
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
	mockFlow.
		EXPECT().
		RevokeToken(ctx, gomock.Any(), gomock.Any()).
		Return(nil, nil).
		Times(1)
	mockTokenCleanUp(mockConfig)
	mockProjectAndOrgCleanUp(mockConfig)
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

	mockApiKeysCleanUp(mockConfig)
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
	}
	ctx := t.Context()
	mockConfig.
		EXPECT().
		AuthType().
		Return(config.ServiceAccount).
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

	mockTokenCleanUp(mockConfig)
	mockProjectAndOrgCleanUp(mockConfig)
	mockApiKeysCleanUp(mockConfig)

	mockConfig.
		EXPECT().
		Delete().
		Return(nil).
		Times(1)

	require.NoError(t, opts.Run(ctx))
}

func mockApiKeysCleanUp(mockConfig *MockConfigDeleter) {
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
