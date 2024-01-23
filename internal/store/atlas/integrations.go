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
	atlasv2 "go.mongodb.org/atlas-sdk/v20231115004/admin"
)

//go:generate mockgen -destination=../../mocks/atlas/mock_integrations.go -package=atlas github.com/mongodb/mongodb-atlas-cli/internal/store/atlas IntegrationLister

type IntegrationLister interface {
	Integrations(string) (*atlasv2.PaginatedIntegration, error)
}

// Integrations encapsulates the logic to manage different cloud providers.
func (s *Store) Integrations(projectID string) (*atlasv2.PaginatedIntegration, error) {
	switch s.service {
	case config.CloudService, config.CloudGovService:
		resp, _, err := s.clientv2.ThirdPartyIntegrationsApi.ListThirdPartyIntegrations(s.ctx, projectID).Execute()
		return resp, err
	default:
		return nil, fmt.Errorf("%w: %s", errUnsupportedService, s.service)
	}
}
