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
	"time"

	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/pointer"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/test"
	atlasv2 "go.mongodb.org/atlas-sdk/v20250312005/admin"
	"go.uber.org/mock/gomock"
)

func TestConfigList_Run(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockStore := NewMockAlertConfigurationLister(ctrl)

	expected := &atlasv2.PaginatedAlertConfig{
		Results: &[]atlasv2.GroupAlertsConfig{
			{
				Id:              pointer.Get("test"),
				GroupId:         pointer.Get("test"),
				Enabled:         pointer.Get(true),
				EventTypeName:   pointer.Get("test"),
				Created:         pointer.Get(time.Now()),
				Matchers:        nil,
				Notifications:   nil,
				Updated:         pointer.Get(time.Now()),
				MetricThreshold: nil,
				Threshold:       nil,
			},
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

	params := &atlasv2.ListAlertConfigurationsApiParams{
		GroupId: listOpts.ProjectID,
		PageNum: &listOpts.PageNum,
	}

	mockStore.
		EXPECT().
		AlertConfigurations(params).
		Return(expected, nil).
		Times(1)

	if err := listOpts.Run(); err != nil {
		t.Fatalf("Run() unexpected error: %v", err)
	}

	t.Log(buf.String())
	test.VerifyOutputTemplate(t, settingsListTemplate, expected)
}
