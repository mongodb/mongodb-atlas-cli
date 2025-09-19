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

	"github.com/mongodb/atlas-cli-core/config"
	atlasv2 "go.mongodb.org/atlas-sdk/v20250312007/admin"
)

// CreateLinkToken encapsulate the logic to manage different cloud providers.
func (s *Store) CreateLinkToken(orgID string, linkToken *atlasv2.TargetOrgRequest) (*atlasv2.TargetOrg, error) {
	if s.service == config.CloudGovService {
		return nil, fmt.Errorf("%w: %s", errUnsupportedService, s.service)
	}
	result, _, err := s.clientv2.CloudMigrationServiceApi.CreateLinkToken(s.ctx, orgID, linkToken).Execute()
	return result, err
}

// DeleteLinkToken encapsulate the logic to manage different cloud providers.
func (s *Store) DeleteLinkToken(orgID string) error {
	if s.service == config.CloudGovService {
		return fmt.Errorf("%w: %s", errUnsupportedService, s.service)
	}
	_, err := s.clientv2.CloudMigrationServiceApi.DeleteLinkTokens(s.ctx, orgID).Execute()
	return err
}
