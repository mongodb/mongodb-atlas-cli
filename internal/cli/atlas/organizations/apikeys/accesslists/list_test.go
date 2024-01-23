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

package accesslists

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/mongodb/mongodb-atlas-cli/internal/flag"
	mocks "github.com/mongodb/mongodb-atlas-cli/internal/mocks/atlas"
	"github.com/mongodb/mongodb-atlas-cli/internal/pointer"
	"github.com/mongodb/mongodb-atlas-cli/internal/test"
	"go.mongodb.org/atlas-sdk/v20231115004/admin"
)

func TestListOpts_Run(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockStore := mocks.NewMockOrganizationAPIKeyAccessListLister(ctrl)

	opts := &ListOpts{
		store: mockStore,
	}
	listOpts := opts.NewListOptions()

	params := &admin.ListApiKeyAccessListsEntriesApiParams{
		OrgId:        opts.OrgID,
		ApiUserId:    opts.id,
		PageNum:      pointer.Get(listOpts.PageNum),
		ItemsPerPage: pointer.Get(listOpts.ItemsPerPage),
	}
	expected := admin.PaginatedApiUserAccessList{
		Results: []admin.UserAccessList{},
	}
	mockStore.
		EXPECT().
		OrganizationAPIKeyAccessLists(params).
		Return(&expected, nil).
		Times(1)

	if err := opts.Run(); err != nil {
		t.Fatalf("Run() unexpected error: %v", err)
	}
	test.VerifyOutputTemplate(t, listTemplate, expected)
}

func TestListBuilder(t *testing.T) {
	test.CmdValidator(
		t,
		ListBuilder(),
		0,
		[]string{flag.OrgID, flag.Output, flag.Page, flag.Limit},
	)
}
