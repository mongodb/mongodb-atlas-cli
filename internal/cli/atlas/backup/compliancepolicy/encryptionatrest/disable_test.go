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

package encryptionatrest

import (
	"context"
	"testing"

	"github.com/andreaangiolillo/mongocli-test/internal/cli"
	"github.com/andreaangiolillo/mongocli-test/internal/flag"
	mocks "github.com/andreaangiolillo/mongocli-test/internal/mocks/atlas"
	"github.com/andreaangiolillo/mongocli-test/internal/test"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	atlasv2 "go.mongodb.org/atlas-sdk/v20231115002/admin"
)

func TestDisableBuilder(t *testing.T) {
	test.CmdValidator(
		t,
		DisableBuilder(),
		0,
		[]string{
			flag.ProjectID,
			flag.Output,
			flag.EnableWatch,
		},
	)
}

func TestDisableOpts_InitStore(t *testing.T) {
	opts := &DisableOpts{}

	require.NoError(t, opts.initStore(context.TODO())())
	assert.NotNil(t, opts.store)
}

func TestDisableOpts_Watcher(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockStore := mocks.NewMockCompliancePolicyEncryptionAtRestDisabler(ctrl)

	opts := &DisableOpts{
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

	res, err := opts.watcher()
	require.NoError(t, err)
	assert.True(t, res)
}

func TestDisableOpts_Run(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockStore := mocks.NewMockCompliancePolicyEncryptionAtRestDisabler(ctrl)
	encryptionAtRestAfter := false

	expected := &atlasv2.DataProtectionSettings20231001{
		EncryptionAtRestEnabled: &encryptionAtRestAfter,
	}

	opts := &DisableOpts{
		store: mockStore,
	}

	mockStore.
		EXPECT().
		DisableEncryptionAtRest(opts.ProjectID).
		Return(expected, nil).
		Times(1)

	if err := opts.Run(); err != nil {
		t.Fatalf("Run() unexpected error: %v", err)
	}
	assert.False(t, *opts.policy.EncryptionAtRestEnabled)
	test.VerifyOutputTemplate(t, disableTemplate, expected)
}

func TestDisableOpts_WatchRun(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockStore := mocks.NewMockCompliancePolicyEncryptionAtRestDisabler(ctrl)

	opts := &DisableOpts{
		store: mockStore,
		WatchOpts: cli.WatchOpts{
			EnableWatch: true,
		},
	}

	expected := &atlasv2.DataProtectionSettings20231001{
		State: atlasv2.PtrString(active),
	}

	mockStore.
		EXPECT().
		DisableEncryptionAtRest(opts.ProjectID).
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

	test.VerifyOutputTemplate(t, disableWatchTemplate, expected)
}
