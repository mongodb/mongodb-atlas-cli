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

//go:generate mockgen -destination=../mocks/mock_data_lake.go -package=mocks github.com/mongodb/mongocli/internal/store DataLakeLister,DataLakeDescriber,DataLakeCreator,DataLakeDeleter,DataLakeUpdater,DataLakeStore

type DataLakeLister interface {
	DataLakes(string) ([]atlas.DataLake, error)
}

type DataLakeDescriber interface {
	DataLake(string, string) (*atlas.DataLake, error)
}

type DataLakeCreator interface {
	CreateDataLake(string, *atlas.DataLakeCreateRequest) (*atlas.DataLake, error)
}

type DataLakeDeleter interface {
	DeleteDataLake(string, string) error
}

type DataLakeUpdater interface {
	UpdateDataLake(string, string, *atlas.DataLakeUpdateRequest) (*atlas.DataLake, error)
}

type DataLakeStore interface {
	DataLakeLister
	DataLakeDescriber
	DataLakeCreator
	DataLakeDeleter
	DataLakeUpdater
}

// CreateDataLake encapsulate the logic to manage different cloud providers
func (s *Store) CreateDataLake(projectID string, dataLake *atlas.DataLakeCreateRequest) (*atlas.DataLake, error) {
	switch s.service {
	case config.CloudService:
		result, _, err := s.client.(*atlas.Client).DataLakes.Create(context.Background(), projectID, dataLake)
		return result, err
	default:
		return nil, fmt.Errorf("unsupported service: %s", s.service)
	}
}

// UpdateDataLake encapsulate the logic to manage different cloud providers
func (s *Store) UpdateDataLake(projectID, name string, dataLake *atlas.DataLakeUpdateRequest) (*atlas.DataLake, error) {
	switch s.service {
	case config.CloudService:
		result, _, err := s.client.(*atlas.Client).DataLakes.Update(context.Background(), projectID, name, dataLake)
		return result, err
	default:
		return nil, fmt.Errorf("unsupported service: %s", s.service)
	}
}

// DeleteDataLake encapsulate the logic to manage different cloud providers
func (s *Store) DeleteDataLake(projectID, name string) error {
	switch s.service {
	case config.CloudService:
		_, err := s.client.(*atlas.Client).DataLakes.Delete(context.Background(), projectID, name)
		return err
	default:
		return fmt.Errorf("unsupported service: %s", s.service)
	}
}

// DataLakes encapsulate the logic to manage different cloud providers
func (s *Store) DataLakes(projectID string) ([]atlas.DataLake, error) {
	switch s.service {
	case config.CloudService:
		result, _, err := s.client.(*atlas.Client).DataLakes.List(context.Background(), projectID)
		return result, err
	default:
		return nil, fmt.Errorf("unsupported service: %s", s.service)
	}
}

// DataLake encapsulate the logic to manage different cloud providers
func (s *Store) DataLake(projectID, name string) (*atlas.DataLake, error) {
	switch s.service {
	case config.CloudService:
		result, _, err := s.client.(*atlas.Client).DataLakes.Get(context.Background(), projectID, name)
		return result, err
	default:
		return nil, fmt.Errorf("unsupported service: %s", s.service)
	}
}
