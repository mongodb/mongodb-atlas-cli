// Copyright 2023 MongoDB Inc
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
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/flag"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/mocks"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/test"
	atlasv2 "go.mongodb.org/atlas-sdk/v20240530002/admin"
)

func TestCreate_Run(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockStore := mocks.NewMockServerlessRestoreJobsCreator(ctrl)

	expected := &atlasv2.ServerlessBackupRestoreJob{}

	t.Run(automatedRestore, func(t *testing.T) {
		listOpts := &CreateOpts{
			store:             mockStore,
			deliveryType:      automatedRestore,
			clusterName:       "Cluster0",
			targetClusterName: "Cluster1",
			targetProjectID:   "1",
		}

		mockStore.
			EXPECT().
			ServerlessCreateRestoreJobs(listOpts.ProjectID, "Cluster0", listOpts.newCloudProviderSnapshotRestoreJob()).
			Return(expected, nil).
			Times(1)

		if err := listOpts.Run(); err != nil {
			t.Fatalf("Run() unexpected error: %v", err)
		}
	})

	t.Run(pointInTimeRestore, func(t *testing.T) {
		listOpts := &CreateOpts{
			store:             mockStore,
			deliveryType:      pointInTimeRestore,
			clusterName:       "Cluster0",
			targetClusterName: "Cluster1",
			targetProjectID:   "1",
		}

		mockStore.
			EXPECT().
			ServerlessCreateRestoreJobs(listOpts.ProjectID, "Cluster0", listOpts.newCloudProviderSnapshotRestoreJob()).
			Return(expected, nil).
			Times(1)

		if err := listOpts.Run(); err != nil {
			t.Fatalf("Run() unexpected error: %v", err)
		}
	})

	t.Run(downloadRestore, func(t *testing.T) {
		listOpts := &CreateOpts{
			store:        mockStore,
			deliveryType: downloadRestore,
			clusterName:  "Cluster0",
		}

		mockStore.
			EXPECT().
			ServerlessCreateRestoreJobs(listOpts.ProjectID, "Cluster0", listOpts.newCloudProviderSnapshotRestoreJob()).
			Return(expected, nil).
			Times(1)

		if err := listOpts.Run(); err != nil {
			t.Fatalf("Run() unexpected error: %v", err)
		}
	})
}

func TestCreateBuilder(t *testing.T) {
	test.CmdValidator(
		t,
		CreateBuilder(),
		0,
		[]string{
			flag.DeliveryType,
			flag.SnapshotID,
			flag.ClusterName,
			flag.TargetProjectID,
			flag.TargetClusterName,
			flag.OplogTS,
			flag.OplogInc,
			flag.ProjectID,
			flag.Output,
		},
	)
}
