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

	"github.com/golang/mock/gomock"
	"github.com/mongodb/mongodb-atlas-cli/mongocli/v2/internal/mocks"
	"go.mongodb.org/ops-manager/opsmngr"
)

func TestUnacknowledge_Run(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockStore := mocks.NewMockAlertAcknowledger(ctrl)

	expected := &opsmngr.Alert{}

	acknowledgeOpts := &UnacknowledgeOpts{
		alertID: "533dc40ae4b00835ff81eaee",
		comment: "Test",
		store:   mockStore,
	}

	ackReq := acknowledgeOpts.newAcknowledgeRequest()

	mockStore.
		EXPECT().
		AcknowledgeAlert(acknowledgeOpts.ProjectID, acknowledgeOpts.alertID, ackReq).
		Return(expected, nil).
		Times(1)

	if err := acknowledgeOpts.Run(); err != nil {
		t.Fatalf("Run() unexpected error: %v", err)
	}
}
