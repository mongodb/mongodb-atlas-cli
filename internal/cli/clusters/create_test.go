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
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/flag"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/mocks"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/pointer"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/test"
	"github.com/spf13/afero"
	"github.com/stretchr/testify/assert"
	atlasv2 "go.mongodb.org/atlas-sdk/v20240530002/admin"
)

func TestCreateOpts_Run(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockStore := mocks.NewMockClusterCreator(ctrl)

	expected := &atlasv2.AdvancedClusterDescription{}

	t.Run("flags run", func(t *testing.T) {
		createOpts := &CreateOpts{
			name:       "ProjectBar",
			region:     "US",
			tier:       atlasM2,
			members:    3,
			diskSizeGB: 10,
			backup:     false,
			mdbVersion: "7.0",
			store:      mockStore,
		}

		cluster, _ := createOpts.newCluster()
		mockStore.
			EXPECT().
			CreateCluster(cluster).Return(expected, nil).
			Times(1)

		if err := createOpts.Run(); err != nil {
			t.Fatalf("Run() unexpected error: %v", err)
		}
	})

	t.Run("file run", func(t *testing.T) {
		appFS := afero.NewMemMapFs()
		// create test file
		fileYML := `
{
  "name": "ProjectBar",
  "diskSizeGB": 10,
  "numShards": 1,
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
		fileName := "atlas_cluster_create_test.json"
		_ = afero.WriteFile(appFS, fileName, []byte(fileYML), 0600)

		createOpts := &CreateOpts{
			filename: fileName,
			fs:       appFS,
			store:    mockStore,
		}

		cluster, _ := createOpts.newCluster()
		mockStore.
			EXPECT().
			CreateCluster(cluster).Return(expected, nil).
			Times(1)
		if err := createOpts.Run(); err != nil {
			t.Fatalf("Run() unexpected error: %v", err)
		}
	})
}

func TestCreateOpts_PostRun(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockStore := mocks.NewMockClusterCreator(ctrl)

	expected := &atlasv2.AdvancedClusterDescription{
		Name: pointer.Get("ProjectBar"),
	}

	buf := new(bytes.Buffer)

	createOpts := &CreateOpts{
		WatchOpts: cli.WatchOpts{
			EnableWatch: false,
			OutputOpts: cli.OutputOpts{
				Template:  createTemplate,
				OutWriter: buf,
			},
		},
		name:  "ProjectBar",
		store: mockStore,
	}

	cluster, _ := createOpts.newCluster()
	mockStore.
		EXPECT().
		CreateCluster(cluster).
		Return(expected, nil).
		Times(1)

	if err := createOpts.Run(); err != nil {
		t.Fatalf("Run() unexpected error: %v", err)
	}

	if err := createOpts.PostRun(); err != nil {
		t.Fatalf("PostRun() unexpected error: %v", err)
	}
	assert.Contains(t, `Cluster 'ProjectBar' is being created.
`, buf.String())
	t.Log(buf.String())
}

func TestCreateOpts_PostRun_EnableWatch(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockStore := &struct {
		*mocks.MockClusterCreator
		*mocks.MockClusterDescriber
	}{
		mocks.NewMockClusterCreator(ctrl),
		mocks.NewMockClusterDescriber(ctrl),
	}

	expected := &atlasv2.AdvancedClusterDescription{
		Name:      pointer.Get("ProjectBar"),
		StateName: pointer.Get("CREATING"),
	}
	expectedIdle := &atlasv2.AdvancedClusterDescription{
		Name:      expected.Name,
		StateName: pointer.Get("IDLE"),
	}

	buf := new(bytes.Buffer)

	createOpts := &CreateOpts{
		GlobalOpts: cli.GlobalOpts{
			ProjectID: "aaaa1e7e0f2912c554080abc",
		},
		WatchOpts: cli.WatchOpts{
			EnableWatch: true,
			OutputOpts: cli.OutputOpts{
				Template:  createTemplate,
				OutWriter: buf,
			},
		},
		name:  "ProjectBar",
		store: mockStore,
	}

	cluster, _ := createOpts.newCluster()
	mockStore.
		MockClusterCreator.
		EXPECT().
		CreateCluster(cluster).
		Return(expected, nil).
		Times(1)

	gomock.InOrder(
		mockStore.
			MockClusterDescriber.
			EXPECT().
			AtlasCluster(createOpts.ProjectID, expected.GetName()).
			Return(expected, nil).
			Times(1),
		mockStore.
			MockClusterDescriber.
			EXPECT().
			AtlasCluster(createOpts.ProjectID, expected.GetName()).
			Return(expectedIdle, nil).
			Times(1),
	)

	if err := createOpts.Run(); err != nil {
		t.Fatalf("Run() unexpected error: %v", err)
	}

	if err := createOpts.PostRun(); err != nil {
		t.Fatalf("PostRun() unexpected error: %v", err)
	}
	assert.Contains(t, `Cluster 'ProjectBar' created successfully.
`, buf.String())
	t.Log(buf.String())
}

func TestCreateBuilder(t *testing.T) {
	test.CmdValidator(
		t,
		CreateBuilder(),
		0,
		[]string{flag.Provider, flag.Region, flag.Members, flag.Tier, flag.DiskSizeGB, flag.MDBVersion, flag.Backup,
			flag.File, flag.TypeFlag, flag.Shards, flag.ProjectID, flag.Output, flag.EnableTerminationProtection, flag.Tag},
	)
}

func TestCreateTemplates(t *testing.T) {
	test.VerifyOutputTemplate(t, createTemplate, &atlasv2.AdvancedClusterDescription{})
	test.VerifyOutputTemplate(t, createWatchTemplate, &atlasv2.AdvancedClusterDescription{})
}
