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

//go:build unit

package maintenance

import (
	"testing"

	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli"
	"go.uber.org/mock/gomock"
)

func TestClearOpts_Run(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockStore := NewMockClearer(ctrl)

	updateOpts := &ClearOpts{
		store: mockStore,
		ProjectOpts: cli.ProjectOpts{
			ProjectID: "21321323343243243",
		},
	}

	mockStore.
		EXPECT().
		ClearMaintenanceWindow(updateOpts.ConfigProjectID()).
		Return(nil).
		Times(1)

	if err := updateOpts.Run(); err != nil {
		t.Fatalf("Run() unexpected error: %v", err)
	}
}
