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

// +build unit

package automation

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/mongodb/mongocli/internal/mocks"
	"github.com/spf13/afero"
	"go.mongodb.org/ops-manager/opsmngr"
)

func TestAutomationUpdate_Run(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockStore := mocks.NewMockAutomationUpdater(ctrl)

	defer ctrl.Finish()

	expected := &opsmngr.AutomationConfig{
		Version: 1,
	}
	appFS := afero.NewMemMapFs()
	// create test file
	fileJSON := `
{
"version": 1
}`
	fileName := "om_automation_test.json"
	_ = afero.WriteFile(appFS, fileName, []byte(fileJSON), 0600)
	createOpts := &UpdateOpts{
		store:    mockStore,
		fs:       appFS,
		filename: fileName,
	}

	mockStore.
		EXPECT().
		UpdateAutomationConfig(createOpts.ProjectID, expected).
		Return(nil).
		Times(1)

	err := createOpts.Run()
	if err != nil {
		t.Fatalf("Run() unexpected error: %v", err)
	}
}
