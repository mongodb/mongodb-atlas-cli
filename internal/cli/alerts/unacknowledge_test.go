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

package alerts

import (
	"testing"

	atlasv2 "go.mongodb.org/atlas-sdk/v20250312005/admin"
	"go.uber.org/mock/gomock"
)

func TestUnacknowledge_Run(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockStore := NewMockAlertAcknowledger(ctrl)

	expected := &atlasv2.AlertViewForNdsGroup{}

	acknowledgeOpts := &UnacknowledgeOpts{
		alertID: "533dc40ae4b00835ff81eaee",
		comment: "Test",
		store:   mockStore,
	}

	ackReq := acknowledgeOpts.newUnacknowledgeRequest()
	params := &atlasv2.AcknowledgeAlertApiParams{
		GroupId:          acknowledgeOpts.ProjectID,
		AlertId:          acknowledgeOpts.alertID,
		AcknowledgeAlert: ackReq,
	}

	mockStore.
		EXPECT().
		AcknowledgeAlert(params).
		Return(expected, nil).
		Times(1)

	if err := acknowledgeOpts.Run(); err != nil {
		t.Fatalf("Run() unexpected error: %v", err)
	}
}
