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
	"fmt"

	"github.com/mongodb/mongodb-atlas-cli/internal/config"
	atlas "go.mongodb.org/atlas/mongodbatlas"
	atlasv2 "go.mongodb.org/atlas/mongodbatlasv2"
)

//go:generate mockgen -destination=../mocks/mock_containers.go -package=mocks github.com/mongodb/mongodb-atlas-cli/internal/store ContainersLister,ContainersDeleter

type ContainersLister interface {
	ContainersByProvider(string, *atlas.ContainersListOptions) ([]interface{}, error)
	AllContainers(string, *atlas.ListOptions) ([]interface{}, error)
}

type ContainersDeleter interface {
	DeleteContainer(string, string) error
}

// ContainersByProvider encapsulates the logic to manage different cloud providers.
func (s *Store) ContainersByProvider(projectID string, opts *atlas.ContainersListOptions) ([]interface{}, error) {
	switch s.service {
	case config.CloudService, config.CloudGovService:
		result, _, err := s.clientv2.NetworkPeeringApi.ListPeeringContainerByCloudProvider(s.ctx, projectID).
			IncludeCount(opts.IncludeCount).
			PageNum(int32(opts.PageNum)).
			ItemsPerPage(int32(opts.ItemsPerPage)).
			ProviderName(opts.ProviderName).
			Execute()

		containers := make([]interface{}, len(result.Results))
		for i, container := range result.Results {
			switch v := container.GetActualInstance().(type) {
			case *atlasv2.AWSCloudProviderContainer:
				containers[i] = *v
			case *atlasv2.AzureCloudProviderContainer:
				containers[i] = *v
			case *atlasv2.GCPCloudProviderContainer:
				containers[i] = *v
			}
		}

		return containers, err
	default:
		return nil, fmt.Errorf("%w: %s", errUnsupportedService, s.service)
	}
}

const maxPerPage = 100

// AzureContainers encapsulates the logic to manage different cloud providers.
func (s *Store) AzureContainers(projectID string) ([]atlasv2.AzureCloudProviderContainer, error) {
	switch s.service {
	case config.CloudService, config.CloudGovService:
		result, _, err := s.clientv2.NetworkPeeringApi.ListPeeringContainerByCloudProvider(s.ctx, projectID).
			PageNum(0).
			ItemsPerPage(maxPerPage).
			ProviderName("Azure").
			Execute()

		containers := make([]atlasv2.AzureCloudProviderContainer, len(result.Results))
		for i, container := range result.Results {
			containers[i] = *container.GetActualInstance().(*atlasv2.AzureCloudProviderContainer)
		}
		return containers, err
	default:
		return nil, fmt.Errorf("%w: %s", errUnsupportedService, s.service)
	}
}

// AWSContainers encapsulates the logic to manage different cloud providers.
func (s *Store) AWSContainers(projectID string) ([]atlasv2.AWSCloudProviderContainer, error) {
	switch s.service {
	case config.CloudService, config.CloudGovService:
		result, _, err := s.clientv2.NetworkPeeringApi.ListPeeringContainerByCloudProvider(s.ctx, projectID).
			PageNum(0).
			ItemsPerPage(maxPerPage).
			ProviderName("AWS").
			Execute()

		containers := make([]atlasv2.AWSCloudProviderContainer, len(result.Results))
		for i, container := range result.Results {
			containers[i] = *container.GetActualInstance().(*atlasv2.AWSCloudProviderContainer)
		}
		return containers, err
	default:
		return nil, fmt.Errorf("%w: %s", errUnsupportedService, s.service)
	}
}

// GCPContainers encapsulates the logic to manage different cloud providers.
func (s *Store) GCPContainers(projectID string) ([]atlasv2.GCPCloudProviderContainer, error) {
	switch s.service {
	case config.CloudService:
		result, _, err := s.clientv2.NetworkPeeringApi.ListPeeringContainerByCloudProvider(s.ctx, projectID).
			PageNum(0).
			ItemsPerPage(maxPerPage).
			ProviderName("GCP").
			Execute()

		containers := make([]atlasv2.GCPCloudProviderContainer, len(result.Results))
		for i, container := range result.Results {
			containers[i] = *container.GetActualInstance().(*atlasv2.GCPCloudProviderContainer)
		}
		return containers, err
	default:
		return nil, fmt.Errorf("%w: %s", errUnsupportedService, s.service)
	}
}

// AllContainers encapsulates the logic to manage different cloud providers.
func (s *Store) AllContainers(projectID string, opts *atlas.ListOptions) ([]interface{}, error) {
	switch s.service {
	case config.CloudService, config.CloudGovService:
		result, _, err := s.clientv2.NetworkPeeringApi.ListPeeringContainers(s.ctx, projectID).
			IncludeCount(opts.IncludeCount).
			PageNum(int32(opts.PageNum)).
			ItemsPerPage(int32(opts.ItemsPerPage)).
			Execute()

		containers := make([]interface{}, len(result.Results))
		for i, container := range result.Results {
			switch v := container.GetActualInstance().(type) {
			case *atlasv2.AWSCloudProviderContainer:
				containers[i] = *v
			case *atlasv2.AzureCloudProviderContainer:
				containers[i] = *v
			case *atlasv2.GCPCloudProviderContainer:
				containers[i] = *v
			}
		}
		return containers, err
	default:
		return nil, fmt.Errorf("%w: %s", errUnsupportedService, s.service)
	}
}

// DeleteContainer encapsulates the logic to manage different cloud providers.
func (s *Store) DeleteContainer(projectID, containerID string) error {
	switch s.service {
	case config.CloudService, config.CloudGovService:
		_, err := s.clientv2.NetworkPeeringApi.DeletePeeringContainer(s.ctx, projectID, containerID).Execute()
		return err
	default:
		return fmt.Errorf("%w: %s", errUnsupportedService, s.service)
	}
}

// Container encapsulates the logic to manage different cloud providers.
func (s *Store) Container(projectID, containerID string) (interface{}, error) {
	switch s.service {
	case config.CloudService, config.CloudGovService:
		result, _, err := s.clientv2.NetworkPeeringApi.GetPeeringContainer(s.ctx, projectID, containerID).Execute()

		switch v := result.GetActualInstance().(type) {
		case *atlasv2.AWSCloudProviderContainer:
			return *v, err
		case *atlasv2.AzureCloudProviderContainer:
			return *v, err
		case *atlasv2.GCPCloudProviderContainer:
			return *v, err
		default:
			return v, err
		}
	default:
		return nil, fmt.Errorf("%w: %s", errUnsupportedService, s.service)
	}
}

// CreateContainer encapsulates the logic to manage different cloud providers.
func (s *Store) CreateContainer(projectID string, container *atlasv2.CloudProviderContainer) (interface{}, error) {
	switch s.service {
	case config.CloudService, config.CloudGovService:
		result, _, err := s.clientv2.NetworkPeeringApi.CreatePeeringContainer(s.ctx, projectID).CloudProviderContainer(*container).Execute()

		switch v := result.GetActualInstance().(type) {
		case *atlasv2.AWSCloudProviderContainer:
			return *v, err
		case *atlasv2.AzureCloudProviderContainer:
			return *v, err
		case *atlasv2.GCPCloudProviderContainer:
			return *v, err
		default:
			return v, err
		}
	default:
		return nil, fmt.Errorf("%w: %s", errUnsupportedService, s.service)
	}
}
