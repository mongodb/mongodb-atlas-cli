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
	atlasv2 "go.mongodb.org/atlas-sdk/v20231115007/admin"
)

//go:generate mockgen -destination=../../mocks/atlas/mock_access_role.go -package=atlas github.com/mongodb/mongodb-atlas-cli/internal/store/atlas CloudProviderAccessRoleLister

type CloudProviderAccessRoleLister interface {
	CloudProviderAccessRoles(string) (*atlasv2.CloudProviderAccessRoles, error)
}

// CloudProviderAccessRoles encapsulates the logic to manage different cloud providers.
func (s *Store) CloudProviderAccessRoles(groupID string) (*atlasv2.CloudProviderAccessRoles, error) {
	result, _, err := s.clientv2.CloudProviderAccessApi.ListCloudProviderAccessRoles(s.ctx, groupID).Execute()
	return result, err
}
