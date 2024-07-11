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

package search

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/mocks"
	"github.com/spf13/afero"
	"github.com/stretchr/testify/require"
	atlasv2 "go.mongodb.org/atlas-sdk/v20240530002/admin"
)

func TestUpdateOpts_Run(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockStore := mocks.NewMockSearchIndexUpdater(ctrl)

	t.Run("flags run", func(t *testing.T) {
		updateOpts := &UpdateOpts{
			store: mockStore,
		}
		updateOpts.Name = testName
		updateOpts.id = "1"

		expected := &atlasv2.ClusterSearchIndex{}

		request, err := updateOpts.NewSearchIndex()
		require.NoError(t, err)
		mockStore.
			EXPECT().
			UpdateSearchIndexes(updateOpts.ConfigProjectID(), updateOpts.clusterName, updateOpts.id, request).
			Return(expected, nil).
			Times(1)

		require.NoError(t, updateOpts.Run())
	})

	t.Run("file run", func(t *testing.T) {
		appFS := afero.NewMemMapFs()
		// create test file
		fileName := "atlas_search_index_update_test.json"
		_ = afero.WriteFile(appFS, fileName, []byte(testJSON), 0600)

		updateOpts := &UpdateOpts{
			store: mockStore,
		}
		updateOpts.id = "1"
		updateOpts.Filename = fileName
		updateOpts.Fs = appFS

		expected := &atlasv2.ClusterSearchIndex{}

		request, err := updateOpts.NewSearchIndex()
		require.NoError(t, err)
		mockStore.
			EXPECT().
			UpdateSearchIndexes(updateOpts.ConfigProjectID(), updateOpts.clusterName, updateOpts.id, request).
			Return(expected, nil).
			Times(1)

		require.NoError(t, updateOpts.Run())
	})
}
