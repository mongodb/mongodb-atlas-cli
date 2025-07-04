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

package databases

import (
	"testing"

	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli"
	atlasv2 "go.mongodb.org/atlas-sdk/v20250312005/admin"
	"go.uber.org/mock/gomock"
)

func TestDatabasesDescribeOpts_Run(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockStore := NewMockProcessDatabaseMeasurementsLister(ctrl)

	listOpts := &DescribeOpts{
		host:  "hard-00-00.mongodb.net",
		port:  27017,
		name:  "test",
		store: mockStore,
		ProjectOpts: cli.ProjectOpts{
			ProjectID: "groupID",
		},
	}

	params := listOpts.NewDatabaseMeasurementsAPIParams("groupID", "hard-00-00.mongodb.net:27017", listOpts.name)
	expected := &atlasv2.ApiMeasurementsGeneralViewAtlas{
		Measurements: &[]atlasv2.MetricsMeasurementAtlas{},
	}
	mockStore.
		EXPECT().ProcessDatabaseMeasurements(params).
		Return(expected, nil).
		Times(1)

	if err := listOpts.Run(); err != nil {
		t.Fatalf("Run() unexpected error: %v", err)
	}
}
