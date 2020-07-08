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

package atlas

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/mongodb/mongocli/internal/mocks"
	"go.mongodb.org/atlas/mongodbatlas"
)

func TestDataLakeUpdate_Run(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockStore := mocks.NewMockDataLakeUpdater(ctrl)

	defer ctrl.Finish()

	expected := mongodbatlas.DataLake{
		CloudProviderConfig: mongodbatlas.CloudProviderConfig{},
		DataProcessRegion:   mongodbatlas.DataProcessRegion{},
		GroupID:             "",
		Hostnames:           nil,
		Name:                "",
		State:               "",
		Storage:             mongodbatlas.Storage{},
	}

	createOpts := &DataLakeUpdateOpts{
		store:      mockStore,
		Name:       "new_data_lake",
		Region:     "some_region",
		Role:       "some::arn",
		TestBucket: "some_bucket",
	}

	updateRequest := &mongodbatlas.DataLakeUpdateRequest{
		CloudProviderConfig: mongodbatlas.CloudProviderConfig{
			AWSConfig: mongodbatlas.AwsCloudProviderConfig{
				IAMAssumedRoleARN: "some::arn",
				TestS3Bucket:      "some_bucket",
			},
		},
		DataProcessRegion: mongodbatlas.DataProcessRegion{
			CloudProvider: AWS,
			Region:        "some_region",
		},
	}

	mockStore.
		EXPECT().
		UpdateDataLake(createOpts.ProjectID, createOpts.Name, updateRequest).
		Return(&expected, nil).
		Times(1)

	err := createOpts.Run()
	if err != nil {
		t.Fatalf("Run() unexpected error: %v", err)
	}
}
