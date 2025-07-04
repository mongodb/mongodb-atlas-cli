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
	atlasv2 "go.mongodb.org/atlas-sdk/v20250312005/admin"
)

// CreateUser encapsulates the logic to manage different cloud providers.
func (s *Store) CreateUser(user *atlasv2.CloudAppUser) (*atlasv2.CloudAppUser, error) {
	result, _, err := s.clientv2.MongoDBCloudUsersApi.CreateUser(s.ctx, user).Execute()
	return result, err
}

// UserByID encapsulates the logic to manage different cloud providers.
func (s *Store) UserByID(userID string) (*atlasv2.CloudAppUser, error) {
	result, _, err := s.clientv2.MongoDBCloudUsersApi.GetUser(s.ctx, userID).Execute()
	return result, err
}

// UserByName encapsulates the logic to manage different cloud providers.
func (s *Store) UserByName(username string) (*atlasv2.CloudAppUser, error) {
	result, _, err := s.clientv2.MongoDBCloudUsersApi.GetUserByUsername(s.ctx, username).Execute()
	return result, err
}

// OrganizationUsers encapsulates the logic to manage different cloud providers.
func (s *Store) OrganizationUsers(organizationID string, opts *ListOptions) (*atlasv2.PaginatedOrgUser, error) {
	res := s.clientv2.MongoDBCloudUsersApi.ListOrganizationUsers(s.ctx, organizationID)
	if opts != nil {
		res = res.ItemsPerPage(opts.ItemsPerPage).PageNum(opts.PageNum).IncludeCount(opts.IncludeCount)
	}
	result, _, err := res.Execute()
	return result, err
}

// TeamUsers encapsulates the logic to manage different cloud providers.
func (s *Store) TeamUsers(orgID, teamID string) (*atlasv2.PaginatedOrgUser, error) {
	result, _, err := s.clientv2.MongoDBCloudUsersApi.ListTeamUsers(s.ctx, orgID, teamID).Execute()
	return result, err
}
