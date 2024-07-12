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

func Test_projectListOpts_Run(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockStore := mocks.NewMockProjectEventLister(ctrl)

	expected := &admin.GroupPaginatedEvent{}
	listOpts := &projectListOpts{
		store: mockStore,
	}
	listOpts.ProjectID = "1"
	anyMock := gomock.Any()
	mockStore.
		EXPECT().ProjectEvents(anyMock).
		Return(expected, nil).
		Times(1)

	if err := listOpts.Run(); err != nil {
		t.Fatalf("Run() unexpected error: %v", err)
	}
}

func Test_projectListOpts_Run_WithDate(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockStore := mocks.NewMockProjectEventLister(ctrl)

	expected := &admin.GroupPaginatedEvent{}
	listOpts := &projectListOpts{
		store: mockStore,
		EventListOpts: EventListOpts{
			MaxDate: "2024-03-18T15:00:03-0000",
			MinDate: "2024-03-18T14:40:03-0000",
		},
	}
	listOpts.ProjectID = "1"
	anyMock := gomock.Any()
	mockStore.
		EXPECT().ProjectEvents(anyMock).
		Return(expected, nil).
		Times(1)

	if err := listOpts.Run(); err != nil {
		t.Fatalf("Run() unexpected error: %v", err)
	}
}

func Test_projectListOpts_Run_WithInvalidDate(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockStore := mocks.NewMockProjectEventLister(ctrl)

	listOpts := &projectListOpts{
		store: mockStore,
		EventListOpts: EventListOpts{
			MaxDate: "2024-03-18T15:00:03+00:00Z",
			MinDate: "2024-03-18T15:00:03+00:00Z",
		},
	}
	listOpts.ProjectID = "1"

	err := listOpts.Run()
	if err == nil {
		t.Fatal("Run() expected error")
	}
	assert.True(t, strings.Contains(err.Error(), "parsing time"))
}

func TestProjectListBuilder(t *testing.T) {
	test.CmdValidator(
		t,
		ProjectListBuilder(),
		0,
		[]string{
			flag.Limit,
			flag.Page,
			flag.Output,
			flag.OmitCount,
			flag.ProjectID,
			flag.TypeFlag,
			flag.MaxDate,
			flag.MinDate,
		},
	)
}

func TestProjectsBuilder(t *testing.T) {
	test.CmdValidator(
		t,
		ProjectsBuilder(),
		1,
		[]string{},
	)
}
