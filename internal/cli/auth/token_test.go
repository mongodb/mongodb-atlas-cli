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
	"testing"

	"github.com/mongodb/atlas-cli-core/config"
	"github.com/mongodb/atlas-cli-core/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
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
	assert.Equal(t, "test-access-token", buf.String())
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

func TestTokenBuilder_PreRunE_ProfileFromContext(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockStore := mocks.NewMockStore(ctrl)

	// Create a test profile
	testProfile := config.NewProfile("test-profile", mockStore)

	cmd := TokenBuilder()

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
