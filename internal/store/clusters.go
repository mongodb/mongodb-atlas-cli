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

	"github.com/mongodb/mongodb-atlas-cli/mongocli/v2/internal/config"
	"go.mongodb.org/ops-manager/opsmngr"
)

//go:generate mockgen -destination=../mocks/mock_clusters.go -package=mocks github.com/mongodb/mongodb-atlas-cli/mongocli/v2/internal/store ClusterLister,OpsManagerClusterDescriber

type ClusterLister interface {
	ProjectClusters(string, *opsmngr.ListOptions) (*opsmngr.Clusters, error)
}

type OpsManagerClusterDescriber interface {
	OpsManagerCluster(string, string) (*opsmngr.Cluster, error)
}

// ProjectClusters encapsulate the logic to manage different cloud providers.
func (s *Store) ProjectClusters(projectID string, opts *opsmngr.ListOptions) (*opsmngr.Clusters, error) {
	switch s.service {
	case config.OpsManagerService, config.CloudManagerService:
		result, _, err := s.client.Clusters.List(s.ctx, projectID, opts)
		return result, err
	default:
		return nil, fmt.Errorf("%w: %s", errUnsupportedService, s.service)
	}
}

// OpsManagerCluster encapsulates the logic to manage different cloud providers.
func (s *Store) OpsManagerCluster(projectID, name string) (*opsmngr.Cluster, error) {
	switch s.service {
	case config.OpsManagerService, config.CloudManagerService:
		result, _, err := s.client.Clusters.Get(s.ctx, projectID, name)
		return result, err
	default:
		return nil, fmt.Errorf("%w: %s", errUnsupportedService, s.service)
	}
}

// ListAllProjectClusters encapsulate the logic to manage different cloud providers.
func (s *Store) ListAllProjectClusters() (*opsmngr.AllClustersProjects, error) {
	switch s.service {
	case config.OpsManagerService, config.CloudManagerService:
		result, _, err := s.client.Clusters.ListAll(s.ctx)
		return result, err
	default:
		return nil, fmt.Errorf("%w: %s", errUnsupportedService, s.service)
	}
}
