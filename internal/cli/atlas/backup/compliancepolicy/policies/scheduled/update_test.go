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

package scheduled

import (
	"testing"

	"github.com/andreaangiolillo/mongocli-test/internal/flag"
	"github.com/andreaangiolillo/mongocli-test/internal/mocks/atlas"
	"github.com/andreaangiolillo/mongocli-test/internal/test"
	"github.com/golang/mock/gomock"
	atlasv2 "go.mongodb.org/atlas-sdk/v20231115002/admin"
)

func TestUpdateOpts_Run(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockStore := atlas.NewMockCompliancePolicyScheduledPolicyUpdater(ctrl)

	opts := &UpdateOpts{
		store:             mockStore,
		scheduledPolicyID: "123",
		frequencyType:     "weekly",
		frequencyInterval: 1,
		retentionUnit:     "days",
		retentionValue:    30,
	}

	policyItem := &atlasv2.BackupComplianceScheduledPolicyItem{
		Id:                &opts.scheduledPolicyID,
		FrequencyType:     opts.frequencyType,
		FrequencyInterval: opts.frequencyInterval,
		RetentionUnit:     opts.retentionUnit,
		RetentionValue:    opts.retentionValue,
	}

	expected := &atlasv2.DataProtectionSettings20231001{}

	mockStore.
		EXPECT().
		UpdateScheduledPolicy("", policyItem).Return(expected, nil).
		Times(1)

	if err := opts.Run(); err != nil {
		t.Fatalf("Run() unexpected error: %v", err)
	}
}

func TestUpdateBuilder(t *testing.T) {
	test.CmdValidator(
		t,
		UpdateBuilder(),
		0,
		[]string{flag.ScheduledPolicyID, flag.FrequencyType, flag.FrequencyInterval, flag.RetentionUnit, flag.RetentionValue, flag.EnableWatch, flag.ProjectID, flag.Output},
	)
}

func TestUpdateTemplate(t *testing.T) {
	test.VerifyOutputTemplate(t, updateTemplate, &atlasv2.DataProtectionSettings20231001{})
}
