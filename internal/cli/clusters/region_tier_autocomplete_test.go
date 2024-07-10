// Copyright 2023 MongoDB Inc
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

//go:build unit

package clusters

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/mocks"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/pointer"
	"github.com/stretchr/testify/require"
	atlasv2 "go.mongodb.org/atlas-sdk/v20240530002/admin"
)

func Test_autoCompleteOpts_tierSuggestions(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockStore := mocks.NewMockCloudProviderRegionsLister(ctrl)

	au := &autoCompleteOpts{

		store: mockStore,
	}
	expected := &atlasv2.PaginatedApiAtlasProviderRegions{
		Results: &[]atlasv2.CloudProviderRegions{
			{
				Provider: pointer.Get("AWS"),
				InstanceSizes: &[]atlasv2.ClusterCloudProviderInstanceSize{
					{
						Name:             pointer.Get("M0"),
						AvailableRegions: nil,
					},
				},
			},
		},
	}
	mockStore.
		EXPECT().
		CloudProviderRegions(au.ProjectID, au.tier, au.providers).
		Return(expected, nil).
		Times(1)

	res, err := au.tierSuggestions("")
	require.NoError(t, err)
	require.Equal(t, []string{"M0"}, res)
}

func Test_autoCompleteOpts_regionSuggestions(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockStore := mocks.NewMockCloudProviderRegionsLister(ctrl)

	au := &autoCompleteOpts{

		store: mockStore,
	}
	expected := &atlasv2.PaginatedApiAtlasProviderRegions{
		Results: &[]atlasv2.CloudProviderRegions{
			{
				Provider: pointer.Get("AWS"),
				InstanceSizes: &[]atlasv2.ClusterCloudProviderInstanceSize{
					{
						Name: pointer.Get("M0"),
						AvailableRegions: &[]atlasv2.AvailableCloudProviderRegion{
							{
								Name:    pointer.Get("EU_EAST"),
								Default: pointer.Get(false),
							},
						},
					},
				},
			},
		},
	}
	mockStore.
		EXPECT().
		CloudProviderRegions(au.ProjectID, au.tier, au.providers).
		Return(expected, nil).
		Times(1)

	res, err := au.regionSuggestions("")
	require.NoError(t, err)
	require.Equal(t, []string{"EU_EAST"}, res)
}
