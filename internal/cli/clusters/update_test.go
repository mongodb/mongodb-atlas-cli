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

package clusters

import (
	"bytes"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/mocks"
	"github.com/spf13/afero"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	atlasClustersPinned "go.mongodb.org/atlas-sdk/v20240530005/admin"
	atlasv2 "go.mongodb.org/atlas-sdk/v20241113005/admin"
)

func TestUpdate_Run(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockStore := mocks.NewMockAtlasClusterGetterUpdater(ctrl)

	expected := &atlasClustersPinned.AdvancedClusterDescription{}

	t.Run("flags run", func(t *testing.T) {
		updateOpts := &UpdateOpts{
			name:       "ProjectBar",
			tier:       atlasM2,
			diskSizeGB: 10,
			mdbVersion: "7.0",
			store:      mockStore,
		}

		mockStore.
			EXPECT().
			AtlasCluster(updateOpts.ProjectID, updateOpts.name).
			Return(expected, nil).
			Times(1)

		updateOpts.patchOpts(expected)

		mockStore.
			EXPECT().
			UpdateCluster(updateOpts.ConfigProjectID(), updateOpts.name, expected).Return(expected, nil).
			Times(1)

		if err := updateOpts.Run(); err != nil {
			t.Fatalf("Run() unexpected error: %v", err)
		}
	})

	t.Run("file run", func(t *testing.T) {
		appFS := afero.NewMemMapFs()
		// create test file
		fileYML := `{
  "name": "ProjectBar",
  "diskSizeGB": 10,
  "numShards": 1,
  "connectionStrings": {
    "standard": "mongodb://clusterm10-shard-00-00.85xn1.mongodb.net:27017,clusterm10-shard-00-01.85xn1.mongodb.net:27017,clusterm10-shard-00-02.85xn1.mongodb.net:27017/?ssl=true\u0026authSource=admin\u0026replicaSet=atlas-zzw0ln-shard-0",
    "standardSrv": "mongodb+srv://clusterm10.85xn1.mongodb.net"
  },
  "providerSettings": {
    "providerName": "AWS",
    "instanceSizeName": "M2",
    "regionName": "US"
  },
  "clusterType" : "REPLICASET",
  "replicationFactor": 3,
  "replicationSpecs": [{
    "numShards": 1,
    "regionsConfig": {
      "US_EAST_1": {
        "analyticsNodes": 0,
        "electableNodes": 3,
        "priority": 7,
        "readOnlyNodes": 0
      }
    },
    "zoneName": "Zone 1"
  }],
  "backupEnabled": false,
  "providerBackupEnabled" : false
}`
		fileName := "atlas_cluster_update_test.json"
		_ = afero.WriteFile(appFS, fileName, []byte(fileYML), 0600)

		buf := new(bytes.Buffer)
		updateOpts := &UpdateOpts{
			filename: fileName,
			fs:       appFS,
			store:    mockStore,
			name:     "ProjectBar",
			OutputOpts: cli.OutputOpts{
				Template:  updateTmpl,
				OutWriter: buf,
			},
		}

		cluster, _ := updateOpts.cluster()
		removeReadOnlyAttributes(cluster)
		mockStore.
			EXPECT().
			UpdateCluster(updateOpts.ConfigProjectID(), updateOpts.name, cluster).
			Return(expected, nil).
			Times(1)

		require.NoError(t, updateOpts.Run())
		assert.Contains(t, buf.String(), "Updating cluster")
	})
}

func TestUpdate_FlexClusterRun(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockStore := mocks.NewMockAtlasClusterGetterUpdater(ctrl)

	expected := &atlasv2.FlexClusterDescription20241113{}

	t.Run("flags run", func(t *testing.T) {
		updateOpts := &UpdateOpts{
			name:          "ProjectBar",
			store:         mockStore,
			isFlexCluster: true,
			ProjectOpts:   cli.ProjectOpts{ProjectID: "test"},
		}

		mockStore.
			EXPECT().
			FlexCluster(updateOpts.ConfigProjectID(), updateOpts.name).
			Return(expected, nil).
			Times(1)

		cluster, _ := updateOpts.newFlexCluster()

		mockStore.
			EXPECT().
			UpdateFlexCluster(updateOpts.ConfigProjectID(), updateOpts.name,
				cluster).
			Return(expected, nil).
			Times(1)

		mockStore.
			EXPECT().
			FlexCluster(updateOpts.ConfigProjectID(), updateOpts.name).
			Return(expected, nil).
			Times(1)

		require.NoError(t, updateOpts.Run())
	})

	t.Run("flags run with existing tags", func(t *testing.T) {
		updateOpts := &UpdateOpts{
			name:          "ProjectBar",
			store:         mockStore,
			isFlexCluster: true,
			ProjectOpts:   cli.ProjectOpts{ProjectID: "test"},
			tag: map[string]string{
				"key1": "value22",
			},
		}

		expectedGet := &atlasv2.FlexClusterDescription20241113{
			Tags: newResourceTags(map[string]string{
				"test1": "value1",
			}),
		}

		expectedPost := &atlasv2.FlexClusterDescription20241113{
			Tags: newResourceTags(map[string]string{
				"test1": "value1",
				"key1":  "value22",
			}),
		}

		mockStore.
			EXPECT().
			FlexCluster(updateOpts.ConfigProjectID(), updateOpts.name).
			Return(expectedGet, nil).
			Times(1)

		cluster, _ := updateOpts.newFlexCluster()

		mockStore.
			EXPECT().
			FlexCluster(updateOpts.ConfigProjectID(), updateOpts.name).
			Return(expectedGet, nil).
			Times(1)

		mockStore.
			EXPECT().
			UpdateFlexCluster(updateOpts.ConfigProjectID(), updateOpts.name,
				cluster).
			Return(expectedPost, nil).
			Times(1)

		require.NoError(t, updateOpts.Run())
	})

	t.Run("file run", func(t *testing.T) {
		appFS := afero.NewMemMapFs()
		// create test file
		fileYML := `{
		  "tags": [
			{
			  "key": "test",
			  "value": "test222"
			}
		  ],
		  "terminationProtectionEnabled": false
}`
		fileName := "atlas_flex_cluster_update_test.json"

		_ = afero.WriteFile(appFS, fileName, []byte(fileYML), 0600)

		buf := new(bytes.Buffer)
		updateOpts := &UpdateOpts{
			filename: fileName,
			fs:       appFS,
			store:    mockStore,
			name:     "ProjectBar",
			OutputOpts: cli.OutputOpts{
				Template:  updateTmpl,
				OutWriter: buf,
			},
			isFlexCluster:               true,
			enableTerminationProtection: true,
			ProjectOpts:                 cli.ProjectOpts{ProjectID: "test"},
			tag: map[string]string{
				"test": "test222",
			},
		}

		cluster, _ := updateOpts.newFlexCluster()
		mockStore.
			EXPECT().
			UpdateFlexCluster(
				updateOpts.ConfigProjectID(),
				updateOpts.name, cluster).
			Return(expected, nil).
			Times(1)

		require.NoError(t, updateOpts.Run())
		assert.Contains(t, buf.String(), "Updating cluster")
	})
}
