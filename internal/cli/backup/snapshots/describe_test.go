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
	"testing"

	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/test"
	"github.com/stretchr/testify/require"
	atlasv2 "go.mongodb.org/atlas-sdk/v20250312008/admin"
	"go.uber.org/mock/gomock"
)

func TestDescribe_Run(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockStore := NewMockDescriber(ctrl)

	var expected atlasv2.DiskBackupReplicaSet

	describeOpts := &DescribeOpts{
		store: mockStore,
	}

	expectedError := &atlasv2.GenericOpenAPIError{}
	expectedError.SetModel(atlasv2.ApiError{ErrorCode: CannotUseNotFlexWithFlexApisErrorCode})

	mockStore.
		EXPECT().
		FlexClusterSnapshot(describeOpts.ProjectID, describeOpts.clusterName, describeOpts.snapshot).
		Return(nil, expectedError).
		Times(1)

	mockStore.
		EXPECT().
		Snapshot(describeOpts.ProjectID, describeOpts.clusterName, describeOpts.snapshot).
		Return(&expected, nil).
		Times(1)

	require.NoError(t, describeOpts.Run())
	test.VerifyOutputTemplate(t, describeTemplate, expected)
}

func TestDescribe_Run_FlexCluster(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockStore := NewMockDescriber(ctrl)
	expected := &atlasv2.FlexBackupSnapshot20241113{}

	describeOpts := &DescribeOpts{
		store: mockStore,
	}

	mockStore.
		EXPECT().
		FlexClusterSnapshot(describeOpts.ProjectID, describeOpts.clusterName, describeOpts.snapshot).
		Return(expected, nil).
		Times(1)

	require.NoError(t, describeOpts.Run())
	test.VerifyOutputTemplate(t, describeTemplateFlex, expected)
}
