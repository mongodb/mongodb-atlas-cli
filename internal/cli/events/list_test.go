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
	"testing"

	"github.com/stretchr/testify/assert"
	"go.mongodb.org/atlas-sdk/v20250312002/admin"
	"go.uber.org/mock/gomock"
)

func TestList_Run(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockStore := NewMockEventLister(ctrl)

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
		assert.Contains(t, err.Error(), "parsing time")
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
		assert.Contains(t, err.Error(), "parsing time")
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
		assert.Contains(t, err.Error(), "parsing time")
	})
}
