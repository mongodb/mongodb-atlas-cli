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

package schedule

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/flag"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/mocks"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/pointer"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/test"
	"github.com/spf13/afero"
	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	atlasv2 "go.mongodb.org/atlas-sdk/v20240530002/admin"
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
		filename:                            "",
		fs:                                  afero.NewMemMapFs(),
	}

	expected := &atlasv2.DiskBackupSnapshotSchedule{}
	cmd := &cobra.Command{}

	mockStore.
		EXPECT().
		UpdateSchedule(opts.ProjectID, opts.clusterName, gomock.Any()).
		Return(expected, nil).
		Times(1)

	err := opts.Run(cmd)
	require.NoError(t, err)
}

func TestUpdateOpts_RunWithFile(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockStore := mocks.NewMockScheduleDescriberUpdater(ctrl)
	fs := afero.NewMemMapFs()

	fileContents := `
{
  "autoExportEnabled": true,
  "export": {
    "exportBucketId": "604f6322dc786a5341d4f7fb",
    "frequencyType": "monthly"
  },
  "policies": [],
  "referenceHourOfDay": 12,
  "referenceMinuteOfHour": 30,
  "restoreWindowDays": 5,
  "updateSnapshots": true,
  "useOrgAndGroupNamesInExportPrefix": true
}`

	fileName := "test.json"
	require.NoError(t, afero.WriteFile(fs, fileName, []byte(fileContents), 0600))

	opts := &UpdateOpts{
		store:       mockStore,
		clusterName: "Test",
		filename:    fileName,
		fs:          fs,
	}

	expected := &atlasv2.DiskBackupSnapshotSchedule{}
	cmd := &cobra.Command{}

	mockStore.
		EXPECT().
		UpdateSchedule(opts.ProjectID, opts.clusterName, gomock.Any()).
		Return(expected, nil).
		Times(1)

	err := opts.Run(cmd)
	require.NoError(t, err)
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
			want:             pointer.Get(true),
		},
		{
			name:             "disable only",
			inputEnableFlag:  false,
			inputDisableFlag: true,
			want:             pointer.Get(false),
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
