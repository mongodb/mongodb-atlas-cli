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

package setup

import (
	"bytes"
	"context"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/mongodb/mongodb-atlas-cli/internal/config"
	"github.com/mongodb/mongodb-atlas-cli/internal/flag"
	"github.com/mongodb/mongodb-atlas-cli/internal/mocks"
	"github.com/mongodb/mongodb-atlas-cli/internal/test"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestBuilder(t *testing.T) {
	test.CmdValidator(
		t,
		Builder(),
		0,
		[]string{
			flag.Region,
			flag.ClusterName,
			flag.Provider,
			flag.AccessListIP,
			flag.Username,
			flag.Password,
			flag.SkipMongosh,
			flag.SkipSampleData,
		},
	)
}

func Test_setupOpts_Run(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockQuickstartFlow := mocks.NewMockFlow(ctrl)
	mockFlow := mocks.NewMockRefresher(ctrl)
	mockConfig := mocks.NewMockProfileReader(ctrl)

	ctx := context.TODO()
	buf := new(bytes.Buffer)

	opts := &Opts{
		quickstart:   mockQuickstartFlow,
		config:       mockConfig,
		skipLogin:    true,
		skipRegister: true,
	}
	opts.register.WithFlow(mockFlow)
	opts.OutWriter = buf

	mockConfig.
		EXPECT().
		OrgID().
		Return("1").
		Times(1)

	mockConfig.
		EXPECT().
		ProjectID().
		Return("1").
		Times(1)

	mockQuickstartFlow.
		EXPECT().
		Run().
		Return(nil).
		Times(1)

	mockQuickstartFlow.
		EXPECT().
		PreRun(ctx, buf).
		Return(nil).
		Times(1)

	require.NoError(t, opts.Run(ctx))
}

func Test_setupOpts_PreRunWithAPIKeys(t *testing.T) {
	t.Cleanup(test.CleanupConfig)
	ctrl := gomock.NewController(t)
	mockQuickstartFlow := mocks.NewMockFlow(ctrl)
	mockFlow := mocks.NewMockRefresher(ctrl)
	ctx := context.TODO()
	buf := new(bytes.Buffer)

	opts := &Opts{
		quickstart: mockQuickstartFlow,
	}
	opts.register.WithFlow(mockFlow)

	config.SetPublicAPIKey("publicKey")
	config.SetPrivateAPIKey("privateKey")
	opts.OutWriter = buf

	require.NoError(t, opts.PreRun(ctx))
	assert.Equal(t, `
You are already authenticated with an API key (Public key: publicKey).

Run "atlas auth setup --profile <profile_name>" to create a new Atlas account on a new Atlas CLI profile.
`, buf.String())
}

func Test_setupOpts_RunSkipRegister(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockQuickstartFlow := mocks.NewMockFlow(ctrl)
	mockFlow := mocks.NewMockRefresher(ctrl)
	ctx := context.TODO()
	buf := new(bytes.Buffer)

	opts := &Opts{
		quickstart: mockQuickstartFlow,
		skipLogin:  true,
	}
	opts.register.WithFlow(mockFlow)

	config.SetAccessToken("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiYWRtaW4iOnRydWV9.TJVA95OrM7E2cBab30RMHrHDcEfxjoYZgeFONFh7HgQ")

	opts.OutWriter = buf
	require.NoError(t, opts.PreRun(ctx))
	assert.True(t, opts.skipRegister)
}
