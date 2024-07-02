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

package events

import (
	"strings"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/flag"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/mocks"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/test"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/atlas-sdk/v20240530002/admin"
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

	t.Run("for a org with dates", func(t *testing.T) {
		expected := &admin.OrgPaginatedEvent{}
		listOpts := &ListOpts{
			store: mockStore,
			EventListOpts: EventListOpts{
				MaxDate: "2024-03-18T15:00:03-0000",
				MinDate: "2024-03-18T14:40:03-0000",
			},
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
	t.Run("for an project with dates", func(t *testing.T) {
		expected := &admin.GroupPaginatedEvent{}
		listOpts := &ListOpts{
			store: mockStore,
			EventListOpts: EventListOpts{
				MaxDate: "2024-03-18T15:00:03-0000",
				MinDate: "2024-03-18T14:40:03-0000",
			},
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

	t.Run("for a org with invalid dates", func(t *testing.T) {
		listOpts := &ListOpts{
			store: mockStore,
			EventListOpts: EventListOpts{
				MaxDate: "2024-03-18T15:00:03+00:00Z",
				MinDate: "2024-03-18T13:00:03+00:00Z",
			},
		}
		listOpts.orgID = "1"

		err := listOpts.Run()
		if err == nil {
			t.Fatal("Run() expected an error, but got none")
		}
		assert.True(t, strings.Contains(err.Error(), "parsing time"))
	})
	t.Run("for an project with invalid dates", func(t *testing.T) {
		listOpts := &ListOpts{
			store: mockStore,
			EventListOpts: EventListOpts{
				MaxDate: "2024-03-18T15:00:03+00:00Z",
				MinDate: "2024-03-18T13:00:03+00:00Z",
			},
		}

		listOpts.projectID = "1"
		err := listOpts.Run()
		if err == nil {
			t.Fatal("Run() expected an error, but got none")
		}
		assert.True(t, strings.Contains(err.Error(), "parsing time"))
	})
}

func TestParseDate(t *testing.T) {
	t.Run("valid date", func(t *testing.T) {
		date := "2024-03-18T15:00:03-0000"
		_, err := parseDate(date)
		if err != nil {
			t.Fatalf("parseDate() unexpected error: %v", err)
		}
	})
	t.Run("invalid date", func(t *testing.T) {
		date := "2024-03-18T15:00:03+00:00Z"
		_, err := parseDate(date)
		if err == nil {
			t.Fatalf("expected error from parseDate() but got none")
		}
		assert.True(t, strings.Contains(err.Error(), "parsing time"))
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
			flag.OmitCount,
			flag.Output,
			flag.ProjectID,
			flag.OrgID,
			flag.TypeFlag,
			flag.MaxDate,
			flag.MinDate,
		},
	)
}
