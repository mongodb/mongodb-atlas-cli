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

package maintenance

import (
	"testing"

	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/pointer"
	atlasv2 "go.mongodb.org/atlas-sdk/v20250312025/admin"
	"go.uber.org/mock/gomock"
)

func TestUpdateOpts_Run(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockStore := NewMockUpdater(ctrl)

	updateOpts := &UpdateOpts{
		store:     mockStore,
		hourOfDay: 2,
		dayOfWeek: 1,
		startASAP: true,
	}
	updateOpts.ProjectID = "21321323343243243"

	mockStore.
		EXPECT().
		UpdateMaintenanceWindow("21321323343243243", &atlasv2.GroupMaintenanceWindow{
			DayOfWeek: pointer.Get(1),
			HourOfDay: pointer.Get(2),
			StartASAP: pointer.Get(true),
		}).
		Return(nil).
		Times(1)

	if err := updateOpts.Run(); err != nil {
		t.Fatalf("Run() unexpected error: %v", err)
	}
}
