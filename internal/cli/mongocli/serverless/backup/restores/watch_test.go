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
	"github.com/mongodb/mongodb-atlas-cli/internal/flag"
	"github.com/mongodb/mongodb-atlas-cli/internal/mocks"
	"github.com/mongodb/mongodb-atlas-cli/internal/pointer"
	"github.com/mongodb/mongodb-atlas-cli/internal/test"
	atlasv2 "go.mongodb.org/atlas-sdk/admin"
)

func TestWatchBuilder(t *testing.T) {
	test.CmdValidator(
		t,
		WatchBuilder(),
		0,
		[]string{
			flag.ClusterName,
			flag.ProjectID,
			flag.Output,
		},
	)
}

func TestWatchOpts_Run(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockStore := mocks.NewMockServerlessRestoreJobsDescriber(ctrl)

	expected := &atlasv2.ServerlessBackupRestoreJob{
		Failed: pointer.Get(true),
	}

	describeOpts := &WatchOpts{
		store:       mockStore,
		clusterName: "Cluster0",
		id:          "1",
	}

	mockStore.
		EXPECT().
		ServerlessRestoreJob(describeOpts.ProjectID, describeOpts.clusterName, describeOpts.id).
		Return(expected, nil).
		Times(1)

	if err := describeOpts.Run(); err != nil {
		t.Fatalf("Run() unexpected error: %v", err)
	}
}
