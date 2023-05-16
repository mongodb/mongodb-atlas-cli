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

package quickstart

import (
	"bytes"
	"context"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/mongodb/mongodb-atlas-cli/internal/flag"
	"github.com/mongodb/mongodb-atlas-cli/internal/mocks"
	"github.com/mongodb/mongodb-atlas-cli/internal/test"
	"github.com/mongodb/mongodb-atlas-cli/internal/validate"
	"github.com/spf13/pflag"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.mongodb.org/atlas/mongodbatlas"
	atlasv2 "go.mongodb.org/atlas/mongodbatlasv2"
)

func TestBuilder(t *testing.T) {
	test.CmdValidator(
		t,
		Builder(),
		0,
		[]string{
			flag.ProjectID,
			flag.Region,
			flag.ClusterName,
			flag.Provider,
			flag.AccessListIP,
			flag.Username,
			flag.Password,
			flag.EnableTerminationProtection,
			flag.SkipMongosh,
			flag.SkipSampleData,
		},
	)
}

func TestQuickstartOpts_Run(t *testing.T) {
	t.Cleanup(test.CleanupConfig)
	ctrl := gomock.NewController(t)
	mockStore := mocks.NewMockAtlasClusterQuickStarter(ctrl)
	mockFlow := mocks.NewMockRefresher(ctrl)

	expectedCluster := &mongodbatlas.AdvancedCluster{
		StateName: "IDLE",
		ConnectionStrings: &mongodbatlas.ConnectionStrings{
			StandardSrv: "",
		},
	}

	expectedDBUser := &atlasv2.DatabaseUser{}

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
	}
	opts.WithFlow(mockFlow)

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

	if err := opts.Run(); err != nil {
		t.Fatalf("Run() unexpected error: %v", err)
	}
}

func TestQuickstartOpts_Run_NotLoggedIn(t *testing.T) {
	t.Cleanup(test.CleanupConfig)
	ctrl := gomock.NewController(t)
	mockStore := mocks.NewMockAtlasClusterQuickStarter(ctrl)
	mockFlow := mocks.NewMockRefresher(ctrl)

	buf := new(bytes.Buffer)
	ctx := context.TODO()
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
	}
	opts.WithFlow(mockFlow)

	require.Error(t, validate.ErrMissingCredentials, opts.quickstartPreRun(ctx, buf))
}

func TestQuickstartOpts_Run_CheckFlagsSet(t *testing.T) {
	t.Cleanup(test.CleanupConfig)
	ctrl := gomock.NewController(t)
	mockStore := mocks.NewMockAtlasClusterQuickStarter(ctrl)
	mockFlow := mocks.NewMockRefresher(ctrl)
	defer ctrl.Finish()

	expectedCluster := &mongodbatlas.AdvancedCluster{
		StateName: "IDLE",
		ConnectionStrings: &mongodbatlas.ConnectionStrings{
			StandardSrv: "",
		},
	}

	expectedDBUser := &atlasv2.DatabaseUser{}

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
	opts.WithFlow(mockFlow)

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

	if err := opts.Run(); err != nil {
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
