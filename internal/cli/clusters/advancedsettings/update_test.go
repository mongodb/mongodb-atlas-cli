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

package advancedsettings

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/flag"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/mocks"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/pointer"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/test"
	atlasv2 "go.mongodb.org/atlas-sdk/v20240530002/admin"
)

func TestUpdate_Run(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockStore := mocks.NewMockAtlasClusterConfigurationOptionsUpdater(ctrl)

	expected := &atlasv2.ClusterDescriptionProcessArgs{
		DefaultReadConcern:               pointer.Get(""),
		DefaultWriteConcern:              pointer.Get(""),
		MinimumEnabledTlsProtocol:        pointer.Get(""),
		SampleSizeBIConnector:            pointer.Get[int](1000),
		SampleRefreshIntervalBIConnector: pointer.Get[int](0),
		NoTableScan:                      pointer.Get(false),
	}

	t.Run("flags run", func(t *testing.T) {
		updateOpts := &UpdateOpts{
			name:  "ProjectBar",
			store: mockStore,
		}

		mockStore.
			EXPECT().
			UpdateAtlasClusterConfigurationOptions(updateOpts.ProjectID, updateOpts.name, updateOpts.newProcessArgs()).
			Return(expected, nil).
			Times(1)

		if err := updateOpts.Run(); err != nil {
			t.Fatalf("Run() unexpected error: %v", err)
		}
	})
}

func TestUpdateBuilder(t *testing.T) {
	test.CmdValidator(
		t,
		UpdateBuilder(),
		0,
		[]string{
			flag.ReadConcern, flag.WriteConcern, flag.TLSProtocol, flag.DisableFailIndexKeyTooLong, flag.EnableFailIndexKeyTooLong,
			flag.OplogMinRetentionHours, flag.EnableJavascript, flag.EnableTableScan,
			flag.DisableJavascript, flag.OplogMinRetentionHours, flag.OplogSizeMB, flag.SampleRefreshIntervalBIConnector,
			flag.SampleSizeBIConnector, flag.DisableTableScan, flag.ProjectID, flag.Output},
	)
}
