// Copyright 2022 MongoDB Inc
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

package restores

import (
	"testing"
	"time"

	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/pointer"
	"github.com/stretchr/testify/require"
	atlasv2 "go.mongodb.org/atlas-sdk/v20250312009/admin"
	"go.uber.org/mock/gomock"
)

func TestWatchOpts_Run(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockStore := NewMockRestoreJobsDescriber(ctrl)

	expected := &atlasv2.DiskBackupSnapshotRestoreJob{
		Failed:     pointer.Get(false),
		FinishedAt: pointer.Get(time.Now()),
	}

	describeOpts := &WatchOpts{
		store:         mockStore,
		clusterName:   "Cluster0",
		id:            "1",
		isFlexCluster: false,
	}

	mockStore.
		EXPECT().
		RestoreJob(describeOpts.ProjectID, describeOpts.clusterName, describeOpts.id).
		Return(expected, nil).
		Times(1)

	if err := describeOpts.Run(); err != nil {
		t.Fatalf("Run() unexpected error: %v", err)
	}
}

func TestWatch_Run_FlexCluster(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockStore := NewMockRestoreJobsDescriber(ctrl)

	watchOpts := &WatchOpts{
		id:            "test",
		store:         mockStore,
		clusterName:   "cluster",
		isFlexCluster: true,
	}

	expected := &atlasv2.FlexBackupRestoreJob20241113{Status: pointer.Get("COMPLETED")}

	mockStore.
		EXPECT().
		RestoreFlexClusterJob(watchOpts.ConfigProjectID(), watchOpts.clusterName, watchOpts.id).
		Return(expected, nil).
		Times(1)

	require.NoError(t, watchOpts.Run())
}
