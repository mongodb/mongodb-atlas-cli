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

package store

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
	atlasClustersPinned "go.mongodb.org/atlas-sdk/v20240530005/admin"
	"go.uber.org/mock/gomock"
)

// Test CreateClusterPerType will use the right method based on the type of the cluster
func TestCreateClusterPerType(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	actualStore, err := New(WithContext(context.Background()))
	require.NoError(t, err)

	// add mocked atlascluster create stores
	actualStore.clientClusters = mocks.NewMockClustersApiService(ctrl)

	actualStore.clientClusters.EXPECT().CreateCluster(gomock.Any()).Return(nil, ErrUnsupportedClusterType)

	cluster := &atlasClustersPinned.AdvancedClusterDescription{}
	mockStore.EXPECT().CreateCluster(cluster).Return(cluster, nil)

}
