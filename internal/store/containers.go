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

//go:generate mockgen -destination=../mocks/mock_containers.go -package=mocks github.com/mongodb/mongocli/internal/store ContainersLister,ContainersDeleter

type ContainersLister interface {
	ContainersByProvider(string, *atlas.ContainersListOptions) ([]atlas.Container, error)
	AllContainers(string, *atlas.ListOptions) ([]atlas.Container, error)
}

type ContainersDeleter interface {
	DeleteContainer(string, string) error
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

const maxPerPage = 100

// AzureContainers encapsulates the logic to manage different cloud providers
func (s *Store) AzureContainers(projectID string) ([]atlas.Container, error) {
	switch s.service {
	case config.CloudService:
		opts := &atlas.ContainersListOptions{
			ProviderName: "Azure",
			ListOptions: atlas.ListOptions{
				PageNum:      0,
				ItemsPerPage: maxPerPage,
			},
		}
		result, _, err := s.client.(*atlas.Client).Containers.List(context.Background(), projectID, opts)
		return result, err
	default:
		return nil, fmt.Errorf("unsupported service: %s", s.service)
	}
}

// AWSContainers encapsulates the logic to manage different cloud providers
func (s *Store) AWSContainers(projectID string) ([]atlas.Container, error) {
	switch s.service {
	case config.CloudService:
		opts := &atlas.ContainersListOptions{
			ProviderName: "AWS",
			ListOptions: atlas.ListOptions{
				PageNum:      0,
				ItemsPerPage: maxPerPage,
			},
		}
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

// DeleteContainer encapsulates the logic to manage different cloud providers
func (s *Store) DeleteContainer(projectID, containerID string) error {
	switch s.service {
	case config.CloudService:
		_, err := s.client.(*atlas.Client).Containers.Delete(context.Background(), projectID, containerID)
		return err
	default:
		return fmt.Errorf("unsupported service: %s", s.service)
	}
}

// Container encapsulates the logic to manage different cloud providers
func (s *Store) Container(projectID, containerID string) (*atlas.Container, error) {
	switch s.service {
	case config.CloudService:
		result, _, err := s.client.(*atlas.Client).Containers.Get(context.Background(), projectID, containerID)
		return result, err
	default:
		return nil, fmt.Errorf("unsupported service: %s", s.service)
	}
}

// CreateContainer encapsulates the logic to manage different cloud providers
func (s *Store) CreateContainer(projectID string, container *atlas.Container) (*atlas.Container, error) {
	switch s.service {
	case config.CloudService:
		result, _, err := s.client.(*atlas.Client).Containers.Create(context.Background(), projectID, container)
		return result, err
	default:
		return nil, fmt.Errorf("unsupported service: %s", s.service)
	}
}
