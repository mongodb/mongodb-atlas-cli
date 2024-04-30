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

package atlas

import (
	"fmt"

	"github.com/andreaangiolillo/mongocli-test/internal/config"
	atlasv2 "go.mongodb.org/atlas-sdk/v20231115002/admin"
)

//go:generate mockgen -destination=../../mocks/atlas/mock_database_roles.go -package=atlas github.com/andreaangiolillo/mongocli-test/internal/store/atlas DatabaseRoleLister

type DatabaseRoleLister interface {
	DatabaseRoles(string) ([]atlasv2.UserCustomDBRole, error)
}

// DatabaseRoles encapsulate the logic to manage different cloud providers.
func (s *Store) DatabaseRoles(projectID string) ([]atlasv2.UserCustomDBRole, error) {
	switch s.service {
	case config.CloudService, config.CloudGovService:
		result, _, err := s.clientv2.CustomDatabaseRolesApi.ListCustomDatabaseRoles(s.ctx, projectID).Execute()
		return result, err
	default:
		return nil, fmt.Errorf("%w: %s", errUnsupportedService, s.service)
	}
}
