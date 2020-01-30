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

package cli

import (
	"testing"

	"github.com/10gen/mcli/internal/config"
	"github.com/10gen/mcli/internal/fixtures"
	"github.com/10gen/mcli/internal/mocks"
	"github.com/golang/mock/gomock"
)

func TestIAMProjectsList_Run(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockStore := mocks.NewMockProjectLister(ctrl)

	defer ctrl.Finish()

	expected := fixtures.Projects()

	t.Run("No OrgID is given", func(t *testing.T) {
		mockStore.
			EXPECT().
			GetAllProjects().
			Return(expected, nil).
			Times(1)

		listOpts := &iamProjectsListOpts{store: mockStore}
		err := listOpts.Run()
		if err != nil {
			t.Fatalf("Run() unexpected error: %v", err)
		}
	})

	t.Run("An OrgID is given for OM", func(t *testing.T) {
		mockStore.
			EXPECT().
			GetOrgProjects("1").
			Return(expected, nil).
			Times(1)

		listOpts := &iamProjectsListOpts{
			orgID: "1",
			store: mockStore,
		}
		config.SetService(config.OpsManagerService)
		err := listOpts.Run()
		if err != nil {
			t.Fatalf("Run() unexpected error: %v", err)
		}
	})

	t.Run("An OrgID is given for Atlas", func(t *testing.T) {
		mockStore.
			EXPECT().
			GetAllProjects().
			Return(expected, nil).
			Times(1)

		listOpts := &iamProjectsListOpts{
			orgID: "1",
			store: mockStore,
		}
		config.SetService(config.CloudService)
		err := listOpts.Run()
		if err != nil {
			t.Fatalf("Run() unexpected error: %v", err)
		}
	})
}
