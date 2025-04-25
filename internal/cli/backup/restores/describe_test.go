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

package restores

import (
	"testing"

	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/mocks"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/pointer"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/test"
	"github.com/stretchr/testify/require"
	atlasv2 "go.mongodb.org/atlas-sdk/v20250312002/admin"
	"go.uber.org/mock/gomock"
)

func TestDescribeOpts_Run(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockStore := mocks.NewMockRestoreJobsDescriber(ctrl)

	expected := &atlasv2.DiskBackupSnapshotRestoreJob{
		Id: pointer.Get("1"),
	}

	describeOpts := &DescribeOpts{
		store:       mockStore,
		clusterName: "Cluster0",
		id:          "1",
	}

	expectedError := &atlasv2.GenericOpenAPIError{}
	expectedError.SetModel(atlasv2.ApiError{ErrorCode: cannotUseNotFlexWithFlexApisErrorCode})

	mockStore.
		EXPECT().
		RestoreFlexClusterJob(describeOpts.ProjectID, describeOpts.clusterName, describeOpts.id).
		Return(nil, expectedError).
		Times(1)

	mockStore.
		EXPECT().
		RestoreJob(describeOpts.ProjectID, describeOpts.clusterName, describeOpts.id).
		Return(expected, nil).
		Times(1)

	require.NoError(t, describeOpts.Run())
	test.VerifyOutputTemplate(t, restoreDescribeTemplate, expected)
}

func TestDescribeOpts_Run_FlexCluster(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockStore := mocks.NewMockRestoreJobsDescriber(ctrl)

	expected := &atlasv2.FlexBackupRestoreJob20241113{
		Id: pointer.Get("1"),
	}

	describeOpts := &DescribeOpts{
		store:       mockStore,
		clusterName: "Cluster0",
		id:          "1",
	}

	mockStore.
		EXPECT().
		RestoreFlexClusterJob(describeOpts.ProjectID, describeOpts.clusterName, describeOpts.id).
		Return(nil, nil).
		Times(1)

	require.NoError(t, describeOpts.Run())

	test.VerifyOutputTemplate(t, restoreDescribeFlexClusterTemplate, expected)
}
