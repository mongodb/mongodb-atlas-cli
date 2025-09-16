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
	atlasv2 "go.mongodb.org/atlas-sdk/v20250312007/admin"
)

// CreateCloudProviderAccessRole encapsulates the logic to manage different cloud providers.
func (s *Store) CreateCloudProviderAccessRole(groupID, provider string) (*atlasv2.CloudProviderAccessRole, error) {
	req := atlasv2.CloudProviderAccessRoleRequest{
		ProviderName: provider,
	}
	result, _, err := s.clientv2.CloudProviderAccessApi.CreateCloudProviderAccessRole(s.ctx, groupID, &req).Execute()
	return result, err
}

// CloudProviderAccessRoles encapsulates the logic to manage different cloud providers.
func (s *Store) CloudProviderAccessRoles(groupID string) (*atlasv2.CloudProviderAccessRoles, error) {
	result, _, err := s.clientv2.CloudProviderAccessApi.ListCloudProviderAccessRoles(s.ctx, groupID).Execute()
	return result, err
}

// DeauthorizeCloudProviderAccessRoles encapsulates the logic to manage different cloud providers.
func (s *Store) DeauthorizeCloudProviderAccessRoles(groupID string, cloudProvider string, roleID string) error {
	_, err := s.clientv2.CloudProviderAccessApi.DeauthorizeCloudProviderAccessRole(s.ctx, groupID, cloudProvider, roleID).Execute()
	return err
}

// AuthorizeCloudProviderAccessRole encapsulates the logic to manage different cloud providers.
func (s *Store) AuthorizeCloudProviderAccessRole(groupID, roleID string, req *atlasv2.CloudProviderAccessRoleRequestUpdate) (*atlasv2.CloudProviderAccessRole, error) {
	result, _, err := s.clientv2.CloudProviderAccessApi.AuthorizeCloudProviderAccessRole(s.ctx, groupID, roleID, req).Execute()
	return result, err
}
