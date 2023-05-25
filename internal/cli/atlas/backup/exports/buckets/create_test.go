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

package buckets

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/mongodb/mongodb-atlas-cli/internal/flag"
	"github.com/mongodb/mongodb-atlas-cli/internal/mocks"
	"github.com/mongodb/mongodb-atlas-cli/internal/test"
	atlasv2 "go.mongodb.org/atlas-sdk/admin"
)

func TestCreateOpts_Run(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockStore := mocks.NewMockExportBucketsCreator(ctrl)

	expected := &atlasv2.DiskBackupSnapshotAWSExportBucket{}

	createOpts := &CreateOpts{
		bucketName:    "testBucket",
		iamRoleID:     "12345678f901a234dbdb00ca",
		cloudProvider: "AWS",
		store:         mockStore,
	}

	bucket := createOpts.newExportBucket()
	mockStore.
		EXPECT().
		CreateExportBucket("", bucket).Return(expected, nil).
		Times(1)

	if err := createOpts.Run(); err != nil {
		t.Fatalf("Run() unexpected error: %v", err)
	}
}

func TestCreateBuilder(t *testing.T) {
	test.CmdValidator(
		t,
		CreateBuilder(),
		0,
		[]string{flag.CloudProvider, flag.IAMRoleID, flag.ProjectID, flag.Output},
	)
}
