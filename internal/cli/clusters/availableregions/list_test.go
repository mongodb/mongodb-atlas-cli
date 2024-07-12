// Copyright 2021 MongoDB Inc
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

package availableregions

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/flag"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/mocks"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/test"
	"go.mongodb.org/atlas-sdk/v20240530002/admin"
)

func TestList_Run_NoFlags(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockStore := mocks.NewMockCloudProviderRegionsLister(ctrl)

	var expected *admin.PaginatedApiAtlasProviderRegions
	var empty []string

	listOpts := &ListOpts{
		store: mockStore,
	}

	mockStore.
		EXPECT().
		CloudProviderRegions(listOpts.ProjectID, "", empty).
		Return(expected, nil).
		Times(1)

	if err := listOpts.Run(); err != nil {
		t.Fatalf("Run() unexpected error: %v", err)
	}
}

func TestList_Run(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockStore := mocks.NewMockCloudProviderRegionsLister(ctrl)

	var expected *admin.PaginatedApiAtlasProviderRegions

	listOpts := &ListOpts{
		store:    mockStore,
		tier:     "M2",
		provider: "AWS",
	}

	mockStore.
		EXPECT().
		CloudProviderRegions(listOpts.ProjectID, listOpts.tier, []string{listOpts.provider}).
		Return(expected, nil).
		Times(1)

	if err := listOpts.Run(); err != nil {
		t.Fatalf("Run() unexpected error: %v", err)
	}
}

func TestListBuilder(t *testing.T) {
	test.CmdValidator(
		t,
		ListBuilder(),
		0,
		[]string{flag.Output, flag.ProjectID, flag.Tier, flag.Provider},
	)
}
