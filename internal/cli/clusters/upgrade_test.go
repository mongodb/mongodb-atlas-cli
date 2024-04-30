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

//go:build unit

package clusters

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/flag"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/mocks"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/test"
	"github.com/spf13/afero"
	"go.mongodb.org/atlas/mongodbatlas"
)

const atlasM10 = "M10"

func TestUpgrade_Run(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockStore := mocks.NewMockAtlasSharedClusterGetterUpgrader(ctrl)

	expected := &mongodbatlas.Cluster{}

	t.Run("flags run", func(t *testing.T) {
		upgradeOpts := &UpgradeOpts{
			name:       "",
			tier:       atlasM10,
			diskSizeGB: 10,
			mdbVersion: "6.0",
			store:      mockStore,
		}

		mockStore.
			EXPECT().
			AtlasSharedCluster(upgradeOpts.ProjectID, upgradeOpts.name).
			Return(expected, nil).
			Times(1)

		upgradeOpts.patchOpts(expected)

		mockStore.
			EXPECT().
			UpgradeCluster(upgradeOpts.name, expected).
			Return(expected, nil).
			Times(1)

		if err := upgradeOpts.Run(); err != nil {
			t.Fatalf("Run() unexpected error: %v", err)
		}
	})

	t.Run("file run", func(t *testing.T) {
		appFS := afero.NewMemMapFs()
		// create test file
		fileYML := `
  {
    "autoScaling": {
            "autoIndexingEnabled": false,
            "compute": {
                    "enabled": false,
                    "scaleDownEnabled": false
            },
            "diskGBEnabled": false
    },
    "backupEnabled": false,
    "biConnector": {
            "enabled": false,
            "readPreference": "secondary"
    },
    "clusterType": "REPLICASET",
    "diskSizeGB": 10,
    "encryptionAtRestProvider": "NONE",
    "labels": [
            {
                    "key": "Infrastructure Tool",
                    "value": "mongoCLI"
            }
    ],
    "groupId": "62ab4d9e22f63b08ef5876f7",
    "mongoDBMajorVersion": "6.0",
    "name": "TestCluster",
    "numShards": 1,
    "paused": false,
    "pitEnabled": false,
    "providerBackupEnabled": false,
    "providerSettings": {
            "instanceSizeName": "M20",
            "providerName": "AWS",
            "regionName": "US_EAST_1"
    },
    "replicationFactor": 3,
    "replicationSpecs": [
            {
                    "id": "62bda0b86068de3e0c2cf036",
                    "numShards": 1,
                    "zoneName": "Zone 1",
                    "regionsConfig": {
                            "US_EAST_1": {
                                    "analyticsNodes": 0,
                                    "electableNodes": 3,
                                    "priority": 7,
                                    "readOnlyNodes": 0
                            }
                    }
            }
    ],
    "srvAddress": "mongodb+srv://cluster3.wb12jif.mongodb-dev.net",
    "links": [
            {
                    "rel": "self",
                    "href": "https://cloud-dev.mongodb.com/api/atlas/v1.0/groups/62ab4d9e22f63b08ef5876f7/clusters/Cluster3"
            },
            {
                    "rel": "http://cloud.mongodb.com/restoreJobs",
                    "href": "https://cloud-dev.mongodb.com/api/atlas/v1.0/groups/62ab4d9e22f63b08ef5876f7/clusters/Cluster3/restoreJobs"
            },
            {
                    "rel": "http://cloud.mongodb.com/snapshots",
                    "href": "https://cloud-dev.mongodb.com/api/atlas/v1.0/groups/62ab4d9e22f63b08ef5876f7/clusters/Cluster3/snapshots"
            }
    ],
    "versionReleaseSystem": "LTS"
}`
		fileName := "atlas_cluster_upgrade_test.json"
		_ = afero.WriteFile(appFS, fileName, []byte(fileYML), 0600)

		upgradeOpts := &UpgradeOpts{
			filename: fileName,
			fs:       appFS,
			store:    mockStore,
		}

		cluster, _ := upgradeOpts.cluster()
		mockStore.
			EXPECT().
			UpgradeCluster("", cluster).
			Return(expected, nil).
			Times(1)

		if err := upgradeOpts.Run(); err != nil {
			t.Fatalf("Run() unexpected error: %v", err)
		}
	})
}

func TestUpgradeBuilder(t *testing.T) {
	test.CmdValidator(
		t,
		UpgradeBuilder(),
		0,
		[]string{flag.Tier, flag.DiskSizeGB, flag.MDBVersion,
			flag.EnableTerminationProtection, flag.DisableTerminationProtection,
			flag.File, flag.Tag, flag.ProjectID, flag.Output},
	)
}
