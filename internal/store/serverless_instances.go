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

	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/config"
	atlasv2 "go.mongodb.org/atlas-sdk/v20240530002/admin"
)

//go:generate mockgen -destination=../mocks/mock_serverless_instances.go -package=mocks github.com/mongodb/mongodb-atlas-cli/atlascli/internal/store ServerlessInstanceLister,ServerlessInstanceDescriber,ServerlessInstanceDeleter,ServerlessInstanceCreator,ServerlessInstanceUpdater

type ServerlessInstanceLister interface {
	ServerlessInstances(string, *ListOptions) (*atlasv2.PaginatedServerlessInstanceDescription, error)
}

type ServerlessInstanceDescriber interface {
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
func (s *Store) ServerlessInstances(projectID string, listOps *ListOptions) (*atlasv2.PaginatedServerlessInstanceDescription, error) {
	if s.service == config.CloudGovService {
		return nil, fmt.Errorf("%w: %s", errUnsupportedService, s.service)
	}
	result, _, err := s.clientv2.ServerlessInstancesApi.ListServerlessInstances(s.ctx, projectID).
		ItemsPerPage(listOps.ItemsPerPage).
		PageNum(listOps.PageNum).
		IncludeCount(listOps.IncludeCount).
		Execute()

	return result, err
}

// GetServerlessInstance encapsulates the logic to manage different cloud providers.
func (s *Store) GetServerlessInstance(projectID, clusterName string) (*atlasv2.ServerlessInstanceDescription, error) {
	if s.service == config.CloudGovService {
		return nil, fmt.Errorf("%w: %s", errUnsupportedService, s.service)
	}
	result, _, err := s.clientv2.ServerlessInstancesApi.GetServerlessInstance(s.ctx, projectID, clusterName).Execute()
	return result, err
}

// DeleteServerlessInstance encapsulate the logic to manage different cloud providers.
func (s *Store) DeleteServerlessInstance(projectID, name string) error {
	if s.service == config.CloudGovService {
		return fmt.Errorf("%w: %s", errUnsupportedService, s.service)
	}
	_, _, err := s.clientv2.ServerlessInstancesApi.DeleteServerlessInstance(s.ctx, projectID, name).Execute()
	return err
}

// CreateServerlessInstance encapsulate the logic to manage different cloud providers.
func (s *Store) CreateServerlessInstance(projectID string, cluster *atlasv2.ServerlessInstanceDescriptionCreate) (*atlasv2.ServerlessInstanceDescription, error) {
	if s.service == config.CloudGovService {
		return nil, fmt.Errorf("%w: %s", errUnsupportedService, s.service)
	}
	result, _, err := s.clientv2.ServerlessInstancesApi.CreateServerlessInstance(s.ctx, projectID, cluster).
		Execute()
	return result, err
}

// UpdateServerlessInstance encapsulate the logic to manage different cloud providers.
func (s *Store) UpdateServerlessInstance(projectID string, instanceName string, req *atlasv2.ServerlessInstanceDescriptionUpdate) (*atlasv2.ServerlessInstanceDescription, error) {
	if s.service == config.CloudGovService {
		return nil, fmt.Errorf("%w: %s", errUnsupportedService, s.service)
	}
	result, _, err := s.clientv2.ServerlessInstancesApi.UpdateServerlessInstance(s.ctx, projectID, instanceName, req).
		Execute()

	return result, err
}
