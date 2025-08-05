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

package setup

import (
	"bytes"
	"testing"

	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/config"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/flag"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/mocks"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/pointer"
	"github.com/spf13/pflag"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	atlasClustersPinned "go.mongodb.org/atlas-sdk/v20240530005/admin"
	atlasv2 "go.mongodb.org/atlas-sdk/v20250312005/admin"
	"go.uber.org/mock/gomock"
)

func Test_setupOpts_PreRunWithAPIKeys(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockFlow := mocks.NewMockRefresher(ctrl)
	ctx := t.Context()
	buf := new(bytes.Buffer)

	opts := &Opts{}

	opts.OutWriter = buf
	opts.register.WithFlow(mockFlow)

	config.SetPublicAPIKey("publicKey")
	config.SetPrivateAPIKey("privateKey")

	require.NoError(t, opts.PreRun(ctx))

	assert.True(t, opts.skipRegister)
	assert.Equal(t, 0, buf.Len())
	assert.True(t, opts.skipLogin)
}

func Test_setupOpts_RunSkipRegister(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockFlow := mocks.NewMockRefresher(ctrl)
	ctx := t.Context()
	buf := new(bytes.Buffer)

	opts := &Opts{
		skipLogin: true,
	}
	opts.register.WithFlow(mockFlow)

	config.SetAccessToken("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiYWRtaW4iOnRydWV9.TJVA95OrM7E2cBab30RMHrHDcEfxjoYZgeFONFh7HgQ")

	opts.OutWriter = buf
	require.NoError(t, opts.PreRun(ctx))
	assert.True(t, opts.skipRegister)
}

func TestCluster_Run(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockStore := NewMockAtlasClusterQuickStarter(ctrl)
	mockFlow := mocks.NewMockRefresher(ctrl)

	expectedCluster := &atlasClustersPinned.AdvancedClusterDescription{
		StateName: pointer.Get("IDLE"),
		ConnectionStrings: &atlasClustersPinned.ClusterConnectionStrings{
			StandardSrv: pointer.Get(""),
		},
	}

	expectedClusterLatest := &atlasv2.ClusterDescription20240805{
		StateName: pointer.Get(expectedCluster.GetStateName()),
		ConnectionStrings: &atlasv2.ClusterConnectionStrings{
			StandardSrv: pointer.Get(expectedCluster.ConnectionStrings.GetStandardSrv()),
		},
	}

	expectedDBUser := &atlasv2.CloudDatabaseUser{}

	var expectedProjectAccessLists *atlasv2.PaginatedNetworkAccess

	opts := &Opts{
		ClusterName:     "ProjectBar",
		Region:          "US",
		store:           mockStore,
		IPAddresses:     []string{"0.0.0.0"},
		DBUsername:      "user",
		DBUserPassword:  "test",
		Provider:        "AWS",
		SkipMongosh:     true,
		SkipSampleData:  true,
		Confirm:         true,
		MDBVersion:      "7.0",
		Tag:             map[string]string{"env": "test"},
		AutoScalingMode: clusterWideScaling,
	}
	opts.register.WithFlow(mockFlow)

	projectIPAccessList := opts.newProjectIPAccessList()

	mockStore.
		EXPECT().
		CreateCluster(opts.newCluster()).Return(expectedCluster, nil).
		Times(1)

	mockStore.
		EXPECT().
		CreateProjectIPAccessList(projectIPAccessList).Return(expectedProjectAccessLists, nil).
		Times(1)

	mockStore.
		EXPECT().
		LatestAtlasCluster(opts.ConfigProjectID(), opts.ClusterName).Return(expectedClusterLatest, nil).
		Times(2)

	mockStore.
		EXPECT().
		CreateDatabaseUser(opts.newDatabaseUser()).Return(expectedDBUser, nil).
		Times(1)

	if err := opts.setupCluster(); err != nil {
		t.Fatalf("Run() unexpected error: %v", err)
	}
}

func TestCluster_Run_LatestAPI(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockStore := NewMockAtlasClusterQuickStarter(ctrl)
	mockFlow := mocks.NewMockRefresher(ctrl)

	expectedCluster := &atlasv2.ClusterDescription20240805{
		StateName: pointer.Get("IDLE"),
		ConnectionStrings: &atlasv2.ClusterConnectionStrings{
			StandardSrv: pointer.Get(""),
		},
	}

	expectedDBUser := &atlasv2.CloudDatabaseUser{}

	var expectedProjectAccessLists *atlasv2.PaginatedNetworkAccess

	opts := &Opts{
		ClusterName:     "ProjectBar",
		Region:          "US",
		store:           mockStore,
		IPAddresses:     []string{"0.0.0.0"},
		DBUsername:      "user",
		DBUserPassword:  "test",
		Provider:        "AWS",
		SkipMongosh:     true,
		SkipSampleData:  true,
		Confirm:         true,
		MDBVersion:      "7.0",
		Tag:             map[string]string{"env": "test"},
		AutoScalingMode: "independentShardingScaling",
	}
	opts.register.WithFlow(mockFlow)

	projectIPAccessList := opts.newProjectIPAccessList()

	mockStore.
		EXPECT().
		CreateClusterLatest(opts.newClusterLatest()).Return(expectedCluster, nil).
		Times(1)

	mockStore.
		EXPECT().
		CreateProjectIPAccessList(projectIPAccessList).Return(expectedProjectAccessLists, nil).
		Times(1)

	mockStore.
		EXPECT().
		LatestAtlasCluster(opts.ConfigProjectID(), opts.ClusterName).Return(expectedCluster, nil).
		Times(2)

	mockStore.
		EXPECT().
		CreateDatabaseUser(opts.newDatabaseUser()).Return(expectedDBUser, nil).
		Times(1)

	if err := opts.setupCluster(); err != nil {
		t.Fatalf("Run() unexpected error: %v", err)
	}
}

func TestCluster_Run_CheckFlagsSet(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockStore := NewMockAtlasClusterQuickStarter(ctrl)
	mockFlow := mocks.NewMockRefresher(ctrl)
	defer ctrl.Finish()

	expectedCluster := &atlasClustersPinned.AdvancedClusterDescription{
		StateName: pointer.Get("IDLE"),
		ConnectionStrings: &atlasClustersPinned.ClusterConnectionStrings{
			StandardSrv: pointer.Get(""),
		},
	}

	expectedClusterLatest := &atlasv2.ClusterDescription20240805{
		StateName: pointer.Get(expectedCluster.GetStateName()),
		ConnectionStrings: &atlasv2.ClusterConnectionStrings{
			StandardSrv: pointer.Get(expectedCluster.ConnectionStrings.GetStandardSrv()),
		},
	}

	expectedDBUser := &atlasv2.CloudDatabaseUser{}

	var expectedProjectAccessLists *atlasv2.PaginatedNetworkAccess

	opts := &Opts{
		ClusterName:                 "ProjectBar",
		Region:                      "US",
		store:                       mockStore,
		IPAddresses:                 []string{"0.0.0.0"},
		DBUsername:                  "user",
		DBUserPassword:              "test",
		Provider:                    "AWS",
		EnableTerminationProtection: true,
		SkipMongosh:                 true,
		SkipSampleData:              true,
		Confirm:                     true,
		MDBVersion:                  "7.0",
		AutoScalingMode:             clusterWideScaling,
	}
	opts.register.WithFlow(mockFlow)

	projectIPAccessList := opts.newProjectIPAccessList()

	mockStore.
		EXPECT().
		CreateCluster(opts.newCluster()).Return(expectedCluster, nil).
		Times(1)

	mockStore.
		EXPECT().
		CreateProjectIPAccessList(projectIPAccessList).Return(expectedProjectAccessLists, nil).
		Times(1)

	mockStore.
		EXPECT().
		LatestAtlasCluster(opts.ConfigProjectID(), opts.ClusterName).Return(expectedClusterLatest, nil).
		Times(2)

	mockStore.
		EXPECT().
		CreateDatabaseUser(opts.newDatabaseUser()).Return(expectedDBUser, nil).
		Times(1)

	cmd := Builder()
	cmd.Flags().VisitAll(func(f *pflag.Flag) {
		f.Changed = false
		_ = cmd.Flags().Set(f.Name, f.DefValue)
	})

	opts.flags = cmd.Flags()

	if err := opts.setupCluster(); err != nil {
		t.Fatalf("Run() unexpected error: %v", err)
	}

	assert.False(t, opts.shouldAskForValue(flag.ClusterName))
	assert.False(t, opts.shouldAskForValue(flag.Region))
	assert.False(t, opts.shouldAskForValue(flag.AccessListIP))
	assert.False(t, opts.shouldAskForValue(flag.Region))
	assert.False(t, opts.shouldAskForValue(flag.Username))
	assert.False(t, opts.shouldAskForValue(flag.Password))
	assert.False(t, opts.shouldAskForValue(flag.EnableTerminationProtection))
}
