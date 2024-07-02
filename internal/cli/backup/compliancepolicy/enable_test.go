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
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/flag"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/mocks"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/test"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	atlasv2 "go.mongodb.org/atlas-sdk/v20240530002/admin"
)

const (
	authorizedUserFirstName = "firstname"
	authorizedUserLastName  = "lastname"
	authorizedEmail         = "firstname.lastname@example.com"
)

func TestEnableBuilder(t *testing.T) {
	test.CmdValidator(
		t,
		EnableBuilder(),
		0,
		[]string{
			flag.ProjectID,
			flag.AuthorizedUserFirstName,
			flag.AuthorizedUserLastName,
			flag.AuthorizedEmail,
			flag.Output,
			flag.EnableWatch,
		},
	)
}

func TestEnableOpts_Watcher(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockStore := mocks.NewMockCompliancePolicyEnabler(ctrl)

	opts := &EnableOpts{
		store:   mockStore,
		confirm: true,
	}
	state := active
	expected := &atlasv2.DataProtectionSettings20231001{
		State: &state,
	}

	mockStore.
		EXPECT().
		DescribeCompliancePolicy(opts.ProjectID).
		Return(expected, nil).
		Times(1)

	_, res, err := opts.enableWatcher()
	if err != nil {
		t.Fatalf("enableWatcher() unexpected error: %v", err)
	}
	assert.True(t, res)
}

func TestEnableOpts_Run(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockStore := mocks.NewMockCompliancePolicyEnabler(ctrl)
	state := active

	expected := &atlasv2.DataProtectionSettings20231001{
		State: &state,
	}

	opts := &EnableOpts{
		store:                   mockStore,
		authorizedUserFirstName: authorizedUserFirstName,
		authorizedUserLastName:  authorizedUserLastName,
		authorizedEmail:         authorizedEmail,
		confirm:                 true,
	}

	mockStore.
		EXPECT().
		EnableCompliancePolicy(opts.ProjectID, authorizedEmail, authorizedUserFirstName, authorizedUserLastName).
		Return(expected, nil).
		Times(1)

	if err := opts.Run(); err != nil {
		t.Fatalf("Run() unexpected error: %v", err)
	}
	test.VerifyOutputTemplate(t, enableTemplate, expected)
}

func TestEnableOpts_Run_invalidEmail(t *testing.T) {
	invalidEmail := "invalidEmail"

	opts := &EnableOpts{
		authorizedEmail: invalidEmail,
	}

	require.Error(t, opts.Run())
}
