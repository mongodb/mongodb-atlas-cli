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

	"github.com/mongodb/mongocli/internal/config"

	"github.com/golang/mock/gomock"
	"github.com/mongodb/mongocli/internal/fixtures"
	"github.com/mongodb/mongocli/internal/mocks"
)

func TestCloudManagerClustersList_Run(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockStore := mocks.NewMockCloudManagerClustersLister(ctrl)

	defer ctrl.Finish()

	t.Run("ProjectID is given", func(t *testing.T) {
		expected := fixtures.AutomationConfig()

		listOpts := &cloudManagerClustersListOpts{
			globalOpts: newGlobalOpts(),
			store:      mockStore,
		}

		listOpts.projectID = "1"
		mockStore.
			EXPECT().
			GetAutomationConfig(listOpts.projectID).
			Return(expected, nil).
			Times(1)

		err := listOpts.Run()
		if err != nil {
			t.Fatalf("Run() unexpected error: %v", err)
		}

	})

	t.Run("No ProjectID is given", func(t *testing.T) {
		expected := fixtures.AllClusters()
		config.SetService(config.OpsManagerService)
		mockStore.
			EXPECT().
			ListAllClustersProjects().
			Return(expected, nil).
			Times(1)

		listOpts := &cloudManagerClustersListOpts{
			globalOpts: newGlobalOpts(),
			store:      mockStore,
		}

		err := listOpts.Run()
		if err != nil {
			t.Fatalf("Run() unexpected error: %v", err)
		}
	})

}
