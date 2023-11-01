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
	"context"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/mongodb/mongodb-atlas-cli/internal/cli"
	"github.com/mongodb/mongodb-atlas-cli/internal/flag"
	mocks "github.com/mongodb/mongodb-atlas-cli/internal/mocks/atlas"
	"github.com/mongodb/mongodb-atlas-cli/internal/test"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	atlasv2 "go.mongodb.org/atlas-sdk/v20231001002/admin"
)

func TestEnableBuilder(t *testing.T) {
	test.CmdValidator(
		t,
		EnableBuilder(),
		0,
		[]string{
			flag.ProjectID,
			flag.Output,
			flag.EnableWatch,
			flag.RestoreWindowDays,
		},
	)
}

func TestEnableOpts_InitStore(t *testing.T) {
	opts := &EnableOpts{}

	require.NoError(t, opts.initStore(context.TODO())())
	assert.NotNil(t, opts.store)
}

func TestEnableOpts_Watcher(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockStore := mocks.NewMockCompliancePolicyPointInTimeRestoresEnabler(ctrl)

	opts := &EnableOpts{
		store: mockStore,
	}

	expected := &atlasv2.DataProtectionSettings{
		State: atlasv2.PtrString(active),
	}

	mockStore.
		EXPECT().
		DescribeCompliancePolicy(opts.ProjectID).
		Return(expected, nil).
		Times(1)

	res, err := opts.watcher()
	require.NoError(t, err)
	assert.True(t, res)
}

func TestEnableOpts_Run(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockStore := mocks.NewMockCompliancePolicyPointInTimeRestoresEnabler(ctrl)
	pointInTimeRestoreBefore := false
	pointInTimeRestoreAfter := true

	initial := &atlasv2.DataProtectionSettings{
		PitEnabled: &pointInTimeRestoreBefore,
	}

	expected := &atlasv2.DataProtectionSettings{
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
	mockStore := mocks.NewMockCompliancePolicyPointInTimeRestoresEnabler(ctrl)

	opts := &EnableOpts{
		store: mockStore,
		WatchOpts: cli.WatchOpts{
			EnableWatch: true,
		},
		restoreWindowDays: 1,
	}

	expected := &atlasv2.DataProtectionSettings{
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

	if err := opts.Run(); err != nil {
		t.Fatalf("Run() unexpected error: %v", err)
	}

	test.VerifyOutputTemplate(t, enableWatchTemplate, expected)
}

func TestRestoreWindowDaysValidator(t *testing.T) {
	tests := []struct {
		name              string
		restoreWindowDays int
		wantErr           bool
	}{
		{
			"valid test",
			1,
			false,
		},
		{
			"invalid, negative number",
			-1,
			true,
		},
		{
			"invalid, zero",
			0,
			true,
		},
	}

	for _, testOptions := range tests {
		opts := &EnableOpts{
			restoreWindowDays: testOptions.restoreWindowDays,
		}

		if testOptions.wantErr {
			require.Error(t, opts.validateRestoreWindowDays())
		} else {
			require.NoError(t, opts.validateRestoreWindowDays())
		}
	}
}
