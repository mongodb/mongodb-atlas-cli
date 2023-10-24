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

package atlas

import (
	"fmt"

	"github.com/mongodb/mongodb-atlas-cli/internal/config"
	"go.mongodb.org/atlas-sdk/v20231001002/admin"
)

//go:generate mockgen -destination=../../mocks/atlas/mock_clusters.go -package=atlas github.com/mongodb/mongodb-atlas-cli/internal/store/atlas ClusterLister,ClusterDescriber,ClusterConfigurationOptionsDescriber

type ClusterLister interface {
	ProjectClusters(string, *ListOptions) (interface{}, error)
}

type ClusterDescriber interface {
	AtlasCluster(string, string) (*admin.AdvancedClusterDescription, error)
}

type ClusterConfigurationOptionsDescriber interface {
	AtlasClusterConfigurationOptions(string, string) (*admin.ClusterDescriptionProcessArgs, error)
}

// ProjectClusters encapsulate the logic to manage different cloud providers.
func (s *Store) ProjectClusters(projectID string, opts *ListOptions) (interface{}, error) {
	switch s.service {
	case config.CloudService, config.CloudGovService:
		res := s.clientv2.ClustersApi.ListClusters(s.ctx, projectID)
		if opts != nil {
			res = res.PageNum(opts.PageNum).ItemsPerPage(opts.ItemsPerPage)
		}
		result, _, err := res.Execute()
		return result, err
	default:
		return nil, fmt.Errorf("%w: %s", errUnsupportedService, s.service)
	}
}

// AtlasCluster encapsulates the logic to manage different cloud providers.
func (s *Store) AtlasCluster(projectID, name string) (*admin.AdvancedClusterDescription, error) {
	switch s.service {
	case config.CloudService, config.CloudGovService:
		result, _, err := s.clientv2.ClustersApi.GetCluster(s.ctx, projectID, name).Execute()
		return result, err
	default:
		return nil, fmt.Errorf("%w: %s", errUnsupportedService, s.service)
	}
}

// AtlasClusterConfigurationOptions encapsulates the logic to manage different cloud providers.
func (s *Store) AtlasClusterConfigurationOptions(projectID, name string) (*admin.ClusterDescriptionProcessArgs, error) {
	switch s.service {
	case config.CloudService, config.CloudGovService:
		result, _, err := s.clientv2.ClustersApi.GetClusterAdvancedConfiguration(s.ctx, projectID, name).Execute()
		return result, err
	default:
		return nil, fmt.Errorf("%w: %s", errUnsupportedService, s.service)
	}
}
