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
	atlasv2 "go.mongodb.org/atlas-sdk/v20231115005/admin"
	atlas "go.mongodb.org/atlas/mongodbatlas"
)

//go:generate mockgen -destination=../mocks/mock_peering_connections.go -package=mocks github.com/mongodb/mongodb-atlas-cli/internal/store PeeringConnectionLister,PeeringConnectionDescriber,PeeringConnectionDeleter,AzurePeeringConnectionCreator,AWSPeeringConnectionCreator,GCPPeeringConnectionCreator,PeeringConnectionCreator,ContainersLister,ContainersDeleter

type PeeringConnectionLister interface {
	PeeringConnections(string, *atlas.ContainersListOptions) ([]atlasv2.BaseNetworkPeeringConnectionSettings, error)
}

type PeeringConnectionDescriber interface {
	PeeringConnection(string, string) (*atlasv2.BaseNetworkPeeringConnectionSettings, error)
}

type PeeringConnectionCreator interface {
	CreateContainer(string, *atlasv2.CloudProviderContainer) (*atlasv2.CloudProviderContainer, error)
	CreatePeeringConnection(string, *atlasv2.BaseNetworkPeeringConnectionSettings) (*atlasv2.BaseNetworkPeeringConnectionSettings, error)
}

type AzurePeeringConnectionCreator interface {
	AzureContainers(string) ([]atlasv2.CloudProviderContainer, error)
	PeeringConnectionCreator
}

type AWSPeeringConnectionCreator interface {
	AWSContainers(string) ([]atlasv2.CloudProviderContainer, error)
	PeeringConnectionCreator
}

type GCPPeeringConnectionCreator interface {
	GCPContainers(string) ([]atlasv2.CloudProviderContainer, error)
	PeeringConnectionCreator
}

type PeeringConnectionDeleter interface {
	DeletePeeringConnection(string, string) error
}

type ContainersLister interface {
	ContainersByProvider(string, *atlas.ContainersListOptions) ([]atlasv2.CloudProviderContainer, error)
	AllContainers(string, *atlas.ListOptions) ([]atlasv2.CloudProviderContainer, error)
}

type ContainersDeleter interface {
	DeleteContainer(string, string) error
}

// PeeringConnections encapsulates the logic to manage different cloud providers.
func (s *Store) PeeringConnections(projectID string, opts *atlas.ContainersListOptions) ([]atlasv2.BaseNetworkPeeringConnectionSettings, error) {
	switch s.service {
	case config.CloudService, config.CloudGovService:
		result, _, err := s.clientv2.NetworkPeeringApi.ListPeeringConnections(s.ctx, projectID).
			ItemsPerPage(opts.ItemsPerPage).
			PageNum(opts.PageNum).
			ProviderName(opts.ProviderName).Execute()
		if err != nil {
			return nil, err
		}
		return result.GetResults(), nil
	default:
		return nil, fmt.Errorf("%w: %s", errUnsupportedService, s.service)
	}
}

// PeeringConnections encapsulates the logic to manage different cloud providers.
func (s *Store) PeeringConnection(projectID, peerID string) (*atlasv2.BaseNetworkPeeringConnectionSettings, error) {
	switch s.service {
	case config.CloudService, config.CloudGovService:
		result, _, err := s.clientv2.NetworkPeeringApi.GetPeeringConnection(s.ctx, projectID, peerID).Execute()
		return result, err
	default:
		return nil, fmt.Errorf("%w: %s", errUnsupportedService, s.service)
	}
}

// DeletePrivateEndpoint encapsulates the logic to manage different cloud providers.
func (s *Store) DeletePeeringConnection(projectID, peerID string) error {
	switch s.service {
	case config.CloudService, config.CloudGovService:
		_, _, err := s.clientv2.NetworkPeeringApi.DeletePeeringConnection(s.ctx, projectID, peerID).Execute()
		return err
	default:
		return fmt.Errorf("%w: %s", errUnsupportedService, s.service)
	}
}

// CreatePeeringConnection encapsulates the logic to manage different cloud providers.
func (s *Store) CreatePeeringConnection(projectID string, peer *atlasv2.BaseNetworkPeeringConnectionSettings) (*atlasv2.BaseNetworkPeeringConnectionSettings, error) {
	switch s.service {
	case config.CloudService, config.CloudGovService:
		result, _, err := s.clientv2.NetworkPeeringApi.CreatePeeringConnection(s.ctx, projectID, peer).Execute()
		return result, err
	default:
		return nil, fmt.Errorf("%w: %s", errUnsupportedService, s.service)
	}
}

// ContainersByProvider encapsulates the logic to manage different cloud providers.
func (s *Store) ContainersByProvider(projectID string, opts *atlas.ContainersListOptions) ([]atlasv2.CloudProviderContainer, error) {
	switch s.service {
	case config.CloudService, config.CloudGovService:
		res := s.clientv2.NetworkPeeringApi.ListPeeringContainerByCloudProvider(s.ctx, projectID)
		if opts != nil {
			res = res.PageNum(opts.PageNum).ItemsPerPage(opts.ItemsPerPage).ProviderName(opts.ProviderName)
		}
		result, _, err := res.Execute()
		if err != nil {
			return nil, err
		}
		return result.GetResults(), nil
	default:
		return nil, fmt.Errorf("%w: %s", errUnsupportedService, s.service)
	}
}

const maxPerPage = 100

// AzureContainers encapsulates the logic to manage different cloud providers.
func (s *Store) AzureContainers(projectID string) ([]atlasv2.CloudProviderContainer, error) {
	switch s.service {
	case config.CloudService, config.CloudGovService:
		result, _, err := s.clientv2.NetworkPeeringApi.ListPeeringContainerByCloudProvider(s.ctx, projectID).
			PageNum(0).
			ItemsPerPage(maxPerPage).
			ProviderName("Azure").
			Execute()
		if err != nil {
			return nil, err
		}
		return result.GetResults(), nil
	default:
		return nil, fmt.Errorf("%w: %s", errUnsupportedService, s.service)
	}
}

// AWSContainers encapsulates the logic to manage different cloud providers.
func (s *Store) AWSContainers(projectID string) ([]atlasv2.CloudProviderContainer, error) {
	switch s.service {
	case config.CloudService, config.CloudGovService:
		result, _, err := s.clientv2.NetworkPeeringApi.ListPeeringContainerByCloudProvider(s.ctx, projectID).
			PageNum(0).
			ItemsPerPage(maxPerPage).
			ProviderName("AWS").
			Execute()

		if err != nil {
			return nil, err
		}
		return result.GetResults(), nil
	default:
		return nil, fmt.Errorf("%w: %s", errUnsupportedService, s.service)
	}
}

// GCPContainers encapsulates the logic to manage different cloud providers.
func (s *Store) GCPContainers(projectID string) ([]atlasv2.CloudProviderContainer, error) {
	switch s.service {
	case config.CloudService:
		result, _, err := s.clientv2.NetworkPeeringApi.ListPeeringContainerByCloudProvider(s.ctx, projectID).
			PageNum(0).
			ItemsPerPage(maxPerPage).
			ProviderName("GCP").
			Execute()
		if err != nil {
			return nil, err
		}
		return result.GetResults(), nil
	default:
		return nil, fmt.Errorf("%w: %s", errUnsupportedService, s.service)
	}
}

// AllContainers encapsulates the logic to manage different cloud providers.
func (s *Store) AllContainers(projectID string, opts *atlas.ListOptions) ([]atlasv2.CloudProviderContainer, error) {
	switch s.service {
	case config.CloudService, config.CloudGovService:
		res := s.clientv2.NetworkPeeringApi.ListPeeringContainers(s.ctx, projectID)
		if opts != nil {
			res = res.PageNum(opts.PageNum).ItemsPerPage(opts.ItemsPerPage)
		}
		result, _, err := res.Execute()
		if err != nil {
			return nil, err
		}
		return result.GetResults(), nil
	default:
		return nil, fmt.Errorf("%w: %s", errUnsupportedService, s.service)
	}
}

// DeleteContainer encapsulates the logic to manage different cloud providers.
func (s *Store) DeleteContainer(projectID, containerID string) error {
	switch s.service {
	case config.CloudService, config.CloudGovService:
		_, _, err := s.clientv2.NetworkPeeringApi.DeletePeeringContainer(s.ctx, projectID, containerID).Execute()
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
		return result, err
	default:
		return nil, fmt.Errorf("%w: %s", errUnsupportedService, s.service)
	}
}

// CreateContainer encapsulates the logic to manage different cloud providers.
func (s *Store) CreateContainer(projectID string, container *atlasv2.CloudProviderContainer) (*atlasv2.CloudProviderContainer, error) {
	switch s.service {
	case config.CloudService, config.CloudGovService:
		result, _, err := s.clientv2.NetworkPeeringApi.CreatePeeringContainer(s.ctx, projectID, container).Execute()
		return result, err
	default:
		return nil, fmt.Errorf("%w: %s", errUnsupportedService, s.service)
	}
}
