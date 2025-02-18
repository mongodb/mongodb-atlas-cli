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

package snapshots

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/mocks"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/pointer"
	"github.com/stretchr/testify/require"
	atlasv2 "go.mongodb.org/atlas-sdk/v20241113005/admin"
)

func TestWatch_Run(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockStore := mocks.NewMockSnapshotsDescriber(ctrl)

	watchOpts := &WatchOpts{
		id:            "test",
		store:         mockStore,
		clusterName:   "cluster",
		isFlexCluster: false,
	}

	expected := &atlasv2.DiskBackupReplicaSet{Status: pointer.Get("completed")}

	mockStore.
		EXPECT().
		Snapshot(watchOpts.ProjectID, watchOpts.clusterName, watchOpts.id).
		Return(expected, nil).
		Times(1)

	require.NoError(t, watchOpts.Run())
}

func TestWatch_Run_FlexCluster(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockStore := mocks.NewMockSnapshotsDescriber(ctrl)

	watchOpts := &WatchOpts{
		id:            "test",
		store:         mockStore,
		clusterName:   "cluster",
		isFlexCluster: true,
	}

	expected := &atlasv2.FlexBackupSnapshot20241113{Status: pointer.Get("COMPLETED")}

	mockStore.
		EXPECT().
		FlexClusterSnapshot(watchOpts.ConfigProjectID(), watchOpts.clusterName, watchOpts.id).
		Return(expected, nil).
		Times(1)

	require.NoError(t, watchOpts.Run())
}
