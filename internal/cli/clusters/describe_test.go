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
	"errors"
	"testing"

	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/pointer"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/test"
	"github.com/stretchr/testify/require"
	atlasClustersPinned "go.mongodb.org/atlas-sdk/v20240530005/admin"
	atlasv2 "go.mongodb.org/atlas-sdk/v20250312012/admin"
	"go.uber.org/mock/gomock"
)

func TestDescribe_Run(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockStore := NewMockClusterDescriber(ctrl)

	expected := &atlasClustersPinned.AdvancedClusterDescription{}

	describeOpts := &DescribeOpts{
		name:  "test",
		store: mockStore,
	}

	mockStore.
		EXPECT().
		AtlasCluster(describeOpts.ProjectID, describeOpts.name).
		Return(expected, nil).
		Times(1)

	require.NoError(t, describeOpts.Run())
	test.VerifyOutputTemplate(t, describeTemplate, expected)
}

func TestDescribe_RunFlexCluster(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockStore := NewMockClusterDescriber(ctrl)

	expected := &atlasv2.FlexClusterDescription20241113{}
	expectedError := &atlasClustersPinned.GenericOpenAPIError{}
	expectedError.SetModel(atlasClustersPinned.ApiError{ErrorCode: pointer.Get(cannotUseFlexWithClusterApisErrorCode)})

	describeOpts := &DescribeOpts{
		name:  "test",
		store: mockStore,
	}

	mockStore.
		EXPECT().
		AtlasCluster(describeOpts.ProjectID, describeOpts.name).
		Return(nil, expectedError).
		Times(1)

	mockStore.
		EXPECT().
		FlexCluster(describeOpts.ProjectID, describeOpts.name).
		Return(expected, nil).
		Times(1)

	require.NoError(t, describeOpts.Run())
	test.VerifyOutputTemplate(t, describeTemplate, expected)
}

func TestDescribe_RunFlexCluster_Error(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockStore := NewMockClusterDescriber(ctrl)

	expected := &atlasv2.FlexClusterDescription20241113{}
	expectedError := errors.New("test")

	describeOpts := &DescribeOpts{
		name:  "test",
		store: mockStore,
	}

	mockStore.
		EXPECT().
		AtlasCluster(describeOpts.ProjectID, describeOpts.name).
		Return(nil, expectedError).
		Times(1)

	mockStore.
		EXPECT().
		FlexCluster(describeOpts.ProjectID, describeOpts.name).
		Return(expected, nil).
		Times(0)

	require.Error(t, describeOpts.Run())
}

func TestDescribe_RunDedicatedCluster_IndependentShardScaling(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockStore := NewMockClusterDescriber(ctrl)

	expected := &atlasv2.ClusterDescription20240805{}

	describeOpts := &DescribeOpts{
		name:            "test",
		store:           mockStore,
		autoScalingMode: independentShardScalingFlag,
	}

	mockStore.
		EXPECT().
		LatestAtlasCluster(describeOpts.ProjectID, describeOpts.name).
		Return(expected, nil).
		Times(1)

	require.NoError(t, describeOpts.Run())
	test.VerifyOutputTemplate(t, describeTemplate, expected)
}
