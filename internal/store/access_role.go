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

	"github.com/mongodb/mongodb-atlas-cli/internal/config"
	atlasv2 "go.mongodb.org/atlas-sdk/admin"
	atlas "go.mongodb.org/atlas/mongodbatlas"
)

//go:generate mockgen -destination=../mocks/mock_access_role.go -package=mocks github.com/mongodb/mongodb-atlas-cli/internal/store CloudProviderAccessRoleCreator,CloudProviderAccessRoleAuthorizer,CloudProviderAccessRoleLister,CloudProviderAccessRoleDeauthorizer

type CloudProviderAccessRoleCreator interface {
	CreateCloudProviderAccessRole(string, string) (*atlasv2.CloudProviderAccessRole, error)
}

type CloudProviderAccessRoleLister interface {
	CloudProviderAccessRoles(string) (*atlasv2.CloudProviderAccessRoles, error)
}

type CloudProviderAccessRoleDeauthorizer interface {
	DeauthorizeCloudProviderAccessRoles(*atlas.CloudProviderDeauthorizationRequest) error
}

// CreateCloudProviderAccessRole encapsulates the logic to manage different cloud providers.
type CloudProviderAccessRoleAuthorizer interface {
	AuthorizeCloudProviderAccessRole(string, string, *atlas.CloudProviderAuthorizationRequest) (*atlasv2.CloudProviderAccessRole, error)
}

// CreateCloudProviderAccessRole encapsulates the logic to manage different cloud providers.
func (s *Store) CreateCloudProviderAccessRole(groupID, provider string) (*atlasv2.CloudProviderAccessRole, error) {
	switch s.service {
	case config.CloudService, config.CloudGovService:
		req := atlasv2.CloudProviderAccessRole{
			CloudProviderAccessAWSIAMRole: &atlasv2.CloudProviderAccessAWSIAMRole{
				ProviderName: provider,
			},
		}
		result, _, err := s.clientv2.CloudProviderAccessApi.CreateCloudProviderAccessRole(s.ctx, groupID, &req).Execute()
		return result, err
	default:
		return nil, fmt.Errorf("%w: %s", errUnsupportedService, s.service)
	}
}

// CloudProviderAccessRoles encapsulates the logic to manage different cloud providers.
func (s *Store) CloudProviderAccessRoles(groupID string) (*atlasv2.CloudProviderAccessRoles, error) {
	switch s.service {
	case config.CloudService, config.CloudGovService:
		result, _, err := s.clientv2.CloudProviderAccessApi.ListCloudProviderAccessRoles(s.ctx, groupID).Execute()
		return result, err
	default:
		return nil, fmt.Errorf("%w: %s", errUnsupportedService, s.service)
	}
}

// DeauthorizeCloudProviderAccessRoles encapsulates the logic to manage different cloud providers.
func (s *Store) DeauthorizeCloudProviderAccessRoles(req *atlas.CloudProviderDeauthorizationRequest) error {
	switch s.service {
	case config.CloudService, config.CloudGovService:
		_, err := s.clientv2.CloudProviderAccessApi.DeauthorizeCloudProviderAccessRole(s.ctx, req.GroupID, req.ProviderName, req.RoleID).Execute()
		return err
	default:
		return fmt.Errorf("%w: %s", errUnsupportedService, s.service)
	}
}

// AuthorizeCloudProviderAccessRole encapsulates the logic to manage different cloud providers.
func (s *Store) AuthorizeCloudProviderAccessRole(groupID, roleID string, req *atlas.CloudProviderAuthorizationRequest) (*atlasv2.CloudProviderAccessRole, error) {
	switch s.service {
	case config.CloudService, config.CloudGovService:
		role := atlasv2.CloudProviderAccessRole{
			CloudProviderAccessAWSIAMRole: &atlasv2.CloudProviderAccessAWSIAMRole{
				ProviderName:      req.ProviderName,
				IamAssumedRoleArn: &req.IAMAssumedRoleARN,
			},
		}
		result, _, err := s.clientv2.CloudProviderAccessApi.AuthorizeCloudProviderAccessRole(s.ctx, groupID, roleID, &role).Execute()
		return result, err
	default:
		return nil, fmt.Errorf("%w: %s", errUnsupportedService, s.service)
	}
}
