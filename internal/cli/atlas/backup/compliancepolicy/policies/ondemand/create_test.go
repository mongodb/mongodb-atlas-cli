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

package ondemand

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/mongodb/mongodb-atlas-cli/internal/flag"
	"github.com/mongodb/mongodb-atlas-cli/internal/mocks/atlas"
	"github.com/mongodb/mongodb-atlas-cli/internal/test"
	atlasv2 "go.mongodb.org/atlas-sdk/v20231115005/admin"
)

func TestCreateOpts_Run(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockStore := atlas.NewMockCompliancePolicyOnDemandPolicyCreator(ctrl)

	createOpts := &CreateOpts{
		store:          mockStore,
		retentionUnit:  "days",
		retentionValue: 30,
	}

	policyItem := &atlasv2.BackupComplianceOnDemandPolicyItem{
		FrequencyType:  onDemandFrequencyType,
		RetentionUnit:  createOpts.retentionUnit,
		RetentionValue: createOpts.retentionValue,
	}

	expected := &atlasv2.DataProtectionSettings20231001{}

	mockStore.
		EXPECT().
		CreateOnDemandPolicy("", policyItem).Return(expected, nil).
		Times(1)

	if err := createOpts.Run(); err != nil {
		t.Fatalf("Run() unexpected error: %v", err)
	}
}

func TestCreateBuilder(t *testing.T) {
	test.CmdValidator(
		t,
		CreateBuilder(),
		0,
		[]string{flag.RetentionUnit, flag.RetentionValue, flag.EnableWatch, flag.ProjectID, flag.Output},
	)
}
func TestCreateTemplate(t *testing.T) {
	test.VerifyOutputTemplate(t, updateTemplate, &atlasv2.DataProtectionSettings20231001{})
}
