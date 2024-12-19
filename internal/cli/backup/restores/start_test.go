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
	"github.com/stretchr/testify/require"
	atlasv2 "go.mongodb.org/atlas-sdk/v20241113004/admin"
)

func TestStart_Run(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockStore := mocks.NewMockRestoreJobsCreator(ctrl)

	expected := &atlasv2.DiskBackupSnapshotRestoreJob{}

	t.Run(automatedRestore, func(t *testing.T) {
		listOpts := &StartOpts{
			store:             mockStore,
			method:            automatedRestore,
			clusterName:       "Cluster0",
			targetClusterName: "Cluster1",
			targetProjectID:   "1",
		}

		expectedError := &atlasv2.GenericOpenAPIError{}
		expectedError.SetModel(atlasv2.ApiError{ErrorCode: cannotUseNotFlexWithFlexApisErrorCode})

		mockStore.
			EXPECT().
			CreateRestoreFlexClusterJobs(listOpts.ProjectID, "Cluster0", listOpts.newFlexBackupRestoreJobCreate()).
			Return(nil, expectedError).
			Times(1)

		mockStore.
			EXPECT().
			CreateRestoreJobs(listOpts.ProjectID, "Cluster0", listOpts.newCloudProviderSnapshotRestoreJob()).
			Return(expected, nil).
			Times(1)

		require.NoError(t, listOpts.Run())
	})

	t.Run("Flex Cluster automated restore job", func(t *testing.T) {
		listOpts := &StartOpts{
			store:             mockStore,
			method:            automatedRestore,
			clusterName:       "Cluster0",
			targetClusterName: "Cluster1",
			targetProjectID:   "1",
		}

		expectedFlex := &atlasv2.FlexBackupRestoreJob20241113{}
		expectedError := &atlasv2.GenericOpenAPIError{}
		expectedError.SetModel(atlasv2.ApiError{ErrorCode: cannotUseNotFlexWithFlexApisErrorCode})

		mockStore.
			EXPECT().
			CreateRestoreFlexClusterJobs(listOpts.ProjectID, "Cluster0", listOpts.newFlexBackupRestoreJobCreate()).
			Return(expectedFlex, nil).
			Times(1)

		require.NoError(t, listOpts.Run())
	})

	t.Run(pointInTimeRestore, func(t *testing.T) {
		listOpts := &StartOpts{
			store:             mockStore,
			method:            pointInTimeRestore,
			clusterName:       "Cluster0",
			targetClusterName: "Cluster1",
			targetProjectID:   "1",
		}

		expectedError := &atlasv2.GenericOpenAPIError{}
		expectedError.SetModel(atlasv2.ApiError{ErrorCode: cannotUseNotFlexWithFlexApisErrorCode})

		mockStore.
			EXPECT().
			CreateRestoreFlexClusterJobs(listOpts.ProjectID, "Cluster0", listOpts.newFlexBackupRestoreJobCreate()).
			Return(nil, expectedError).
			Times(1)

		mockStore.
			EXPECT().
			CreateRestoreJobs(listOpts.ProjectID, "Cluster0", listOpts.newCloudProviderSnapshotRestoreJob()).
			Return(expected, nil).
			Times(1)

		require.NoError(t, listOpts.Run())
	})

	t.Run(downloadRestore, func(t *testing.T) {
		listOpts := &StartOpts{
			store:       mockStore,
			method:      downloadRestore,
			clusterName: "Cluster0",
		}

		expectedError := &atlasv2.GenericOpenAPIError{}
		expectedError.SetModel(atlasv2.ApiError{ErrorCode: cannotUseNotFlexWithFlexApisErrorCode})

		mockStore.
			EXPECT().
			CreateRestoreFlexClusterJobs(listOpts.ProjectID, "Cluster0", listOpts.newFlexBackupRestoreJobCreate()).
			Return(nil, expectedError).
			Times(1)

		mockStore.
			EXPECT().
			CreateRestoreJobs(listOpts.ProjectID, "Cluster0", listOpts.newCloudProviderSnapshotRestoreJob()).
			Return(expected, nil).
			Times(1)

		require.NoError(t, listOpts.Run())
	})
}
