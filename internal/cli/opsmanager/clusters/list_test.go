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

//go:build unit

package clusters

import (
	"bytes"
	"testing"

	"github.com/andreangiolillo/mongocli-test/internal/config"
	"github.com/andreangiolillo/mongocli-test/internal/flag"
	"github.com/andreangiolillo/mongocli-test/internal/mocks"
	"github.com/andreangiolillo/mongocli-test/internal/test"
	"github.com/andreangiolillo/mongocli-test/internal/test/fixture"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"go.mongodb.org/ops-manager/opsmngr"
)

func TestList_Run(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockStore := mocks.NewMockCloudManagerClustersLister(ctrl)

	t.Run("clusters for project simple view", func(t *testing.T) {
		expected := &opsmngr.Clusters{}

		listOpts := &ListOpts{
			store: mockStore,
		}
		buf := new(bytes.Buffer)
		listOpts.OutWriter = buf
		listOpts.ProjectID = "1"
		mockStore.
			EXPECT().
			ProjectClusters(listOpts.ProjectID, nil).
			Return(expected, nil).
			Times(1)

		require.NoError(t, listOpts.Run())
		test.VerifyOutputTemplate(t, listTemplate, expected)
	})

	t.Run("clusters for project json view", func(t *testing.T) {
		expected := fixture.AutomationConfig()
		listOpts := &ListOpts{
			store: mockStore,
		}
		buf := new(bytes.Buffer)
		listOpts.OutWriter = buf
		listOpts.ProjectID = "1"
		listOpts.Output = config.JSON
		mockStore.
			EXPECT().
			GetAutomationConfig(listOpts.ProjectID).
			Return(expected, nil).
			Times(1)
		require.NoError(t, listOpts.Run())
	})
}

func TestListBuilder(t *testing.T) {
	test.CmdValidator(t,
		ListBuilder(),
		0,
		[]string{
			flag.Output,
			flag.ProjectID,
		})
}
