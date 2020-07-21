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
	"context"
	"fmt"

	"github.com/mongodb/mongocli/internal/config"
	atlas "go.mongodb.org/atlas/mongodbatlas"
)

//go:generate mockgen -destination=../mocks/mock_search.go -package=mocks github.com/mongodb/mongocli/internal/store SearchIndexLister,SearchIndexCreator,SearchIndexDescriber,SearchIndexDeleter

type SearchIndexLister interface {
	SearchIndexes(string, string, string, string, *atlas.ListOptions) ([]*atlas.SearchIndex, error)
}

type SearchIndexCreator interface {
	CreateSearchIndexes(string, string, *atlas.SearchIndex) (*atlas.SearchIndex, error)
}

type SearchIndexDescriber interface {
	SearchIndex(string, string, string) (*atlas.SearchIndex, error)
}

type SearchIndexDeleter interface {
	DeleteSearchIndex(string, string, string) error
}

// SearchIndexes encapsulate the logic to manage different cloud providers
func (s *Store) SearchIndexes(projectID, clusterName, dbName, collName string, opts *atlas.ListOptions) ([]*atlas.SearchIndex, error) {
	switch s.service {
	case config.CloudService:
		result, _, err := s.client.(*atlas.Client).Search.ListIndexes(context.Background(), projectID, clusterName, dbName, collName, opts)
		return result, err
	default:
		return nil, fmt.Errorf("unsupported service: %s", s.service)
	}
}

// CreateSearchIndexes encapsulate the logic to manage different cloud providers
func (s *Store) CreateSearchIndexes(projectID, clusterName string, index *atlas.SearchIndex) (*atlas.SearchIndex, error) {
	switch s.service {
	case config.CloudService:
		result, _, err := s.client.(*atlas.Client).Search.CreateIndex(context.Background(), projectID, clusterName, index)
		return result, err
	default:
		return nil, fmt.Errorf("unsupported service: %s", s.service)
	}
}

// SearchIndex encapsulate the logic to manage different cloud providers
func (s *Store) SearchIndex(projectID, clusterName, indexID string) (*atlas.SearchIndex, error) {
	switch s.service {
	case config.CloudService:
		index, _, err := s.client.(*atlas.Client).Search.GetIndex(context.Background(), projectID, clusterName, indexID)
		return index, err
	default:
		return nil, fmt.Errorf("unsupported service: %s", s.service)
	}
}

// DeleteSearchIndex encapsulate the logic to manage different cloud providers
func (s *Store) DeleteSearchIndex(projectID, clusterName, indexID string) error {
	switch s.service {
	case config.CloudService:
		_, err := s.client.(*atlas.Client).Search.DeleteIndex(context.Background(), projectID, clusterName, indexID)
		return err
	default:
		return fmt.Errorf("unsupported service: %s", s.service)
	}
}
