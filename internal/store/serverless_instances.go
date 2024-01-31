// Copyright 2021 MongoDB Inc
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

//go:generate mockgen -destination=../mocks/mock_serverless_instances.go -package=mocks github.com/mongodb/mongodb-atlas-cli/internal/store ServerlessInstanceLister,ServerlessInstanceDescriber,ServerlessInstanceDeleter,ServerlessInstanceCreator,ServerlessInstanceUpdater

type ServerlessInstanceLister interface {
	ServerlessInstances(string, *atlas.ListOptions) (*atlasv2.PaginatedServerlessInstanceDescription, error)
}

type ServerlessInstanceDescriber interface {
	ServerlessInstance(string, string) (*atlas.Cluster, error)
	GetServerlessInstance(string, string) (*atlasv2.ServerlessInstanceDescription, error)
}

type ServerlessInstanceDeleter interface {
	DeleteServerlessInstance(string, string) error
}

type ServerlessInstanceCreator interface {
	CreateServerlessInstance(string, *atlasv2.ServerlessInstanceDescriptionCreate) (*atlasv2.ServerlessInstanceDescription, error)
}

type ServerlessInstanceUpdater interface {
	UpdateServerlessInstance(string, string, *atlasv2.ServerlessInstanceDescriptionUpdate) (*atlasv2.ServerlessInstanceDescription, error)
}

// ServerlessInstances encapsulates the logic to manage different cloud providers.
func (s *Store) ServerlessInstances(projectID string, listOps *atlas.ListOptions) (*atlasv2.PaginatedServerlessInstanceDescription, error) {
	switch s.service {
	case config.CloudService:
		result, _, err := s.clientv2.ServerlessInstancesApi.ListServerlessInstances(s.ctx, projectID).
			ItemsPerPage(listOps.ItemsPerPage).
			PageNum(listOps.PageNum).
			Execute()

		return result, err
	default:
		return nil, fmt.Errorf("%w: %s", errUnsupportedService, s.service)
	}
}

// ServerlessInstance encapsulates the logic to manage different cloud providers.
func (s *Store) GetServerlessInstance(projectID, clusterName string) (*atlasv2.ServerlessInstanceDescription, error) {
	switch s.service {
	case config.CloudService:
		result, _, err := s.clientv2.ServerlessInstancesApi.GetServerlessInstance(s.ctx, projectID, clusterName).Execute()
		return result, err
	default:
		return nil, fmt.Errorf("%w: %s", errUnsupportedService, s.service)
	}
}

// Used by Kubernetes v1 ServerlessInstance encapsulates the logic to manage different cloud providers.
func (s *Store) ServerlessInstance(projectID, clusterName string) (*atlas.Cluster, error) {
	switch s.service {
	case config.CloudService:
		result, _, err := s.client.(*atlas.Client).ServerlessInstances.Get(s.ctx, projectID, clusterName)
		return result, err
	default:
		return nil, fmt.Errorf("%w: %s", errUnsupportedService, s.service)
	}
}

// DeleteServerlessInstance encapsulate the logic to manage different cloud providers.
func (s *Store) DeleteServerlessInstance(projectID, name string) error {
	switch s.service {
	case config.CloudService:
		_, _, err := s.clientv2.ServerlessInstancesApi.DeleteServerlessInstance(s.ctx, projectID, name).Execute()
		return err
	default:
		return fmt.Errorf("%w: %s", errUnsupportedService, s.service)
	}
}

// CreateServerlessInstance encapsulate the logic to manage different cloud providers.
func (s *Store) CreateServerlessInstance(projectID string, cluster *atlasv2.ServerlessInstanceDescriptionCreate) (*atlasv2.ServerlessInstanceDescription, error) {
	switch s.service {
	case config.CloudService:
		result, _, err := s.clientv2.ServerlessInstancesApi.CreateServerlessInstance(s.ctx, projectID, cluster).
			Execute()
		return result, err
	default:
		return nil, fmt.Errorf("%w: %s", errUnsupportedService, s.service)
	}
}

// UpdateServerlessInstance encapsulate the logic to manage different cloud providers.
func (s *Store) UpdateServerlessInstance(projectID string, instanceName string, req *atlasv2.ServerlessInstanceDescriptionUpdate) (*atlasv2.ServerlessInstanceDescription, error) {
	switch s.service {
	case config.CloudService:
		result, _, err := s.clientv2.ServerlessInstancesApi.UpdateServerlessInstance(s.ctx, projectID, instanceName, req).
			Execute()

		return result, err
	default:
		return nil, fmt.Errorf("%w: %s", errUnsupportedService, s.service)
	}
}
