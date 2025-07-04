// Copyright 2023 MongoDB Inc
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

package pointintimerestore

import (
	"testing"

	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/test"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	atlasv2 "go.mongodb.org/atlas-sdk/v20250312005/admin"
	"go.uber.org/mock/gomock"
)

func TestEnableOpts_Watcher(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockStore := NewMockCompliancePolicyPointInTimeRestoresEnabler(ctrl)

	opts := &EnableOpts{
		store: mockStore,
	}

	expected := &atlasv2.DataProtectionSettings20231001{
		State: atlasv2.PtrString(active),
	}

	mockStore.
		EXPECT().
		DescribeCompliancePolicy(opts.ProjectID).
		Return(expected, nil).
		Times(1)

	_, res, err := opts.watcher()
	require.NoError(t, err)
	assert.True(t, res)
}

func TestEnableOpts_Run(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockStore := NewMockCompliancePolicyPointInTimeRestoresEnabler(ctrl)
	pointInTimeRestoreBefore := false
	pointInTimeRestoreAfter := true

	initial := &atlasv2.DataProtectionSettings20231001{
		PitEnabled: &pointInTimeRestoreBefore,
	}

	expected := &atlasv2.DataProtectionSettings20231001{
		State:      atlasv2.PtrString(active),
		PitEnabled: &pointInTimeRestoreAfter,
	}

	opts := &EnableOpts{
		store:             mockStore,
		policy:            initial,
		restoreWindowDays: 1,
	}

	mockStore.
		EXPECT().
		EnablePointInTimeRestore(opts.ProjectID, 1).
		Return(expected, nil).
		Times(1)

	assert.False(t, *opts.policy.PitEnabled)

	if err := opts.Run(); err != nil {
		t.Fatalf("Run() unexpected error: %v", err)
	}
	assert.True(t, *opts.policy.PitEnabled)
	test.VerifyOutputTemplate(t, enableTemplate, expected)
}

func TestEnableOpts_WatchRun(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockStore := NewMockCompliancePolicyPointInTimeRestoresEnabler(ctrl)

	opts := &EnableOpts{
		store: mockStore,
		WatchOpts: cli.WatchOpts{
			EnableWatch: true,
		},
		restoreWindowDays: 1,
	}

	expected := &atlasv2.DataProtectionSettings20231001{
		State: atlasv2.PtrString(active),
	}

	mockStore.
		EXPECT().
		EnablePointInTimeRestore(opts.ProjectID, 1).
		Return(expected, nil).
		Times(1)
	mockStore.
		EXPECT().
		DescribeCompliancePolicy(opts.ProjectID).
		Return(expected, nil).
		Times(1)

	require.NoError(t, opts.Run())
	test.VerifyOutputTemplate(t, enableWatchTemplate, expected)
}

func Test_validateRestoreWindowDays(t *testing.T) {
	tests := []struct {
		name              string
		restoreWindowDays int
		wantErr           require.ErrorAssertionFunc
	}{
		{
			"valid test",
			1,
			require.NoError,
		},
		{
			"invalid, negative number",
			-1,
			require.Error,
		},
		{
			"invalid, zero",
			0,
			require.Error,
		},
	}

	for _, tc := range tests {
		tt := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			opts := &EnableOpts{
				restoreWindowDays: tt.restoreWindowDays,
			}
			tt.wantErr(t, opts.validateRestoreWindowDays())
		})
	}
}
