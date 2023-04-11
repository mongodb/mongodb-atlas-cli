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

//go:generate mockgen -destination=../mocks/mock_peering_connections.go -package=mocks github.com/mongodb/mongodb-atlas-cli/internal/store PeeringConnectionLister,PeeringConnectionDescriber,PeeringConnectionDeleter,AzurePeeringConnectionCreator,AWSPeeringConnectionCreator,GCPPeeringConnectionCreator,PeeringConnectionCreator

type PeeringConnectionLister interface {
	PeeringConnections(string, *atlas.ContainersListOptions) ([]interface{}, error)
}

type PeeringConnectionDescriber interface {
	PeeringConnection(string, string) (interface{}, error)
}

type PeeringConnectionCreator interface {
	CreateContainer(string, *atlasv2.CloudProviderContainer) (interface{}, error)
	CreatePeeringConnection(string, *atlasv2.ContainerPeerViewRequest) (interface{}, error)
}

type AzurePeeringConnectionCreator interface {
	AzureContainers(string) ([]*atlasv2.AzureCloudProviderContainer, error)
	PeeringConnectionCreator
}

type AWSPeeringConnectionCreator interface {
	AWSContainers(string) ([]*atlasv2.AWSCloudProviderContainer, error)
	PeeringConnectionCreator
}

type GCPPeeringConnectionCreator interface {
	GCPContainers(string) ([]*atlasv2.GCPCloudProviderContainer, error)
	PeeringConnectionCreator
}

type PeeringConnectionDeleter interface {
	DeletePeeringConnection(string, string) error
}

// PeeringConnections encapsulates the logic to manage different cloud providers.
func (s *Store) PeeringConnections(projectID string, opts *atlas.ContainersListOptions) ([]interface{}, error) {
	switch s.service {
	case config.CloudService, config.CloudGovService:
		result, _, err := s.clientv2.NetworkPeeringApi.ListPeeringConnections(s.ctx, projectID).
			IncludeCount(opts.IncludeCount).
			ItemsPerPage(int32(opts.ItemsPerPage)).
			PageNum(int32(opts.PageNum)).
			ProviderName(opts.ProviderName).Execute()

		var connections []interface{}
		switch v := result.GetActualInstance().(type) {
		case *atlasv2.PaginatedAWSPeerVpc:
			for _, connection := range v.Results {
				connections = append(connections, connection)
			}
		case *atlasv2.PaginatedAzurePeerNetwork:
			for _, connection := range v.Results {
				connections = append(connections, connection)
			}
		case *atlasv2.PaginatedGCPPeerVpc:
			for _, connection := range v.Results {
				connections = append(connections, connection)
			}
		}

		return connections, err
	default:
		return nil, fmt.Errorf("%w: %s", errUnsupportedService, s.service)
	}
}

// PeeringConnections encapsulates the logic to manage different cloud providers.
func (s *Store) PeeringConnection(projectID, peerID string) (interface{}, error) {
	switch s.service {
	case config.CloudService, config.CloudGovService:
		result, _, err := s.clientv2.NetworkPeeringApi.GetPeeringConnection(s.ctx, projectID, peerID).Execute()
		return result.GetActualInstance(), err
	default:
		return nil, fmt.Errorf("%w: %s", errUnsupportedService, s.service)
	}
}

// DeletePrivateEndpoint encapsulates the logic to manage different cloud providers.
func (s *Store) DeletePeeringConnection(projectID, peerID string) error {
	switch s.service {
	case config.CloudService, config.CloudGovService:
		_, err := s.clientv2.NetworkPeeringApi.DeletePeeringConnection(s.ctx, projectID, peerID).Execute()
		return err
	default:
		return fmt.Errorf("%w: %s", errUnsupportedService, s.service)
	}
}

// CreatePeeringConnection encapsulates the logic to manage different cloud providers.
func (s *Store) CreatePeeringConnection(projectID string, peer *atlasv2.ContainerPeerViewRequest) (interface{}, error) {
	switch s.service {
	case config.CloudService, config.CloudGovService:
		result, _, err := s.clientv2.NetworkPeeringApi.CreatePeeringConnection(s.ctx, projectID).ContainerPeerViewRequest(*peer).Execute()
		return result.GetActualInstance(), err
	default:
		return nil, fmt.Errorf("%w: %s", errUnsupportedService, s.service)
	}
}
