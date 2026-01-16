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
	atlasv2 "go.mongodb.org/atlas-sdk/v20250312012/admin"
)

// SearchIndexes encapsulate the logic to manage different cloud providers.
func (s *Store) SearchIndexes(projectID, clusterName, dbName, collName string) ([]atlasv2.SearchIndexResponse, error) {
	result, _, err := s.clientv2.AtlasSearchApi.ListSearchIndex(s.ctx, projectID, clusterName, collName, dbName).Execute()
	return result, err
}

// CreateSearchIndexes encapsulate the logic to manage different cloud providers.
func (s *Store) CreateSearchIndexes(projectID, clusterName string, index *atlasv2.SearchIndexCreateRequest) (*atlasv2.SearchIndexResponse, error) {
	result, _, err := s.clientv2.AtlasSearchApi.CreateClusterSearchIndex(s.ctx, projectID, clusterName, index).Execute()
	return result, err
}

// SearchIndex encapsulate the logic to manage different cloud providers.
func (s *Store) SearchIndex(projectID, clusterName, indexID string) (*atlasv2.SearchIndexResponse, error) {
	index, _, err := s.clientv2.AtlasSearchApi.GetClusterSearchIndex(s.ctx, projectID, clusterName, indexID).Execute()
	return index, err
}

// UpdateSearchIndexes encapsulate the logic to manage different cloud providers.
func (s *Store) UpdateSearchIndexes(projectID, clusterName, indexID string, index *atlasv2.SearchIndexUpdateRequest) (*atlasv2.SearchIndexResponse, error) {
	result, _, err := s.clientv2.AtlasSearchApi.UpdateClusterSearchIndex(s.ctx, projectID, clusterName, indexID, index).Execute()
	return result, err
}

// DeleteSearchIndex encapsulate the logic to manage different cloud providers.
func (s *Store) DeleteSearchIndex(projectID, clusterName, indexID string) error {
	_, err := s.clientv2.AtlasSearchApi.DeleteClusterSearchIndex(s.ctx, projectID, clusterName, indexID).Execute()
	return err
}
