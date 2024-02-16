// Copyright 2024 MongoDB Inc
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
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
	atlasv2 "go.mongodb.org/atlas-sdk/v20231115007/admin"
)

//go:generate mockgen -destination=../mocks/mock_search_nodes.go -package=mocks github.com/mongodb/mongodb-atlas-cli/internal/store SearchNodesLister,SearchNodesCreator,SearchNodesUpdater,SearchNodesDeleter

type SearchNodesLister interface {
	SearchNodes(string, string) (*atlasv2.ApiSearchDeploymentResponse, error)
}

type SearchNodesCreator interface {
	CreateSearchNodes(string, string, *atlasv2.ApiSearchDeploymentRequest) (*atlasv2.ApiSearchDeploymentResponse, error)
	SearchNodes(string, string) (*atlasv2.ApiSearchDeploymentResponse, error)
}

type SearchNodesUpdater interface {
	UpdateSearchNodes(string, string, *atlasv2.ApiSearchDeploymentRequest) (*atlasv2.ApiSearchDeploymentResponse, error)
	SearchNodes(string, string) (*atlasv2.ApiSearchDeploymentResponse, error)
}

type SearchNodesDeleter interface {
	DeleteSearchNodes(string, string) error
	SearchNodes(string, string) (*atlasv2.ApiSearchDeploymentResponse, error)
}

// SearchNodes encapsulate the logic to manage different cloud providers.
func (s *Store) SearchNodes(projectID, clusterName string) (*atlasv2.ApiSearchDeploymentResponse, error) {
	switch s.service {
	case config.CloudService, config.CloudGovService:
		result, _, err := s.clientv2.AtlasSearchApi.GetAtlasSearchDeployment(s.ctx, projectID, clusterName).Execute()
		return result, err
	default:
		return nil, fmt.Errorf("%w: %s", errUnsupportedService, s.service)
	}
}
func (s *Store) CreateSearchNodes(projectID, clusterName string, spec *atlasv2.ApiSearchDeploymentRequest) (*atlasv2.ApiSearchDeploymentResponse, error) {
	switch s.service {
	case config.CloudService, config.CloudGovService:
		result, _, err := s.clientv2.AtlasSearchApi.UpdateAtlasSearchDeployment(s.ctx, projectID, clusterName, spec).Execute()
		return result, err
	default:
		return nil, fmt.Errorf("%w: %s", errUnsupportedService, s.service)
	}
}

func (s *Store) UpdateSearchNodes(projectID, clusterName string, spec *atlasv2.ApiSearchDeploymentRequest) (*atlasv2.ApiSearchDeploymentResponse, error) {
	switch s.service {
	case config.CloudService, config.CloudGovService:
		result, _, err := s.clientv2.AtlasSearchApi.UpdateAtlasSearchDeployment(s.ctx, projectID, clusterName, spec).Execute()
		return result, err
	default:
		return nil, fmt.Errorf("%w: %s", errUnsupportedService, s.service)
	}
}

func (s *Store) DeleteSearchNodes(projectID, clusterName string) error {
	switch s.service {
	case config.CloudService, config.CloudGovService:
		_, err := s.clientv2.AtlasSearchApi.DeleteAtlasSearchDeployment(s.ctx, projectID, clusterName).Execute()
		return err
	default:
		return fmt.Errorf("%w: %s", errUnsupportedService, s.service)
	}
}
