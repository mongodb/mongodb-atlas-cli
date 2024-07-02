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

package projects

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/flag"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/mocks"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/test"
	"github.com/stretchr/testify/require"
	atlasv2 "go.mongodb.org/atlas-sdk/v20240530002/admin"
)

func TestList_Run(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockStore := mocks.NewMockOrgProjectLister(ctrl)
	expected := &atlasv2.PaginatedAtlasGroup{}
	t.Run("No OrgID is given", func(t *testing.T) {
		listOpts := &ListOpts{
			store: mockStore,
		}
		mockStore.
			EXPECT().
			Projects(listOpts.NewAtlasListOptions()).
			Return(expected, nil).
			Times(1)
		require.NoError(t, listOpts.Run())
	})
	t.Run("An OrgID is given for Atlas", func(t *testing.T) {
		listOpts := &ListOpts{
			store: mockStore,
		}
		listOpts.OrgID = "1"

		mockStore.
			EXPECT().
			GetOrgProjects(listOpts.OrgID, listOpts.NewAtlasListOptions()).
			Return(expected, nil).
			Times(1)
		require.NoError(t, listOpts.Run())
	})
}

func TestListBuilder(t *testing.T) {
	test.CmdValidator(
		t,
		ListBuilder(),
		0,
		[]string{flag.Page, flag.Limit, flag.OmitCount, flag.OrgID, flag.Output},
	)
}
