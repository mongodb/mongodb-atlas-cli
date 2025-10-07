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

	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/pointer"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/test"
	"github.com/stretchr/testify/require"
	atlasClustersPinned "go.mongodb.org/atlas-sdk/v20240530005/admin"
	atlasv2 "go.mongodb.org/atlas-sdk/v20250312008/admin"
	"go.uber.org/mock/gomock"
)

func TestList_RunDedicatedCluster(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockStore := NewMockClusterLister(ctrl)

	expected := &atlasClustersPinned.PaginatedAdvancedClusterDescription{
		Results: &[]atlasClustersPinned.AdvancedClusterDescription{
			{
				Name: pointer.Get("test"),
				Id:   pointer.Get("123"),
			},
		},
	}

	listOpts := &ListOpts{
		store: mockStore,
	}

	mockStore.
		EXPECT().
		ProjectClusters(listOpts.ProjectID, listOpts.NewAtlasListOptions()).
		Return(expected, nil).
		Times(1)

	require.NoError(t, listOpts.Run())
}

func TestList_RunFlexCluster(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockStore := NewMockClusterLister(ctrl)

	expected := &atlasv2.PaginatedFlexClusters20241113{
		Results: &[]atlasv2.FlexClusterDescription20241113{
			{
				Name: pointer.Get("test"),
				Id:   pointer.Get("123"),
			},
		},
	}

	listOpts := &ListOpts{
		store: mockStore,
		tier:  atlasFlex,
	}

	mockStore.
		EXPECT().
		ListFlexClusters(listOpts.newListFlexClustersAPIParams()).
		Return(expected, nil).
		Times(1)

	require.NoError(t, listOpts.Run())
}

func TestList_RunDedicatedCluster_IndependentShardScaling(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockStore := NewMockClusterLister(ctrl)

	expected := &atlasv2.PaginatedClusterDescription20240805{
		Results: &[]atlasv2.ClusterDescription20240805{
			{
				Name: pointer.Get("test"),
				Id:   pointer.Get("123"),
			},
			{
				Name: pointer.Get("nonISS"),
				Id:   pointer.Get("124"),
			},
		},
	}

	listOpts := &ListOpts{
		store:           mockStore,
		autoScalingMode: independentShardScalingFlag,
	}

	mockStore.
		EXPECT().
		GetClusterAutoScalingConfig(listOpts.ProjectID, *expected.GetResults()[0].Name).
		Return(&atlasv2.ClusterDescriptionAutoScalingModeConfiguration{
			AutoScalingMode: pointer.Get(independentShardScalingFlag),
		}, nil).
		Times(1)

	mockStore.
		EXPECT().
		GetClusterAutoScalingConfig(listOpts.ProjectID, *expected.GetResults()[1].Name).
		Return(&atlasv2.ClusterDescriptionAutoScalingModeConfiguration{
			AutoScalingMode: pointer.Get(clusterWideScalingFlag),
		}, nil).
		Times(1)

	mockStore.
		EXPECT().
		LatestProjectClusters(listOpts.ProjectID, listOpts.NewAtlasListOptions()).
		Return(expected, nil).
		Times(1)

	require.NoError(t, listOpts.Run())
}

func TestListTemplate(t *testing.T) {
	test.VerifyOutputTemplate(t, listTemplate, atlasClustersPinned.PaginatedAdvancedClusterDescription{})
}
