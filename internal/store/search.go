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

package store

import (
	atlasv2 "go.mongodb.org/atlas-sdk/v20240530002/admin"
)

//go:generate mockgen -destination=../mocks/mock_search.go -package=mocks github.com/mongodb/mongodb-atlas-cli/atlascli/internal/store SearchIndexLister,SearchIndexCreator,SearchIndexDescriber,SearchIndexUpdater,SearchIndexDeleter

type SearchIndexLister interface {
	SearchIndexes(string, string, string, string) ([]atlasv2.ClusterSearchIndex, error)
}

type SearchIndexCreator interface {
	CreateSearchIndexes(string, string, *atlasv2.ClusterSearchIndex) (*atlasv2.ClusterSearchIndex, error)
	SearchIndex(string, string, string) (*atlasv2.ClusterSearchIndex, error)
}

type SearchIndexDescriber interface {
	SearchIndex(string, string, string) (*atlasv2.ClusterSearchIndex, error)
}

type SearchIndexUpdater interface {
	UpdateSearchIndexes(string, string, string, *atlasv2.ClusterSearchIndex) (*atlasv2.ClusterSearchIndex, error)
}

type SearchIndexDeleter interface {
	DeleteSearchIndex(string, string, string) error
}

// SearchIndexes encapsulate the logic to manage different cloud providers.
func (s *Store) SearchIndexes(projectID, clusterName, dbName, collName string) ([]atlasv2.ClusterSearchIndex, error) {
	result, _, err := s.clientv2.AtlasSearchApi.ListAtlasSearchIndexesDeprecated(s.ctx, projectID, clusterName, collName, dbName).Execute()
	return result, err
}

// CreateSearchIndexes encapsulate the logic to manage different cloud providers.
func (s *Store) CreateSearchIndexes(projectID, clusterName string, index *atlasv2.ClusterSearchIndex) (*atlasv2.ClusterSearchIndex, error) {
	result, _, err := s.clientv2.AtlasSearchApi.CreateAtlasSearchIndexDeprecated(s.ctx, projectID, clusterName, index).Execute()
	return result, err
}

// SearchIndex encapsulate the logic to manage different cloud providers.
func (s *Store) SearchIndex(projectID, clusterName, indexID string) (*atlasv2.ClusterSearchIndex, error) {
	index, _, err := s.clientv2.AtlasSearchApi.GetAtlasSearchIndexDeprecated(s.ctx, projectID, clusterName, indexID).Execute()
	return index, err
}

// UpdateSearchIndexes encapsulate the logic to manage different cloud providers.
func (s *Store) UpdateSearchIndexes(projectID, clusterName, indexID string, index *atlasv2.ClusterSearchIndex) (*atlasv2.ClusterSearchIndex, error) {
	result, _, err := s.clientv2.AtlasSearchApi.UpdateAtlasSearchIndexDeprecated(s.ctx, projectID, clusterName, indexID, index).Execute()
	return result, err
}

// DeleteSearchIndex encapsulate the logic to manage different cloud providers.
func (s *Store) DeleteSearchIndex(projectID, clusterName, indexID string) error {
	_, _, err := s.clientv2.AtlasSearchApi.DeleteAtlasSearchIndexDeprecated(s.ctx, projectID, clusterName, indexID).Execute()
	return err
}
