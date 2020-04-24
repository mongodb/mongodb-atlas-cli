// Copyright 2020 MongoDB Inc
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
package cli

import (
	"testing"

	"github.com/golang/mock/gomock"
	om "github.com/mongodb/go-client-mongodb-ops-manager/opsmngr"
	"github.com/mongodb/mongocli/internal/mocks"
)

func TestOpsManagerLogsCollectOpts_Run(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockStore := mocks.NewMockLogs(ctrl)

	defer ctrl.Finish()

	expected := &om.LogCollectionJob{ID: "1"}

	listOpts := &opsManagerLogsCollectOpts{
		redacted:                  false,
		sizeRequestedPerFileBytes: 64,
		resourceType:              "CLUSTER",
		resourceName:              "",
		logTypes:                  []string{"AUTOMATION_AGENT"},
		store:                     mockStore,
	}

	mockStore.
		EXPECT().Collect(listOpts.projectID, listOpts.newLog()).
		Return(expected, nil).
		Times(1)

	if err := listOpts.Run(); err != nil {
		t.Fatalf("Run() unexpected error: %v", err)
	}
}
