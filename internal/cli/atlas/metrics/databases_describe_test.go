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

package metrics

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/mongodb/mongocli/internal/mocks"
	"go.mongodb.org/atlas/mongodbatlas"
)

func TestDatabasesDescribeOpts_Run(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockStore := mocks.NewMockProcessDatabaseMeasurementsLister(ctrl)
	defer ctrl.Finish()

	listOpts := &DatabasesDescribeOpts{
		host:  "hard-00-00.mongodb.net",
		port:  27017,
		name:  "test",
		store: mockStore,
	}

	opts := listOpts.NewProcessMetricsListOptions()
	expected := &mongodbatlas.ProcessDatabaseMeasurements{
		ProcessMeasurements: &mongodbatlas.ProcessMeasurements{
			Measurements: []*mongodbatlas.Measurements{},
		},
	}
	mockStore.
		EXPECT().ProcessDatabaseMeasurements(listOpts.ProjectID, listOpts.host, listOpts.port, listOpts.name, opts).
		Return(expected, nil).
		Times(1)

	if err := listOpts.Run(); err != nil {
		t.Fatalf("Run() unexpected error: %v", err)
	}
}
