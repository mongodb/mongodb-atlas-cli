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

package events

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/mongodb/mongodb-atlas-cli/internal/flag"
	mocks "github.com/mongodb/mongodb-atlas-cli/internal/mocks/atlas"
	"github.com/mongodb/mongodb-atlas-cli/internal/test"
	"go.mongodb.org/atlas-sdk/v20231115001/admin"
)

func TestList_Run(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockStore := mocks.NewMockEventLister(ctrl)

	t.Run("for a org", func(t *testing.T) {
		expected := &admin.OrgPaginatedEvent{}
		listOpts := &ListOpts{
			store: mockStore,
		}
		listOpts.orgID = "1"
		anyMock := gomock.Any()
		mockStore.
			EXPECT().OrganizationEvents(anyMock).
			Return(expected, nil).
			Times(1)

		err := listOpts.Run()
		if err != nil {
			t.Fatalf("Run() unexpected error: %v", err)
		}
	})
	t.Run("for an project", func(t *testing.T) {
		expected := &admin.GroupPaginatedEvent{}
		listOpts := &ListOpts{
			store: mockStore,
		}

		anyMock := gomock.Any()
		listOpts.projectID = "1"
		mockStore.
			EXPECT().ProjectEvents(anyMock).
			Return(expected, nil).
			Times(1)

		err := listOpts.Run()
		if err != nil {
			t.Fatalf("Run() unexpected error: %v", err)
		}
	})
}

func TestListBuilder(t *testing.T) {
	test.CmdValidator(
		t,
		ListBuilder(),
		0,
		[]string{
			flag.Limit,
			flag.Page,
			flag.Output,
			flag.ProjectID,
			flag.OrgID,
			flag.TypeFlag,
			flag.MaxDate,
			flag.MinDate,
		},
	)
}
