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
	"bytes"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/mocks"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/pointer"
	"github.com/stretchr/testify/assert"
	atlasv2 "go.mongodb.org/atlas-sdk/v20240530002/admin"
)

func TestCreateOpts_Run(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockStore := mocks.NewMockSnapshotsCreator(ctrl)

	t.Run("diskBackupReplicaSet run", func(t *testing.T) {
		expected := &atlasv2.DiskBackupSnapshot{
			Id: pointer.Get("DiskBackupReplicaSetId"),
		}
		buf := new(bytes.Buffer)

		createOpts := &CreateOpts{
			store:           mockStore,
			clusterName:     "",
			desc:            "",
			retentionInDays: 0,
			OutputOpts: cli.OutputOpts{
				Template:  createTemplate,
				OutWriter: buf,
			},
		}

		mockStore.
			EXPECT().
			CreateSnapshot(createOpts.ProjectID, createOpts.clusterName, createOpts.newCloudProviderSnapshot()).Return(expected, nil).
			Times(1)

		if err := createOpts.Run(); err != nil {
			t.Fatalf("Run() unexpected error: %v", err)
		}

		assert.Equal(t, `Snapshot 'DiskBackupReplicaSetId' created.
`, buf.String())
	})

	t.Run("diskBackupReplicaSet run", func(t *testing.T) {
		expected := &atlasv2.DiskBackupSnapshot{
			Id: pointer.Get("DiskBackupShardedClusterSnapshotId"),
		}

		buf := new(bytes.Buffer)

		createOpts := &CreateOpts{
			store:           mockStore,
			clusterName:     "",
			desc:            "",
			retentionInDays: 0,
			OutputOpts: cli.OutputOpts{
				Template:  createTemplate,
				OutWriter: buf,
			},
		}
		mockStore.
			EXPECT().
			CreateSnapshot(createOpts.ProjectID, createOpts.clusterName, createOpts.newCloudProviderSnapshot()).Return(expected, nil).
			Times(1)

		if err := createOpts.Run(); err != nil {
			t.Fatalf("Run() unexpected error: %v", err)
		}

		assert.Equal(t, `Snapshot 'DiskBackupShardedClusterSnapshotId' created.
`, buf.String())
	})
}
