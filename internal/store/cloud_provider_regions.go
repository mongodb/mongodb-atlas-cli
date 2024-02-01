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

	"github.com/andreangiolillo/mongocli-test/internal/config"
	atlasv2 "go.mongodb.org/atlas-sdk/v20231115002/admin"
)

//go:generate mockgen -destination=../mocks/mock_cloud_provider_regions.go -package=mocks github.com/andreangiolillo/mongocli-test/internal/store CloudProviderRegionsLister

type CloudProviderRegionsLister interface {
	CloudProviderRegions(string, string, []string) (*atlasv2.PaginatedApiAtlasProviderRegions, error)
}

// CloudProviderRegions encapsulates the logic to manage different cloud providers.
func (s *Store) CloudProviderRegions(projectID, tier string, providerName []string) (*atlasv2.PaginatedApiAtlasProviderRegions, error) {
	switch s.service {
	case config.CloudService, config.CloudGovService:
		result, _, err := s.clientv2.ClustersApi.ListCloudProviderRegions(s.ctx, projectID).Providers(providerName).Tier(tier).Execute()
		return result, err
	default:
		return nil, fmt.Errorf("%w: %s", errUnsupportedService, s.service)
	}
}
