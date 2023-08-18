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

package update

import (
	"context"
	"errors"
	"net/http"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/mongodb/mongodb-atlas-cli/internal/flag"
	mocks "github.com/mongodb/mongodb-atlas-cli/internal/mocks/atlas"
	"github.com/mongodb/mongodb-atlas-cli/internal/test"
	"github.com/stretchr/testify/assert"
	atlasv2 "go.mongodb.org/atlas-sdk/v20230201004/admin"
)

type MockUpdateStore struct {
	*mocks.MockCompliancePolicyDescriber
	*mocks.MockCompliancePolicyItemUpdater
	*mocks.MockProjectLister
}

func NewMockUpdateStore(ctrl *gomock.Controller) *MockUpdateStore {
	return &MockUpdateStore{
		MockCompliancePolicyDescriber:   mocks.NewMockCompliancePolicyDescriber(ctrl),
		MockCompliancePolicyItemUpdater: mocks.NewMockCompliancePolicyItemUpdater(ctrl),
		MockProjectLister:               mocks.NewMockProjectLister(ctrl),
	}
}

func TestEnableBuilder(t *testing.T) {
	test.CmdValidator(
		t,
		Builder(),
		0,
		[]string{
			flag.ProjectID,
			flag.Output,
			flag.EnableWatch,
			flag.File,
		},
	)
}

func TestInitStore(t *testing.T) {
	opts := &Opts{}
	ctx := context.Background()

	if err := opts.initStore(ctx)(); err != nil {
		t.Fatalf("initStore()() unexpected error: %v", err)
	}
	assert.NotNil(t, opts.store)
}

func TestEnableOpts_Run(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockStore := NewMockUpdateStore(ctrl)

	opts := &Opts{
		store: mockStore,
	}

	policyItem := atlasv2.NewDiskBackupApiPolicyItem(1, "daily", "days", 1)

	expected := atlasv2.NewDataProtectionSettings()

	mockStore.
		MockCompliancePolicyItemUpdater.
		EXPECT().
		UpdatePolicyItem(opts.projectID, policyItem).
		Return(expected, nil, nil).
		Times(1)

	err := opts.Run(policyItem)
	if err != nil {
		t.Fatalf("Run() unexpected error: %v", err)
	}
}

func TestEnableOpts_Run_Fail_code_500(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockStore := NewMockUpdateStore(ctrl)

	opts := &Opts{
		store: mockStore,
	}

	policyItem := atlasv2.NewDiskBackupApiPolicyItem(1, "daily", "days", 1)

	httpResponse := &http.Response{
		StatusCode: http.StatusInternalServerError,
	}

	mockError := errors.New("network error")

	mockStore.
		MockCompliancePolicyItemUpdater.
		EXPECT().
		UpdatePolicyItem(opts.projectID, policyItem).
		Return(nil, httpResponse, mockError).
		Times(1)

	err := opts.Run(policyItem)
	assert.ErrorContains(t, err, errorCode500Template)
}
