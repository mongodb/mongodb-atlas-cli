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

package compliancepolicy

import (
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	mocks "github.com/mongodb/mongodb-atlas-cli/internal/mocks/atlas"
	"github.com/stretchr/testify/assert"
	atlasv2 "go.mongodb.org/atlas-sdk/v20230201004/admin"
)

func TestCopyProtectionOpts_PreRun(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockStore := mocks.NewMockCompliancePolicy(ctrl)

	opts := &CopyProtectionOpts{
		store: mockStore,
	}
	expected := &atlasv2.DataProtectionSettings{}

	mockStore.
		EXPECT().
		DescribeCompliancePolicy(opts.ProjectID).
		Return(expected, nil).
		Times(1)

	if err := opts.PreRun(); err != nil {
		t.Fatalf("PreRun() unexpected error: %v", err)
	}

	assert.Equal(t, opts.policy, expected)
}

func TestCopyProtectionOpts_PreRun_fail(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockStore := mocks.NewMockCompliancePolicy(ctrl)

	opts := &CopyProtectionOpts{
		store: mockStore,
	}

	mockStore.
		EXPECT().
		DescribeCompliancePolicy(opts.ProjectID).
		Return(nil, errors.New("network error")).
		Times(1)

	err := opts.PreRun()
	assert.Error(t, err)
}

func TestCopyProtectionOpts_Watcher(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockStore := mocks.NewMockCompliancePolicy(ctrl)

	opts := &CopyProtectionOpts{
		store: mockStore,
	}
	state := active
	expected := &atlasv2.DataProtectionSettings{
		State: &state,
	}

	mockStore.
		EXPECT().
		DescribeCompliancePolicy(opts.ProjectID).
		Return(expected, nil).
		Times(1)

	res, err := opts.copyProtectionWatcher()
	if err != nil {
		t.Fatalf("copyProtectionWatcher() unexpected error: %v", err)
	}
	assert.True(t, res)
}

func TestCopyProtectionOpts_Run(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockStore := mocks.NewMockCompliancePolicy(ctrl)
	state := active
	copyprotectionBefore := false
	copyprotectionAfter := true

	initial := &atlasv2.DataProtectionSettings{
		CopyProtectionEnabled: &copyprotectionBefore,
	}

	expected := &atlasv2.DataProtectionSettings{
		State:                 &state,
		CopyProtectionEnabled: &copyprotectionAfter,
	}

	opts := &CopyProtectionOpts{
		store:  mockStore,
		policy: initial,
		enable: true,
	}

	mockStore.
		EXPECT().
		UpdateCompliancePolicy(opts.ProjectID, opts.policy).
		Return(expected, nil).
		Times(1)

	assert.False(t, *opts.policy.CopyProtectionEnabled)

	if err := opts.Run(); err != nil {
		t.Fatalf("Run() unexpected error: %v", err)
	}
	assert.True(t, *opts.policy.CopyProtectionEnabled)
}
