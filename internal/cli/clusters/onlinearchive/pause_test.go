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

package onlinearchive

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/flag"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/mocks"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/pointer"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/test"
	atlasv2 "go.mongodb.org/atlas-sdk/v20240805003/admin"
)

func TestPauseBuilder(t *testing.T) {
	test.CmdValidator(
		t,
		PauseBuilder(),
		0,
		[]string{
			flag.ClusterName,
			flag.Output,
			flag.ProjectID,
		},
	)
}

func TestPause_Run(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockStore := mocks.NewMockOnlineArchiveUpdater(ctrl)

	updateOpts := &PauseOpts{
		id:    "1",
		store: mockStore,
	}

	expected := &atlasv2.BackupOnlineArchive{
		Id:    &updateOpts.id,
		State: pointer.Get("PAUSING"),
	}

	mockStore.
		EXPECT().
		UpdateOnlineArchive(updateOpts.ConfigProjectID(), updateOpts.clusterName, expected).
		Return(expected, nil).
		Times(1)

	if err := updateOpts.Run(); err != nil {
		t.Fatalf("Run() unexpected error: %v", err)
	}
}
