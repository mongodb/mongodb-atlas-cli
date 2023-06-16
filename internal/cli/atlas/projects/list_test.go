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
	"github.com/mongodb/mongodb-atlas-cli/internal/config"
	mocks "github.com/mongodb/mongodb-atlas-cli/internal/mocks/atlas"
	atlasv2 "go.mongodb.org/atlas-sdk/admin"
)

func TestList_Run(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockStore := mocks.NewMockProjectLister(ctrl)

	expected := &atlasv2.PaginatedAtlasGroup{}

	t.Run("No ConfigOrgID is given", func(t *testing.T) {
		listOpts := &ListOpts{
			store: mockStore,
		}
		mockStore.
			EXPECT().
			Projects(listOpts.NewListOptions()).
			Return(expected, nil).
			Times(1)

		if err := listOpts.Run(); err != nil {
			t.Fatalf("Run() unexpected error: %v", err)
		}
	})

	t.Run("An ConfigOrgID is given for Atlas", func(t *testing.T) {
		listOpts := &ListOpts{
			store: mockStore,
		}
		listOpts.OrgID = "1"

		mockStore.
			EXPECT().
			Projects(listOpts.NewListOptions()).
			Return(expected, nil).
			Times(1)

		config.SetService(config.CloudService)
		if err := listOpts.Run(); err != nil {
			t.Fatalf("Run() unexpected error: %v", err)
		}
	})
}
