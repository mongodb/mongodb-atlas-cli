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

package restores

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/mocks"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/test"
	"github.com/stretchr/testify/require"
	atlasv2 "go.mongodb.org/atlas-sdk/v20250312001/admin"
)

func TestListOpts_Run(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockStore := mocks.NewMockRestoreJobsLister(ctrl)

	expected := &atlasv2.PaginatedCloudBackupRestoreJob{}

	listOpts := &ListOpts{
		store:       mockStore,
		clusterName: "Cluster0",
	}

	expectedError := &atlasv2.GenericOpenAPIError{}
	expectedError.SetModel(atlasv2.ApiError{ErrorCode: cannotUseNotFlexWithFlexApisErrorCode})

	mockStore.
		EXPECT().
		RestoreFlexClusterJobs(listOpts.newListFlexBackupRestoreJobsAPIParams()).
		Return(nil, expectedError).
		Times(1)

	mockStore.
		EXPECT().
		RestoreJobs(listOpts.ProjectID, "Cluster0", listOpts.NewListOptions()).
		Return(expected, nil).
		Times(1)

	require.NoError(t, listOpts.Run())
	test.VerifyOutputTemplate(t, restoreListTemplate, expected)
}

func TestListOpts_Run_FlexCluster(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockStore := mocks.NewMockRestoreJobsLister(ctrl)

	expected := &atlasv2.PaginatedApiAtlasFlexBackupRestoreJob20241113{}

	listOpts := &ListOpts{
		store:       mockStore,
		clusterName: "Cluster0",
	}

	mockStore.
		EXPECT().
		RestoreFlexClusterJobs(listOpts.newListFlexBackupRestoreJobsAPIParams()).
		Return(nil, nil).
		Times(1)

	require.NoError(t, listOpts.Run())
	test.VerifyOutputTemplate(t, restoreListFlexClusterTemplate, expected)
}
