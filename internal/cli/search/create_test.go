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
	"strings"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/mocks"
	"github.com/spf13/afero"
	atlasv2 "go.mongodb.org/atlas-sdk/v20240530002/admin"
)

const testName = "default"
const testJSON = `{"name":"default"}`
const testInvalidJSON = `{"name:"default"}`
const fileName = "atlas_search_index_create_test.json"

func TestCreateOpts_Run(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockStore := mocks.NewMockSearchIndexCreator(ctrl)

	t.Run("flags run", func(t *testing.T) {
		opts := &CreateOpts{
			store: mockStore,
		}
		opts.Name = testName

		request, err := opts.NewSearchIndex()
		if err != nil {
			t.Fatalf("newSearchIndex() unexpected error: %v", err)
		}
		expected := &atlasv2.ClusterSearchIndex{}
		mockStore.
			EXPECT().
			CreateSearchIndexes(opts.ProjectID, opts.clusterName, request).
			Return(expected, nil).
			Times(1)

		if err := opts.Run(); err != nil {
			t.Fatalf("Run() unexpected error: %v", err)
		}
	})

	t.Run("file run", func(t *testing.T) {
		appFS := afero.NewMemMapFs()
		// create test file
		_ = afero.WriteFile(appFS, fileName, []byte(testJSON), 0600)

		opts := &CreateOpts{
			store: mockStore,
		}
		opts.Filename = fileName
		opts.Fs = appFS

		expected := &atlasv2.ClusterSearchIndex{}
		request, err := opts.NewSearchIndex()
		if err != nil {
			t.Fatalf("newSearchIndex() unexpected error: %v", err)
		}
		mockStore.
			EXPECT().
			CreateSearchIndexes(opts.ProjectID, opts.clusterName, request).Return(expected, nil).
			Times(1)
		if err := opts.Run(); err != nil {
			t.Fatalf("Run() unexpected error: %v", err)
		}
	})

	t.Run("invalid file run", func(t *testing.T) {
		appFS := afero.NewMemMapFs()
		// create test file
		_ = afero.WriteFile(appFS, fileName, []byte(testInvalidJSON), 0600)

		opts := &CreateOpts{
			store: mockStore,
		}
		opts.Filename = fileName
		opts.Fs = appFS

		_, err := opts.NewSearchIndex()
		if err == nil {
			t.Fatalf("newSearchIndex() expected error")
		}

		expectedError := "failed to parse JSON file due to"
		if !strings.Contains(err.Error(), expectedError) {
			t.Fatalf("newSearchIndex() unexpected error: %v expected: %s", err, expectedError)
		}
	})
}
