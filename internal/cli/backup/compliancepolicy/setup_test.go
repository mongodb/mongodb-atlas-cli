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

package compliancepolicy

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/flag"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/mocks"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/test"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	atlasv2 "go.mongodb.org/atlas-sdk/v20240530002/admin"
)

func TestSetupBuilder(t *testing.T) {
	test.CmdValidator(
		t,
		SetupBuilder(),
		0,
		[]string{
			flag.ProjectID,
			flag.Output,
			flag.File,
			flag.Force,
			flag.EnableWatch,
		},
	)
}

// Tests that setupWatcher() returns true when status == "ACTIVE".
func TestSetupOpts_Watcher(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockStore := mocks.NewMockCompliancePolicyUpdater(ctrl)
	state := active

	opts := &SetupOpts{
		store: mockStore,
	}

	expected := &atlasv2.DataProtectionSettings20231001{
		State: &state,
	}

	mockStore.
		EXPECT().
		DescribeCompliancePolicy(opts.ProjectID).
		Return(expected, nil).
		Times(1)

	_, res, err := opts.setupWatcher()
	require.NoError(t, err)
	assert.True(t, res)
}

// Verifies the output template.
func TestSetupOpts_Run(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockStore := mocks.NewMockCompliancePolicyUpdater(ctrl)
	state := active

	opts := &SetupOpts{
		store:   mockStore,
		confirm: true,
		policy:  new(atlasv2.DataProtectionSettings20231001),
	}

	tests := map[string]*atlasv2.DataProtectionSettings20231001{
		"no policy": {
			State: &state,
		},
		"with scheduled policy": {
			State: &state,
			ScheduledPolicyItems: &[]atlasv2.BackupComplianceScheduledPolicyItem{
				*atlasv2.NewBackupComplianceScheduledPolicyItem(1, "daily", "weeks", 1),
			},
		},
		"with ondemand policy": {
			State:              &state,
			OnDemandPolicyItem: atlasv2.NewBackupComplianceOnDemandPolicyItem(1, "ondemand", "weeks", 1),
		},
	}

	for name, expected := range tests {
		t.Run(name, func(t *testing.T) {
			mockStore.
				EXPECT().
				UpdateCompliancePolicy(opts.ProjectID, opts.policy).
				Return(expected, nil).
				Times(1)

			require.NoError(t, opts.Run())
			test.VerifyOutputTemplate(t, setupTemplate, expected)
		})
	}
}

// Verifies the output template when using --watch.
func TestSetupOpts_WatchRun(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockStore := mocks.NewMockCompliancePolicyUpdater(ctrl)
	state := active

	opts := &SetupOpts{
		store:   mockStore,
		confirm: true,
		policy:  new(atlasv2.DataProtectionSettings20231001),
		WatchOpts: cli.WatchOpts{
			EnableWatch: true,
		},
	}

	expected := &atlasv2.DataProtectionSettings20231001{
		State: &state,
	}

	mockStore.
		EXPECT().
		UpdateCompliancePolicy(opts.ProjectID, opts.policy).
		Return(expected, nil).
		Times(1)
	mockStore.
		EXPECT().
		DescribeCompliancePolicy(opts.ProjectID).
		Return(expected, nil).
		Times(1)

	require.NoError(t, opts.Run())
	test.VerifyOutputTemplate(t, setupWatchTemplate, expected)
}
