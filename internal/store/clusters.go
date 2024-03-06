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
	"go.mongodb.org/ops-manager/opsmngr"
)

//go:generate mockgen -destination=../mocks/mock_clusters.go -package=mocks github.com/mongodb/mongodb-atlas-cli/internal/store ClusterLister,OpsManagerClusterDescriber,AtlasSharedClusterDescriber

type ClusterLister interface {
	ProjectClusters(string, *atlas.ListOptions) (interface{}, error)
}

type OpsManagerClusterDescriber interface {
	OpsManagerCluster(string, string) (*opsmngr.Cluster, error)
}

type AtlasSharedClusterDescriber interface {
	AtlasSharedCluster(string, string) (*atlas.Cluster, error)
}

// AtlasSharedCluster encapsulates the logic to fetch details of one shared cluster.
func (s *Store) AtlasSharedCluster(projectID, name string) (*atlas.Cluster, error) {
	switch s.service {
	case config.OpsManagerService, config.CloudManagerService:
		result, _, err := s.client.(*atlas.Client).Clusters.Get(s.ctx, projectID, name)
		return result, err
	default:
		return nil, fmt.Errorf("%w: %s", errUnsupportedService, s.service)
	}
}

// ProjectClusters encapsulate the logic to manage different cloud providers.
func (s *Store) ProjectClusters(projectID string, opts *atlas.ListOptions) (interface{}, error) {
	switch s.service {
	case config.OpsManagerService, config.CloudManagerService:
		result, _, err := s.client.(*opsmngr.Client).Clusters.List(s.ctx, projectID, opts)
		return result, err
	default:
		return nil, fmt.Errorf("%w: %s", errUnsupportedService, s.service)
	}
}

// OpsManagerCluster encapsulates the logic to manage different cloud providers.
func (s *Store) OpsManagerCluster(projectID, name string) (*opsmngr.Cluster, error) {
	switch s.service {
	case config.OpsManagerService, config.CloudManagerService:
		result, _, err := s.client.(*opsmngr.Client).Clusters.Get(s.ctx, projectID, name)
		return result, err
	default:
		return nil, fmt.Errorf("%w: %s", errUnsupportedService, s.service)
	}
}

// ListAllProjectClusters encapsulate the logic to manage different cloud providers.
func (s *Store) ListAllProjectClusters() (*opsmngr.AllClustersProjects, error) {
	switch s.service {
	case config.OpsManagerService, config.CloudManagerService:
		result, _, err := s.client.(*opsmngr.Client).Clusters.ListAll(s.ctx)
		return result, err
	default:
		return nil, fmt.Errorf("%w: %s", errUnsupportedService, s.service)
	}
}
