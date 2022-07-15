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

package schedule

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/mongodb/mongodb-atlas-cli/internal/cli"
	"github.com/mongodb/mongodb-atlas-cli/internal/flag"
	"github.com/mongodb/mongodb-atlas-cli/internal/mocks"
	"github.com/mongodb/mongodb-atlas-cli/internal/test"
	"github.com/openlyinc/pointy"
	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
	atlas "go.mongodb.org/atlas/mongodbatlas"
)

func TestUpdateOpts_Run(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockStore := mocks.NewMockScheduleDescriberUpdater(ctrl)

	opts := &UpdateOpts{
		store:                               mockStore,
		clusterName:                         "Test",
		exportBucketID:                      "604f6322dc786a5341d4f7fb",
		exportFrequencyType:                 "monthly",
		backupPolicy:                        []string{},
		referenceHourOfDay:                  12,
		referenceMinuteOfHour:               30,
		restoreWindowDays:                   5,
		autoExport:                          true,
		noAutoExport:                        false,
		updateSnapshots:                     true,
		noUpdateSnapshots:                   false,
		useOrgAndGroupNamesInExportPrefix:   true,
		noUseOrgAndGroupNamesInExportPrefix: false,
	}

	expected := &atlas.CloudProviderSnapshotBackupPolicy{}
	cmd := &cobra.Command{}

	mockStore.
		EXPECT().
		UpdateSchedule(opts.ProjectID, opts.clusterName, gomock.Any()).
		Return(expected, nil).
		Times(1)

	err := opts.Run(cmd)
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

func TestReturnMockValueForSetting(t *testing.T) {
	tests := []struct {
		name             string
		inputEnableFlag  bool
		inputDisableFlag bool
		want             *bool
	}{
		{
			name:             "both true",
			inputEnableFlag:  true,
			inputDisableFlag: true,
			want:             nil,
		},
		{
			name:             "enable only",
			inputEnableFlag:  true,
			inputDisableFlag: false,
			want:             pointy.Bool(true),
		},
		{
			name:             "disable only",
			inputEnableFlag:  false,
			inputDisableFlag: true,
			want:             pointy.Bool(false),
		},
		{
			name:             "both false",
			inputEnableFlag:  false,
			inputDisableFlag: false,
			want:             nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := cli.ReturnValueForSetting(tt.inputEnableFlag, tt.inputDisableFlag)
			assert.Equalf(t, tt.want, got, "returnValueForSetting()")
		})
	}
}
