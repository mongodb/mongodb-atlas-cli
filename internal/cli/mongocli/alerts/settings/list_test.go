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

package settings

import (
	"bytes"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/mongodb/mongodb-atlas-cli/mongocli/v2/internal/cli"
	"github.com/mongodb/mongodb-atlas-cli/mongocli/v2/internal/mocks"
	"github.com/mongodb/mongodb-atlas-cli/mongocli/v2/internal/pointer"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/atlas/mongodbatlas"
)

func TestConfigList_Run(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockStore := mocks.NewMockAlertConfigurationLister(ctrl)

	expected := []mongodbatlas.AlertConfiguration{
		{
			ID:                     "test",
			GroupID:                "test",
			AlertConfigID:          "test",
			EventTypeName:          "test",
			Created:                "test",
			Status:                 "test",
			AcknowledgedUntil:      "test",
			AcknowledgementComment: "test",
			AcknowledgingUsername:  "test",
			Updated:                "test",
			Resolved:               "test",
			LastNotified:           "test",
			HostnameAndPort:        "test",
			HostID:                 "test",
			ReplicaSetName:         "test",
			MetricName:             "test",
			Enabled:                pointer.Get(true),
			ClusterID:              "test",
			ClusterName:            "test",
			SourceTypeName:         "test",
			CurrentValue: &mongodbatlas.CurrentValue{
				Number: pointer.Get(1.2),
				Units:  "test",
			},
			Matchers:        nil,
			MetricThreshold: nil,
			Threshold:       nil,
			Notifications:   nil,
		},
	}

	buf := new(bytes.Buffer)

	listOpts := &ListOpts{
		store: mockStore,
		OutputOpts: cli.OutputOpts{
			Template:  settingsListTemplate,
			OutWriter: buf,
		},
	}

	mockStore.
		EXPECT().
		AlertConfigurations(listOpts.ProjectID, listOpts.NewListOptions()).
		Return(expected, nil).
		Times(1)

	if err := listOpts.Run(); err != nil {
		t.Fatalf("Run() unexpected error: %v", err)
	}

	assert.Equal(t, `ID     TYPE   ENABLED
test   test   true
`, buf.String())
	t.Log(buf.String())
}
