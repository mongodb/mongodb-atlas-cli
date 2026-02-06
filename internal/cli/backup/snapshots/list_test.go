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

package snapshots

import (
	"bytes"
	"testing"

	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/pointer"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/test"
	"github.com/stretchr/testify/require"
	atlasv2 "go.mongodb.org/atlas-sdk/v20250312013/admin"
	"go.uber.org/mock/gomock"
)

func TestList_Run(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockStore := NewMockLister(ctrl)
	buf := new(bytes.Buffer)
	expected := &atlasv2.PaginatedCloudBackupReplicaSet{
		Results: &[]atlasv2.DiskBackupReplicaSet{
			{
				CloudProvider: pointer.Get("AWS"),
				Id:            pointer.Get("5f9b0b5e0b5e9d6b6e0b5e9d"),
				SnapshotType:  pointer.Get("cloud"),
			},
			*atlasv2.NewDiskBackupReplicaSet(),
		},
	}

	listOpts := &ListOpts{
		store:       mockStore,
		clusterName: "Cluster0",
		OutputOpts: cli.OutputOpts{
			Template:  listTemplate,
			OutWriter: buf,
		},
	}

	expectedError := &atlasv2.GenericOpenAPIError{}
	expectedError.SetModel(atlasv2.ApiError{ErrorCode: CannotUseNotFlexWithFlexApisErrorCode})

	mockStore.
		EXPECT().
		FlexClusterSnapshots(listOpts.newListFlexBackupsAPIParams()).
		Return(nil, expectedError).
		Times(1)

	mockStore.
		EXPECT().
		Snapshots(listOpts.ProjectID, "Cluster0", listOpts.NewAtlasListOptions()).
		Return(expected, nil).
		Times(1)

	require.NoError(t, listOpts.Run())
	test.VerifyOutputTemplate(t, listTemplate, expected)
}

func TestList_Run_FlexCluster(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockStore := NewMockLister(ctrl)
	expected := &atlasv2.PaginatedApiAtlasFlexBackupSnapshot20241113{}
	buf := new(bytes.Buffer)
	listOpts := &ListOpts{
		store:       mockStore,
		clusterName: "Cluster0",
		OutputOpts: cli.OutputOpts{
			Template:  listTemplate,
			OutWriter: buf,
		},
	}

	mockStore.
		EXPECT().
		FlexClusterSnapshots(listOpts.newListFlexBackupsAPIParams()).
		Return(expected, nil).
		Times(1)

	require.NoError(t, listOpts.Run())
	test.VerifyOutputTemplate(t, listTemplate, expected)
}

func TestList_Run_ClusterNotFoundFallback(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockStore := NewMockLister(ctrl)
	buf := new(bytes.Buffer)
	expected := &atlasv2.PaginatedCloudBackupReplicaSet{
		Results: &[]atlasv2.DiskBackupReplicaSet{
			{
				CloudProvider: pointer.Get("AWS"),
				Id:            pointer.Get("5f9b0b5e0b5e9d6b6e0b5e9d"),
				SnapshotType:  pointer.Get("cloud"),
			},
			*atlasv2.NewDiskBackupReplicaSet(),
		},
	}

	listOpts := &ListOpts{
		store:       mockStore,
		clusterName: "Cluster0",
		OutputOpts: cli.OutputOpts{
			Template:  listTemplate,
			OutWriter: buf,
		},
	}

	expectedError := &atlasv2.GenericOpenAPIError{}
	expectedError.SetModel(atlasv2.ApiError{ErrorCode: ClusterNotFoundErrorCode})

	mockStore.
		EXPECT().
		FlexClusterSnapshots(listOpts.newListFlexBackupsAPIParams()).
		Return(nil, expectedError).
		Times(1)

	mockStore.
		EXPECT().
		Snapshots(listOpts.ProjectID, "Cluster0", listOpts.NewAtlasListOptions()).
		Return(expected, nil).
		Times(1)

	require.NoError(t, listOpts.Run())
	test.VerifyOutputTemplate(t, listTemplate, expected)
}
