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

package setup

import (
	"bytes"
	"context"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/config"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/flag"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/mocks"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/pointer"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/test"
	"github.com/spf13/pflag"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	atlasv2 "go.mongodb.org/atlas-sdk/v20240530002/admin"
)

func TestBuilder(t *testing.T) {
	test.CmdValidator(
		t,
		Builder(),
		0,
		[]string{
			flag.ClusterName,
			flag.Tier,
			flag.Provider,
			flag.Region,
			flag.AccessListIP,
			flag.Username,
			flag.Password,
			flag.EnableTerminationProtection,
			flag.SkipSampleData,
			flag.SkipMongosh,
			flag.Force,
			flag.CurrentIP,
			flag.Tag,
			flag.Default,
			flag.ProjectID,
		},
	)
}

func Test_setupOpts_PreRunWithAPIKeys(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockFlow := mocks.NewMockRefresher(ctrl)
	ctx := context.TODO()
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
	ctx := context.TODO()
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
	mockStore := mocks.NewMockAtlasClusterQuickStarter(ctrl)
	mockFlow := mocks.NewMockRefresher(ctrl)

	expectedCluster := &atlasv2.AdvancedClusterDescription{
		StateName: pointer.Get("IDLE"),
		ConnectionStrings: &atlasv2.ClusterConnectionStrings{
			StandardSrv: pointer.Get(""),
		},
	}

	expectedDBUser := &atlasv2.CloudDatabaseUser{}

	var expectedProjectAccessLists *atlasv2.PaginatedNetworkAccess

	opts := &Opts{
		ClusterName:    "ProjectBar",
		Region:         "US",
		store:          mockStore,
		IPAddresses:    []string{"0.0.0.0"},
		DBUsername:     "user",
		DBUserPassword: "test",
		Provider:       "AWS",
		SkipMongosh:    true,
		SkipSampleData: true,
		Confirm:        true,
		Tag:            map[string]string{"env": "test"},
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
		AtlasCluster(opts.ConfigProjectID(), opts.ClusterName).Return(expectedCluster, nil).
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
	mockStore := mocks.NewMockAtlasClusterQuickStarter(ctrl)
	mockFlow := mocks.NewMockRefresher(ctrl)
	defer ctrl.Finish()

	expectedCluster := &atlasv2.AdvancedClusterDescription{
		StateName: pointer.Get("IDLE"),
		ConnectionStrings: &atlasv2.ClusterConnectionStrings{
			StandardSrv: pointer.Get(""),
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
		AtlasCluster(opts.ConfigProjectID(), opts.ClusterName).Return(expectedCluster, nil).
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
