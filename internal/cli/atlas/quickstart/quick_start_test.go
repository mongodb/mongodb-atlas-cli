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

package quickstart

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/mongodb/mongocli/internal/convert"
	"github.com/mongodb/mongocli/internal/mocks"
	"go.mongodb.org/atlas/mongodbatlas"
)

func TestQuickstartOpts_Run(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockStore := mocks.NewMockAtlasClusterQuickStarter(ctrl)
	defer ctrl.Finish()

	expectedCluster := &mongodbatlas.Cluster{
		StateName: "IDLE",
	}
	expectedDBUser := &mongodbatlas.DatabaseUser{}
	var expectedWhitelist []mongodbatlas.ProjectIPWhitelist

	opts := &Opts{
		clusterName: "ProjectBar",
		region:      "US",
		store:       mockStore,
		ipAddress:   "0.0.0.0",
		dbUsername:  "user",
	}

	whitelist := opts.newWhitelist()

	mockStore.
		EXPECT().
		CreateCluster(opts.newCluster()).Return(expectedCluster, nil).
		Times(1)

	mockStore.
		EXPECT().
		CreateProjectIPAccessList(whitelist).Return(expectedWhitelist, nil).
		Times(1)

	mockStore.
		EXPECT().
		CreateDatabaseUser(opts.newDatabaseUser()).Return(expectedDBUser, nil).
		Times(1)

	mockStore.
		EXPECT().
		DatabaseUser(convert.AdminDB, opts.ConfigProjectID(), opts.dbUsername).Return(nil, nil).
		Times(1)

	mockStore.
		EXPECT().
		AtlasCluster(opts.ConfigProjectID(), opts.clusterName).Return(expectedCluster, nil).
		Times(2)

	err := opts.Run()
	if err != nil {
		t.Fatalf("Run() unexpected error: %v", err)
	}
}
