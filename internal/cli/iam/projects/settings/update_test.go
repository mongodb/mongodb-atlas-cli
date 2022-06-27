// Copyright 2022 MongoDB Inc
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
// +build unit

package settings

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/mongodb/mongodb-atlas-cli/internal/flag"
	"github.com/mongodb/mongodb-atlas-cli/internal/mocks"
	"github.com/mongodb/mongodb-atlas-cli/internal/test"
	"github.com/stretchr/testify/assert"
	atlas "go.mongodb.org/atlas/mongodbatlas"
)

func TestUpdateOpts_Run(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockStore := mocks.NewMockProjectSettingsGetterUpdater(ctrl)
	defer ctrl.Finish()

	opts := &UpdateOpts{
		store:                                    mockStore,
		enableCollectDatabaseSpecificsStatistics: true,
		disableCollectDatabaseSpecificsStatistics: false,
		enableDataExplorer:                        false,
		disableDataExplorer:                       true,
		enablePerformanceAdvisor:                  true,
		disablePerformanceAdvisor:                 false,
		enableSchemaAdvisor:                       true,
		disableSchemaAdvisor:                      false,
		enableRealtimePerformancePanel:            true,
		disableRealtimePerformancePanel:           false,
		projectSettings: &atlas.ProjectSettings{
			IsCollectDatabaseSpecificsStatisticsEnabled: returnMockValueForSetting(true),
			IsDataExplorerEnabled:                       returnMockValueForSetting(true),
			IsPerformanceAdvisorEnabled:                 returnMockValueForSetting(false),
			IsRealtimePerformancePanelEnabled:           returnMockValueForSetting(false),
			IsSchemaAdvisorEnabled:                      returnMockValueForSetting(true),
		},
	}

	expected := &atlas.ProjectSettings{}

	mockStore.
		EXPECT().
		UpdateProjectSettings(opts.ProjectID, opts.projectSettings).
		Return(expected, nil).
		Times(1)

	err := opts.Run()
	assert.NoError(t, err)
}

func TestUpdateBuilder(t *testing.T) {
	test.CmdValidator(
		t,
		UpdateBuilder(),
		0,
		[]string{flag.ProjectID, flag.Output},
	)
}

func returnMockValueForSetting(value bool) *bool {
	var valueToSet bool
	if value == true {
		valueToSet = true
		return &valueToSet
	}
	valueToSet = false
	return &valueToSet
}
