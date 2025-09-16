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

package clusters

import (
	"testing"

	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli"
	"github.com/stretchr/testify/require"
	atlasv2 "go.mongodb.org/atlas-sdk/v20250312007/admin"
	"go.uber.org/mock/gomock"
)

func TestDelete_Run(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockStore := NewMockClusterDeleter(ctrl)

	deleteOpts := &DeleteOpts{
		DeleteOpts: &cli.DeleteOpts{
			Confirm: true,
			Entry:   "test",
		},
		store: mockStore,
	}

	mockStore.
		EXPECT().
		DeleteCluster(deleteOpts.ProjectID, deleteOpts.Entry).
		Return(nil).
		Times(1)

	require.NoError(t, deleteOpts.Run())
}

func TestDelete_RunFlexCluster(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockStore := NewMockClusterDeleter(ctrl)

	expectedError := &atlasv2.GenericOpenAPIError{}
	expectedError.SetModel(atlasv2.ApiError{ErrorCode: cannotUseFlexWithClusterApisErrorCode})

	deleteOpts := &DeleteOpts{
		DeleteOpts: &cli.DeleteOpts{
			Confirm: true,
			Entry:   "test",
		},
		store: mockStore,
	}

	mockStore.
		EXPECT().
		DeleteCluster(deleteOpts.ProjectID, deleteOpts.Entry).
		Return(expectedError).
		Times(1)

	mockStore.
		EXPECT().
		DeleteFlexCluster(deleteOpts.ProjectID, deleteOpts.Entry).
		Return(nil).
		Times(1)

	require.NoError(t, deleteOpts.Run())
}
