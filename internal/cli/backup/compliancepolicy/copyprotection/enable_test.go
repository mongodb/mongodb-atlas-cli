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

package copyprotection

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
	mockStore := NewMockCompliancePolicyCopyProtectionEnabler(ctrl)

	opts := &EnableOpts{
		store: mockStore,
	}
	expected := &atlasv2.DataProtectionSettings20231001{}
	expected.SetState(active)

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
	mockStore := NewMockCompliancePolicyCopyProtectionEnabler(ctrl)

	expected := &atlasv2.DataProtectionSettings20231001{}
	expected.SetState(active)
	expected.SetCopyProtectionEnabled(true)

	opts := &EnableOpts{
		store: mockStore,
	}

	mockStore.
		EXPECT().
		EnableCopyProtection(opts.ProjectID).
		Return(expected, nil).
		Times(1)

	if err := opts.Run(); err != nil {
		t.Fatalf("Run() unexpected error: %v", err)
	}
	assert.True(t, opts.policy.GetCopyProtectionEnabled())
	test.VerifyOutputTemplate(t, enableTemplate, expected)
}

func TestEnableOpts_WatchRun(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockStore := NewMockCompliancePolicyCopyProtectionEnabler(ctrl)

	opts := &EnableOpts{
		store: mockStore,
		WatchOpts: cli.WatchOpts{
			EnableWatch: true,
		},
	}

	expected := &atlasv2.DataProtectionSettings20231001{}
	expected.SetState(active)

	mockStore.
		EXPECT().
		EnableCopyProtection(opts.ProjectID).
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
