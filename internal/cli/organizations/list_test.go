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

package organizations

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/flag"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/mocks"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/pointer"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/test"
	atlasv2 "go.mongodb.org/atlas-sdk/v20240530002/admin"
)

func TestList_Run(t *testing.T) {
	tests := []struct {
		name      string
		expected  *atlasv2.PaginatedOrganization
		returnErr error
	}{
		{
			name: "non-nil result",
			expected: &atlasv2.PaginatedOrganization{
				Results: &[]atlasv2.AtlasOrganization{{}, {}, {}},
			},
		},
		{
			name: "nil result",
			expected: &atlasv2.PaginatedOrganization{
				Results: nil,
			},
		},
		{
			name:     "no results",
			expected: &atlasv2.PaginatedOrganization{},
		},
		{
			name: "no results",
			expected: &atlasv2.PaginatedOrganization{
				Results: &[]atlasv2.AtlasOrganization{},
			},
		},
		{
			name: "with results",
			expected: &atlasv2.PaginatedOrganization{
				Results: &[]atlasv2.AtlasOrganization{
					{
						Id:   pointer.Get("test"),
						Name: "test",
					},
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mockStore := mocks.NewMockOrganizationLister(ctrl)
			listOpts := &ListOpts{store: mockStore}

			mockStore.
				EXPECT().
				Organizations(listOpts.newOrganizationListOptions()).
				Return(tt.expected, nil).
				Times(1)

			if err := listOpts.Run(); err != nil {
				t.Fatalf("Run() unexpected error: %v", err)
			}

			err := listOpts.Print(tt.expected)
			if err != nil {
				t.Fatalf("Print() unexpected error: %v", err)
			}

			test.VerifyOutputTemplate(t, listTemplate, tt.expected)
		})
	}
}

func TestListBuilder(t *testing.T) {
	test.CmdValidator(
		t,
		ListBuilder(),
		0,
		[]string{flag.Page, flag.Limit, flag.OmitCount, flag.IncludeDeleted, flag.Output, flag.Name},
	)
}
