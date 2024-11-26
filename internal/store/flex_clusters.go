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

	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/config"
	atlas "go.mongodb.org/atlas-sdk/v20241113001/admin"
)

//go:generate mockgen -destination=../mocks/mock_flex_clusters.go -package=mocks github.com/mongodb/mongodb-atlas-cli/atlascli/internal/store FlexClusterLister,FlexClusterDescriber,FlexClusterCreator,FlexClusterUpdater,FlexClusterUpgrader

type FlexClusterLister interface {
	ListFlexClusters(*atlas.ListFlexClustersApiParams) (*atlas.PaginatedFlexClusters20241113, error)
}

type FlexClusterDescriber interface {
	FlexCluster(string, string) (*atlas.FlexClusterDescription20241113, error)
}

type FlexClusterCreator interface {
	CreateCluster(string, *atlas.FlexClusterDescriptionCreate20241113) (*atlas.FlexClusterDescription20241113, error)
}

type FlexClusterUpdater interface {
	CreateCluster(string, string, *atlas.FlexClusterDescriptionUpdate20241113) (*atlas.FlexClusterDescription20241113, error)
}

type FlexClusterUpgrader interface {
	UpgradeFlexCluster(string, *atlas.FlexClusterDescriptionUpdate20241113) (*atlas.FlexClusterDescription20241113, error)
}

// ListFlexClusters encapsulate the logic to manage different cloud providers.
func (s *Store) ListFlexClusters(opts *atlas.ListFlexClustersApiParams) (*atlas.PaginatedFlexClusters20241113, error) {
	if s.service == config.CloudGovService {
		return nil, fmt.Errorf("%w: %s", errUnsupportedService, s.service)
	}

	result, _, err := s.clientv2.FlexClustersApi.ListFlexClustersWithParams(s.ctx, opts).Execute()
	return result, err
}

// FlexCluster encapsulate the logic to manage different cloud providers.
func (s *Store) FlexCluster(groupId, name string) (*atlas.FlexClusterDescription20241113, error) {
	if s.service == config.CloudGovService {
		return nil, fmt.Errorf("%w: %s", errUnsupportedService, s.service)
	}

	result, _, err := s.clientv2.FlexClustersApi.GetFlexCluster(s.ctx, groupId, name).Execute()
	return result, err
}

// CreateFlexCluster encapsulate the logic to manage different cloud providers.
func (s *Store) CreateFlexCluster(groupId string, flexClusterDescriptionCreate20241113 *atlas.FlexClusterDescriptionCreate20241113) (*atlas.FlexClusterDescription20241113, error) {
	if s.service == config.CloudGovService {
		return nil, fmt.Errorf("%w: %s", errUnsupportedService, s.service)
	}

	result, _, err := s.clientv2.FlexClustersApi.CreateFlexCluster(s.ctx, groupId, flexClusterDescriptionCreate20241113).Execute()
	return result, err
}

// UpdateFlexCluster encapsulate the logic to manage different cloud providers.
func (s *Store) UpdateFlexCluster(groupId, name string, flexClusterDescriptionUpdate20241113 *atlas.FlexClusterDescriptionUpdate20241113) (*atlas.FlexClusterDescription20241113, error) {
	if s.service == config.CloudGovService {
		return nil, fmt.Errorf("%w: %s", errUnsupportedService, s.service)
	}

	result, _, err := s.clientv2.FlexClustersApi.UpdateFlexCluster(s.ctx, groupId, name, flexClusterDescriptionUpdate20241113).Execute()
	return result, err
}

// UpgradeFlexCluster encapsulate the logic to manage different cloud providers.
func (s *Store) UpgradeFlexCluster(groupId string, flexClusterDescriptionUpdate20241113 *atlas.FlexClusterDescription20241113) (*atlas.FlexClusterDescription20241113, error) {
	if s.service == config.CloudGovService {
		return nil, fmt.Errorf("%w: %s", errUnsupportedService, s.service)
	}

	result, _, err := s.clientv2.FlexClustersApi.UpgradeFlexCluster(s.ctx, groupId, flexClusterDescriptionUpdate20241113).Execute()
	return result, err
}
