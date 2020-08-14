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

//go:generate mockgen -destination=../mocks/mock_containers.go -package=mocks github.com/mongodb/mongocli/internal/store ContainersLister

type ContainersLister interface {
	ContainersByProvider(string, *atlas.ContainersListOptions) ([]atlas.Container, error)
	AllContainers(string, *atlas.ListOptions) ([]atlas.Container, error)
}

// ContainersByProvider encapsulates the logic to manage different cloud providers
func (s *Store) ContainersByProvider(projectID string, opts *atlas.ContainersListOptions) ([]atlas.Container, error) {
	switch s.service {
	case config.CloudService:
		result, _, err := s.client.(*atlas.Client).Containers.List(context.Background(), projectID, opts)
		return result, err
	default:
		return nil, fmt.Errorf("unsupported service: %s", s.service)
	}
}

// AllContainers encapsulates the logic to manage different cloud providers
func (s *Store) AllContainers(projectID string, opts *atlas.ListOptions) ([]atlas.Container, error) {
	switch s.service {
	case config.CloudService:
		result, _, err := s.client.(*atlas.Client).Containers.ListAll(context.Background(), projectID, opts)
		return result, err
	default:
		return nil, fmt.Errorf("unsupported service: %s", s.service)
	}
}
