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
	"errors"
	"testing"

	"github.com/spf13/afero"
	atlasv2 "go.mongodb.org/atlas-sdk/v20250312005/admin"
	"go.uber.org/mock/gomock"
)

const (
	testName        = "default"
	testJSON        = `{"name":"default"}`
	testInvalidJSON = `{"name:"default"}`
	fileName        = "atlas_search_index_create_test.json"
)

func TestCreateOpts_Run(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockStore := NewMockCreator(ctrl)

	t.Run("flags run", func(t *testing.T) {
		opts := &CreateOpts{
			store: mockStore,
		}
		opts.Name = testName

		request, err := opts.CreateSearchIndex()
		if err != nil {
			t.Fatalf("newSearchIndex() unexpected error: %v", err)
		}
		expected := &atlasv2.ClusterSearchIndex{}
		mockStore.
			EXPECT().
			CreateSearchIndexesDeprecated(opts.ProjectID, opts.clusterName, request).
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
		request, err := opts.CreateSearchIndex()
		if err != nil {
			t.Fatalf("CreateSearchIndex() unexpected error: %v", err)
		}
		mockStore.
			EXPECT().
			CreateSearchIndexesDeprecated(opts.ProjectID, opts.clusterName, request).Return(expected, nil).
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

		_, err := opts.CreateSearchIndex()
		if err == nil {
			t.Fatalf("CreateSearchIndex() expected error")
		}

		if !errors.Is(err, errFailedToLoadIndexMessage) {
			t.Fatalf("CreateSearchIndex() unexpected error: %v expected: %s", err, errFailedToLoadIndexMessage)
		}
	})
}
