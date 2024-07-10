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
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/test"
	atlasv2 "go.mongodb.org/atlas-sdk/v20240530002/admin"
)

func TestCreateBuilder(t *testing.T) {
	test.CmdValidator(
		t,
		CreateBuilder(),
		0,
		[]string{
			flag.ClusterName,
			flag.Database,
			flag.Collection,
			flag.DateField,
			flag.DateFormat,
			flag.ArchiveAfter,
			flag.Partition,
			flag.ProjectID,
			flag.Output,
		},
	)
}

func TestCreateOpts_Run(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockStore := mocks.NewMockOnlineArchiveCreator(ctrl)

	opts := &CreateOpts{
		store: mockStore,
	}

	request, _ := opts.newOnlineArchive()

	expected := &atlasv2.BackupOnlineArchive{}
	mockStore.
		EXPECT().
		CreateOnlineArchive(opts.ProjectID, opts.clusterName, request).
		Return(expected, nil).
		Times(1)

	if err := opts.Run(); err != nil {
		t.Fatalf("Run() unexpected error: %v", err)
	}
}
