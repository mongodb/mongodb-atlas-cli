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

	"go.uber.org/mock/gomock"
)

func TestUpdateOpts_Run(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockStore := NewMockUpdater(ctrl)

	updateOpts := &UpdateOpts{
		store:     mockStore,
		hourOfDay: 2,
		dayOfWeek: 1,
	}
	updateOpts.ProjectID = "21321323343243243"

	mockStore.
		EXPECT().
		UpdateMaintenanceWindow(updateOpts.ConfigProjectID(), updateOpts.newMaintenanceWindow()).
		Return(nil).
		Times(1)

	if err := updateOpts.Run(); err != nil {
		t.Fatalf("Run() unexpected error: %v", err)
	}
}

func TestUpdateOpts_newMaintenanceWindow(t *testing.T) {
	updateOpts := &UpdateOpts{
		dayOfWeek: 1,
		hourOfDay: 2,
		startASAP: true,
	}

	window := updateOpts.newMaintenanceWindow()

	if window.DayOfWeek == nil || *window.DayOfWeek != 1 {
		t.Errorf("newMaintenanceWindow() DayOfWeek = %v, want 1", window.DayOfWeek)
	}
	if window.HourOfDay == nil || *window.HourOfDay != 2 {
		t.Errorf("newMaintenanceWindow() HourOfDay = %v, want 2", window.HourOfDay)
	}
	if window.StartASAP == nil || !*window.StartASAP {
		t.Errorf("newMaintenanceWindow() StartASAP = %v, want true", window.StartASAP)
	}
}
