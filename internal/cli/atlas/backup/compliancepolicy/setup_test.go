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
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/mongodb/mongodb-atlas-cli/internal/cli"
	"github.com/mongodb/mongodb-atlas-cli/internal/flag"
	mocks "github.com/mongodb/mongodb-atlas-cli/internal/mocks/atlas"
	"github.com/mongodb/mongodb-atlas-cli/internal/test"
	"github.com/stretchr/testify/assert"
	atlasv2 "go.mongodb.org/atlas-sdk/v20230201008/admin"
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

	expected := &atlasv2.DataProtectionSettings{
		State: &state,
	}

	mockStore.
		EXPECT().
		DescribeCompliancePolicy(opts.ProjectID).
		Return(expected, nil).
		Times(1)

	res, err := opts.setupWatcher()
	if err != nil {
		t.Fatalf("setupWatcher() unexpected error: %v", err)
	}
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
		policy:  new(atlasv2.DataProtectionSettings),
	}

	expected := &atlasv2.DataProtectionSettings{
		State: &state,
	}

	mockStore.
		EXPECT().
		UpdateCompliancePolicy(opts.ProjectID, opts.policy).
		Return(expected, nil).
		Times(1)

	if err := opts.Run(); err != nil {
		t.Fatalf("run() unexpected error: %v", err)
	}

	test.VerifyOutputTemplate(t, setupTemplate, expected)
}

// Verifies the output template when using --watch.
func TestSetupOpts_WatchRun(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockStore := mocks.NewMockCompliancePolicyUpdater(ctrl)
	state := active

	opts := &SetupOpts{
		store:   mockStore,
		confirm: true,
		policy:  new(atlasv2.DataProtectionSettings),
		WatchOpts: cli.WatchOpts{
			EnableWatch: true,
		},
	}

	expected := &atlasv2.DataProtectionSettings{
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

	if err := opts.Run(); err != nil {
		t.Fatalf("run() unexpected error: %v", err)
	}

	test.VerifyOutputTemplate(t, setupWatchTemplate, expected)
}
