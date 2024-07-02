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

package databases

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/flag"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/mocks"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/pointer"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/test"
	atlasv2 "go.mongodb.org/atlas-sdk/v20240530002/admin"
)

func TestDatabasesListsOpts_Run(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockStore := mocks.NewMockProcessDatabaseLister(ctrl)

	expected := &atlasv2.PaginatedDatabase{
		Results: &[]atlasv2.MesurementsDatabase{
			{
				DatabaseName: pointer.Get("test"),
			},
		},
	}

	listOpts := &ListsOpts{
		host:  "hard-00-00.mongodb.net",
		port:  27017,
		store: mockStore,
	}

	opts := listOpts.NewListOptions()
	mockStore.
		EXPECT().ProcessDatabases(listOpts.ProjectID, listOpts.host, listOpts.port, opts).
		Return(expected, nil).
		Times(1)

	if err := listOpts.Run(); err != nil {
		t.Fatalf("Run() unexpected error: %v", err)
	}
	test.VerifyOutputTemplate(t, databasesListTemplate, expected)
}

func TestListBuilder(t *testing.T) {
	test.CmdValidator(
		t,
		ListBuilder(),
		0,
		[]string{flag.Page, flag.Limit, flag.OmitCount, flag.ProjectID, flag.Output},
	)
}
