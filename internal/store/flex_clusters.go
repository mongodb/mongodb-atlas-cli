// Copyright 2024 MongoDB Inc
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
	atlasv2 "go.mongodb.org/atlas-sdk/v20250312005/admin"
)

// ListFlexClusters encapsulate the logic to manage different cloud providers.
func (s *Store) ListFlexClusters(opts *atlasv2.ListFlexClustersApiParams) (*atlasv2.PaginatedFlexClusters20241113, error) {
	if s.service == config.CloudGovService {
		return nil, fmt.Errorf("%w: %s", errUnsupportedService, s.service)
	}

	result, _, err := s.clientv2.FlexClustersApi.ListFlexClustersWithParams(s.ctx, opts).Execute()
	return result, err
}

// FlexCluster encapsulate the logic to manage different cloud providers.
func (s *Store) FlexCluster(groupID, name string) (*atlasv2.FlexClusterDescription20241113, error) {
	if s.service == config.CloudGovService {
		return nil, fmt.Errorf("%w: %s", errUnsupportedService, s.service)
	}

	result, _, err := s.clientv2.FlexClustersApi.GetFlexCluster(s.ctx, groupID, name).Execute()
	return result, err
}

// CreateFlexCluster encapsulate the logic to manage different cloud providers.
func (s *Store) CreateFlexCluster(groupID string, flexClusterDescriptionCreate20241113 *atlasv2.FlexClusterDescriptionCreate20241113) (*atlasv2.FlexClusterDescription20241113, error) {
	if s.service == config.CloudGovService {
		return nil, fmt.Errorf("%w: %s", errUnsupportedService, s.service)
	}

	result, _, err := s.clientv2.FlexClustersApi.CreateFlexCluster(s.ctx, groupID, flexClusterDescriptionCreate20241113).Execute()
	return result, err
}

// UpdateFlexCluster encapsulate the logic to manage different cloud providers.
func (s *Store) UpdateFlexCluster(groupID, name string, flexClusterDescriptionUpdate20241113 *atlasv2.FlexClusterDescriptionUpdate20241113) (*atlasv2.FlexClusterDescription20241113, error) {
	if s.service == config.CloudGovService {
		return nil, fmt.Errorf("%w: %s", errUnsupportedService, s.service)
	}

	result, _, err := s.clientv2.FlexClustersApi.UpdateFlexCluster(s.ctx, groupID, name, flexClusterDescriptionUpdate20241113).Execute()
	return result, err
}

// UpgradeFlexCluster encapsulate the logic to manage different cloud providers.
func (s *Store) UpgradeFlexCluster(groupID string, flexClusterDescriptionUpdate20241113 *atlasv2.AtlasTenantClusterUpgradeRequest20240805) (*atlasv2.FlexClusterDescription20241113, error) {
	if s.service == config.CloudGovService {
		return nil, fmt.Errorf("%w: %s", errUnsupportedService, s.service)
	}

	result, _, err := s.clientv2.FlexClustersApi.UpgradeFlexCluster(s.ctx, groupID, flexClusterDescriptionUpdate20241113).Execute()
	return result, err
}

// DeleteFlexCluster encapsulate the logic to manage different cloud providers.
func (s *Store) DeleteFlexCluster(groupID, name string) error {
	if s.service == config.CloudGovService {
		return fmt.Errorf("%w: %s", errUnsupportedService, s.service)
	}

	_, err := s.clientv2.FlexClustersApi.DeleteFlexCluster(s.ctx, groupID, name).Execute()
	return err
}
